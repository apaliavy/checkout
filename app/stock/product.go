package stock

import "fmt"

type SKU string

type Product struct {
	SKU       SKU
	UnitPrice int
}

type ProductsCollection map[SKU]Product

func NewProductsCollection(products ...Product) ProductsCollection {
	collection := make(ProductsCollection)
	for _, p := range products {
		collection[p.SKU] = p
	}

	return collection
}

func (p ProductsCollection) GetUnitPrice(sku SKU) (int, error) {
	product, ok := p[sku]
	if !ok {
		return 0, fmt.Errorf("failed to get product price - unknown product(sku:%s)", sku)
	}
	return product.UnitPrice, nil
}
