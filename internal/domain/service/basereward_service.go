package service

import (
	"chemi123/reward/internal/domain/entity"
	repository "chemi123/reward/internal/domain/repository_inerface"
)

type BaseRewardService struct {
	repository repository.BaseRewardRepositoryInterface
}

func NewBaseRewardService(repository repository.BaseRewardRepositoryInterface) *BaseRewardService {
	return &BaseRewardService{
		repository: repository,
	}
}

func (s *BaseRewardService) GetBaseReward() (entity.BaseReward, error) {
	return s.repository.GetById(1)
}
