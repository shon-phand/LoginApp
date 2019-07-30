package handler

import (
	"LoginApp/platform/signup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistrationPost(user *signup.Repo) gin.HandlerFunc {

	return func(c *gin.Context) {
		firstname := c.PostForm("first_name")
		lastname := c.PostForm("last_name")
		email := c.PostForm("email")
		password := c.PostForm("password")

		newuser := signup.User{firstname, lastname, email, password}
		user.Register(newuser)
		c.JSON(http.StatusOK, "user created")
	}

}
