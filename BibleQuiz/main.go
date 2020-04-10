package main

import (
	"BibleQuiz/controllers"
	"BibleQuiz/models"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	r := gin.New()
	r.Use(gin.Logger())
	r.LoadHTMLGlob("BibleQuiz/view/*.html")
	db := models.SetupModels()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	r.GET("/", controllers.MainPage)
	r.POST("/next", controllers.ValidateNames)
	r.POST("/submit", controllers.Form)
	r.Run(":" + port)
}
