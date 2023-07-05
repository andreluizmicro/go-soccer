//go:build wireinject
// +build wireinject

package main

import (
	CountryRepositoryInterface "github.com/andreluizmicro/go-soccer/internal/country/domain/repository"
	CountryRepository "github.com/andreluizmicro/go-soccer/internal/country/infrastructure/repository"
	CountryUseCase "github.com/andreluizmicro/go-soccer/internal/country/usecase"
	"github.com/google/wire"
)

var setCountryRepositoryDependency = wire.NewSet(
	CountryRepository.NewCountryRepository,
	wire.Bind(new(CountryRepositoryInterface.CountryRepositoryInterface), new(*CountryRepository.CountryRepository)),
)

func NewCreateCountryUseCase() *CountryUseCase.CreateCountryUseCase {
	wire.Build(
		setCountryRepositoryDependency,
		CountryUseCase.NewCreateCountryUseCase,
	)
	return &CountryUseCase.CreateCountryUseCase{}
}
