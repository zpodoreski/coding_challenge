package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/coding_challenge/functions/add"
	"github.com/coding_challenge/functions/divide"
	"github.com/coding_challenge/functions/multiply"
	"github.com/coding_challenge/functions/subtract"
)

type Server struct{}

// Setting server
func New() *Server {
	s := &Server{}
	return s
}

func (s *Server) Serve(ctx context.Context) error {
	mux := http.NewServeMux()

	// List of all url paths on which server is listening and their methods
	mux.Handle("/add", http.HandlerFunc(add.Add))
	mux.Handle("/subtract", http.HandlerFunc(subtract.Subtract))
	mux.Handle("/multiply", http.HandlerFunc(multiply.Multiply))
	mux.Handle("/divide", http.HandlerFunc(divide.Divide))

	srv := &http.Server{
		Handler: mux,
	}

	// Starting HTTP server to listen od declared paths
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Printf("Server started")

	// Listening on channel for server shutdown signal
	<-ctx.Done()
	log.Printf("Server stopped")

	// Starting context with timeout for graceful shutdown
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer func() {
		cancel()
	}()

	// Shutting down server
	err := srv.Shutdown(ctxShutDown)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server Shutdown Failed: %+s", err)
	}

	return nil
}
