package app

import (
	"context"
	grpcapp "homework-backend/internal/app/grpc"
	"homework-backend/internal/services/auth"
	"homework-backend/internal/services/post"
	"homework-backend/internal/storage/pgx"
	"homework-backend/internal/storage/redis"
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
	storagePathForRead string,
	cachePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := pgx.New(storagePath)
	if err != nil {
		panic(err)
	}

	storageForRead, err := pgx.New(storagePathForRead)
	if err != nil {
		panic(err)
	}

	cache, err := redis.New(cachePath)
	if err != nil {
		log.Log(context.Background(), slog.LevelError, err.Error())
	}

	authService := auth.New(log, storage, storageForRead, storage, tokenTTL)
	postService := post.New(log, storage, storageForRead, cache)

	grpcApp := grpcapp.New(log, authService, postService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
