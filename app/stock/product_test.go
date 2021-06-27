package stock_test

import (
	"testing"

	"github.com/apaliavy/checkout/app/stock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewProductsCollection(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name                  string
		products              []stock.Product
		expectedCollectionLen int
		customAssertFn        func(collection stock.ProductsCollection)
	}{
		{
			name:                  "check I can create an empty collection",
			products:              []stock.Product{},
			expectedCollectionLen: 0,
		},
		{
			name: "check I can create collection with single element",
			products: []stock.Product{{
				SKU:       "A",
				UnitPrice: 50,
			}},
			expectedCollectionLen: 1,
		},
		{
			name: "check I can create collection with multiple elements",
			products: []stock.Product{{
				SKU:       "A",
				UnitPrice: 50,
			}, {
				SKU:       "B",
				UnitPrice: 35,
			}},
			expectedCollectionLen: 2,
		},
		{
			name: "check there are no duplicated elements in collection",
			products: []stock.Product{{
				SKU:       "A",
				UnitPrice: 50,
			}, {
				SKU:       "A",
				UnitPrice: 35,
			}},
			expectedCollectionLen: 1,
			customAssertFn: func(collection stock.ProductsCollection) {
				assert.Equal(t, stock.SKU("A"), collection["A"].SKU)
				assert.Equal(t, 35, collection["A"].UnitPrice)
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			collection := stock.NewProductsCollection(tc.products...)

			assert.Len(t, collection, tc.expectedCollectionLen)
			if tc.customAssertFn != nil {
				tc.customAssertFn(collection)
			}
		})
	}
}

func TestProductsCollection_GetUnitPrice(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		collection    stock.ProductsCollection
		sku           stock.SKU
		expectError   bool
		expectedPrice int
	}{
		{
			name:        "try to get a unit price from an empty products collection",
			collection:  stock.NewProductsCollection(),
			sku:         stock.SKU("A"),
			expectError: true,
		},
		{
			name:        "try to get a unit price by SKU which doesn't exist in the collection",
			collection:  stock.NewProductsCollection(stock.Product{SKU: "B", UnitPrice: 20}),
			sku:         stock.SKU("A"),
			expectError: true,
		},
		{
			name:          "get a unit price - happy path",
			collection:    stock.NewProductsCollection(stock.Product{SKU: "B", UnitPrice: 20}),
			sku:           stock.SKU("B"),
			expectError:   false,
			expectedPrice: 20,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			price, err := tc.collection.GetUnitPrice(tc.sku)
			if tc.expectError {
				require.Error(t, err)
				return
			}

			assert.Equal(t, tc.expectedPrice, price)
		})
	}
}
