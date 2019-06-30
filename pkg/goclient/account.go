package goclient

import (
	"encoding/hex"
	"fmt"

	"github.com/codemaveric/libra-go/pkg/common"
)

type Account struct {
	libraClient *LibraClient
}

func accountResourcePath() string {
	// Hardcoded Resource Path, because for now recomputing the resource path gives the same result everytime.
	return "01217da6c6b3e19f1825cfb2676daecce3bf3de03cf26647c78df00b371b25cc97"
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
