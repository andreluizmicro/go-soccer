package repository

import (
	"gorm.io/gorm"
)

type StadiumRepository struct {
	DB *gorm.DB
}

func NewStadiumRepository(db *gorm.DB) *StadiumRepository {
	return &StadiumRepository{
		DB: db,
	}
}
