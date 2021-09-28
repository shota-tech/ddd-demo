package repository

import "github.com/shota-tech/ddd-demo/domain/model"

type UserRepository interface {
	FindById(int) (*model.User, error)
	FindAll() ([]model.User, error)
	Create(*model.User) (int, error)
	Update(int, *model.User) error
	Delete(int) error
}
