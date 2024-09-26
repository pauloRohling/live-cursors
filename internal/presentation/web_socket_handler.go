package presentation

import (
	"github.com/gorilla/websocket"
	"live-cursors/internal/domain/client"
	"live-cursors/internal/domain/generator"
	"live-cursors/internal/model"
	"log"
	"net/http"
)

type WebSocketHandler struct {
	nameGenerator  generator.Generator[string]
	colorGenerator generator.Generator[string]
	manager        client.Manager
	upgrader       *websocket.Upgrader
}

func NewWebSocketHandler(
	nameGenerator generator.Generator[string],
	colorGenerator generator.Generator[string],
	manager client.Manager,
) *WebSocketHandler {
	upgrader := &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	return &WebSocketHandler{
		nameGenerator:  nameGenerator,
		colorGenerator: colorGenerator,
		upgrader:       upgrader,
		manager:        manager,
	}
}

func (handler *WebSocketHandler) Handle(w http.ResponseWriter, r *http.Request) {
	conn, err := handler.upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newClient, err := handler.createClient(conn)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = handler.manager.Add(newClient)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer func(newClient *client.WebSocketClient) {
		if err = handler.manager.Remove(newClient.ID); err != nil {
			log.Printf("Error during closing connection: %s", err.Error())
		}
	}(newClient)

	message := model.NewMessage(newClient, model.MessageTypeUser)
	payload, err := serialize(message)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	handler.manager.Broadcast(payload, nil)

	log.Println("WebSocketClient Connected as ", newClient.ID)
	handler.reader(newClient)
}

func (handler *WebSocketHandler) createClient(conn *websocket.Conn) (*client.WebSocketClient, error) {
	name, err := handler.nameGenerator.Generate()
	if err != nil {
		return nil, err
	}

	color, err := handler.colorGenerator.Generate()
	if err != nil {
		return nil, err
	}

	return client.NewWebSocketClient(name, color, conn), nil
}

func (handler *WebSocketHandler) reader(newClient *client.WebSocketClient) {
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

		handler.manager.Broadcast(payload, &newClient.ID)
	}
}
