package api_test

import (
	"avalanchego-client/pkg/api"
	"avalanchego-client/pkg/network"
	"fmt"
	"testing"
)

func TestAvmNFT(t *testing.T) {
	calltests := []struct {
		name     string
		response string
		fn       func(api.Calls) error
	}{
		{
			name: "CreateNFTAsset",
			response: `{
				"jsonrpc": "2.0",
				"result": {
					"assetID": "assetID",
					"changeAddr": "changeAddr"
				},
				"id": 1
			}`,
			fn: func(c api.Calls) error {
				assetID, changeAddr, err := c.CreateNFTAsset(api.NFTAsset{})
				if err != nil {
					return err
				}

				if assetID != "assetID" || changeAddr != "changeAddr" {
					return fmt.Errorf("invald return values: assetID: %s, changeAddr: %s", assetID, changeAddr)
				}

				return nil
			},
		},
		{
			name: "MintNFT",
			response: `{
				"jsonrpc":"2.0",
				"id"     :1,
				"result" :{
					"txID":"2oGdPdfw2qcNUHeqjw8sU2hPVrFyNUTgn6A8HenDra7oLCDtja",
					"changeAddr": "X-avax1turszjwn05lflpewurw96rfrd3h6x8flgs5uf8"
				}
			}`,
			fn: func(c api.Calls) error {
				txID, changeAddr, err := c.MintNFT(api.MintNFT{})
				if err != nil {
					return err
				}

				if txID != "2oGdPdfw2qcNUHeqjw8sU2hPVrFyNUTgn6A8HenDra7oLCDtja" || changeAddr != "X-avax1turszjwn05lflpewurw96rfrd3h6x8flgs5uf8" {
					return fmt.Errorf("invald return values: assetID: %s, changeAddr: %s", txID, changeAddr)
				}

				return nil
			},
		},
	}

	for _, tt := range calltests {
		_ = t.Run(tt.name, func(t *testing.T) {
			server := NewServer(tt.response)
			defer server.Close()

			client, err := network.NewNodeClient(server.URL, 1)
			if err != nil {
				t.Error(err)
			}

			api := api.NewAPICalls(client)
			err = tt.fn(api)
			if err != nil {
				t.Error(err)
			}
		})
	}
}
