package handlers

import (
	"encoding/json"
	"golang-beginner-chap24/collections"
	"golang-beginner-chap24/services"
	"golang-beginner-chap24/utils"
	"net/http"
)

type OrderHandler struct {
	OrderService *services.OrderService
}

func NewOrderHandler(orderService *services.OrderService) *OrderHandler {
	return &OrderHandler{OrderService: orderService}
}

func (h *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Invalid request method", nil)
		return
	}

	order := collections.Order{}
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	err := h.OrderService.CreateOrder(order)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, "Order created successfully", nil)
}
