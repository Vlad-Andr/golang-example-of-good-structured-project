package main

import (
	"best-structure-example/internal/config"
	"best-structure-example/internal/handler"
	"best-structure-example/internal/kafka"
	"best-structure-example/internal/repository"
	"best-structure-example/internal/server"
	"best-structure-example/internal/service"
	"context"
	"go.uber.org/fx"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Configure logger
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	// Start application with fx
	app := fx.New(
		// Provide dependencies
		fx.Provide(
			// Supply the logger
			func() *log.Logger {
				return logger
			},
			// Load configurations
			config.NewConfig,
			config.ProvideServerConfig,
			config.ProvideKafkaConfig,
			// Create repository
			repository.NewRepository,
			// Create Kafka producer
			kafka.NewProducer,
			// Create Kafka consumer
			kafka.NewConsumer,
			// Create service
			service.NewService,
			// Create HTTP handler
			handler.NewHandler,
			handler.ProvideHttpHandler,
			// Create HTTP server
			server.NewServer,
		),
		// Register lifecycle hooks
		fx.Invoke(registerHooks),
	)

	// Start the application
	app.Run()
}

// registerHooks sets up the application lifecycle hooks
func registerHooks(
	lc fx.Lifecycle,
	logger *log.Logger,
	server *server.Server,
	consumer kafka.Consumer,
	service service.Service,
) {
	// Create context for shutdown
	_, cancel := context.WithCancel(context.Background())

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Start Kafka consumer
			logger.Println("Starting Kafka consumer...")
			if err := consumer.Start(ctx, service.ProcessMessage); err != nil {
				return err
			}

			// Start HTTP server
			logger.Printf("Starting HTTP server on %s", server.Addr)
			go func() {
				if err := server.Start(); err != nil {
					logger.Printf("Server error: %v", err)
				}
			}()

			// Handle shutdown signals
			go func() {
				sigCh := make(chan os.Signal, 1)
				signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
				sig := <-sigCh
				logger.Printf("Received signal: %s", sig)
				cancel()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Println("Shutting down application...")

			// Shutdown HTTP server
			if err := server.Shutdown(ctx); err != nil {
				logger.Printf("Server shutdown error: %v", err)
			}

			// Stop Kafka consumer
			if err := consumer.Stop(); err != nil {
				logger.Printf("Consumer shutdown error: %v", err)
			}

			logger.Println("Application stopped")
			return nil
		},
	})
}
