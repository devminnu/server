package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
)

func main() {
	// Create a simple handler
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("request received")
		w.Write([]byte("Hello, HTTP/3 World!"))
		log.Println("response sent")
	})

	// Load TLS certificates (self-signed or from a CA)
	certFile := "cert.pem"
	keyFile := "key.pem"
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// TLS configuration for HTTP/2 and HTTP/3
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		NextProtos:   []string{"h3", "http/1.1"}, // HTTP/3 and HTTP/1.1 protocols
	}

	// Create the HTTP/3 server
	server := http3.Server{
		Addr:      ":443",    // Listen on port 443 for HTTPS/QUIC traffic
		Handler:   mux,       // Your HTTP handler (mux)
		TLSConfig: tlsConfig, // TLS configuration
	}

	// Start the HTTP/3 server
	log.Println("Starting HTTP/3 server on https://gamesinfotech.com")
	err = server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
	}
}
