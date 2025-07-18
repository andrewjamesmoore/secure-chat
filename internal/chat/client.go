package chat

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// - loads the server’s certificate and returns a TLS config with RootCAs
func loadClientTLSConfig(certPath string) *tls.Config {
	caCert, err := os.ReadFile(certPath)
	if err != nil {
		log.Fatalf("unable to read cert.pem: %v", err)
	}

	caPool := x509.NewCertPool()
	caPool.AppendCertsFromPEM(caCert)

	return &tls.Config{RootCAs: caPool}
}

// - handles reading and printing messages from the server
func listenForMessages(conn *tls.Conn) {
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Println("Disconnected by server")
			os.Exit(0)
		} else if err != nil {
			log.Printf("read error: %v", err)
			continue
		}
		fmt.Print(msg)
	}
}

// - handles user input and sending it to the server
func sendMessages(conn *tls.Conn) {
	stdin := bufio.NewScanner(os.Stdin)

	reader := bufio.NewReader(os.Stdin)
	username, _ := reader.ReadString('\n')
	fmt.Fprintln(conn, username)
	username = strings.TrimSpace(username)

	fmt.Printf("Welcome, %s! Begin typing to send secure messages:\n", username)

	for stdin.Scan() {
		text := stdin.Text()
		if _, err := fmt.Fprintln(conn, text); err != nil {
			log.Println("write error:", err)
		}
	}
}

// - connects to the server and starts listening
func ConnectClient(certPath string) {
	config := loadClientTLSConfig(certPath)

	host := os.Getenv("EC2_PUBLIC_IP")
	port := os.Getenv("PORT")

	if host == "" || port == "" {
		log.Fatal("Missing EC2_PUBLIC_IP or PORT in environment")
	}

	address := fmt.Sprintf("%s:%s", host, port)
	log.Println("Connecting to:", address)

	conn, err := tls.Dial("tcp", address, config)
	if err != nil {
		log.Fatalf("client dial error: %v", err)
	}
	defer conn.Close()

	log.Println("Connected to secure chat.")
	go listenForMessages(conn)
	sendMessages(conn)
}
