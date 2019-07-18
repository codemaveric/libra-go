package goclient

import (
	"encoding/hex"
	"fmt"

	"github.com/codemaveric/libra-go/gowrapper"
	"github.com/codemaveric/libra-go/pkg/common"
	"github.com/codemaveric/libra-go/pkg/crypto"
	"github.com/codemaveric/libra-go/pkg/types"
	"github.com/golang/protobuf/proto"
)


func accountResourcePath() string {
	// Hardcoded Resource Path, because for now recomputing the resource path gives the same result everytime.
	return "01217da6c6b3e19f1825cfb2676daecce3bf3de03cf26647c78df00b371b25cc97"
}

func decodeSignedTransactionWP(signedTransactionWP *gowrapper.SignedTransactionWithProof) (*types.SignedTransactionWithProof, error) {
	// Decode Transaction
	signedTransaction := signedTransactionWP.GetSignedTransaction()
	rawTxnBytes := signedTransaction.GetRawTxnBytes()
	transaction, err := decodeRawTransactionBytes(rawTxnBytes)
	if err != nil {
		return nil, err
	}

	libraSignTransaction := &types.SignedTransaction{
		RawTransaction: transaction,
		PublicKey:      &crypto.PublicKey{Value: signedTransaction.GetSenderPublicKey()},
		Signature:      &crypto.Signature{Value: signedTransaction.GetSenderSignature()},
	}

	// Decode Events
	eventList := []*types.ContractEvent{}
	events := signedTransactionWP.GetEvents()
	if events != nil {
		for _, v := range events.GetEvents() {
			eventList = append(eventList, &types.ContractEvent{
				AccountAddress: v.AccessPath.GetAddress(),
				EventData:      v.GetEventData(),
				SequenceNumber: v.GetSequenceNumber(),
				Path:           v.AccessPath.GetPath(),
			})
		}
	}
	libraSignedTransactionWP := &types.SignedTransactionWithProof{
		SignedTransaction: libraSignTransaction,
		Proof:             signedTransactionWP.GetProof(),
		Events:            eventList,
	}

	return libraSignedTransactionWP, nil
}

func decodeRawTransactionBytes(rawTxnBytes []byte) (*types.RawTransaction, error) {
	var rawTxn gowrapper.RawTransaction
	err := proto.Unmarshal(rawTxnBytes, &rawTxn)
	if err != nil {
		return nil, err
	}
	rawProgram := rawTxn.GetProgram()
	arguments := []*types.LibraProgramArgument{}
	for _, v := range rawProgram.GetArguments() {
		arguments = append(arguments, &types.LibraProgramArgument{Type: types.LibraProgramArgumentType(v.Type), Data: v.Data})
	}
	program := &types.LibraProgram{
		Arguments: arguments,
		Code:      rawProgram.GetCode(),
		Modules:   rawProgram.GetModules(),
	}

	libraTransaction := &types.RawTransaction{
		Sender:             rawTxn.GetSenderAccount(),
		SequenceNumber:     rawTxn.GetSequenceNumber(),
		MaxGasAmount:       rawTxn.GetMaxGasAmount(),
		GasUnitPrice:       rawTxn.GetGasUnitPrice(),
		TransactionProgram: program,
		ExpirationTime:     rawTxn.GetExpirationTime(),
	}

	return libraTransaction, nil

}

func decodeAccountStateBlob(blob []byte) map[string][]byte {
	canonicalSerializer := common.NewCanonicalSerializer(blob)
	bloblen := int(canonicalSerializer.Read32())

	state := make(map[string][]byte)

	for i := 0; i < bloblen; i++ {

		keyLen := int(canonicalSerializer.Read32())

		keyBuffer := make([]byte, keyLen)
		for j := 0; j < keyLen; j++ {
			keyBuffer[j] = canonicalSerializer.Read8()
		}
		valueLen := int(canonicalSerializer.Read32())
		valueBuffer := make([]byte, valueLen)
		for k := 0; k < valueLen; k++ {
			valueBuffer[k] = canonicalSerializer.Read8()
		}
		fmt.Println(hex.EncodeToString(keyBuffer))
		state[hex.EncodeToString(keyBuffer)] = valueBuffer[:]
	}
	return state
}
