package taxCalculator

import (
	"fmt"
	"testing"
)

func TestReport(t *testing.T) {
	tests := []struct {
		income float64
		tax    float64
		want   []float64
	}{
		{income: 150000.0, tax: 0.0, want: []float64{0.0, 0.0, 0.0, 0.0, 0.0}},
	}
	for _, test := range tests {
		test_description := fmt.Sprintf("tax level should be %v when income is %v",
			test.want, test.income,
		)
		t.Run(test_description, func(t *testing.T) {
			m := mockTaxCalculator{}
			want := test.tax
			m.CalculateTaxShouldReturn(test.tax)

			incomeTaxReport := Report(m)

			got := incomeTaxReport.Tax

			if got != want {
				t.Errorf("got = %v, want %v", got, want)
			}
		})
	}

}
