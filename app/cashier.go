package app

import (
	"github.com/apaliavy/checkout/app/stock"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// nolint: lll
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 --fake-name CalculatorMock -o ../testing/mocks/calculator/calculator.go . Calculator

// Calculator describes prices calculation engine.
// Give it product SKU and items quantity in cart and it gives you total price for selected SKU.
// Calculator returns an error if it doesn't know item with such SKU value.
type Calculator interface {
	CalculateItemsPrice(sku stock.SKU, quantity int) (int, error)
}

// Cashier implements Checkout interface.
// Calculator is used internally to calculate total price of all Cashier scanned items.
type Cashier struct {
	calculator Calculator
	items      map[string]int
	logger     logrus.FieldLogger
}

// NewCashier gives you a pointer to the Cashier instance with provided Calculator.
func NewCashier(c Calculator) *Cashier {
	return &Cashier{
		calculator: c,
		items:      make(map[string]int),
		logger:     logrus.New(),
	}
}

// Scan an item and put it into cart.
// Each time when you scan item with existing sku it increases items count.
func (c *Cashier) Scan(sku string) {
	c.logger.Infof("scanning #%s", sku)

	itemsInCard := c.items[sku]
	c.items[sku] = itemsInCard + 1
}

// GetTotalPrice returns total price of items in your cart (including special offers).
func (c *Cashier) GetTotalPrice() (int, error) {
	total := 0
	for sku, count := range c.items {
		amount, err := c.calculator.CalculateItemsPrice(stock.SKU(sku), count)
		if err != nil {
			return 0, errors.Wrap(err, "failed to get item price")
		}
		total += amount
	}

	return total, nil
}
