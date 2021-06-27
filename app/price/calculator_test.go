package price_test

import (
	"testing"

	"github.com/apaliavy/checkout/app/discount"
	"github.com/apaliavy/checkout/app/price"
	"github.com/apaliavy/checkout/app/stock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCalculator(t *testing.T) {
	t.Parallel()

	calculator := price.NewCalculator(stock.NewProductsCollection(), discount.NewSpecialOffersCollection())
	assert.NotNil(t, calculator)
}

func TestCalculator_CalculateItemsPrice(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		calculator    *price.Calculator
		sku           stock.SKU
		quantity      int
		expectError   bool
		expectedTotal int
	}{
		{
			name: "I have one product in cart, no special offers given",
			calculator: price.NewCalculator(
				stock.NewProductsCollection(stock.Product{SKU: "A", UnitPrice: 50}),
				discount.NewSpecialOffersCollection(),
			),
			sku:           stock.SKU("A"),
			quantity:      1,
			expectedTotal: 50,
		},
		{
			name: "I have a few items in the cart, no items satisfy special offers rules",
			calculator: price.NewCalculator(
				stock.NewProductsCollection(stock.Product{SKU: "A", UnitPrice: 50}),
				discount.NewSpecialOffersCollection(discount.SpecialOffer{SKU: "A", Quantity: 3, SpecialPrice: 130}),
			),
			sku:           stock.SKU("A"),
			quantity:      2,
			expectedTotal: 100,
		},
		{
			name: "I have a few items in the cart, special offer applied for all of them",
			calculator: price.NewCalculator(
				stock.NewProductsCollection(stock.Product{SKU: "A", UnitPrice: 50}),
				discount.NewSpecialOffersCollection(discount.SpecialOffer{SKU: "A", Quantity: 3, SpecialPrice: 130}),
			),
			sku:           stock.SKU("A"),
			quantity:      3,
			expectedTotal: 130,
		},
		{
			name: "I have a few items in the cart, special offer applied for some of them",
			calculator: price.NewCalculator(
				stock.NewProductsCollection(stock.Product{SKU: "A", UnitPrice: 50}),
				discount.NewSpecialOffersCollection(discount.SpecialOffer{SKU: "A", Quantity: 3, SpecialPrice: 130}),
			),
			sku:           stock.SKU("A"),
			quantity:      4,
			expectedTotal: 180,
		},
		{
			name: "I have a few items in the cart, but some items aren't found in the products list",
			calculator: price.NewCalculator(
				stock.NewProductsCollection(stock.Product{SKU: "B", UnitPrice: 50}),
				discount.NewSpecialOffersCollection(discount.SpecialOffer{SKU: "B", Quantity: 3, SpecialPrice: 130}),
			),
			sku:         stock.SKU("A"),
			quantity:    4,
			expectError: true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			total, err := tc.calculator.CalculateItemsPrice(tc.sku, tc.quantity)
			if tc.expectError {
				require.Error(t, err)
				return
			}

			assert.Equal(t, tc.expectedTotal, total)
		})
	}
}
