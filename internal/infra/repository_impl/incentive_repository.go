package repository_impl

import (
	"chemi123/reward/internal/domain/entity"
	"database/sql"
)

type IncentiveRepository struct {
	db *sql.DB
}

func NewIncentiveRepository(db *sql.DB) *IncentiveRepository {
	return &IncentiveRepository{
		db: db,
	}
}

func (repository *IncentiveRepository) GetById(id int32) (entity.Incenntive, error) {
	var incentive entity.Incenntive

	row := repository.db.QueryRow("SELECT id, fee FROM incentive WHERE id = ?", id)
	if err := row.Scan(&incentive.Id, &incentive.Fee); err != nil {
		return incentive, err
	}

	return incentive, nil
}
