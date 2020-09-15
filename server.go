package main

import (
	"fmt"
	"log"
	"net"
)

// Create a new user struct in the server.
func (s *Server) NewUser(conn net.Conn) {
	log.Printf("New User Conected: %s", conn.RemoteAddr())
	fmt.Fprintln(conn, startMsg)

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

// Function for sending message to users in the room.
func (s *Server) SendMessage(msg string, u *User) {
	for _, room := range s.Rooms {
		for _, user := range room.Users {
			if user == u {
				for _, rU := range room.Users {
					if rU != u {
						fmt.Fprintf(rU.Conn, "> %s: %s\n", u.Nick, msg)
					}
				}
			}
		}
	}
}

// Console command '/showrooms' show all rooms on the server.
func (s *Server) ShowRooms(u *User) {
	if len(s.Rooms) == 0 {
		fmt.Fprintln(u.Conn, "~ There's no rooms yet.")
	} else {
		for _, room := range s.Rooms {
			fmt.Fprintln(u.Conn, "- ", room.Name)
		}
	}
}

// To create a new room use '/mkroom' command.
func (s *Server) MakeRoom(name string) {
	room := &Room{
		Name: name,
	}
	s.Rooms = append(s.Rooms, room)
}

// Function which allows join the room.
func (s *Server) JoinRoom(name string, u *User) {
	for i, room := range s.Rooms {
		if room.Name == name {
			s.Rooms[i].Users = append(s.Rooms[i].Users, u)
			return
		}
	}
	fmt.Fprintln(u.Conn, "This room doesn't exist.")
}

// Function which delete user from the room.
func (s *Server) LeaveRoom(u *User) {
	for _, room := range s.Rooms {
		for i, roomUser := range room.Users {
			if roomUser == u {
				room.Users = append(room.Users[:i], room.Users[i+1:]...)
			}
		}
	}
}

// 'DeleteUser' delete user from server whan he use '/exit' command.
func (s *Server) DeleteUser(u *User) {
	s.LeaveRoom(u)
	for i, user := range s.Users {
		if user == u {
			s.Users = append(s.Users[:i], s.Users[i+1:]...)
		}
	}
}
