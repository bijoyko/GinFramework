package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("mysql", "root:Lolol_123@tcp(127.0.0.1:3306)/localdevelopment?charset=utf8mb4&parseTime=True&loc=Local")
	db.SingularTable(true)

	//defer db.Close()

	if err != nil {
		fmt.Println("Connection Failed to Open")
	}
	fmt.Println("Connection Established")
	return db
}
