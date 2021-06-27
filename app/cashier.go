package app

import (
	"github.com/apaliavy/checkout/app/stock"

	"github.com/pkg/errors"
)

// nolint: lll
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 --fake-name CalculatorMock -o ../testing/mocks/calculator/calculator.go . Calculator

type Calculator interface {
	CalculateItemsPrice(sku stock.SKU, quantity int) (int, error)
}

type Cashier struct {
	calculator Calculator
	items      map[string]int
}

func NewCashier(c Calculator) *Cashier {
	return &Cashier{
		calculator: c,
		items:      make(map[string]int),
	}
}

func (c *Cashier) Scan(sku string) {
	itemsInCard := c.items[sku]
	c.items[sku] = itemsInCard + 1
}

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
