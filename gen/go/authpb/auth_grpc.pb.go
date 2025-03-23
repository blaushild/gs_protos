// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: authpb/auth.proto

package auth

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Authenticator_Register_FullMethodName = "/auth.Authenticator/Register"
	Authenticator_Login_FullMethodName    = "/auth.Authenticator/Login"
)

// AuthenticatorClient is the client API for Authenticator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthenticatorClient interface {
	Register(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*UserID, error)
	Login(ctx context.Context, in *UserLoginData, opts ...grpc.CallOption) (*Token, error)
}

type authenticatorClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthenticatorClient(cc grpc.ClientConnInterface) AuthenticatorClient {
	return &authenticatorClient{cc}
}

func (c *authenticatorClient) Register(ctx context.Context, in *UserData, opts ...grpc.CallOption) (*UserID, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UserID)
	err := c.cc.Invoke(ctx, Authenticator_Register_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authenticatorClient) Login(ctx context.Context, in *UserLoginData, opts ...grpc.CallOption) (*Token, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Token)
	err := c.cc.Invoke(ctx, Authenticator_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthenticatorServer is the server API for Authenticator service.
// All implementations must embed UnimplementedAuthenticatorServer
// for forward compatibility.
type AuthenticatorServer interface {
	Register(context.Context, *UserData) (*UserID, error)
	Login(context.Context, *UserLoginData) (*Token, error)
	mustEmbedUnimplementedAuthenticatorServer()
}

// UnimplementedAuthenticatorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAuthenticatorServer struct{}

func (UnimplementedAuthenticatorServer) Register(context.Context, *UserData) (*UserID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedAuthenticatorServer) Login(context.Context, *UserLoginData) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedAuthenticatorServer) mustEmbedUnimplementedAuthenticatorServer() {}
func (UnimplementedAuthenticatorServer) testEmbeddedByValue()                       {}

// UnsafeAuthenticatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthenticatorServer will
// result in compilation errors.
type UnsafeAuthenticatorServer interface {
	mustEmbedUnimplementedAuthenticatorServer()
}

func RegisterAuthenticatorServer(s grpc.ServiceRegistrar, srv AuthenticatorServer) {
	// If the following call pancis, it indicates UnimplementedAuthenticatorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Authenticator_ServiceDesc, srv)
}

func _Authenticator_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authenticator_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Register(ctx, req.(*UserData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Authenticator_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserLoginData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthenticatorServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Authenticator_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthenticatorServer).Login(ctx, req.(*UserLoginData))
	}
	return interceptor(ctx, in, info, handler)
}

// Authenticator_ServiceDesc is the grpc.ServiceDesc for Authenticator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Authenticator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Authenticator",
	HandlerType: (*AuthenticatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Authenticator_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Authenticator_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "authpb/auth.proto",
}
