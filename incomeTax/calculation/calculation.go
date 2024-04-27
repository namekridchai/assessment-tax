package incomeTaxCalculation

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

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

type tax struct {
	TotalIncome float64
	Tax         float64
}

func CSVCalulation(c echo.Context) error {

	fh, err := c.FormFile("taxFile")
	if err != nil {
		fmt.Println("here")
		return c.JSON(http.StatusInternalServerError, err)

	}

	f, err := fh.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	body, err := io.ReadAll(f)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	myString := string(body[:])
	stringByline := strings.Split(myString, "\n")
	header := stringByline[0]

	mapIndexToColName := make(map[int]string)

	columnNames := strings.Split(header, ",")
	for i := 0; i < len(columnNames); i++ {
		mapIndexToColName[i] = columnNames[i]
	}

	mapColToFn := make(map[string]func(*incomeTaxCalculator.IncomeTaxCalculator, string) error)
	mapColToFn["totalIncome"] = setTotalIncome
	mapColToFn["wht"] = setTotalWht
	mapColToFn["donation"] = setDonation

	var taxes []tax
	for line := 1; line < len(stringByline); line++ {
		calculator := &incomeTaxCalculator.IncomeTaxCalculator{}
		content := strings.Split(stringByline[line], ",")
		for col := 0; col < len(content); col++ {
			colName := strings.TrimSpace(mapIndexToColName[col])
			err := mapColToFn[colName](calculator, content[col])
			if err != nil {
				return c.JSON(http.StatusBadRequest, err.Error())
			}
		}
		calculator.SetPersonalAllowance(incomeTaxAllowance.GetPersonalDeduction())
		calculator.SetAdminKrcp(incomeTaxAllowance.GetKrcp())
		t := tax{calculator.TotalIncome, calculator.CalculateTax()}
		taxes = append(taxes, t)
	}

	reponse := struct {
		Taxes []tax `json:"taxes"`
	}{Taxes: taxes}
	return c.JSON(http.StatusCreated, reponse)
}

func setTotalIncome(c *incomeTaxCalculator.IncomeTaxCalculator, income string) error {
	incomeFloat, err := strconv.ParseFloat(strings.TrimSpace(income), 64)
	if err != nil {
		return fmt.Errorf("income wrong format")
	}
	c.TotalIncome = incomeFloat
	return nil
}

func setTotalWht(c *incomeTaxCalculator.IncomeTaxCalculator, wht string) error {
	whtFloat, err := strconv.ParseFloat(strings.TrimSpace(wht), 64)
	if err != nil {
		return fmt.Errorf("income wrong format")
	}
	c.Wht = whtFloat
	return nil
}

func setDonation(c *incomeTaxCalculator.IncomeTaxCalculator, donation string) error {
	donationFloat, err := strconv.ParseFloat(strings.TrimSpace(donation), 64)
	if err != nil {
		return fmt.Errorf("income wrong format")
	}

	a := incomeTaxAllowance.Allowance{AllowanceType: "donation", Amount: donationFloat}
	c.AddAllowance(a)
	return nil
}
