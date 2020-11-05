package controllers

import (
	"fileprotector/services/crypto"
	"fileprotector/services/file"
	"fileprotector/services/password"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("frontend/src/templates/index.html"))

	files, err := ioutil.ReadDir(os.Getenv("ROOT_PATH"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	filenames := []os.FileInfo{}
	dirs := []os.FileInfo{}
	for _, finfo := range files {
		if finfo.IsDir() {
			dirs = append(dirs, finfo)
			continue
		}
		if finfo.Name() == ".lock" {
			continue
		}

		filenames = append(filenames, finfo)
	}
	params := make(map[string]interface{})
	params["filenames"] = filenames
	params["dirs"] = dirs
	tmpl.Execute(w, params)

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

func GetFile(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		http.Error(w, "filename is required", http.StatusBadRequest)
		return
	}
	filename = strings.Trim(filename, "/")

	filePath := os.Getenv("ROOT_PATH") + "/" + filename
	bts, err := file.ReadFile(filePath)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	key := password.GetPassword()
	fileData, err := crypto.Decrypt(key, bts)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = w.Write(fileData)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
