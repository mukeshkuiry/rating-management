// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: management.proto

package grpc_mngm

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

const (
	RatingManagement_GetScores_FullMethodName     = "/grpc_mngm.RatingManagement/GetScores"
	RatingManagement_AssignProblem_FullMethodName = "/grpc_mngm.RatingManagement/AssignProblem"
)

// RatingManagementClient is the client API for RatingManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RatingManagementClient interface {
	GetScores(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreResponse, error)
	AssignProblem(ctx context.Context, in *AssignProblemRequest, opts ...grpc.CallOption) (*AssignProblemResponse, error)
}

type ratingManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewRatingManagementClient(cc grpc.ClientConnInterface) RatingManagementClient {
	return &ratingManagementClient{cc}
}

func (c *ratingManagementClient) GetScores(ctx context.Context, in *GetScoreRequest, opts ...grpc.CallOption) (*GetScoreResponse, error) {
	out := new(GetScoreResponse)
	err := c.cc.Invoke(ctx, RatingManagement_GetScores_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ratingManagementClient) AssignProblem(ctx context.Context, in *AssignProblemRequest, opts ...grpc.CallOption) (*AssignProblemResponse, error) {
	out := new(AssignProblemResponse)
	err := c.cc.Invoke(ctx, RatingManagement_AssignProblem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RatingManagementServer is the server API for RatingManagement service.
// All implementations must embed UnimplementedRatingManagementServer
// for forward compatibility
type RatingManagementServer interface {
	GetScores(context.Context, *GetScoreRequest) (*GetScoreResponse, error)
	AssignProblem(context.Context, *AssignProblemRequest) (*AssignProblemResponse, error)
	mustEmbedUnimplementedRatingManagementServer()
}

// UnimplementedRatingManagementServer must be embedded to have forward compatible implementations.
type UnimplementedRatingManagementServer struct {
}

func (UnimplementedRatingManagementServer) GetScores(context.Context, *GetScoreRequest) (*GetScoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetScores not implemented")
}
func (UnimplementedRatingManagementServer) AssignProblem(context.Context, *AssignProblemRequest) (*AssignProblemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AssignProblem not implemented")
}
func (UnimplementedRatingManagementServer) mustEmbedUnimplementedRatingManagementServer() {}

// UnsafeRatingManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RatingManagementServer will
// result in compilation errors.
type UnsafeRatingManagementServer interface {
	mustEmbedUnimplementedRatingManagementServer()
}

func RegisterRatingManagementServer(s grpc.ServiceRegistrar, srv RatingManagementServer) {
	s.RegisterService(&RatingManagement_ServiceDesc, srv)
}

func _RatingManagement_GetScores_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetScoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingManagementServer).GetScores(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingManagement_GetScores_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingManagementServer).GetScores(ctx, req.(*GetScoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RatingManagement_AssignProblem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AssignProblemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RatingManagementServer).AssignProblem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RatingManagement_AssignProblem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RatingManagementServer).AssignProblem(ctx, req.(*AssignProblemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RatingManagement_ServiceDesc is the grpc.ServiceDesc for RatingManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RatingManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_mngm.RatingManagement",
	HandlerType: (*RatingManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetScores",
			Handler:    _RatingManagement_GetScores_Handler,
		},
		{
			MethodName: "AssignProblem",
			Handler:    _RatingManagement_AssignProblem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "management.proto",
}
