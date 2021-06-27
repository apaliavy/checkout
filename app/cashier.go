package app

import "github.com/apaliavy/checkout/app/price"

// Checkout system that handles price schemes such as "pineapples cost 50, three pineapples cost 130."
// An example of how app interface could look like
type Checkout interface {
	Scan(item string)
	GetTotalPrice() (int, error)
}

type Cashier struct {
	// todo: implement
}

// todo: implement
func NewCashier(calc *price.Calculator) *Cashier {
	return &Cashier{}
}

func (c *Cashier) Scan(item string) {}

func (c *Cashier) GetTotalPrice() (int, error) {
	return 0, nil
}
