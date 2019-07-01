package crypto

import (
	"hash"

	"golang.org/x/crypto/sha3"
)

const HASH_LENGTH = 32

const LIBRA_HASH_SUFFIX = "@@$$LIBRA$$@@"

const (
	RAW_TRANSACTION = "RawTransaction"
)

type HashValue struct {
	hash [HASH_LENGTH]byte
}

func from_sha3(buffer []byte) *HashValue {
	state := sha3.New256()
	digest := state.Sum(buffer)
	var hash [HASH_LENGTH]byte
	copy(hash[:], digest)
	return &HashValue{hash: hash}
}

type CryptoHasher struct {
	state hash.Hash
}

func NewCryptoHasher(salt []byte) *CryptoHasher {
	state := sha3.New256()
	if len(salt) != 0 {
		hashSuf := []byte(LIBRA_HASH_SUFFIX)
		salt = append(salt, hashSuf...)
		hash := from_sha3(salt).hash
		state.Write(hash[:])
	}

	return &CryptoHasher{state: state}
}

func (c *CryptoHasher) Hash(data []byte) *HashValue {
	digest := c.state.Sum(data)
	var hash [HASH_LENGTH]byte
	copy(hash[:], digest)
	return &HashValue{hash: hash}
}

// type CryptoHash struct {
// 	hasher *CryptoHasher
// }

// func NewCryptoHash(salt string) *CryptoHash {
// 	hasher := NewCryptoHasher([]byte(salt))
// 	return &CryptoHash{hasher: hasher}
// }

// func (c *CryptoHash) Hash(data []byte) *HashValue {
// 	digest := c.hasher.state.Sum(data)
// 	var hash [HASH_LENGTH]byte
// 	copy(hash[:], digest)
// 	return &HashValue{hash: hash}
// }
