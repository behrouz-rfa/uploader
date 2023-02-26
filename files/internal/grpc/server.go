package grpc

import (
	"uploader/files/filespb"
	"uploader/files/internal/application"
)

type server struct {
	app application.App
	filespb.UnimplementedFilesServiceServer
}
