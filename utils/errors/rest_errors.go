package errors

import (
	"errors"
	"net/http"
)

//NewError returns a new instance of an error
func NewError(msg string) error {
	return errors.New(msg)
}

//RestErr defines common structure for handling error messages
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

//NewBadRequestError shortcut for bad error message
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

//NewNotFoundError shortcut for not found error message
func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "bad_request",
	}
}

//NewInternalServerError shortcut for internal server error message
func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
