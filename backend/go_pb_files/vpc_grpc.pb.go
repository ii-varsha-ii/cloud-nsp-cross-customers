// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.11
// source: protos/vpc.proto

package protos

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
	VpcService_CreateVpcWithSubnets_FullMethodName = "/VpcService/CreateVpcWithSubnets"
	VpcService_GetVpcBasedOnRegion_FullMethodName  = "/VpcService/GetVpcBasedOnRegion"
	VpcService_GetSubnetsInVpc_FullMethodName      = "/VpcService/GetSubnetsInVpc"
	VpcService_CreateSubnet_FullMethodName         = "/VpcService/CreateSubnet"
)

// VpcServiceClient is the client API for VpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VpcServiceClient interface {
	CreateVpcWithSubnets(ctx context.Context, in *CreateVpcRequest, opts ...grpc.CallOption) (*Vpc, error)
	GetVpcBasedOnRegion(ctx context.Context, in *RegionAndVpcRequest, opts ...grpc.CallOption) (*ListOfVpcs, error)
	GetSubnetsInVpc(ctx context.Context, in *RegionAndVpcRequest, opts ...grpc.CallOption) (*ListOfSubnets, error)
	CreateSubnet(ctx context.Context, in *CreateSubnetRequest, opts ...grpc.CallOption) (*Subnet, error)
}

type vpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVpcServiceClient(cc grpc.ClientConnInterface) VpcServiceClient {
	return &vpcServiceClient{cc}
}

func (c *vpcServiceClient) CreateVpcWithSubnets(ctx context.Context, in *CreateVpcRequest, opts ...grpc.CallOption) (*Vpc, error) {
	out := new(Vpc)
	err := c.cc.Invoke(ctx, VpcService_CreateVpcWithSubnets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vpcServiceClient) GetVpcBasedOnRegion(ctx context.Context, in *RegionAndVpcRequest, opts ...grpc.CallOption) (*ListOfVpcs, error) {
	out := new(ListOfVpcs)
	err := c.cc.Invoke(ctx, VpcService_GetVpcBasedOnRegion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vpcServiceClient) GetSubnetsInVpc(ctx context.Context, in *RegionAndVpcRequest, opts ...grpc.CallOption) (*ListOfSubnets, error) {
	out := new(ListOfSubnets)
	err := c.cc.Invoke(ctx, VpcService_GetSubnetsInVpc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vpcServiceClient) CreateSubnet(ctx context.Context, in *CreateSubnetRequest, opts ...grpc.CallOption) (*Subnet, error) {
	out := new(Subnet)
	err := c.cc.Invoke(ctx, VpcService_CreateSubnet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VpcServiceServer is the server API for VpcService service.
// All implementations must embed UnimplementedVpcServiceServer
// for forward compatibility
type VpcServiceServer interface {
	CreateVpcWithSubnets(context.Context, *CreateVpcRequest) (*Vpc, error)
	GetVpcBasedOnRegion(context.Context, *RegionAndVpcRequest) (*ListOfVpcs, error)
	GetSubnetsInVpc(context.Context, *RegionAndVpcRequest) (*ListOfSubnets, error)
	CreateSubnet(context.Context, *CreateSubnetRequest) (*Subnet, error)
	mustEmbedUnimplementedVpcServiceServer()
}

// UnimplementedVpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVpcServiceServer struct {
}

func (UnimplementedVpcServiceServer) CreateVpcWithSubnets(context.Context, *CreateVpcRequest) (*Vpc, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVpcWithSubnets not implemented")
}
func (UnimplementedVpcServiceServer) GetVpcBasedOnRegion(context.Context, *RegionAndVpcRequest) (*ListOfVpcs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVpcBasedOnRegion not implemented")
}
func (UnimplementedVpcServiceServer) GetSubnetsInVpc(context.Context, *RegionAndVpcRequest) (*ListOfSubnets, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubnetsInVpc not implemented")
}
func (UnimplementedVpcServiceServer) CreateSubnet(context.Context, *CreateSubnetRequest) (*Subnet, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubnet not implemented")
}
func (UnimplementedVpcServiceServer) mustEmbedUnimplementedVpcServiceServer() {}

// UnsafeVpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VpcServiceServer will
// result in compilation errors.
type UnsafeVpcServiceServer interface {
	mustEmbedUnimplementedVpcServiceServer()
}

func RegisterVpcServiceServer(s grpc.ServiceRegistrar, srv VpcServiceServer) {
	s.RegisterService(&VpcService_ServiceDesc, srv)
}

func _VpcService_CreateVpcWithSubnets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VpcServiceServer).CreateVpcWithSubnets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VpcService_CreateVpcWithSubnets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VpcServiceServer).CreateVpcWithSubnets(ctx, req.(*CreateVpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VpcService_GetVpcBasedOnRegion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegionAndVpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VpcServiceServer).GetVpcBasedOnRegion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VpcService_GetVpcBasedOnRegion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VpcServiceServer).GetVpcBasedOnRegion(ctx, req.(*RegionAndVpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VpcService_GetSubnetsInVpc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegionAndVpcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VpcServiceServer).GetSubnetsInVpc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VpcService_GetSubnetsInVpc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VpcServiceServer).GetSubnetsInVpc(ctx, req.(*RegionAndVpcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VpcService_CreateSubnet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubnetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VpcServiceServer).CreateSubnet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VpcService_CreateSubnet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VpcServiceServer).CreateSubnet(ctx, req.(*CreateSubnetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VpcService_ServiceDesc is the grpc.ServiceDesc for VpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "VpcService",
	HandlerType: (*VpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVpcWithSubnets",
			Handler:    _VpcService_CreateVpcWithSubnets_Handler,
		},
		{
			MethodName: "GetVpcBasedOnRegion",
			Handler:    _VpcService_GetVpcBasedOnRegion_Handler,
		},
		{
			MethodName: "GetSubnetsInVpc",
			Handler:    _VpcService_GetSubnetsInVpc_Handler,
		},
		{
			MethodName: "CreateSubnet",
			Handler:    _VpcService_CreateSubnet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/vpc.proto",
}
