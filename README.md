# Libra Golang Client
Libra Golang Client is library to interact with Libra Blockchain

> Note: The project is still under some major development! The Package is not stable and will keep changing!

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
	"time"

	"github.com/codemaveric/libra-go/pkg/goclient"
	"github.com/codemaveric/libra-go/pkg/librawallet"
)

func main() {
  // Instantiate LibraWallet with mnemonic 
	wallet := librawallet.NewWalletLibrary("times good gospel coin social media giant")
	address, _, err := wallet.NewAddress() // Generate a new address
	if err != nil {
		log.Print(err)
	}
	log.Print(address.ToString())
  // Instantiate LibraClient with TestNet Configuration
	libraClient := goclient.NewLibraClient(goclient.LibraClientConfig{Network: TestNet})
  
	// Mint coins on testnet to reciever address, amount is in microlibra
	libraClient.MintWithFaucetService(address.ToString(), 25000000, true)
  
	accState, err := libraClient.GetAccountState(address.ToString())

	log.Print(accState.Balance)
	
}

```

## Roadmap

Libra Transfer, so you can transfer coin from your account to other account and so much more. Your contribution is welcome!!!

## Contribution
Feel free to contribute by opening issues or PR's.

## License
MIT
