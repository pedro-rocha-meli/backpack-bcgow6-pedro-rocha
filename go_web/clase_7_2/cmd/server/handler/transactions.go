package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/go_web/clase_7_2/cmd/internal/transactions"
)

type Request struct {
	Id       int     `json:"id"`
	Code     string  `json:"code" binding:"required"`
	Currency string  `json:"currency" binding:"required"`
	Amount   float64 `json:"amount" binding:"required"`
	Sender   string  `json:"sender" binding:"required"`
	Date     string  `json:"date" binding:"required"`
}

type Handler struct {
	service transactions.Service
}

func NewHandler(s transactions.Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("token") != "1234" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token invalido",
			})
			return
		}

		transactions, err := h.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, transactions)
	}
}

func (h *Handler) Store() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "1234" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		var request Request
		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		transaction, err := h.service.Store(request.Id, request.Code, request.Currency, request.Amount, request.Sender, request.Date)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, transaction)
	}
}
