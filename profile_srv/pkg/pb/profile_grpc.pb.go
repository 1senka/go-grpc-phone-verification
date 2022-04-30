// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: pkg/pb/profile.proto

package profilepb

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

// ProfileServiceClient is the client API for ProfileService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileServiceClient interface {
	ClientCreateProfile(ctx context.Context, in *ClientCreateProfileRequest, opts ...grpc.CallOption) (*ClientCreateProfileResponse, error)
	ClientGetProfile(ctx context.Context, in *ClientGetProfileRequest, opts ...grpc.CallOption) (*ClientGetProfileResponse, error)
	TherapistCreateProfile(ctx context.Context, in *TherapistCreateProfileRequest, opts ...grpc.CallOption) (*TherapistCreateProfileResponse, error)
	TherapistGetProfile(ctx context.Context, in *TherapistGetProfileRequest, opts ...grpc.CallOption) (*TherapistGetProfileResponse, error)
	ClientUpdateProfile(ctx context.Context, in *ClientUpdateProfileRequest, opts ...grpc.CallOption) (*ClientUpdateProfileResponse, error)
	TherapistUpdateProfile(ctx context.Context, in *TherapistUpdateProfileRequest, opts ...grpc.CallOption) (*TherapistUpdateProfileResponse, error)
}

type profileServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileServiceClient(cc grpc.ClientConnInterface) ProfileServiceClient {
	return &profileServiceClient{cc}
}

func (c *profileServiceClient) ClientCreateProfile(ctx context.Context, in *ClientCreateProfileRequest, opts ...grpc.CallOption) (*ClientCreateProfileResponse, error) {
	out := new(ClientCreateProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/ClientCreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) ClientGetProfile(ctx context.Context, in *ClientGetProfileRequest, opts ...grpc.CallOption) (*ClientGetProfileResponse, error) {
	out := new(ClientGetProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/ClientGetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) TherapistCreateProfile(ctx context.Context, in *TherapistCreateProfileRequest, opts ...grpc.CallOption) (*TherapistCreateProfileResponse, error) {
	out := new(TherapistCreateProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/TherapistCreateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) TherapistGetProfile(ctx context.Context, in *TherapistGetProfileRequest, opts ...grpc.CallOption) (*TherapistGetProfileResponse, error) {
	out := new(TherapistGetProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/TherapistGetProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) ClientUpdateProfile(ctx context.Context, in *ClientUpdateProfileRequest, opts ...grpc.CallOption) (*ClientUpdateProfileResponse, error) {
	out := new(ClientUpdateProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/ClientUpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileServiceClient) TherapistUpdateProfile(ctx context.Context, in *TherapistUpdateProfileRequest, opts ...grpc.CallOption) (*TherapistUpdateProfileResponse, error) {
	out := new(TherapistUpdateProfileResponse)
	err := c.cc.Invoke(ctx, "/profile.ProfileService/TherapistUpdateProfile", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileServiceServer is the server API for ProfileService service.
// All implementations must embed UnimplementedProfileServiceServer
// for forward compatibility
type ProfileServiceServer interface {
	ClientCreateProfile(context.Context, *ClientCreateProfileRequest) (*ClientCreateProfileResponse, error)
	ClientGetProfile(context.Context, *ClientGetProfileRequest) (*ClientGetProfileResponse, error)
	TherapistCreateProfile(context.Context, *TherapistCreateProfileRequest) (*TherapistCreateProfileResponse, error)
	TherapistGetProfile(context.Context, *TherapistGetProfileRequest) (*TherapistGetProfileResponse, error)
	ClientUpdateProfile(context.Context, *ClientUpdateProfileRequest) (*ClientUpdateProfileResponse, error)
	TherapistUpdateProfile(context.Context, *TherapistUpdateProfileRequest) (*TherapistUpdateProfileResponse, error)
	mustEmbedUnimplementedProfileServiceServer()
}

// UnimplementedProfileServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileServiceServer struct {
}

func (UnimplementedProfileServiceServer) ClientCreateProfile(context.Context, *ClientCreateProfileRequest) (*ClientCreateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientCreateProfile not implemented")
}
func (UnimplementedProfileServiceServer) ClientGetProfile(context.Context, *ClientGetProfileRequest) (*ClientGetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientGetProfile not implemented")
}
func (UnimplementedProfileServiceServer) TherapistCreateProfile(context.Context, *TherapistCreateProfileRequest) (*TherapistCreateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TherapistCreateProfile not implemented")
}
func (UnimplementedProfileServiceServer) TherapistGetProfile(context.Context, *TherapistGetProfileRequest) (*TherapistGetProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TherapistGetProfile not implemented")
}
func (UnimplementedProfileServiceServer) ClientUpdateProfile(context.Context, *ClientUpdateProfileRequest) (*ClientUpdateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClientUpdateProfile not implemented")
}
func (UnimplementedProfileServiceServer) TherapistUpdateProfile(context.Context, *TherapistUpdateProfileRequest) (*TherapistUpdateProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TherapistUpdateProfile not implemented")
}
func (UnimplementedProfileServiceServer) mustEmbedUnimplementedProfileServiceServer() {}

// UnsafeProfileServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileServiceServer will
// result in compilation errors.
type UnsafeProfileServiceServer interface {
	mustEmbedUnimplementedProfileServiceServer()
}

func RegisterProfileServiceServer(s grpc.ServiceRegistrar, srv ProfileServiceServer) {
	s.RegisterService(&ProfileService_ServiceDesc, srv)
}

func _ProfileService_ClientCreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientCreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).ClientCreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/ClientCreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).ClientCreateProfile(ctx, req.(*ClientCreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_ClientGetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientGetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).ClientGetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/ClientGetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).ClientGetProfile(ctx, req.(*ClientGetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_TherapistCreateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TherapistCreateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).TherapistCreateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/TherapistCreateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).TherapistCreateProfile(ctx, req.(*TherapistCreateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_TherapistGetProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TherapistGetProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).TherapistGetProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/TherapistGetProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).TherapistGetProfile(ctx, req.(*TherapistGetProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_ClientUpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClientUpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).ClientUpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/ClientUpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).ClientUpdateProfile(ctx, req.(*ClientUpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileService_TherapistUpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TherapistUpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileServiceServer).TherapistUpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/profile.ProfileService/TherapistUpdateProfile",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileServiceServer).TherapistUpdateProfile(ctx, req.(*TherapistUpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileService_ServiceDesc is the grpc.ServiceDesc for ProfileService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "profile.ProfileService",
	HandlerType: (*ProfileServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClientCreateProfile",
			Handler:    _ProfileService_ClientCreateProfile_Handler,
		},
		{
			MethodName: "ClientGetProfile",
			Handler:    _ProfileService_ClientGetProfile_Handler,
		},
		{
			MethodName: "TherapistCreateProfile",
			Handler:    _ProfileService_TherapistCreateProfile_Handler,
		},
		{
			MethodName: "TherapistGetProfile",
			Handler:    _ProfileService_TherapistGetProfile_Handler,
		},
		{
			MethodName: "ClientUpdateProfile",
			Handler:    _ProfileService_ClientUpdateProfile_Handler,
		},
		{
			MethodName: "TherapistUpdateProfile",
			Handler:    _ProfileService_TherapistUpdateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/profile.proto",
}
