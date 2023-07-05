package models

import "github.com/andreluizmicro/go-soccer/internal/manegment/domain/entity"

type StadiumInterface interface {
	Create(stadium *entity.Stadium) error
	FindById(id string) (*entity.Stadium, error)
}
