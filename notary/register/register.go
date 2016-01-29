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
	"log"
	"net"
	"net/mail"
	"net/smtp"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/openpgp/packet"
)

var db *leveldb.DB

// XXX move all gloabl vars to config:
var hostName string
var authC authConfig
var debug bool = true

type authConfig struct {
	Identity string
	Username string
	Password string
	Host     string
	Port     string
}

func init() {
	// XXX move to config
	tokenDbPath := "dbs/pendingTokens"
	var err error
	db, err = leveldb.OpenFile(tokenDbPath, nil)

	if err != nil {
		log.Fatalf("Couldn't find path %s", tokenDbPath)
	}

	authC = authConfig{}
	_, err = toml.DecodeFile("../smtp.toml", &authC)
	if err != nil {
		panic("Couldn't read config:" + err.Error())
	}
	fmt.Printf("config: %+v", authC)
	hostName = "localhost:8080"
}

func validData(userMail string, publicKey string) (bool, *openpgp.Entity) { // XXX error handling
	e, err := mail.ParseAddress(userMail)
	if err != nil {
		return false, nil
	}
	// XXX should we just allow keys with no identity? /*|| len(e.Name) > 0*/
	if len(e.Address) == 0 || strings.EqualFold(e.Name, userMail) {
		return false, nil
	}
	entityList, err := openpgp.ReadArmoredKeyRing(bytes.NewBufferString(publicKey))
	if err != nil || entityList == nil || len(entityList) == 0 {
		return false, nil
	}
	if len(entityList) != 1 {
		return false, nil
	}
	return true, entityList[0]
}

func sendConfirmationLink(userMail string, userEntity openpgp.Entity) error {
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return err
	}
	sum := sha256.Sum256(token)
	// XXX use TLS
	url := fmt.Sprintf("http://%s/confirm?t=%s", hostName, base64.URLEncoding.EncodeToString(sum[:]))
	msg := "TODO use e-mail template ... \n" + url

	buf := new(bytes.Buffer)
	armored, _ := armor.Encode(buf, "PGP MESSAGE", nil)
	plaintext, err := openpgp.Encrypt(armored, openpgp.EntityList{&userEntity}, nil, // TODO use server's private key to sign msg
		nil, nil)
	if err != nil {
		return err
	}

	fmt.Fprintf(plaintext, msg)
	if err != nil {
		return err
	}
	err = plaintext.Close()
	if err != nil {
		return err
	}
	err = armored.Close()
	if err != nil {
		return err
	}

	// send e-mail and write token into pending tokens DB
	go func() {
		msgBytes, err := ioutil.ReadAll(buf)
		if err != nil {
			log.Print(err)
			return
		}
		fmt.Println(string(msgBytes))
		auth := smtp.PlainAuth(authC.Identity,
			authC.Username,
			authC.Password,
			authC.Host)

		smtpServer := net.JoinHostPort(authC.Host, authC.Port)
		if !debug {
			err = smtp.SendMail(
				smtpServer,
				auth,
				authC.Username,
				[]string{userMail},
				msgBytes,
			)
			if err != nil {
				log.Print(err)
				return
			}
		}
		saveToken(token, userMail, userEntity)
	}()

	return nil
}

func saveToken(token []byte, uAdress string, entitiy openpgp.Entity) error {
	buf := new(bytes.Buffer)
	entitiy.Serialize(buf)
	// XXX other info we need to store: creation time, username (?)
	return db.Put(token, buf.Bytes(), &opt.WriteOptions{Sync: true})
}

func storePendingUserToMerkle(token []byte) error {
	b, err := db.Get(token, nil)

	if err != nil {
		db.Delete(token, nil)
		return err
	}

	pReader := packet.NewReader(bytes.NewReader(b))
	entity, err := openpgp.ReadEntity(pReader)
	if err != nil {
		db.Delete(token, nil)
		return err
	}
	// extract first id
	var uMail string
	for id := range entity.Identities {
		uMail = id
	}

	fmt.Printf("TODO store entity in merkle prefix tree (e-mail=%v and key's fingerprint=%v)", uMail, entity.PrimaryKey.Fingerprint)
	db.Delete(token, nil)

	return nil
}
