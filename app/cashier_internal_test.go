package app

import (
	"testing"

	"github.com/apaliavy/checkout/app/discount"
	"github.com/apaliavy/checkout/app/price"
	"github.com/apaliavy/checkout/app/stock"

	"github.com/stretchr/testify/assert"
)

func TestCashier_Scan(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name               string
		itemsToScan        []string
		expectedCartLen    int
		expectedItemsCount map[string]int
	}{
		{
			name:               "scan single item",
			itemsToScan:        []string{"A"},
			expectedCartLen:    1,
			expectedItemsCount: map[string]int{"A": 1},
		},
		{
			name:               "scan a few unique items",
			itemsToScan:        []string{"A", "B"},
			expectedCartLen:    2,
			expectedItemsCount: map[string]int{"A": 1, "B": 1},
		},
		{
			name:               "scan two equal items",
			itemsToScan:        []string{"A", "A"},
			expectedCartLen:    1,
			expectedItemsCount: map[string]int{"A": 2},
		},
		{
			name:               "scan a few items - equal and unique",
			itemsToScan:        []string{"A", "A", "B", "B", "C", "A", "D"},
			expectedCartLen:    4,
			expectedItemsCount: map[string]int{"A": 3, "B": 2, "C": 1, "D": 1},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cashier := NewCashier(price.NewCalculator(
				stock.NewProductsCollection(),
				discount.NewSpecialOffersCollection(),
			))

			for _, i := range tc.itemsToScan {
				cashier.Scan(i)
			}

			assert.Len(t, cashier.items, tc.expectedCartLen)
			for sku, count := range tc.expectedItemsCount {
				assert.Equal(t, count, cashier.items[sku])
			}
		})
	}
}
