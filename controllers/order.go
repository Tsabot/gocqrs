package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gocqrs/bus"
	"packages.hetic.net/gocqrs/models"
)

// CreateOrder handle request to call order handler
func CreateOrder(c *gin.Context) {
	createdBy := c.PostForm("createdBy")

	var order models.Instruction

	order.Data = createdBy

	order.Operation = "Create"

	orderCreated := bus.OrderBusCommandHandler(order)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Created order successfully",
		"content": orderCreated,
	})
}

// GetOrders handle request to call order handler
func GetOrders(c *gin.Context) {
	var order models.Instruction

	order.Operation = "Selects"

	ordersCreated := bus.OrderBusQueryHandler(order, c)

	c.JSON(http.StatusCreated, gin.H{
		"success": true,
		"message": "Found orders successfully",
		"content": ordersCreated,
	})
}
