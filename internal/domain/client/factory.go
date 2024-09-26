package client

import (
	"github.com/gorilla/websocket"
)

type Factory interface {
	Create(conn *websocket.Conn) (Client, error)
}
