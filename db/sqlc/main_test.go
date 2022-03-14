package db

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	"github.com/khilmi-aminudin/simplebankv1/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannt load config : ", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
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
