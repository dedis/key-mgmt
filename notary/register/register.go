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
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	pgp "golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

var db *leveldb.DB

// XXX move all gloabl vars to config:
var hostName string
var auth smtp.Auth

func init() {
	// XXX move to config
	tokenDbPath := "pendingTokens"
	var err error
	db, err = leveldb.OpenFile(tokenDbPath, nil)

	if err != nil {
		panic(fmt.Sprintf("Couldn't find path %s", tokenDbPath))
	}
	// TODO write config file
	auth = smtp.PlainAuth("",
		"user@example.com",
		"password",
		"mail.example.com")
	hostName = "localhost:8080"

}

func validData(userMail string, publicKey string) (bool, pgp.EntityList) { // XXX error handling
	e, err := mail.ParseAddress(userMail)
	if err != nil {
		return false, nil
	}
	// XXX should we just allow keys with no identity? /*|| len(e.Name) > 0*/
	if len(e.Address) == 0 || strings.EqualFold(e.Name, userMail) {
		return false, nil
	}
	// TODO
	entity, err := pgp.ReadArmoredKeyRing(bytes.NewBufferString(publicKey))
	if err != nil || entity == nil || len(entity) == 0 {
		return false, nil
	}
	if len(entity) != 1 {
		return false, nil
	}
	return true, entity
}

func sendConfirmationLink(userMail string, userEntity pgp.EntityList) error {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return err
	}
	sum := sha256.Sum256(token)
	// XXX use TLS
	url := fmt.Sprintf("http://%s/confirm?t=%s", hostName, base64.URLEncoding.EncodeToString(sum[:]))
	msg := "TODO ... \n" + url

	fmt.Println("Clear text", msg)
	buf := new(bytes.Buffer)
	a, _ := armor.Encode(buf, "PGP MESSAGE", nil)
	plaintext, err := pgp.Encrypt(a, userEntity, nil, /*use server's key to sign msg*/
		nil, nil)
	if err != nil {
		return err
	}

	//_, err = plaintext.Write([]byte(msg))
	fmt.Fprintf(plaintext, msg)
	if err != nil {
		return err
	}
	err = plaintext.Close()
	a.Close()
	if err != nil {
		return err
	}

	// TODO write mail in (go routine) and write token into pending tokens DB
	fmt.Println("TODO send message")
	// Encode to base64
	bytes, err := ioutil.ReadAll(buf)
	if err != nil {
		return err
	}
	fmt.Println(string(bytes))

	return nil
}

func saveToken(token []byte, uAdress string, entitiy pgp.Entity) error {
	buf := new(bytes.Buffer)
	entitiy.Serialize(buf)
	// XXX other Info we need to store: creation time, username (?)
	return db.Put(token, buf.Bytes(), &opt.WriteOptions{Sync: true})
}

func storePendingUserToMerkle(token []byte) error {
	b, err := db.Get(token, nil)

	if err != nil {
		db.Delete(token, nil)
		return err
	}

	pReader := packet.NewReader(bytes.NewReader(b))
	entity, err := pgp.ReadEntity(pReader)
	if err != nil {
		db.Delete(token, nil)
		return err
	}
	// extract first key
	var uMail string
	for id := range entity.Identities {
		uMail = id
	}

	fmt.Printf("TODO store entity in merkle prefix tree (e-mail=%v and key's fingerprint=%v)", uMail, entity.PrimaryKey.Fingerprint)
	db.Delete(token, nil)

	return nil
}
