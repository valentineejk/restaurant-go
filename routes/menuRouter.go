package routes

import "github.com/gin-gonic/gin"

func MenuRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/Menus", controller.GetInvoices())
	incomingRoutes.GET("/invoices/:invoicse_id", controller.GetInvoice())
	incomingRoutes.POST("/invoices", controller.CreateInvoice())
	incomingRoutes.PUT("/invoices/:invoices_id", controller.UpdateInvoice())
}
