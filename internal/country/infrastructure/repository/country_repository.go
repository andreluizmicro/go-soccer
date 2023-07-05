package repository

import (
	"database/sql"
	"fmt"

	"github.com/andreluizmicro/go-soccer/internal/country/domain/entity"
)

type CountryRepository struct {
	DB *sql.DB
}

func NewCountryRepository(db *sql.DB) *CountryRepository {
	return &CountryRepository{
		DB: db,
	}
}

func (repository *CountryRepository) Create(country *entity.Country) error {
	fmt.Println("TSTEEEEE")
	return nil
}

func (repository *CountryRepository) FindAll(page, limit int, sort string) ([]entity.Country, error) {
	var countries []entity.Country
	return countries, nil
	// var countries []entity.Country
	// var err error
	// if sort != "" && sort != "asc" && sort != "desc" {
	// 	sort = "asc"
	// }
	// if page != 0 && limit != 0 {
	// 	err = repository.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at" + sort).Find(countries).Error
	// } else {
	// 	err = repository.DB.Order("created_at" + sort).Find(countries).Error
	// }
	// return countries, err
}

func (repository *CountryRepository) FindByID(id string) (*entity.Country, error) {
	var country entity.Country
	return &country, nil
	// err := repository.DB.First(&country, "id = ?", id).Error
	// return &country, err
}

func (repository *CountryRepository) Update(country *entity.Country) error {
	return nil
	// _, err := repository.FindByID(country.ID.String())
	// if err != nil {
	// 	return err
	// }
	// return repository.DB.Save(country).Error
}

func (repository *CountryRepository) Delete(id string) error {
	return nil
	// country, err := repository.FindByID(id)
	// if err != nil {
	// 	return err
	// }
	// return repository.DB.Delete(country).Error
}
