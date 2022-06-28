package service

import (
	"chemi123/reward/internal/domain/entity"
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

func (s *BaseRewardService) GetBaseReward() (entity.BaseReward, error) {
	return s.baseRewardRepository.GetById(1)
}
