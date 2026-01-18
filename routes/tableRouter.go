package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)
func TableRoutes(router *gin.Engine) {
	router.GET("/tables" , controllers.GetTables())
	router.GET("/tables/:table_id", controllers.GetTableById())
	router.POST("/tables", controllers.CreateTable())
	router.PATCH("/tables/:table_id", controllers.UpdateTable())
}