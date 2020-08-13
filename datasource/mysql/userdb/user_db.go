package userdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysqlUserUsername = "root"
	mysqlUserPassword = "root"
	mysqlUserHost     = "127.0.0.1:3306"
	mysqlUserSchema   = "users_db"
)

var (
	//Client represents an user database connection
	Client *sql.DB
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", mysqlUserUsername, mysqlUserPassword, mysqlUserHost, mysqlUserSchema)
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")
}
