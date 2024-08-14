package usecase_test

import (
	domain "task-manager/Domain"
	usecase "task-manager/Usecase"
	"task-manager/mocks/repomocks"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

// UseCasetestSuite struct to hold any shared resources or setup for the tests
type UseCasetestSuite struct {
	suite.Suite
	mockTaskRepository *repomocks.TaskRepository
	mockUserRepository *repomocks.UserRepository
	taskUsecase        domain.TaskUsecase
	userUsecase        domain.UserUsecase
}

// SetupTest runs before each test case
func (s *UseCasetestSuite) SetupTest() {
	s.mockTaskRepository = new(repomocks.TaskRepository)
	s.mockUserRepository = new(repomocks.UserRepository)
	s.taskUsecase = usecase.NewTaskUsecase(s.mockTaskRepository, 30*time.Second)
	s.userUsecase = usecase.NewUserUsecase(s.mockUserRepository, 30*time.Second)
}

// TearDownTest runs after each test case
func (s *UseCasetestSuite) TearDownTest() {
	// Clean up resources if needed
}

// TestRunSuite runs the test suite
func TestRunSuite(t *testing.T) {
	suite.Run(t, new(UseCasetestSuite))
}
