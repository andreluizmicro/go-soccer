package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/andreluizmicro/go-soccer/internal/shared/domain/repository"
	usecase "github.com/andreluizmicro/go-soccer/internal/shared/usercase/country"
)

type WebCountryHandler struct {
	CountryRepository repository.CountryRepositoryInterface
}

func NewWebCountryHandler(CountryRepository repository.CountryRepositoryInterface) *WebCountryHandler {
	return &WebCountryHandler{
		CountryRepository: CountryRepository,
	}
}

func (handler *WebCountryHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.CountryInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	usecase := usecase.NewCreateCountryUseCase(handler.CountryRepository)
	output, err := usecase.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
