package handlers

import (
	"net/http"
	"github.com/tobyjsullivan/btcpayments/btc"
	"fmt"
	"encoding/json"
)

func WalletsPostHandler(w http.ResponseWriter, req *http.Request) {
	wallet, err := btc.GenerateWallet()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "We encountered an unexpected error: %s", err.Error())
		return
	}

	display := &walletDisplay{
		Address: wallet.GetAddress(),
		PrivateKeyWIF: wallet.GetPrivateKeyWIF()}

	out, err := json.Marshal(display)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error marshalling response: %s", err.Error())
		return
	}

	fmt.Fprintln(w, string(out))
}

type walletDisplay struct {
	Address string `json:"address"`
	PrivateKeyWIF string `json:"private-key-wif"`
}
