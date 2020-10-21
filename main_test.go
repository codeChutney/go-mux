package main

import (
	"fmt"
	"log"

	"os"
	"testing"
)

var a Application

func TestMain(m *testing.M) {

	a.Initialize("postgres", "Pass2020!", "postgres")
	ensureTableExists()

	fmt.Printf("Start TestMain func with %v and %T\n", m, m)
	code := m.Run()
	fmt.Printf("After running TestMain code is %v and it is of type %T\n", code, code)
	os.Exit(code)
}

func ensureTableExists() {

	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products 
(
	id SERIAL,
	name TEXT NOT NULL
)`
