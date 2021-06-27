package stock

type SKU string

type Product struct {
	SKU       SKU
	UnitPrice int
}

type ProductsCollection map[SKU]Product

func NewProductsCollection(products ...Product) ProductsCollection {
	return ProductsCollection{} // todo: implement
}

func (p ProductsCollection) GetUnitPrice(sku SKU) (int, error) {
	return 0, nil // todo: implement
}
