package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"transactions-service/config"
	"transactions-service/db"
	_ "transactions-service/swagger"
)

// @title			Transactions Service
// @version		v1
// @description	Exposed APIs for accounts and transactions related operations
// @contact.name	Jetin Biraj
// @contact.email	birajdarjk1106@gmail.com
func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	if err := config.Set(); err != nil {
		return err
	}

	httpServer, err := buildHTTPServer()
	if err != nil {
		return err
	}

	defer func() {
		if err = closePostgres(); err != nil {
			log.Printf("error closing postgres: %v", err)
		}
	}()

	serverErr := make(chan error, 1)
	go func() {
		log.Printf("Application server staring on port %v", httpServer.Addr)
		serverErr <- httpServer.ListenAndServe()
	}()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	select {
	case err = <-serverErr:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil

	case <-ctx.Done():
		log.Printf("Shutdown signal received, stopping server...")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err = httpServer.Shutdown(shutdownCtx); err != nil {
			return err
		}

		err = <-serverErr
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			return err
		}

		log.Printf("Server shutdown completed")
		return nil
	}
}

func closePostgres() error {
	if db.Pg == nil {
		return nil
	}
	return db.Pg.Close()
}
