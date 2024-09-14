package app

import (
	"log/slog"
	"time"

	"google.golang.org/grpc"
)

type App struct {
	GRPCSrv *grpc.Server
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
