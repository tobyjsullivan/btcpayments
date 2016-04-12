package main

import (
	"os"

	"github.com/codegangsta/negroni"
)

func main() {
	router := buildRouter()

	n := negroni.Classic()
	n.UseHandler(router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	n.Run(":" + port)
}
