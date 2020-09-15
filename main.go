package main

import (
	"log"
	"net"
)

func main() {
	s := &Server{}

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

		go s.NewUser(conn)
	}
}
