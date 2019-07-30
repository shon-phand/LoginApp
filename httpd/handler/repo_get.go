package handler

import (
	"LoginApp/platform/signup"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllRegisterUsers(user *signup.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := user.GetAllRegisterd()
		c.JSON(http.StatusOK, results)
	}
}
