package network

import (
	"fmt"

	"github.com/ybbus/jsonrpc"
)

type rpcClient struct {
	endpoint  string
	networkID int
}

func (rpc *rpcClient) NewRequest(basepath, rpcmethod string, input interface{}, output interface{}) (err error) {
	endpoint := fmt.Sprintf("%s/%s", rpc.endpoint, basepath)
	rsp, err := jsonrpc.NewClient(endpoint).Call(rpcmethod, input)
	if err != nil {
		return err
	}

	if rsp.Error != nil {
		return fmt.Errorf("rpc client returned error: %d, %s", rsp.Error.Code, rsp.Error.Message)
	}

	err = rsp.GetObject(&output)
	if err != nil {
		return fmt.Errorf("rpc client returned invalid JSON object: %v", rsp.Result)
	}

	return nil
}

func (rpc *rpcClient) GetEndpoint() string {
	return rpc.endpoint
}

func (rpc *rpcClient) GetNetworkID() int {
	return rpc.networkID
}

func newRPCClient(endpoint string, networkID int) (*rpcClient, error) {
	return &rpcClient{
		endpoint:  endpoint,
		networkID: networkID,
	}, nil
}
