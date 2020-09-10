package main

import (
	"net"
)

var ()

var server = Server{}

var startMsg = `
   Hello!
   Chat have this commands:
   - \setnick
   - \join
   - \rmroom
   - \showrooms
   - \msg
   - \exit
   - \help
`

type Server struct {
	Users []User
	Rooms []*Room
}

type User struct {
	Connection net.Conn
	Nick       string
}

type Room struct {
	Name  string
	Users []*User
}
