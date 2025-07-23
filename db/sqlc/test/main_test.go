package test

import (
	"database/sql"
	_ "github.com/lib/pq" // PostgreSQL driver
	"github.com/marko/simplebank/db/sqlc"
	"log"
	"os"
	"testing"
)

var testQueries *db.Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
