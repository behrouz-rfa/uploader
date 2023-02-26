package grpc

import (
	"context"
	"google.golang.org/grpc"
	"uploader/files/filespb"
	"uploader/files/internal/application"
)

type server struct {
	app application.App
	filespb.UnimplementedFilesServiceServer
}

func (s server) Upload(ctx context.Context, request *filespb.UploadRequest) (*filespb.UploadResponse, error) {
	//TODO implement me
	panic("implement me")
}

func RegisterServer(app application.App, register grpc.ServiceRegistrar) error {
	filespb.RegisterFilesServiceServer(register, server{app: app})
	return nil
}

var _ filespb.FilesServiceServer = (*server)(nil)
