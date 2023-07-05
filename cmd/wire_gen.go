// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	repository2 "github.com/andreluizmicro/go-soccer/internal/country/domain/repository"
	"github.com/andreluizmicro/go-soccer/internal/country/infrastructure/repository"
	"github.com/andreluizmicro/go-soccer/internal/country/usecase"
	"github.com/google/wire"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func NewCreateCountryUseCase() *usecase.CreateCountryUseCase {
	countryRepository := repository.NewCountryRepository()
	createCountryUseCase := usecase.NewCreateCountryUseCase(countryRepository)
	return createCountryUseCase
}

// wire.go:

var setCountryRepositoryDependency = wire.NewSet(repository.NewCountryRepository, wire.Bind(new(repository2.CountryRepositoryInterface), new(*repository.CountryRepository)))
