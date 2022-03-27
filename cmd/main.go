package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/coding_challenge/cmd/server"
)

func main() {
	// Creating channel for service termination
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Setting service working time
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Minute)

	// Listening for service stop signal
	go func() {
		oscall := <-c
		log.Printf("Server interupted:%+v", oscall)
		cancel()
	}()

	// Initializing server
	s := server.New()

	// Starting listening for HTTP requests
	if err := s.Serve(ctx); err != nil {
		log.Printf("Failed to serve:+%v\n", err)
	}
}
