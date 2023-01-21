package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/trungthienpyf/simplebank/api"
	db "github.com/trungthienpyf/simplebank/db/sqlc"
	"github.com/trungthienpyf/simplebank/util"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	serverAddress="0.0.0.0:8080"
)

func main() {
	config,err :=util.LoadConfig(".")
	if err != nil{
		log.Fatal("cannot load config:",err)
	}
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err =server.Start(config.ServerAddress)
	if err !=nil{
		log.Fatal("cannot start server:",err)
	}
}