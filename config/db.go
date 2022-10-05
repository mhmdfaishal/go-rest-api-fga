package config

import (
	"fmt"
	"rest-api-go/structs"

	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB{
	dbHost := "localhost"
	dbPort := "3306"
	dbUser := "root"
	dbPass := ""
	dbName := "go_rest"

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open("mysql",dsn)
	if err != nil {
		panic("Failed to connect to Database")
	}
	
	db.AutoMigrate(structs.Person{})

	return db
}