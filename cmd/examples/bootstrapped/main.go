package main

import (
	"avalanchego-client/pkg/api"
	"avalanchego-client/pkg/network"
	"fmt"
	"log"
)

func main() {
	client, err := network.NewNodeClient("http://127.0.0.1:9650", 1)
	if err != nil {
		log.Fatal(err)
	}

	nodeclient := api.NewAPICalls(client)
	ok, err := nodeclient.InfoIsBootstrapped("X")
	checkerr(err)
	if !ok {
		log.Fatal(fmt.Errorf("chain X does not bootstrapped yet"))
	}

	fmt.Println("chain X bootstrapped!")

	ok, err = nodeclient.InfoIsBootstrapped("P")
	checkerr(err)
	if !ok {
		log.Fatal(fmt.Errorf("chain P does not bootstrapped yet"))
	}

	fmt.Println("chain P bootstrapped!")

	ok, err = nodeclient.InfoIsBootstrapped("C")
	checkerr(err)
	if !ok {
		log.Fatal(fmt.Errorf("chain C does not bootstrapped yet"))
	}

	fmt.Println("chain C bootstrapped!")
}

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
