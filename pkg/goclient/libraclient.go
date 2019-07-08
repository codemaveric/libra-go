package goclient

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/codemaveric/libra-go/gowrapper"
	"github.com/codemaveric/libra-go/pkg/librawallet"
	"github.com/codemaveric/libra-go/pkg/types"
	"google.golang.org/grpc"
)

const (
	DefaultFaucetServerHost  string = "faucet.testnet.libra.org"
	DefaultTestnetServerHost string = "ac.testnet.libra.org"
)

type LibraClient struct {
	client gowrapper.AdmissionControlClient
}

func NewLibraClient(config LibraClientConfig) *LibraClient {
	if config.Host == "" {
		config.Host = DefaultTestnetServerHost
	}
	if config.Port == "" {
		config.Port = "80"
	}
	address := fmt.Sprintf("%s:%s", config.Host, config.Port)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	client := gowrapper.NewAdmissionControlClient(conn)
	return &LibraClient{client: client}
}

func (l *LibraClient) TransferCoins(sender *librawallet.Account, receipientAddress string, amount uint64, gasUnitPrice uint64, maxGasAmount uint64, isBlocking bool) error {
	program, err := encodeTransferProgram(receipientAddress, amount)
	if err != nil {
		return err
	}
	req, err := createSubmitTransactionReq(program, sender, gasUnitPrice, maxGasAmount)
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
	res, err := l.client.SubmitTransaction(ctx, req, grpc.WaitForReady(true))
	if err != nil {
		return err
	}
	acStatus := res.GetAcStatus()
	// If status is not Accepted return error else increase the sequence number of the sender
	if acStatus != gowrapper.AdmissionControlStatus_Accepted {
		return errors.New(fmt.Sprintf("Transaction failed with status: %s", gowrapper.AdmissionControlStatus_name[int32(acStatus)]))
	}
	sender.Sequence += 1
	return nil
}

func (l *LibraClient) GetAccountState(address string) (*AccountState, error) {
	res, err := l.GetAccountBlob(address)
	if err != nil {
		return nil, err
	}
	account := &AccountState{}
	payload := res.AccountStateWithProof.Blob.Blob
	accountTree := decodeAccountStateBlob(payload)
	ap := accountResourcePath()
	err = account.Deserialize(accountTree[ap])
	return account, err
}

func (g *LibraClient) GetAccountBlob(address string) (*gowrapper.GetAccountStateResponse, error) {
	decoded, err := hex.DecodeString(address)
	if err != nil {
		return nil, err
	}
	accountState := &gowrapper.GetAccountStateRequest{Address: decoded}
	accountStateReq := &gowrapper.RequestItem_GetAccountStateRequest{GetAccountStateRequest: accountState}
	requestItem := &gowrapper.RequestItem{RequestedItems: accountStateReq}

	res, err := g.GetLatestWithProof([]*gowrapper.RequestItem{requestItem})
	if err != nil {
		return nil, err
	}
	responseItems := res.ResponseItems[0].ResponseItems
	response := responseItems.(*gowrapper.ResponseItem_GetAccountStateResponse)
	return response.GetAccountStateResponse, nil
}

func (g *LibraClient) GetLatestWithProof(requestItems []*gowrapper.RequestItem) (*gowrapper.UpdateToLatestLedgerResponse, error) {
	req := &gowrapper.UpdateToLatestLedgerRequest{ClientKnownVersion: 0, RequestedItems: requestItems}
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	res, err := g.client.UpdateToLatestLedger(ctx, req, grpc.WaitForReady(true))
	if err != nil {
		return nil, err
	}

	err = verify(req, res)
	return res, err
}

func verify(req *gowrapper.UpdateToLatestLedgerRequest, resp *gowrapper.UpdateToLatestLedgerResponse) error {
	ledgeInfo, signatures := resp.LedgerInfoWithSigs.LedgerInfo, resp.LedgerInfoWithSigs.Signatures

	// Verify that the same or a newer ledger info is returned.
	if ledgeInfo.GetVersion() <= req.GetClientKnownVersion() {
		return errors.New("Got stale ledger_info with version {}, known version: {}.")
	}
	// Verify ledger info signatures.
	if !(ledgeInfo.GetVersion() == 0 && len(signatures) == 0) {

	}
	if len(req.GetRequestedItems()) != len(resp.GetResponseItems()) {
		return errors.New("Number of request items ({}) does not match that of response items ({}).")
	}
	return nil
}

// Mint coins on testnet to reciever address
// num_coin should be in microlibra.
func (l *LibraClient) MintWithFaucetService(address string, num_coins uint64, is_blocking bool) error {
	furl := fmt.Sprintf("http://%s?amount=%d&address=%s", DefaultFaucetServerHost, num_coins, address)
	client := http.Client{Timeout: time.Second * 5}
	res, err := client.Get(furl)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return errors.New("Failed to query remote faucet server")
	}
	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	hdfj := string(payload)
	sequence, _ := strconv.Atoi(hdfj)
	//accountAddress := types.NewAccountAddress(address)
	if is_blocking {
		l.waitForTransaction(make([]byte, 32), uint64(sequence))
	}
	return nil
}

func (l *LibraClient) waitForTransaction(address types.AccountAddress, sequenceNumber uint64) {
	maxIteration := 50
	for {
		maxIteration--
		seqNo, _ := l.GetSequenceNumber(address)

		if seqNo >= sequenceNumber {
			break
		}
		if maxIteration <= 0 {
			log.Print("wait_for_transaction timeout")
			break
		}
		time.Sleep(time.Millisecond * 100)
	}
}

func (l *LibraClient) GetSequenceNumber(address types.AccountAddress) (uint64, error) {
	account, err := l.GetAccountState(address.ToString())

	if err != nil {
		return 0, err
	}
	return account.SequenceNumber, nil
}
