package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayhelloName(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Bijoy is taking over",
	})

}

func LoadLogin(c *gin.Context) {
	t, _ := template.ParseFiles("Login.html")
	t.Execute(c.Writer, nil)
}

func login(c *gin.Context) {
	a := c.PostForm("username")
	b := c.PostForm("password")
	fmt.Println(a)
	fmt.Println(b)
	t, _ := template.ParseFiles("ThankYou.html")
	t.Execute(c.Writer, nil)
}

func main() {
	r := gin.Default()
	r.GET("/", sayhelloName)
	r.GET("/login", LoadLogin)
	r.POST("/login", login)
	r.Run(":9090")
}
