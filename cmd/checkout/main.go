package main

import "github.com/sirupsen/logrus"

//
// | SKU           | Unit Price    | Special Price |
// | ------------- | ------------- | ------------- |
// |       A       |       50      |   3  for 130  |
// |       B       |       30      |   2  for 45   |
// |       C       |       20      |               |
// |       D       |       15      |               |
//

// Entities in the system:
// Product - has SKU assigned and price (// todo: SKU - separate type?)

// SpecialOffer - has SKU, items quantity and special price for this exact quantity

// Pricing - should be independent from checkout, so we'll have something like a pricing.Calculator

// Checkout - Scans items (SKU) and returns a total price (// todo: better naming, cashier?)
// Checkout must has access to our calculator, ie ch.calculator.CalculateTotalPrice()

func main() {
	logger := logrus.New()
	logger.Info("check app is up and running")

	// some pseudocode first..

	// products := cart.Products(
	//   cart.Product("A", 50),
	//   cart.Product("B", 30),
	//   cart.Product("C", 20),
	//   cart.Product("D", 15),
	//)

	// offers := pricing.SpecialOffers(
	//   pricing.SpecialOffer("A", 3, 130),
	//   pricing.SpecialOffer("B", 2, 45)
	//)

	// calculator := pricing.Calculator(products, offers)

	// cashier := app.Checkout(calculator)
	// cashier.Scan("A")
	// cashier.Scan("B")
	// cashier.Scan("A")
	// cashier.Scan("C")
	// cashier.Scan("D")
	// cashier.Scan("A")
	// todo: what if we have 4 "A" items? (3 for special price, 1 for regular one)
	// (calculator responsible)?

	// total, err := cashier.GetTotalPrice()
	// if err != nil {
	//   (some error handler, probably logger.WithError(err).Fatal("... message ... "))
	// }

	// display results
	// logger.Info("the result is: %d", total)
}
