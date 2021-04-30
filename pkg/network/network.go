package network

//NodeClient network connection of node
type NodeClient interface {
	NewRequest(basepath, rpcmethod string, input interface{}, output interface{}) error
	GetEndpoint() string
	GetNetworkID() int
}

//NewNodeClient initialize network connection
func NewNodeClient(endpoint string, networkID int) (NodeClient, error) {
	return newRPCClient(endpoint, networkID)
}
