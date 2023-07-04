package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const countryName = "Brasil"
const countryCapital = "Brasília"
const countryArea = "8.510.000 km²"
const countryLanguage = "Portuguese"
const countryPopulation = 10000

func TestNewCountry(t *testing.T) {
	params := Params{
		Name:         "Brasil",
		Capital:      "Brasília",
		Area:         "8.510.000 km²",
		Languages:    "Portuguese",
		Timezone:     "UTC",
		Continent:    "Americano",
		OfficalColor: "verde e amarelo",
		Population:   10000,
	}

	country, err := NewCountry(params)
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.Equal(t, countryName, country.Name)
	assert.Equal(t, countryCapital, country.Capital)
	assert.Equal(t, countryArea, country.Area)
	assert.Equal(t, countryLanguage, country.Languages)
	assert.Equal(t, countryPopulation, country.Population)
}
