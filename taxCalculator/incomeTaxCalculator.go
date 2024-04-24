package taxCalculator

import "strings"

type TaxCalculator interface {
	CalculateTax() float64
}

type incomeTaxCalculatorInterface interface {
	CalculateTax() float64
	Wht() float64
	Allowances() []allowance
	PersonalAllowance() float64
	AdminKrcp() float64
	TotalIncome() float64
	NetIncome() float64
}

type IncomeTaxCalculator struct {
	totalIncome       float64
	wht               float64
	allowances        []allowance
	personalAllowance float64
	adminKrcp         float64
}

func (i IncomeTaxCalculator) Wht() float64 {
	return i.wht
}

func (i IncomeTaxCalculator) Allowances() []allowance {
	return i.allowances
}

func (i IncomeTaxCalculator) PersonalAllowance() float64 {
	return i.personalAllowance
}

func (i IncomeTaxCalculator) AdminKrcp() float64 {
	return i.adminKrcp
}

func (i IncomeTaxCalculator) TotalIncome() float64 {
	return i.totalIncome
}

func (i *IncomeTaxCalculator) addAllowance(a allowance) {
	i.allowances = append(i.allowances, a)
}

func (i *IncomeTaxCalculator) NetIncome() float64 {
	netIncome := max(i.TotalIncome()-i.personalAllowance, 0)
	allowanceMap := make(map[string]float64)
	allowanceMap["donation"] = 100000.0
	allowanceMap["k-receipt"] = i.adminKrcp

	for _, a := range i.Allowances() {
		netIncome -= min(a.Amount, allowanceMap[strings.ToLower(a.AllowanceType)])
	}

	return netIncome
}

func (i IncomeTaxCalculator) CalculateTax() float64 {

	netIncome := i.NetIncome()

	out := taxStep1(netIncome) + taxStep2(netIncome) + taxStep3(netIncome) +
		taxStep4(netIncome)

	return out - i.Wht()

}

func taxStep1(netIncome float64) float64 {
	if 150000 < netIncome && netIncome <= 500000 {
		return (netIncome - 150000) * 0.1
	} else if netIncome > 500000 {
		return (500000 - 150000) * 0.1
	}
	return 0
}

func taxStep2(netIncome float64) float64 {
	if 500000 < netIncome && netIncome <= 1000000 {
		return (netIncome - 500000) * 0.15
	} else if netIncome > 1000000 {
		return (1000000 - 500000) * 0.15
	}
	return 0
}

func taxStep3(netIncome float64) float64 {
	if 1000000 < netIncome && netIncome <= 2000000 {
		return (netIncome - 1000000) * 0.2
	} else if netIncome > 1000000 {
		return (2000000 - 1000000) * 0.2
	}
	return 0
}

func taxStep4(netIncome float64) float64 {
	if netIncome > 2000000 {
		return (netIncome - 2000000) * 0.35
	}
	return 0

}
