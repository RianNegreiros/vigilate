package database

import (
	"database/sql"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func waitForDatabase(db *sql.DB) {
	const maxAttempts = 10
	for i := 0; i < maxAttempts; i++ {
		err := db.Ping()
		if err == nil {
			log.Println("Successfully connected to database")
			return
		}
		log.Println("Failed to ping database:", err)
		time.Sleep(time.Second * 5)
	}
	log.Fatalf("Failed to connect to database after %d attempts, exiting.", maxAttempts)
}

func NewDatabase() (*Database, error) {
	databaseURL := os.Getenv("DATABASE_URL")

	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)

	if os.Getenv("APP_ENV") == "docker" {
		waitForDatabase(db)
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}

func (d *Database) Ping() error {
	return d.db.Ping()
}
