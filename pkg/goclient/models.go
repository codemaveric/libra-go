package goclient

import (
	"encoding/hex"

	"github.com/codemaveric/libra-go/pkg/common"
)

const AccountStatePath string = "01217da6c6b3e19f1825cfb2676daecce3bf3de03cf26647c78df00b371b25cc97"

type AccountState struct {
	AuthenticationKey   string `json:"authentication_key"`
	Balance             uint64 `json:"balance"`
	RecievedEventsCount uint64 `json:"received_events_count"`
	SentEventsCount     uint64 `json:"sent_events_count"`
	SequenceNumber      uint64 `json:sequence_number`
}

func (a *AccountState) Deserialize(payload []byte) error {
	canonicalSerializer := common.NewCanonicalSerializer(payload)
	authenticationKeyLen := canonicalSerializer.Read32()
	data := canonicalSerializer.ReadXBytes(uint64(authenticationKeyLen))
	a.AuthenticationKey = hex.EncodeToString(data)
	a.Balance = canonicalSerializer.Read64()
	a.RecievedEventsCount = canonicalSerializer.Read64()
	a.SentEventsCount = canonicalSerializer.Read64()
	a.SequenceNumber = canonicalSerializer.Read64()
	return nil
}

func (a *AccountState) Serialize() []byte {
	return nil
}
