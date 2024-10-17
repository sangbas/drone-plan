package entity

import "github.com/google/uuid"

type Tree struct {
	Id       uuid.UUID
	EstateId uuid.UUID
	XAxis    int
	YAxis    int
	Height   int
}

func NewTree(estateId uuid.UUID, xAxis, yAxis, height int) *Tree {
	return &Tree{
		Id:       uuid.New(),
		EstateId: estateId,
		XAxis:    xAxis,
		YAxis:    yAxis,
		Height:   height,
	}
}
