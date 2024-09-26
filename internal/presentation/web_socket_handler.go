package presentation

import (
	"github.com/gorilla/websocket"
	"live-cursors/internal/domain/client"
	"live-cursors/internal/model"
	"log"
	"net/http"
)

type WebSocketHandler struct {
	factory  client.Factory
	manager  client.Manager
	upgrader *websocket.Upgrader
}

func NewWebSocketHandler(factory client.Factory, manager client.Manager) *WebSocketHandler {
	upgrader := &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return &WebSocketHandler{
		factory:  factory,
		upgrader: upgrader,
		manager:  manager,
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
	selfMessage := model.NewMessage(newClient, model.MessageTypeSelf)
	payload, err := serialize(selfMessage)
	if err != nil {
		return err
	}

	if err = newClient.Send(payload); err != nil {
		return err
	}

	userMessage := model.NewMessage(newClient, model.MessageTypeUser)
	payload, err = serialize(userMessage)
	if err != nil {
		return err
	}

	clientID := newClient.GetID()
	handler.manager.Broadcast(payload, &clientID)

	otherUsersInRoom := handler.manager.GetAll()
	for _, otherClient := range otherUsersInRoom {
		if otherClient.GetID() == newClient.GetID() {
			continue
		}

		otherUserMessage := model.NewMessage(otherClient, model.MessageTypeUser)
		payload, err = serialize(otherUserMessage)
		if err != nil {
			return err
		}

		if err = otherClient.Send(payload); err != nil {
			return err
		}
	}

	return nil
}

func (handler *WebSocketHandler) listenPositions(newClient client.Client) {
	clientID := newClient.GetID()

	for {
		rawPosition, err := newClient.Read()
		if err != nil {
			log.Println(err)
			return
		}

		position, err := deserialize[model.Position](rawPosition)
		if err != nil {
			log.Println(err)
			return
		}

		message := model.NewMessage(position, model.MessageTypePosition)
		payload, err := serialize(message)
		if err != nil {
			log.Println(err)
			return
		}

		handler.manager.Broadcast(payload, &clientID)
	}
}
