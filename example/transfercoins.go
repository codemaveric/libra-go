package main

import (
	"fmt"
	"github.com/codemaveric/libra-go/pkg/goclient"
	"github.com/codemaveric/libra-go/pkg/librawallet"
	"log"
	"strings"
)

func main() {
	// Change the mnemonic to something else,
	// or you will get address used by others
	mnemonic := "present good satochi coin future media giant"
	wallet := librawallet.NewWalletLibrary(mnemonic)

	// create a new address
	address1, childNum, err := wallet.NewAddress()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Adddress1: %s\n", address1.ToString())

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

	// Set the current account sequence
	sourceAccount.Sequence = SourceaccState.SequenceNumber

	// The balance will not be 500000000 if you don't change the mnemonic,
	// cause the account has been used multiple times
	fmt.Printf("Address1 balance: %d\n", SourceaccState.Balance)

	// Create another address
	address2, childNum, err := wallet.NewAddress()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Adddress2: %s\n", address2.ToString())

	// Generate KeyPair with mnemonic and childNum
	keyPair = librawallet.GenerateKeyPair(strings.Split(mnemonic, " "), childNum)

	// Create Account from KeyPair
	destAccount := librawallet.NewAccountFromKeyPair(keyPair)

	var amount uint64 = 100000000
	// Transfer Coins from source account to destination address
	err = libraClient.TransferCoins(sourceAccount, destAccount.Address.ToString(), amount, 0, 10000, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Transfer %d coins from account1 to account2\n", amount)

	// Get Account State
	SourceaccState, err = libraClient.GetAccountState(sourceAccount.Address.ToString())
	if err != nil {
		log.Fatal(err)
	}

	// The balance should be 400000000 now
	fmt.Printf("Address1 balance: %d\n", SourceaccState.Balance)

	// Get Account State
	destaccState, err := libraClient.GetAccountState(destAccount.Address.ToString())
	if err != nil {
		log.Fatal(err)
	}

	// The balance should be 100000000 now
	fmt.Printf("Address2 balance: %d\n", destaccState.Balance)

	// Get Account Transaction by Sequence Number
	transaction, err := libraClient.GetAccountTransaction(sourceAccount.Address.ToString(), 0, true)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(transaction)
}
