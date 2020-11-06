package controllers

import (
	"fileprotector/services/crypto"
	"fileprotector/services/file"
	"fileprotector/services/password"
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

	selectedDir := strings.Trim(r.URL.Query().Get("dir"), "/")

	target := os.Getenv("ROOT_PATH")
	if selectedDir != "" {
		target += "/" + selectedDir
	}
	files, err := ioutil.ReadDir(target)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
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
	directoryPath := os.Getenv("ROOT_PATH")

	r.ParseForm()
	selectedDir := r.FormValue("selectedDir")
	if selectedDir != "" {
		directoryPath += "/" + selectedDir
	}
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
	if selectedDir == "" {
		_, _ = w.Write([]byte(`<meta http-equiv="refresh" content="1; url = /" />`))
	} else {
		_, _ = w.Write([]byte(`<meta http-equiv="refresh" content="1; url = /?dir=` + selectedDir + `" />`))
	}
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
