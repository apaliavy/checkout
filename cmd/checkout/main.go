package main

import (
	"github.com/apaliavy/checkout/app"
	"github.com/apaliavy/checkout/app/discount"
	"github.com/apaliavy/checkout/app/price"
	"github.com/apaliavy/checkout/app/stock"

	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()

	products := stock.NewProductsCollection(
		stock.Product{SKU: "A", UnitPrice: 50},
		stock.Product{SKU: "B", UnitPrice: 30},
		stock.Product{SKU: "C", UnitPrice: 20},
		stock.Product{SKU: "D", UnitPrice: 15},
	)

	specialOffers := discount.NewSpecialOffersCollection(
		discount.SpecialOffer{SKU: "A", Quantity: 3, SpecialPrice: 130},
		discount.SpecialOffer{SKU: "B", Quantity: 2, SpecialPrice: 45},
	)

	priceCalculator := price.NewCalculator(products, specialOffers)

	cashier := app.NewCashier(priceCalculator)
	cashier.Scan("A")
	cashier.Scan("B")
	cashier.Scan("A")
	cashier.Scan("C")
	cashier.Scan("D")
	cashier.Scan("A")

	total, err := cashier.GetTotalPrice()
	if err != nil {
		logger.WithError(err).Fatal("failed to calculate total price")
	}

	logger.Infof("the total price is: %d", total)
}
