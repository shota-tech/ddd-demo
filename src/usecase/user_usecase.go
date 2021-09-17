package usecase

import (
	"github.com/shota-tech/layered-architecture-demo/src/domain/model"
	"github.com/shota-tech/layered-architecture-demo/src/domain/repository"
)

type UserUsecase interface {
	GetUserByID(int) (*model.User, error)
	GetUserList() ([]model.User, error)
	AddUser(*model.User) (int, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (uu userUsecase) GetUserByID(id int) (*model.User, error) {
	user, err := uu.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (uu userUsecase) GetUserList() ([]model.User, error) {
	users, err := uu.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uu userUsecase) AddUser(user *model.User) (int, error) {
	id, err := uu.userRepository.Save(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
