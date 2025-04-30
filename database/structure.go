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
	UserId   string    `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Comments []Comment `json:"comments,omitempty"`
}

type Comment struct {
	CommentId string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	Content   string `json:"content,omitempty"`
	UserId    string `json:"user_id,omitempty"`
}
