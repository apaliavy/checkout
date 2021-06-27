package stock

import "fmt"

// SKU - Stock Keeping Unit. In terms of the application - Product identifier (single char like "A", "B", "C", etc).
type SKU string

// Product - supermarket product, identified by SKU and it costs come price (per unit).
type Product struct {
	SKU       SKU
	UnitPrice int
}

// ProductsCollection - Products map (SKU is an identifier)
type ProductsCollection map[SKU]Product

// NewProductsCollection gives you collection of Products.
func NewProductsCollection(products ...Product) ProductsCollection {
	collection := make(ProductsCollection)
	for _, p := range products {
		collection[p.SKU] = p
	}

	return collection
}

// GetUnitPrice perform product lookup (by SKU) and gives you item price (happy path).
// If item isn't found it returns an error.
func (p ProductsCollection) GetUnitPrice(sku SKU) (int, error) {
	product, ok := p[sku]
	if !ok {
		return 0, fmt.Errorf("failed to get product price - unknown product(sku:%s)", sku)
	}
	return product.UnitPrice, nil
}
