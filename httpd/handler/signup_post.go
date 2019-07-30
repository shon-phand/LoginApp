package handler

import (
	"LoginApp/platform/signup"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegistrationPost(user *signup.Repo) gin.HandlerFunc {

	return func(c *gin.Context) {
		firstname := c.PostForm("first_name")
		lastname := c.PostForm("last_name")
		email := c.PostForm("email")
		password := c.PostForm("password")

		/* password hashing mechanism */
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

		newuser := signup.User{firstname, lastname, email, string(hashedPassword)}
		user.Register(newuser)
		c.JSON(http.StatusOK, "user created")
	}

}
