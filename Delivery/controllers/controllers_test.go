package controllers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"task-manager/Delivery/controllers"
	domain "task-manager/Domain"
	"task-manager/mocks/usecasemock"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TaskControllerSuite struct {
	suite.Suite
	taskUsecase   *usecasemock.TaskUsecase // Pointer to the mock
	taskCtrl      *controllers.TaskController
	testingServer *httptest.Server
	userUsecase   *usecasemock.UserUsecase // Pointer to the mock
}

func (suite *TaskControllerSuite) SetupTest() {
	suite.taskUsecase = new(usecasemock.TaskUsecase) // Initialize the mock
	suite.userUsecase = new(usecasemock.UserUsecase) // Initialize the mock
	suite.taskCtrl = controllers.NewTaskController(suite.taskUsecase, suite.userUsecase)

	router := gin.Default()
	router.POST("/tasks", suite.taskCtrl.PostTask)
	router.GET("/tasks/:id", suite.taskCtrl.GetTasksById)
	router.GET("/tasks", suite.taskCtrl.GetTasks)
	router.PUT("/tasks/:id", suite.taskCtrl.PutTask)
	router.DELETE("/tasks/:id", suite.taskCtrl.DeleteTask)

	router.POST("/register", suite.taskCtrl.RegisterUser)
	router.POST("/login", suite.taskCtrl.LoginUser)

	router.DELETE("/users/:id", suite.taskCtrl.DeleteUser)

	suite.testingServer = httptest.NewServer(router)
}

func (suite *TaskControllerSuite) TearDownTest() {
	suite.testingServer.Close()
}

func TestTaskControllerSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerSuite))
}

func (suite *TaskControllerSuite) TestPostTask() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		UserID:      primitive.NewObjectID(),
		Title:       "task1",
		Description: "description",
		DueDate:     time.Now(),
		Status:      "Completed",
	}

	// Mock the AddTask use case
	suite.taskUsecase.On("AddTask", mock.Anything, mock.AnythingOfType("domain.Task")).Return(nil)

	// Marshal the task to JSON
	requestBody, err := json.Marshal(&task)
	suite.NoError(err, "cannot marshal struct to json")

	// Send POST request
	response, err := http.Post(fmt.Sprintf("%s/tasks", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	// Check the response status code
	suite.Equal(http.StatusCreated, response.StatusCode, "Expected status OK")

	// Define the expected response structure
	var responseBody domain.Response
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "cannot decode response")
	log.Println(responseBody)
}

func (suite *TaskControllerSuite) TestGetTasks() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "description",
		DueDate:     time.Now(),
		Status:      "Completed",
		UserID:      primitive.NewObjectID(),
	}

	// Mock the GetAllTasks use case
	suite.taskUsecase.On("GetAllTasks", mock.Anything, mock.Anything, mock.Anything).Return([]domain.Task{task}, nil)

	// Send GET request
	response, err := http.Get(fmt.Sprintf("%s/tasks", suite.testingServer.URL))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	// Check the response status code
	suite.Equal(http.StatusOK, response.StatusCode, "Expected status OK")

	// Define the expected response structure
	var responseBody domain.Response
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "cannot decode response")

	// Assert that the response matches expected values
	log.Println(responseBody)

}

func (suite *TaskControllerSuite) TestGetTask() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "description",
		DueDate:     time.Now(),
		Status:      "Completed",
		UserID:      primitive.NewObjectID(),
	}

	suite.taskUsecase.On("GetTask", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(task, nil)

	response, err := http.Get(fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, task.ID.Hex()))
	suite.NoError(err, "no error when calling the endpoint")

	// Check the response status code
	suite.Equal(http.StatusOK, response.StatusCode, "Expected status OK")

	responseBody := domain.Response{}

	err = json.NewDecoder(response.Body).Decode(&responseBody)
	suite.NoError(err, "cannot decode response")
}

func (suite *TaskControllerSuite) TestSetTasks() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "description",
		DueDate:     time.Now(),
		Status:      "Completed",
		UserID:      primitive.NewObjectID(),
	}

	suite.taskUsecase.On("SetTask", mock.Anything, task.ID.Hex(), mock.Anything, mock.Anything).Return(nil)

	requestBody, err := json.Marshal(&task)
	suite.NoError(err, "can not marshal struct to json")

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, task.ID.Hex()), bytes.NewBuffer(requestBody))
	suite.NoError(err, "can not create PUT request")

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(req)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	suite.Equal(http.StatusOK, response.StatusCode, "it must be status OK")

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	// You can add further assertions here if needed
}

func (suite *TaskControllerSuite) TestDeleteTask() {
	task := domain.Task{
		ID:          primitive.NewObjectID(),
		Title:       "task1",
		Description: "description",
		DueDate:     time.Now(),
		Status:      "Completed",
		UserID:      primitive.NewObjectID(),
	}

	suite.taskUsecase.On("DeleteTask", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/tasks/%s", suite.testingServer.URL, task.ID.Hex()), nil)
	suite.NoError(err, "can not create DELETE request")

	client := &http.Client{}
	response, err := client.Do(req)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

}

func (suite *TaskControllerSuite) TestRegisterUserDb() {
	user := domain.User{
		Email:    "tester1@gmail.com",
		Password: "password",
		IsAdmin:  true,
	}
	suite.userUsecase.On("RegisterUserDb", mock.Anything, user).Return(nil)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/register", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusCreated, response.StatusCode)
	suite.Equal("User registered successfully", responseBody.Message)
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestLogin() {
	user := domain.User{
		Email:    "tester1@gmail.com",
		Password: "password",
	}
	token_string := "the-token"
	suite.userUsecase.On("LoginUserDb", mock.Anything, user).Return(token_string, mock.Anything, nil)

	requestBody, err := json.Marshal(&user)
	suite.NoError(err, "can not marshal struct to json")

	response, err := http.Post(fmt.Sprintf("%s/login", suite.testingServer.URL), "application/json", bytes.NewBuffer(requestBody))
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("User logged in successfully", responseBody.Message)
	suite.userUsecase.AssertExpectations(suite.T())
}

func (suite *TaskControllerSuite) TestDeleteUser() {
	userid := primitive.NewObjectID().Hex()

	suite.userUsecase.On("DeleteUser", mock.Anything, mock.Anything).Return(nil)

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/users/%s", suite.testingServer.URL, userid), nil)
	suite.NoError(err, "can not create Delete request")

	client := &http.Client{}
	response, err := client.Do(req)
	suite.NoError(err, "no error when calling the endpoint")
	defer response.Body.Close()

	responseBody := domain.Response{}
	json.NewDecoder(response.Body).Decode(&responseBody)

	suite.Equal(http.StatusOK, response.StatusCode)
	suite.Equal("user deleted", responseBody.Message)
	suite.userUsecase.AssertExpectations(suite.T())
}
