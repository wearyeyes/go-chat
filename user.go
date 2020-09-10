package main

import (
	"fmt"
)

func (u *User) showRooms() {
	for _, room := range server.Rooms {
		fmt.Fprintln(u.Connection, room.Name)
	}
}