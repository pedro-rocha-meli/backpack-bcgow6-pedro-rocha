package main

import (
	"github.com/gin-gonic/gin"
)

//Ejercicio 3
type Transaction struct {
	Id       int
	Code     string
	Currency string
	Amount   float64
	Sender   string
	Date     string
}

func main() {

	//Ejercicio 2
	router := gin.Default()

	router.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome user",
		})
	})

	
	//Ejercicio 3
	transactions := []Transaction{
		{
			Id: 1,
			Code: "ARX-12",
			Currency: "BTC",
			Amount: 0.00032,
			Sender: "aximBodok",
			Date: "22-03-21",
		},
	}

	transactionRoutes := router.Group("/transactions")
	{
		transactionRoutes.GET("/GetAll", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": transactions,
			})
		})
	}


	router.Run()
}
