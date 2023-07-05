package controllers

import (
	"github.com/andreluizmicro/go-soccer/internal/shared/domain/repository"
)

type WebCountryHandler struct {
	CountryRepository repository.CountryRepositoryInterface
}

func NewWebCountryHandler(CountryRepository repository.CountryRepositoryInterface) *WebCountryHandler {
	return &WebCountryHandler{
		CountryRepository: CountryRepository,
	}
}
