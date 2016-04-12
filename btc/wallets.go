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
	rand := rand.Reader
	priv, err := btckey.GenerateKey(rand)

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

func (w *walletImpl) GetAddress() string {
	return w.btckey.ToAddress()
}

func (w *walletImpl) GetPrivateKeyWIF() string {
	return w.btckey.ToWIF()
}


