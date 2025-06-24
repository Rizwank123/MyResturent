package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rizwank123/myResturent/internal/dependency"
	"github.com/rizwank123/myResturent/internal/http/swagger"
	"github.com/rizwank123/myResturent/internal/pkg/config"
)

func main() {
	option := loadOption()
	// Intilize config
	cfg, err := dependency.NewConfig(option)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	// Intilize database
	db, err := dependency.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Intilize ResturnetApi
	api, err := dependency.NewResturnetApi(cfg, db)
	if err != nil {
		log.Fatalf("Failed to create ResturnetApi: %v", err)
	}

	e := echo.New()
	e.HideBanner = true

	// Set up the middleware
	api.SetupMiddleware(e)

	// Set up the swagger documentation
	swagger.SetupSwagger(cfg, e)

	// Set up the routes
	api.SetupRoutes(e)

	// Start the server in a goroutine to handle graceful shutdown
	go func() {
		e.Logger.Info(e.Start(fmt.Sprintf("0.0.0.0:%d", cfg.AppPort)))
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

	log.Println("Server gracefully stopped")

}

func loadOption() config.Options {
	cfgSource := os.Getenv(config.SourceKey)
	if cfgSource == "" {
		cfgSource = config.SourceEnv
	}
	cfgOptions := config.Options{
		ConfigSource: cfgSource,
	}
	switch cfgSource {
	case config.SourceEnv:
		cfgOptions.ConfigFile = ".env"
	}
	return cfgOptions
}
