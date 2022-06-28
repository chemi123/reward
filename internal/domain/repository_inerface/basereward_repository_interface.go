package repository_inerface

import "chemi123/reward/internal/domain/data"

type BaseRewardRepositoryInterface interface {
	GetById(id int32) (data.BaseReward, error)
}
