package main

import (
	"fmt"
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
	address, childNum, err := wallet.NewAddress()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Adddress: %s\n", address.ToString())

	// Generate KeyPair with mnemonic and childNum
	keyPair := librawallet.GenerateKeyPair(strings.Split(mnemonic, " "), childNum)

	// Create Account from KeyPair
	sourceAccount := librawallet.NewAccountFromKeyPair(keyPair)
	fmt.Printf("Address from KeyPair: %s\n", sourceAccount.Address.ToString())

	// show the private key
	fmt.Printf("Private key: %x\n", sourceAccount.KeyPair.PrivateKey.Value)
}


