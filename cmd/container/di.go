package container

import "github.com/andreluizmicro/go-soccer/internal/country/usecase"

var CountryService *usecase.CreateCountryUseCase

func Testando(teste *usecase.CreateCountryUseCase) {
	CountryService = teste
}
