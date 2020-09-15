# TCP chat in golang

It's a simple console TCP chat with the ability to create rooms. 
It has a few commands for work:
- /msg [message]- to send a message to the users in the room.
- /showrooms - to see the list of available rooms on the server.
- /mkroom [name of room] - to create a new room.
- /join [name of room] - to join the room.
- /leaveroom - to leave the room.
- /exit - to quit chat.

This chat works on 8090 port and you can join to this using **telnet**. Just start the server (`go run .`) and in another terminals use `telnet localhost 8090` console line to connect to the server.
