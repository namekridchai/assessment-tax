package incomeTaxCalculator

import (
	"fmt"
	"testing"

	incomeTaxAllowance "github.com/namekridchai/assessment_tax/incomeTax/allowance"
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
			a := incomeTaxAllowance.Allowance{AllowanceType: "donation", Amount: 0.0}
			incomeTaxCalculator := IncomeTaxCalculator{TotalIncome: test.income, Wht: 0.0}
			incomeTaxCalculator.AddAllowance(a)
			incomeTaxCalculator.personalAllowance = test.personalAllowance

			want := test.want

			got := incomeTaxCalculator.CalculateTax()

			if got != want {
				t.Errorf("got = %v, want %v", got, want)
			}
		})
	}

}
func TestCalculateTaxWithWht(t *testing.T) {
	tests := []struct {
		totalIncome float64
		wht         float64
		want        float64
	}{
		{totalIncome: 3000000.0, wht: 60000.0, want: 600000.0},
		{totalIncome: 3000000.0, wht: 660000.0, want: 0},
	}
	for _, test := range tests {
		test_description := fmt.Sprintf("should return %v when income is %v and wht is %v",
			test.want, test.totalIncome, test.wht,
		)
		t.Run(test_description, func(t *testing.T) {

			incomeTaxCalculator := IncomeTaxCalculator{TotalIncome: test.totalIncome, Wht: test.wht}

			want := test.want

			got := incomeTaxCalculator.CalculateTax()

			if got != want {
				t.Errorf("got = %v, want %v", got, want)
			}
		})
	}

}

func TestCalculateTaxWithDonationAllowance(t *testing.T) {
	tests := []struct {
		totalIncome         float64
		dontation_allowance float64
		want                float64
	}{
		{totalIncome: 3100000, dontation_allowance: 100000.0, want: 660000.0},
		{totalIncome: 3100000, dontation_allowance: 200000.0, want: 660000.0},
	}

	for _, test := range tests {
		test_description := fmt.Sprintf("should return %v when income is %v and donation allowance is %v",
			test.want, test.totalIncome, test.dontation_allowance,
		)
		t.Run(test_description, func(t *testing.T) {

			incomeTaxCalculator := IncomeTaxCalculator{TotalIncome: test.totalIncome, Wht: 0.0}
			a := incomeTaxAllowance.Allowance{AllowanceType: "donation", Amount: test.dontation_allowance}
			incomeTaxCalculator.AddAllowance(a)

			want := test.want

			got := incomeTaxCalculator.CalculateTax()

			if got != want {
				t.Errorf("got = %v, want %v", got, want)
			}
		})
	}

}

func TestCalculateTaxWithNonExistAllowance(t *testing.T) {

	test_description := fmt.Sprintf("should return %v when income is %v and  allowance is %v",
		660000.0, 3000000.0, 100000.0,
	)
	t.Run(test_description, func(t *testing.T) {

		incomeTaxCalculator := IncomeTaxCalculator{TotalIncome: 3000000.0, Wht: 0.0}
		a := incomeTaxAllowance.Allowance{AllowanceType: "donate to aj.dang guitar", Amount: 100000.0}
		incomeTaxCalculator.AddAllowance(a)

		want := 660000.0

		got := incomeTaxCalculator.CalculateTax()

		if got != want {
			t.Errorf("got = %v, want %v", got, want)
		}
	})
}

func TestCalculateTaxWithKrcpAllowance(t *testing.T) {
	tests := []struct {
		totalIncome float64
		adminKrcp   float64
		krcp        float64
		want        float64
	}{
		{totalIncome: 3100000, adminKrcp: 100000.0, krcp: 100000.0, want: 660000},
		{totalIncome: 3090000, adminKrcp: 100000.0, krcp: 90000.0, want: 660000},
		{totalIncome: 3100000, adminKrcp: 100000.0, krcp: 200000.0, want: 660000},
	}

	for _, test := range tests {
		test_description := fmt.Sprintf("should return %v when income is %v admin allowance is %v user allowance is %v",
			test.want, test.totalIncome, test.adminKrcp, test.krcp)
		t.Run(test_description, func(t *testing.T) {
			incomeTaxCalculator := IncomeTaxCalculator{TotalIncome: test.totalIncome, Wht: 0.0}
			a := incomeTaxAllowance.Allowance{AllowanceType: "k-receipt", Amount: test.krcp}
			incomeTaxCalculator.AddAllowance(a)

			incomeTaxCalculator.adminKrcp = test.adminKrcp

			want := 660000.0

			got := incomeTaxCalculator.CalculateTax()

			if got != want {
				t.Errorf("got = %v, want %v", got, want)
			}
		})
	}

}

func TestCalculateTaxWithMultipleAllowance(t *testing.T) {

	test_description := fmt.Sprintf("should return %v when income is %v krcp is %v donation is %v",
		660000.0, 3200000.0, 100000.0, 100000.0)
	t.Run(test_description, func(t *testing.T) {
		incomeTaxCalculator := IncomeTaxCalculator{TotalIncome: 3200000.0, Wht: 0.0}

		a1 := incomeTaxAllowance.Allowance{AllowanceType: "k-receipt", Amount: 100000.0}
		a2 := incomeTaxAllowance.Allowance{AllowanceType: "donation", Amount: 100000.0}

		incomeTaxCalculator.AddAllowance(a1)
		incomeTaxCalculator.AddAllowance(a2)

		incomeTaxCalculator.adminKrcp = 100000.0

		want := 660000.0

		got := incomeTaxCalculator.CalculateTax()

		if got != want {
			t.Errorf("got = %v, want %v", got, want)
		}
	})

}
