package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"rsc.io/quote"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, quote.Go())
}

func main() {
	var envPort = os.Getenv("PORT")
	var port string
	if envPort == "" {
		port = "8080"
	} else {
		port = envPort
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
