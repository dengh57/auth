// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package authorization

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

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthorizationClient interface {
	AddModel(ctx context.Context, in *AddModelReq, opts ...grpc.CallOption) (*AddModelRes, error)
	AddPolicy(ctx context.Context, in *AddPolicyReq, opts ...grpc.CallOption) (*AddPolicyRes, error)
	DeletePolicy(ctx context.Context, in *DeletePolicyReq, opts ...grpc.CallOption) (*Empty, error)
	Authorization(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRes, error)
}

type authorizationClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthorizationClient(cc grpc.ClientConnInterface) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) AddModel(ctx context.Context, in *AddModelReq, opts ...grpc.CallOption) (*AddModelRes, error) {
	out := new(AddModelRes)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/AddModel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) AddPolicy(ctx context.Context, in *AddPolicyReq, opts ...grpc.CallOption) (*AddPolicyRes, error) {
	out := new(AddPolicyRes)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/AddPolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) DeletePolicy(ctx context.Context, in *DeletePolicyReq, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/DeletePolicy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authorizationClient) Authorization(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*AuthRes, error) {
	out := new(AuthRes)
	err := c.cc.Invoke(ctx, "/authorization.Authorization/Authorization", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
// All implementations must embed UnimplementedAuthorizationServer
// for forward compatibility
type AuthorizationServer interface {
	AddModel(context.Context, *AddModelReq) (*AddModelRes, error)
	AddPolicy(context.Context, *AddPolicyReq) (*AddPolicyRes, error)
	DeletePolicy(context.Context, *DeletePolicyReq) (*Empty, error)
	Authorization(context.Context, *AuthReq) (*AuthRes, error)
	mustEmbedUnimplementedAuthorizationServer()
}

// UnimplementedAuthorizationServer must be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (UnimplementedAuthorizationServer) AddModel(context.Context, *AddModelReq) (*AddModelRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddModel not implemented")
}
func (UnimplementedAuthorizationServer) AddPolicy(context.Context, *AddPolicyReq) (*AddPolicyRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPolicy not implemented")
}
func (UnimplementedAuthorizationServer) DeletePolicy(context.Context, *DeletePolicyReq) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePolicy not implemented")
}
func (UnimplementedAuthorizationServer) Authorization(context.Context, *AuthReq) (*AuthRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authorization not implemented")
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

func _Authorization_AddModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddModelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).AddModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/AddModel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).AddModel(ctx, req.(*AddModelReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_AddPolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).AddPolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/AddPolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).AddPolicy(ctx, req.(*AddPolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_DeletePolicy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeletePolicyReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).DeletePolicy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/DeletePolicy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).DeletePolicy(ctx, req.(*DeletePolicyReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authorization_Authorization_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Authorization(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/authorization.Authorization/Authorization",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Authorization(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Authorization_ServiceDesc is the grpc.ServiceDesc for Authorization service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authorization_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authorization.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddModel",
			Handler:    _Authorization_AddModel_Handler,
		},
		{
			MethodName: "AddPolicy",
			Handler:    _Authorization_AddPolicy_Handler,
		},
		{
			MethodName: "DeletePolicy",
			Handler:    _Authorization_DeletePolicy_Handler,
		},
		{
			MethodName: "Authorization",
			Handler:    _Authorization_Authorization_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authorization/auth.proto",
}