package taxCalculator

type TaxLevel struct {
	Level string
	Tax   float64
}

type IncomeTaxReport struct {
	Tax       float64
	TaxLevels []TaxLevel
}

func CreateReport(calcultor incomeTaxCalculatorInterface) (report IncomeTaxReport) {
	r := IncomeTaxReport{Tax: calcultor.CalculateTax()}
	netIncome := calcultor.TotalIncome()
	taxlevel := []TaxLevel{
		{"", 0},
		{"", taxStep1(netIncome)},
		{"", taxStep2(netIncome)},
		{"", taxStep3(netIncome)},
		{"", 0},
	}

	r.TaxLevels = taxlevel

	return r
}
