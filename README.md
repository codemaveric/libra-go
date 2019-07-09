# Libra Golang Client
Libra Golang Client is library to interact with Libra Blockchain

> Note: The project is still under major development! The Package is not stable and will keep changing!

## Installation
Run to install the package

```
go get github.com/codemaveric/libra-go
```

## Usage
Example

```Go
package main

import (
	"log"
	"strings"

	"github.com/codemaveric/libra-go/pkg/goclient"
	"github.com/codemaveric/libra-go/pkg/librawallet"
)

func main() {
 	// I will advice you to change the mnemonic to something else
	mnemonic := "present good satochi coin future media giant"
	wallet := librawallet.NewWalletLibrary(mnemonic)
	address, childNum, err := wallet.NewAddress()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(address.ToString())

	// Generate Keypair with mnemonic and childNum
	keyPair := librawallet.GenerateKeyPair(strings.Split(mnemonic, " "), childNum)

	// Create Account from KeyPair
	sourceAccount := librawallet.NewAccountFromKeyPair(keyPair)
	// Libra Client Configuration
	config := goclient.LibraClientConfig{
		Host:    "ac.testnet.libra.org",
		Port:    "80",
		Network: goclient.TestNet,
	}
	// Instantiate LibraClient with Configuration
	libraClient := goclient.NewLibraClient(config)

	// Mint Coin from Test Faucet to account generated
	err = libraClient.MintWithFaucetService(sourceAccount.Address.ToString(), 500000000, true)

	if err != nil {
		log.Fatal(err)
	}

	// Get Account State
	SourceaccState, err := libraClient.GetAccountState(sourceAccount.Address.ToString())
	if err != nil {
		log.Fatal(err)
	}
	log.Print(SourceaccState.Balance)

	// Set the current account sequence
	sourceAccount.Sequence = SourceaccState.SequenceNumber

	// Transfer Coins from source account to destination address
	err = libraClient.TransferCoins(sourceAccount, "f4aebe371e4176143c3409122d0adf43c0e00a6552b5b0ae9980d8981fcd0221", 11000000, 0, 10000, true)
	if err != nil {
		log.Fatal(err)
	}	
}

```

## Roadmap

1) Transaction History.
2) Get Transaction Details with tran_id

## Contribution
Feel free to contribute by opening issues or PR's.

## License
MIT
