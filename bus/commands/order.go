package commands

import (
	"fmt"

	"packages.hetic.net/gocqrs/models"
)

func CreateOrderCommand(order models.Instruction) (models.CreatedOrder, error) {
	createdOrder, err := models.CreateOrder(order.Data)

	if err != nil {
		fmt.Print(err)
	}

	return createdOrder, err
}
