package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/bernie-pham/cloneSimpleBank/ultilities"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	var config ultilities.Config
	config, err = ultilities.LoadConfig("../..")

	testDB, err = sql.Open(config.DRIVER_NAME, config.DRIVER_SOURCE)
	if err != nil {
		log.Fatal("Cannot connect to database with this issue: ", err)
	}
	testQueries = New(testDB)

	os.Exit(m.Run())
}
