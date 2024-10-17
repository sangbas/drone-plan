package entity

import "github.com/google/uuid"

type Estate struct {
	ID     uuid.UUID
	Length int
	Width  int
}

func NewEstate(length int, width int) *Estate {
	return &Estate{
		ID:     uuid.New(),
		Length: length,
		Width:  width,
	}
}
