package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	
	app := application {
		logger: logger,
	}
	
	srv := http.Server {
		Addr: ":4000",
		ReadTimeout: 5 *time.Second,
		WriteTimeout: 10 * time.Second,
		Handler: app.routes(),
	}
	
	logger.Info("starting server", slog.String("addr", ":4000"))
	if err := srv.ListenAndServe(); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}