package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tobyjsullivan/btcpayments/handlers"
)

func buildRouter() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", placeholderHandler).Methods("GET")
	router.HandleFunc("/wallets", handlers.WalletsPostHandler).Methods("POST")

	return router
}

func placeholderHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Try the endpoint: curl -X POST 127.0.0.1:3000/wallets")
}
