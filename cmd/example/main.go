package main

import (
	"avalanchego-client/pkg/api"
	"avalanchego-client/pkg/network"
	"log"
)

func main() {
	client, err := network.NewNodeClient("http://10.10.10.107:9650", 1)
	if err != nil {
		log.Fatal(err)
	}

	api := api.NewAPICalls(client)
	ok, err := api.IsBootstrapped()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("bootstraped is %t", ok)
}
