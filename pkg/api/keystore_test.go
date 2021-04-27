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

func TestKeyStoreExportUser(t *testing.T) {
	ts := NewServer(`{
		"jsonrpc":"2.0",
		"id"     :1,
		"result" :{
			"user":"4CsUh5sfVwz2jNrJXBVpoPtDsb4tZksWykqmxC5CXoDEERyhoRryq62jYTETYh53y13v7NzeReisi",
			"encoding":"cb58"
		}
	}`)
	defer ts.Close()

	client, err := network.NewNodeClient(ts.URL, 1)
	if err != nil {
		t.Error(err)
	}

	api := api.NewAPICalls(client)
	user, encoding, err := api.ExportUser("username", "password")

	if user != "4CsUh5sfVwz2jNrJXBVpoPtDsb4tZksWykqmxC5CXoDEERyhoRryq62jYTETYh53y13v7NzeReisi" {
		t.Errorf("unexpected value :%s", user)
	}

	if encoding != "cb58" {
		t.Errorf("unexpected value :%s", encoding)
	}
}
