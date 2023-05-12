package main

import (
	"fmt"
	"log"
	"net"
)

// Telnet Server

func handleConnection(conn net.Conn) error {
	log.Println("Connection Established")

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return err
		}
		if n == 0 {
			log.Println("Zero bytes, closing connection")
			break
		}

		msg := buf[0 : n-2]
		log.Println("Received message:", (msg))

		resp := fmt.Sprintf("You Said, \"%s\"\r\n", msg)
		n, err = conn.Write([]byte(resp))
		if err != nil {
			return err
		}
		if n == 0 {
			log.Println("Zero bytes, closing connection")
			break
		}
	}

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
