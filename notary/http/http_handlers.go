package http

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"net/http"

	"github.com/dedis/key-mgmt/notary/register"
)

// Register handles requests in which the user sends a public key
func Register(w http.ResponseWriter, r *http.Request) {
	// TODO read e-mail and public GPG key from request
	if r.Method == "GET" {
		// XXX for testing purposes: simple form where user can post their email and
		// public key. Move this to the client (mailvelope/end-to-end) and only
		// handle POST requests here!
		t, _ := template.ParseFiles("test.html")
		t.Execute(w, nil)
	} else if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)

		//r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
		// attention: If you do not call ParseForm method, the following data can not be obtained form
		fmt.Println(r.Form) // print information on server side.
		fmt.Println("path", r.URL.Path)
		fmt.Println("scheme", r.URL.Scheme)
		fmt.Println(r.Form["url_long"])
		userMail := r.Form.Get("email")
		if userMail == "" {
			return
		}
		file, _, err := r.FormFile("ascii-public-key")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		ok, entity := register.ValidData(userMail, file)
		if !ok {
			// XXX return some reasonable error data
			return
		}
		// if debug == false, will send an email accodring to smtp.toml
		if err := register.SendConfirmationLink(userMail, *entity); err != nil {
			// XXX could not send mail
			return
		}

		// XXX use protobuf instead of form
		fmt.Fprintf(w, "Check your e-mail for an encrypted confirmation challenge")
	}
}

// Confirm handles confirmation links  which are send in encrypted e-mail
// challenges: (e.g. https://<hostname>/confirm?t=<superRandomToken>)
// TODO return JSON or JWS instead of strings
func Confirm(w http.ResponseWriter, r *http.Request) {
	// TODO parse confirmation token, delete from tokenDb an write user-key-binding
	// to merkle tree
	tokenStr := r.URL.Query().Get("t")
	token, err := base64.URLEncoding.DecodeString(tokenStr)
	if err != nil {
		return
	}
	if len(token) == register.TokenLen { // more validation
		if err := register.StorePendingUser(token); err != nil {
			fmt.Fprint(w, "Could not register")
		} else {
			fmt.Fprint(w, "Successfully stored user-key binding")
		}
	} else {
		fmt.Fprintf(w, "Invalid token len=%d", len(token))
	}

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
