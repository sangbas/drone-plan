package entity

import "github.com/google/uuid"

type Estate struct {
	Id     uuid.UUID
	Length int
	Width  int
}

func NewEstate(length int, width int) *Estate {
	return &Estate{
		Id:     uuid.New(),
		Length: length,
		Width:  width,
	}
}
