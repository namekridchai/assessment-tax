package taxCalculator

import (
	"fmt"
	"testing"
)

func TestReport(t *testing.T) {
	test_description := fmt.Sprintf("tax should be %v when income is %v",
		0.0, 150000.0,
	)
	t.Run(test_description, func(t *testing.T) {

		income := 150000.0
		want := 0.0

		incomeTaxReport := Report(nil, 0.0, income)
		got := incomeTaxReport.Tax

		if got != want {
			t.Errorf("got = %v, want %v", got, want)
		}
	})
}
