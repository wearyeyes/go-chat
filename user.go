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
		case strings.HasPrefix(msg, "/msg"):
			s.SendMessage(strings.SplitN(msg, " ", 2)[1], u)
		case msg == "/showrooms":
			s.ShowRooms(u)
		case strings.HasPrefix(msg, "/mkroom"):
			smsg := strings.SplitN(msg, " ", 2)
			s.MakeRoom(smsg[1])
		case strings.HasPrefix(msg, "/join"):
			s.JoinRoom(strings.SplitN(msg, " ", 2)[1], u)
		case msg == "/leaveroom":
			s.LeaveRoom(u)
		case msg == "/exit":
			log.Printf("User %q has disconnected", u.Nick)
			s.DeleteUser(u)
			u.Conn.Close()
			return
		}
	}
}
