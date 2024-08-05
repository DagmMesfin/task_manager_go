package data

import (
	"context"
	"errors"
	"log"
	"net/http"
	"task-manager/models"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserManager struct {
	client *mongo.Client
}

func NewUserManager(mongoClient *mongo.Client) *UserManager {
	return &UserManager{
		client: mongoClient,
	}
}

func (taskmgr *UserManager) RegisterUserDb(user models.User) (int, error) {

	collection := taskmgr.client.Database("task-manager").Collection("users")

	ere := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Err()

	if ere == nil {
		return http.StatusBadRequest, errors.New("user already exists with same email")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user.ID = primitive.NewObjectID()

	user.Password = string(hashedPassword)

	_, erro := collection.InsertOne(context.TODO(), user)

	if erro != nil {
		return http.StatusBadRequest, erro
	}

	return http.StatusOK, nil

}

func (taskmgr *UserManager) LoginUserDb(user models.User) (int, error, string) {
	collection := taskmgr.client.Database("task-manager").Collection("users")

	var jwtSecret = []byte("your_jwt_secret")

	var existingUser models.User

	// User login logic
	collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)

	log.Println(existingUser, user)
	if bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password)) != nil {

		return http.StatusUnauthorized, errors.New("Invalid email or password"), ""
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"_id":     existingUser.ID,
		"email":   existingUser.Email,
		"isadmin": existingUser.IsAdmin,
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return http.StatusInternalServerError, errors.New("Internal server error"), ""
	}

	return http.StatusOK, nil, jwtToken

}
