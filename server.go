package main

import (
	"fmt"
	"log"
	"net"
)

func (s *Server) NewUser(conn net.Conn) {
	log.Printf("New User Conected: %s", conn.RemoteAddr())

	var nick string
	fmt.Fprint(conn, "\nEnter your nickname: ")
	fmt.Fscanln(conn, &nick)

	user := &User{
		Conn: conn,
		Nick: nick,
	}

	s.Users = append(s.Users, user)

	user.ReadInput(s)
}

func (s *Server) SendMessage(msg string, u *User) {
	for _, user := range s.Users {
		if user != u {
			fmt.Fprintf(user.Conn, "> %s: %s\n", u.Nick, msg)
		}
	}
}