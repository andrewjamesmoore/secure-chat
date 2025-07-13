package chat

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
)

var (
	clients   = make(map[net.Conn]string)
	clientsMu sync.Mutex
)

// - returning the TLS configuration to be used by tls.Listen
func loadServerTLSConfig(certPath, keyPath string) *tls.Config {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		log.Fatalf("failed to load certs: %v", err)
	}
	return &tls.Config{Certificates: []tls.Certificate{cert}}
}

// - prompting the user to enter a username
func getUsername(conn net.Conn) (string, error) {
	conn.Write([]byte("Enter a username:\n"))
	reader := bufio.NewReader(conn)
	username, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(username), nil
}

// - storing the client connection and username in the shared map
func registerClient(conn net.Conn, username string) {
	clientsMu.Lock()
	clients[conn] = username
	clientsMu.Unlock()
}

// - reading messages from the connection
func handleMessages(conn net.Conn, username string) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		msg := scanner.Text()
		fullMsg := fmt.Sprintf("%s: %s\n", username, msg)
		log.Print(fullMsg)
		broadcast(fullMsg, conn)
	}
	
	if err := scanner.Err(); err != nil {
		log.Printf("client error (%s): %v", username, err)
	}
}

// - sending a message to all connected clients except the sender
func broadcast(message string, sender net.Conn) {
	clientsMu.Lock()
	defer clientsMu.Unlock()
	for client := range clients {
		if client != sender {
			client.Write([]byte(message))
		}
	}
}

// - removing the client from the shared map
func unregisterClient(conn net.Conn) {
	clientsMu.Lock()
	username := clients[conn]
	delete(clients, conn)
	clientsMu.Unlock()
	log.Printf("%s disconnected", username)
}

// - reading username, registering client, listening for messages
func handleClient(conn net.Conn) {
	defer func() {
		unregisterClient(conn)
		conn.Close()
		}()
		
		username, err := getUsername(conn)
		if err != nil {
			log.Println("username input error:", err)
			return
		}
		
		registerClient(conn, username)
		log.Printf("%s connected", username)
		
		handleMessages(conn, username)
	}
	
// - handling each client in a new goroutine
func RunServer(certPath, keyPath string, port string) {
	config := loadServerTLSConfig(certPath, keyPath)

	ln, err := tls.Listen("tcp", port, config)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	log.Printf("Server listening on port%s", port)
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		go handleClient(conn)
	}
}
