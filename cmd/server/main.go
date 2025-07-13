package main

import (
	"log"

	"github.com/andrewjamesmoore/secure-chat/internal/chat"
)

var port string = ":8443"

func main() {
	log.Println("Server listening on port %s", port)
	chat.RunServer("certs/cert.pem", "certs/key.pem", port)
}
