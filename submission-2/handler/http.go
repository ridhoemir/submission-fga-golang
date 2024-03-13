package handler

import (
	"submission-2/handler/order"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(orderHandler *order.Handler) {
	r := gin.Default()

	r.GET("/orders", orderHandler.GetAllOrders)
	r.POST("/orders", orderHandler.CreateOrder)
	r.PUT("/orders/:id", orderHandler.UpdateOrder)
	r.DELETE("/orders/:id", orderHandler.DeleteOrder)
	r.Run(":8000")
}
