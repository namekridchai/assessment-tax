package incomeTaxCalculation

import (
	"net/http"

	"github.com/labstack/echo/v4"
	incomeTaxAllowance "github.com/namekridchai/assessment_tax/incomeTax/allowance"
	incomeTaxCalculator "github.com/namekridchai/assessment_tax/incomeTax/calculator"
	incomeTaxReport "github.com/namekridchai/assessment_tax/incomeTax/report"
)

func Calculation(c echo.Context) error {

	var calculator incomeTaxCalculator.IncomeTaxCalculator
	err := c.Bind(&calculator)
	if err != nil {
		return err
	}
	calculator.SetPersonalAllowance(incomeTaxAllowance.GetPersonalDeduction())
	calculator.SetAdminKrcp(incomeTaxAllowance.GetKrcp())
	incomeTaxReport := incomeTaxReport.CreateReport(calculator)
	return c.JSON(http.StatusCreated, incomeTaxReport)

}
