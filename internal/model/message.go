package model

import "time"

type MessageType string

const (
	MessageTypePosition MessageType = "position"
	MessageTypeUser     MessageType = "user"
)

type Message[T any] struct {
	Data      T           `json:"data"`
	Type      MessageType `json:"type"`
	Timestamp int64       `json:"timestamp"`
}

func NewMessage[T any](data T, messageType MessageType) *Message[T] {
	return &Message[T]{
		Data:      data,
		Type:      messageType,
		Timestamp: time.Now().UTC().UnixMilli(),
	}
}
