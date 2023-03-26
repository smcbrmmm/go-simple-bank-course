package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"simplebank/api"
	db "simplebank/db/sqlc"
	"simplebank/util"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
)

func main() {

	s1 := rand.NewSource(time.Now().UnixMicro())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(1000))

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	store := db.NewStore(conn)

	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server at port : 8000")
	}

}
