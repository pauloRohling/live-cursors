package client

import "github.com/google/uuid"

type Manager interface {
	Add(newClient *Client) error
	Remove(id uuid.UUID) error
	Get(id uuid.UUID) *Client
	GetAll() []*Client
	Broadcast(message []byte, ignoreId *uuid.UUID)
}
