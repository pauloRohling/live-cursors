package client

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRandomFactory_Create(t *testing.T) {
	nameGenerator := NewMockGenerator("Name")
	colorGenerator := NewMockGenerator("Color")
	factory := NewRandomFactory(nameGenerator, colorGenerator)

	conn := &websocket.Conn{}
	client, err := factory.Create(conn)
	assert.NoError(t, err)
	assert.NotNil(t, client.GetID())
}
