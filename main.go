package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/rizkiamr/go-bookshelf-api/api"
	db "github.com/rizkiamr/go-bookshelf-api/db/sqlc"
	"github.com/rizkiamr/go-bookshelf-api/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	DbSource := "postgresql://" + config.DbUser + ":" + config.DbPassword + "@" + config.DbHost + ":" + config.DbPort + "/" + config.DbName + "?sslmode=disable"

	conn, err := sql.Open(config.DbDriver, DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	ServiceAddressPort := config.ServiceAddress + ":" + config.ServicePort

	err = server.Start(ServiceAddressPort)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
