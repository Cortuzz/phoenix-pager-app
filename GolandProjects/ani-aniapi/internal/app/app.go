package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log/slog"
	// TODO: "ani-aniapi/internal/gateway"
	// TODO: pbContent "ani-aniapi/pkg/proto/anipj/content"
	"strconv"
)

type App struct {
	log    *slog.Logger
	server *gin.Engine
	port   int
}

func CreateHttpServer(log *slog.Logger, contentGrpc *grpc.ClientConn, port int) *App {
	// TODO: contentClient := pb.NewContentServiceClient(contentGrpc)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// TODO: r.GET("/incidents/:id", handler.GetIncident)

	err := r.Run(":" + strconv.Itoa(port))

	if err != nil {
		return nil
	}

	return &App{
		log:    log,
		server: r,
		port:   port,
	}
}
