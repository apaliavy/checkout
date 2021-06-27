package price

import (
	"github.com/apaliavy/checkout/app/discount"
	"github.com/apaliavy/checkout/app/stock"
)

type Calculator struct {
	products      stock.ProductsCollection
	specialOffers discount.SpecialOffersCollection
}

func NewCalculator(products stock.ProductsCollection, offers discount.SpecialOffersCollection) *Calculator {
	return &Calculator{
		products:      products,
		specialOffers: offers,
	}
}

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
