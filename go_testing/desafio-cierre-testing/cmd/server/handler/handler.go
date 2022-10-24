package handler

import (
	"example.com/internal/products"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	svc products.Service
}

func NewHandler(s products.Service) *Handler {
	return &Handler{
		svc: s,
	}
}

func (h *Handler) GetProducts(ctx *gin.Context) {
	sellerID := ctx.Query("seller_id")
	if sellerID == "" {
		ctx.JSON(400, gin.H{"error": "seller_id query param is required"})
		return
	}
	products, err := h.svc.GetAllBySeller(sellerID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, products)

}
