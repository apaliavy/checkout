package discount_test

import (
	"testing"

	"github.com/apaliavy/checkout/app/discount"
	"github.com/apaliavy/checkout/app/stock"

	"github.com/stretchr/testify/assert"
)

func TestNewSpecialOffersCollection(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name                  string
		specialOffers         []discount.SpecialOffer
		expectedCollectionLen int
		customAssertFn        func(collection discount.SpecialOffersCollection)
	}{
		{
			name:                  "check I can create an empty collection",
			specialOffers:         []discount.SpecialOffer{},
			expectedCollectionLen: 0,
		},
		{
			name: "check I can create collection with single element",
			specialOffers: []discount.SpecialOffer{{
				SKU:          "A",
				Quantity:     1,
				SpecialPrice: 50,
			}},
			expectedCollectionLen: 1,
		},
		{
			name: "check I can create collection with multiple elements",
			specialOffers: []discount.SpecialOffer{{
				SKU:          "A",
				Quantity:     3,
				SpecialPrice: 50,
			}, {
				SKU:          "B",
				Quantity:     2,
				SpecialPrice: 35,
			}},
			expectedCollectionLen: 2,
		},
		{
			name: "check there are no duplicated elements in collection",
			specialOffers: []discount.SpecialOffer{{
				SKU:          "A",
				Quantity:     3,
				SpecialPrice: 50,
			}, {
				SKU:          "A",
				Quantity:     2,
				SpecialPrice: 35,
			}},
			expectedCollectionLen: 1,
			customAssertFn: func(collection discount.SpecialOffersCollection) {
				assert.Equal(t, stock.SKU("A"), collection["A"].SKU)
				assert.Equal(t, 2, collection["A"].Quantity)
				assert.Equal(t, 35, collection["A"].SpecialPrice)
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			collection := discount.NewSpecialOffersCollection(tc.specialOffers...)

			assert.Len(t, collection, tc.expectedCollectionLen)
			if tc.customAssertFn != nil {
				tc.customAssertFn(collection)
			}
		})
	}
}

func TestSpecialOffersCollection_Apply(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                 string
		offers               discount.SpecialOffersCollection
		SKU                  stock.SKU
		quantity             int
		expectedTotal        int
		expectedQuantityLeft int
	}{
		{
			name:                 "special offers is empty",
			offers:               discount.NewSpecialOffersCollection(),
			SKU:                  stock.SKU("A"),
			quantity:             3,
			expectedQuantityLeft: 3,
			expectedTotal:        0,
		},
		{
			name: "special offers not found for specified SKU",
			offers: discount.NewSpecialOffersCollection(
				discount.SpecialOffer{SKU: "A", Quantity: 2, SpecialPrice: 130},
			),
			SKU:                  stock.SKU("B"),
			quantity:             3,
			expectedQuantityLeft: 3,
			expectedTotal:        0,
		},
		{
			name: "special offer found but items quantity is not enough",
			offers: discount.NewSpecialOffersCollection(
				discount.SpecialOffer{SKU: "A", Quantity: 2, SpecialPrice: 130},
			),
			SKU:                  stock.SKU("A"),
			quantity:             1,
			expectedQuantityLeft: 1,
			expectedTotal:        0,
		},
		{
			name: "special offer found, quantities are equal",
			offers: discount.NewSpecialOffersCollection(
				discount.SpecialOffer{SKU: "A", Quantity: 2, SpecialPrice: 130},
			),
			SKU:                  stock.SKU("A"),
			quantity:             2,
			expectedQuantityLeft: 0,
			expectedTotal:        130,
		},
		{
			name: "special offer found, quantity is 2x greater",
			offers: discount.NewSpecialOffersCollection(
				discount.SpecialOffer{SKU: "A", Quantity: 2, SpecialPrice: 130},
			),
			SKU:                  stock.SKU("A"),
			quantity:             4,
			expectedQuantityLeft: 0,
			expectedTotal:        260,
		},
		{
			name: "special offer found, items count in cart is N+1 (1 item left)",
			offers: discount.NewSpecialOffersCollection(
				discount.SpecialOffer{SKU: "A", Quantity: 2, SpecialPrice: 130},
			),
			SKU:                  stock.SKU("A"),
			quantity:             3,
			expectedQuantityLeft: 1,
			expectedTotal:        130,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			total, qLeft := tc.offers.Apply(tc.SKU, tc.quantity)

			assert.Equal(t, tc.expectedTotal, total)
			assert.Equal(t, tc.expectedQuantityLeft, qLeft)
		})
	}
}
