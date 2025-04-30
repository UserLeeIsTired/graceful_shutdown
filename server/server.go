package server

import (
	"context"
	"net/http"

	"github.com/UserLeeIsTired/graceful_shutdown/database"
	"github.com/UserLeeIsTired/graceful_shutdown/server/handlers"
	"github.com/go-chi/chi/v5"
)

func ShutdownHandler(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-ctx.Done():
			// Respond immediately if the server is shutting down
			http.Error(w, "Server is shutting down", http.StatusServiceUnavailable)
			return
		default:
			// Continue to the next handler if the server is running
			next.ServeHTTP(w, r)
		}
	})
}

func NewServer(ctx context.Context, db *database.Database, addr string) {
	r := chi.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return ShutdownHandler(ctx, next)
	})

	r.Route("/users", func(r chi.Router) {
		handlers.UserRoutes(r, db)
	})

	r.Route("/comments", func(r chi.Router) {
		handlers.CommentRoute(r, db)
	})

	http.ListenAndServe(addr, r)
}
