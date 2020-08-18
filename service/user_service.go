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
	if err := u.Save(); err != nil {
		return nil, err
	}
	return &u, nil
}

//GetUser return an user given its ID
func GetUser(userID int64) (*user.User, *errors.RestErr) {
	if userID <= 0 {
		return nil, errors.NewBadRequestError("ID must be greater than 0")
	}
	u := user.User{ID: userID}
	return u.Get()
}

//UpdateUser updates user
func UpdateUser(isPartial bool, u user.User) (*user.User, *errors.RestErr) {
	current, err := GetUser(u.ID)
	if err != nil {
		return nil, err
	}
	if err := current.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if u.FirstName != "" {
			current.FirstName = u.FirstName
		}
		if u.LastName != "" {
			current.LastName = u.LastName
		}
		if u.Email != "" {
			current.Email = u.Email
		}

	} else {
		current.FirstName = u.FirstName
		current.LastName = u.LastName
		current.Email = u.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}
	return current, nil
}

//DeleteUser remove user
func DeleteUser(userID int64) *errors.RestErr {
	user := &user.User{ID: userID}
	return user.Delete()
}
