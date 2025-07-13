package main

import (
	"github.com/andrewjamesmoore/secure-chat/internal/chat"
)

func main() {
	chat.ConnectClient("certs/cert.pem", ":8443")
}
