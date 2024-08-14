package usecase_test

import (
	"context"
	domain "task-manager/Domain"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *UseCasetestSuite) TestTaskUsecase_GetAllTasks() {

	userid := primitive.NewObjectID()

	expectedTasks := []domain.Task{
		{ID: userid, Title: "Test Task"},
	}

	s.Run("success", func() {
		s.mockTaskRepository.On("GetAllTasks", true, userid).Return(expectedTasks, nil).Once()

		tasks, err := s.taskUsecase.GetAllTasks(context.Background(), true, userid)

		assert.Nil(s.T(), err)
		assert.Equal(s.T(), expectedTasks, tasks)
		s.mockTaskRepository.AssertExpectations(s.T())
	})

	s.Run("error", func() {
		s.mockTaskRepository.On("GetAllTasks", true, userid).Return([]domain.Task{}, domain.ErrNoTasksFound).Once()

		tasks, err := s.taskUsecase.GetAllTasks(context.Background(), true, userid)

		assert.Error(s.T(), err.Unwrap())
		assert.NotEqual(s.T(), expectedTasks, tasks)
		s.mockTaskRepository.AssertExpectations(s.T())
	})
}

func (s *UseCasetestSuite) TestTaskUsecase_AddTask() {
	task := domain.Task{
		Title:       "New Task",
		Description: "This is a new task.",
		UserID:      primitive.NewObjectID(),
		Status:      "Completed",
		DueDate:     time.Now(),
	}

	s.Run("success", func() {
		s.mockTaskRepository.On("AddTask", task).Return(nil).Once()

		err := s.taskUsecase.AddTask(context.Background(), task)

		assert.Nil(s.T(), err)

		s.mockTaskRepository.AssertExpectations(s.T())
	})

	s.Run("error", func() {
		s.mockTaskRepository.On("AddTask", task).Return(domain.ErrTaskInsertionFailed).Once()

		err := s.taskUsecase.AddTask(context.Background(), task)

		assert.Error(s.T(), err.Unwrap())

		s.mockTaskRepository.AssertExpectations(s.T())
	})
}

func (s *UseCasetestSuite) TestTaskUsecase_GetTask() {

	taskID := primitive.NewObjectID().Hex()
	userid := primitive.NewObjectID().Hex()
	expectedTask := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "New Task",
		Description: "New Description",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "Completed",
	}
	emptyTask := *new(domain.Task)

	s.Run("success", func() {
		s.mockTaskRepository.On("GetTask", taskID, true, userid).Return(expectedTask, nil).Once()

		task, err := s.taskUsecase.GetTask(context.Background(), taskID, true, userid)

		assert.Nil(s.T(), err)
		assert.Equal(s.T(), expectedTask, task)
		s.mockTaskRepository.AssertExpectations(s.T())
	})

	s.Run("error", func() {
		s.mockTaskRepository.On("GetTask", taskID, true, userid).Return(emptyTask, domain.ErrTaskNotFound).Once()

		task, err := s.taskUsecase.GetTask(context.Background(), taskID, true, userid)

		assert.Error(s.T(), err.Unwrap())

		assert.NotEqual(s.T(), expectedTask, task)
		s.mockTaskRepository.AssertExpectations(s.T())
	})
}

func (s *UseCasetestSuite) TestTaskUsecase_SetTask() {

	taskID := primitive.NewObjectID().Hex()
	updatedTask := domain.Task{
		Title:       "Updated Task",
		Description: "Updated Description",
		DueDate:     time.Now().Add(24 * time.Hour),
		Status:      "Completed",
	}

	s.Run("success", func() {
		s.mockTaskRepository.On("SetTask", taskID, updatedTask, true).Return(nil).Once()

		err := s.taskUsecase.SetTask(context.Background(), taskID, updatedTask, true)

		assert.Nil(s.T(), err)

		s.mockTaskRepository.AssertExpectations(s.T())
	})

	s.Run("error", func() {
		s.mockTaskRepository.On("SetTask", taskID, updatedTask, true).Return(domain.ErrTaskUpdateFailed).Once()

		err := s.taskUsecase.SetTask(context.Background(), taskID, updatedTask, true)

		assert.Error(s.T(), err.Unwrap())

		s.mockTaskRepository.AssertExpectations(s.T())
	})
}

func (s *UseCasetestSuite) TestTaskUsecase_DeleteTask() {

	taskID := primitive.NewObjectID().Hex()
	userid := primitive.NewObjectID().Hex()

	s.Run("success", func() {
		s.mockTaskRepository.On("DeleteTask", taskID, userid, true).Return(nil).Once()

		err := s.taskUsecase.DeleteTask(context.Background(), taskID, userid, true)

		assert.Nil(s.T(), err)

		s.mockTaskRepository.AssertExpectations(s.T())
	})

	s.Run("error", func() {
		s.mockTaskRepository.On("DeleteTask", taskID, userid, true).Return(domain.ErrTaskDeletionFailed).Once()

		err := s.taskUsecase.DeleteTask(context.Background(), taskID, userid, true)

		assert.Error(s.T(), err.Unwrap())

		s.mockTaskRepository.AssertExpectations(s.T())
	})

}
