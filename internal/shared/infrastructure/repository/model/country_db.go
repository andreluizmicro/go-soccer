package model

import (
	"github.com/andreluizmicro/go-soccer/internal/shared/domain/entity"
	"gorm.io/gorm"
)

type CountryModel struct {
	DB *gorm.DB
}

func NewCountryModel(db *gorm.DB) *CountryModel {
	return &CountryModel{
		DB: db,
	}
}

func (model *CountryModel) Create(country *entity.Country) error {
	return model.DB.Create(country).Error
}

func (model *CountryModel) FindAll(page, limit int, sort string) ([]entity.Country, error) {
	var countries []entity.Country
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = model.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at" + sort).Find(countries).Error
	} else {
		err = model.DB.Order("created_at" + sort).Find(countries).Error
	}
	return countries, err
}

func (model *CountryModel) FindByID(id string) (*entity.Country, error) {
	var country entity.Country
	err := model.DB.First(&country, "id = ?", id).Error
	return &country, err
}

func (model *CountryModel) Update(country entity.Country) error {
	_, err := model.FindByID(country.ID.String())
	if err != nil {
		return err
	}
	return model.DB.Save(country).Error
}

func (model *CountryModel) Delete(id string) error {
	country, err := model.FindByID(id)
	if err != nil {
		return err
	}
	return model.DB.Delete(country).Error
}
