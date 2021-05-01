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

func (c Calls) InfoGetNetworkID(alias string) (networkID int, err error) {
	input := P{
		"alias": alias,
	}
	output := P{}
	err = c.client.NewRequest("ext/info", "info.getNetworkID", input, &output)
	if err != nil {
		return networkID, err
	}
	err = output.Error()
	return networkID, err
}

func (c Calls) InfoGetNetworkName() (networkName string, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/info", "info.getNetworkName", input, &output)
	if err != nil {
		return networkName, err
	}
	err = output.OutStr("networkName", &networkName).Error()
	return networkName, err
}

func (c Calls) InfoGetNodeID() (nodeID string, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/info", "info.getNodeID", input, &output)
	if err != nil {
		return nodeID, err
	}
	err = output.OutStr("nodeID", &nodeID).Error()
	return nodeID, err
}

func (c Calls) InfoGetNodeIP() (ip string, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/info", "info.getNodeIP", input, &output)
	if err != nil {
		return ip, err
	}
	err = output.OutStr("ip", &ip).Error()
	return ip, err
}

func (c Calls) InfoGetNodeVersion() (version string, err error) {
	input := P{}
	output := P{}
	err = c.client.NewRequest("ext/info", "info.getNodeVersion", input, &output)
	if err != nil {
		return version, err
	}
	err = output.OutStr("version", &version).Error()
	return version, err
}

func (c Calls) InfoIsBootstrapped(chain string) (isBootstrapped bool, err error) {
	input := P{
		"chain": chain,
	}
	output := P{}
	err = c.client.NewRequest("ext/info", "info.isBootstrapped", input, &output)
	if err != nil {
		return isBootstrapped, err
	}
	err = output.OutBool("isBootstrapped", &isBootstrapped).Error()
	return isBootstrapped, err
}

func (c Calls) InfoPeers(nodeIDs []string) (peers InfoPeersArgs, err error) {
	input := P{
		"nodeIDs": nodeIDs,
	}

	err = c.client.NewRequest("ext/info", "info.peers", input, &peers)
	if err != nil {
		return peers, err
	}

	return peers, err
}

func (c Calls) InfoGetTxFee(nodeIDs []string) (creationTxFee uint64, txFee uint64, err error) {
	input := P{
		"nodeIDs": nodeIDs,
	}
	output := P{}
	err = c.client.NewRequest("ext/info", "info.getTxFee", input, &output)
	if err != nil {
		return creationTxFee, txFee, err
	}
	err = output.Error()
	return creationTxFee, txFee, err
}
