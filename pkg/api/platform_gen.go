// Code generated for package api by gen.go DO NOT EDIT. (@generated)
package api

func (c Calls) PlatformAddDelegator(delegator AddDelegatorArgs) (txID string, changeAddr string, err error) {
	input := delegator
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.addDelegator", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) PlatformAddValidator(validator AddValidatorArgs) (txID string, changeAddr string, err error) {
	input := validator
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.addValidator", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) PlatformAddSubnetValidator(subnetvalidator AddSubnetValidatorArgs) (txID string, changeAddr string, err error) {
	input := subnetvalidator
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.addSubnetValidator", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) PlatformCreateAddress(username string, password string) (address string, err error) {
	input := P{
		"username": username,
		"password": password,
	}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.createAddress", input, &output)
	if err != nil {
		return address, err
	}
	err = output.OutStr("address", &address).Error()
	return address, err
}

func (c Calls) PlatformCreateBlockchain(blockchain CreateBlockchainArgs) (txID string, changeAddr string, err error) {
	input := blockchain
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.createBlockchain", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) PlatformCreateSubnet(subnet CreateSubnetArgs) (txID string, err error) {
	input := subnet
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.createSubnet", input, &output)
	if err != nil {
		return txID, err
	}
	err = output.OutStr("txID", &txID).Error()
	return txID, err
}

func (c Calls) PlatformExportAVAX(avax PlatformExportAVAXArgs) (txID string, changeAddr string, err error) {
	input := avax
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.exportAVAX", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}
