package main

import (
	"log"
	"net"
)

// Telnet Server

func handleConnection(conn net.Conn) error {
	log.Println("Connection Established")
	return nil
}
func StartServer() error {
	log.Println("Starting Server")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		return err
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection", err)
			continue
		}
		go func() {
			if err := handleConnection(conn); err != nil {
				log.Println("Error handling exception", err)
				return
			}
		}()
	}
}

func main() {
	err := StartServer()
	if err != nil {
		log.Fatal(err)
	}
}
