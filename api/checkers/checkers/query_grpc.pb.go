// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: checkers/checkers/query.proto

package checkers

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
	Query_Params_FullMethodName        = "/checkers.checkers.Query/Params"
	Query_SystemInfo_FullMethodName    = "/checkers.checkers.Query/SystemInfo"
	Query_StoredGame_FullMethodName    = "/checkers.checkers.Query/StoredGame"
	Query_StoredGameAll_FullMethodName = "/checkers.checkers.Query/StoredGameAll"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// Queries a SystemInfo by index.
	SystemInfo(ctx context.Context, in *QueryGetSystemInfoRequest, opts ...grpc.CallOption) (*QueryGetSystemInfoResponse, error)
	// Queries a list of StoredGame items.
	StoredGame(ctx context.Context, in *QueryGetStoredGameRequest, opts ...grpc.CallOption) (*QueryGetStoredGameResponse, error)
	StoredGameAll(ctx context.Context, in *QueryAllStoredGameRequest, opts ...grpc.CallOption) (*QueryAllStoredGameResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, Query_Params_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) SystemInfo(ctx context.Context, in *QueryGetSystemInfoRequest, opts ...grpc.CallOption) (*QueryGetSystemInfoResponse, error) {
	out := new(QueryGetSystemInfoResponse)
	err := c.cc.Invoke(ctx, Query_SystemInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StoredGame(ctx context.Context, in *QueryGetStoredGameRequest, opts ...grpc.CallOption) (*QueryGetStoredGameResponse, error) {
	out := new(QueryGetStoredGameResponse)
	err := c.cc.Invoke(ctx, Query_StoredGame_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) StoredGameAll(ctx context.Context, in *QueryAllStoredGameRequest, opts ...grpc.CallOption) (*QueryAllStoredGameResponse, error) {
	out := new(QueryAllStoredGameResponse)
	err := c.cc.Invoke(ctx, Query_StoredGameAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// Queries a SystemInfo by index.
	SystemInfo(context.Context, *QueryGetSystemInfoRequest) (*QueryGetSystemInfoResponse, error)
	// Queries a list of StoredGame items.
	StoredGame(context.Context, *QueryGetStoredGameRequest) (*QueryGetStoredGameResponse, error)
	StoredGameAll(context.Context, *QueryAllStoredGameRequest) (*QueryAllStoredGameResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (UnimplementedQueryServer) SystemInfo(context.Context, *QueryGetSystemInfoRequest) (*QueryGetSystemInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SystemInfo not implemented")
}
func (UnimplementedQueryServer) StoredGame(context.Context, *QueryGetStoredGameRequest) (*QueryGetStoredGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoredGame not implemented")
}
func (UnimplementedQueryServer) StoredGameAll(context.Context, *QueryAllStoredGameRequest) (*QueryAllStoredGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StoredGameAll not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_Params_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_SystemInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetSystemInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).SystemInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_SystemInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).SystemInfo(ctx, req.(*QueryGetSystemInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StoredGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetStoredGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StoredGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_StoredGame_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StoredGame(ctx, req.(*QueryGetStoredGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_StoredGameAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllStoredGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).StoredGameAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_StoredGameAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).StoredGameAll(ctx, req.(*QueryAllStoredGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "checkers.checkers.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "SystemInfo",
			Handler:    _Query_SystemInfo_Handler,
		},
		{
			MethodName: "StoredGame",
			Handler:    _Query_StoredGame_Handler,
		},
		{
			MethodName: "StoredGameAll",
			Handler:    _Query_StoredGameAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "checkers/checkers/query.proto",
}
