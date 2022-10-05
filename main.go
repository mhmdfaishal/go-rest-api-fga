package main

import (
	"rest-api-go/config"
	"rest-api-go/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := config.DBInit()

	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.GET("/person/:id", inDB.GetPerson)
	router.GET("/persons", inDB.GetPersons)
	router.POST("/person", inDB.CreatePerson)
	router.PUT("/person/:id", inDB.UpdatePerson)
	router.DELETE("/person/:id", inDB.DeletePerson)

	router.Run(":3000")
}
