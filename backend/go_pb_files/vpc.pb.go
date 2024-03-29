// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.11
// source: protos/vpc.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegionAndVpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region string  `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	VpcId  *string `protobuf:"bytes,2,opt,name=vpc_id,json=vpcId,proto3,oneof" json:"vpc_id,omitempty"`
}

func (x *RegionAndVpcRequest) Reset() {
	*x = RegionAndVpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vpc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegionAndVpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegionAndVpcRequest) ProtoMessage() {}

func (x *RegionAndVpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vpc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegionAndVpcRequest.ProtoReflect.Descriptor instead.
func (*RegionAndVpcRequest) Descriptor() ([]byte, []int) {
	return file_protos_vpc_proto_rawDescGZIP(), []int{0}
}

func (x *RegionAndVpcRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *RegionAndVpcRequest) GetVpcId() string {
	if x != nil && x.VpcId != nil {
		return *x.VpcId
	}
	return ""
}

type ListOfVpcs struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Vpc []*Vpc `protobuf:"bytes,1,rep,name=vpc,proto3" json:"vpc,omitempty"`
}

func (x *ListOfVpcs) Reset() {
	*x = ListOfVpcs{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vpc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfVpcs) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfVpcs) ProtoMessage() {}

func (x *ListOfVpcs) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vpc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfVpcs.ProtoReflect.Descriptor instead.
func (*ListOfVpcs) Descriptor() ([]byte, []int) {
	return file_protos_vpc_proto_rawDescGZIP(), []int{1}
}

func (x *ListOfVpcs) GetVpc() []*Vpc {
	if x != nil {
		return x.Vpc
	}
	return nil
}

type ListOfSubnets struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnets []*Subnet `protobuf:"bytes,1,rep,name=subnets,proto3" json:"subnets,omitempty"`
}

func (x *ListOfSubnets) Reset() {
	*x = ListOfSubnets{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vpc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfSubnets) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfSubnets) ProtoMessage() {}

func (x *ListOfSubnets) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vpc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfSubnets.ProtoReflect.Descriptor instead.
func (*ListOfSubnets) Descriptor() ([]byte, []int) {
	return file_protos_vpc_proto_rawDescGZIP(), []int{2}
}

func (x *ListOfSubnets) GetSubnets() []*Subnet {
	if x != nil {
		return x.Subnets
	}
	return nil
}

type Vpc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CidrBlock string    `protobuf:"bytes,2,opt,name=cidr_block,json=cidrBlock,proto3" json:"cidr_block,omitempty"`
	Name      string    `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Region    string    `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	Subnets   []*Subnet `protobuf:"bytes,5,rep,name=subnets,proto3" json:"subnets,omitempty"`
}

func (x *Vpc) Reset() {
	*x = Vpc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vpc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vpc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vpc) ProtoMessage() {}

func (x *Vpc) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vpc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vpc.ProtoReflect.Descriptor instead.
func (*Vpc) Descriptor() ([]byte, []int) {
	return file_protos_vpc_proto_rawDescGZIP(), []int{3}
}

func (x *Vpc) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Vpc) GetCidrBlock() string {
	if x != nil {
		return x.CidrBlock
	}
	return ""
}

func (x *Vpc) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Vpc) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Vpc) GetSubnets() []*Subnet {
	if x != nil {
		return x.Subnets
	}
	return nil
}

type Subnet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AvailabilityZone string `protobuf:"bytes,3,opt,name=availability_zone,json=availabilityZone,proto3" json:"availability_zone,omitempty"`
	CidrBlock        string `protobuf:"bytes,4,opt,name=cidr_block,json=cidrBlock,proto3" json:"cidr_block,omitempty"`
	VpcId            string `protobuf:"bytes,5,opt,name=vpc_id,json=vpcId,proto3" json:"vpc_id,omitempty"`
	RouteTableId     string `protobuf:"bytes,6,opt,name=route_table_id,json=routeTableId,proto3" json:"route_table_id,omitempty"`
}

func (x *Subnet) Reset() {
	*x = Subnet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vpc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subnet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subnet) ProtoMessage() {}

func (x *Subnet) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vpc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subnet.ProtoReflect.Descriptor instead.
func (*Subnet) Descriptor() ([]byte, []int) {
	return file_protos_vpc_proto_rawDescGZIP(), []int{4}
}

func (x *Subnet) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Subnet) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subnet) GetAvailabilityZone() string {
	if x != nil {
		return x.AvailabilityZone
	}
	return ""
}

func (x *Subnet) GetCidrBlock() string {
	if x != nil {
		return x.CidrBlock
	}
	return ""
}

func (x *Subnet) GetVpcId() string {
	if x != nil {
		return x.VpcId
	}
	return ""
}

func (x *Subnet) GetRouteTableId() string {
	if x != nil {
		return x.RouteTableId
	}
	return ""
}

type CreateVpcRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region  string                 `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Name    string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	IpCidr  string                 `protobuf:"bytes,3,opt,name=ip_cidr,json=ipCidr,proto3" json:"ip_cidr,omitempty"`
	Subnets []*CreateSubnetRequest `protobuf:"bytes,4,rep,name=subnets,proto3" json:"subnets,omitempty"`
}

func (x *CreateVpcRequest) Reset() {
	*x = CreateVpcRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vpc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateVpcRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateVpcRequest) ProtoMessage() {}

func (x *CreateVpcRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vpc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateVpcRequest.ProtoReflect.Descriptor instead.
func (*CreateVpcRequest) Descriptor() ([]byte, []int) {
	return file_protos_vpc_proto_rawDescGZIP(), []int{5}
}

func (x *CreateVpcRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CreateVpcRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateVpcRequest) GetIpCidr() string {
	if x != nil {
		return x.IpCidr
	}
	return ""
}

func (x *CreateVpcRequest) GetSubnets() []*CreateSubnetRequest {
	if x != nil {
		return x.Subnets
	}
	return nil
}

type CreateSubnetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Region           string `protobuf:"bytes,1,opt,name=region,proto3" json:"region,omitempty"`
	Name             string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AvailabilityZone string `protobuf:"bytes,3,opt,name=availability_zone,json=availabilityZone,proto3" json:"availability_zone,omitempty"`
	CidrBlock        string `protobuf:"bytes,4,opt,name=cidr_block,json=cidrBlock,proto3" json:"cidr_block,omitempty"`
	VpcId            string `protobuf:"bytes,5,opt,name=vpc_id,json=vpcId,proto3" json:"vpc_id,omitempty"`
	Access           string `protobuf:"bytes,6,opt,name=access,proto3" json:"access,omitempty"`
}

func (x *CreateSubnetRequest) Reset() {
	*x = CreateSubnetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_vpc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSubnetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSubnetRequest) ProtoMessage() {}

func (x *CreateSubnetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_vpc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSubnetRequest.ProtoReflect.Descriptor instead.
func (*CreateSubnetRequest) Descriptor() ([]byte, []int) {
	return file_protos_vpc_proto_rawDescGZIP(), []int{6}
}

func (x *CreateSubnetRequest) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *CreateSubnetRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateSubnetRequest) GetAvailabilityZone() string {
	if x != nil {
		return x.AvailabilityZone
	}
	return ""
}

func (x *CreateSubnetRequest) GetCidrBlock() string {
	if x != nil {
		return x.CidrBlock
	}
	return ""
}

func (x *CreateSubnetRequest) GetVpcId() string {
	if x != nil {
		return x.VpcId
	}
	return ""
}

func (x *CreateSubnetRequest) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

var File_protos_vpc_proto protoreflect.FileDescriptor

var file_protos_vpc_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x70, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x54, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x56,
	0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x12, 0x1a, 0x0a, 0x06, 0x76, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x05, 0x76, 0x70, 0x63, 0x49, 0x64, 0x88, 0x01, 0x01, 0x42, 0x09, 0x0a,
	0x07, 0x5f, 0x76, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x0a, 0x4c, 0x69, 0x73, 0x74,
	0x4f, 0x66, 0x56, 0x70, 0x63, 0x73, 0x12, 0x16, 0x0a, 0x03, 0x76, 0x70, 0x63, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x04, 0x2e, 0x56, 0x70, 0x63, 0x52, 0x03, 0x76, 0x70, 0x63, 0x22, 0x32,
	0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12,
	0x21, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x07, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6e, 0x65,
	0x74, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x03, 0x56, 0x70, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x69,
	0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x0a, 0x07, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52,
	0x07, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x22, 0xb5, 0x01, 0x0a, 0x06, 0x53, 0x75, 0x62,
	0x6e, 0x65, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x7a, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79,
	0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x69, 0x64, 0x72, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x70, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x70, 0x63, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x6f,
	0x75, 0x74, 0x65, 0x5f, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x64,
	0x22, 0x87, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x70, 0x63, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x63, 0x69, 0x64, 0x72, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x69, 0x70, 0x43, 0x69, 0x64, 0x72, 0x12, 0x2e, 0x0a, 0x07, 0x73, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x07, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x22, 0xbc, 0x01, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x2b,
	0x0a, 0x11, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x7a,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x76, 0x61, 0x69, 0x6c,
	0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x5a, 0x6f, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x69, 0x64, 0x72, 0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x69, 0x64, 0x72, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x76, 0x70,
	0x63, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x70, 0x63, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xdf, 0x01, 0x0a, 0x0a, 0x56, 0x70,
	0x63, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x56, 0x70, 0x63, 0x57, 0x69, 0x74, 0x68, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73,
	0x12, 0x11, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x56, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x04, 0x2e, 0x56, 0x70, 0x63, 0x12, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x56, 0x70, 0x63, 0x42, 0x61, 0x73, 0x65, 0x64, 0x4f, 0x6e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x12, 0x14, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x64, 0x56, 0x70, 0x63, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x56,
	0x70, 0x63, 0x73, 0x12, 0x37, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74,
	0x73, 0x49, 0x6e, 0x56, 0x70, 0x63, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x41,
	0x6e, 0x64, 0x56, 0x70, 0x63, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x73, 0x12, 0x2d, 0x0a, 0x0c,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x12, 0x14, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x07, 0x2e, 0x53, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x42, 0x42, 0x5a, 0x40, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x69, 0x2d, 0x76, 0x61, 0x72,
	0x73, 0x68, 0x61, 0x2d, 0x69, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2d, 0x6e, 0x73, 0x70,
	0x2d, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x73,
	0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_vpc_proto_rawDescOnce sync.Once
	file_protos_vpc_proto_rawDescData = file_protos_vpc_proto_rawDesc
)

func file_protos_vpc_proto_rawDescGZIP() []byte {
	file_protos_vpc_proto_rawDescOnce.Do(func() {
		file_protos_vpc_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_vpc_proto_rawDescData)
	})
	return file_protos_vpc_proto_rawDescData
}

var file_protos_vpc_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protos_vpc_proto_goTypes = []interface{}{
	(*RegionAndVpcRequest)(nil), // 0: RegionAndVpcRequest
	(*ListOfVpcs)(nil),          // 1: ListOfVpcs
	(*ListOfSubnets)(nil),       // 2: ListOfSubnets
	(*Vpc)(nil),                 // 3: Vpc
	(*Subnet)(nil),              // 4: Subnet
	(*CreateVpcRequest)(nil),    // 5: CreateVpcRequest
	(*CreateSubnetRequest)(nil), // 6: CreateSubnetRequest
}
var file_protos_vpc_proto_depIdxs = []int32{
	3, // 0: ListOfVpcs.vpc:type_name -> Vpc
	4, // 1: ListOfSubnets.subnets:type_name -> Subnet
	4, // 2: Vpc.subnets:type_name -> Subnet
	6, // 3: CreateVpcRequest.subnets:type_name -> CreateSubnetRequest
	5, // 4: VpcService.CreateVpcWithSubnets:input_type -> CreateVpcRequest
	0, // 5: VpcService.GetVpcBasedOnRegion:input_type -> RegionAndVpcRequest
	0, // 6: VpcService.GetSubnetsInVpc:input_type -> RegionAndVpcRequest
	6, // 7: VpcService.CreateSubnet:input_type -> CreateSubnetRequest
	3, // 8: VpcService.CreateVpcWithSubnets:output_type -> Vpc
	1, // 9: VpcService.GetVpcBasedOnRegion:output_type -> ListOfVpcs
	2, // 10: VpcService.GetSubnetsInVpc:output_type -> ListOfSubnets
	4, // 11: VpcService.CreateSubnet:output_type -> Subnet
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_protos_vpc_proto_init() }
func file_protos_vpc_proto_init() {
	if File_protos_vpc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_vpc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegionAndVpcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_vpc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfVpcs); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_vpc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfSubnets); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_vpc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vpc); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_vpc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subnet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_vpc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateVpcRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_vpc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSubnetRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_protos_vpc_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_vpc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_vpc_proto_goTypes,
		DependencyIndexes: file_protos_vpc_proto_depIdxs,
		MessageInfos:      file_protos_vpc_proto_msgTypes,
	}.Build()
	File_protos_vpc_proto = out.File
	file_protos_vpc_proto_rawDesc = nil
	file_protos_vpc_proto_goTypes = nil
	file_protos_vpc_proto_depIdxs = nil
}
