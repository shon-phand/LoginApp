package main

import (
	"LoginApp/httpd/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.LoadHTMLGlob("/home/shon/Documents/Go_practise/LoginApp/resources/template/*")
	r.GET("/", handler.Homepage())
	r.GET("/login", handler.LoginPage())
	r.GET("/signup", handler.RegistrationPage())

	r.Run(":8083")
}
