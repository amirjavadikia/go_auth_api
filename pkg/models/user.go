package models

import (
	"github.com/amirjavadi/go_auth_api/pkg/config"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
	Password string `json:"password"`
	gorm.Model
}

func init() {
	var err error
	err = config.Connect()
	if err != nil {
		log.Printf("Could not connect to database: %v", err)
	}
	db := config.GetDb()
	err = db.AutoMigrate(&User{}).Error
	if err != nil {
		log.Printf("Could not migrate database: %v", err)
	}
	log.Println("Database connected and migrated successfully!")
}
