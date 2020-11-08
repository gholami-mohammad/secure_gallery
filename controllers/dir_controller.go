package controllers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func Mkdir(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.FormValue("name")
	name = strings.Trim(name, "/")

	if name == "" {
		http.Error(w, "directory name is required", http.StatusUnprocessableEntity)
		return
	}

	directoryPath := os.Getenv("ROOT_PATH")
	target := directoryPath + "/" + name
	if _, err := os.Stat(target); err == nil {
		// directory exists
		http.Error(w, "direcory exists", http.StatusBadRequest)
		return
	}
	err := os.MkdirAll(target, os.ModePerm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	info, _ := os.Stat(target)

	var resp struct {
		ModTime string `json:"mod_time"`
	}
	resp.ModTime = info.ModTime().Format("2006-01-02")

	bts, err := json.Marshal(&resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bts)

}
