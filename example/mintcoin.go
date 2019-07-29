package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/codemaveric/libra-go/pkg/goclient"
	"github.com/codemaveric/libra-go/pkg/librawallet"
)

func main() {
	// Change the mnemonic to something else,
	// or you will get address used by others
	mnemonic := "present good satochi coin future media giant"
	wallet := librawallet.NewWalletLibrary(mnemonic)

	// create a new address
	address, childNum, err := wallet.NewAddress()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Adddress: %s\n", address.ToString())

	// Generate KeyPair with mnemonic and childNum
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
	fmt.Println("Mint coin successfully.")

	// Get Account State
	SourceaccState, err := libraClient.GetAccountState(sourceAccount.Address.ToString())
	if err != nil {
		log.Fatal(err)
	}

	// The balance will not be 500000000 if you don't change the mnemonic,
	// cause the account has been used multiple times
	fmt.Printf("balance: %d\n", SourceaccState.Balance)
}
