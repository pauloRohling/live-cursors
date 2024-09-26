package client

import (
	"github.com/google/uuid"
)

type MockClient struct {
	ID uuid.UUID
}

func NewMockClient() *MockClient {
	return &MockClient{ID: uuid.New()}
}

func (client *MockClient) GetID() uuid.UUID {
	return client.ID
}

func (client *MockClient) Send(message []byte) error {
	return nil
}

func (client *MockClient) Read() ([]byte, error) {
	return nil, nil
}

func (client *MockClient) Close() error {
	return nil
}
