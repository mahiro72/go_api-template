package http

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi/v5"
	"github.com/mahiro72/go_api-template/pkg/infrastructure/db/postgresql"
	"github.com/mahiro72/go_api-template/pkg/web/http/controller"
)

func InitRouter() {
	db, err := postgresql.NewDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	srv := &http.Server{Addr: "0.0.0.0:8080", Handler: newHandlers(db)}

	// graceful shutdown
	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}

func newHandlers(db *sql.DB) http.Handler {
	router := chi.NewRouter()

	router.Group(func(ur chi.Router) {
		ur.Get("/users/{userID}", controller.GetUser(db))
		ur.Post("/users", controller.CreateUser(db))
		ur.Put("/users/{userID}", controller.UpdateUser(db))
		ur.Delete("/users/{userID}", controller.DeleteUser(db))
	})
	return router
}
