package app

import (
	grpcapp "homework-dialogs/internal/app/grpc"
	"homework-dialogs/internal/services/auth"
	"homework-dialogs/internal/services/dialog"
	"homework-dialogs/internal/storage/pgx"

	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort uint,
	storagePath string,
	authPah string,

) *App {
	storage, err := pgx.New(storagePath)
	if err != nil {
		panic(err)
	}

	ars := auth.New(authPah)

	dialogService := dialog.New(log, storage, storage)

	grpcApp := grpcapp.New(log, dialogService, grpcPort, ars)

	return &App{
		GRPCServer: grpcApp,
	}
}
