package types

import (
	"encoding/hex"

	"golang.org/x/crypto/sha3"
)

const (
	ADDRESS_LENGTH int = 32
)

type AccountAddress []byte

func NewAccountAddress(address string) AccountAddress {
	res, _ := hex.DecodeString(address)
	return res
}
func (a AccountAddress) ToString() string {
	return hex.EncodeToString(a)
}

func (a AccountAddress) IsValidBytes() bool {
	return len(a) == ADDRESS_LENGTH
}

// Get Address from secret key in hex
func GetAddressFromSecret(secretKey string) AccountAddress {
	keybytes, _ := hex.DecodeString(secretKey)
	publicKey := make([]byte, ADDRESS_LENGTH)
	copy(publicKey, keybytes[32:])
	digest := sha3.Sum256(publicKey)
	return digest[:]
}
