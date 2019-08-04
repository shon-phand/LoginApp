package handler

import (
	"LoginApp/platform/login"
	"database/sql"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func LoginPost(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		sid, err := c.Cookie("session")
		//	fmt.Println("err,sid:", err, sid)
		if err != nil || sid == "" {
			//fmt.Println("err,sid::", err, sid)
			username := c.PostForm("username")
			password := c.PostForm("password")

			user, err := login.GetUserByUsername(username, db)

			if err == nil {

				err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
				if err == nil {

					cookie := uuid.NewV4().String()
					c.SetCookie("session", cookie, 300, "/", "", false, true)
					c.Redirect(303, "/homepage?sid="+cookie)

				} else {
					c.String(400, "password not matched")
				}

			} else {
				c.JSON(404, "username not found")
			}

		}
	}
}
