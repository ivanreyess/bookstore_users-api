package mysqlutils

import (
	"fmt"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/ivanreyess/bookstore_users-api/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

//ParseError formats mysql errors in a standard way
func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewNotFoundError("no record matching given ID")
		}
		return errors.NewInternalServerError("error parsing databse response")
	}
	switch sqlErr.Number {
	case 1062:
		return errors.NewBadRequestError(fmt.Sprintf("invalida data"))
	}
	return errors.NewInternalServerError("error procesing request")
}
