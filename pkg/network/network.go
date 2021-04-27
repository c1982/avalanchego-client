package network

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
