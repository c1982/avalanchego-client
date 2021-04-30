// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) AvmCreateAddress(username string, password string) (address string, err error) {
	input := P{
		"username": username,
		"password": password,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.createAddress", input, &output)
	if err != nil {
		return address, err
	}
	err = output.OutStr("address", &address).Error()
	return address, err
}

func (c Calls) AvmCreateFixedCapAsset(asset FixedCapAssetArgs) (assetID string, changeAddr string, err error) {
	input := asset
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.createFixedCapAsset", input, &output)
	if err != nil {
		return assetID, changeAddr, err
	}
	err = output.OutStr("assetID", &assetID).OutStr("changeAddr", &changeAddr).Error()
	return assetID, changeAddr, err
}

func (c Calls) AvmMint(mint AvmMintArgs) (txID string, changeAddr string, err error) {
	input := mint
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.mint", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) AvmCreateVariableCapAsset(asset CreateVariableCapAssetArgs) (assetID string, changeAddr string, err error) {
	input := asset
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.createVariableCapAsset", input, &output)
	if err != nil {
		return assetID, changeAddr, err
	}
	err = output.OutStr("assetID", &assetID).OutStr("changeAddr", &changeAddr).Error()
	return assetID, changeAddr, err
}

func (c Calls) AvmCreateNFTAsset(asset NFTAssetArgs) (assetID string, changeAddr string, err error) {
	input := asset
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.createNFTAsset", input, &output)
	if err != nil {
		return assetID, changeAddr, err
	}
	err = output.OutStr("assetID", &assetID).OutStr("changeAddr", &changeAddr).Error()
	return assetID, changeAddr, err
}

func (c Calls) AvmMintNFT(mint MintNFTArgs) (txID string, changeAddr string, err error) {
	input := mint
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.mintNFT", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) AvmExport(mint ExportArgs) (txID string, changeAddr string, err error) {
	input := mint
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.export", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) AvmExportAVAX(mint ExportAVAXArgs) (txID string, changeAddr string, err error) {
	input := mint
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.exportAVAX", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) AvmExportKey(username string, password string, address string) (privateKey string, err error) {
	input := P{
		"username": username,
		"password": password,
		"address":  address,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.exportKey", input, &output)
	if err != nil {
		return privateKey, err
	}
	err = output.OutStr("privateKey", &privateKey).Error()
	return privateKey, err
}

func (c Calls) AvmGetAllBalances(address string) (balances []BalancesResponse, err error) {
	input := P{
		"address": address,
	}

	err = c.client.NewRequest("ext/bc/X", "avm.getAllBalances", input, &balances)
	if err != nil {
		return balances, err
	}

	return balances, err
}

func (c Calls) AvmGetAssetDescription(assetid string) (assetID string, name string, symbol string, denomination string, err error) {
	input := P{
		"assetid": assetid,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.getAssetDescription", input, &output)
	if err != nil {
		return assetID, name, symbol, denomination, err
	}
	err = output.OutStr("assetID", &assetID).OutStr("name", &name).OutStr("symbol", &symbol).OutStr("denomination", &denomination).Error()
	return assetID, name, symbol, denomination, err
}

func (c Calls) AvmGetBalance(address string, assetID string) (balance BalanceResponse, err error) {
	input := P{
		"address": address,
		"assetID": assetID,
	}

	err = c.client.NewRequest("ext/bc/X", "avm.getBalance", input, &balance)
	if err != nil {
		return balance, err
	}

	return balance, err
}

func (c Calls) AvmGetTx(txid string, encode string) (tx string, encoding string, err error) {
	input := P{
		"txid":   txid,
		"encode": encode,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.getTx", input, &output)
	if err != nil {
		return tx, encoding, err
	}
	err = output.OutStr("tx", &tx).OutStr("encoding", &encoding).Error()
	return tx, encoding, err
}

func (c Calls) AvmGetTxStatus(txID string) (status string, err error) {
	input := P{
		"txID": txID,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.getTxStatus", input, &output)
	if err != nil {
		return status, err
	}
	err = output.OutStr("status", &status).Error()
	return status, err
}

func (c Calls) AvmGetUTXOs(utxo GetUTXOsArgs) (utxorsp GetUTXOsResponse, err error) {
	input := utxo

	err = c.client.NewRequest("ext/bc/X", "avm.getUTXOs", input, &utxorsp)
	if err != nil {
		return utxorsp, err
	}

	return utxorsp, err
}

func (c Calls) AvmImport(to string, sourceChain string, username string, password string) (txID string, err error) {
	input := P{
		"to":          to,
		"sourceChain": sourceChain,
		"username":    username,
		"password":    password,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.import", input, &output)
	if err != nil {
		return txID, err
	}
	err = output.OutStr("txID", &txID).Error()
	return txID, err
}

func (c Calls) AvmImportAVAX(to string, sourceChain string, username string, password string) (txID string, err error) {
	input := P{
		"to":          to,
		"sourceChain": sourceChain,
		"username":    username,
		"password":    password,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.importAVAX", input, &output)
	if err != nil {
		return txID, err
	}
	err = output.OutStr("txID", &txID).Error()
	return txID, err
}

func (c Calls) AvmImportKey(privateKey string, username string, password string) (address string, err error) {
	input := P{
		"privateKey": privateKey,
		"username":   username,
		"password":   password,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.importKey", input, &output)
	if err != nil {
		return address, err
	}
	err = output.OutStr("address", &address).Error()
	return address, err
}

func (c Calls) AvmIssueTx(tx string, encoding string) (txID string, err error) {
	input := P{
		"tx":       tx,
		"encoding": encoding,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.issueTx", input, &output)
	if err != nil {
		return txID, err
	}
	err = output.OutStr("txID", &txID).Error()
	return txID, err
}

func (c Calls) AvmListAddresses(username string, password string) (addresses ListAddressesResponse, err error) {
	input := P{
		"username": username,
		"password": password,
	}

	err = c.client.NewRequest("ext/bc/X", "avm.listAddresses", input, &addresses)
	if err != nil {
		return addresses, err
	}

	return addresses, err
}

func (c Calls) AvmSend(send SendArgs) (txID string, changeAddr string, err error) {
	input := send
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.send", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) AvmSendMultiple(sendmultiple SendMultipleArgs) (txID string, changeAddr string, err error) {
	input := sendmultiple
	output := P{}
	err = c.client.NewRequest("ext/bc/X", "avm.sendMultiple", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) AvmSendNFT(sendNFT SendNFTArgs) (txID string, changeAddr string, err error) {
	input := sendNFT

	err = c.client.NewRequest("ext/bc/X", "avm.sendNFT", input, &txID)
	if err != nil {
		return txID, changeAddr, err
	}

	return txID, changeAddr, err
}

func (c Calls) WalletIssueTx(tx string, encoding string) (txID string, err error) {
	input := P{
		"tx":       tx,
		"encoding": encoding,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/X/wallet", "wallet.issueTx", input, &output)
	if err != nil {
		return txID, err
	}
	err = output.OutStr("txID", &txID).Error()
	return txID, err
}

func (c Calls) WalletSend(send SendWalletArgs) (txID string, changeAddr string, err error) {
	input := send
	output := P{}
	err = c.client.NewRequest("ext/bc/X/wallet", "wallet.send", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) WalletSendMultiple(send SendMultipleArgs) (txID string, changeAddr string, err error) {
	input := send
	output := P{}
	err = c.client.NewRequest("ext/bc/X/wallet", "wallet.sendMultiple", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}
