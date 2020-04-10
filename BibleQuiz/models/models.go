package models

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
