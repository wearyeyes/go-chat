package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConnection(c net.Conn) {
	defer c.Close()
	log.Println("Conected")

	fmt.Fprintln(c, startMsg)
	
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

		for _, user := range server.Users{
			if user.Connection != c {
				fmt.Fprintln(user.Connection, msg)
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

		var user = User{
			Connection: conn,
		}

		server.Users = append(server.Users, user)

		go handleConnection(conn)
	}
}
