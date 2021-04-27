package api

import (
	"avalanchego-client/pkg/network"
	"encoding/json"
)

type AvmMintArgs struct {
	Amount     int64    `json:"amount"`
	Assetid    string   `json:"assetID"`
	To         string   `json:"to"`
	From       []string `json:"from"`
	Changeaddr string   `json:"changeAddr"`
	Username   string   `json:"username"`
	Password   string   `json:"password"`
}

//AVMBuildGenesis Given a JSON representation of this Virtual Machine’s genesis state, create the byte representation of that state.
func (c Calls) AVMBuildGenesis(networkID int, data json.RawMessage, encodetype string) (bytes, encoding string, err error) {
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

//AVMCreateAddress Create a new address controlled by the given user.
func (c Calls) AVMCreateAddress(username, password string) (address string, err error) {
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

//AVMCreateFixedCapAsset Create a new fixed-cap, fungible asset. A quantity of it is created at initialization and then no more is ever created. The asset can be sent with avm.send.
func (c Calls) AVMCreateFixedCapAsset(asset FixedCapAssetArgs) (assetID, changeAddr string, err error) {
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

//AVMMint Mint units of a variable-cap asset
func (c Calls) AVMMint(mint AvmMintArgs) (txID, changeAddr string, err error) {
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

//AVMCreateVariableCapAsset Create a new variable-cap, fungible asset. No units of the asset exist at initialization. Minters can mint units of this asset using avm.mint.
func (c Calls) AVMCreateVariableCapAsset(asset CreateVariableCapAssetArgs) (assetID, changeAddr string, err error) {
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

//AVMExport Send a non-AVAX from the X-Chain to the P-Chain or C-Chain. After calling this method, you must call avax.import on the C-Chain to complete the transfer.
func (c Calls) AVMExport(export ExportArgs) (txID, changeAddr string, err error) {
	rsp, err := c.client.NewRequestStruct("ext/bc/X", "avm.export", export)
	if err != nil {
		return txID, changeAddr, err
	}

	err = rsp.
		OutStr("txID", &txID).
		OutStr("changeAddr", &changeAddr).
		Error()

	return txID, changeAddr, err
}

//AVMExportAVAX Send AVAX from the X-Chain to another chain. After calling this method, you must call import on the other chain to complete the transfer.
func (c Calls) AVMExportAVAX(args ExportAVAXArgs) (txID, changeAddr string, err error) {
	rsp, err := c.client.NewRequestStruct("ext/bc/X", "avm.exportAVAX", args)
	if err != nil {
		return txID, changeAddr, err
	}

	err = rsp.
		OutStr("txID", &txID).
		OutStr("changeAddr", &changeAddr).
		Error()

	return txID, changeAddr, err
}

//AVMExportKey Get the private key that controls a given address. The returned private key can be added to a user with avm.importKey.
func (c Calls) AVMExportKey(username, password, address string) (privateKey string, err error) {
	rsp := network.P{}
	err = c.client.NewRequestFor(&rsp, "ext/bc/X", "avm.exportKey", network.P{
		"username": username,
		"password": password,
		"address":  address,
	})
	if err != nil {
		return privateKey, err
	}

	err = rsp.
		OutStr("privateKey", &privateKey).
		Error()

	return privateKey, err
}
