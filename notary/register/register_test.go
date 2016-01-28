package register

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"testing"
)

var dummyTestKey = `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: Mailvelope v1.3.4
Comment: https://www.mailvelope.com

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
=YSEB
-----END PGP PUBLIC KEY BLOCK-----

`

func TestValidData(t *testing.T) {
	ok, entityList := validData("user@freetopia.org", dummyTestKey)
	if !ok {
		t.Fatal("Could not parse key or email")
	}
	fmt.Printf("FYI: got %s\nwith len=%d", entityList, len(entityList))
}

func TestSaveAndReadTokenData(t *testing.T) {
	uMail := "user@freetopia.org"
	ok, entityList := validData(uMail, dummyTestKey)
	if !ok {
		t.Fatal("Could not parse key or email")
	}
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		t.Fatal("Could not generate token", err)
	}
	// XXX md5 secure enough for this purpose?
	entity := *entityList[0]
	sum := md5.Sum(token)
	if err := saveToken(sum[:], uMail, entity); err != nil {
		t.Fatal("Could not save token:", err)
	}
	sendConfirmationLink(uMail, entityList)
	if err := storePendingUserToMerkle(sum[:]); err != nil {
		t.Fatal("Could not store pending user, did not find token:", err)
	}
}
