package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/kombolewis/kompani/utils"
	_ "github.com/lib/pq"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load configuration", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
