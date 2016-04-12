package btc

import (
	"github.com/tobyjsullivan/btckeygenie/btckey"
	"crypto/rand"
)

type Wallet interface {
	GetAddress() string
	GetPrivateKeyWIF() string
}

func GenerateWallet() (Wallet, error) {
	priv, err := generateBtcKey(rand.Reader)

	if err != nil {
		return nil, err
	}

	wallet := &walletImpl{
		btckey: priv}

	return wallet, nil
}

type walletImpl struct {
	btckey btckey.PrivateKey
}

// Produces a compressed bitcoin address
func (w *walletImpl) GetAddress() string {
	return w.btckey.ToAddress()
}

func (w *walletImpl) GetPrivateKeyWIF() string {
	return w.btckey.ToWIF()
}

// Overrides for testing
var (
	generateBtcKey = btckey.GenerateKey
)

