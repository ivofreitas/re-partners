package domain

import "re-partners/internal/config"

// PackSizes defines the available pack sizes in descending order for ease of calculation and to minimize the number of packs.
var PackSizes = config.GetEnv().Server.PackSizes

// Order represents the total number of items a customer has ordered.
type Order struct {
	TotalItems int `json:"total_items" validate:"required,gt=0"` // Total number of items ordered by the customer
}

// Fulfillment represents the result of calculating the required packs for an order.
type Fulfillment struct {
	Packs map[int]int `json:"packs"` // Pack size to count mapping
}
