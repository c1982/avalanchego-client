package api_test

import (
	"avalanchego-client/pkg/api"
	"avalanchego-client/pkg/network"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func NewServer(responsebody string) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, responsebody)
	}))
}

func TestBootStrap(t *testing.T) {
	ts := NewServer(`{"jsonrpc":"2.0","result":{"isBootstrapped":false},"id":1}`)
	defer ts.Close()

	client, err := network.NewNodeClient(ts.URL, 1)
	if err != nil {
		t.Error(err)
	}

	api := api.NewAPICalls(client)
	ok, err := api.IsBootstrapped()
	if err != nil {
		t.Error(err)
	}

	if ok {
		t.Error("IsBootstrapped error: expect: false, got: true")
	}
}
