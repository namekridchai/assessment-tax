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
	taxlevel := []TaxLevel{
		{"", 0},
		{"", 0},
		{"", 0},
		{"", 0},
		{"", 0},
	}
	netIncome := calcultor.TotalIncome()

	if 150000 < netIncome && netIncome <= 500000 {
		taxlevel[1].Tax = r.Tax
	}

	if 500000 < netIncome && netIncome <= 1000000 {
		taxlevel[1].Tax = taxStep1(netIncome)
		taxlevel[2].Tax = taxStep2(netIncome)
	}

	r.TaxLevels = taxlevel

	return r
}
