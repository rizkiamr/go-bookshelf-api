package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rizkiamr/go-bookshelf-api/api"
	db "github.com/rizkiamr/go-bookshelf-api/db/sqlc"
)

const (
	DbDriver      = "postgres"
	DbSource      = "postgresql://postgres:postgres@127.0.0.1:5432/bookshelf?sslmode=disable"
	ServerAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(DbDriver, DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
