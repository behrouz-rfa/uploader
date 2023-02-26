package rest

import (
	"context"
	"github.com/go-chi/chi/v5"
)

func RegisterGateway(ctx context.Context, mux *chi.Mux, h Handler) error {
	const apiRoot = "/api/files"

	mux.Post("/api/files", h.Upload)

	//gateway := runtime.NewServeMux()
	//err := filespb.RegisterFilesServiceHandlerFromEndpoint(ctx, gateway, grpcAddr, []grpc.DialOption{
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//})
	//if err != nil {
	//	return err
	//}
	//
	//// mount the GRPC gateway
	//mux.Mount(apiRoot, h.Upload)

	return nil
}
