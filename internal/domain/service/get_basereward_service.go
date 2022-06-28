package service

import (
	"chemi123/reward/internal/domain/data"
	repository "chemi123/reward/internal/domain/repository_inerface"
)

type BaseRewardService struct {
	baseRewardRepository repository.BaseRewardRepositoryInterface
}

func NewBaseRewardService(baseRewardRepository repository.BaseRewardRepositoryInterface) *BaseRewardService {
	return &BaseRewardService{
		baseRewardRepository: baseRewardRepository,
	}
}

func (s *BaseRewardService) GetBaseReward() (data.BaseReward, error) {
	return s.baseRewardRepository.GetById(1)
}
