package librawallet

import (
	"errors"
	"strings"

	"github.com/codemaveric/libra-go/pkg/types"
)

type WalletLibrary struct {
	Mnemonic   Mnemonic
	KeyFactory *KeyFactory
	KeyLeaf    uint64
	AddressMap map[string]uint64 //string is hex string of AccountAddress
}

func NewWalletLibrary(mnemonicStr string) *WalletLibrary {
	var mnemonic Mnemonic
	if mnemonicStr == "" {
		mnemonic = generateMnemonic()
	} else {
		mnemonic = strings.Split(mnemonicStr, " ")
	}
	seed := NewSeed(mnemonic, "")
	return &WalletLibrary{
		Mnemonic:   mnemonic,
		KeyFactory: NewKeyFactory(seed),
		KeyLeaf:    0,
		AddressMap: make(map[string]uint64),
	}
}

func (w *WalletLibrary) GenerateAddress(depth uint64) error {
	if w.KeyLeaf > depth {
		return errors.New("Addresses already generated up to the supplied depth")
	}
	_, _, err := w.NewAddress()
	if err != nil {
		return err
	}

	return nil
}

func (w *WalletLibrary) NewAddress() (types.AccountAddress, uint64, error) {
	key := w.KeyFactory.GenerateKey(w.KeyLeaf)
	address := key.GetAddress()
	child := w.KeyLeaf
	w.KeyLeaf += 1
	if _, ok := w.AddressMap[address.ToString()]; ok {
		return address, child, errors.New("This address is already in your wallet")
	}
	w.AddressMap[address.ToString()] = child
	return address, child, nil
}
