package domain

import (
	"errors"
	"net/http"
)

// AppError is a custom error type that includes an error message and an associated HTTP status code.
type AppError struct {
	message    string
	statusCode int
	err        error
}

// Error implements the error interface.
func (e *AppError) Status() int {
	return e.statusCode
}

// Error implements the error interface.
func (e *AppError) Message() string {
	return e.message
}

// Unwrap allows the underlying error to be accessible.
func (e *AppError) Unwrap() error {
	return e.err
}

// NewAppError creates a new AppError.
func NewAppError(message string, statusCode int, err error) *AppError {
	return &AppError{
		message:    message,
		statusCode: statusCode,
		err:        err,
	}
}

var (
	ErrTaskNotFound = NewAppError("task not found", http.StatusNotFound, errors.New("task not found"))

	ErrUnauthorizedAccess = NewAppError("you don't have the privilege to perform this action", http.StatusForbidden, errors.New("unauthorized access"))

	ErrNoTasksFound = NewAppError("no tasks found", http.StatusNotFound, errors.New("no tasks found"))

	ErrTaskInsertionFailed = NewAppError("failed to insert the task", http.StatusInternalServerError, errors.New("task insertion failed"))

	ErrTaskUpdateFailed = NewAppError("failed to update the task", http.StatusInternalServerError, errors.New("task update failed"))

	ErrTaskDeletionFailed = NewAppError("failed to delete the task", http.StatusInternalServerError, errors.New("task deletion failed"))

	//user errors
	ErrUserExists = NewAppError("user already exists with the same email", http.StatusBadRequest, errors.New("user exists"))

	ErrInvalidCredentials = NewAppError("invalid email or password", http.StatusUnauthorized, errors.New("invalid credentials"))

	ErrUserNotFound = NewAppError("user not found", http.StatusNotFound, errors.New("user not found"))

	ErrUserDeletionFailed = NewAppError("failed to delete the user", http.StatusInternalServerError, errors.New("user deletion failed"))

	ErrUserRegistrationFailed = NewAppError("failed to register the user", http.StatusInternalServerError, errors.New("user registration failed"))

	ErrInternalServerError = NewAppError("internal server error", http.StatusInternalServerError, errors.New("internal server error"))
)
