package model

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Color string    `json:"color"`
	conn  *websocket.Conn
}

func NewClient(name string, color string, conn *websocket.Conn) *Client {
	return &Client{
		ID:    uuid.New(),
		Name:  name,
		Color: color,
		conn:  conn,
	}
}

func (client *Client) GetID() uuid.UUID {
	return client.ID
}

func (client *Client) Send(message []byte) error {
	return client.conn.WriteMessage(websocket.TextMessage, message)
}

func (client *Client) Read() ([]byte, error) {
	_, message, err := client.conn.ReadMessage()
	return message, err
}

func (client *Client) Close() error {
	return client.conn.Close()
}
