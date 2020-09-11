package bus

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gocqrs/bus/commands"
	"packages.hetic.net/gocqrs/bus/queries"
	"packages.hetic.net/gocqrs/models"
)

func OrderBusCommandHandler(order models.Instruction) models.CreatedOrder {
	var createdOrder models.CreatedOrder

	switch order.Operation {
	case "Create":
		fmt.Println("Créer une commande")
		createdOrder, _ = commands.CreateOrderCommand(order)
		return createdOrder
	default:
		fmt.Println("Mirmph")
	}

	return createdOrder
}

func OrderBusQueryHandler(order models.Instruction, c *gin.Context) []models.CreatedOrder {
	var createdOrders []models.CreatedOrder

	switch order.Operation {
	case "Selects":
		fmt.Println("Récupérer les commandes")

		order.Data = c

		queries.OrderQueryChan <- order

		createdOrders := <-queries.OrderQueryChan

		return createdOrders.Data.([]models.CreatedOrder)
	default:
		fmt.Println("Mirmph")
	}

	return createdOrders
}
