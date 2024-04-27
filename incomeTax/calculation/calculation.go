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

	mapColToFn := make(map[string]func(*incomeTaxCalculator.IncomeTaxCalculator, string))
	mapColToFn["totalIncome"] = setTotalIncome
	mapColToFn["wht"] = setTotalWht
	mapColToFn["donation"] = setDonation

	var taxes []tax
	for line := 1; line < len(stringByline); line++ {
		calculator := &incomeTaxCalculator.IncomeTaxCalculator{}
		content := strings.Split(stringByline[line], ",")
		for col := 0; col < len(content); col++ {
			colName := strings.TrimSpace(mapIndexToColName[col])
			mapColToFn[colName](calculator, content[col])
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

func setTotalIncome(c *incomeTaxCalculator.IncomeTaxCalculator, income string) {
	incomeFloat, _ := strconv.ParseFloat(strings.TrimSpace(income), 64)
	c.TotalIncome = incomeFloat

}

func setTotalWht(c *incomeTaxCalculator.IncomeTaxCalculator, wht string) {
	whtFloat, _ := strconv.ParseFloat(strings.TrimSpace(wht), 64)
	c.Wht = whtFloat

}

func setDonation(c *incomeTaxCalculator.IncomeTaxCalculator, donation string) {
	donationFloat, _ := strconv.ParseFloat(strings.TrimSpace(donation), 64)

	a := incomeTaxAllowance.Allowance{AllowanceType: "donation", Amount: donationFloat}
	c.AddAllowance(a)
}
