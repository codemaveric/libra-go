package librawallet

import (
	"github.com/codemaveric/libra-go/pkg/crypto"
	"github.com/codemaveric/libra-go/pkg/types"
	"golang.org/x/crypto/sha3"
)

type Account struct {
	Address  types.AccountAddress
	KeyPair  *crypto.KeyPair
	Sequence uint64
}

func AccountFromSecret(privateKey crypto.PrivateKey) *Account {
	keyPair := crypto.NewKeyPair(privateKey.Value)
	digest := sha3.Sum256(keyPair.PublicKey.Value)
	address := digest[:]
	return &Account{Address: address, KeyPair: keyPair}
}
