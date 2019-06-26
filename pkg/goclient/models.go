package goclient

import (
	"github.com/codemaveric/libra-go/pkg/common"
)

type AccountState struct {
	AuthenticationKey   []byte `json:"authentication_key"`
	Balance             uint64 `json:"balance"`
	RecievedEventsCount uint64 `json:"received_events_count"`
	SentEventsCount     uint64 `json:"sent_events_count"`
	SequenceNumber      uint64 `json:sequence_number`
}

func (a *AccountState) Deserialize(payload []byte) error {
	canonicalSerializer := common.NewCanonicalSerializer(payload)
	canonicalSerializer.ReadXBytes(81)
	a.Balance = canonicalSerializer.Read64()
	a.RecievedEventsCount = canonicalSerializer.Read64()
	a.SentEventsCount = canonicalSerializer.Read64()
	a.SequenceNumber = canonicalSerializer.Read64()
	return nil
}

func (a *AccountState) Serialize() []byte {
	return nil
}
