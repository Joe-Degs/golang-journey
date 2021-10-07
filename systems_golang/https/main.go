package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello")
	})
	err := http.ListenAndServeTLS(":443", "keys/server.crt", "keys/server.key", nil)
	if err != nil {
		log.Fatal("Listen and Server TLS error", err)
	}
}
