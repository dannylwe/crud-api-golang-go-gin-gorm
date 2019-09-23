package routers

import (
	"github.com/danny/todo/todo/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api/v1/todos")
	{
		v1.GET("/", handlers.SanityCheck)
		v1.GET("/people", handlers.GetPeople)
		v1.POST("/people", handlers.CreatePerson)
	}
	return r
}
