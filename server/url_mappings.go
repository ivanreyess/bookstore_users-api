package server

import (
	"github.com/ivanreyess/bookstore_users-api/controller/ping"
	"github.com/ivanreyess/bookstore_users-api/controller/user"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)
	router.POST("/users", user.Create)
	router.GET("/users/:user_id", user.Get)
	router.PUT("/users/:user_id", user.Update)
	router.PATCH("/users/:user_id", user.Update)
	router.DELETE("/users/:user_id", user.Delete)
}
