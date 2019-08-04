package handler

import (
	"LoginApp/platform/signup"
	"fmt"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func LoginPost(user *signup.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {

		sid, err := c.Cookie("session")
		fmt.Println("err,sid:", err, sid)
		if err != nil || sid == "" {
			var msg string
			var flag int
			fmt.Println("err,sid::", err, sid)
			username := c.PostForm("username")
			password := c.PostForm("password")
			fmt.Println("username and password", username, password)
			for _, v := range user.Users {

				fmt.Println("in for loop")
				fmt.Println("before checking condition:", username, v.Email)
				if username == v.Email {
					fmt.Println("v.email", v.Email)
					// password matching with hashed password
					err = bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(password))
					fmt.Println("password err : ", err)
					if err == nil {
						c.SetCookie("session", "sidnumber", 600, "/", "", false, false)
						fmt.Println("cookie set redirecting to homepage")
						fmt.Println("before redirecting", username, password, v.Email, v.Password)
						c.Redirect(303, "/homepage")
						break
					} else {
						msg = "password not matched"
						return
					}

				} else {
					msg = "username not found"
					fmt.Println(msg)
				}

			}
			if flag == 1 {
				c.JSON(404, msg)
			} else if flag == 2 {

			}
		} else {
			c.Redirect(302, "/homepage")
		}

	}
}
