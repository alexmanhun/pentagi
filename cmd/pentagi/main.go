package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/vxcontrol/pentagi/pkg/version"
)

func main() {
	// Print version information on startup
	v := version.Get()
	fmt.Printf("Starting pentagi %s (commit: %s, built: %s)\n",
		v.Version, v.GitCommit, v.BuildDate)

	// Create a root context that is cancelled on OS interrupt signals
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Handle graceful shutdown on SIGINT and SIGTERM
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		fmt.Printf("\nReceived signal %s, shutting down...\n", sig)
		cancel()
	}()

	if err := run(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

// run is the main entry point for the application logic.
// It initialises all subsystems and blocks until the context is cancelled.
func run(ctx context.Context) error {
	fmt.Println("pentagi is running. Press Ctrl+C to stop.")

	// Block until context is cancelled (e.g. by OS signal)
	<-ctx.Done()

	fmt.Println("Shutdown complete.")
	return nil
}
