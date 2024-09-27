package client

import (
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type InMemoryManager struct {
	clients map[uuid.UUID]Client
	mutex   *sync.Mutex
}

func NewInMemoryManager() *InMemoryManager {
	return &InMemoryManager{
		clients: make(map[uuid.UUID]Client),
		mutex:   &sync.Mutex{},
	}
}

func (manager *InMemoryManager) Add(client Client) error {
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

	if client, ok := manager.clients[id]; ok {
		if err := client.Close(); err != nil {
			return err
		}
		delete(manager.clients, id)
	}

	return nil
}

func (manager *InMemoryManager) Get(id uuid.UUID) Client {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()
	return manager.clients[id]
}

func (manager *InMemoryManager) GetAll() []Client {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	clients := make([]Client, len(manager.clients))

	i := 0
	for _, client := range manager.clients {
		clients[i] = client
		i++
	}

	return clients
}

func (manager *InMemoryManager) Broadcast(message []byte, ignoreId *uuid.UUID) {
	manager.mutex.Lock()
	defer manager.mutex.Unlock()

	for id, client := range manager.clients {
		if ignoreId != nil && *ignoreId == id {
			continue
		}

		// If there is an error sending the message, assume the client
		// has disconnected and remove it from the manager
		if err := client.Send(message); err != nil {
			delete(manager.clients, id)
		}
	}
}
