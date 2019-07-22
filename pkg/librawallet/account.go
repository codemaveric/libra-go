package librawallet

import (
	"strings"

	"github.com/codemaveric/libra-go/pkg/crypto"
	"github.com/codemaveric/libra-go/pkg/types"
	"golang.org/x/crypto/sha3"
)

type Account struct {
	Address  types.AccountAddress
	KeyPair  *crypto.KeyPair
	Sequence uint64
}

func GenerateKeyPair(mnemonic Mnemonic, childNum uint64) *crypto.KeyPair {
	seed := NewSeed(mnemonic, "")
	keyfactory := NewKeyFactory(seed)
	priveKey := keyfactory.GenerateKey(childNum)
	return crypto.NewKeyPair(priveKey.PrivateKey)
}

// Create Account from Mnemonic and child number.
func NewAccount(mnemonic string, childNumber uint64) *Account {
	seed := NewSeed(strings.Split(mnemonic, " "), "")
	keyfactory := NewKeyFactory(seed)
	privateKey := keyfactory.GenerateKey(childNumber)
	return NewAccountFromSecret(privateKey.ToString())
}

// Create account from key pair
func NewAccountFromKeyPair(keyPair *crypto.KeyPair) *Account {
	digest := sha3.Sum256(keyPair.PublicKey.Value)
	address := digest[:]
	return &Account{Address: address, KeyPair: keyPair}
}

// Create Account from Secret Key.
func NewAccountFromSecret(secret string) *Account {
	keyPair := crypto.NewKeyPair([]byte(secret))
	return NewAccountFromKeyPair(keyPair)
}
