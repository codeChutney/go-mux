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
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {

	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {

	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products 
(
	id SERIAL,
	name TEXT NOT NULL
)`
