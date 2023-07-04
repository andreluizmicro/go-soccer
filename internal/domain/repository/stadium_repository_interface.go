package models

import (
	domain_entity "github.com/andreluizmicro/go-soccer/internal/domain/entity"
)

type StadiumInterface interface {
	Create(stadium *domain_entity.Stadium) error
	FindById(id string) (*domain_entity.Stadium, error)
}
