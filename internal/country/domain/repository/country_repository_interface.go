package repository

import "github.com/andreluizmicro/go-soccer/internal/country/domain/entity"

type CountryRepositoryInterface interface {
	Create(country *entity.Country) error
	FindAll(page, limit int, sort string) ([]entity.Country, error)
	FindByID(id string) (*entity.Country, error)
	Update(country *entity.Country) error
	Delete(id string) error
}
