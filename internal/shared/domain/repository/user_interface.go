package repository

import "github.com/andreluizmicro/go-soccer/internal/shared/domain/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
