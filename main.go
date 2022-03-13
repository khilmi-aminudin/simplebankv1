package main

import (
	"database/sql"
	"log"
	"time"

	"github.com/khilmi-aminudin/simplebankv1/api"
	db "github.com/khilmi-aminudin/simplebankv1/db/sqlc"
	_ "github.com/lib/pq"
)

const (
	serverAddres = "0.0.0.0:8080"
)

func main() {
	dbConnection, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/simplebank?sslmode=disable")
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	dbConnection.SetMaxOpenConns(500)
	dbConnection.SetMaxIdleConns(50)
	dbConnection.SetConnMaxIdleTime(time.Minute * 10)
	dbConnection.SetConnMaxLifetime(time.Hour * 1)

	store := db.NewStore(dbConnection)
	server := api.NewServer(store)

	err = server.Start(serverAddres)
	if err != nil {
		log.Fatalf("cannot start server, err : %v", err)
	}

}
