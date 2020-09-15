package main

import (
	"bufio"
	"log"
	"strings"
)

func (u *User) ReadInput(s *Server) {
	reader := bufio.NewReader(u.Conn)
	for {
		msg, _ := reader.ReadString('\n')

		msg = strings.TrimSpace(msg)

		switch {
		case msg == "/exit":
			log.Printf("User %q has disconnected", u.Nick)
			u.Conn.Close()
			return
		default: 
			s.SendMessage(msg, u) // Send messages to the users on server.
		}
	}
}
