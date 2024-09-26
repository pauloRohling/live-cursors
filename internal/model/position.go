package model

import "github.com/google/uuid"

type Position struct {
	ID uuid.UUID `json:"id"`
	X  int       `json:"x"`
	Y  int       `json:"y"`
}
