package message

import (
	"iter"
	"live-cursors/internal/domain/client"
	"live-cursors/internal/model"
	"live-cursors/pkg/json"
)

type Producer struct {
	manager client.Manager
}

func NewProducer(manager client.Manager) *Producer {
	return &Producer{manager: manager}
}

func (producer *Producer) Position(client client.Client, positionMessage []byte) error {
	for otherClient := range producer.getAllExcept(client) {
		if err := otherClient.Send(positionMessage); err != nil {
			return err
		}
	}

	return nil
}

func (producer *Producer) Self(client client.Client) error {
	message := model.NewMessage(client, model.MessageTypeSelf)
	payload, err := json.Serialize(message)
	if err != nil {
		return err
	}
	return client.Send(payload)
}

func (producer *Producer) Client(client client.Client) error {
	message := model.NewMessage(client, model.MessageTypeClient)
	payload, err := json.Serialize(message)
	if err != nil {
		return err
	}

	for otherClient := range producer.getAllExcept(client) {
		if err = otherClient.Send(payload); err != nil {
			return err
		}
	}

	return nil
}

func (producer *Producer) Remove(client client.Client) error {
	message := model.NewMessage(client, model.MessageTypeRemove)
	payload, err := json.Serialize(message)
	if err != nil {
		return err
	}

	for otherClient := range producer.getAllExcept(client) {
		if err = otherClient.Send(payload); err != nil {
			return err
		}
	}

	return nil
}

func (producer *Producer) CurrentClients(client client.Client) error {
	for otherClient := range producer.getAllExcept(client) {
		message := model.NewMessage(otherClient, model.MessageTypeClient)
		payload, err := json.Serialize(message)
		if err != nil {
			return err
		}

		if err = client.Send(payload); err != nil {
			return err
		}
	}

	return nil
}

func (producer *Producer) getAllExcept(ignoredClient client.Client) iter.Seq[client.Client] {
	clients := producer.manager.GetAll()

	return func(yield func(client.Client) bool) {
		for _, otherClient := range clients {
			if ignoredClient != nil && otherClient.GetID() == ignoredClient.GetID() {
				continue
			}

			if !yield(otherClient) {
				return
			}
		}
	}
}
