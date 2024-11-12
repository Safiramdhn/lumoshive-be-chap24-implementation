package handlers

import (
	"golang-beginner-chap24/services"
	"golang-beginner-chap24/utils"
	"net/http"

	"go.uber.org/zap"
)

type OrderItemHandler struct {
	OrderItemService *services.OrderItemService
	Log              *zap.Logger
}

func NewOrderItemHandler(service *services.OrderItemService, log *zap.Logger) *OrderItemHandler {
	return &OrderItemHandler{
		OrderItemService: service,
		Log:              log,
	}
}

func (h *OrderItemHandler) GetRatingAverageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.Log.Error("Invalid request method", zap.String("hanlder", "OrderItemHandler"), zap.String("function", "GetRatingAverageHandler"))
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Invalid request method", nil)
		return
	}

	ratingAvg, err := h.OrderItemService.GetRatingAverage()
	if err != nil {
		h.Log.Error("Error getting rating average", zap.String("handler", "OrderItemHandler"), zap.String("function", "GetRatingAverageHandler"), zap.Error(err))
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Rating average", ratingAvg)
}
