package controllers

import (
	"log"
	"net/http"
	"os"
	"strings"
)

func DeletePath(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	path = strings.Trim(path, "/")

	if path == "" {
		http.Error(w, "file or directory name is required", http.StatusUnprocessableEntity)
		return
	}

	directoryPath := os.Getenv("ROOT_PATH")
	target := directoryPath + "/" + path
	_, err := os.Stat(target)
	if err != nil {
		http.Error(w, "no such file or directory", http.StatusNotFound)
		return
	}

	err = os.RemoveAll(target)
	if err != nil {
		log.Println(err)
		http.Error(w, "unable to remove the file or directory", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("file/directory deleted"))
}
