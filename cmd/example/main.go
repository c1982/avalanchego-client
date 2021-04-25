package main

import (
	"avalanchego-client/pkg/api"
	"avalanchego-client/pkg/network"
	"log"
)

func main() {

	username := "user1"
	password := "dPeGbERNC4fDIN9F1X0S3KLi28oWYEf/1lO4FUEc"
	endpoint := "http://10.10.10.107:9650"

	client, err := network.NewNodeClient(endpoint, 1)
	if err != nil {
		log.Fatal(err)
	}

	api := api.NewAPICalls(client)
	ok, err := api.IsBootstrapped()
	if err != nil {
		log.Fatal(err)
	}

	err = api.CreateUser(username, password)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("bootstraped is %t", ok)
}
