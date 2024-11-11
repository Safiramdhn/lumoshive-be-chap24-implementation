package collections

type Address struct {
	ID         int    `json:"addressId"`
	CustomerID int    `json:"customerId"`
	Street     string `json:"street"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	Country    string `json:"country"`
}
