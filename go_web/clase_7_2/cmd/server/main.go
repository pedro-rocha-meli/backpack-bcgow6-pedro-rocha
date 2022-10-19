package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/go_web/clase_7_2/cmd/internal/transactions"
	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/go_web/clase_7_2/cmd/pkg/store"
	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/go_web/clase_7_2/cmd/server/handler"

	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/override/docs"
)


// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("an error occurred while trying to load the environment variables")
	}

	db := store.New(store.FileType, "./transactions.json")
	repository := transactions.NewRepository(db)
	service := transactions.NewService(repository)
	handler := handler.NewHandler(service)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	transactionRoutes := router.Group("/transactions")
	{
		transactionRoutes.GET("/", handler.GetAll())
		transactionRoutes.POST("/", handler.Store())
		transactionRoutes.PUT("/:id", handler.Update())
		transactionRoutes.DELETE("/:id", handler.Delete())
		transactionRoutes.PATCH("/:id", handler.UpdateCode())
	}

	router.Run()
}