package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	driverSource = "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable"
	driverName   = "postgres"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(driverName, driverSource)
	if err != nil {
		log.Fatal("Cannot connect to database with this issue: ", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
