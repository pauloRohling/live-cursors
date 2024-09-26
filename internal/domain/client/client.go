package client

import "github.com/google/uuid"

type Client interface {
	GetID() uuid.UUID
	Send(message []byte) error
	Read() ([]byte, error)
	Close() error
}
