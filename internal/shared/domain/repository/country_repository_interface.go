package repository

import "github.com/andreluizmicro/go-soccer/internal/shared/domain/entity"

type CountryRepositoryInterface interface {
	Create(country *entity.Country) error
	FindAll(page, limit int, sort string) ([]entity.Country, error)
	FindByID(id string) (*entity.Country, error)
	Update(country *entity.Country) error
	Delete(id string) error
}
