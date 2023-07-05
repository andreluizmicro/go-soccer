package test

import (
	"fmt"
	"testing"

	"github.com/andreluizmicro/go-soccer/internal/country/infrastructure/repository"
	country "github.com/andreluizmicro/go-soccer/internal/country/usecase"
)

func TestGetRequest(t *testing.T) {
	repository := repository.NewCountryRepository()
	service := country.NewCreateCountryUseCase(repository)

	input := country.CountryInputDTO{
		Name:         "Brasil",
		Capital:      "Brasilia",
		Area:         "10000x1000",
		Language:     "portuguese",
		Timezone:     "UTC",
		Continent:    "Americano",
		OfficalColor: "verde amarelo",
		Population:   10,
	}

	outrput, err := service.Execute(input)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(outrput.Name)
}
