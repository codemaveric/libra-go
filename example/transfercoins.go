package main

import (
	"fmt"
	"log"

	"github.com/codemaveric/libra-go/pkg/goclient"
	"github.com/codemaveric/libra-go/pkg/librawallet"
)

func main() {
	// Change the mnemonic to something else,
	// or you will get address used by others.
	mnemonic := "present good satochi coin future media giant"
	wallet := librawallet.NewWalletLibrary(mnemonic)

	// create a new address
	address1, childNum, err := wallet.NewAddress()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Source Account Address: %s\n", address1.ToString())

	// Create account  with mnemonic and childNum
	sourceAccount := librawallet.NewAccount(mnemonic, childNum)
	// If you have your secrey key you can create account from it.
	// secreyKey := "hex string of secret key here"
	// sourceAccount := librawallet.NewAccountFromSecret(secretKey)

	// Libra Client Configuration
	config := goclient.LibraClientConfig{
		Host:    "ac.testnet.libra.org",
		Port:    "80",
		Network: goclient.TestNet,
	}
	// Instantiate LibraClient with Configuration
	libraClient := goclient.NewLibraClient(config)

	// Amount in Micro-Libra
	var amount uint64 = 100000000
	destinationAddress := "f4aebe371e4176143c3409122d0adf43c0e00a6552b5b0ae9980d8981fcd0221"

	// Note: Make sure the current source account sequence is set
	// Get the Account State and Set the current account sequence.
	// Transfer Coins from source account to destination address.
	err = libraClient.TransferCoins(sourceAccount, destinationAddress, amount, 0, 10000, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transfer %d coins from %x to %s", amount, sourceAccount.Address, destinationAddress)
}
