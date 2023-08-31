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
)

type application struct {
	logger    *log.Logger
	companion *companionResponses
	apiLogger *apiLogManager
	config    config
}

type config struct {
	env  string
	port string
}

// newApp function returns a new instance of Mock Wallet application
func newApp(config config) *application {
	return &application{
		logger:    log.New(os.Stdout, "", log.LstdFlags|log.Llongfile),
		config:    config,
		companion: defaultCompanionResponses(),
	}
}

func (app *application) run() error {
	server := &http.Server{
		Addr:         app.config.port,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownErrChan := make(chan error)

	go func() {
		exitChan := make(chan os.Signal, 1)

		signal.Notify(exitChan, syscall.SIGTERM, syscall.SIGINT)

		sig := <-exitChan
		app.logger.Printf("system signal %s was received\nthe application is stopping...", sig.String())

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		shutdownErrChan <- server.Shutdown(ctx)
	}()

	app.logger.Printf("application is running on port %s", app.config.port)
	err := server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownErrChan
	if err != nil {
		return err
	}

	return nil
}
