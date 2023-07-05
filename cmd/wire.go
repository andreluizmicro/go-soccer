//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	SharedDomainRepositoryInterface "github.com/andreluizmicro/go-soccer/internal/shared/domain/repository"
	SharedRepository "github.com/andreluizmicro/go-soccer/internal/shared/infrastructure/repository"
	"github.com/andreluizmicro/go-soccer/internal/shared/usercase/country"
	"github.com/google/wire"
)

var setCountryRepositoryDependency = wire.NewSet(
	SharedRepository.NewCountryRepository,
	wire.Bind(new(SharedDomainRepositoryInterface.CountryRepositoryInterface), new(*SharedRepository.CountryRepository)),
)

func NewCreateCountryUseCase(db *sql.DB) *country.CreateCountryUseCase {
	wire.Build(
		setCountryRepositoryDependency,
		country.NewCreateCountryUseCase,
	)
	return &country.CreateCountryUseCase{}
}
