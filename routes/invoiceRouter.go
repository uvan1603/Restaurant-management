package routes

import (
	"github.com/gin-gonic/gin"
	"restaurant-management/controllers"
)
func InvoiceRoutes(router *gin.Engine) {
	router.GET("/invoices" , controllers.GetInvoices())
	router.GET("/invoices/:invoice_id", controllers.GetInvoiceById())
	router.POST("/invoices" , controllers.CreateInvoice())
	router.PATCH("/invoices/:invoice_id" , controllers.UpdateInvoice())
}