package grpc

import (
	"context"
	"uploader/gates/gatespb"
	"uploader/gates/internal/application"
)

type server struct {
	app application.App
	gatespb.UnimplementedGatesServiceServer
}

func (s server) Files(ctx context.Context, request *gatespb.FilesRequest) (*gatespb.FilesResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) FromBucket(ctx context.Context, request *gatespb.FromBucketRequest) (*gatespb.FromBucketResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) Transcode(ctx context.Context, request *gatespb.TranscodeRequest) (*gatespb.TranscodeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (s server) mustEmbedUnimplementedGatesServiceServer() {
	//TODO implement me
	panic("implement me")
}
