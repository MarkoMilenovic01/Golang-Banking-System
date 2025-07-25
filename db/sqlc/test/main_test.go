package test

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/marko/simplebank/db/sqlc"
	"github.com/marko/simplebank/util"
	"log"
	"os"
	"testing"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../../../")
	if err != nil {
		log.Fatal(err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
