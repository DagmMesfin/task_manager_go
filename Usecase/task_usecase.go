package usecase

import (
	"context"
	domain "task-manager/Domain"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type taskUsecase struct {
	taskRepository domain.TaskRepository
	contextTimeout time.Duration
}

func NewTaskUsecase(taskRepository domain.TaskRepository, timeout time.Duration) domain.TaskUsecase {
	return &taskUsecase{
		taskRepository: taskRepository,
		contextTimeout: timeout,
	}
}

func (tuse *taskUsecase) GetAllTasks(c context.Context, isadmin bool, userid primitive.ObjectID) ([]domain.Task, *domain.AppError) {
	_, cancel := context.WithTimeout(c, tuse.contextTimeout)
	defer cancel()
	return tuse.taskRepository.GetAllTasks(isadmin, userid)
}

func (tuse *taskUsecase) GetTask(c context.Context, id string, isadmin bool, userid string) (domain.Task, *domain.AppError) {
	_, cancel := context.WithTimeout(c, tuse.contextTimeout)
	defer cancel()
	return tuse.taskRepository.GetTask(id, isadmin, userid)
}

func (tuse *taskUsecase) AddTask(c context.Context, task domain.Task) *domain.AppError {
	_, cancel := context.WithTimeout(c, tuse.contextTimeout)
	defer cancel()
	return tuse.taskRepository.AddTask(task)

}

func (tuse *taskUsecase) SetTask(c context.Context, id string, updatedTask domain.Task, isadmin bool) *domain.AppError {
	_, cancel := context.WithTimeout(c, tuse.contextTimeout)
	defer cancel()
	return tuse.taskRepository.SetTask(id, updatedTask, isadmin)
}

func (tuse *taskUsecase) DeleteTask(c context.Context, id string, userid string, isadmin bool) *domain.AppError {
	_, cancel := context.WithTimeout(c, tuse.contextTimeout)
	defer cancel()
	return tuse.taskRepository.DeleteTask(id, userid, isadmin)
}
