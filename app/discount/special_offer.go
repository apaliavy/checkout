package discount

import "github.com/apaliavy/checkout/app/stock"

type SpecialOffer struct {
	SKU          stock.SKU
	Quantity     int
	SpecialPrice int
}

type SpecialOffersCollection map[stock.SKU]SpecialOffer

func NewSpecialOffersCollection(offers ...SpecialOffer) SpecialOffersCollection {
	collection := make(SpecialOffersCollection)
	for _, o := range offers {
		collection[o.SKU] = o
	}

	return collection
}

// nolint: gocritic
func (so SpecialOffersCollection) Apply(sku stock.SKU, quantity int) (int, int) {
	offer, ok := so[sku]
	if !ok {
		return 0, quantity
	}

	price := 0
	for ; quantity >= offer.Quantity; quantity -= offer.Quantity {
		price += offer.SpecialPrice
	}

	return price, quantity
}
