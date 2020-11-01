package server

import (
	"fileprotector/controllers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()

	router.HandleFunc("/", controllers.Index)
	router.HandleFunc("/upload", controllers.Upload)

	fs := http.FileServer(http.Dir("./frontend/dist/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	fmt.Println("server started at: http://127.0.0.1:1313")
	log.Fatalln(http.ListenAndServe(":1313", router))
}
