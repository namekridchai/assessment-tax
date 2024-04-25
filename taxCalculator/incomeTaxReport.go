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

	netIncome := calcultor.NetIncome()

	taxlevel := []TaxLevel{
		{"0-150,000", 0},
		{"150,001-500,000", taxStep1(netIncome)},
		{"500,001-1,000,000", taxStep2(netIncome)},
		{"1,000,001-2,000,000", taxStep3(netIncome)},
		{"2,000,001 ขึ้นไป", taxStep4(netIncome)},
	}

	r.TaxLevels = taxlevel

	return r
}
