package service

import (
	"github.com/ivanreyess/bookstore_users-api/model/user"
	"github.com/ivanreyess/bookstore_users-api/utils/errors"
)

//CreateUser creates a new user reference
func CreateUser(u user.User) (*user.User, *errors.RestErr) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	return &u, nil
}
