package errors

import "net/http"

//RestErr defines common structure for handling error messages
type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"code"`
	Error   string `json:"error"`
}

//NewBadRequestError shortcut for bad request message
func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
