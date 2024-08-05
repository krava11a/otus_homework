package app

import (
	"context"
	grpcapp "homework-backend/internal/app/grpc"
	"homework-backend/internal/models"
	"homework-backend/internal/services/auth"
	"homework-backend/internal/services/post"
	"homework-backend/internal/storage/pgx"
	"homework-backend/internal/storage/rabbit"
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
	queuePath string,
	tokenTTL time.Duration,
	app models.App,
) *App {
	storage, err := pgx.New(storagePath, app)
	if err != nil {
		panic(err)
	}

	storageForRead, err := pgx.New(storagePathForRead, app)
	if err != nil {
		panic(err)
	}

	cache, err := redis.New(cachePath)
	if err != nil {
		log.Log(context.Background(), slog.LevelError, err.Error())
	}

	rqueue, err := rabbit.New(queuePath)
	if err != nil {
		log.Log(context.Background(), slog.LevelError, err.Error())
	}

	authService := auth.New(log, storage, storageForRead, storage, tokenTTL)
	postService := post.New(log, storage, storageForRead, cache, rqueue)

	grpcApp := grpcapp.New(log, authService, postService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
