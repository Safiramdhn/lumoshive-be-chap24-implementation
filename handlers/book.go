package handlers

import (
	"fmt"
	"golang-beginner-chap24/collections"
	"golang-beginner-chap24/services"
	"golang-beginner-chap24/utils"
	"io"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type BookHandler struct {
	BookService services.BookService
}

func NewBookHandler(bookService services.BookService) *BookHandler {
	return &BookHandler{BookService: bookService}
}

func (h *BookHandler) CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.RespondWithJSON(w, http.StatusMethodNotAllowed, "Invalid request method", nil)
		return
	}

	if err := r.ParseForm(); err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request form", nil)
		return
	}

	name := r.Form.Get("bookName")
	category := r.Form.Get("bookCategory")
	categoryId, _ := strconv.Atoi(category)
	author := r.Form.Get("author")
	price := r.Form.Get("price")
	priceFloat, _ := strconv.ParseFloat(price, 64)
	discount := r.Form.Get("discount")
	discountInt, _ := strconv.Atoi(discount)
	quantity := r.Form.Get("quantity")
	quantityInt, _ := strconv.Atoi(quantity)

	coverFile, _, err := r.FormFile("cover")
	if err != nil {
		http.Error(w, "Unable to retrieve cover file", http.StatusBadRequest)
		return
	}
	pdfFile, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Unable to retrieve PDF file", http.StatusBadRequest)
		return
	}
	defer coverFile.Close()
	defer pdfFile.Close()

	coverBytes, err := io.ReadAll(coverFile)
	if err != nil {
		http.Error(w, "Unable to read cover file", http.StatusInternalServerError)
		return
	}
	fileBytes, err := io.ReadAll(pdfFile)
	if err != nil {
		http.Error(w, "Unable to read PDF file", http.StatusInternalServerError)
		return
	}

	bookInput := collections.Book{
		Name: name,
		Category: collections.Category{
			ID: categoryId,
		},
		Author:   author,
		Price:    priceFloat,
		Discount: discountInt,
		Cover:    coverBytes,
		File:     fileBytes,
		Quantity: quantityInt,
	}

	err = h.BookService.CreateBook(bookInput)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusBadRequest, "Invalid request method", nil)
	}
}

func (h *BookHandler) GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	// Get all books from the database
	books, err := h.BookService.GetAllBooks()
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	for i, book := range books {
		// Convert price to a formatted string with thousands separator
		formattedPrice := fmt.Sprintf("Rp%s", formatPrice(book.Price))
		books[i].FormattedPrice = formattedPrice
	}
	err = templates.ExecuteTemplate(w, "book-list-view", map[string]interface{}{
		"Books": books,
	})
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}
}

func (h *BookHandler) EditBookFormHandler(w http.ResponseWriter, r *http.Request) {
	// Get the book ID from the URL parameters using chi's Param function
	bookId := chi.URLParam(r, "id")

	// Retrieve the book and categories from the service
	book, categories, err := h.BookService.GetBookByID(bookId)
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	fmt.Printf("Book: %+v, Categories: %+v", book, categories)

	// Ensure the book has a valid Category field and pass the data to the template
	err = templates.ExecuteTemplate(w, "edit-book-form", map[string]interface{}{
		"Book":       book,
		"Categories": categories,
	})
	if err != nil {
		utils.RespondWithJSON(w, http.StatusInternalServerError, err.Error(), nil)
	}
}

func formatPrice(price float64) string {
	// Convert price to string with commas
	return strconv.FormatFloat(price, 'f', 0, 64)
}
