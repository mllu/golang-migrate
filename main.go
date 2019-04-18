package main

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello World")
	// m, err := migrate.New(
	// 	"file:///migrations",
	// 	"postgres://air:air@127.0.0.1:5432/database?sslmode=disable")

	db, err := sql.Open("postgres", "postgres://air:air@127.0.0.1:5432/database?sslmode=disable")
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file:///Users/mllu/go/src/github.robot.car/meng-lin-lu/golang-migrate/migrations",
		"postgres", driver)
	if err != nil {
		fmt.Println(err)
	}
	m.Steps(1)

	return
}
