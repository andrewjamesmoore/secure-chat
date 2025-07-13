# 🧪 Secure Chat – Experimental TLS Chat Server in Go

### A simple secure chat server and client built with Go using TLS.

Weekend learning experiment to explore low-level networking in Go — specifically using `net` and `tls` to create a secure, terminal-based chat server and client.

### What it does

- TLS-encrypted server/client communication
- Prompts users for a username
- Broadcasts messages to all connected users
- Built with `net`, `tls`, and Go’s standard library — no external packages

### What I Learned

- How to work with raw TCP sockets in Go (net.Conn, tls.Conn).
- The difference between SSL and TLS and why Common Name (CN) is deprecated in favor of Subject Alternative Name (SAN).
- Importance of encryption (inspecting network traffic with TLS on and off to see the difference).
- How to generate and use self-signed TLS certificates properly.
- How frustrating certs, Nginx config, and reverse proxies can be when juggling local and remote development.
- Structuring a Go project
