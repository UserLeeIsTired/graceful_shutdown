package database

import (
	"context"
	"database/sql"
)

type Database struct {
	db      *sql.DB
	Context context.Context
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
