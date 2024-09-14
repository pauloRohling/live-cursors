package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Message struct {
	UserUuid string `json:"userUuid"`
	Point    Point  `json:"point"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[int]*websocket.Conn)

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	clientId := rand.Int()
	clients[clientId] = ws
	log.Println("Client Connected as ", clientId)

	defer delete(clients, clientId)
	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, payload, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		for _, otherConn := range clients {
			if otherConn == conn {
				continue
			}

			if err = otherConn.WriteMessage(messageType, payload); err != nil {
				log.Println(err)
				return
			}
		}
	}
}

func main() {
	fmt.Println("Starting server...")
	http.HandleFunc("/", wsEndpoint)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
