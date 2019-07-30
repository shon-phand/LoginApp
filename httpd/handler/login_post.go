package handler

import (
	"LoginApp/platform/signup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPost(user *signup.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")

		for _, v := range user.Users {
			if username == v.Email {
				if password == v.Password {
					c.HTML(http.StatusOK, "index.gohtml", nil)
				} else {
					c.JSON(http.StatusNotFound, "password not matched")

				}

			} else {
				c.JSON(http.StatusNotFound, "username not found")

			}
		}
	}
}
