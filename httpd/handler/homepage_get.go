package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Homepage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.HTML(http.StatusOK, "homepage.gohtml", nil)
	}
}
