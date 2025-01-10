package client

import (
	"github.com/stretchr/testify/assert"
	"live-cursors/internal/domain/client"
	"testing"
)

type MockConnection struct {
}

func (conn *MockConnection) ReadMessage() (messageType int, p []byte, err error) {
	return 0, nil, nil
}

func (conn *MockConnection) WriteMessage(messageType int, data []byte) error {
	return nil
}

func (conn *MockConnection) Close() error {
	return nil
}

var connection = &MockConnection{}

func TestInMemoryManager_Add(t *testing.T) {
	manager := NewInMemoryManager()
	newClient := client.NewClient("Test", "#000FFF", connection)
	err := manager.Add(newClient)
	assert.NoError(t, err)
}

func TestInMemoryManager_Remove(t *testing.T) {
	manager := NewInMemoryManager()
	newClient := client.NewClient("Test", "#000FFF", connection)
	err := manager.Add(newClient)
	assert.NoError(t, err)

	err = manager.Remove(newClient.ID)
	assert.NoError(t, err)
}

func TestInMemoryManager_Get(t *testing.T) {
	manager := NewInMemoryManager()
	newClient := client.NewClient("Test", "#000FFF", connection)
	err := manager.Add(newClient)
	assert.NoError(t, err)

	retrievedClient := manager.Get(newClient.ID)
	assert.NotNil(t, retrievedClient)
	assert.Equal(t, newClient.ID, retrievedClient.GetID())
}
