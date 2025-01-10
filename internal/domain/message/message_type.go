package message

type Type string

const (
	PositionType Type = "position"
	ClientType   Type = "client"
	SelfType     Type = "self"
	RemoveType   Type = "remove"
)
