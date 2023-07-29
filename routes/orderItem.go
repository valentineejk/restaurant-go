package routes

import (
	"restaurant-go/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/ordersItems", controllers.GetOrderItems())
	incomingRoutes.GET("/ordersItems/:orderItem_id", controllers.GetOrderItem())
	incomingRoutes.GET("/ordersItems-order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/ordersItems", controllers.CreateOrderItem())
	incomingRoutes.PUT("/ordersItems/:orderItem_id", controllers.UpdateOrderItem())
}
