package presentation

import (
	"github.com/gorilla/websocket"
	"live-cursors/internal/domain/client"
	"log"
	"net/http"
)

type WebSocketHandler struct {
	factory  client.Factory
	manager  client.Manager
	producer client.Producer
	upgrader *websocket.Upgrader
}

func NewWebSocketHandler(factory client.Factory, manager client.Manager, producer client.Producer) *WebSocketHandler {
	upgrader := &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return &WebSocketHandler{
		factory:  factory,
		manager:  manager,
		producer: producer,
		upgrader: upgrader,
	}
}

func (handler *WebSocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := handler.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newClient, err := handler.factory.Create(conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = handler.manager.Add(newClient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func(newClient client.Client) {
		if err = handler.manager.Remove(newClient.GetID()); err != nil {
			log.Printf("Error during closing connection: %s", err.Error())
		}

		// TODO Send a message to all clients to remove the client
	}(newClient)

	if err = handler.sendMessages(newClient); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Println("Client Connected as ", newClient.GetID())
	handler.listenPositions(newClient)
}

func (handler *WebSocketHandler) sendMessages(newClient client.Client) error {
	if err := handler.producer.ProduceSelf(newClient); err != nil {
		return err
	}

	if err := handler.producer.ProduceUser(newClient); err != nil {
		return err
	}

	return handler.producer.ProduceCurrentUsers(newClient)
}

func (handler *WebSocketHandler) listenPositions(newClient client.Client) {
	for {
		rawPosition, err := newClient.Read()
		if err != nil {
			log.Println(err)
			return
		}

		if err = handler.producer.ProducePosition(newClient, rawPosition); err != nil {
			log.Println(err)
			return
		}
	}
}
