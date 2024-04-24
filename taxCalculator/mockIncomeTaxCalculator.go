package taxCalculator

type mockIncomeTaxCalculator struct {
	tax               float64
	totalIncome       float64
	wht               float64
	allowances        []allowance
	personalAllowance float64
	adminKrcp         float64
}

func (m mockIncomeTaxCalculator) Wht() float64 {
	return m.wht
}

func (m mockIncomeTaxCalculator) Allowances() []allowance {
	return m.allowances
}

func (m mockIncomeTaxCalculator) PersonalAllowance() float64 {
	return m.personalAllowance
}

func (m mockIncomeTaxCalculator) AdminKrcp() float64 {
	return m.adminKrcp
}

func (m mockIncomeTaxCalculator) TotalIncome() float64 {
	return m.totalIncome
}

func (m mockIncomeTaxCalculator) CalculateTax() float64 {
	return m.tax
}

func (m *mockIncomeTaxCalculator) CalculateTaxShouldReturn(tax float64) {
	m.tax = tax
}
