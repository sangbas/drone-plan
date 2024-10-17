// This file contains types that are used in the repository layer.
package repository

type GetTestByIdInput struct {
	Id string
}

type GetTestByIdOutput struct {
	Name string
}

type EstateStat struct {
	Count  int
	Max    int
	Min    int
	Median int
}
