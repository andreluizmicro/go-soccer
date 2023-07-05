package entity

import (
	"github.com/andreluizmicro/go-soccer/pkg/entity"
)

type Coach struct {
	ID          entity.ID `json:"id"`
	Name        string    `json:"name"`
	Nationality string    `json:"nationality"`
	Age         uint      `json:"age"`
}

func NewCoach(name, nationality string, age uint) (*Coach, error) {
	return &Coach{
		ID:          entity.NewID(),
		Name:        name,
		Nationality: nationality,
		Age:         age,
	}, nil
}
