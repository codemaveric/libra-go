package librawallet

import "strings"

type Mnemonic []string

func (m Mnemonic) ToBytes() []byte {
	wordString := m.ToString()

	return []byte(wordString)
}

func (m Mnemonic) ToString() string {
	return strings.Join(m, " ")
}
