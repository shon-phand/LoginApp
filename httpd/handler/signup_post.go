package handler

import (
	"LoginApp/platform/signup"
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func RegistrationPost(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		firstname := c.PostForm("first_name")
		lastname := c.PostForm("last_name")
		email := c.PostForm("email")
		password := c.PostForm("password")

		/* password hashing mechanism */
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
		var ID = 3
		var newuser signup.Data
		newuser.ID = ID
		newuser.Firstname = firstname
		newuser.Lastname = lastname
		newuser.Email = email
		newuser.Password = string(hashedPassword)
		newuser.Create = time.Now().Format("Mon Jan _2 15:04:05 2006")
		newuser.Update = time.Now().Format("Mon Jan _2 15:04:05 2006")

		err := signup.RegisterInDB(newuser, db)
		if err == nil {
			c.HTML(http.StatusOK, "login.gohtml", nil)
		} else {
			c.JSON(http.StatusNoContent, "registration failed")
		}
		//newuser := signup.User{firstname, lastname, email, string(hashedPassword)}
		//user.Register(newuser)

	}

}
