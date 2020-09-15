package main

import (
	"net"
)

var startMsg = `
	Hello!
	This chat has next commands:
	- /msg
	- /showrooms
	- /mkroom
	- /leaveroom
	- /exit
`

type Server struct {
	Users []*User
	Rooms []*Room
}

type User struct {
	Conn net.Conn
	Nick string
}

type Room struct {
	Name  string
	Users []*User
}
