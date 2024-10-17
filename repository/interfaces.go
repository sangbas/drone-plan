// This file contains the interfaces for the repository layer.
// The repository layer is responsible for interacting with the database.
// For testing purpose we will generate mock implementations of these
// interfaces using mockgen. See the Makefile for more information.
package repository

import (
	"context"

	"github.com/SawitProRecruitment/UserService/entity"
	"github.com/google/uuid"
)

type RepositoryInterface interface {
	GetTestById(ctx context.Context, input GetTestByIdInput) (output GetTestByIdOutput, err error)
	CreateEstate(ctx context.Context, estate *entity.Estate) (id uuid.UUID, err error)
	CreateTree(ctx context.Context, tree *entity.Tree) (id uuid.UUID, err error)
	GetEstateById(ctx context.Context, id uuid.UUID) (estate *entity.Estate, err error)
	GetTreesByEstateId(ctx context.Context, id uuid.UUID) (trees []entity.Tree, err error)
	GetEstateStat(ctx context.Context, id uuid.UUID) (estateStat EstateStat, err error)
}
