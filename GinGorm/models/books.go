package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}

type CreateBookInput struct {
	ID     uint   `json:"id"`
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	Year   string `json:"year"`
}

//this struct is used as a schema in the updatebooks patch function
type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	Year   string `json:"year"`
}
