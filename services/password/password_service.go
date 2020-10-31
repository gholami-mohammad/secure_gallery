package password

import (
	"encoding/hex"
	"fileprotector/services/crypto"
	"fileprotector/services/keyprovider"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"golang.org/x/crypto/ssh/terminal"
)

var password []byte

func GetPassword() []byte {
	return password
}

func ReadPasswordInput() []byte {
	fmt.Println("If you had not set any password, this password will be used for encryption.")
	fmt.Print("Enter Your password: ")
	bts, err := terminal.ReadPassword(0)
	if err != nil {
		log.Fatalln(err)
		return nil
	}
	fmt.Println("")
	password = keyprovider.GetEncryptioKey(bts)
	return password
}

func CheckPassword(pass []byte) {
	fname := os.Getenv("ROOT_PATH") + "/.lock"
	if _, err := os.Stat(fname); err != nil {
		err = setAppInitializationPass()
		if err != nil {
			log.Fatalln(err)
			return
		}
		return
	}

	// check that exising password is valid
	_, err := isPasswordValid()
	if err != nil {
		log.Fatalln(err)
		return
	}

}

func setAppInitializationPass() error {
	fname := os.Getenv("ROOT_PATH") + "/.lock"
	fmt.Println("Directory is empty, this password will be set as your encryption password")
	err := os.MkdirAll(os.Getenv("ROOT_PATH"), os.ModePerm)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()

	cipher, err := crypto.Encrypt(password, password)
	if err != nil {
		return err
	}

	str := hex.EncodeToString(cipher)
	_, err = f.WriteString(str)
	if err != nil {
		return err
	}
	return nil
}

func isPasswordValid() (bool, error) {
	fname := os.Getenv("ROOT_PATH") + "/.lock"
	// check that exising password is valid
	f, err := os.OpenFile(fname, os.O_RDONLY, os.ModePerm)
	if err != nil {
		return false, err
	}
	defer f.Close()
	bts, err := ioutil.ReadAll(f)
	if err != nil {
		return false, err
	}

	cipher, err := hex.DecodeString(string(bts))
	if err != nil {
		return false, err
	}
	_, err = crypto.Decrypt(password, cipher)
	if err != nil {
		return false, err
	}

	return true, nil
}
