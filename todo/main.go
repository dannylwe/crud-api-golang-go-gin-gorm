package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

// Person is a datastruct used to hold person info
type Person struct {
	ID        int    `gorm:"AUTO_INCREMENT" json:"id"`
	FirstName string `form:"firstname" json:"firstname"`
	LastName  string `form:"lastname" json:"lastname"`
}

func main() {
	const PORT = "8009"
	var err error
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Printf("could not create sqlitedb %v", err)
	}
	defer db.Close()
	db.AutoMigrate(&Person{})

	r := gin.Default()
	v1 := r.Group("api/v1/todos")
	{
		v1.GET("/", SanityCheck)
		v1.GET("/people", GetPeople)
		v1.POST("/people", CreatePerson)
	}

	r.Run(":" + PORT)
}

// SanityCheck checks whether server is running
func SanityCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "world"})
}

// GetPeople getall people records
func GetPeople(c *gin.Context) {
	var people []Person
	if err := db.Find(&people).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"list of people": people})
	}
}

// CreatePerson create a single person
func CreatePerson(c *gin.Context) {

	var user Person
	c.BindJSON(&user)

	if user.FirstName != "" && user.LastName != "" {
		// INSERT INTO "users" (name) VALUES (user.Name);
		db.Create(&user)
		// Display error
		c.JSON(http.StatusCreated, gin.H{"user created": user})
	} else {
		// Display error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
	}
}
