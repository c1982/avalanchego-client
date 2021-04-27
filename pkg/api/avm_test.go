package api_test

import (
	"avalanchego-client/pkg/api"
	"avalanchego-client/pkg/network"
	"fmt"
	"testing"
)

func TestAvmGeneral(t *testing.T) {
	calltests := []struct {
		name     string
		response string
		fn       func(api.Calls) error
	}{
		{
			name: "AVMExport",
			response: `{
				"jsonrpc": "2.0",
				"result": {
					"txID": "2Eu16yNaepP57XrrJgjKGpiEDandpiGWW8xbUm6wcTYny3fejj",
					"changeAddr": "X-avax1turszjwn05lflpewurw96rfrd3h6x8flgs5uf8"
				},
				"id": 1
			}`,
			fn: func(c api.Calls) error {
				txtID, changeAddr, err := c.AVMExport(api.ExportArgs{})
				if err != nil {
					return nil
				}

				if txtID != "2Eu16yNaepP57XrrJgjKGpiEDandpiGWW8xbUm6wcTYny3fejj" || changeAddr != "X-avax1turszjwn05lflpewurw96rfrd3h6x8flgs5uf8" {
					return fmt.Errorf("invald return values: txtID: %s, changeAddr: %s", txtID, changeAddr)
				}

				return nil
			},
		},
		{
			name: "AVMExportAVAX",
			response: `{
				"jsonrpc": "2.0",
				"result": {
					"txID": "25VzbNzt3gi2vkE3Kr6H9KJeSR2tXkr8FsBCm3vARnB5foLVmx",
					"changeAddr": "X-avax1turszjwn05lflpewurw96rfrd3h6x8flgs5uf8"
				},
				"id": 1
			}`,
			fn: func(c api.Calls) error {
				txtID, changeAddr, err := c.AVMExportAVAX(api.ExportAVAXArgs{})
				if err != nil {
					return nil
				}
				if txtID != "25VzbNzt3gi2vkE3Kr6H9KJeSR2tXkr8FsBCm3vARnB5foLVmx" || changeAddr != "X-avax1turszjwn05lflpewurw96rfrd3h6x8flgs5uf8" {
					return fmt.Errorf("invald return values: txtID: %s, changeAddr: %s", txtID, changeAddr)
				}

				return nil
			},
		}, {
			name: "AVMExportKey",
			response: `{
				"jsonrpc":"2.0",
				"id"     :1,
				"result" :{
					"privateKey":"PrivateKey-2w4XiXxPfQK4TypYqnohRL8DRNTz9cGiGmwQ1zmgEqD9c9KWLq"
				}
			}`,
			fn: func(c api.Calls) error {
				privateKey, err := c.AVMExportKey("username", "password", "address")
				if err != nil {
					return nil
				}
				if privateKey != "PrivateKey-2w4XiXxPfQK4TypYqnohRL8DRNTz9cGiGmwQ1zmgEqD9c9KWLq" {
					return fmt.Errorf("invald return value: privateKey: %s", privateKey)
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
