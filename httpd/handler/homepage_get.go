package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage() gin.HandlerFunc {

	return func(c *gin.Context) {
		sid, err := c.Cookie("session")
		if err != nil || sid == "" {
			c.Redirect(302, "/login")
		} else {
			sessionname, _ := c.Cookie("session")
			c.HTML(http.StatusOK, "homepage.gohtml", sessionname)
		}
	}
}
