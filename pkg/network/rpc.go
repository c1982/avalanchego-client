package network

import (
	"encoding/json"
	"fmt"

	"github.com/ybbus/jsonrpc"
)

type rpcClient struct {
	endpoint  string
	networkID int
}

func (rpc *rpcClient) NewRequest(basepath, rpcmethod string, params interface{}) (response []byte, err error) {
	endpoint := fmt.Sprintf("%s/%s", rpc.endpoint, basepath)
	rsp, err := jsonrpc.NewClient(endpoint).Call(rpcmethod, params)
	if err != nil {
		return response, err
	}

	if rsp.Error != nil {
		return response, fmt.Errorf("rpcClient returned error: %d, %s", rsp.Error.Code, rsp.Error.Message)
	}

	response, err = json.Marshal(rsp.Result)
	if err != nil {
		return response, fmt.Errorf("rpcClient returned invalid JSON object: %v", rsp.Result)
	}

	return response, nil
}

func (rpc *rpcClient) NewRequestFor(out interface{}, basepath, rpcmethod string, params P) error {
	rsp, err := rpc.NewRequest(basepath, rpcmethod, params)
	if err != nil {
		return err
	}

	err = json.Unmarshal(rsp, out)
	if err != nil {
		return err
	}

	return nil
}

func (rpc *rpcClient) NewRequestStruct(basepath, rpcmethod string, params interface{}) (results P, err error) {
	rsp, err := rpc.NewRequest(basepath, rpcmethod, params)
	if err != nil {
		return results, err
	}

	results = P{}
	err = json.Unmarshal(rsp, &results)
	if err != nil {
		return results, err
	}

	return results, nil
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
