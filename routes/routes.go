package routes

import (
	"80HW/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/items", controllers.GetItems)
	router.POST("/items", controllers.CreateItem)
	router.GET("/items/:id", controllers.GetItem)
	router.PUT("/items/:id", controllers.UpdateItem)
	router.DELETE("/items/:id", controllers.DeleteItem)
}
