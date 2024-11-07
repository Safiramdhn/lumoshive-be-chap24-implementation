package collections

type Book struct {
	ID       string   `json:"bookId"`
	Name     string   `json:"bookName"`
	Author   string   `json:"bookAuthor"`
	Price    float64  `json:"price"`
	Discount int      `json:"discount"`
	Cover    []byte   `json:"bookCover"`
	File     []byte   `json:"bookFile"`
	Quantity int      `json:"bookQuantity"`
	Category Category `json:"category"`
}
