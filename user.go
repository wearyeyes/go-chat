package main

import (
	"bufio"
	"log"
	"strings"
)

// This function requests for data input from each user.
func (u *User) ReadInput(s *Server) {
	reader := bufio.NewReader(u.Conn)
	for {
		// Read data input from console.
		msg, _ := reader.ReadString('\n')
		msg = strings.TrimSpace(msg)

		// Define the command that the user wants to use.
		switch {
		case strings.HasPrefix(msg, "/msg"):
			// Use only message without command by SplitN func.
			s.SendMessage(strings.SplitN(msg, " ", 2)[1], u)
		case msg == "/showrooms":
			s.ShowRooms(u)
		case strings.HasPrefix(msg, "/mkroom"):
			// Use only the name of room from console input.
			s.MakeRoom(strings.SplitN(msg, " ", 2)[1])
		case strings.HasPrefix(msg, "/join"):
			// Use only the name of room by SlitN func.
			s.JoinRoom(strings.SplitN(msg, " ", 2)[1], u)
		case msg == "/leaveroom":
			s.LeaveRoom(u)
		case msg == "/exit":
			// Message to the server console.
			log.Printf("User %q has disconnected", u.Nick)
			s.DeleteUser(u)
			u.Conn.Close()
			return
		}
	}
}
