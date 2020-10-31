package crypto

import (
	"encoding/hex"
	"fileprotector/services/keyprovider"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestEncrypt(t *testing.T) {
	Convey("invalid key size", t, func() {
		key := []byte("1")

		data := []byte(hex.EncodeToString([]byte("something")))
		_, err := Encrypt(key, data)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "invalid key size 1")
	})
	Convey("pass", t, func() {
		key := keyprovider.GetEncryptioKey([]byte("password"))

		data := []byte(hex.EncodeToString([]byte("something")))
		_, err := Encrypt(key, data)
		So(err, ShouldBeNil)
	})
}

func TestDecrypt(t *testing.T) {
	Convey("invalid key size", t, func() {
		key := []byte("1")

		_, err := Decrypt(key, []byte("CIPHERDATA"))
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "invalid key size 1")
	})
	Convey("invalid password", t, func() {
		key := keyprovider.GetEncryptioKey([]byte("password"))
		data := []byte(hex.EncodeToString([]byte("something")))
		cipherdata, _ := Encrypt(key, data)

		key2 := keyprovider.GetEncryptioKey([]byte("otherpass"))
		_, err := Decrypt(key2, cipherdata)
		So(err, ShouldNotBeNil)
		So(err.Error(), ShouldContainSubstring, "cipher: message authentication failed")
	})
	Convey("pass", t, func() {
		key := keyprovider.GetEncryptioKey([]byte("password"))
		data := []byte(hex.EncodeToString([]byte("something")))
		cipherdata, _ := Encrypt(key, data)

		_, err := Decrypt(key, cipherdata)
		So(err, ShouldBeNil)
	})
}
