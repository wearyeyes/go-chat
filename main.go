package main

import (
	"log"
	"net"
)

func handleConnection(c net.Conn) {
	c.Close()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Println("Server started!")

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}
