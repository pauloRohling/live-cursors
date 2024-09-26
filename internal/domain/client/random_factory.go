package client

import (
	"github.com/gorilla/websocket"
	"live-cursors/internal/model"
)

type RandomFactory struct {
	nameGenerator  Generator[string]
	colorGenerator Generator[string]
}

func NewRandomFactory(nameGenerator Generator[string], colorGenerator Generator[string]) *RandomFactory {
	return &RandomFactory{
		nameGenerator:  nameGenerator,
		colorGenerator: colorGenerator,
	}
}

func (factory *RandomFactory) Create(conn *websocket.Conn) (Client, error) {
	name, err := factory.nameGenerator.Generate()
	if err != nil {
		return nil, err
	}

	color, err := factory.colorGenerator.Generate()
	if err != nil {
		return nil, err
	}

	return model.NewClient(name, color, conn), nil
}
