// register contains all functionality to register a new email-to-key-binding
//
// The protocol can be summarized as follows:
//
//
//                Client                              Server
//  1) initiate     |                                   |
//                  |    register(mailAddr, PK) (http)  |
//                  | ------------------------------->  | Register:
//                  |                                   | 1) validate data
//                  |                                   |   (whitelisted mail & gpg public-key)
//                  |                                   | 2) generate token, encrypt mail using PK containing URL with token
//                  |    e-mail & sk challenge (smtp)   |
//                  | <------------------------------   |
//                  |                                   |
//                  |                                   |
// 2) decrypt mail  |   response (GET req with token)   |
//    using sk      | ------------------------------->  | if request arrives in time:
//                  |                                   | add to tree, otherwise
//                  |                                   | delete token from DB
//                  |          notify client            |
//                  | <- - - - - - - - - - - - - - - -  |
package register

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/mail"

	"github.com/syndtr/goleveldb/leveldb"
	pgp "golang.org/x/crypto/openpgp"
)

var db *leveldb.DB

// XXX move to config
var hostName string

func init() {
	// XXX move to config
	tokenDbPath := "tokenDb"
	var err error
	db, err = leveldb.OpenFile(tokenDbPath, nil)
	if err != nil {
		panic(fmt.Sprintf("Couldn't find path %s", tokenDbPath))
	}

}

func validData(userMail string, publicKey string) (bool, pgp.EntityList) {
	e, err := mail.ParseAddress(userMail)
	if err != nil {
		return false, nil
	}
	if len(e.Address) == 0 || len(e.Name) > 0 {
		return false, nil
	}
	// TODO
	entity, err := pgp.ReadArmoredKeyRing(bytes.NewBufferString(publicKey))
	if err != nil || entity == nil || len(entity) == 0 {
		return false, nil
	}
	return true, entity
}

func sendConfirmationLink(userMail, userEntity pgp.EntityList) error {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return err
	}
	// XXX md5 secure enough for this purpose?
	sum := md5.Sum(token)
	// XXX use TLS
	url := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("http://%s/confirm?t=%s", hostName, sum)))
	msg := "TODO ... \n" + url

	buf := new(bytes.Buffer)
	w, err := pgp.Encrypt(buf, userEntity, nil, nil, nil)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}

	// TODO write mail in (go routine) and write token into pending tokens DB

	return nil
}
