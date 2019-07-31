package main

import (
	"LoginApp/httpd/handler"
	"LoginApp/platform/signup"

	"github.com/gin-gonic/gin"
)

func main() {

	user := signup.New()

	r := gin.Default()
	r.LoadHTMLGlob("/home/shon/Documents/Go_practise/LoginApp/resources/template/*")
	r.GET("/", handler.WelcomePage())
	r.GET("/login", handler.LoginPage())
	r.GET("/signup", handler.RegistrationPage())
	r.GET("/homepage", handler.Homepage())
	r.GET("/repo", handler.AllRegisterUsers(user))

	r.POST("/signup", handler.RegistrationPost(user))
	r.POST("/login", handler.LoginPost(user))

	r.Run(":8080")
}
