package discount

import "github.com/apaliavy/checkout/app/stock"

// SpecialOffer describes a special offer for some products (by SKU).
// For example: buy three As and they’ll cost you 130 (instead of 150).
type SpecialOffer struct {
	SKU          stock.SKU
	Quantity     int
	SpecialPrice int
}

// SpecialOffersCollection - a set of special offers (SKU is an identifier)
type SpecialOffersCollection map[stock.SKU]SpecialOffer

// NewSpecialOffersCollection returns you collection of special offers
func NewSpecialOffersCollection(offers ...SpecialOffer) SpecialOffersCollection {
	collection := make(SpecialOffersCollection)
	for _, o := range offers {
		collection[o.SKU] = o
	}

	return collection
}

// Apply a special offer to the product.
// Returns a special price if special offer for given SKU was found.
// Second return parameter describes amount of items non-covered by special offer.
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
