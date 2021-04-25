package network

import "fmt"

//P node input/output data transfer object
type P map[string]interface{}

//Success looks for the "success" key.
func (p P) Success() bool {
	v, ok := (p["success"]).(bool)
	if !ok {
		return false
	}

	return v
}

//GetString get string value from map
func (p P) GetString(key string) (string, error) {
	v, exists := p[key]
	if !exists {
		return "", fmt.Errorf("key not found: %s", key)
	}

	str, cast := v.(string)
	if !cast {
		return "", fmt.Errorf("key (%s) cannot cast to string", key)
	}

	return str, nil
}

//NodeClient network connection of node
type NodeClient interface {
	NewRequest(basepath, rpcmethod string, params interface{}) ([]byte, error)
	NewRequestStruct(basepath, rpcmethod string, params interface{}) (P, error)
	NewRequestFor(out interface{}, basepath, rpcmethod string, params P) error
	GetEndpoint() string
	GetNetworkID() int
}

//NewNodeClient initialize network connection
func NewNodeClient(endpoint string, networkID int) (NodeClient, error) {
	return newRPCClient(endpoint, networkID)
}
