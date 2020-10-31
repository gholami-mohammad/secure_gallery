package keyprovider

import (
	"crypto/sha1"

	"golang.org/x/crypto/pbkdf2"
)

var nonce, salte []byte

func GetNonce() []byte {
	if nonce == nil {
		nonce = pbkdf2.Key([]byte("noncepass"), GetSalt(), 4096, 12, sha1.New)
	}
	return nonce
}

func GetSalt() []byte {
	if salte == nil {
		salte = []byte("salt")
	}
	return salte
}

func GetEncryptioKey(password []byte) []byte {
	return pbkdf2.Key(password, GetSalt(), 4096, 32, sha1.New)
}
