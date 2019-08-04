package handler

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PasswordReset(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		c1, err := c.Cookie("selfserve")
		newpass := c.PostForm("newpass")
		if err == nil {
			stmt, err := db.Prepare("update users set password= ? where email= ?;")
			if err != nil {
				fmt.Print(err.Error())
			}
			password, err := bcrypt.GenerateFromPassword([]byte(newpass), 14)
			_, err = stmt.Exec(password, c1)
			if err != nil {
				fmt.Print(err.Error())
				c.HTML(http.StatusOK, "login.gohtml", "password reset failed")
			} else {
				c.SetCookie("selfserve", "", -1, "/", "", false, false)
				msg := "password reset successfully"
				c.HTML(303, "login.gohtml", msg)
			}
		} else {
			c.String(400, "please verify user username")
		}
	}
}
