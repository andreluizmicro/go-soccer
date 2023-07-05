package usecase

import (
	"github.com/andreluizmicro/go-soccer/internal/country/domain/entity"
	"github.com/andreluizmicro/go-soccer/internal/country/domain/repository"
)

type CountryInputDTO struct {
	Name         string `json:"name"`
	Capital      string `json:"capital"`
	Area         string `json:"area"`
	Language     string `json:"language"`
	Timezone     string `json:"timezone"`
	Continent    string `json:"continent"`
	OfficalColor string `json:"offical_color"`
	Population   int    `json:"population"`
}

type CountryOutputDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Capital      string `json:"capital"`
	Area         string `json:"area"`
	Language     string `json:"language"`
	Timezone     string `json:"timezone"`
	Continent    string `json:"continent"`
	OfficalColor string `json:"offical_color"`
	Population   int    `json:"population"`
}

type CreateCountryUseCase struct {
	CountryRepository repository.CountryRepositoryInterface
}

func NewCreateCountryUseCase(CountryRepository repository.CountryRepositoryInterface) *CreateCountryUseCase {
	return &CreateCountryUseCase{
		CountryRepository: CountryRepository,
	}
}

func (usecase *CreateCountryUseCase) Execute(input CountryInputDTO) (CountryOutputDTO, error) {
	params := entity.Params{
		Name:         input.Name,
		Capital:      input.Capital,
		Area:         input.Area,
		Language:     input.Language,
		Timezone:     input.Timezone,
		Continent:    input.Continent,
		OfficalColor: input.OfficalColor,
		Population:   input.Population,
	}

	country, err := entity.NewCountry(params)
	if err != nil {
		return CountryOutputDTO{}, err
	}

	if err := usecase.CountryRepository.Create(country); err != nil {
		return CountryOutputDTO{}, err
	}

	return CountryOutputDTO{
		Name:         country.Name,
		Capital:      input.Capital,
		Area:         input.Area,
		Language:     input.Language,
		Timezone:     input.Timezone,
		Continent:    input.Continent,
		OfficalColor: input.OfficalColor,
		Population:   input.Population,
	}, err
}
