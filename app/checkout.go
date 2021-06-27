package app

// Checkout system that handles pricing schemes such as "pineapples cost 50, three pineapples cost 130."
// An example of how app interface could look like
type Checkout interface {
	Scan(item string)
	GetTotalPrice() int
}
