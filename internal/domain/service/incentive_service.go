package service

import (
	"chemi123/reward/internal/domain/entity"
	repository "chemi123/reward/internal/domain/repository_interface"
)

type IncentiveService struct {
	repository repository.IncentiveRepositoryInterface
}

func NewIncentiveService(repository repository.IncentiveRepositoryInterface) *IncentiveService {
	return &IncentiveService{
		repository: repository,
	}
}

func (s *IncentiveService) GetIncentive() (entity.Incenntive, error) {
	return s.repository.GetById(1)
}
