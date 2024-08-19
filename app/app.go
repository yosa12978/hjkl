package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yosa12978/hjkl/config"
	"github.com/yosa12978/hjkl/data"
	"github.com/yosa12978/hjkl/logging"
	"github.com/yosa12978/hjkl/router"
)

func Run() error {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer cancel()

	// Initializing database connections
	data.Postgres(ctx)
	data.Redis(ctx)

	// Initializing logger
	logger := logging.NewJsonLogger(os.Stdout)

	handler := router.New(
		router.WithLogger(logger),
	)

	cfg := config.Read()
	addr := fmt.Sprintf("%s:%d", cfg.App.Host, cfg.App.Port)
	server := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("server started", "address", addr)
		if err := server.ListenAndServe(); err != nil {
			errCh <- err
		}
		close(errCh)
	}()

	var err error
	select {
	case err = <-errCh:
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(
			context.Background(),
			10*time.Second,
		)
		defer cancel()
		err = server.Shutdown(timeout)
	}

	return err
}
