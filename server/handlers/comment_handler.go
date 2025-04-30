package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/UserLeeIsTired/graceful_shutdown/database"
	"github.com/go-chi/chi/v5"
)

func CommentRoute(r chi.Router, db *database.Database) {

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		GetAllCommentsWithUser(w, r, db)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		CreateComment(w, r, db)
	})

	r.Get("/{commentId}", func(w http.ResponseWriter, r *http.Request) {
		GetCommentByCommentId(w, r, db)
	})
}

func CreateComment(w http.ResponseWriter, r *http.Request, db *database.Database) {
	var comment database.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	commentId, err := db.CreateComment(
		comment.UserId,
		comment.Title,
		comment.Content,
	)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteHeader(w, http.StatusCreated, commentId)
}

func GetAllCommentsWithUser(w http.ResponseWriter, r *http.Request, db *database.Database) {
	comments, err := db.GetAllCommentsWithUser()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteHeader(w, http.StatusOK, comments)
}

func GetCommentByCommentId(w http.ResponseWriter, r *http.Request, db *database.Database) {
	commentId := chi.URLParam(r, "commentId")
	comment, err := db.GetCommentByCommentId(commentId)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	WriteHeader(w, http.StatusOK, comment)
}
