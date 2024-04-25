package incomeTaxReport

import (
	calculator "github.com/namekridchai/assessment_tax/incomeTax/calculator"
)

type incomeTaxCalculatorInterface calculator.IncomeTaxCalculatorInterface

type TaxLevel struct {
	Level string
	Tax   float64
}

type IncomeTaxReport struct {
	Tax       float64
	TaxLevels []TaxLevel
	TaxRefund float64
}

func CreateReport(calcultor incomeTaxCalculatorInterface) (report IncomeTaxReport) {
	r := IncomeTaxReport{Tax: calcultor.CalculateTax()}

	netIncome := calcultor.NetIncome()

	taxlevel := []TaxLevel{
		{"0-150,000", 0},
		{"150,001-500,000", calculator.TaxStep1(netIncome)},
		{"500,001-1,000,000", calculator.TaxStep2(netIncome)},
		{"1,000,001-2,000,000", calculator.TaxStep3(netIncome)},
		{"2,000,001 ขึ้นไป", calculator.TaxStep4(netIncome)},
	}

	wht := calcultor.Wht()
	taxlevelIndex := len(taxlevel) - 1
	for wht > 0 && taxlevelIndex >= 0 {
		if taxlevel[taxlevelIndex].Tax != 0 {
			currentTax := taxlevel[taxlevelIndex].Tax
			newTax := currentTax - wht
			wht -= currentTax
			taxlevel[taxlevelIndex].Tax = max(newTax, 0)
		}
		taxlevelIndex--

	}

	r.TaxLevels = taxlevel
	wht = max(0, wht)
	r.TaxRefund = wht

	return r
}
