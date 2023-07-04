package entity

import "github.com/andreluizmicro/go-soccer/pkg/entity"

type Stadium struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Capacity int       `json:"capacity"`
}

func NewStadium(name, location string, capacity int) (*Stadium, error) {
	return &Stadium{
		ID:       entity.NewID(),
		Name:     name,
		Location: location,
		Capacity: capacity,
	}, nil
}
