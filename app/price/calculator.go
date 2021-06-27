package price

import (
	"github.com/apaliavy/checkout/app/discount"
	"github.com/apaliavy/checkout/app/stock"
)

type Calculator struct {
	// todo: implement
}

func NewCalculator(products stock.ProductsCollection, offers discount.SpecialOffersCollection) *Calculator {
	return &Calculator{} // todo: implement
}
