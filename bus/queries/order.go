package queries

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"packages.hetic.net/gocqrs/models"
)

var OrderQueryChan = make(chan models.Instruction)

func InitRoutine() {
	go func() {
		order := <-OrderQueryChan

		ginContext := order.Data.(*gin.Context)

		GetOrdersQueryHandler(ginContext)
	}()
}

func GetOrdersQueryHandler(c *gin.Context) {
	createdOrders, err := models.GetOrders(c)

	if err != nil {
		fmt.Print(err)
	}

	var instruction models.Instruction

	instruction.Data = createdOrders

	OrderQueryChan <- instruction
}
