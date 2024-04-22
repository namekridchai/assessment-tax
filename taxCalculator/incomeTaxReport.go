package taxCalculator

type TaxLevel struct {
	Level string
	Tax   float64
}

type IncomeTaxReport struct {
	Tax       float64
	TaxLevels []TaxLevel
}

func Report(calcultor IncomeTaxCalculator) (report IncomeTaxReport) {
	return IncomeTaxReport{Tax: calcultor.CalculateTax()}
}
