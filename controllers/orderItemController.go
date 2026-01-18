package controllers

import "github.com/gin-gonic/gin"

func GetOrderItems() gin.HandlerFunc {
	return func( c *gin.Context){

	}

}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func( c *gin.Context){
		
	}
}

func ItemsByOrder(id string) (OrderItems []primitive.M, err error) {
	
}

func GetOrderItemById() gin.HandlerFunc {
	return func( c *gin.Context) {

	}

}



func CreateOrderItem() gin.HandlerFunc {
	return func( c *gin.Context){

	}

}

func UpdateOrderItem() gin.HandlerFunc {
	return func( c *gin.Context){
		
	}
}

 