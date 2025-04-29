package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

func NewDatabase(ctx context.Context) (*Database, error) {

	err := godotenv.Load(".env")

	if err != nil {
		return nil, err
	}

	var (
		host     = os.Getenv("HOST")
		port     = os.Getenv("PORT")
		user     = os.Getenv("DATABASE_USER")
		password = os.Getenv("PASSWORD")
		dbname   = os.Getenv("DBNAME")
	)

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		dbname,
	)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return &Database{db: db, Context: ctx}, nil
}

func (d *Database) Close() {
	if d != nil {
		d.db.Close()
	}
}
