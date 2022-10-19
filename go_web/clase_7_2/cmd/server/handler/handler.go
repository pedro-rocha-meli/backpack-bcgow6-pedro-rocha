package handler

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/go_web/clase_7_2/cmd/internal/transactions"
	"github.com/pedro-rocha-meli/backpack-bcgow6-pedro-rocha/go_web/clase_7_2/cmd/pkg/web"
)

type Request struct {
	Id       int     `json:"id"`
	Code     string  `json:"code" validate:"required"`
	Currency string  `json:"currency" validate:"required"`
	Amount   float64 `json:"amount" validate:"required"`
	Sender   string  `json:"sender" validate:"required"`
	Date     string  `json:"date" validate:"required"`
}

type Handler struct {
	service transactions.Service
}

func NewHandler(s transactions.Service) *Handler {
	return &Handler{
		service: s,
	}
}

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /products [get]
func (h *Handler) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("token") != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "invalid token"))
			return
		}

		transactions, err := h.service.GetAll()

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, transactions, ""))
	}
}

// StoreProducts godoc
// @Summary Store products
// @Tags Products
// @Description store products
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param product body request true "Product to store"
// @Success 200 {object} web.Response
// @Router /products [post]
func (h *Handler) Store() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("token") != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "invalid token"))
			return
		}

		var request Request

		if err := ctx.Bind(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}

		if err := request.validate(); err != nil {
			valError := err.(validator.ValidationErrors)
			errors := []string{}
			for _, e := range valError {
				errors = append(errors, "the field "+e.Field()+" is required")
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"error": errors,
			})
			return
		}
		
		transaction, err := h.service.Store(request.Code, request.Currency, request.Amount, request.Sender, request.Date)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusInternalServerError, nil, err.Error()))
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, transaction, ""))
	}
}

func (h *Handler) Update() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if ctx.Request.Header.Get("token") != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "invalid token"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusBadRequest, nil, "invalid id"))
			return
		}

		var request Request

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}

		if err := request.validate(); err != nil {
			valError := err.(validator.ValidationErrors)
			errors := []string{}
			for _, e := range valError {
				errors = append(errors, "the field "+e.Field()+" is required")
			}
			ctx.JSON(http.StatusBadRequest, gin.H{
				"code": http.StatusBadRequest,
				"error": errors,
			})
			return
		}

		transaction, err := h.service.Update(id, request.Code, request.Currency, request.Amount, request.Sender, request.Date)

		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusBadRequest, nil, "invalid id"))
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, transaction, ""))
	}
}

func (h *Handler) Delete() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		if ctx.Request.Header.Get("token") != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "invalid token"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusUnauthorized, nil, "invalid id"))
			return
		}

		if err := h.service.Delete(id); err != nil {
			ctx.JSON(http.StatusNotFound, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}

		msj := fmt.Sprintf("the transaction with id: %d was deleted successfully", id)

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, msj, ""))

	}
}
func (h *Handler) UpdateCode() gin.HandlerFunc {

	return func(ctx *gin.Context) {

		if ctx.Request.Header.Get("token") != os.Getenv("TOKEN") {
			ctx.JSON(http.StatusUnauthorized, web.NewResponse(http.StatusUnauthorized, nil, "invalid token"))
			return
		}

		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusUnauthorized, nil, "invalid id"))
			return
		}

		var request Request

		if err := ctx.ShouldBindJSON(&request); err != nil {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusUnauthorized, nil, err.Error()))
			return
		}

		if request.Code == "" {
			ctx.JSON(http.StatusBadRequest, web.NewResponse(http.StatusUnauthorized, nil, "the code is required"))
			return
		}

		transaction, err := h.service.UpdateCode(id, request.Code)

		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, web.NewResponse(http.StatusOK, transaction, ""))
	}
}

func (r *Request) validate() error {
    validate := validator.New()
    return validate.Struct(r)
}