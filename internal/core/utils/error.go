package utils

import (
	"errors"
	"net/http"
)

// Custom errors to be use accross the application
var (
	// ErrUnauthorized is an error for when the user is unauthorized
	ErrUnauthorized = errors.New("user is unauthorized to access the resource")
	// ErrForbidden is an error for when the user is forbidden to access the resource
	ErrForbidden = errors.New("user is forbidden to access the resource")
)

// HandleError handles the error and returns the appropriate http status code
var ErrorMapToHttp = map[error]int{
	ErrUnauthorized: http.StatusUnauthorized,
	ErrForbidden:    http.StatusForbidden,
}
