package main

import (
	"github.com/apaliavy/checkout/app"
	"github.com/apaliavy/checkout/app/discount"
	"github.com/apaliavy/checkout/app/price"
	"github.com/apaliavy/checkout/app/stock"

	"github.com/sirupsen/logrus"
)

// Entities in the system:
// Product - has SKU assigned and price (// todo: SKU - separate type?)

// SpecialOffer - has SKU, items quantity and special price for this exact quantity

// Pricing - should be independent from checkout, so we'll have something like a price.Calculator

// Checkout - Scans items (SKU) and returns a total price (// todo: better naming, cashier?)
// Checkout must has access to our calculator, ie ch.calculator.CalculateTotalPrice()

func main() {
	logger := logrus.New()
	logger.Info("check app is up and running")

	// some pseudocode first..

	// must be a collection + simple iteration through the list
	products := stock.NewProductsCollection(
		stock.Product{SKU: "A", UnitPrice: 50},
		stock.Product{SKU: "B", UnitPrice: 30},
		stock.Product{SKU: "C", UnitPrice: 20},
		stock.Product{SKU: "D", UnitPrice: 15},
	)
	// it should be possible to get a UnitPrice from collection by SKU (products.GetPriceBySKU(sku))
	// we'll use it to calculate a "regular" price

	// the same here:
	specialOffers := discount.NewSpecialOffersCollection(
		discount.SpecialOffer{SKU: stock.SKU("A"), Quantity: 3, SpecialPrice: 130},
		discount.SpecialOffer{SKU: stock.SKU("B"), Quantity: 2, SpecialPrice: 45},
	)

	// to avoid problems with 4 "A" items? (3 for special price, 1 for regular one) case, it should be possible to "Apply" offer to an item
	// if we have some products left - apply regular price, ie:
	// price, quantityLeft := offer.Apply(sku, quantity)

	priceCalculator := price.NewCalculator(products, specialOffers)

	cashier := app.NewCashier(priceCalculator)
	cashier.Scan("A")
	cashier.Scan("B")
	cashier.Scan("A")
	cashier.Scan("C")
	cashier.Scan("D")
	cashier.Scan("A")
	// todo: calculate amount of each element:  A: 4, B: 1, C: 1, D: 1

	// total, err := cashier.GetTotalPrice()
	// if err != nil {
	//   (some error handler, probably logger.WithError(err).Fatal("... message ... "))
	// }

	// display results
	// logger.Info("the result is: %d", total)
}
