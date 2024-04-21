package taxCalculator

import (
	"fmt"
	"testing"
)

func TestCalculateTax(t *testing.T) {

	tests := []struct {
		income            float64
		personalAllowance float64
		want              float64
	}{
		{income: 0.0, personalAllowance: 60000, want: 0.0},
		{income: 150000.0, personalAllowance: 60000, want: 0.0},
		{income: 149999.0, personalAllowance: 60000, want: 0.0},
		{income: 1.0, personalAllowance: 60000, want: 0.0},
		{income: 150001.0, personalAllowance: 60000, want: 0.0},
		{income: 150002.0, personalAllowance: 60000, want: 0.0},
		{income: 210000.0, personalAllowance: 60000, want: 0.0},
		{income: 209999.0, personalAllowance: 60000, want: 0.0},
		{income: 150001.0, personalAllowance: 0, want: 0.1},
		{income: 150002.0, personalAllowance: 0, want: 0.2},
		{income: 500000.0, personalAllowance: 0, want: 35000},
		{income: 499999.0, personalAllowance: 0, want: 34999.9},
		{income: 500000.0, personalAllowance: 60000, want: 29000.0},
		{income: 500001.0, personalAllowance: 0, want: 35000.15},
		{income: 500002.0, personalAllowance: 0, want: 35000.30},
		{income: 1000000.0, personalAllowance: 0, want: 110000},
		{income: 1000000.0, personalAllowance: 60000, want: 101000},
		{income: 1000001.0, personalAllowance: 0, want: 110000.2},
		{income: 1000002.0, personalAllowance: 0, want: 110000.4},
		{income: 2000000.0, personalAllowance: 0, want: 310000},
		{income: 2000001.0, personalAllowance: 0, want: 310000.35},
		{income: 3000000.0, personalAllowance: 0, want: 660000},
	}

	for _, test := range tests {
		test_description := fmt.Sprintf("should return %v when income is %v and personal allowance is %v",
			test.want, test.income, test.personalAllowance,
		)
		t.Run(test_description, func(t *testing.T) {
			a := allowance{AllowanceType: "donation", Amount: 0.0}
			incomeTaxCalculator := IncomeTaxCalculator{TotalIncome: test.income, Wht: 0.0}
			incomeTaxCalculator.addAllowance(a)

			want := test.want

			got := incomeTaxCalculator.CalculateTax(test.personalAllowance)

			if got != want {
				t.Errorf("got = %v, want %v", got, want)
			}
		})
	}

}
func TestCalculateTaxWithWht(t *testing.T) {
	test_description := fmt.Sprintf("should return %v when income is %v and personal allowance is %v",
		600000.0, 3000000.0, 0.0,
	)
	t.Run(test_description, func(t *testing.T) {

		incomeTaxCalculator := IncomeTaxCalculator{TotalIncome: 3000000.0, Wht: 60000.0}

		want := 600000.0

		got := incomeTaxCalculator.CalculateTax(0)

		if got != want {
			t.Errorf("got = %v, want %v", got, want)
		}
	})
}