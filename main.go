package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/UserLeeIsTired/graceful_shutdown/database"
	"github.com/UserLeeIsTired/graceful_shutdown/server"
)

func main() {
	// Create a context with cancel
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Listen for OS signals to trigger shutdown
	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		fmt.Println("Shutdown signal received. Initiating graceful shutdown...")
		cancel() // Cancel the context
	}()

	// Initialize the database
	d, err := database.NewDatabase(ctx)
	if err != nil {
		panic(err)
	}
	defer d.Close()

	// Start the server
	go server.NewServer(ctx, d, ":8080")

	// Simulate some database activity
	go func() {
		for range ctx.Done() {
			fmt.Println("Context canceled, stopping database operations...")
			return
		}
	}()

	// Wait for the context to be canceled
	<-ctx.Done()

	// Graceful shutdown with a timeout
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	fmt.Println("Waiting for ongoing operations to complete...")
	select {
	case <-shutdownCtx.Done():
		if shutdownCtx.Err() == context.DeadlineExceeded {
			fmt.Println("Shutdown timeout reached. Forcing shutdown.")
		} else {
			fmt.Println("All operations completed. Shutting down.")
		}
	}

	fmt.Println("Application has shut down.")
}
