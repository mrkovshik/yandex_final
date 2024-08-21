// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: proto/memento.proto

package proto

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
	Memento_AddUser_FullMethodName        = "/api.grpc.Memento/AddUser"
	Memento_GetToken_FullMethodName       = "/api.grpc.Memento/GetToken"
	Memento_AddCredential_FullMethodName  = "/api.grpc.Memento/AddCredential"
	Memento_GetCredentials_FullMethodName = "/api.grpc.Memento/GetCredentialsByUserID"
	Memento_AddCardData_FullMethodName    = "/api.grpc.Memento/AddCardData"
	Memento_GetCards_FullMethodName       = "/api.grpc.Memento/GetCards"
	Memento_AddVariousData_FullMethodName = "/api.grpc.Memento/AddVariousData"
)

// MementoClient is the client API for Memento service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MementoClient interface {
	AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error)
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
	AddCredential(ctx context.Context, in *AddCredentialRequest, opts ...grpc.CallOption) (*AddCredentialResponse, error)
	GetCredentials(ctx context.Context, in *GetCredentialsRequest, opts ...grpc.CallOption) (*GetCredentialsResponse, error)
	AddCardData(ctx context.Context, in *AddCardDataRequest, opts ...grpc.CallOption) (*AddCardDataResponse, error)
	GetCards(ctx context.Context, in *GetCardsRequest, opts ...grpc.CallOption) (*GetCardsResponse, error)
	AddVariousData(ctx context.Context, in *AddVariousDataRequest, opts ...grpc.CallOption) (*AddVariousDataResponse, error)
}

type mementoClient struct {
	cc grpc.ClientConnInterface
}

func NewMementoClient(cc grpc.ClientConnInterface) MementoClient {
	return &mementoClient{cc}
}

func (c *mementoClient) AddUser(ctx context.Context, in *AddUserRequest, opts ...grpc.CallOption) (*AddUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddUserResponse)
	err := c.cc.Invoke(ctx, Memento_AddUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mementoClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, Memento_GetToken_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mementoClient) AddCredential(ctx context.Context, in *AddCredentialRequest, opts ...grpc.CallOption) (*AddCredentialResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCredentialResponse)
	err := c.cc.Invoke(ctx, Memento_AddCredential_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mementoClient) GetCredentials(ctx context.Context, in *GetCredentialsRequest, opts ...grpc.CallOption) (*GetCredentialsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCredentialsResponse)
	err := c.cc.Invoke(ctx, Memento_GetCredentials_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mementoClient) AddCardData(ctx context.Context, in *AddCardDataRequest, opts ...grpc.CallOption) (*AddCardDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddCardDataResponse)
	err := c.cc.Invoke(ctx, Memento_AddCardData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mementoClient) GetCards(ctx context.Context, in *GetCardsRequest, opts ...grpc.CallOption) (*GetCardsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCardsResponse)
	err := c.cc.Invoke(ctx, Memento_GetCards_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mementoClient) AddVariousData(ctx context.Context, in *AddVariousDataRequest, opts ...grpc.CallOption) (*AddVariousDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddVariousDataResponse)
	err := c.cc.Invoke(ctx, Memento_AddVariousData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MementoServer is the server API for Memento service.
// All implementations must embed UnimplementedMementoServer
// for forward compatibility.
type MementoServer interface {
	AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error)
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
	AddCredential(context.Context, *AddCredentialRequest) (*AddCredentialResponse, error)
	GetCredentials(context.Context, *GetCredentialsRequest) (*GetCredentialsResponse, error)
	AddCardData(context.Context, *AddCardDataRequest) (*AddCardDataResponse, error)
	GetCards(context.Context, *GetCardsRequest) (*GetCardsResponse, error)
	AddVariousData(context.Context, *AddVariousDataRequest) (*AddVariousDataResponse, error)
	mustEmbedUnimplementedMementoServer()
}

// UnimplementedMementoServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMementoServer struct{}

func (UnimplementedMementoServer) AddUser(context.Context, *AddUserRequest) (*AddUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUser not implemented")
}
func (UnimplementedMementoServer) GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToken not implemented")
}
func (UnimplementedMementoServer) AddCredential(context.Context, *AddCredentialRequest) (*AddCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCredential not implemented")
}
func (UnimplementedMementoServer) GetCredentials(context.Context, *GetCredentialsRequest) (*GetCredentialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCredentialsByUserID not implemented")
}
func (UnimplementedMementoServer) AddCardData(context.Context, *AddCardDataRequest) (*AddCardDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCardData not implemented")
}
func (UnimplementedMementoServer) GetCards(context.Context, *GetCardsRequest) (*GetCardsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCards not implemented")
}
func (UnimplementedMementoServer) AddVariousData(context.Context, *AddVariousDataRequest) (*AddVariousDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVariousData not implemented")
}
func (UnimplementedMementoServer) mustEmbedUnimplementedMementoServer() {}
func (UnimplementedMementoServer) testEmbeddedByValue()                 {}

// UnsafeMementoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MementoServer will
// result in compilation errors.
type UnsafeMementoServer interface {
	mustEmbedUnimplementedMementoServer()
}

func RegisterMementoServer(s grpc.ServiceRegistrar, srv MementoServer) {
	// If the following call pancis, it indicates UnimplementedMementoServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Memento_ServiceDesc, srv)
}

func _Memento_AddUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MementoServer).AddUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Memento_AddUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MementoServer).AddUser(ctx, req.(*AddUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memento_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MementoServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Memento_GetToken_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MementoServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memento_AddCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCredentialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MementoServer).AddCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Memento_AddCredential_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MementoServer).AddCredential(ctx, req.(*AddCredentialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memento_GetCredentials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCredentialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MementoServer).GetCredentials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Memento_GetCredentials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MementoServer).GetCredentials(ctx, req.(*GetCredentialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memento_AddCardData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCardDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MementoServer).AddCardData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Memento_AddCardData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MementoServer).AddCardData(ctx, req.(*AddCardDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memento_GetCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MementoServer).GetCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Memento_GetCards_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MementoServer).GetCards(ctx, req.(*GetCardsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Memento_AddVariousData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVariousDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MementoServer).AddVariousData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Memento_AddVariousData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MementoServer).AddVariousData(ctx, req.(*AddVariousDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Memento_ServiceDesc is the grpc.ServiceDesc for Memento service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Memento_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.grpc.Memento",
	HandlerType: (*MementoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddUser",
			Handler:    _Memento_AddUser_Handler,
		},
		{
			MethodName: "GetToken",
			Handler:    _Memento_GetToken_Handler,
		},
		{
			MethodName: "AddCredential",
			Handler:    _Memento_AddCredential_Handler,
		},
		{
			MethodName: "GetCredentialsByUserID",
			Handler:    _Memento_GetCredentials_Handler,
		},
		{
			MethodName: "AddCardData",
			Handler:    _Memento_AddCardData_Handler,
		},
		{
			MethodName: "GetCards",
			Handler:    _Memento_GetCards_Handler,
		},
		{
			MethodName: "AddVariousData",
			Handler:    _Memento_AddVariousData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/memento.proto",
}
