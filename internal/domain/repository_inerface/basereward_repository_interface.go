package repository_inerface

import "chemi123/reward/internal/domain/entity"

type BaseRewardRepositoryInterface interface {
	GetById(id int32) (entity.BaseReward, error)
}
