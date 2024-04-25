package incomeTaxAllowance

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Allowance struct {
	AllowanceType string
	Amount        float64
}

const connStr = "user=postgres password=postgres dbname=ktaxes sslmode=disable"

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
