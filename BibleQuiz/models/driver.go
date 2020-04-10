package models

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "youth_bijoy:Mavericklolol123@tcp(ipckalyan.com:3306)/youth_quiz?charset=utf8mb4&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Connection established")
	//db.Table("quiz-data").AutoMigrate(&models.Quiz{})
	return db
}
