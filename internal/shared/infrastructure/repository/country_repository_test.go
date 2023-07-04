package repository

import (
	"testing"

	"github.com/andreluizmicro/go-soccer/internal/shared/domain/entity"
	"github.com/andreluizmicro/go-soccer/internal/shared/infrastructure/repository/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestShouldCrateNewProduct(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.Country{})

	params := entity.Params{
		Name:         "Brasil",
		Capital:      "Brasília",
		Area:         "8.510.000 km²",
		Languages:    "Portuguese",
		Timezone:     "UTC",
		Continent:    "Americano",
		OfficalColor: "verde e amarelo",
		Population:   100000,
	}

	country, err := entity.NewCountry(params)
	assert.Nil(t, err)
	assert.NotEmpty(t, country.ID)
	CountryDB := model.NewCountryModel(db)
	err = CountryDB.Create(country)
	assert.NoError(t, err)
	assert.NotEmpty(t, country.ID)
}
