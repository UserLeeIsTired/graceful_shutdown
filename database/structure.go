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
	UserId int    `json:"id"`
	Name   string `json:"name"`
}

type Comment struct {
	CommentId int    `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	UserId    string `json:"user_id"`
}
