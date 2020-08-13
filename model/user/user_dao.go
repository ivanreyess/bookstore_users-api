package user

import (
	"fmt"

	"github.com/ivanreyess/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

//Get a single user given its ID
func (u *User) Get() (*User, *errors.RestErr) {
	result := usersDB[u.ID]
	if result == nil {
		return nil, errors.NewNotFoundError(fmt.Sprintf("user %d not found", u.ID))
	}
	u.FirstName = result.FirstName
	u.LastName = result.LastName
	u.Email = result.Email
	u.DateCreated = result.DateCreated
	return u, nil
}

//Save user in database
func (u *User) Save() *errors.RestErr {
	current := usersDB[u.ID]
	if current != nil {
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", u.ID))
	}
	usersDB[u.ID] = u
	return nil
}
