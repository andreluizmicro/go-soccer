package entity

import (
	country "github.com/andreluizmicro/go-soccer/internal/shared/domain/entity"
	"github.com/andreluizmicro/go-soccer/pkg/entity"
)

type Props struct {
	Name        string
	Symbol      string
	Country     country.Country
	Coach       Coach
	Players     []Player
	Stadium     Stadium
	FoundedYear uint
	Titles      uint
}

type Team struct {
	ID          entity.ID       `json:"id"`
	Name        string          `json:"name"`
	Country     country.Country `json:"country"`
	Coach       Coach           `json:"coach"`
	Players     []Player        `json:"players"`
	Stadium     Stadium         `json:"stadium"`
	FoundedYear uint            `json:"founded_year"`
	Titles      uint            `json:" titles"`
	Symbol      string          `json:"symbol"`
}

func NewTeam(props Props) (*Team, error) {
	return &Team{
		ID:          entity.NewID(),
		Name:        props.Name,
		Country:     props.Country,
		Coach:       props.Coach,
		Players:     props.Players,
		Stadium:     props.Stadium,
		FoundedYear: props.FoundedYear,
		Titles:      props.Titles,
		Symbol:      props.Symbol,
	}, nil
}
