//go:generate go run gen.go

package api

import (
	"avalanchego-client/pkg/network"
)

//Calls command struct
type Calls struct {
	client network.NodeClient
}

//NewAPICalls access to avalanche node functions
func NewAPICalls(client network.NodeClient) Calls {
	return Calls{client: client}
}
