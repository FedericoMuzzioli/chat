package main

import (
	"log"
	"net"
)

func handleConnection(_conn net.Conn) {
	n, err := _conn.Write("Hi")
	if err != nil {
		log.Printf("Could not wrtite to: %1", _conn.RemoteAddr())
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("ERROR : %1", err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("ERROR temp: %1", err)
			// handle error
		}
		go handleConnection(conn)
	}
}
