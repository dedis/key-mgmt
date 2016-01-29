package http

import (
	"fmt"
	"net/http"

	"github.com/dedis/key-mgmt/notary/register"
)

// Register handles requests in which the user sends a public key
func Register(w http.ResponseWriter, r *http.Request) {
	// TODO read e-mail and public GPG key from request
	fmt.Fprintf(w, "TODO handle %s!", r.URL.Path[1:])
}

// Confirm handles confirmation links  which are send in encrypted e-mail
// challenges: (e.g. https://<hostname>/confirm?t=<superRandomToken>)
func Confirm(w http.ResponseWriter, r *http.Request) {
	// TODO parse confirmation token, delete from tokenDb an write user-key-binding
	// to merkle tree
	tokenStr := r.URL.Query().Get("t")
	token := []byte(tokenStr)
	if len(token) != register.TokenLen { // more validation
		// return some JSON
		return
	}
	register.StorePendingUser(token)

	//fmt.Fprintf(w, "TODO handle %s with query %s!", r.URL.Path[1:], r.URL.Query())
}

// Update XXX later
/*func Update(w http.ResponseWriter, r *http.Request) {
	// TODO read e-mail and public GPG key from request
	// is this different from Register (how)?
	fmt.Fprintf(w, "TODO handle %s!", r.URL.Path[1:])
}*/

// Lookup returns a public-key bound to the requested user-id (e-mail address)
func Lookup(w http.ResponseWriter, r *http.Request) {
	// TODO read e-mail
	fmt.Fprintf(w, "TODO handle %s!", r.URL.Path[1:])
}
