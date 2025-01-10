package client

import (
	"fmt"
	"github.com/google/uuid"
	"live-cursors/internal/domain/client"
	"maps"
	"slices"
	"sync"
)

type InMemoryManager struct {
	clients map[uuid.UUID]*client.Client
	mutex   *sync.Mutex
}

func NewInMemoryManager() *InMemoryManager {
	return &InMemoryManager{
		clients: make(map[uuid.UUID]*client.Client),
		mutex:   &sync.Mutex{},
	}
}

func (manager *InMemoryManager) Add(client *client.Client) error {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	if _, ok := manager.clients[client.GetID()]; ok {
		return fmt.Errorf("client already exists")
	}

	manager.clients[client.GetID()] = client
	return nil
}

func (manager *InMemoryManager) Remove(id uuid.UUID) error {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	if aClient, ok := manager.clients[id]; ok {
		if err := aClient.Close(); err != nil {
			return err
		}
		delete(manager.clients, id)
	}

	return nil
}

func (manager *InMemoryManager) Get(id uuid.UUID) *client.Client {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	return manager.clients[id]
}

func (manager *InMemoryManager) GetAll() []*client.Client {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	return slices.Collect(maps.Values(manager.clients))
}

func (manager *InMemoryManager) Broadcast(message []byte, ignoreId *uuid.UUID) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	for id, aClient := range manager.clients {
		if ignoreId != nil && *ignoreId == id {
			continue
		}

		// If there is an error sending the message, assume the client
		// has disconnected and remove it from the manager
		if err := aClient.Send(message); err != nil {
			delete(manager.clients, id)
		}
	}
}
