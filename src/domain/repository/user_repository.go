package repository

import "github.com/shota-tech/layered-architecture-demo/src/domain/model"

type UserRepository interface {
	FindById(int) (*model.User, error)
	FindAll() ([]model.User, error)
	Save(*model.User) (int, error)
}
