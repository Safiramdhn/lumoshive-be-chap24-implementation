package handlers

import (
	"encoding/json"
	"golang-beginner-chap24/collections"
	"golang-beginner-chap24/services"
	"golang-beginner-chap24/utils"
	"net/http"

	"go.uber.org/zap"
)

type OrderHandler struct {
	OrderService *services.OrderService
	Log          *zap.Logger
}

func NewOrderHandler(orderService *services.OrderService, log *zap.Logger) *OrderHandler {
	return &OrderHandler{OrderService: orderService, Log: log}
}

func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.Log.Error("Invalid request method", zap.String("hanlder", "OrderHandler"), zap.String("function", "CreateOrderHandler"))
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Invalid request method", nil)
		return
	}

	order := collections.Order{}
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		h.Log.Error("Invalid request body"+err.Error(), zap.String("hanlder", "OrderHandler"), zap.String("function", "CreateOrderHandler"))
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err := h.OrderService.CreateOrder(order)
	if err != nil {
		h.Log.Error(err.Error(), zap.String("hanlder", "OrderHandler"), zap.String("function", "CreateOrderHandler"))
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, "Order created successfully", nil)
}
