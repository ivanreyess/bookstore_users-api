package user

import (
	"net/http"
	"strconv"

	"github.com/ivanreyess/bookstore_users-api/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/ivanreyess/bookstore_users-api/model/user"
	"github.com/ivanreyess/bookstore_users-api/service"
)

//GetUser get a user given its ID
func GetUser(c *gin.Context) {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("user ID should be a number")
		c.JSON(err.Status, err)
		return
	}
	u, userErr := service.GetUser(userID)
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	c.JSON(http.StatusOK, u)
}

//CreateUser create a new user
func CreateUser(c *gin.Context) {
	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	result, saveErr := service.CreateUser(u)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}
