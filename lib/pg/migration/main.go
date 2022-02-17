package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var dsn string

func init() {
	_ = godotenv.Load(".env")
	dsn = os.Getenv("PG_URL")
	fmt.Println("dsn", dsn)
}

func main() {
	fmt.Println("migrate start")
	m, err := migrate.New(
		"file://lib/pg/migration/migrations",
		dsn)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("migrate finish")
}
