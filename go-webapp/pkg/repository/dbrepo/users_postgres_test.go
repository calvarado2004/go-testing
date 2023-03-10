package dbrepo

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
	"log"
	"os"
	"testing"
)

//integration tests for Postgres dbrepo

var (
	host     = "localhost"
	user     = "postgres"
	password = "postgres"
	dbname   = "users_test"
	port     = "5435"
	dsn      = "host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=UTC connect_timeout=5"
)

var resource *dockertest.Resource
var pool *dockertest.Pool
var testDB *sql.DB

// TestMain is the entry point for all tests
func TestMain(m *testing.M) {

	// connect to docker; fail if docker is not running
	p, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %v", err)
	}

	pool = p

	// set up docker options, specifying image, port, and env vars
	dockerOpts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        "14.5",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbname,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{
					HostIP:   "0.0.0.0",
					HostPort: port,
				},
			},
		},
	}

	// get docker image
	resource, err = pool.RunWithOptions(&dockerOpts)
	if err != nil {
		log.Fatalf("Could not start resource: %v", err)
	}

	// start the docker container and wait for it to be ready
	if err = pool.Retry(func() error {
		var err error
		testDB, err = sql.Open("pgx", fmt.Sprintf(dsn, host, port, user, password, dbname))
		if err != nil {
			log.Println("Could not connect to postgres yet")
			return err
		}
		return testDB.Ping()
	}); err != nil {
		_ = pool.Purge(resource)
		log.Fatalf("Could not connect to postgres at all: %s", err)
	}

	// populate the test database

	// run the tests

	code := m.Run()

	// clean up after tests

	os.Exit(code)
}
