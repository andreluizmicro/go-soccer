package repository

import (
	"github.com/andreluizmicro/go-soccer/internal/shared/domain/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (repository *UserRepository) Create(user *entity.User) error {
	return repository.DB.Create(user).Error
}

func (repository *UserRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := repository.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
