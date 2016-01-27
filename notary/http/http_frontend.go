package http

import (
	"fmt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO read e-mail and public GPG key from request
	fmt.Fprintf(w, "TODO handle %s!", r.URL.Path[1:])
}

func Confirm(w http.ResponseWriter, r *http.Request) {
	// TODO parse confirmation token, delete from tokenDb an write user-key-binding
	// to merkle tree
	fmt.Fprintf(w, "TODO handle %s with query %s!", r.URL.Path[1:], r.URL.Query())
}

func Update(w http.ResponseWriter, r *http.Request) {
	// TODO read e-mail and public GPG key from request
	// is this different from Register (how)?
	fmt.Fprintf(w, "TODO handle %s!", r.URL.Path[1:])
}

func Lookup(w http.ResponseWriter, r *http.Request) {
	// TODO read e-mail
	fmt.Fprintf(w, "TODO handle %s!", r.URL.Path[1:])
}
