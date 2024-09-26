package client

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type WebSocketClient struct {
	ID    uuid.UUID
	Name  string
	Color string
	conn  *websocket.Conn
}

func NewWebSocketClient(name string, color string, socket *websocket.Conn) *WebSocketClient {
	return &WebSocketClient{
		ID:    uuid.New(),
		Name:  name,
		Color: color,
		conn:  socket,
	}
}

func (client *WebSocketClient) GetID() uuid.UUID {
	return client.ID
}

func (client *WebSocketClient) Send(message []byte) error {
	return client.conn.WriteMessage(websocket.TextMessage, message)
}

func (client *WebSocketClient) Read() ([]byte, error) {
	_, message, err := client.conn.ReadMessage()
	return message, err
}

func (client *WebSocketClient) Close() error {
	return client.conn.Close()
}
