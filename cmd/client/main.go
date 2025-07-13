package main

import (
	"github.com/andrewjamesmoore/secure-chat/internal/chat"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	certPath := "certs/cert.pem"
	chat.ConnectClient(certPath)
}
