package file

import (
	"io"
	"io/ioutil"
	"os"
)

func ReadFile(filePath string) (bts []byte, err error) {
	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	bts, err = ioutil.ReadAll(file)
	if err != nil {
		return
	}

	return
}

func WriteFile(filePath string, data []byte) error {
	fileToWrite, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}
	defer fileToWrite.Close()
	_, err = fileToWrite.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}
	_, err = fileToWrite.Write(data)
	if err != nil {
		return err
	}
	return fileToWrite.Sync()
}
