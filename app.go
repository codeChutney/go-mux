package main

import (
	"database/sql"
	"fmt"

	"github.com/gorilla/mux"
	//yes, comments can be added here
	_ "github.com/lib/pq"
)

type Application struct {
	//What's the * here
	Router *mux.Router
	DB     *sql.DB
}

//TODO - come back and implement these

func (a *Application) Initialize(username, password, dbname string) {
	fmt.Printf("Initialize the application using username %v and dbname %v", username, dbname)
}

func (a *Application) Run(address string) {}
