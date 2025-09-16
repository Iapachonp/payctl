/*
Copyright © 2025 sir-pachis
*/
package main

import (
	"fmt"
	"payctl/cmd"
	"payctl/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fmt.Println("Setting up payctl...")
	db := database.Open()
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		fmt.Printf("failed to start up sqlite driver %v" , err)
	}
	m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        "sqlite3", driver)
	if err != nil {
		fmt.Printf("Failed to init db driver %v \n" , err)
	}
	if err = m.Up(); (err != nil && err != migrate.ErrNoChange) {
		fmt.Printf("Failed to init db migrations %v \n" , err)
	} // or m.Steps(2) if you want to explicitly set the number of migrations to run
	fmt.Println("payctl started successfully")
	cmd.Execute()
}
