package taxCalculator

type TaxLevel struct {
	Level string
	Tax   float64
}

type IncomeTaxReport struct {
	Tax       float64
	TaxLevels []TaxLevel
}

func Report(calcultor TaxCalculator) (report IncomeTaxReport) {
	r := IncomeTaxReport{Tax: calcultor.CalculateTax()}
	taxlevel := []TaxLevel{
		{"", 0},
		{"", 0},
		{"", 0},
		{"", 0},
		{"", 0},
	}
	r.TaxLevels = taxlevel

	return r
}
