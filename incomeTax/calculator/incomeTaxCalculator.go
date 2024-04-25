package incomeTaxCalculator

import (
	"strings"

	incomeTaxAllowance "github.com/namekridchai/assessment_tax/incomeTax/allowance"
)

type TaxCalculator interface {
	CalculateTax() float64
}

type allowance incomeTaxAllowance.Allowance

type IncomeTaxCalculatorInterface interface {
	CalculateTax() float64
	GetWht() float64
	NetIncome() float64
}

type IncomeTaxCalculator struct {
	TotalIncome       float64
	Wht               float64
	Allowances        []allowance
	personalAllowance float64
	adminKrcp         float64
}

func (i IncomeTaxCalculator) GetWht() float64 {
	return i.Wht
}

func (i *IncomeTaxCalculator) addAllowance(a allowance) {
	i.Allowances = append(i.Allowances, a)
}

func (i IncomeTaxCalculator) NetIncome() float64 {
	netIncome := max(i.TotalIncome-i.personalAllowance, 0)
	allowanceMap := make(map[string]float64)
	allowanceMap["donation"] = 100000.0
	allowanceMap["k-receipt"] = i.adminKrcp

	for _, a := range i.Allowances {
		netIncome -= min(a.Amount, allowanceMap[strings.ToLower(a.AllowanceType)])
	}
	return netIncome
}

func (i IncomeTaxCalculator) CalculateTax() float64 {

	netIncome := i.NetIncome()

	out := TaxStep1(netIncome) + TaxStep2(netIncome) + TaxStep3(netIncome) +
		TaxStep4(netIncome)

	return max(out-i.Wht, 0)

}

func (i *IncomeTaxCalculator) SetPersonalAllowance(personalAllowance float64) {
	i.personalAllowance = personalAllowance

}

func (i *IncomeTaxCalculator) SetAdminKrcp(adminKrcp float64) {
	i.adminKrcp = adminKrcp
}

func TaxStep1(netIncome float64) float64 {
	if 150000 < netIncome && netIncome <= 500000 {
		return (netIncome - 150000) * 0.1
	} else if netIncome > 500000 {
		return (500000 - 150000) * 0.1
	}
	return 0
}

func TaxStep2(netIncome float64) float64 {
	if 500000 < netIncome && netIncome <= 1000000 {
		return (netIncome - 500000) * 0.15
	} else if netIncome > 1000000 {
		return (1000000 - 500000) * 0.15
	}
	return 0
}

func TaxStep3(netIncome float64) float64 {
	if 1000000 < netIncome && netIncome <= 2000000 {
		return (netIncome - 1000000) * 0.2
	} else if netIncome > 1000000 {
		return (2000000 - 1000000) * 0.2
	}
	return 0
}

func TaxStep4(netIncome float64) float64 {
	if netIncome > 2000000 {
		return (netIncome - 2000000) * 0.35
	}
	return 0

}
