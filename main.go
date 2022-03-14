package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/khilmi-aminudin/simplebankv1/api"
	db "github.com/khilmi-aminudin/simplebankv1/db/sqlc"
	"github.com/khilmi-aminudin/simplebankv1/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannt load config : ", err)
	}

	dbConnection, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	dbConnection.SetMaxOpenConns(500)
	dbConnection.SetMaxIdleConns(50)
	dbConnection.SetConnMaxIdleTime(time.Minute * 10)
	dbConnection.SetConnMaxLifetime(time.Hour * 1)

	store := db.NewStore(dbConnection)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalf("cannot start server, err : %v", err)
	}

}
