package user

import (
	"github.com/ivanreyess/bookstore_users-api/datasource/mysql/userdb"
	"github.com/ivanreyess/bookstore_users-api/utils/dateutils"
	"github.com/ivanreyess/bookstore_users-api/utils/errors"
	"github.com/ivanreyess/bookstore_users-api/utils/mysqlutils"
)

const (
	queryInsertUser  = "INSERT INTO user(first_name, last_name, email, date_created) VALUES(?,?,?,?)"
	indexEmailUnique = "email_UNIQUE"
	queryGetUser     = "SELECT u.id, u.first_name, u.last_name, u.email, u.date_created FROM user u WHERE u.id=?"
	queryUpdateUser  = "UPDATE user SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser  = "DELETE FROM user WHERE id=?"
	errorNoRows      = "no rows in result set"
)

//Get a single user given its ID
func (u *User) Get() (*User, *errors.RestErr) {
	stmt, err := userdb.Client.Prepare(queryGetUser)
	if err != nil {
		return nil, mysqlutils.ParseError(err)
	}
	defer stmt.Close()
	result := stmt.QueryRow(u.ID)
	if err := result.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated); err != nil {
		return nil, mysqlutils.ParseError(err)
	}
	return u, nil
}

//Save user in database
func (u *User) Save() *errors.RestErr {
	stmt, err := userdb.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	u.DateCreated = dateutils.GetNowString()
	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated)
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	u.ID = userID
	return nil
}

//Update user in database
func (u *User) Update() *errors.RestErr {
	stmt, err := userdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.ID)
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}

//Delete user in database
func (u *User) Delete() *errors.RestErr {
	stmt, err := userdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.ID)
	if err != nil {
		return mysqlutils.ParseError(err)
	}
	return nil
}
