package handlers

import (
	"fmt"
	"log"
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
	err := c.BindJSON(&user)
	if err != nil {
		log.Fatal(err)
	}

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

// GetSinglePerson get a single person by Id
func GetSinglePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := common.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		c.JSON(http.StatusOK, gin.H{"single person": person})
	}
}

// UpdateSinglePerson update a single Person record
func UpdateSinglePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := common.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		err := c.BindJSON(&person)
		if err != nil {
			log.Fatal(err)
		}
		common.DB.Save(&person)
		c.JSON(http.StatusOK, gin.H{"updated person": person})
	}
}

// DeleteSinglePerson delete a single person record
func DeleteSinglePerson(c *gin.Context) {
	id := c.Param("id")
	var person models.Person
	if err := common.DB.Where("id = ?", id).First(&person).Error; err != nil {
		c.JSON(http.StatusNotFound,
			gin.H{"message": "user with Id does not exist"})
	} else {
		d := common.DB.Where("id = ?", id).Delete(&person)
		c.JSON(http.StatusOK, gin.H{"deleted user with id " + id: d})
	}
}
