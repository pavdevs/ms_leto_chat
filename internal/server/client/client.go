package client

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	user *User
	conn *websocket.Conn
}

func NewClient(user *User, conn *websocket.Conn) *Client {

	return &Client{
		user: user,
		conn: conn,
	}
}

func (c *Client) Send(message []byte) error {
	return c.conn.WriteMessage(websocket.TextMessage, message)
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) GetUser() *User {
	return c.user
}
