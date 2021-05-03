package main

import (
	"avalanchego-client/pkg/api"
	"avalanchego-client/pkg/network"
	"fmt"
	"log"
)

func main() {

	username := "user1"
	password := "dPeGbERNC4fDIN9F1X0S3KLi28oWYEf/1lO4FUEc"
	client, err := network.NewNodeClient("http://127.0.0.1:9650", 1)
	if err != nil {
		log.Fatal(err)
	}

	nodeclient := api.NewAPICalls(client)
	ok, err := nodeclient.InfoIsBootstrapped("X")
	checkerr(err)
	if !ok {
		log.Fatal(fmt.Errorf("Chanin-X does not bootstrapped yet"))
	}

	fmt.Println("Chain-X is bootstrapped")

	//1. Create address for subnet control keys
	address1, err := nodeclient.PlatformCreateAddress(username, password)
	checkerr(err)
	fmt.Printf("address1 created: %s\n", address1)

	address2, err := nodeclient.PlatformCreateAddress(username, password)
	checkerr(err)
	fmt.Printf("address2 created: %s\n", address2)

	//2. Create Subnet
	txid, err := nodeclient.PlatformCreateSubnet(api.CreateSubnetArgs{
		Controlkeys: []string{address1, address2},
		Threshold:   2,
		Username:    username,
		Password:    password,
	})
	checkerr(err)

	fmt.Printf("transaction ID: %s\n", txid)

	subnets, err := nodeclient.PlatformGetSubnets([]string{})
	checkerr(err)

	if len(subnets.Subnets[0].Controlkeys) != 2 {
		fmt.Printf("subnet cannot validated: have: 2 got: %d\n", len(subnets.Subnets[0].Controlkeys))
	} else {
		fmt.Println("subnet created successfully")
	}
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
