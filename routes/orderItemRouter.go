package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)
func OrderItemRoutes(router *gin.Engine) {
	router.GET("/orderItems" , controllers.GetOrderItems())
	router.GET("/orderItems/:orderItem_id", controllers.GetOrderItemById())
	router.POST("/orderItems", controllers.CreateOrderItem())
	router.PATCH("/orderItems/:orderItem_id", controllers.UpdateOrderItem())
	router.GET("/orderItems-order/:orderId",controllers.GetOrderItemsByOrder())
}