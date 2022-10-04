package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/go_web/clase_7_2/cmd/internal/transactions"
	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/go_web/clase_7_2/cmd/server/handler"
)

func main() {
	repository := transactions.NewRepository()
	service := transactions.NewService(repository)
	handler := handler.NewHandler(service)

	router := gin.Default()

	transactionRoutes := router.Group("/transactions")
	{
		transactionRoutes.GET("/", handler.GetAll())
		transactionRoutes.POST("/", handler.Store())
	}

	router.Run()
}
