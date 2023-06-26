package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/rizkiamr/go-bookshelf-api/util"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		panic(err)
	}

	DbSource := "postgresql://" + config.DbUser + ":" + config.DbPassword + "@" + config.DbHost + ":" + config.DbPort + "/" + config.DbName + "?sslmode=disable"

	conn, err := sql.Open(config.DbDriver, DbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
