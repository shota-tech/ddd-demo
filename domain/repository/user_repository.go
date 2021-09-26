package repository

import "github.com/shota-tech/layered-architecture-demo/domain/model"

type UserRepository interface {
	FindById(int) (*model.User, error)
	FindAll() ([]model.User, error)
	Save(*model.User) (int, error)
}
