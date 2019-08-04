package main

import (
	"LoginApp/httpd/handler"
	"LoginApp/platform/DB"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	//ser := signup.New()
	db, err := DB.Start()

	if err != nil {
		fmt.Printf(err.Error())
	}

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
	r.POST("/login", handler.LoginPost(db))

	r.Run(":8182")
}
