package usecase

import (
	"context"
	domain "task-manager/Domain"
	"time"
)

type userUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, timeout time.Duration) domain.UserUsecase {
	return &userUsecase{
		userRepository: userRepository,
		contextTimeout: timeout,
	}
}

func (useru *userUsecase) RegisterUserDb(c context.Context, user domain.User) *domain.AppError {
	_, cancel := context.WithTimeout(c, useru.contextTimeout)
	defer cancel()
	return useru.userRepository.RegisterUserDb(user)
}

func (useru *userUsecase) LoginUserDb(c context.Context, user domain.User) (string, interface{}, *domain.AppError) {
	_, cancel := context.WithTimeout(c, useru.contextTimeout)
	defer cancel()
	return useru.userRepository.LoginUserDb(user)

}

func (useru *userUsecase) DeleteUser(c context.Context, id string) *domain.AppError {
	_, cancel := context.WithTimeout(c, useru.contextTimeout)
	defer cancel()
	return useru.userRepository.DeleteUser(id)

}
