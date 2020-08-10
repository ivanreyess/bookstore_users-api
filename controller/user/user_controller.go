package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ivanreyess/bookstore_users-api/model/user"
	"io/ioutil"
	"net/http"
)

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}

func CreateUser(c *gin.Context) {
	var u user.User
	fmt.Println(u)
	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusInternalServerError, "could not read user")
	}
	err = json.Unmarshal(bytes, &u)
	if err != nil {
		c.String(http.StatusInternalServerError, "could not unmarshal user")
	}
	fmt.Println(err)
	c.String(http.StatusNotImplemented, "implement me!")
}
