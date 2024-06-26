package cmd

import (
	"homework-backend/internal/app"
	grpcweb "homework-backend/internal/app/grpc-web"
	"homework-backend/internal/config"
	"log/slog"
	"os"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func Run() {
	// TODO: инициализировать объект конфига
	cfg := config.MustLoad()

	// TODO: инициализировать логгер
	log := setupLogger(cfg.Env)
	// TODO: инициализировать приложение (app)
	var application *app.App
	if cfg.StoragePathForRead == "" {
		application = app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.StoragePath, cfg.CachePath, cfg.QueuePath, cfg.TokenTTL)
	}
	application = app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.StoragePathForRead, cfg.CachePath, cfg.QueuePath, cfg.TokenTTL)

	// TODO: запустить gRPC-сервер приложения
	go application.GRPCServer.MustRun()

	grpcweb.Run(cfg.GRPC.Port, cfg.GRPC.WebPort, cfg.GRPC.WsPort, cfg.QueuePath, log)

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
