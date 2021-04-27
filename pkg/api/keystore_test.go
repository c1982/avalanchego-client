package api_test

import (
	"avalanchego-client/pkg/api"
	"avalanchego-client/pkg/network"
	"fmt"
	"testing"
)

func TestKeyStore(t *testing.T) {
	calltests := []struct {
		name     string
		response string
		fn       func(api.Calls) error
	}{
		{
			name: "CreateUser",
			response: `{
				"jsonrpc":"2.0",
				"id"     :1,
				"result" :{
					"success":false
				}
			}`,
			fn: func(c api.Calls) error {
				err := c.CreateUser("username", "password")
				if err == nil {
					return err
				}

				return nil
			},
		}, {
			name: "ListUsers",
			response: `{
				"jsonrpc":"2.0",
				"id"     :1,
				"result" :{
					"users":[
						"user1",
						"user2",
						"user3",
						"user4"
					]
				}
			}`,
			fn: func(c api.Calls) error {
				users, err := c.ListUsers()
				if err != nil {
					return err
				}

				if len(users) != 4 {
					return fmt.Errorf("invalid user count. have: 4, got: %d", len(users))
				}
				return nil
			},
		}, {
			name: "KeyStoreExportUser",
			response: `{
				"jsonrpc":"2.0",
				"id"     :1,
				"result" :{
					"user":"4CsUh5sfVwz2jNrJXBVpoPtDsb4tZksWykqmxC5CXoDEERyhoRryq62jYTETYh53y13v7NzeReisi",
					"encoding":"cb58"
				}
			}`,
			fn: func(c api.Calls) error {
				user, encoding, err := c.ExportUser("username", "password")
				if err != nil {
					return err
				}
				if user != "4CsUh5sfVwz2jNrJXBVpoPtDsb4tZksWykqmxC5CXoDEERyhoRryq62jYTETYh53y13v7NzeReisi" {
					return fmt.Errorf("unexpected value :%s", user)
				}

				if encoding != "cb58" {
					return fmt.Errorf("unexpected value :%s", encoding)
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
