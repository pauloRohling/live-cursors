package client

import "github.com/google/uuid"

type Manager interface {
	Add(client Client) error
	Remove(id uuid.UUID) error
	Get(id uuid.UUID) Client
	Broadcast(message []byte, ignoreId uuid.UUID)
}
