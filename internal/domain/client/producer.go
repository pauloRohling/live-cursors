package client

type Producer interface {
	ProducePosition(client Client, positionMessage []byte) error
	ProduceSelf(client Client) error
	ProduceUser(client Client) error
	ProduceCurrentUsers(client Client) error
}
