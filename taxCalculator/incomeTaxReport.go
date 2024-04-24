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

	if 150000 < calcultor.TotalIncome() && calcultor.TotalIncome() <= 500000 {
		taxlevel[1].Tax = r.Tax
	}

	r.TaxLevels = taxlevel

	return r
}
