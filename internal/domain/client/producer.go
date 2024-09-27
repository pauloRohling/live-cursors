package client

type Producer interface {
	Position(client Client, positionMessage []byte) error
	Self(client Client) error
	Client(client Client) error
	Remove(client Client) error
	CurrentClients(client Client) error
}
