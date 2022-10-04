package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type Transaction struct {
	Id       int     `json:"id"`
	Code     string  `json:"code" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Sender   string  `json:"sender" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

func main() {

	var transactions []Transaction
	var lastID int
	router := gin.Default()

	router.POST("/transactions", func(ctx *gin.Context) {

		var request Transaction

		if ctx.GetHeader("token") != "1234" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "no tiene permisos para realizar la peticion solicitada",
			})

			return
		}

		if err := ctx.ShouldBindJSON(&request); err != nil {

			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("el campo %s es requerido", e.Field())
				errorMessages = append(errorMessages, errorMessage)
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": errorMessages,
			})
			return
		}

		lastID++
		request.Id = lastID
		transactions = append(transactions, request)
		ctx.JSON(200, request)
	})

	router.Run()
}
