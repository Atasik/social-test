package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"social/internal/config"
	delivery "social/internal/http/v1"
	"social/internal/repository"
	"social/internal/server"
	"social/internal/service"
	"syscall"
	"time"

	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

const (
	timeout        = 5 * time.Second
	connectRetries = 5
	retryInterval  = 10 * time.Second
)

func Run(configDir string) {
	zapLogger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("Error occurred while loading zapLogger: %s\n", err.Error())
		return
	}
	defer zapLogger.Sync() //nolint:errcheck
	logger := zapLogger.Sugar()

	cfg, err := config.InitConfig(configDir)
	if err != nil {
		logger.Errorf("Error occurred while loading config: %s\n", err.Error())
		return
	}

	repos := repository.NewRepository(
		repository.NewPostMemoryRepo(),
	)

	services := service.NewService(
		service.NewPostService(repos),
	)

	validate := validator.New()

	handler := delivery.NewHandler(services, logger, validate)

	mux := handler.InitRoutes()

	srv := server.NewServer(cfg, mux)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("panic occurred: %s\n", err)
			}
		}()
		if err := srv.Run(); err != nil {
			logger.Errorf("Failed to start server: %s\n", err.Error())
		}
	}()

	logger.Info("Application is running")

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	logger.Info("Application is shutting down")

	if err := srv.Shutdown(ctx); err != nil {
		logger.Error(err.Error())
	}
}
