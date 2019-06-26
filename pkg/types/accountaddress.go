package types

import "encoding/hex"

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

//TODO: Implement Get Account Address from secret key
func (a AccountAddress) FromSecret(secretKeyHex string) AccountAddress {
	return make([]byte, 32)
}
