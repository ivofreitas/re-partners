package domain

var PackSizes = [5]int{5000, 2000, 1000, 500, 250}

type Order struct {
	TotalItems int `json:"total_items"` // Total number of items ordered by the customer
}

type Fulfillment struct {
	Packs map[int]int `json:"packs"` // Pack size to count
}
