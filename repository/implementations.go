package repository

import (
	"context"
	"database/sql"
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
	_, err = r.Db.ExecContext(ctx, "INSERT INTO estate (id, length, width) VALUES ($1, $2, $3)", estate.Id, estate.Length, estate.Width)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateEstate: %v", err)
	}

	id = estate.Id

	return
}

func (r *Repository) CreateTree(ctx context.Context, tree *entity.Tree) (id uuid.UUID, err error) {
	_, err = r.Db.ExecContext(ctx, "INSERT INTO tree (id, estate_id, x_axis, y_axis, height) VALUES ($1, $2, $3, $4, $5)", tree.Id, tree.EstateId, tree.XAxis, tree.YAxis, tree.Height)
	if err != nil {
		return uuid.Nil, fmt.Errorf("CreateTree: %v", err)
	}

	id = tree.Id

	return
}

func (r *Repository) GetEstateById(ctx context.Context, id uuid.UUID) (estate *entity.Estate, err error) {
	var length, width int
	err = r.Db.QueryRowContext(ctx, "SELECT length, width FROM estate WHERE id = $1", id).Scan(&length, &width)
	if err != nil {
		return
	}
	estate = &entity.Estate{
		Id:     id,
		Length: length,
		Width:  width,
	}

	return
}

func (r *Repository) GetTreesByEstateId(ctx context.Context, id uuid.UUID) (trees []entity.Tree, err error) {
	rows, err := r.Db.QueryContext(ctx, "SELECT * FROM tree WHERE estate_id = $1", id)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var id, estateId uuid.UUID
		var xAxis, yAxis, height int
		err = rows.Scan(&id, &estateId, &xAxis, &yAxis, &height)
		if err != nil {
			return
		}
		tree := entity.Tree{
			Id:       id,
			EstateId: estateId,
			XAxis:    xAxis,
			YAxis:    yAxis,
			Height:   height,
		}
		trees = append(trees, tree)
	}

	return
}

func (r *Repository) GetEstateStat(ctx context.Context, id uuid.UUID) (estateStat EstateStat, err error) {
	var count int
	var max, min, median sql.NullInt32
	query := `SELECT count(height), max(height), min(height), PERCENTILE_CONT(0.5) WITHIN GROUP(ORDER BY height) 
				FROM tree t 
				WHERE estate_id = $1`
	err = r.Db.QueryRowContext(ctx, query, id).Scan(&count, &max, &min, &median)
	if err != nil {
		return
	}
	estateStat = EstateStat{
		Count:  count,
		Max:    int(max.Int32),
		Min:    int(min.Int32),
		Median: int(median.Int32),
	}

	return
}
