package usecase

import (
	"github.com/shota-tech/ddd-demo/domain/model"
	"github.com/shota-tech/ddd-demo/domain/repository"
)

type UserUsecase interface {
	GetUserByID(int) (*model.User, error)
	GetUserList() ([]model.User, error)
	AddUser(*model.User) (int, error)
	EditUser(int, *model.User) error
	DeleteUser(int) error
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
	}
}

func (u *userUsecase) GetUserByID(id int) (*model.User, error) {
	user, err := u.userRepository.FindById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userUsecase) GetUserList() ([]model.User, error) {
	users, err := u.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userUsecase) AddUser(user *model.User) (int, error) {
	id, err := u.userRepository.Create(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *userUsecase) EditUser(id int, user *model.User) error {
	err := u.userRepository.Update(id, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) DeleteUser(id int) error {
	err := u.userRepository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
