package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemoryManager_Add(t *testing.T) {
	manager := NewInMemoryManager()
	client := NewMockClient()
	err := manager.Add(client)
	assert.NoError(t, err)
}

func TestInMemoryManager_Remove(t *testing.T) {
	manager := NewInMemoryManager()
	client := NewMockClient()
	err := manager.Add(client)
	assert.NoError(t, err)

	err = manager.Remove(client.ID)
	assert.NoError(t, err)
}

func TestInMemoryManager_Get(t *testing.T) {
	manager := NewInMemoryManager()
	client := NewMockClient()
	err := manager.Add(client)
	assert.NoError(t, err)

	retrievedClient := manager.Get(client.ID)
	assert.NotNil(t, retrievedClient)
	assert.Equal(t, client.ID, retrievedClient.GetID())
}
