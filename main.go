package main

import (
	"log"
	"net"
)

func main() {
	// Create a server.
	s := &Server{}

	// Create tcp-connection.
	l, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	log.Println("Server started!")

	for {
		// Accept connection from new users.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Initialize parallel function for each user.
		go s.NewUser(conn)
	}
}
