package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	dbURL := os.Getenv("DATABASE_URL")

	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	store := NewLessonStore(db)
	handler := &LessonHandler{store: store}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("GET /lessons", handler.List)
	mux.HandleFunc("GET /lessons/{id}", handler.Get)
	mux.HandleFunc("POST /lessons", handler.Create)
	mux.HandleFunc("PATCH /lessons/{id}", handler.SetStatus)
	mux.HandleFunc("DELETE /lessons/{id}", handler.Delete)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("server started on :%s", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}