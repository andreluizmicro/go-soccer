package entity

import (
	country "github.com/andreluizmicro/go-soccer/internal/shared/domain/entity"
	"github.com/andreluizmicro/go-soccer/pkg/entity"
)

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

func NewTeam(name, symbol string, country country.Country, coach Coach, players []Player, stadium Stadium, foundedYear, titles uint) (*Team, error) {
	return &Team{
		ID:          entity.NewID(),
		Name:        name,
		Country:     country,
		Coach:       coach,
		Players:     players,
		Stadium:     stadium,
		FoundedYear: foundedYear,
		Titles:      titles,
		Symbol:      symbol,
	}, nil
}
