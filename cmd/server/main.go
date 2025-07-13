package main

import (
	"github.com/andrewjamesmoore/secure-chat/internal/chat"
)

func main() {
	chat.RunServer("certs/cert.pem", "certs/key.pem", ":8443")
}
