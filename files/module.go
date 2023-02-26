package files

import (
	"context"
	"uploader/files/internal/application"
	"uploader/files/internal/logging"
	"uploader/files/internal/postgres"
	"uploader/files/internal/rest"
	"uploader/internal/monolith"
	"uploader/internal/storage"
)

type Module struct {
}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) error {
	postgres.Migrate(mono.DB())
	fileInfo := postgres.NewFileInfoRepository(mono.DB())
	var app application.App
	app = application.NewApplication(fileInfo)
	app = logging.NewApplication(app, mono.Logger())

	fileStorage := storage.New(mono.Minio(), mono.Config().Minio.EndPoint, mono.Config().Minio.PublicEndPoint)

	handler := rest.New(app, fileStorage)
	//if err := grpc.RegisterServer(app, mono.RPC()); err != nil {
	//	return err
	//}
	if err := rest.RegisterGateway(ctx, mono.Mux(), handler); err != nil {
		return err
	}
	if err := rest.RegisterSwagger(mono.Mux()); err != nil {
		return err
	}

	return nil
}
