package types

import (
	"github.com/codemaveric/libra-go/pkg/crypto"
)

type LibraProgram struct {
	Code      []byte
	Arguments []*LibraProgramArgument
	Modules   [][]byte
}

type LibraProgramArgument struct {
	Type LibraProgramArgumentType
	Data []byte
}

type LibraArgument interface {
	IsArgumentType()
}

type LibraArgument_UInt uint64
type LibraArgument_String string
type LibraArgument_Address AccountAddress
type LibraArgument_ByteArray []byte

func (l LibraArgument_UInt) IsArgumentType()      {}
func (l LibraArgument_String) IsArgumentType()    {}
func (l LibraArgument_Address) IsArgumentType()   {}
func (l LibraArgument_ByteArray) IsArgumentType() {}

type LibraProgramArgumentType uint64

const (
	U64       LibraProgramArgumentType = 0
	ADDRESS   LibraProgramArgumentType = 1
	STRING    LibraProgramArgumentType = 2
	BYTEARRAY LibraProgramArgumentType = 3
)

type RawTransaction struct {
	// Sender's Address.
	Sender AccountAddress
	// Sequence Number of this transaction corresponding to sender's account.
	SequenceNumber uint64
	// The transaction program to execute.
	TransactionProgram *LibraProgram
	// Maximal total gas specified by wallet to spend for this transaction.
	MaxGasAmount uint64
	// Maximal price can be paid per gas.
	GasUnitPrice uint64
	// Expiration time for this transaction.
	ExpirationTime uint64
}

type SignedTransaction struct {
	// The raw transaction.
	RawTransaction *RawTransaction
	// Sender's public key. When checking the signature, we first need to check whether this key
	// is indeed the pre-image of the pubkey hash stored under sender's account.
	PublicKey *crypto.PublicKey
	// Signature of the transaction that correspond with the public key.
	Signature *crypto.Signature
	// The original raw bytes generated from the wallet.
	RawTxnBytes []byte
}

// `TransactionInfo` is the object we store in the transaction accumulator. It
// consists of the transaction as well as the execution result of this
// transaction. This are later returned to the client so that a client can
// validate the tree
type TransactionInfo struct {
	// Hash of the signed transaction that is stored
	SignTransactionHash []byte
	// The root hash of Sparse Merkle Tree describing the world state at the end
	// of this transaction
	StateRootHash []byte
	// The root hash of Merkle Accumulator storing all events emitted during this
	// transaction.
	EventRootHash []byte
	// The amount of gas used by this transaction.
	GasUsed uint64
}

type SignedTransactionWithProof struct {
	// The version of the returned signed transaction. This is also like transaction Id
	Version uint64
	// The transaction itself.
	SignedTransaction *SignedTransaction
	// The proof authenticating the signed transaction.
	Proof SignedTransactionProof
	// The events yielded by executing the transaction.
	Events []*ContractEvent
}

type SignedTransactionProof interface {
}

type ContractEvent struct {
	// The data payload of the event.
	EventData []byte
	// Address to access path.
	AccountAddress AccountAddress
	// The number of messages that have been emitted to the path previously.
	SequenceNumber uint64
	// The path that the event was emitted to.
	Path []byte
}
