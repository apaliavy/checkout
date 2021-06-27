package app_test

import (
	"fmt"
	"testing"

	"github.com/apaliavy/checkout/testing/mocks/calculator"

	"github.com/stretchr/testify/require"

	"github.com/apaliavy/checkout/app"
	"github.com/stretchr/testify/assert"
)

func TestNewCashier(t *testing.T) {
	t.Parallel()

	cashier := app.NewCashier(&calculator.CalculatorMock{})
	assert.NotNil(t, cashier)
}

func TestCashier_GetTotalPrice(t *testing.T) {
	t.Parallel()

	happyCalculator := &calculator.CalculatorMock{}
	happyCalculator.CalculateItemsPriceReturns(100, nil)

	grumpyCalculator := &calculator.CalculatorMock{}
	grumpyCalculator.CalculateItemsPriceReturns(0, fmt.Errorf("just an error"))

	cases := []struct {
		name               string
		calculator         app.Calculator
		itemsToScan        []string
		expectedTotalPrice int
		expectError        bool
	}{
		{
			name:               "check cashier returns zero when 0 elements scanned",
			calculator:         &calculator.CalculatorMock{},
			itemsToScan:        []string{},
			expectedTotalPrice: 0,
		},
		{
			name:               "check cashier returns non-zero price",
			calculator:         happyCalculator,
			itemsToScan:        []string{"A", "B", "A", "C"},
			expectedTotalPrice: 300,
		},
		{
			name:               "check cashier returns error",
			calculator:         grumpyCalculator,
			itemsToScan:        []string{"A", "B", "A", "C"},
			expectedTotalPrice: 0,
			expectError:        true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			cashier := app.NewCashier(tc.calculator)
			for _, i := range tc.itemsToScan {
				cashier.Scan(i)
			}

			total, err := cashier.GetTotalPrice()

			if tc.expectError {
				require.Error(t, err)
				return
			}

			assert.Equal(t, tc.expectedTotalPrice, total)
		})
	}
}
