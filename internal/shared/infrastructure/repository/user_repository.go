package repository

import (
	"errors"

	"github.com/andreluizmicro/go-soccer/internal/shared/domain/entity"
	"github.com/andreluizmicro/go-soccer/internal/shared/infrastructure/repository/model"
)

type UserRepository struct {
	model model.UserModel
}

func NewUserRepository(model *model.UserModel) *UserRepository {
	return &UserRepository{
		model: *model,
	}
}

func (repository *UserRepository) Create(user *entity.User) error {
	return repository.model.Create(user)
}

func (repository *UserRepository) FindByEmail(email string) (*entity.User, error) {
	user, err := repository.model.FindByEmail(email)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}
