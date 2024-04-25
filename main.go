package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	calculation "github.com/namekridchai/assessment_tax/incomeTax/calculation"
)

const portNum string = ":8080"
const connStr = "user=postgres password=postgres dbname=ktaxes sslmode=disable"

var db *sql.DB

func UpdatePersonalDeduction(c echo.Context) error {

	var request struct{ Amount float64 }
	err := c.Bind(&request)
	if err != nil {
		return err
	}

	if request.Amount < 0 || request.Amount > 100000 {
		return c.JSON(http.StatusBadRequest, "personal allowance should between 1 and 100000")
	}

	statementUpdate, err := db.Prepare("UPDATE public.allowance_master set personal = $1")
	if err != nil {
		return err
	}
	defer statementUpdate.Close()

	_, err = statementUpdate.Exec(request.Amount)
	if err != nil {
		return err
	}

	response := struct{ PersonalDeduction float64 }{request.Amount}
	return c.JSON(http.StatusCreated, response)

}

func UpdateKrcp(c echo.Context) error {

	var request struct{ Amount float64 }
	err := c.Bind(&request)
	if err != nil {
		return err
	}

	if request.Amount < 0 || request.Amount > 100000 {
		return c.JSON(http.StatusBadRequest, "k_receipt should between 1 and 100000")
	}

	statementUpdate, err := db.Prepare("UPDATE public.allowance_master set k_receipt = $1")
	if err != nil {
		return err
	}
	defer statementUpdate.Close()

	_, err = statementUpdate.Exec(request.Amount)
	if err != nil {
		return err
	}

	response := struct {
		KReceipt float64 `json:"kReceipt"`
	}{request.Amount}
	fmt.Println(response)
	return c.JSON(http.StatusCreated, response)

}

func HandleBasicAuth(username string, password string, c echo.Context) (bool, error) {
	if username == "adminTax" && password == "admin!" {
		return true, nil
	}
	return false, nil
}

func main() {

	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()

	if pingErr != nil {
		log.Fatal(pingErr)
	}

	fmt.Println("Connected! to db")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Go Bootcamp!")
	})
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(HandleBasicAuth))
	g.POST("/deductions/personal", UpdatePersonalDeduction)
	g.POST("/deductions/k-receipt", UpdateKrcp)
	e.POST("tax/calculations", calculation.Calculation)

	go func() {
		e.Logger.Fatal(e.Start(portNum))
	}()

	gracefulStop := make(chan os.Signal, 1)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)

	<-gracefulStop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down server %s", err)
	} else {
		fmt.Println("Server gracefully stopped")
	}

	if err := db.Close(); err != nil {
		fmt.Printf("Error closing db connection %s", err)
	} else {
		fmt.Println("DB connection gracefully closed")
	}
}
