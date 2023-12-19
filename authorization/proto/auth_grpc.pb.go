// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: auth.proto

package auth_proto

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
	Authorization_GetId_FullMethodName                  = "/auth.Authorization/GetId"
	Authorization_GetIdsAndPaths_FullMethodName         = "/auth.Authorization/GetIdsAndPaths"
	Authorization_GetAuthorizationStatus_FullMethodName = "/auth.Authorization/GetAuthorizationStatus"
)

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationClient interface {
	GetId(ctx context.Context, in *FindIdRequest, opts ...grpc.CallOption) (*FindIdResponse, error)
	GetIdsAndPaths(ctx context.Context, in *NamesAndPathsListRequest, opts ...grpc.CallOption) (*NamesAndPathsResponse, error)
	GetAuthorizationStatus(ctx context.Context, in *AuthorizationCheckRequest, opts ...grpc.CallOption) (*AuthorizationCheckResponse, error)
}

type authorizationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationClient(cc grpc.ClientConnInterface) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) GetId(ctx context.Context, in *FindIdRequest, opts ...grpc.CallOption) (*FindIdResponse, error) {
	out := new(FindIdResponse)
	err := c.cc.Invoke(ctx, Authorization_GetId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) GetIdsAndPaths(ctx context.Context, in *NamesAndPathsListRequest, opts ...grpc.CallOption) (*NamesAndPathsResponse, error) {
	out := new(NamesAndPathsResponse)
	err := c.cc.Invoke(ctx, Authorization_GetIdsAndPaths_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) GetAuthorizationStatus(ctx context.Context, in *AuthorizationCheckRequest, opts ...grpc.CallOption) (*AuthorizationCheckResponse, error) {
	out := new(AuthorizationCheckResponse)
	err := c.cc.Invoke(ctx, Authorization_GetAuthorizationStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
// All implementations must embed UnimplementedAuthorizationServer
// for forward compatibility
type AuthorizationServer interface {
	GetId(context.Context, *FindIdRequest) (*FindIdResponse, error)
	GetIdsAndPaths(context.Context, *NamesAndPathsListRequest) (*NamesAndPathsResponse, error)
	GetAuthorizationStatus(context.Context, *AuthorizationCheckRequest) (*AuthorizationCheckResponse, error)
	mustEmbedUnimplementedAuthorizationServer()
}

// UnimplementedAuthorizationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (UnimplementedAuthorizationServer) GetId(context.Context, *FindIdRequest) (*FindIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetId not implemented")
}
func (UnimplementedAuthorizationServer) GetIdsAndPaths(context.Context, *NamesAndPathsListRequest) (*NamesAndPathsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetIdsAndPaths not implemented")
}
func (UnimplementedAuthorizationServer) GetAuthorizationStatus(context.Context, *AuthorizationCheckRequest) (*AuthorizationCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthorizationStatus not implemented")
}
func (UnimplementedAuthorizationServer) mustEmbedUnimplementedAuthorizationServer() {}

// UnsafeAuthorizationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthorizationServer will
// result in compilation errors.
type UnsafeAuthorizationServer interface {
	mustEmbedUnimplementedAuthorizationServer()
}

func RegisterAuthorizationServer(s grpc.ServiceRegistrar, srv AuthorizationServer) {
	s.RegisterService(&Authorization_ServiceDesc, srv)
}

func _Authorization_GetId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).GetId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authorization_GetId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).GetId(ctx, req.(*FindIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_GetIdsAndPaths_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NamesAndPathsListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).GetIdsAndPaths(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authorization_GetIdsAndPaths_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).GetIdsAndPaths(ctx, req.(*NamesAndPathsListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_GetAuthorizationStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizationCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).GetAuthorizationStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authorization_GetAuthorizationStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).GetAuthorizationStatus(ctx, req.(*AuthorizationCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Authorization_ServiceDesc is the grpc.ServiceDesc for Authorization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authorization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetId",
			Handler:    _Authorization_GetId_Handler,
		},
		{
			MethodName: "GetIdsAndPaths",
			Handler:    _Authorization_GetIdsAndPaths_Handler,
		},
		{
			MethodName: "GetAuthorizationStatus",
			Handler:    _Authorization_GetAuthorizationStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "auth.proto",
}