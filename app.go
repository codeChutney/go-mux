package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	//yes, comments can be added here
	_ "github.com/lib/pq"
)

type Application struct {
	//What's the * here
	Router *mux.Router
	DB     *sql.DB
}

func (a *Application) Initialize(username, password, dbname string) {
	fmt.Printf("Initialize the application using username %v and dbname %v\n", username, dbname)

	//create a connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", username, password, dbname)

	//there is a type error(?!)
	var err error

	a.DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

func (a *Application) Run(address string) {
	fmt.Printf("Attempting to run application on %v\n", address)
}
