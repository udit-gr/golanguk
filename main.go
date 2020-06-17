package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/udit-gr/golanguk/postgres"
	"log"
)

var db *sql.DB

func init() {

	var err error

	connectionString := "user=udit dbname=employee sslmode=disable"
	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("[ERR] Error while connecting to db, Err : ", err)
	}

}

func main() {

	log.Printf("[INFO] Transaction exection started")

	employeeInfo := postgres.EmployeeDetails{
		Name:       "Ben",
		Age:        20,
		Gender:     "Male",
		Department: "Technology",
		Position:   "Software Engineer",
		Manager:    "Smith",
	}

	err := postgres.InsertEmployeeDetails(db, employeeInfo)
	if err != nil {
		log.Printf("[ERR] Error while inserting employee details, Err : %+v", err)
		log.Printf("[INFO] Transaction execution failed")
		return
	}

	log.Printf("INFO] Transaction successfully executed")

}
