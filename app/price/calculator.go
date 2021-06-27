package price

import (
	"github.com/apaliavy/checkout/app/discount"
	"github.com/apaliavy/checkout/app/stock"
)

// Calculator is used to calculate total cost of given products (based on SKU and items quantity).
type Calculator struct {
	products      stock.ProductsCollection
	specialOffers discount.SpecialOffersCollection
}

// NewCalculator returns a pointer to Calculator with given products and offers.
func NewCalculator(products stock.ProductsCollection, offers discount.SpecialOffersCollection) *Calculator {
	return &Calculator{
		products:      products,
		specialOffers: offers,
	}
}

// CalculateItemsPrice performs total cost calculation.
// Give it product SKU and items quantity in cart and it gives you total price for selected SKU.
// Returns an error if it doesn't know item with such SKU value.
func (calc *Calculator) CalculateItemsPrice(sku stock.SKU, quantity int) (int, error) {
	specialPrice, quantity := calc.specialOffers.Apply(sku, quantity)
	if quantity == 0 {
		// all items are covered by special offers, there's no need to get a regular price
		return specialPrice, nil
	}

	unitPrice, err := calc.products.GetUnitPrice(sku)
	if err != nil {
		return 0, err
	}

	return specialPrice + (unitPrice * quantity), nil
}
