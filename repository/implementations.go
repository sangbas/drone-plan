package repository

import (
	"context"
	"fmt"

	"github.com/SawitProRecruitment/UserService/entity"
	"github.com/google/uuid"
)

func (r *Repository) GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error) {
	err = r.Db.QueryRowContext(ctx, "SELECT name FROM test WHERE id = $1", input.Id).Scan(&output.Name)
	if err != nil {
		return
	}
	return
}

func (r *Repository) CreateEstate(ctx context.Context, estate *entity.Estate) (id uuid.UUID, err error) {
	_, err = r.Db.ExecContext(ctx, "INSERT INTO estate (id, length, width) VALUES ($1, $2, $3)", estate.ID, estate.Length, estate.Width)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateEstate: %v", err)
	}

	id = estate.ID

	return
}

func (r *Repository) CreateTree(ctx context.Context, tree entity.Tree) (id uuid.UUID, err error) {
	return
}
