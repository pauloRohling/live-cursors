package message

import "time"

type Message[T any] struct {
	Data      T     `json:"data"`
	Type      Type  `json:"type"`
	Timestamp int64 `json:"timestamp"`
}

func NewMessage[T any](data T, messageType Type) *Message[T] {
	return &Message[T]{
		Data:      data,
		Type:      messageType,
		Timestamp: time.Now().UTC().UnixMilli(),
	}
}
