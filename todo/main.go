package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

type Person struct {
	ID        uint   `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func main() {
	const PORT = "8009"
	db, err := gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Printf("could not create sqlitedb %v", err)
	}
	defer db.Close()
	db.AutoMigrate(&Person{})

	r := gin.Default()
	r.GET("/", GetProjects)
	r.POST("/people", CreatePerson)
	r.Run(":" + PORT)
}

func GetProjects(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "world"})
}

func CreatePerson(c *gin.Context) {
	// var person Person
	// c.BindJSON(&person)
	// db.Create(&person)
	// c.JSON(200, person)
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "handler not implemented yet",
	})
}
