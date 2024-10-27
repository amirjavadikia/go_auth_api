package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() error {
	d, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/go_auth_api?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Failed to connect to database")
		return err
	}
	db = d
	return nil
}

func GetDb() *gorm.DB {
	return db
}
