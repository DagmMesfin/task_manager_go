package repositories

import (
	"context"
	"log"
	domain "task-manager/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRepository struct {
	database     domain.Database
	collection   domain.Collection
	pass_service domain.PasswordService
}

func NewUserRepository(mongoDatabase domain.Database, password_service domain.PasswordService) domain.UserRepository {
	return &UserRepository{
		database:     mongoDatabase,
		collection:   mongoDatabase.Collection("users"),
		pass_service: password_service,
	}
}

func (userepo *UserRepository) RegisterUserDb(user domain.User) *domain.AppError {

	collection := userepo.collection

	var usero domain.User

	collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&usero)

	if usero.Email == user.Email {
		return domain.ErrUserExists
	}

	hashedPassword, err := userepo.pass_service.PasswordHasher(user.Password)
	if err != nil {
		return domain.ErrUserRegistrationFailed
	}

	user.Password = hashedPassword

	user.ID = primitive.NewObjectID()

	_, erro := collection.InsertOne(context.TODO(), user)

	if erro != nil {
		return domain.ErrUserRegistrationFailed
	}

	return nil

}

func (userepo *UserRepository) LoginUserDb(user domain.User) (string, interface{}, *domain.AppError) {

	collection := userepo.collection

	var existingUser domain.User

	result := collection.FindOne(context.TODO(), bson.M{"email": user.Email}).Decode(&existingUser)

	log.Println(existingUser, user)

	if userepo.pass_service.PasswordComparator(existingUser.Password, user.Password) {
		return "", result, domain.ErrUnauthorizedAccess
	}

	jwtToken, err := userepo.pass_service.TokenGenerator(existingUser.ID, existingUser.Email, existingUser.IsAdmin)

	if err != nil {
		return "", result, domain.ErrInternalServerError
	}

	return jwtToken, existingUser, nil

}

func (userepo *UserRepository) DeleteUser(id string) *domain.AppError {
	collection := userepo.collection

	ido, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": ido}

	result, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil || result == 0 {
		return domain.ErrUserNotFound
	}

	return nil

}
