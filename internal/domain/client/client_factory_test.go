package client

import (
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"testing"
)

type MockGenerator struct {
	result string
}

func NewMockGenerator(result string) *MockGenerator {
	return &MockGenerator{result: result}
}

func (generator MockGenerator) Generate() (string, error) {
	return generator.result, nil
}

func TestRandomFactory_Create(t *testing.T) {
	nameGenerator := NewMockGenerator("Name")
	colorGenerator := NewMockGenerator("Color")
	factory := NewDefaultFactory(nameGenerator, colorGenerator)

	conn := &websocket.Conn{}
	client, err := factory.Create(conn)
	assert.NoError(t, err)
	assert.NotNil(t, client.GetID())
}
