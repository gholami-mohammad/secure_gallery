package controllers

import (
	"fileprotector/services/crypto"
	"fileprotector/services/file"
	"fileprotector/services/password"
	"fmt"
	"html/template"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func Upload(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(200000)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	directoryPath := fmt.Sprintf("%s", os.Getenv("ROOT_PATH"))
	err = os.MkdirAll(directoryPath, os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fhs := r.MultipartForm.File["files"]
	for _, fh := range fhs {
		err = uploadMultipartFile(*fh, directoryPath)
		if err != nil {
			continue
		}
	}
	_ = r.MultipartForm.RemoveAll()

	w.Header().Set("content-type", "text/html")
	_, _ = w.Write([]byte(`<meta http-equiv="refresh" content="1; url = /" />`))
	_, _ = w.Write([]byte(`files uploaded, redirecting...`))
}

func uploadMultipartFile(fh multipart.FileHeader, target string) error {
	uploadedFile, err := fh.Open()
	if err != nil {
		return err
	}

	defer uploadedFile.Close()

	bts, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		return err
	}

	key := password.GetPassword()
	cipherData, err := crypto.Encrypt(key, bts)
	if err != nil {
		return err
	}

	fileName := file.GenerateUniqueFilename(target, fh.Filename, 1)
	filePath := target + "/" + fileName

	err = file.WriteFile(filePath, cipherData)
	if err != nil {
		return err
	}

	return nil
}
