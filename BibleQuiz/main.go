package main

import (
	"BibleQuiz/controllers"
	"BibleQuiz/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db := models.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/", controllers.MainPage)
	r.POST("/next", controllers.ValidateNames)
	r.POST("/submit", controllers.Form)
	r.Run(":9090")
}
