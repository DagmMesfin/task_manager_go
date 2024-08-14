package infrastructure

import (
	domain "task-manager/Domain"

	"golang.org/x/crypto/bcrypt"
)

type PasswordService struct {
}

func NewPasswordService() domain.PasswordService {
	return &PasswordService{}

}

// compares the inputted password from the existing hash
func (ps *PasswordService) PasswordComparator(hash string, password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil
}

// hashes the password with a SHA-256 encryption
func (ps *PasswordService) PasswordHasher(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
