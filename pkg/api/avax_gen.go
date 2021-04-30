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

func (c Calls) AvaxExportKey(address string, username string, password string) (privateKey string, privateKeyHex string, err error) {
	input := P{
		"address":  address,
		"username": username,
		"password": password,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/C/avax", "avax.exportKey", input, &output)
	if err != nil {
		return privateKey, privateKeyHex, err
	}
	err = output.OutStr("privateKey", &privateKey).OutStr("privateKeyHex", &privateKeyHex).Error()
	return privateKey, privateKeyHex, err
}
