package incomeTaxCalculator

type MockIncomeTaxCalculator struct {
	tax         float64
	TotalIncome float64
	wht         float64
	netIncome   float64
}

func (m MockIncomeTaxCalculator) Wht() float64 {
	return m.wht
}

func (m MockIncomeTaxCalculator) CalculateTax() float64 {
	return m.tax
}

func (m MockIncomeTaxCalculator) NetIncome() float64 {
	return m.netIncome
}

func (m *MockIncomeTaxCalculator) CalculateTaxShouldReturn(tax float64) {
	m.tax = tax
}

func (m *MockIncomeTaxCalculator) NetIncomeShouldReturn(netIncome float64) {
	m.netIncome = netIncome
}

func (m *MockIncomeTaxCalculator) SetWht(wht float64) {
	m.wht = wht
}
