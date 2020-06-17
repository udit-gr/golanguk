package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/udit-gr/postgres"
	"log"
)

func init() {

	connectionString := "user=udit dbname=employee sslmode=disable"
	_, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("[ERR] Error while connecting to db, Err : ", err)
	}

}

func main() {

	log.Printf("[INFO] Transaction exection started")

	log.Printf("INFO] Transaction successfully executed")

}
