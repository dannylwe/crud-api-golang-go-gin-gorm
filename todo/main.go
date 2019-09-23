package main

import (
	"fmt"

	"github.com/danny/todo/todo/common"
	"github.com/danny/todo/todo/handlers"
	"github.com/danny/todo/todo/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Person is a datastruct used to hold person info

func main() {

	var err error
	common.DB, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Printf("could not create sqlitedb %v", err)
	}
	defer common.DB.Close()
	common.DB.AutoMigrate(&models.Person{})

	r := gin.Default()
	v1 := r.Group("api/v1/todos")
	{
		v1.GET("/", handlers.SanityCheck)
		v1.GET("/people", handlers.GetPeople)
		v1.POST("/people", handlers.CreatePerson)
	}

	r.Run(":" + common.PORT)
}
