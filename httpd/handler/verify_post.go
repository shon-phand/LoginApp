package handler

import (
	"LoginApp/platform/login"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyUsername(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.PostForm("username")

		exist := login.VerifyUsername(username, db)

		if exist {
			c.SetCookie("selfserve", username, 300, "/", "", false, true)
			c.HTML(http.StatusOK, "confirmpass.gohtml", username)
		} else {
			c.String(400, "username does not exist")
		}
	}
}
