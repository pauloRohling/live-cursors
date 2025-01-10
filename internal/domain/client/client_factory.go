package client

import (
	"github.com/gorilla/websocket"
)

type Generator[T any] interface {
	Generate() (T, error)
}

type Factory interface {
	Create(conn *websocket.Conn) (*Client, error)
}

type DefaultFactory struct {
	nameGenerator  Generator[string]
	colorGenerator Generator[string]
}

func NewDefaultFactory(nameGenerator Generator[string], colorGenerator Generator[string]) *DefaultFactory {
	return &DefaultFactory{
		nameGenerator:  nameGenerator,
		colorGenerator: colorGenerator,
	}
}

func (factory *DefaultFactory) Create(conn *websocket.Conn) (*Client, error) {
	name, err := factory.nameGenerator.Generate()
	if err != nil {
		return nil, err
	}

	color, err := factory.colorGenerator.Generate()
	if err != nil {
		return nil, err
	}

	return NewClient(name, color, conn), nil
}
