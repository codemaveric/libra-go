package librawallet

import (
	"strings"

	"github.com/tyler-smith/go-bip39"
)

type Mnemonic []string

func generateMnemonic() Mnemonic {
	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)
	return strings.Split(mnemonic, " ")
}

func (m Mnemonic) ToBytes() []byte {
	wordString := m.ToString()

	return []byte(wordString)
}

func (m Mnemonic) ToString() string {
	return strings.Join(m, " ")
}
