package handler

import (
	"LoginApp/platform/signup"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AllRegisterUsers(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		results := signup.GetAllRegisterdInDB(db)
		//fmt.Println(results)
		c.JSON(http.StatusOK, results)
	}
}
