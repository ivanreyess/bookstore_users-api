package user

import (
	"fmt"

	"github.com/ivanreyess/bookstore_users-api/datasource/mysql/userdb"
	"github.com/ivanreyess/bookstore_users-api/utils/errors"
	"github.com/ivanreyess/bookstore_users-api/utils/mysqlutils"
)

const (
	queryInsertUser       = "INSERT INTO user(first_name, last_name, email, date_created, status, password) VALUES(?,?,?,?,?,?)"
	indexEmailUnique      = "email_UNIQUE"
	queryGetUser          = "SELECT u.id, u.first_name, u.last_name, u.email, u.date_created, status FROM user u WHERE u.id=?"
	queryFindUserByStatus = "SELECT u.id, u.first_name, u.last_name, u.email, u.date_created, status FROM user u WHERE u.status=?"
	queryUpdateUser       = "UPDATE user SET first_name=?, last_name=?, email=? WHERE id=?"
	queryDeleteUser       = "DELETE FROM user WHERE id=?"
	errorNoRows           = "no rows in result set"
)

//Get a single user given its ID
func (u *User) Get() (*User, *errors.RestErr) {
	if err := userdb.Client.Ping(); err != nil {
		panic(err)
	}
	stmt, err := userdb.Client.Prepare(queryGetUser)
	if err != nil {
		return nil, mysqlutils.ParseError(err)
	}
	defer stmt.Close()
	result := stmt.QueryRow(u.ID)
	if err := result.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); err != nil {
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
	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated, u.Status, u.Password)
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

//FindByStatus get a user slice given the status
func (u *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := userdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		return nil, errors.NewInternalServerError(err.Error())
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			return nil, mysqlutils.ParseError(err)
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
