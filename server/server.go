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

	router.HandleFunc("/", controllers.Index).Methods(http.MethodGet)
	router.HandleFunc("/upload", controllers.Upload).Methods(http.MethodPost)
	router.HandleFunc("/getfile", controllers.GetFile).Methods(http.MethodGet)
	router.HandleFunc("/mkdir", controllers.Mkdir).Methods(http.MethodPost)
	router.HandleFunc("/del", controllers.DeletePath).Methods(http.MethodDelete)

	fs := http.FileServer(http.Dir("./frontend/dist/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

	fmt.Println("server started at: http://127.0.0.1:1313")
	log.Fatalln(http.ListenAndServe(":1313", router))
}
