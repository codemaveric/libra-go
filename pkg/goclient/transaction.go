package goclient

import (
	"encoding/binary"
	"encoding/hex"
	"time"

	"github.com/codemaveric/libra-go/gowrapper"
	"github.com/codemaveric/libra-go/pkg/crypto"
	"github.com/codemaveric/libra-go/pkg/librawallet"
	"github.com/golang/protobuf/proto"
)

const GAS_UNIT_PRICE uint64 = 0
const MAX_GAS_AMOUNT uint64 = 10000
const TX_EXPIRATION int64 = 100

// Const of transaction program
const HEX_PEER_TO_PEER_TRANSFER_CODE = "4c49425241564d0a010007014a00000004000000034e000000060000000c54000000050000000d5900000004000000055d0000002900000004860000002000000007a60000000d00000000000001000200010300020002040203020402063c53454c463e0c4c696272614163636f756e74046d61696e0f7061795f66726f6d5f73656e64657200000000000000000000000000000000000000000000000000000000000000000001020004000c000c01110102"

type LibraProgram struct {
	Code []byte
}

func encodeTransferProgram(receiverAddress string, numCoins uint64) (*gowrapper.Program, error) {
	code, _ := hex.DecodeString(HEX_PEER_TO_PEER_TRANSFER_CODE)
	address, _ := hex.DecodeString(receiverAddress)
	amount := make([]byte, 8)
	binary.LittleEndian.PutUint64(amount, numCoins)
	arg1 := &gowrapper.TransactionArgument{Type: gowrapper.TransactionArgument_ADDRESS, Data: address}
	arg2 := &gowrapper.TransactionArgument{Type: gowrapper.TransactionArgument_U64, Data: amount}
	arg := []*gowrapper.TransactionArgument{arg1, arg2}
	modules := [][]byte{} // Empty modules
	pr := &gowrapper.Program{Code: code, Arguments: arg, Modules: modules}

	return pr, nil
}

func createSubmitTransactionReq(program *gowrapper.Program, sender *librawallet.Account, gasUnitPrice, maxGasAmount uint64) (*gowrapper.SubmitTransactionRequest, error) {
	raw_txn := &gowrapper.RawTransaction{
		SequenceNumber: sender.Sequence,
		SenderAccount:  sender.Address,
		MaxGasAmount:   maxGasAmount,
		GasUnitPrice:   gasUnitPrice,
		Payload:        &gowrapper.RawTransaction_Program{program},
		ExpirationTime: uint64(time.Now().Unix() + TX_EXPIRATION),
	}
	proto.Marshal(raw_txn)
	txn_bytes, err := proto.Marshal(raw_txn)
	if err != nil {
		return nil, err
	}
	//Initialize a new RAW_TRANSACTINO Hasher
	cryptoHasher := crypto.NewCryptoHasher([]byte(crypto.RAW_TRANSACTION))
	hash := cryptoHasher.Hash(txn_bytes) // Hash Transaction byte

	signature := crypto.SignMessage(hash, sender.KeyPair.PrivateKey)

	signedTxn := &gowrapper.SignedTransaction{
		SenderPublicKey: sender.KeyPair.PublicKey.Value,
		RawTxnBytes:     txn_bytes,
		SenderSignature: signature.Value,
	}
	req := &gowrapper.SubmitTransactionRequest{
		SignedTxn: signedTxn,
	}
	return req, nil
}
