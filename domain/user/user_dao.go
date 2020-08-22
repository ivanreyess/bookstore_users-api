package user

import (
	"fmt"

	"github.com/ivanreyess/bookstore_users-api/datasource/mysql/userdb"
	"github.com/ivanreyess/bookstore_users-api/logger"
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
func (u *User) Get() *errors.RestErr {
	stmt, err := userdb.Client.Prepare(queryGetUser)
	if err != nil {
		logger.Error("error when trying to prepare get user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	result := stmt.QueryRow(u.ID)
	if err := result.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.DateCreated, &u.Status); err != nil {
		logger.Error("error when trying to get user by id", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

//Save user in database
func (u *User) Save() *errors.RestErr {
	stmt, err := userdb.Client.Prepare(queryInsertUser)
	if err != nil {
		logger.Error("error when trying to prepare insert user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	insertResult, err := stmt.Exec(u.FirstName, u.LastName, u.Email, u.DateCreated, u.Status, u.Password)
	if err != nil {
		logger.Error("error when trying execute insert user statement", err)
		return errors.NewInternalServerError("database error")
	}
	userID, err := insertResult.LastInsertId()
	if err != nil {
		logger.Error("error when trying to get last user id after creating user", err)
		return mysqlutils.ParseError(err)
	}
	u.ID = userID
	return nil
}

//Update user in database
func (u *User) Update() *errors.RestErr {
	stmt, err := userdb.Client.Prepare(queryUpdateUser)
	if err != nil {
		logger.Error("error when trying to prepare update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.FirstName, u.LastName, u.Email, u.ID)
	if err != nil {
		logger.Error("error when trying to execute update user statement", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

//Delete user in database
func (u *User) Delete() *errors.RestErr {
	stmt, err := userdb.Client.Prepare(queryDeleteUser)
	if err != nil {
		logger.Error("error when trying to prepare delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.ID)
	if err != nil {
		logger.Error("error when trying to execute delete user statement", err)
		return errors.NewInternalServerError("database error")
	}
	return nil
}

//FindByStatus get a user slice given the status
func (u *User) FindByStatus(status string) ([]User, *errors.RestErr) {
	stmt, err := userdb.Client.Prepare(queryFindUserByStatus)
	if err != nil {
		logger.Error("error when trying to prepare find user by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer stmt.Close()
	rows, err := stmt.Query(status)
	if err != nil {
		logger.Error("error when trying to execute find user by status statement", err)
		return nil, errors.NewInternalServerError("database error")
	}
	defer rows.Close()
	results := make([]User, 0)
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status); err != nil {
			logger.Error("error when trying to scan rows from find user by status statement", err)
			return nil, errors.NewInternalServerError("database error")
		}
		results = append(results, user)
	}
	if len(results) == 0 {
		//no log over here is internal
		return nil, errors.NewNotFoundError(fmt.Sprintf("no users matching status %s", status))
	}
	return results, nil
}
