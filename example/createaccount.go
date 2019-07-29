package main

import (
	"fmt"
	"log"

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

	// Get Account with Child number
	account, err := wallet.GetAccount(childNum)
	// If you have your secrey key you can create account from it.
	// secreyKey := "hex string of secret key here"
	// account, err := librawallet.NewAccountFromSecret(secretKey)

	if err != nil {
		log.Fatal(err)
	}

	// Show the private key
	fmt.Printf("Private key: %x\n", account.KeyPair.PrivateKey.Value)
}
