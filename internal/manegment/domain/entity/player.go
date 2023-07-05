package entity

import "github.com/andreluizmicro/go-soccer/pkg/entity"

type Player struct {
	ID            entity.ID `json:"id"`
	Name          string    `json:"name"`
	Age           uint      `json:"age"`
	Nationality   string    `json:"nationality"`
	Position      uint      `json:"position"`
	OverallRating int64     `json:"overall_rating"`
}

func NewPlayer(name, nationality string, overallRating int64, age, position uint) (*Player, error) {
	return &Player{
		ID:            entity.NewID(),
		Name:          name,
		Age:           age,
		Nationality:   nationality,
		Position:      position,
		OverallRating: overallRating,
	}, nil
}
