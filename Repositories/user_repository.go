package repositories

import (
	"context"
	"errors"
	"log"
	"net/http"
	domain "task-manager/Domain"
	infrastructure "task-manager/Infrastructure"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func NewUserRepository(mongoClient *mongo.Client) domain.UserRepository {
	return &UserRepository{
		client:     mongoClient,
		database:   mongoClient.Database("task-manager"),
		collection: mongoClient.Database("task-manager").Collection("tasks"),
	}
}

func (userepo *UserRepository) RegisterUserDb(user domain.User) (int, error) {

	collection := userepo.collection

	ere := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Err()

	if ere == nil {
		return http.StatusBadRequest, errors.New("user already exists with same email")
	}

	hashedPassword, err := infrastructure.PasswordHasher(user.Password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	user.Password = hashedPassword

	user.ID = primitive.NewObjectID()

	_, erro := collection.InsertOne(context.TODO(), user)

	if erro != nil {
		return http.StatusBadRequest, erro
	}

	return http.StatusOK, nil

}

func (userepo *UserRepository) LoginUserDb(user domain.User) (int, string, error) {

	collection := userepo.collection

	var existingUser domain.User

	collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)

	log.Println(existingUser, user)

	if infrastructure.PasswordComparator(existingUser.Password, user.Password) {
		return http.StatusUnauthorized, "", errors.New("invalid email or password")
	}

	jwtToken, err := infrastructure.TokenGenerator(existingUser.ID, existingUser.Email, existingUser.IsAdmin)

	if err != nil {
		return http.StatusInternalServerError, "", errors.New("internal server error")
	}

	return http.StatusOK, jwtToken, nil

}

func (userepo *UserRepository) DeleteUser(id string) (int, error) {
	collection := userepo.collection

	ido, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ido}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil || result.DeletedCount == 0 {
		return 404, errors.New("user not found")
	}

	return 200, nil

}
