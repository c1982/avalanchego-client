# avalanchego-client

The Avalanche Platform Go Library 

## TODO's

* Admin API
* Auth API
* AVM API (X-Chain)
* EVM API (C-Chain)
* Health API
* Info API
* Keystore API
* Metrics API
* PlatformVM API (P-Chain)

### Documents

## Requirements

Go 1.15+

## Installation

> go get github.com/c1982/avalanchego-client

## Usage

Check out the cmd/examples/ directory to see how to use this libray.

```golang
package main

import (
	"avalanchego-client/pkg/api"
	"avalanchego-client/pkg/network"
	"log"
)

func main() {

	username := "user1"
	password := "password"
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
```

### Creating a Local Testnet

To create a single node testnet, run:

```sh
./avalanchego --network-id=local --staking-enabled=false --snow-sample-size=1 --snow-quorum-size=1 --http-host=ENTER-YOUR-LOCAL-IP-ADDRESS
```

This launches an Avalanche network with one node.