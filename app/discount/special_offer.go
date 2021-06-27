package discount

import "github.com/apaliavy/checkout/app/stock"

type SpecialOffer struct {
	SKU          stock.SKU
	Quantity     int
	SpecialPrice int
}

type SpecialOffersCollection map[stock.SKU]SpecialOffer

func NewSpecialOffersCollection(offers ...SpecialOffer) SpecialOffersCollection {
	return SpecialOffersCollection{} // todo: implement
}

func (so SpecialOffersCollection) Apply(sku stock.SKU, quantity int) (int, int) {
	return 0, 0 // todo: implement
}
