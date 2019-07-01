package crypto

import (
	"golang.org/x/crypto/ed25519"
)

type PrivateKey struct {
	Value ed25519.PrivateKey
}

type PublicKey struct {
	Value ed25519.PublicKey
}

type Signature struct {
	Value []byte
}

type KeyPair struct {
	PrivateKey *PrivateKey
	PublicKey  *PublicKey
}

func NewKeyPair(secret []byte) *KeyPair {
	privateKey := &PrivateKey{Value: secret}
	publicKey := privateKey.Value.Public().(ed25519.PublicKey)
	return &KeyPair{PrivateKey: privateKey, PublicKey: &PublicKey{Value: publicKey}}
}

func SignMessage(msg *HashValue, privateKey *PrivateKey) *Signature {
	signature := ed25519.Sign(privateKey.Value, msg.hash[:])
	return &Signature{Value: signature}
}
