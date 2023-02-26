// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: gatespb/api.proto

package gatespb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// GatesServiceClient is the client API for GatesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatesServiceClient interface {
	Files(ctx context.Context, in *FilesRequest, opts ...grpc.CallOption) (*FilesResponse, error)
	FromBucket(ctx context.Context, in *FromBucketRequest, opts ...grpc.CallOption) (*FromBucketResponse, error)
	Transcode(ctx context.Context, in *TranscodeRequest, opts ...grpc.CallOption) (*TranscodeResponse, error)
}

type gatesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatesServiceClient(cc grpc.ClientConnInterface) GatesServiceClient {
	return &gatesServiceClient{cc}
}

func (c *gatesServiceClient) Files(ctx context.Context, in *FilesRequest, opts ...grpc.CallOption) (*FilesResponse, error) {
	out := new(FilesResponse)
	err := c.cc.Invoke(ctx, "/gatespb.GatesService/Files", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatesServiceClient) FromBucket(ctx context.Context, in *FromBucketRequest, opts ...grpc.CallOption) (*FromBucketResponse, error) {
	out := new(FromBucketResponse)
	err := c.cc.Invoke(ctx, "/gatespb.GatesService/FromBucket", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatesServiceClient) Transcode(ctx context.Context, in *TranscodeRequest, opts ...grpc.CallOption) (*TranscodeResponse, error) {
	out := new(TranscodeResponse)
	err := c.cc.Invoke(ctx, "/gatespb.GatesService/Transcode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatesServiceServer is the server API for GatesService service.
// All implementations must embed UnimplementedGatesServiceServer
// for forward compatibility
type GatesServiceServer interface {
	Files(context.Context, *FilesRequest) (*FilesResponse, error)
	FromBucket(context.Context, *FromBucketRequest) (*FromBucketResponse, error)
	Transcode(context.Context, *TranscodeRequest) (*TranscodeResponse, error)
	mustEmbedUnimplementedGatesServiceServer()
}

// UnimplementedGatesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatesServiceServer struct {
}

func (UnimplementedGatesServiceServer) Files(context.Context, *FilesRequest) (*FilesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Files not implemented")
}
func (UnimplementedGatesServiceServer) FromBucket(context.Context, *FromBucketRequest) (*FromBucketResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FromBucket not implemented")
}
func (UnimplementedGatesServiceServer) Transcode(context.Context, *TranscodeRequest) (*TranscodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transcode not implemented")
}
func (UnimplementedGatesServiceServer) mustEmbedUnimplementedGatesServiceServer() {}

// UnsafeGatesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatesServiceServer will
// result in compilation errors.
type UnsafeGatesServiceServer interface {
	mustEmbedUnimplementedGatesServiceServer()
}

func RegisterGatesServiceServer(s grpc.ServiceRegistrar, srv GatesServiceServer) {
	s.RegisterService(&GatesService_ServiceDesc, srv)
}

func _GatesService_Files_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatesServiceServer).Files(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gatespb.GatesService/Files",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatesServiceServer).Files(ctx, req.(*FilesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatesService_FromBucket_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FromBucketRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatesServiceServer).FromBucket(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gatespb.GatesService/FromBucket",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatesServiceServer).FromBucket(ctx, req.(*FromBucketRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatesService_Transcode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TranscodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatesServiceServer).Transcode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gatespb.GatesService/Transcode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatesServiceServer).Transcode(ctx, req.(*TranscodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GatesService_ServiceDesc is the grpc.ServiceDesc for GatesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gatespb.GatesService",
	HandlerType: (*GatesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Files",
			Handler:    _GatesService_Files_Handler,
		},
		{
			MethodName: "FromBucket",
			Handler:    _GatesService_FromBucket_Handler,
		},
		{
			MethodName: "Transcode",
			Handler:    _GatesService_Transcode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gatespb/api.proto",
}