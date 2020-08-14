package user

import (
	"fmt"
	"strings"

	"github.com/ivanreyess/bookstore_users-api/datasource/mysql/userdb"
	"github.com/ivanreyess/bookstore_users-api/utils/dateutils"
	"github.com/ivanreyess/bookstore_users-api/utils/errors"
)

const (
	queryInsertUser  = "INSERT INTO user(first_name, last_name, email, date_created) VALUES(?,?,?,?)"
	indexEmailUnique = "email_UNIQUE"
)

var (
	usersDB = make(map[int64]*User)
)

//Get a single user given its ID
func (u *User) Get() (*User, *errors.RestErr) {
	if err := userdb.Client.Ping(); err != nil {
		panic(err)
	}
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
	stmt, err := userdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	u.DateCreated = dateutils.GetNowString()
	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), indexEmailUnique) {
			return errors.NewInternalServerError(fmt.Sprintf("email %s already exists", u.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error trying to save user: %s", err.Error()))
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error trying to save user: %s", err.Error()))
	}
	u.ID = userID
	if current != nil {
		if current.Email == u.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", u.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", u.ID))
	}
	return nil
}
