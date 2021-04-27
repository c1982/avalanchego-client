package api

import (
	"avalanchego-client/pkg/network"
)

//IsBootstrapped to check whether a given chain is done bootstrapping
func (c Calls) IsBootstrapped() (isbootstrapped bool, err error) {
	rsp, err := c.client.NewRequestStruct("ext/info", "info.isBootstrapped", network.P{"chain": "X"})
	if err != nil {
		return false, err
	}

	err = rsp.OutBool("isBootstrapped", &isbootstrapped).Error()
	return isbootstrapped, err
}
