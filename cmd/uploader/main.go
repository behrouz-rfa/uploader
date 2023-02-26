package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/minio/minio-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"uploader/files"
	"uploader/internal/config"
	"uploader/internal/logger"
	"uploader/internal/monolith"
	"uploader/internal/rpc"
	"uploader/internal/waiter"
	"uploader/internal/web"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() (err error) {
	var cfg config.AppConfig
	// parse config/env/...
	cfg, err = config.InitConfig()
	if err != nil {
		return err
	}

	m := app{cfg: cfg}

	// init infrastructure...
	sqlDB, err := sql.Open("pgx", cfg.PG.Conn)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	m.db = gormDB
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(sqlDB)
	m.logger = logger.New(logger.LogConfig{
		Environment: cfg.Environment,
		LogLevel:    logger.Level(cfg.LogLevel),
	})

	m.minioClient, err = initMinio(cfg.Minio)
	if err != nil {
		return
	}
	m.rpc = initRpc(cfg.Rpc)
	m.mux = initMux(cfg.Web)
	m.waiter = waiter.New(waiter.CatchSignals())

	// init modules
	m.modules = []monolith.Module{
		&files.Module{},
	}
	if err := m.startupModules(); err != nil {
		return err
	}
	//Mount general web resources
	m.mux.Mount("/", http.FileServer(http.FS(web.WebUI)))

	fmt.Println("started mallbots application")
	defer fmt.Println("stopped mallbots application")

	m.waiter.Add(
		m.waitForWeb,
		m.waitForRPC,
	)
	return m.waiter.Wait()
}

func initMinio(cfg config.MinioConfig) (*minio.Client, error) {
	client, err := minio.NewV2(cfg.EndPoint, cfg.AccessKey, cfg.SecretKey, cfg.Secure)
	return client, err
}
func initRpc(_ rpc.RpcConfig) *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}

func initMux(_ web.WebConfig) *chi.Mux {
	return chi.NewMux()
}
