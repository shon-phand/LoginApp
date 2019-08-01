package handler

import (
	"LoginApp/platform/signup"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginPost(user *signup.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {

		sid, err := c.Cookie("session")

		if err != nil || sid == "" {

			username := c.PostForm("username")
			password := c.PostForm("password")

			for _, v := range user.Users {
				if username == v.Email {
					// password matching with hashed password
					err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(password))
					if err == nil {
						c.SetCookie("session", "sidnumber", 600, "/", "", false, false)
						c.Redirect(301, "/homepage")
					} else {
						c.JSON(http.StatusNotFound, "password not matched")

					}

				} else {
					c.JSON(http.StatusNotFound, "username not found")

				}
			}
		} else {
			c.Redirect(302, "/homepage")
		}
	}
}
