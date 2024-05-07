package main

import (
	"log"
	"net"
)

func handleConnection(_conn net.Conn) {
	defer _conn.Close()
	message := []byte("Hi muzzio")
	n, err := _conn.Write(message)
	if err != nil {
		log.Printf("Could not wrtite to: %s", _conn.RemoteAddr())
	}
	if n < len(message) {
		log.Printf("Could not wrtite entire message: %d out of %d", n, len(message))
	}
}

const PORT = "8080"

func main() {
	ln, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("ERROR : %s", err)
	}
	log.Printf("Listening to TCP: %s", PORT)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("ERROR temp: %s", err)
			// handle error
		}
		go handleConnection(conn)
	}
}
