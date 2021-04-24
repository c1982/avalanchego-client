package api

import (
	"avalanchego-client/pkg/network"
	"fmt"
)

//IsBootstrapped to check whether a given chain is done bootstrapping
func (c Calls) IsBootstrapped() (bool, error) {
	out := network.P{}
	err := c.client.NewRequestFor(&out, "ext/info", "info.isBootstrapped", network.P{"chain": "X"})
	if err != nil {
		return false, err
	}

	val, isbool := (out["isBootstrapped"]).(bool)
	if !isbool {
		return false, fmt.Errorf("invalid type of key isBootstrapped")
	}

	return val, nil
}
