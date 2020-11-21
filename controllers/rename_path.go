package controllers

import (
	"net/http"
	"os"
	"strings"
)

func Rename(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	oldPath := r.Form.Get("old_path")
	if oldPath == "" {
		http.Error(w, "old_path is required", http.StatusBadRequest)
		return
	}
	oldPath = strings.Trim(oldPath, "/")

	newPath := r.Form.Get("new_path")
	if newPath == "" {
		http.Error(w, "New Name is required", http.StatusBadRequest)
		return
	}
	newPath = strings.Trim(newPath, "/")

	oldPath = os.Getenv("ROOT_PATH") + "/" + oldPath
	newPath = os.Getenv("ROOT_PATH") + "/" + newPath
	if oldPath == newPath {
		_, _ = w.Write([]byte("done"))
		return
	}

	if _, err := os.Stat(oldPath); err != nil {
		http.Error(w, "No such file or directory", http.StatusNotFound)
		return
	}

	if _, err := os.Stat(newPath); err == nil {
		// file or directory exists with the same name
		http.Error(w, "file or directory exists with the same name", http.StatusUnprocessableEntity)
		return
	}
	err := os.Rename(oldPath, newPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = w.Write([]byte("done"))
}
