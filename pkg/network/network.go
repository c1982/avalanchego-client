package network

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

//NodeClient network connection of node
type NodeClient interface {
	NewRequest(basepath, rpcmethod string, params P) ([]byte, error)
	NewRequestFor(out interface{}, basepath, rpcmethod string, params P) error
	GetEndpoint() string
	GetNetworkID() int
}

//NewNodeClient initialize network connection
func NewNodeClient(endpoint string, networkID int) (NodeClient, error) {
	return newRPCClient(endpoint, networkID)
}
