package repository_interface

import "chemi123/reward/internal/domain/entity"

type IncentiveRepositoryInterface interface {
	GetById(id int32) (entity.Incenntive, error)
}
