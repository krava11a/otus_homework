package cmd

import (
	"homework-backend/internal/app"
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
		application = app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.StoragePath, cfg.TokenTTL)
	}
	application = app.New(log, cfg.GRPC.Port, cfg.StoragePath, cfg.StoragePathForRead, cfg.TokenTTL)

	// TODO: запустить gRPC-сервер приложения
	application.GRPCServer.MustRun()

	// lis, err := net.Listen("tcp", "[::1]:9001")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }

	// grpcServer := grpc.NewServer()
	// service := &services.AuthorizationServer{}

	// proto.RegisterAuthorizationServiceServer(grpcServer, service)
	// err = grpcServer.Serve(lis)

	// if err != nil {
	// 	log.Fatalf("Error strating server: %v", err)
	// }
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
