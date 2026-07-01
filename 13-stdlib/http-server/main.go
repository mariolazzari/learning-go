package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"time"
)

// handler
func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ciao Mario"))
}

// handler
func handleGreat(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	fmt.Fprintf(w, "Ciao %s\n", name)
}

// middleware
func RequestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		dur := time.Since(start)
		slog.Info("request time",
			"path", r.URL.Path,
			"duration", dur)
	})
}

func main() {
	// router
	mux := http.NewServeMux()
	// handlers
	mux.HandleFunc("GET /hello", handleHello)
	mux.HandleFunc("GET /hello/{name}", handleGreat)
	// middlewares
	wrappedMux := RequestTimer(mux)

	// server
	srv := http.Server{
		Addr:         ":8081",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  90 * time.Second,
		Handler:      wrappedMux,
	}

	// start server
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("server: %v", err)
	}
}
