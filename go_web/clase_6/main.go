package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Ejercicio 3
type Transaction struct {
	Id       int `json:"id"`
	Code     string `json:"code"`
	Currency string `json:"currency"`
	Amount   float64 `json:"amount"`
	Sender   string `json:"sender"`
	Date     string `json:"date"`
}

func main() {

	//Ejercicio 2 TM
	router := gin.Default()

	router.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Welcome user",
		})
	})

	//Ejercicio 3 TM
	data, err := os.ReadFile("./Transactions.json")

	if err != nil {
		panic(err)
	}

	var transactions []Transaction

	json.Unmarshal(data, &transactions)

	transactionRoutes := router.Group("/transactions")
	{
		transactionRoutes.GET("/GetAll", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": transactions,
			})
		})

		
		//Ejercicio 2 TT
		transactionRoutes.GET(":id", func(ctx *gin.Context) {
			id, err := strconv.Atoi(ctx.Param("id"))

			if err != nil {
				panic(err)
			}

			var matchedTransactions []Transaction

			for i := range transactions {
				if transactions[i].Id == id {
					matchedTransactions = append(matchedTransactions, transactions[i])
				}
			}

			if matchedTransactions != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"content": matchedTransactions,
				})
			} else {
				ctx.JSON(http.StatusNotFound, gin.H{
					"message": "Transaction not found",
				})
			}
		})
	}

	router.Run()
}
