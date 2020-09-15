package main

import (
	"net"
)

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
