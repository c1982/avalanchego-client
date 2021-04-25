package api

import (
	"avalanchego-client/pkg/network"
	"encoding/json"
)

//BuildGenesis Given a JSON representation of this Virtual Machine’s genesis state, create the byte representation of that state.
func (c Calls) BuildGenesis(networkID int, data json.RawMessage, encodetype string) (bytes, encoding string, err error) {
	jsondata, err := json.Marshal(data)
	if err != nil {
		return bytes, encodetype, err
	}

	out := network.P{}
	err = c.client.NewRequestFor(&out, "ext/vm/avm", "avm.buildGenesis",
		network.P{
			"networkID":   networkID,
			"genesisData": string(jsondata),
			"encoding":    encodetype,
		})
	if err != nil {
		return bytes, encodetype, err
	}

	bytes, err = out.GetString("bytes")
	if err != nil {
		return bytes, encodetype, err
	}

	encoding, err = out.GetString("encoding")
	if err != nil {
		return bytes, encodetype, err
	}

	return bytes, encoding, err
}

//CreateAddress Create a new address controlled by the given user.
func (c Calls) CreateAddress(username, password string) (address string, err error) {
	out := network.P{}
	err = c.client.NewRequestFor(&out, "ext/bc/X", "avm.createAddress",
		network.P{
			"username": username,
			"password": password,
		})
	if err != nil {
		return address, err
	}

	address, err = out.GetString("address")
	if err != nil {
		return address, err
	}

	return address, nil
}

//CreateFixedCapAsset Create a new fixed-cap, fungible asset. A quantity of it is created at initialization and then no more is ever created. The asset can be sent with avm.send.
func (c Calls) CreateFixedCapAsset(asset FixedCapAsset) (assetID, changeAddr string, err error) {
	out, err := c.client.NewRequestStruct("ext/bc/X", "avm.createFixedCapAsset", asset)
	if err != nil {
		return assetID, changeAddr, err
	}

	assetID, err = out.GetString("assetID")
	if err != nil {
		return assetID, changeAddr, err
	}

	assetID, err = out.GetString("changeAddr")
	if err != nil {
		return assetID, changeAddr, err
	}

	return assetID, changeAddr, nil
}

//Mint Mint units of a variable-cap asset
func (c Calls) Mint(mint AvmMint) (txID, changeAddr string, err error) {
	out, err := c.client.NewRequestStruct("ext/bc/X", "avm.mint", mint)
	if err != nil {
		return txID, changeAddr, err
	}

	txID, err = out.GetString("txID")
	if err != nil {
		return txID, changeAddr, err
	}

	changeAddr, err = out.GetString("changeAddr")
	if err != nil {
		return txID, changeAddr, err
	}

	return txID, changeAddr, nil
}

//CreateVariableCapAsset Create a new variable-cap, fungible asset. No units of the asset exist at initialization. Minters can mint units of this asset using avm.mint.
func (c Calls) CreateVariableCapAsset(asset CreateVariableCapAsset) (assetID, changeAddr string, err error) {
	out, err := c.client.NewRequestStruct("ext/bc/X", "amv.createVariableCapAsset", asset)
	if err != nil {
		return assetID, changeAddr, err
	}

	assetID, err = out.GetString("assetID")
	if err != nil {
		return assetID, changeAddr, err
	}

	changeAddr, err = out.GetString("changeAddr")
	if err != nil {
		return assetID, changeAddr, err
	}

	return assetID, changeAddr, nil
}

//CreateNFTAsset Create a new non-fungible asset. No units of the asset exist at initialization. Minters can mint units of this asset using avm.mintNFT
func (c Calls) CreateNFTAsset(asset NFTAsset) (assetID, changeAddr string, err error) {
	out, err := c.client.NewRequestStruct("ext/bc/X", "amv.createVariableCapAsset", asset)
	if err != nil {
		return assetID, changeAddr, err
	}

	assetID, err = out.GetString("assetID")
	if err != nil {
		return assetID, changeAddr, err
	}

	changeAddr, err = out.GetString("changeAddr")
	if err != nil {
		return assetID, changeAddr, err
	}

	return assetID, changeAddr, nil
}
