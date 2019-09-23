package routers

import (
	"github.com/danny/todo/todo/handlers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// APIBase base route of api
const APIBase = "api/v1/todos"

// SetupRouter setup gin router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group(APIBase)
	{
		v1.GET("/", handlers.SanityCheck)
		v1.GET("/people", handlers.GetPeople)
		v1.POST("/people", handlers.CreatePerson)
		v1.GET("/people/:id", handlers.GetSinglePerson)
		v1.PUT("/people/:id", handlers.UpdateSinglePerson)
		v1.DELETE("/people/:id", handlers.DeleteSinglePerson)
	}
	r.Use(cors.Default())
	return r
}
