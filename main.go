package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rizkiamr/go-bookshelf-api/api"
	constants "github.com/rizkiamr/go-bookshelf-api/constant"
	db "github.com/rizkiamr/go-bookshelf-api/db/sqlc"
)

func main() {
	conn, err := sql.Open(constants.DbDriver, constants.DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(constants.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
