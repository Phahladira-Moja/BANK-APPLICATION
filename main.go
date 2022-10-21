package main

import (
	"database/sql"
	"github.com/phahladira-moja/simple-bank-application/api"
	db "github.com/phahladira-moja/simple-bank-application/db/sqlc"
	"github.com/phahladira-moja/simple-bank-application/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
