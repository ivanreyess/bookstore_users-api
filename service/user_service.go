package service

import (
	"github.com/ivanreyess/bookstore_users-api/model/user"
	"github.com/ivanreyess/bookstore_users-api/utils/errors"
)

//CreateUser creates a new user reference
func CreateUser(user user.User) (*user.User, *errors.RestErr) {
	return &user, nil
}
