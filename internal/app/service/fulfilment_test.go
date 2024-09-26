package service

import (
	"re-partners/internal/domain"
	"testing"
)

func TestFulfillment_Calculate(t *testing.T) {
	// Create a new instance of the fulfillment service
	service := New()

	tests := []struct {
		name         string
		itemsOrdered int
		expected     *domain.Fulfillment
	}{
		{
			name:         "Single item should use the smallest pack (250)",
			itemsOrdered: 1,
			expected: &domain.Fulfillment{
				Packs: map[int]int{
					250: 1,
				},
			},
		},
		{
			name:         "Order of 240 should use one 250 pack, since it's the closest match",
			itemsOrdered: 240,
			expected: &domain.Fulfillment{
				Packs: map[int]int{
					250: 1,
				},
			},
		},
		{
			name:         "Exact match for a pack size of 250",
			itemsOrdered: 250,
			expected: &domain.Fulfillment{
				Packs: map[int]int{
					250: 1,
				},
			},
		},
		{
			name:         "Order slightly above 250 should use the next larger pack (500)",
			itemsOrdered: 251,
			expected: &domain.Fulfillment{
				Packs: map[int]int{
					500: 1,
				},
			},
		},
		{
			name:         "Order with 499 should use the next larger pack (500)",
			itemsOrdered: 499,
			expected: &domain.Fulfillment{
				Packs: map[int]int{
					500: 1,
				},
			},
		},
		{
			name:         "Order slightly above 500 should use next small pack (250)",
			itemsOrdered: 501,
			expected: &domain.Fulfillment{
				Packs: map[int]int{
					250: 1,
					500: 1,
				},
			},
		},
		{
			name:         "Order slightly above 1000 should use one 1000 pack and one 250 pack",
			itemsOrdered: 1001,
			expected: &domain.Fulfillment{
				Packs: map[int]int{
					250:  1,
					1000: 1,
				},
			},
		},
		{
			name:         "Large order requiring multiple large packs and a small pack for leftover (12001 items)",
			itemsOrdered: 12001,
			expected: &domain.Fulfillment{
				Packs: map[int]int{
					5000: 2,
					2000: 1,
					250:  1,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := service.Calculate(test.itemsOrdered)

			// Check individual pack counts
			for size, count := range test.expected.Packs {
				if result.Packs[size] != count {
					t.Errorf("For %d items, expected %d packs of size %d but got %d", test.itemsOrdered, count, size, result.Packs[size])
				}
			}
		})
	}
}
