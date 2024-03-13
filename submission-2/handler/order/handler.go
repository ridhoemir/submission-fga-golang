package order

import (
	"net/http"
	"strconv"
	orderDto "submission-2/dto/order"
	"submission-2/service/order"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *order.Service
}

func NewHandler(svc *order.Service) *Handler {
	return &Handler{
		service: svc,
	}
}

func (h *Handler) CreateOrder(ctx *gin.Context) {
	var req orderDto.OrderCreateRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	data := req.ToOrder()
	res, err := h.service.CreateOrder(data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) GetAllOrders(ctx *gin.Context) {
	res, err := h.service.GetAllOrders()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Internal server error",
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) UpdateOrder(ctx *gin.Context) {
	var req orderDto.OrderCreateRequest
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID invalid",
		})
		return
	}

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	data := req.ToOrder()
	res, err := h.service.UpdateOrder(data, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

func (h *Handler) DeleteOrder(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "ID invalid",
		})
		return
	}

	message, err := h.service.DeleteOrder(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": message,
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}
