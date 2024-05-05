package app

import (
	grpcapp "homework-backend/internal/app/grpc"
	"homework-backend/internal/services/auth"
	"homework-backend/internal/storage/pgx"
	"log/slog"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort uint,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := pgx.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
