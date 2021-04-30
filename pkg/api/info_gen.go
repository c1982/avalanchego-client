// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) InfoGetBlockchainID(alias string) (blockchainID string, err error) {
	input := P{
		"alias": alias,
	}
	output := P{}
	err = c.client.NewRequest("ext/info", "info.getBlockchainID", input, &output)
	if err != nil {
		return blockchainID, err
	}
	err = output.OutStr("blockchainID", &blockchainID).Error()
	return blockchainID, err
}
