package model

type MessageType string

const (
	MessageTypePosition MessageType = "position"
	MessageTypeUser     MessageType = "user"
	MessageTypeSelf     MessageType = "self"
)
