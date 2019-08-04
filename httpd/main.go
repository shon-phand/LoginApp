package main

import (
	"LoginApp/httpd/handler"
	"LoginApp/platform/DB"
	"LoginApp/platform/signup"

	"github.com/gin-gonic/gin"
)

func main() {

	user := signup.New()
	db := DB.Start()
	defer db.Close()
	r := gin.Default()
	r.LoadHTMLGlob("/home/shon/Documents/Go_practise/LoginApp/resources/template/*")
	r.GET("/", handler.WelcomePage())
	r.GET("/login", handler.LoginPage())
	r.GET("/signup", handler.RegistrationPage())
	r.GET("/homepage", handler.Homepage())
	r.GET("/repo", handler.AllRegisterUsers(db))
	r.GET("/logout", handler.Logout())

	r.POST("/signup", handler.RegistrationPost(db))
	r.POST("/login", handler.LoginPost(user))

	r.Run(":8181")
}
