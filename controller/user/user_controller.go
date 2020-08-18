package user

import (
	"net/http"
	"strconv"

	"github.com/ivanreyess/bookstore_users-api/utils/errors"

	"github.com/gin-gonic/gin"
	"github.com/ivanreyess/bookstore_users-api/model/user"
	"github.com/ivanreyess/bookstore_users-api/service"
)

//getUserIDParam get and validates user ID parameter
func getUserIDParam(userIDParam string) (int64, *errors.RestErr) {
	userID, err := strconv.ParseInt(userIDParam, 10, 64)
	if err != nil {
		err := errors.NewBadRequestError("user ID should be a number")
		return 0, err
	}
	return userID, nil
}

//Get get a user given its ID
func Get(c *gin.Context) {
	userID, err := getUserIDParam(c.Param("user_id"))
	if err != nil {
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

//Create create a new user
func Create(c *gin.Context) {
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

//Update update a new user with a given ID
func Update(c *gin.Context) {
	userID, err := getUserIDParam(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	var u user.User
	if err := c.ShouldBindJSON(&u); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	u.ID = userID

	isPartial := c.Request.Method == http.MethodPatch

	result, updateErr := service.UpdateUser(isPartial, u)
	if updateErr != nil {
		c.JSON(updateErr.Status, updateErr)
		return
	}
	c.JSON(http.StatusOK, result)
}

//Delete remove user with a given ID
func Delete(c *gin.Context) {
	userID, err := getUserIDParam(c.Param("user_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	userErr := service.DeleteUser(userID)
	if userErr != nil {
		c.JSON(userErr.Status, userErr)
		return
	}
	c.JSON(http.StatusOK, map[string] string{"status": "deleted"})
}

//Search get all users by a given status
func Search(c *gin.Context){
	status := c.Query("status")
	users, err := service.Search(status)
	if err !=nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, users)
}