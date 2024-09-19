package model

import "github.com/google/uuid"

type User struct {
	Id    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Color string    `json:"color"`
}

func NewUser(name string, color string) *User {
	return &User{
		Id:    uuid.New(),
		Name:  name,
		Color: color,
	}
}
