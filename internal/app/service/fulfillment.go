package service

import (
	"re-partners/internal/domain"
)

type Fulfillment interface {
	Calculate(itemsOrdered int) *domain.Fulfillment
}

type fulfillment struct{}

func New() Fulfillment {
	return &fulfillment{}
}

// Calculate determines the minimum number of items and packs needed to fulfill the order
func (f *fulfillment) Calculate(itemsOrdered int) *domain.Fulfillment {
	result := &domain.Fulfillment{
		Packs: make(map[int]int),
	}

	remainingItems := itemsOrdered

	for _, size := range domain.PackSizes {
		remainder := remainingItems % size

		count := remainingItems / size
		if count > 0 {
			result.Packs[size] = count
		}
		remainingItems = remainder
	}

	if remainingItems > 250 {
		result.Packs[500]++
	} else if remainingItems > 0 {
		result.Packs[250]++
	}

	for i := len(domain.PackSizes) - 1; i > 0; i-- {
		if count, ok := result.Packs[domain.PackSizes[i]]; ok && count > 0 {
			smallerItemsCount := domain.PackSizes[i] * result.Packs[domain.PackSizes[i]]
			packSize := domain.PackSizes[i-1]
			if packSize == smallerItemsCount {
				result.Packs[packSize]++
				delete(result.Packs, domain.PackSizes[i])
			}
		}
	}

	return result
}
