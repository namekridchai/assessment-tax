package taxCalculator

type TaxLevel struct {
	Level string
	Tax   float64
}

type IncomeTaxReport struct {
	Tax       float64
	TaxLevels []TaxLevel
}

func Report(a []allowance, wht float64, income float64) (i IncomeTaxReport) {
	incomeTaxReport := IncomeTaxReport{}
	incomeTaxReport.Tax = taxStep1(income) + taxStep2(income)
	return incomeTaxReport
}
