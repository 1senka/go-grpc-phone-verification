// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: client/clientpb/client.proto

package clientpb

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

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientServiceClient interface {
	PhoneProve(ctx context.Context, in *ClientPhoneProofRequest, opts ...grpc.CallOption) (*ClientPhoneProofResponse, error)
	FirstCred(ctx context.Context, in *ClientFirstCredRequest, opts ...grpc.CallOption) (*ClientFirstCredResponse, error)
	PhoneToken(ctx context.Context, in *ClientPhoneTokenRequest, opts ...grpc.CallOption) (*ClientPhoneTokenResponse, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) PhoneProve(ctx context.Context, in *ClientPhoneProofRequest, opts ...grpc.CallOption) (*ClientPhoneProofResponse, error) {
	out := new(ClientPhoneProofResponse)
	err := c.cc.Invoke(ctx, "/client.ClientService/PhoneProve", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) FirstCred(ctx context.Context, in *ClientFirstCredRequest, opts ...grpc.CallOption) (*ClientFirstCredResponse, error) {
	out := new(ClientFirstCredResponse)
	err := c.cc.Invoke(ctx, "/client.ClientService/FirstCred", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) PhoneToken(ctx context.Context, in *ClientPhoneTokenRequest, opts ...grpc.CallOption) (*ClientPhoneTokenResponse, error) {
	out := new(ClientPhoneTokenResponse)
	err := c.cc.Invoke(ctx, "/client.ClientService/PhoneToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClientServiceServer is the server API for ClientService service.
// All implementations must embed UnimplementedClientServiceServer
// for forward compatibility
type ClientServiceServer interface {
	PhoneProve(context.Context, *ClientPhoneProofRequest) (*ClientPhoneProofResponse, error)
	FirstCred(context.Context, *ClientFirstCredRequest) (*ClientFirstCredResponse, error)
	PhoneToken(context.Context, *ClientPhoneTokenRequest) (*ClientPhoneTokenResponse, error)
	//mustEmbedUnimplementedClientServiceServer()
}

// UnimplementedClientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (UnimplementedClientServiceServer) PhoneProve(context.Context, *ClientPhoneProofRequest) (*ClientPhoneProofResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhoneProve not implemented")
}
func (UnimplementedClientServiceServer) FirstCred(context.Context, *ClientFirstCredRequest) (*ClientFirstCredResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FirstCred not implemented")
}
func (UnimplementedClientServiceServer) PhoneToken(context.Context, *ClientPhoneTokenRequest) (*ClientPhoneTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PhoneToken not implemented")
}
func (UnimplementedClientServiceServer) mustEmbedUnimplementedClientServiceServer() {}

// UnsafeClientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientServiceServer will
// result in compilation errors.
type UnsafeClientServiceServer interface {
	mustEmbedUnimplementedClientServiceServer()
}

func RegisterClientServiceServer(s grpc.ServiceRegistrar, srv ClientServiceServer) {
	s.RegisterService(&ClientService_ServiceDesc, srv)
}

func _ClientService_PhoneProve_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientPhoneProofRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).PhoneProve(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientService/PhoneProve",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).PhoneProve(ctx, req.(*ClientPhoneProofRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_FirstCred_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientFirstCredRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).FirstCred(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientService/FirstCred",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).FirstCred(ctx, req.(*ClientFirstCredRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_PhoneToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientPhoneTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).PhoneToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.ClientService/PhoneToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).PhoneToken(ctx, req.(*ClientPhoneTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClientService_ServiceDesc is the grpc.ServiceDesc for ClientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "client.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PhoneProve",
			Handler:    _ClientService_PhoneProve_Handler,
		},
		{
			MethodName: "FirstCred",
			Handler:    _ClientService_FirstCred_Handler,
		},
		{
			MethodName: "PhoneToken",
			Handler:    _ClientService_PhoneToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/clientpb/client.proto",
}
