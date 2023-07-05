package entity

import "github.com/andreluizmicro/go-soccer/pkg/entity"

type Params struct {
	ID           string
	Name         string
	Capital      string
	Area         string
	Language     string
	Timezone     string
	Continent    string
	OfficalColor string
	Population   int
}

type Country struct {
	ID           entity.ID `json:"id"`
	Name         string    `json:"name"`
	Capital      string    `json:"capital"`
	Area         string    `json:"area"`
	Language     string    `json:"language"`
	Timezone     string    `json:"timezone"`
	Continent    string    `json:"continent"`
	OfficalColor string    `json:"officalcolor"`
	Population   int       `json:"population"`
}

func NewCountry(params Params) (*Country, error) {
	return &Country{
		ID:           entity.NewID(),
		Name:         params.Name,
		Capital:      params.Capital,
		Population:   params.Population,
		Area:         params.Area,
		Language:     params.Language,
		Timezone:     params.Timezone,
		Continent:    params.Continent,
		OfficalColor: params.OfficalColor,
	}, nil
}
