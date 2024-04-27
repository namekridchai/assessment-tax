package incomeTaxAllowance

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type Allowance struct {
	AllowanceType string
	Amount        float64
}

var connStr = os.Getenv("DATABASE_URL")
var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}
