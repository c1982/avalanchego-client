package api

import (
	"avalanchego-client/pkg/network"
	"encoding/json"
)

type AvmMint struct {
	Amount     int64    `json:"amount"`
	Assetid    string   `json:"assetID"`
	To         string   `json:"to"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

//BuildGenesis Given a JSON representation of this Virtual Machine’s genesis state, create the byte representation of that state.
func (c Calls) BuildGenesis(networkID int, data json.RawMessage, encodetype string) (bytes, encoding string, err error) {
	jsondata, err := json.Marshal(data)
	if err != nil {
		return bytes, encodetype, err
	}

	rsp, err := c.client.NewRequestStruct("ext/vm/avm", "avm.buildGenesis",
		network.P{
			"networkID":   networkID,
			"genesisData": string(jsondata),
			"encoding":    encodetype,
		})
	if err != nil {
		return bytes, encodetype, err
	}

	err = rsp.
		OutStr("bytes", &bytes).
		OutStr("encoding", &encoding).
		Error()

	return bytes, encoding, err
}

//CreateAddress Create a new address controlled by the given user.
func (c Calls) CreateAddress(username, password string) (address string, err error) {
	rsp, err := c.client.NewRequestStruct("ext/bc/X", "avm.createAddress",
		network.P{
			"username": username,
			"password": password,
		})
	if err != nil {
		return address, err
	}

	err = rsp.OutStr("address", &address).Error()
	return address, err
}

//CreateFixedCapAsset Create a new fixed-cap, fungible asset. A quantity of it is created at initialization and then no more is ever created. The asset can be sent with avm.send.
func (c Calls) CreateFixedCapAsset(asset FixedCapAsset) (assetID, changeAddr string, err error) {
	rsp, err := c.client.NewRequestStruct("ext/bc/X", "avm.createFixedCapAsset", asset)
	if err != nil {
		return assetID, changeAddr, err
	}

	err = rsp.
		OutStr("assetID", &assetID).
		OutStr("changeAddr", &changeAddr).
		Error()

	return assetID, changeAddr, err
}

//Mint Mint units of a variable-cap asset
func (c Calls) Mint(mint AvmMint) (txID, changeAddr string, err error) {
	rsp, err := c.client.NewRequestStruct("ext/bc/X", "avm.mint", mint)
	if err != nil {
		return txID, changeAddr, err
	}

	err = rsp.
		OutStr("txID", &txID).
		OutStr("changeAddr", &changeAddr).
		Error()

	return txID, changeAddr, err
}

//CreateVariableCapAsset Create a new variable-cap, fungible asset. No units of the asset exist at initialization. Minters can mint units of this asset using avm.mint.
func (c Calls) CreateVariableCapAsset(asset CreateVariableCapAsset) (assetID, changeAddr string, err error) {
	rsp, err := c.client.NewRequestStruct("ext/bc/X", "amv.createVariableCapAsset", asset)
	if err != nil {
		return assetID, changeAddr, err
	}

	err = rsp.
		OutStr("assetID", &assetID).
		OutStr("changeAddr", &changeAddr).
		Error()

	return assetID, changeAddr, err
}
