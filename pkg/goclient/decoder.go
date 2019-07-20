package goclient

import (
	"encoding/hex"
	"errors"

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

func decodeTransactionsListWP(transactionListWP *gowrapper.TransactionListWithProof) ([]*types.SignedTransactionWithProof, error) {

	if transactionListWP == nil {
		return nil, errors.New("Empty TransactionsListWithProof")
	}

	signedTransactionWP := []*types.SignedTransactionWithProof{}
	signedTransactions := transactionListWP.GetTransactions()

	if len(signedTransactions) < 1 {
		return nil, errors.New("Transactions not found")
	}

	firstTransactionVersion := transactionListWP.GetFirstTransactionVersion().Value
	eventsVersion := transactionListWP.GetEventsForVersions()
	var eventList []*gowrapper.EventsList

	if eventsVersion != nil {
		eventList = eventsVersion.GetEventsForVersion()
	}

	for i := 0; i < len(signedTransactions); i++ {
		var events []*types.ContractEvent
		// TODO: I think I should come back and check this later
		if eventsVersion != nil {
			events = decodeEventList(eventList[i])
		}
		signedTransaction, _ := decodeSignedTransaction(signedTransactions[i])
		signedTransactionWP = append(signedTransactionWP, &types.SignedTransactionWithProof{
			SignedTransaction: signedTransaction,
			Version:           firstTransactionVersion,
			Events:            events,
		})
		firstTransactionVersion++
	}

	return signedTransactionWP, nil
}

func decodeSignedTransactionWP(signedTransactionWP *gowrapper.SignedTransactionWithProof) (*types.SignedTransactionWithProof, error) {
	// Decode Transaction
	libraSignTransaction, err := decodeSignedTransaction(signedTransactionWP.GetSignedTransaction())

	if err != nil {
		return nil, err
	}
	// Decode Events
	eventList := decodeEventList(signedTransactionWP.GetEvents())

	libraSignedTransactionWP := &types.SignedTransactionWithProof{
		Version:           signedTransactionWP.GetVersion(),
		SignedTransaction: libraSignTransaction,
		Proof:             signedTransactionWP.GetProof(),
		Events:            eventList,
	}

	return libraSignedTransactionWP, nil
}

func decodeEventList(events *gowrapper.EventsList) []*types.ContractEvent {
	eventList := []*types.ContractEvent{}
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
	return eventList
}

func decodeSignedTransaction(signedTransaction *gowrapper.SignedTransaction) (*types.SignedTransaction, error) {
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

	return libraSignTransaction, nil
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
