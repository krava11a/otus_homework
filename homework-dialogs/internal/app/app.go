package app

import (
	"fmt"
	grpcapp "homework-dialogs/internal/app/grpc"
	"homework-dialogs/internal/services/auth"
	"homework-dialogs/internal/services/dialog"
	"homework-dialogs/internal/storage/pgx"
	"homework-dialogs/internal/storage/tarantool"

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
	_, err := pgx.New(storagePath)
	if err != nil {
		fmt.Printf("pgx.new error %s", err)
	}

	tara, err := tarantool.New(storagePath)
	if err != nil {
		fmt.Printf("tarantool.new error %s", err)
	}

	ars := auth.New(authPah)

	dialogService := dialog.New(log, tara, tara)

	grpcApp := grpcapp.New(log, dialogService, grpcPort, ars)

	return &App{
		GRPCServer: grpcApp,
	}
}
