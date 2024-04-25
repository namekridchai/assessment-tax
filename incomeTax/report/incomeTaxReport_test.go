package incomeTaxReport

import (
	"fmt"
	"testing"

	calculator "github.com/namekridchai/assessment_tax/incomeTax/calculator"
)

func TestCreateReport(t *testing.T) {
	tests := []struct {
		income float64
		tax    float64
		want   [5]float64
	}{
		{income: 150000.0, want: [5]float64{0.0, 0.0, 0.0, 0.0, 0.0}},
		{income: 500000.0, want: [5]float64{0.0, 35000.0, 0.0, 0.0, 0.0}},
		{income: 499999.0, want: [5]float64{0.0, 34999.9, 0.0, 0.0, 0.0}},
		{income: 150001.0, want: [5]float64{0.0, 0.1, 0.0, 0.0, 0.0}},
		{income: 500001.0, want: [5]float64{0.0, 35000, 0.15, 0.0, 0.0}},
		{income: 500002.0, want: [5]float64{0.0, 35000, 0.30, 0.0, 0.0}},
		{income: 1000000.0, want: [5]float64{0.0, 35000, 75000, 0.0, 0.0}},
		{income: 1000001.0, want: [5]float64{0.0, 35000, 75000, 0.2, 0.0}},
		{income: 1000002.0, want: [5]float64{0.0, 35000, 75000, 0.4, 0.0}},
		{income: 2000000.0, want: [5]float64{0.0, 35000, 75000, 200000.0, 0.0}},
		{income: 2000001.0, want: [5]float64{0.0, 35000, 75000, 200000.0, 0.35}},
		{income: 3000000.0, want: [5]float64{0.0, 35000, 75000, 200000.0, 350000.0}},
	}
	for _, test := range tests {
		test_description := fmt.Sprintf("tax level should be %v when income is %v",
			test.want, test.income,
		)
		t.Run(test_description, func(t *testing.T) {
			m := calculator.MockIncomeTaxCalculator{TotalIncome: test.income}
			want := test.want
			m.NetIncomeShouldReturn(test.income)

			incomeTaxReport := CreateReport(m)

			taxLevels := incomeTaxReport.TaxLevels
			var got [5]float64
			for i := 0; i < len(taxLevels); i++ {
				got[i] = taxLevels[i].Tax
			}

			if got != want {
				t.Errorf("got = %v, want %v", got, want)
			}
		})
	}

}

func TestCreateReportWithWht(t *testing.T) {
	tests := []struct {
		income float64
		wht    float64
		want   [5]float64
	}{
		{income: 150000.0, wht: 1000, want: [5]float64{0, 0, 0, 0, 0}},
		{income: 500000.0, wht: 1000, want: [5]float64{0, 34000, 0, 0, 0}},
		{income: 500000.0, wht: 35000, want: [5]float64{0, 0, 0, 0, 0}},
		{income: 500000.0, wht: 37000, want: [5]float64{0, 0, 0, 0, 0}},
		{income: 1000000.0, wht: 0, want: [5]float64{0, 35000, 75000, 0, 0}},
		{income: 1000000.0, wht: 1000, want: [5]float64{0, 35000, 74000, 0, 0}},
		{income: 1000000.0, wht: 75000, want: [5]float64{0, 35000, 0, 0.0, 0}},
		{income: 1000000.0, wht: 76000, want: [5]float64{0, 34000, 0, 0.0, 0}},
		{income: 1000000.0, wht: 110000, want: [5]float64{0, 0, 0, 0.0, 0}},
		{income: 1000000.0, wht: 120000, want: [5]float64{0, 0, 0, 0.0, 0}},
		{income: 2000000.0, wht: 0, want: [5]float64{0, 35000, 75000, 200000, 0}},
		{income: 2000000.0, wht: 100000, want: [5]float64{0, 35000, 75000, 100000, 0}},
		{income: 2000000.0, wht: 200000, want: [5]float64{0, 35000, 75000, 0.0, 0}},
		{income: 2000000.0, wht: 210000, want: [5]float64{0, 35000, 65000, 0.0, 0}},
		{income: 2000000.0, wht: 275000, want: [5]float64{0, 35000, 0, 0, 0}},
		{income: 2000000.0, wht: 280000, want: [5]float64{0, 30000, 0, 0, 0}},
		{income: 2000000.0, wht: 310000, want: [5]float64{0, 0, 0, 0, 0}},
		{income: 2000000.0, wht: 320000, want: [5]float64{0, 0, 0, 0, 0}},
		{income: 3000000.0, wht: 0, want: [5]float64{0.0, 35000, 75000, 200000, 350000}},
		{income: 3000000.0, wht: 50000, want: [5]float64{0, 35000, 75000, 200000, 300000}},
		{income: 3000000.0, wht: 350000, want: [5]float64{0, 35000, 75000, 200000, 0}},
		{income: 3000000.0, wht: 450000, want: [5]float64{0, 35000, 75000, 100000, 0}},
		{income: 3000000.0, wht: 550000, want: [5]float64{0, 35000, 75000, 0, 0}},
		{income: 3000000.0, wht: 560000, want: [5]float64{0, 35000, 65000, 0, 0}},
		{income: 3000000.0, wht: 625000, want: [5]float64{0, 35000, 0, 0, 0}},
		{income: 3000000.0, wht: 630000, want: [5]float64{0, 30000, 0, 0, 0}},
		{income: 3000000.0, wht: 660000, want: [5]float64{0, 0, 0, 0, 0}},
		{income: 3000000.0, wht: 680000, want: [5]float64{0, 0, 0, 0, 0}},
	}
	for _, test := range tests {
		test_description := fmt.Sprintf("tax level should be %v when income is %v",
			test.want, test.income,
		)
		t.Run(test_description, func(t *testing.T) {
			m := calculator.MockIncomeTaxCalculator{TotalIncome: test.income}
			m.SetWht(test.wht)
			want := test.want
			m.NetIncomeShouldReturn(test.income)

			incomeTaxReport := CreateReport(m)

			taxLevels := incomeTaxReport.TaxLevels
			var got [5]float64
			for i := 0; i < len(taxLevels); i++ {
				got[i] = taxLevels[i].Tax
			}

			if got != want {
				t.Errorf("got = %v, want %v", got, want)
			}
		})
	}

}

func TestCreateReportWithRefund(t *testing.T) {
	tests := []struct {
		income float64
		wht    float64
		want   float64
	}{
		{income: 150000.0, wht: 1000, want: 1000},
		{income: 500000.0, wht: 36000, want: 1000},
		{income: 500000.0, wht: 35000, want: 0},
		{income: 500000.0, wht: 34000, want: 0},
	}
	for _, test := range tests {
		test_description := fmt.Sprintf("tax level should be %v when income is %v",
			test.want, test.income,
		)
		t.Run(test_description, func(t *testing.T) {
			m := calculator.MockIncomeTaxCalculator{TotalIncome: test.income}
			m.SetWht(test.wht)
			want := test.want
			m.NetIncomeShouldReturn(test.income)

			incomeTaxReport := CreateReport(m)

			got := incomeTaxReport.TaxRefund
			if got != want {
				t.Errorf("got = %v, want %v", got, want)
			}
		})
	}

}
