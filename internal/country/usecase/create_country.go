package country

import (
	"github.com/andreluizmicro/go-soccer/internal/country/domain/entity"
	"github.com/andreluizmicro/go-soccer/internal/country/domain/repository"
)

type CountryInputDTO struct {
	Name         string `json:"name"`
	Capital      string `json:"capital"`
	Area         string `json:"area"`
	Languages    string `json:"language"`
	Timezone     string `json:"timezone"`
	Continent    string `json:"continent"`
	OfficalColor string `json:"offical_color"`
	Population   int    `json:"population"`
}

type CountryOutputDTO struct {
	Name string `json:"name"`
}

type CreateCountryUseCase struct {
	CountryRepository repository.CountryRepositoryInterface
}

func NewCreateCountryUseCase(CountryRepository repository.CountryRepositoryInterface) *CreateCountryUseCase {
	return &CreateCountryUseCase{
		CountryRepository: CountryRepository,
	}
}

func (service *CreateCountryUseCase) Execute(input CountryInputDTO) (CountryOutputDTO, error) {
	params := entity.Params{
		Name:         input.Name,
		Capital:      input.Capital,
		Area:         input.Area,
		Languages:    input.Languages,
		Timezone:     input.Timezone,
		Continent:    input.Continent,
		OfficalColor: input.OfficalColor,
		Population:   input.Population,
	}

	country, err := entity.NewCountry(params)

	if err := service.CountryRepository.Create(country); err != nil {
		return CountryOutputDTO{}, err
	}

	return CountryOutputDTO{
		Name: country.Name,
	}, err
}
