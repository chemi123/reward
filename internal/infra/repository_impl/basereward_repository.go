package repository_impl

import (
	"chemi123/reward/internal/domain/data"
	"database/sql"
)

type BaseRewardRepository struct {
	db *sql.DB
}

func NewBaseRewardRepository(db *sql.DB) *BaseRewardRepository {
	return &BaseRewardRepository{
		db: db,
	}
}

func (repository BaseRewardRepository) GetById(id int32) (data.BaseReward, error) {
	var basereward data.BaseReward

	row := repository.db.QueryRow("SELECT id, message FROM basereward WHERE id = ?", id)
	if err := row.Scan(&basereward.Id, &basereward.Message); err != nil {
		return basereward, err
	}

	return basereward, nil
}
