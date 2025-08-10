package main

import (
	"ani-aniapi/internal/app"
	"ani-aniapi/internal/config"
	"ani-aniapi/pkg/logger"
	"google.golang.org/grpc"
	"log/slog"
)

func main() {
	cfg := config.MustLoad()
	log := logger.SetupLogger(cfg.Env)

	contentGrpc := connectToGrpc(cfg.Services.ContentServiceAddress, log)
	defer contentGrpc.Close()

	app.CreateHttpServer(log, contentGrpc, cfg.Rest.Port)
}

func connectToGrpc(addr string, log *slog.Logger) *grpc.ClientConn {
	grpcConn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Error("Failed to connect to gRPC server: %v", err)
	}

	return grpcConn
}
