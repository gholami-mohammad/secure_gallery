package main

import (
	"fileprotector/server"
	"fileprotector/services/password"
	"log"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.Lshortfile)
	godotenv.Load()
}

func main() {
	pass := password.ReadPasswordInput()
	password.CheckPassword(pass)
	server.Run()
}
