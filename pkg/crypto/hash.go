package crypto

import (
	"golang.org/x/crypto/sha3"
	"hash"
)

const HASH_LENGTH = 32

const LIBRA_HASH_SUFFIX = "@@$$LIBRA$$@@"

const (
	RAW_TRANSACTION = "RawTransaction"
)

type HashValue struct {
	hash [HASH_LENGTH]byte
}

func (h *HashValue) GetValue() [HASH_LENGTH]byte {
	return h.hash
}

func from_sha3(buffer []byte) *HashValue {
	state := sha3.New256()
	state.Write(buffer)
	var digest [HASH_LENGTH]byte
	state.Sum(digest[:0])
	return &HashValue{hash: digest}
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
		// log.Println(hash)
		state.Write(hash[:])
	}
	return &CryptoHasher{state: state}
}

func (c *CryptoHasher) Hash(data []byte) *HashValue {
	c.state.Write(data)
	var digest [HASH_LENGTH]byte
	c.state.Sum(digest[:0])
	return &HashValue{hash: digest}
}
