package main

import (

	"github.com/codegangsta/negroni"
	"os"
)

func main() {
	router := buildRouter()


	n := negroni.Classic()
	n.UseHandler(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	n.Run(":"+port)
}
