package main

import (
	"log"
	"net"
)

type MessageType int

const (
	ClientConnected MessageType = iota + 1
	DeleteCLient
	NewMessage
)

type Client struct {
	conn     net.Conn
	outgoing chan string
}

type Message struct {
	Type    MessageType
	Conn    net.Conn
	Message string
}

func server(messages chan Message) {
	clients := map[string]net.Conn{}
	for {
		msg := <-messages
		switch msg.Type {
		case ClientConnected:
			clients[msg.Conn.RemoteAddr().String()] = msg.Conn
		case DeleteCLient:
			msg.Conn.Close()
			delete(clients, msg.Conn.RemoteAddr().String())
		case NewMessage:
			for _, conn := range clients {
				log.Printf("Read  wrtite from %s: %s", conn.RemoteAddr(), msg.Message)
				_, err := conn.Write([]byte(msg.Message))
				if err != nil {
					log.Printf("Could not wrtite to %s: %s", conn.RemoteAddr(), err)
				}
			}
		}
	}
}

func client(_conn net.Conn, messages chan Message) {
	buffer := make([]byte, 64)
	for {
		n, err := _conn.Read(buffer)
		if err != nil {
			messages <- Message{
				Type: DeleteCLient,
				Conn: _conn,
			}
			return
		}
		messages <- Message{
			Type:    NewMessage,
			Conn:    _conn,
			Message: string(buffer[0:n]),
		}
	}
}

const PORT = "8080"

func main() {
	ln, err := net.Listen("tcp", ":"+PORT)
	if err != nil {
		log.Fatalf("ERROR : %s", err)
	}
	log.Printf("Listening to TCP: %s", PORT)
	messages := make(chan Message)
	go server(messages)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("ERROR temp: %s", err)
			// handle error
		}
		log.Printf("Accepted Connection from: %s", conn.RemoteAddr())

		messages <- Message{
			Type: ClientConnected,
			Conn: conn,
		}

		go client(conn, messages)
	}
}
