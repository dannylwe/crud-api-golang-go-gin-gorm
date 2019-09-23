package handlers

import (
	"fmt"
	"net/http"

	"github.com/danny/todo/todo/common"
	"github.com/danny/todo/todo/models"
	"github.com/gin-gonic/gin"
)

// SanityCheck checks whether server is running
func SanityCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "world"})
}

// GetPeople getall people records
func GetPeople(c *gin.Context) {
	var people []models.Person
	if err := common.DB.Find(&people).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"list of people": people})
	}
}

// CreatePerson create a single person
func CreatePerson(c *gin.Context) {

	var user models.Person
	c.BindJSON(&user)

	if user.FirstName != "" && user.LastName != "" {
		// INSERT INTO "users" (name) VALUES (user.Name);
		common.DB.Create(&user)
		// Display error
		c.JSON(http.StatusCreated, gin.H{"user created": user})
	} else {
		// Display error
		c.JSON(http.StatusBadRequest, gin.H{"error": "Fields are empty"})
	}
}
