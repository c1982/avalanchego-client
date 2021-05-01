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

func (c Calls) PlatformExportKey(username string, password string, address string) (privateKey string, err error) {
	input := P{
		"username": username,
		"password": password,
		"address":  address,
	}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.exportKey", input, &output)
	if err != nil {
		return privateKey, err
	}
	err = output.OutStr("privateKey", &privateKey).Error()
	return privateKey, err
}

func (c Calls) PlatformGetBalance(address string) (balance GetBalanceResponse, err error) {
	input := P{
		"address": address,
	}

	err = c.client.NewRequest("ext/P", "platform.getBalance", input, &balance)
	if err != nil {
		return balance, err
	}

	return balance, err
}

func (c Calls) PlatformGetBlockchains() (blockchains GetBlockchainsResponse, err error) {
	input := P{}

	err = c.client.NewRequest("ext/P", "platform.getBlockchains", input, &blockchains)
	if err != nil {
		return blockchains, err
	}

	return blockchains, err
}

func (c Calls) PlatformGetBlockchainStatus(blockchainID string) (status string, err error) {
	input := P{
		"blockchainID": blockchainID,
	}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.getBlockchainStatus", input, &output)
	if err != nil {
		return status, err
	}
	err = output.OutStr("status", &status).Error()
	return status, err
}

func (c Calls) PlatformGetCurrentSupply() (supply string, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.getCurrentSupply", input, &output)
	if err != nil {
		return supply, err
	}
	err = output.OutStr("supply", &supply).Error()
	return supply, err
}

func (c Calls) PlatformGetCurrentValidators() (validators GetCurrentValidatorsResponse, err error) {
	input := P{}

	err = c.client.NewRequest("ext/P", "platform.getCurrentValidators", input, &validators)
	if err != nil {
		return validators, err
	}

	return validators, err
}

func (c Calls) PlatformGetHeight() (height int, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.getHeight", input, &output)
	if err != nil {
		return height, err
	}
	err = output.Error()
	return height, err
}

func (c Calls) PlatformGetMinStake() (minValidatorStake uint64, minDelegatorStake uint64, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.getMinStake", input, &output)
	if err != nil {
		return minValidatorStake, minDelegatorStake, err
	}
	err = output.Error()
	return minValidatorStake, minDelegatorStake, err
}

func (c Calls) PlatformGetPendingValidators() (pendinvalidators GetPendingValidatorsResponse, err error) {
	input := P{}

	err = c.client.NewRequest("ext/P", "platform.getPendingValidators", input, &pendinvalidators)
	if err != nil {
		return pendinvalidators, err
	}

	return pendinvalidators, err
}

func (c Calls) PlatformGetStakingAssetID(subnetID string) (assetID string, err error) {
	input := P{
		"subnetID": subnetID,
	}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.getStakingAssetID", input, &output)
	if err != nil {
		return assetID, err
	}
	err = output.OutStr("assetID", &assetID).Error()
	return assetID, err
}

func (c Calls) PlatformGetSubnets(ids []string) (subnets GetSubnetsResponse, err error) {
	input := P{
		"ids": ids,
	}

	err = c.client.NewRequest("ext/P", "platform.getSubnets", input, &subnets)
	if err != nil {
		return subnets, err
	}

	return subnets, err
}

func (c Calls) PlatformGetStake(addresses GetStakeArgs) (staked int, err error) {
	input := addresses
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.getStake", input, &output)
	if err != nil {
		return staked, err
	}
	err = output.Error()
	return staked, err
}

func (c Calls) PlatformGetTotalStake() (stake int, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.getTotalStake", input, &output)
	if err != nil {
		return stake, err
	}
	err = output.Error()
	return stake, err
}

func (c Calls) PlatformGetTx(txID string, encoding string) (tx string, encode string, err error) {
	input := P{
		"txID":     txID,
		"encoding": encoding,
	}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.getTx", input, &output)
	if err != nil {
		return tx, encode, err
	}
	err = output.OutStr("tx", &tx).OutStr("encode", &encode).Error()
	return tx, encode, err
}

func (c Calls) PlatformGetTxStatus(txID string) (status string, err error) {
	input := P{
		"txID": txID,
	}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.getTxStatus", input, &output)
	if err != nil {
		return status, err
	}
	err = output.OutStr("status", &status).Error()
	return status, err
}

func (c Calls) PlatformGetUTXOs(addresses PlatformGetUTXOsArgs) (utxos PlatformGetUTXOsResponse, err error) {
	input := P{
		"addresses": addresses,
	}
	output := P{}
	err = c.client.NewRequest("ext/bc/P", "platform.getUTXOs", input, &output)
	if err != nil {
		return utxos, err
	}
	err = output.Error()
	return utxos, err
}

func (c Calls) PlatformImportAVAX(imports PlatformImportAVAXArgs) (txID string, changeAddr string, err error) {
	input := imports
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.importAVAX", input, &output)
	if err != nil {
		return txID, changeAddr, err
	}
	err = output.OutStr("txID", &txID).OutStr("changeAddr", &changeAddr).Error()
	return txID, changeAddr, err
}

func (c Calls) PlatformImportKey(username string, password string, privateKey string) (address string, err error) {
	input := P{
		"username":   username,
		"password":   password,
		"privateKey": privateKey,
	}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.importKey", input, &output)
	if err != nil {
		return address, err
	}
	err = output.OutStr("address", &address).Error()
	return address, err
}

func (c Calls) PlatformIssueTx(tx string, encoding string) (txID string, err error) {
	input := P{
		"tx":       tx,
		"encoding": encoding,
	}
	output := P{}
	err = c.client.NewRequest("ext/P", "platform.issueTx", input, &output)
	if err != nil {
		return txID, err
	}
	err = output.OutStr("txID", &txID).Error()
	return txID, err
}

func (c Calls) PlatformListAddresses(username string, password string) (addresses []string, err error) {
	input := P{
		"username": username,
		"password": password,
	}

	err = c.client.NewRequest("ext/bc/P", "platform.listAddresses", input, &addresses)
	if err != nil {
		return addresses, err
	}

	return addresses, err
}

func (c Calls) PlatformSampleValidators(size int, subnetID string) (validators []string, err error) {
	input := P{
		"size":     size,
		"subnetID": subnetID,
	}

	err = c.client.NewRequest("ext/P", "platform.sampleValidators", input, &validators)
	if err != nil {
		return validators, err
	}

	return validators, err
}

func (c Calls) PlatformValidatedBy(subnetID string) (blockchainIDs []string, err error) {
	input := P{
		"subnetID": subnetID,
	}

	err = c.client.NewRequest("ext/P", "platform.validatedBy", input, &blockchainIDs)
	if err != nil {
		return blockchainIDs, err
	}

	return blockchainIDs, err
}
