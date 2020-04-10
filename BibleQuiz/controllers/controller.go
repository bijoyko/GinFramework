package controllers

import (
	"BibleQuiz/models"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

var Name models.Names
var Score int

func MainPage(c *gin.Context) {
	t, _ := template.ParseFiles("view/main.html")
	t.Execute(c.Writer, nil)
}

func ValidateNames(c *gin.Context) {
	if c.PostForm("Firstname") != "" && c.PostForm("Lastname") != "" && c.PostForm("Language") != "" {
		Name = models.Names{
			Firstname: c.PostForm("Firstname"),
			Lastname:  c.PostForm("Lastname"),
			Language:  c.PostForm("Language"),
		}
		if c.PostForm("Language") == "English" {
			t, _ := template.ParseFiles("view/FormEnglish.html")
			t.Execute(c.Writer, nil)
		} else {
			t, _ := template.ParseFiles("view/FormMalayalam.html")
			t.Execute(c.Writer, nil)
		}

	} else {
		text1 := "Please enter a valid first Name and last name. Also select a language of your choice to proceed\t Blank values are not allowed"
		t, _ := template.ParseFiles("view/main.html")
		t.Execute(c.Writer, text1)
	}
}
func Form(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)

	input := models.EnterQuizData{
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

	quizD := models.Quiz{Firstname: input.Firstname, Lastname: input.Lastname, Language: input.Language, Q1: input.Q1, Q2: input.Q2,
		Q3: input.Q3, Q4: input.Q4, Q5: input.Q5, Q6: input.Q6, Q7: input.Q7, Q8: input.Q8, Q9: input.Q9, Q10: input.Q10}
	db.Table("quiztable").Create(&quizD)

	//answers submitted by the client
	checker := []string{c.PostForm("Q1"), c.PostForm("Q2"), c.PostForm("Q3"), c.PostForm("Q4"), c.PostForm("Q5"), c.PostForm("Q6"), c.PostForm("Q7"),
		c.PostForm("Q8"), c.PostForm("Q9"), c.PostForm("Q10")}

	var quiz models.Quiz
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
	t, _ := template.ParseFiles("view/ThankYou.html")
	t.Execute(c.Writer, Score)

}
