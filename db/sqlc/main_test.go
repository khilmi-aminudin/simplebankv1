package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testDB.SetMaxOpenConns(500)
	testDB.SetMaxIdleConns(50)
	testDB.SetConnMaxIdleTime(time.Minute * 10)
	testDB.SetConnMaxLifetime(time.Hour * 1)

	testQueries = New(testDB)

	os.Exit(m.Run())
}
