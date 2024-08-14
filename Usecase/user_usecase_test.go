package usecase_test

import (
	"context"
	domain "task-manager/Domain"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *UseCasetestSuite) TestUserUsecase_RegisterUserDb() {
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@example.com",
		Password: "password",
	}

	s.Run("success", func() {
		s.mockUserRepository.On("RegisterUserDb", user).Return(nil).Once()

		err := s.userUsecase.RegisterUserDb(context.Background(), user)

		s.Nil(err)
		s.mockUserRepository.AssertExpectations(s.T())
	})

	s.Run("error", func() {
		s.mockUserRepository.On("RegisterUserDb", user).Return(domain.ErrUserRegistrationFailed).Once()

		err := s.userUsecase.RegisterUserDb(context.Background(), user)

		s.Error(err.Unwrap())
		s.mockUserRepository.AssertExpectations(s.T())
	})

}

func (s *UseCasetestSuite) TestUserUsecase_LoginUserDb() {
	user := domain.User{
		Email:    "test@example.com",
		Password: "password",
	}

	expectedToken := "test-token"

	s.Run("success", func() {
		s.mockUserRepository.On("LoginUserDb", user).Return(expectedToken, mock.Anything, nil).Once()

		token, result, err := s.userUsecase.LoginUserDb(context.Background(), user)

		s.NotNil(result)
		s.Nil(err)
		s.Equal(expectedToken, token)
		s.mockUserRepository.AssertExpectations(s.T())
	})

	s.Run("error", func() {
		s.mockUserRepository.On("LoginUserDb", user).Return("", mock.Anything, domain.ErrInvalidCredentials).Once()

		token, _, err := s.userUsecase.LoginUserDb(context.Background(), user)

		s.Error(err.Unwrap())
		s.Empty(token)
		s.mockUserRepository.AssertExpectations(s.T())
	})

}

func (s *UseCasetestSuite) TestUserUsecase_DeleteUser() {
	userID := primitive.NewObjectID().Hex()

	s.Run("success", func() {
		s.mockUserRepository.On("DeleteUser", userID).Return(nil).Once()

		err := s.userUsecase.DeleteUser(context.Background(), userID)

		s.Nil(err)
		s.mockUserRepository.AssertExpectations(s.T())
	})

	s.Run("error", func() {
		s.mockUserRepository.On("DeleteUser", userID).Return(domain.ErrInvalidCredentials).Once()

		err := s.userUsecase.DeleteUser(context.Background(), userID)

		s.Error(err.Unwrap())
		s.mockUserRepository.AssertExpectations(s.T())
	})

}
