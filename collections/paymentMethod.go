package collections

type PaymentMethod struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
	Photo    string `json:"photo"`
}
