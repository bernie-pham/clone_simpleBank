package main

import (
	"database/sql"
	"log"

	"github.com/bernie-pham/cloneSimpleBank/api"
	db "github.com/bernie-pham/cloneSimpleBank/db/sqlc"
	"github.com/bernie-pham/cloneSimpleBank/ultilities"

	_ "github.com/lib/pq"
)

func main() {
	config, err := ultilities.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config: ", err)
	}

	conn, err := sql.Open(config.DRIVER_NAME, config.DRIVER_SOURCE)
	if err != nil {
		log.Fatal("Cannot connect to database with this issue: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.SERVER_ADDRESS)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
