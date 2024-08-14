package repositories_test

import (
	domain "task-manager/Domain"
	repositories "task-manager/Repositories"
	"task-manager/mocks/databasemock"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TaskRepositorySuite defines the test suite for the TaskRepository
type TaskRepositorySuite struct {
	suite.Suite
	databaseHelper   *databasemock.Database
	collectionHelper *databasemock.Collection
	taskRepo         domain.TaskRepository
	mockSingleResult *databasemock.SingleResult
}

// SetupTest initializes the necessary components for the test suite
func (suite *TaskRepositorySuite) SetupTest() {
	suite.databaseHelper = new(databasemock.Database)
	suite.collectionHelper = new(databasemock.Collection)
	suite.mockSingleResult = new(databasemock.SingleResult)
	db := suite.databaseHelper

	suite.databaseHelper.On("Collection", mock.Anything).Return(suite.collectionHelper)
	// Initialize the domain with the mocked database and collection
	suite.taskRepo = repositories.NewTaskRepository(db)
}

// TestCreateUser_Success tests the successful creation of a user
func (suite *TaskRepositorySuite) TestGetAllTasks() {
	userid := primitive.NewObjectID()

	cursor := new(databasemock.Cursor)

	suite.collectionHelper.On("Find", mock.Anything, mock.Anything).Return(cursor, nil)

	cursor.On("All", mock.Anything, mock.Anything).Return(nil)

	_, err := suite.taskRepo.GetAllTasks(true, userid)

	// Assert that no error was returned
	suite.Nil(err)

	// Assert that the expectations were met
	suite.collectionHelper.AssertExpectations(suite.T())
}

func (suite *TaskRepositorySuite) TestGetTask() {
	userid := primitive.NewObjectID().Hex()
	taskid := primitive.NewObjectID().Hex()

	suite.mockSingleResult.On("Decode", mock.Anything).Return(nil).Once()

	suite.collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult).Once()

	_, err := suite.taskRepo.GetTask(taskid, true, userid)

	// Assert that no error was returned
	suite.Nil(err)

	// Assert that the expectations were met
	suite.collectionHelper.AssertExpectations(suite.T())
}
func (suite *TaskRepositorySuite) TestAddTask() {
	userid := primitive.NewObjectID().Hex()
	taskid := primitive.NewObjectID().Hex()

	suite.mockSingleResult.On("Decode", mock.Anything).Return(nil).Once()

	suite.collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult).Once()

	_, err := suite.taskRepo.GetTask(taskid, true, userid)

	// Assert that no error was returned
	suite.Nil(err)

	// Assert that the expectations were met
	suite.collectionHelper.AssertExpectations(suite.T())
}
func (suite *TaskRepositorySuite) TestSetTask() {
	taskid := primitive.NewObjectID()
	updatedTask := domain.Task{ID: taskid, Title: "Test Task"}
	updateres := new(mongo.UpdateResult)
	updateres.MatchedCount = 1

	suite.collectionHelper.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(updateres, nil).Once()

	err := suite.taskRepo.SetTask(taskid.Hex(), updatedTask, true)

	// Assert that no error was returned
	suite.Nil(err)

	// Assert that the expectations were met
	suite.collectionHelper.AssertExpectations(suite.T())
}

func (suite *TaskRepositorySuite) TestDeleteTask() {
	taskid := primitive.NewObjectID()
	userid := primitive.NewObjectID()

	suite.collectionHelper.On("DeleteOne", mock.Anything, mock.Anything).Return(int64(1), nil).Once()

	err := suite.taskRepo.DeleteTask(taskid.Hex(), userid.Hex(), true)

	// Assert that no error was returned
	suite.Nil(err)

	// Assert that the expectations were met
	suite.collectionHelper.AssertExpectations(suite.T())
}

// TestTaskRepositorySuite runs the TaskRepositorySuite
func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
}
