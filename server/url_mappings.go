package server

import (
	"github.com/ivanreyess/bookstore_users-api/controller/ping"
	"github.com/ivanreyess/bookstore_users-api/controller/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", user.CreateUser)
	router.GET("/users/:user_id", user.GetUser)
}
