package taxCalculator

import (
	"fmt"
	"testing"
)

func TestReport(t *testing.T) {
	tests := []struct {
		income float64
		tax    float64
		want   [5]float64
	}{
		{income: 150000.0, tax: 0.0, want: [5]float64{0.0, 0.0, 0.0, 0.0, 0.0}},
		{income: 500000.0, tax: 35000.0, want: [5]float64{0.0, 35000.0, 0.0, 0.0, 0.0}},
		{income: 499999.0, tax: 34999.9, want: [5]float64{0.0, 34999.9, 0.0, 0.0, 0.0}},
		{income: 150001.0, tax: 0.1, want: [5]float64{0.0, 0.1, 0.0, 0.0, 0.0}},
		{income: 500001.0, tax: 35000.15, want: [5]float64{0.0, 35000, 0.15, 0.0, 0.0}},
		{income: 500002.0, tax: 35000.30, want: [5]float64{0.0, 35000, 0.30, 0.0, 0.0}},
		{income: 1000000.0, tax: 110000, want: [5]float64{0.0, 35000, 75000, 0.0, 0.0}},
		{income: 1000001.0, tax: 110000.2, want: [5]float64{0.0, 35000, 75000, 0.2, 0.0}},
		{income: 1000002.0, tax: 110000.4, want: [5]float64{0.0, 35000, 75000, 0.4, 0.0}},
		{income: 2000000.0, tax: 310000, want: [5]float64{0.0, 35000, 75000, 200000.0, 0.0}},
		{income: 2000001.0, tax: 310000.35, want: [5]float64{0.0, 35000, 75000, 200000.0, 0.35}},
		{income: 3000000.0, tax: 660000.0, want: [5]float64{0.0, 35000, 75000, 200000.0, 350000.0}},
	}
	for _, test := range tests {
		test_description := fmt.Sprintf("tax level should be %v when income is %v",
			test.want, test.income,
		)
		t.Run(test_description, func(t *testing.T) {
			m := mockIncomeTaxCalculator{totalIncome: test.income}
			want := test.want
			m.CalculateTaxShouldReturn(test.tax)

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
