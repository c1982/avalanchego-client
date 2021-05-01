// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) AvaxExport(to string, amount int64, assetid string, username string, password string) (txID string, err error) {
	input := P{
		"to":       to,
		"amount":   amount,
		"assetid":  assetid,
		"username": username,
		"password": password,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/C/avax", "avax.export", input, &output)
	if err != nil {
		return txID, err
	}
	err = output.OutStr("txID", &txID).Error()
	return txID, err
}

func (c Calls) AvaxExportKey(username string, password string, address string) (privateKey string, privateKeyHex string, err error) {
	input := P{
		"username": username,
		"password": password,
		"address":  address,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/C/avax", "avax.exportKey", input, &output)
	if err != nil {
		return privateKey, privateKeyHex, err
	}
	err = output.OutStr("privateKey", &privateKey).OutStr("privateKeyHex", &privateKeyHex).Error()
	return privateKey, privateKeyHex, err
}

func (c Calls) AvaxGetUTXOs(UTXOs AVAXGetUTXOsArgs) (utxos AVAXGetUTXOsResponse, err error) {
	input := UTXOs

	err = c.client.NewRequest("ext/bc/C/avax", "avax.getUTXOs", input, &utxos)
	if err != nil {
		return utxos, err
	}

	return utxos, err
}

func (c Calls) AvaxImport(to string, sourceChain string, username string, password string) (txID string, err error) {
	input := P{
		"to":          to,
		"sourceChain": sourceChain,
		"username":    username,
		"password":    password,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/C/avax", "avax.import", input, &output)
	if err != nil {
		return txID, err
	}
	err = output.OutStr("txID", &txID).Error()
	return txID, err
}

func (c Calls) AvaxImportKey(username string, password string, privateKey string) (address string, err error) {
	input := P{
		"username":   username,
		"password":   password,
		"privateKey": privateKey,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/C/avax", "avax.importKey", input, &output)
	if err != nil {
		return address, err
	}
	err = output.OutStr("address", &address).Error()
	return address, err
}
