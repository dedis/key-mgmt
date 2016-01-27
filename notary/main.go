package main

import "net/http"
import http_front "github.com/dedis/key-mgmt/notary/http"

func main() {

	// TODO read config (port, public-key etc)

	http.HandleFunc("/register", http_front.Register)
	http.HandleFunc("/update", http_front.Update)
	http.HandleFunc("/confirm", http_front.Update)
	http.HandleFunc("/lookup", http_front.Lookup)

	http.ListenAndServe(":8080", nil)
}
