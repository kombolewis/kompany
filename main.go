package main

import (
	"database/sql"
	"log"

	"github.com/kombolewis/kompani/api"
	db "github.com/kombolewis/kompani/db/sqlc"
	"github.com/kombolewis/kompani/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load configuration", err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start()
	if err != nil {
		log.Fatal("cannot start server", err)
	}

}
