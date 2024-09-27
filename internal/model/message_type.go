package model

type MessageType string

const (
	MessageTypePosition MessageType = "position"
	MessageTypeClient   MessageType = "client"
	MessageTypeSelf     MessageType = "self"
)
