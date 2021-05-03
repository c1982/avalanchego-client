package main

import (
	"avalanchego-client/pkg/api"
	"avalanchego-client/pkg/network"
	"fmt"
	"log"
)

func main() {
	username := "test-user-3"
	password := "JZNwwhk2coyIR4fGMedPMSdC1GX2CTs+LsimZZjT"
	address := "X-fuji18vz4vzzvwnf8gk2ac2f0qmguuxsfzcrwjalzg4"
	client, err := network.NewNodeClient("http://192.168.1.25:9650", 1)
	if err != nil {
		log.Fatal(err)
	}

	nodeclient := api.NewAPICalls(client)
	ok, err := nodeclient.InfoIsBootstrapped("X")
	checkerr(err)
	if !ok {
		log.Fatal(fmt.Errorf("chain X does not bootstrapped yet"))
	}

	ok, err = nodeclient.KeystoreCreateUser(username, password)
	checkerr(err)
	if !ok {
		log.Fatal(fmt.Errorf("user do not created"))
	}

	fmt.Println("keystore user has been created")

	address, err = nodeclient.AvmCreateAddress(username, password)
	checkerr(err)

	fmt.Printf("address has been created: %s\n", address)

	balances, err := nodeclient.AvmGetAllBalances(address)
	checkerr(err)

	//Please Go to https://faucet.avax-test.network/ and send yourself AWAX
	fmt.Printf("asset count: %d\n", len(balances.Balances))
	for i := 0; i < len(balances.Balances); i++ {
		fmt.Printf("%s = %s", balances.Balances[i].Asset, balances.Balances[i].Balance)
	}
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
