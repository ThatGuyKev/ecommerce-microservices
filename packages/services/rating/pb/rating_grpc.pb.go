// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: _proto/rating.proto

package pb

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

// RatingServiceClient is the client API for RatingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingServiceClient interface {
	GetRatingsAdmin(ctx context.Context, in *GetRatingsAdminReq, opts ...grpc.CallOption) (*GetRatingsAdminRes, error)
}

type ratingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingServiceClient(cc grpc.ClientConnInterface) RatingServiceClient {
	return &ratingServiceClient{cc}
}

func (c *ratingServiceClient) GetRatingsAdmin(ctx context.Context, in *GetRatingsAdminReq, opts ...grpc.CallOption) (*GetRatingsAdminRes, error) {
	out := new(GetRatingsAdminRes)
	err := c.cc.Invoke(ctx, "/rating.RatingService/GetRatingsAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingServiceServer is the server API for RatingService service.
// All implementations must embed UnimplementedRatingServiceServer
// for forward compatibility
type RatingServiceServer interface {
	GetRatingsAdmin(context.Context, *GetRatingsAdminReq) (*GetRatingsAdminRes, error)
	mustEmbedUnimplementedRatingServiceServer()
}

// UnimplementedRatingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRatingServiceServer struct {
}

func (UnimplementedRatingServiceServer) GetRatingsAdmin(context.Context, *GetRatingsAdminReq) (*GetRatingsAdminRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRatingsAdmin not implemented")
}
func (UnimplementedRatingServiceServer) mustEmbedUnimplementedRatingServiceServer() {}

// UnsafeRatingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingServiceServer will
// result in compilation errors.
type UnsafeRatingServiceServer interface {
	mustEmbedUnimplementedRatingServiceServer()
}

func RegisterRatingServiceServer(s grpc.ServiceRegistrar, srv RatingServiceServer) {
	s.RegisterService(&RatingService_ServiceDesc, srv)
}

func _RatingService_GetRatingsAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRatingsAdminReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingServiceServer).GetRatingsAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rating.RatingService/GetRatingsAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingServiceServer).GetRatingsAdmin(ctx, req.(*GetRatingsAdminReq))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingService_ServiceDesc is the grpc.ServiceDesc for RatingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rating.RatingService",
	HandlerType: (*RatingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRatingsAdmin",
			Handler:    _RatingService_GetRatingsAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "_proto/rating.proto",
}