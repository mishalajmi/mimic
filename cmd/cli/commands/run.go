package commands

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mishalalajmi/mimic/internal/project"
	"github.com/mishalalajmi/mimic/internal/server"
	"github.com/spf13/cobra"
)

var (
	projectPath string
	port        int
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "Start the mimic mock server",
	RunE:  run,
}

func init() {
	rootCommand.AddCommand(runCommand)

	runCommand.Flags().StringVar(&projectPath, "path", ".", "Path to project and mock definitions")
	runCommand.Flags().IntVar(&port, "port", 4010, "Port to run the mimic mock server on")
}

func run(cmd *cobra.Command, args []string) error {
	ctx, stop := signal.NotifyContext(
		cmd.Context(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	p, err := project.Load(projectPath)
	if err != nil {
		return fmt.Errorf("loading project: %w", err)
	}

	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Starting Mimic project %q on %s...\n", p.Name, addr)

	s := server.New(addr, p)

	return runServer(ctx, s)
}

func runServer(ctx context.Context, s *server.Server) error {
	errCh := make(chan error, 1)

	go func() {
		errCh <- s.Start()
	}()

	select {
	case err := <-errCh:
		if errors.Is(err, server.ErrServerClosed) {
			return nil
		}

		return err

	case <-ctx.Done():
		fmt.Print("\nReceived shutdown signal, shutting down...\n")
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			5*time.Second,
		)
		defer cancel()

		if err := s.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shutting down server: %w", err)
		}

		return nil
	}
}
