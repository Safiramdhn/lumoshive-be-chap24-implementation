package handlers

import (
	"fmt"
	"golang-beginner-chap24/collections"
	"golang-beginner-chap24/services"
	"golang-beginner-chap24/utils"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type PaymentMethodHandler struct {
	PaymentMethodService *services.PaymentMethodService
}

func NewPaymentMethodHandler(paymentMethodService *services.PaymentMethodService) *PaymentMethodHandler {
	return &PaymentMethodHandler{PaymentMethodService: paymentMethodService}
}

func (h *PaymentMethodHandler) CreatePaymentMethodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Invalid request method", nil)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Unable to parse form", err.Error())
		return
	}

	paymentMethodName := r.FormValue("name")
	if paymentMethodName == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Item name is required", nil)
		return
	}

	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Unable to get file from form", err.Error())
		return
	}
	defer file.Close()

	// Define upload path and ensure directory exists
	uploadPath := "./uploads"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to create upload directory", err.Error())
		return
	}

	// Create file on server
	filePath := filepath.Join(uploadPath, fileHeader.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Unable to save file", err.Error())
		return
	}
	defer out.Close()

	// Copy uploaded file content to destination file
	if _, err := io.Copy(out, file); err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to copy file content", err.Error())
		return
	}

	// Construct the photo URL
	photoURL := fmt.Sprintf("http://%s/uploads/%s", r.Host, fileHeader.Filename)

	// Save the payment method to the database
	paymentMethod := collections.PaymentMethod{
		Name:  paymentMethodName,
		Photo: photoURL,
	}

	if err := h.PaymentMethodService.CreatePaymentMethod(paymentMethod); err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to save payment method", err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusCreated, "Payment method created successfully", paymentMethod)
}

func (h *PaymentMethodHandler) GetAllPaymentMethodsHandler(w http.ResponseWriter, r *http.Request) {
	// Get all payment methods from the database
	paymentMethods, err := h.PaymentMethodService.GetAllPaymentMethods()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, "Payment methods retrieved successfully", paymentMethods)
}

func (h *PaymentMethodHandler) GetPaymentMethodByIdHandler(w http.ResponseWriter, r *http.Request) {
	// Get the payment method ID from the URL parameter
	idParam := chi.URLParam(r, "id")
	paymentMethodId, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid payment method ID", nil)
		return
	}

	// Get the payment method from the database
	paymentMethod, err := h.PaymentMethodService.GetPaymentMethodById(paymentMethodId)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusNotFound, "Payment method not found", nil)
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Payment method successfully retrieved", paymentMethod)
}

func (h *PaymentMethodHandler) UpdatePaymentMethodHandler(w http.ResponseWriter, r *http.Request) {
	// Get the payment method ID from the URL parameter
	idParam := chi.URLParam(r, "id")
	paymentMethodId, err := strconv.Atoi(idParam)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid payment method ID", nil)
		return
	}

	if r.Method != http.MethodPut {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Invalid request method", nil)
		return
	}

	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Unable to parse form", err.Error())
		return
	}

	paymentMethodName := r.FormValue("name")
	if paymentMethodName == "" {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Item name is required", nil)
		return
	}

	file, fileHeader, err := r.FormFile("photo")
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Unable to get file from form", err.Error())
		return
	}
	defer file.Close()

	// Define upload path and ensure directory exists
	uploadPath := "./uploads"
	if err := os.MkdirAll(uploadPath, os.ModePerm); err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to create upload directory", err.Error())
		return
	}

	// Create file on server
	filePath := filepath.Join(uploadPath, fileHeader.Filename)
	out, err := os.Create(filePath)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Unable to save file", err.Error())
		return
	}
	defer out.Close()

	// Copy uploaded file content to destination file
	if _, err := io.Copy(out, file); err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to copy file content", err.Error())
		return
	}

	// Construct the photo URL
	photoURL := fmt.Sprintf("http://%s/uploads/%s", r.Host, fileHeader.Filename)

	// Save the payment method to the database
	paymentMethod := collections.PaymentMethod{
		Name:  paymentMethodName,
		Photo: photoURL,
	}

	if err := h.PaymentMethodService.UpdatePaymentMethod(paymentMethodId, paymentMethod); err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to update payment method", err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Payment method updated successfully", paymentMethod)
}

func (h *PaymentMethodHandler) DeletePaymentMethodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Invalid request method", nil)
		return
	}

	// Get the payment method ID from the URL parameters
	paymentMethodId, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid payment method ID", nil)
		return
	}

	// Delete the payment method from the database
	if err := h.PaymentMethodService.DeletePaymentMethod(paymentMethodId); err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, "Failed to delete payment method", err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, "Payment method deleted successfully", nil)
}
