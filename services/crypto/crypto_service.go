package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"fileprotector/services/keyprovider"
)

func Encrypt(key, data []byte) (cipherData []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	return gcm.Seal(nil, keyprovider.GetNonce(), data, nil), nil
}

func Decrypt(key, cipherData []byte) (orginal []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	return gcm.Open(nil, keyprovider.GetNonce(), cipherData, nil)
}
