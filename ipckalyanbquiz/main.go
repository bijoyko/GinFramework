package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var Name Names
var Score int

type Quiz struct {
	ID        int    `json:"id" gorm:"primary_key"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Language  string `json:"language"`
	Q1        string `json:"q1"`
	Q2        string `json:"q2"`
	Q3        string `json:"q3"`
	Q4        string `json:"q4"`
	Q5        string `json:"q5"`
	Q6        string `json:"q6"`
	Q7        string `json:"q7"`
	Q8        string `json:"q8"`
	Q9        string `json:"q9"`
	Q10       string `json:"q10"`
}

type EnterQuizData struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Language  string `json:"language" binding:"required"`
	Q1        string `json:"q1" binding:"required"`
	Q2        string `json:"q2" binding:"required"`
	Q3        string `json:"q3" binding:"required"`
	Q4        string `json:"q4" binding:"required"`
	Q5        string `json:"q5" binding:"required"`
	Q6        string `json:"q6" binding:"required"`
	Q7        string `json:"q7" binding:"required"`
	Q8        string `json:"q8" binding:"required"`
	Q9        string `json:"q9" binding:"required"`
	Q10       string `json:"q10" binding:"required"`
}

type Names struct {
	Firstname string
	Lastname  string
	Language  string
}

type Answers struct {
	Q1  string `json:"q1" binding:"required"`
	Q2  string `json:"q2" binding:"required"`
	Q3  string `json:"q3" binding:"required"`
	Q4  string `json:"q4" binding:"required"`
	Q5  string `json:"q5" binding:"required"`
	Q6  string `json:"q6" binding:"required"`
	Q7  string `json:"q7" binding:"required"`
	Q8  string `json:"q8" binding:"required"`
	Q9  string `json:"q9" binding:"required"`
	Q10 string `json:"q10" binding:"required"`
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("*.html")
	db := SetupModels()
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.GET("/", MainPage)
	router.POST("/next", ValidateNames)
	router.POST("/submit", Form)
	router.Run(":" + port)
}

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "youth_bijoy:Mavericklolol123@tcp(ipckalyan.com:3306)/youth_quiz?charset=utf8mb4&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connection established")
	return db
}

func MainPage(c *gin.Context) {
	t, _ := template.ParseFiles("main.html")
	t.Execute(c.Writer, nil)
}

func ValidateNames(c *gin.Context) {
	if c.PostForm("Firstname") != "" && c.PostForm("Lastname") != "" && c.PostForm("Language") != "" {
		Name = Names{
			Firstname: c.PostForm("Firstname"),
			Lastname:  c.PostForm("Lastname"),
			Language:  c.PostForm("Language"),
		}
		if c.PostForm("Language") == "English" {
			t, _ := template.ParseFiles("FormEnglish.html")
			t.Execute(c.Writer, nil)
		} else {
			t, _ := template.ParseFiles("FormMalayalam.html")
			t.Execute(c.Writer, nil)
		}

	} else {
		text1 := "Please enter a valid first Name and last name. Also select a language of your choice to proceed\t Blank values are not allowed"
		t, _ := template.ParseFiles("main.html")
		t.Execute(c.Writer, text1)
	}
}
func Form(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	input := EnterQuizData{
		Firstname: Name.Firstname,
		Lastname:  Name.Lastname,
		Language:  Name.Language,
		Q1:        c.PostForm("Q1"),
		Q2:        c.PostForm("Q2"),
		Q3:        c.PostForm("Q3"),
		Q4:        c.PostForm("Q4"),
		Q5:        c.PostForm("Q5"),
		Q6:        c.PostForm("Q6"),
		Q7:        c.PostForm("Q7"),
		Q8:        c.PostForm("Q8"),
		Q9:        c.PostForm("Q9"),
		Q10:       c.PostForm("Q10"),
	}

	quizD := Quiz{Firstname: input.Firstname, Lastname: input.Lastname, Language: input.Language, Q1: input.Q1, Q2: input.Q2,
		Q3: input.Q3, Q4: input.Q4, Q5: input.Q5, Q6: input.Q6, Q7: input.Q7, Q8: input.Q8, Q9: input.Q9, Q10: input.Q10}
	db.Table("quiztable").Create(&quizD)

	//answers submitted by the client
	checker := []string{c.PostForm("Q1"), c.PostForm("Q2"), c.PostForm("Q3"), c.PostForm("Q4"), c.PostForm("Q5"), c.PostForm("Q6"), c.PostForm("Q7"),
		c.PostForm("Q8"), c.PostForm("Q9"), c.PostForm("Q10")}

	var quiz Quiz
	// id=1 or 2 or ... based on the quiz and entry in answerquiz table
	if err := db.Table("answersquiz").Where("id = 1").First(&quiz).Error; err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	//answers pulled up from database which has right answers
	givenAnswers := []string{quiz.Q1, quiz.Q2, quiz.Q3, quiz.Q4, quiz.Q5, quiz.Q6, quiz.Q7, quiz.Q8, quiz.Q9, quiz.Q10}
	Score = 0
	// compare answers given by client with the right answers and increase the score
	for i, _ := range checker {
		if givenAnswers[i] == checker[i] {
			Score++
		}
	}
	t, _ := template.ParseFiles("ThankYou.html")
	t.Execute(c.Writer, Score)

}
