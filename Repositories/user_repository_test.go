package repositories_test

import (
	domain "task-manager/Domain"
	repositories "task-manager/Repositories"
	"task-manager/mocks"
	"task-manager/mocks/databasemock"
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepositorySuite defines the test suite for the UserRepository
type UserRepositorySuite struct {
	suite.Suite
	databaseHelper   *databasemock.Database
	collectionHelper *databasemock.Collection
	userRepo         domain.UserRepository
	mockSingleResult *databasemock.SingleResult
	pass_service     *mocks.PasswordService
}

// SetupTest initializes the necessary components for the test suite
func (suite *UserRepositorySuite) SetupTest() {
	suite.databaseHelper = new(databasemock.Database)
	suite.collectionHelper = new(databasemock.Collection)
	suite.mockSingleResult = new(databasemock.SingleResult)

	db := suite.databaseHelper
	suite.pass_service = new(mocks.PasswordService)

	suite.databaseHelper.On("Collection", mock.Anything).Return(suite.collectionHelper)
	// Initialize the domain with the mocked database and collection
	suite.userRepo = repositories.NewUserRepository(db, suite.pass_service)
}

// TestCreateUser_Success tests the successful creation of a user
func (suite *UserRepositorySuite) TestCreateUser() {
	mockUser := &domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@gmail.com",
		Password: "password",
		IsAdmin:  true,
	}

	mockUserID := mockUser.ID.Hex()

	suite.mockSingleResult.On("Decode", mock.Anything).Return(nil).Once()

	// Set up the expectation for the InsertOne method
	suite.collectionHelper.On("InsertOne", mock.Anything, mock.Anything).Return(mockUserID, nil).Once()

	// Set up the expectation for the FindOne method
	suite.collectionHelper.On("FindOne", mock.Anything, mock.Anything).Return(suite.mockSingleResult, nil).Once()

	suite.pass_service.On("PasswordHasher", mockUser.Password).Return("hashed-password", nil)
	// Call the method under test
	err := suite.userRepo.RegisterUserDb(*mockUser)

	// Assert that no error was returned
	suite.Nil(err)

	// Assert that the expectations were met
	suite.collectionHelper.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestLoginUser() {
	// Create a mock user with a known ID, email, and password
	mockUser := &domain.User{
		ID:       primitive.NewObjectID(),
		Email:    "test@gmail.com",
		Password: "password",
	}

	// Mock the FindOne method to return a result
	suite.collectionHelper.On("FindOne", mock.Anything, mock.MatchedBy(func(filter interface{}) bool {
		return true // You can refine this to match a specific filter if needed
	})).Return(suite.mockSingleResult, nil).Once()

	// Mock the Decode method to decode into the mock user
	suite.mockSingleResult.On("Decode", mock.AnythingOfType("*domain.User")).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.User)
		*arg = *mockUser
	}).Return(nil).Once()

	// Mock the PasswordComparator to return true, indicating the password matches
	suite.pass_service.On("PasswordComparator", mock.Anything, mock.Anything).Return(false).Once()

	// Mock the TokenGenerator to return a specific token
	suite.pass_service.On("TokenGenerator", mock.Anything, mock.Anything, mock.Anything).Return("my-token", nil).Once()

	// Call the method under test
	token, _, err := suite.userRepo.LoginUserDb(*mockUser)

	// Assert that no error was returned
	suite.Nil(err)
	suite.Equal("my-token", token)

	// Assert that all expectations were met
	suite.collectionHelper.AssertExpectations(suite.T())
}

func (suite *UserRepositorySuite) TestDeleteUser() {

	userID := primitive.NewObjectID().Hex()

	// Set up the expectation for the FindOne method
	suite.collectionHelper.On("DeleteOne", mock.Anything, mock.Anything).Return(int64(1), nil).Once()

	// Call the method under test
	err := suite.userRepo.DeleteUser(userID)

	// Assert that no error was returned
	suite.Nil(err)

	// Assert that the expectations were met
	suite.collectionHelper.AssertExpectations(suite.T())
}

// TestUserRepositorySuite runs the UserRepositorySuite
func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
