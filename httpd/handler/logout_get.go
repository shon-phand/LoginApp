package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		sid, _ := c.Cookie("session")
		fmt.Println("sid = ", sid)
		c.SetCookie("session", "", -1, "/", "", false, false)
		sid2, _ := c.Cookie("session")
		fmt.Println("sid2 = ", sid2)
		c.Redirect(302, "/login")
	}
}
