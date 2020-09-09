package main

import (
	"log"
	"net"
	"bufio"
	"strings"
	"fmt"
)

var users []net.Conn

func handleConnection(c net.Conn) {
	defer c.Close()
	log.Println("Conected")
	for { 
		msg, err := bufio.NewReader(c).ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		msg = strings.TrimSpace(msg)
		if msg == `\quit` {
			c.Close()
			break
		}

		for _, user := range users {
			if user != c {
				fmt.Fprintln(user, msg)
			}
			
		}
	}
	
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

		users = append(users, conn)

		go handleConnection(conn)
	}
}
