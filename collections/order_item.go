package collections

type OrderItem struct {
	ID       int     `json:"orderItemId"`
	OrderID  int     `json:"orderId"`
	BookID   string  `json:"bookId"`
	BookName string  `json:"bookName"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
	Subtotal float64 `json:"subtotal"`
	Discount float64 `json:"discount"`
}
