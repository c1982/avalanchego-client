package api_test

import (
	"avalanchego-client/pkg/api"
	"avalanchego-client/pkg/network"
	"testing"
)

func TestKeyStoreCreateUser(t *testing.T) {
	ts := NewServer(`{
		"jsonrpc":"2.0",
		"id"     :1,
		"result" :{
			"success":false
		}
	}`)
	defer ts.Close()

	client, err := network.NewNodeClient(ts.URL, 1)
	if err != nil {
		t.Error(err)
	}

	api := api.NewAPICalls(client)
	err = api.CreateUser("username", "password")
	if err == nil {
		t.Error(err)
	}
}

func TestKeyStoreUserList(t *testing.T) {
	ts := NewServer(`{
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
	}`)
	defer ts.Close()

	client, err := network.NewNodeClient(ts.URL, 1)
	if err != nil {
		t.Error(err)
	}

	api := api.NewAPICalls(client)
	users, err := api.ListUsers()
	if err != nil {
		t.Error(err)
	}

	if len(users) != 4 {
		t.Errorf("invalid user count. have: 4, got: %d", len(users))
	}
}
