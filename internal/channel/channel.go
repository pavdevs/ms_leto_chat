package channel

import "golang.org/x/net/websocket"

type Channel struct {
	Conn *websocket.Conn
	User *User
}

func NewChannel(conn *websocket.Conn, user *User) *Channel {
	return &Channel{
		Conn: conn,
		User: user,
	}
}
