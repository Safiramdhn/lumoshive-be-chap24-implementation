package collections

type Customer struct {
	ID      int     `json:"customerId"`
	Name    string  `json:"customerName"`
	Phone   string  `json:"customerPhone"`
	Address Address `json:"shippingAddress"`
}
