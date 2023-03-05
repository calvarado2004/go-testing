package main

import (
	"database/sql"
	"github.com/calvarado2004/go-testing/pkg/db"
	"log"
	"os"
	"testing"
)

var app application

// TestMain is the entry point for all tests.
func TestMain(m *testing.M) {

	pathToTemplates = "./../../templates/"

	app.Session = getSession()

	app.DSN = "host=localhost port=5432 user=postgres password=postgres dbname=users sslmode=disable timezone=UTC connect_timeout=5"

	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(conn)

	app.DB = db.PostgresConn{DB: conn}

	os.Exit(m.Run())
}
