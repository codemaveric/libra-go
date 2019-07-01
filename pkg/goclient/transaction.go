package goclient

import (
	"encoding/json"

	"github.com/codemaveric/libra-go/gowrapper"
	"github.com/codemaveric/libra-go/pkg/crypto"
	"github.com/codemaveric/libra-go/pkg/librawallet"
)

const GAS_UNIT_PRICE uint64 = 0
const MAX_GAS_AMOUNT uint64 = 10000
const TX_EXPIRATION uint64 = 100

type LibraProgram struct {
	Code []byte
}

func encodeTransferProgram(receiverAddress string, numCoins uint64) (*gowrapper.Program, error) {

	return nil, nil
}

func createSubmitTransactionReq(program *gowrapper.Program, sender *librawallet.Account, gasUnitPrice, maxGasAmount uint64) (*gowrapper.SubmitTransactionRequest, error) {
	raw_txn := &gowrapper.RawTransaction{
		SequenceNumber: sender.Sequence,
		SenderAccount:  sender.Address,
		MaxGasAmount:   maxGasAmount,
		GasUnitPrice:   gasUnitPrice,
		Payload:        &gowrapper.RawTransaction_Program{program},
		ExpirationTime: 0,
	}
	txn_bytes, err := json.Marshal(&raw_txn)
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
