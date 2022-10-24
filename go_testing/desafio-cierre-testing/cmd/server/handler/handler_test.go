package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/internal/domain"
	"example.com/internal/products"
	"example.com/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockRepository mocks.MockRepository) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	service := products.NewService(&mockRepository)
	handler := NewHandler(service)

	router := gin.Default()

	pr := router.Group("/products")
	pr.GET("", handler.GetProducts)

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestGetProducts(t *testing.T) {
	database := []domain.Product{
		{
			ID: "MA1",
			SellerID: "HDS24",
			Description: "Apple",
			Price: 3.23,
		},
	}

	mockRepository := mocks.MockRepository{
		Data: database,
	}

	router := createServer(mockRepository)
	request, responseRecorder := createRequestTest(http.MethodGet, "/products?seller_id=HDS24", "")

	router.ServeHTTP(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
}

func TestGetProductsWithError(t *testing.T) {
	database := []domain.Product{
		{
			ID: "MA1",
			SellerID: "HDS24",
			Description: "Apple",
			Price: 3.23,
		},
	}

	mockRepository := mocks.MockRepository{
		Data: database,
	}

	r := createServer(mockRepository)

	request400, responseRecorder400 := createRequestTest(http.MethodGet, "/products", "")
	request500, responseRecorder500 := createRequestTest(http.MethodGet, "/products?seller_id=not_found", "")

	r.ServeHTTP(responseRecorder400, request400)
	r.ServeHTTP(responseRecorder500, request500)

	assert.Equal(t, http.StatusBadRequest, responseRecorder400.Code)
	assert.Equal(t, http.StatusInternalServerError, responseRecorder500.Code)
}