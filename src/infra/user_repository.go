package infra

import (
	"github.com/shota-tech/layered-architecture-demo/src/domain/model"
	"github.com/shota-tech/layered-architecture-demo/src/domain/repository"
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (r *userRepository) FindById(id int) (*model.User, error) {
	// TODO
	dummyUser := &model.User{
		ID:    1,
		Name:  "user1",
		Email: "user1@sample.com",
	}
	return dummyUser, nil
}

func (r *userRepository) FindAll() ([]model.User, error) {
	// TODO
	dummyUsers := []model.User{
		{
			ID:    1,
			Name:  "user1",
			Email: "user1@sample.com",
		},
		{
			ID:    2,
			Name:  "user2",
			Email: "user2@sample.com",
		},
		{
			ID:    3,
			Name:  "user3",
			Email: "user3@sample.com",
		},
	}
	return dummyUsers, nil
}

func (r *userRepository) Save(*model.User) (int, error) {
	// TODO
	return 1, nil
}
