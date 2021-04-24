# avalanchego-client
avalanchego

### Creating a Local Testnet

To create a single node testnet, run:

```sh
./avalanchego --network-id=local --staking-enabled=false --snow-sample-size=1 --snow-quorum-size=1 --http-host=ENTER-YOUR-LOCAL-IP-ADDRESS
```

This launches an Avalanche network with one node.