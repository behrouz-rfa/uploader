package monolith

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/minio/minio-go"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"gorm.io/gorm"

	"uploader/internal/config"
	"uploader/internal/waiter"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *gorm.DB
	Logger() zerolog.Logger
	Mux() *chi.Mux
	RPC() *grpc.Server
	Waiter() waiter.Waiter
	Minio() *minio.Client
}

type Module interface {
	Startup(context.Context, Monolith) error
}
