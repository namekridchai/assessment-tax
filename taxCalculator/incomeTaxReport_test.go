package taxCalculator

import (
	"fmt"
	"testing"
)

func TestReport(t *testing.T) {
	tests := []struct {
		income float64
		want   float64
	}{
		{income: 150000.0, want: 0.0},
		{income: 500000.0, want: 35000.0},
		{income: 1000000.0, want: 110000.0},
		{income: 2000000.0, want: 310000.0},
	}
	for _, test := range tests {
		test_description := fmt.Sprintf("tax should be %v when income is %v",
			test.want, test.income,
		)
		t.Run(test_description, func(t *testing.T) {

			income := test.income
			want := test.want

			incomeTaxReport := Report(nil, 0.0, income)
			got := incomeTaxReport.Tax

			if got != want {
				t.Errorf("got = %v, want %v", got, want)
			}
		})
	}

}
