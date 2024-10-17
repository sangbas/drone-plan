package entity

import "github.com/google/uuid"

type Tree struct {
	id       uuid.UUID
	estateId uuid.UUID
	xAxis    int
	yAxis    int
	height   int
}

func NewTree(estateId uuid.UUID, xAxis, yAxis, height int) *Tree {
	return &Tree{
		id:       uuid.New(),
		estateId: estateId,
		xAxis:    xAxis,
		height:   height,
	}
}
