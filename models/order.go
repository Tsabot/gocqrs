package models

import (
	"github.com/gin-gonic/gin"
)

type Instruction struct {
	Data      interface{}
	Operation string
}

type Order struct {
	CreatedBy string
}

type CreatedOrder struct {
	ID        int
	CreatedBy string
}

// CreateOrder will add a new order to the DB
func CreateOrder(createdBy interface{}) (CreatedOrder, error) {
	var createdOrder CreatedOrder

	createdBy = createdBy.(string)

	createOrderSQL := `
	INSERT INTO orders (created_by)
	VALUES ($1) RETURNING *;`

	orderRow := db.QueryRow(createOrderSQL, createdBy)
	err := orderRow.Scan(&createdOrder.ID, &createdOrder.CreatedBy)

	if err != nil {
		print(err)
		return createdOrder, err
	}

	return createdOrder, nil
}

// GetOrders will return an order from the DB
func GetOrders(c *gin.Context) ([]CreatedOrder, error) {
	var orders []CreatedOrder

	getOrdersSQL := `
	SELECT * FROM orders;`

	ordersRow, queryErr := db.QueryContext(c, getOrdersSQL)

	if queryErr != nil {
		return orders, queryErr
	}

	for ordersRow.Next() {
		var order CreatedOrder
		if orderErr := ordersRow.Scan(&order.ID, &order.CreatedBy); orderErr != nil {
			return orders, orderErr
		}
		orders = append(orders, order)
	}

	return orders, nil
}
