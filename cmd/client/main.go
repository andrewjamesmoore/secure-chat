package main

import (
	"log"

	"github.com/andrewjamesmoore/secure-chat/internal/chat"
)

func main() {
	log.Println("Connected to secure chat.")
	chat.ConnectClient("certs/cert.pem", ":8443")
}
