package repository

import (
	"github.com/andreluizmicro/go-soccer/internal/shared/domain/entity"
	"github.com/andreluizmicro/go-soccer/internal/shared/infrastructure/repository/model"
)

type CountryRepository struct {
	model *model.CountryModel
}

func NewCountryRepository(db *model.CountryModel) *CountryRepository {
	return &CountryRepository{
		model: db,
	}
}

func (repository *CountryRepository) Create(country *entity.Country) error {
	return repository.model.Create(country)
}

func (repository *CountryRepository) FindAll(page, limit int, sort string) ([]entity.Country, error) {
	return repository.model.FindAll(page, limit, sort)
}

func (repository *CountryRepository) FindByID(id string) (*entity.Country, error) {
	country, err := repository.model.FindByID(id)
	if err != nil {
		return nil, err
	}
	return country, nil
}

func (repository *CountryRepository) Update(country *entity.Country) error {
	return repository.model.Update(*country)
}

func (repository *CountryRepository) Delete(id string) error {
	return repository.model.Delete(id)
}
