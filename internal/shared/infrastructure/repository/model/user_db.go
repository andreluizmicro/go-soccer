package model

import (
	"github.com/andreluizmicro/go-soccer/internal/shared/domain/entity"
	"gorm.io/gorm"
)

type UserModel struct {
	DB *gorm.DB
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		DB: db,
	}
}

func (model *UserModel) Create(user *entity.User) error {
	return model.DB.Create(user).Error
}

func (model *UserModel) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	if err := model.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
