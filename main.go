package main

import (
	"database/sql"
	"fmt"
	server "go-kit-demo/server"
	"os"

	"github.com/go-kit/log"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5455
	user     = "postgres"
	password = "postgres"
	dbname   = "demo"
)

func GetDbConn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("connected fail!")
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("ping fail!")
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db
}

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	db := GetDbConn()
	defer db.Close()
	server.HttpRun(db, logger)
	// server.GrpcRun(db, logger)
}
