package client

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID     uuid.UUID
	Name   string
	Color  string
	socket *websocket.Conn
}

func NewClient(name string, color string, socket *websocket.Conn) *Client {
	return &Client{
		ID:     uuid.New(),
		Name:   name,
		Color:  color,
		socket: socket,
	}
}

func (client *Client) Send(message []byte) error {
	return client.socket.WriteMessage(websocket.TextMessage, message)
}

func (client *Client) Read() ([]byte, error) {
	_, message, err := client.socket.ReadMessage()
	return message, err
}
