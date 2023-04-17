// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.11
// source: protos/organization.proto

package protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Organization struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Arn                string `protobuf:"bytes,2,opt,name=arn,proto3" json:"arn,omitempty"`
	MasterAccountArn   string `protobuf:"bytes,3,opt,name=master_account_arn,json=masterAccountArn,proto3" json:"master_account_arn,omitempty"`
	MasterAccountId    string `protobuf:"bytes,4,opt,name=master_account_id,json=masterAccountId,proto3" json:"master_account_id,omitempty"`
	MasterAccountEmail string `protobuf:"bytes,5,opt,name=master_account_email,json=masterAccountEmail,proto3" json:"master_account_email,omitempty"`
}

func (x *Organization) Reset() {
	*x = Organization{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_organization_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Organization) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Organization) ProtoMessage() {}

func (x *Organization) ProtoReflect() protoreflect.Message {
	mi := &file_protos_organization_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Organization.ProtoReflect.Descriptor instead.
func (*Organization) Descriptor() ([]byte, []int) {
	return file_protos_organization_proto_rawDescGZIP(), []int{0}
}

func (x *Organization) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Organization) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

func (x *Organization) GetMasterAccountArn() string {
	if x != nil {
		return x.MasterAccountArn
	}
	return ""
}

func (x *Organization) GetMasterAccountId() string {
	if x != nil {
		return x.MasterAccountId
	}
	return ""
}

func (x *Organization) GetMasterAccountEmail() string {
	if x != nil {
		return x.MasterAccountEmail
	}
	return ""
}

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Arn             string `protobuf:"bytes,2,opt,name=arn,proto3" json:"arn,omitempty"`
	Email           string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Name            string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Status          string `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	JoinedMethod    string `protobuf:"bytes,6,opt,name=joined_method,json=joinedMethod,proto3" json:"joined_method,omitempty"`
	JoinedTimestamp string `protobuf:"bytes,7,opt,name=joined_timestamp,json=joinedTimestamp,proto3" json:"joined_timestamp,omitempty"`
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_organization_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_protos_organization_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_protos_organization_proto_rawDescGZIP(), []int{1}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

func (x *Account) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Account) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Account) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Account) GetJoinedMethod() string {
	if x != nil {
		return x.JoinedMethod
	}
	return ""
}

func (x *Account) GetJoinedTimestamp() string {
	if x != nil {
		return x.JoinedTimestamp
	}
	return ""
}

type ListOfAccounts struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accounts []*Account `protobuf:"bytes,1,rep,name=accounts,proto3" json:"accounts,omitempty"`
}

func (x *ListOfAccounts) Reset() {
	*x = ListOfAccounts{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_organization_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfAccounts) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfAccounts) ProtoMessage() {}

func (x *ListOfAccounts) ProtoReflect() protoreflect.Message {
	mi := &file_protos_organization_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfAccounts.ProtoReflect.Descriptor instead.
func (*ListOfAccounts) Descriptor() ([]byte, []int) {
	return file_protos_organization_proto_rawDescGZIP(), []int{2}
}

func (x *ListOfAccounts) GetAccounts() []*Account {
	if x != nil {
		return x.Accounts
	}
	return nil
}

type AccountId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
}

func (x *AccountId) Reset() {
	*x = AccountId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_organization_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountId) ProtoMessage() {}

func (x *AccountId) ProtoReflect() protoreflect.Message {
	mi := &file_protos_organization_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountId.ProtoReflect.Descriptor instead.
func (*AccountId) Descriptor() ([]byte, []int) {
	return file_protos_organization_proto_rawDescGZIP(), []int{3}
}

func (x *AccountId) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

type CreateAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email                  string  `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	AccountName            string  `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	RoleName               *string `protobuf:"bytes,3,opt,name=role_name,json=roleName,proto3,oneof" json:"role_name,omitempty"`
	IamUserAccessToBilling string  `protobuf:"bytes,4,opt,name=iam_user_access_to_billing,json=iamUserAccessToBilling,proto3" json:"iam_user_access_to_billing,omitempty"`
}

func (x *CreateAccountRequest) Reset() {
	*x = CreateAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_organization_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountRequest) ProtoMessage() {}

func (x *CreateAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_organization_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountRequest.ProtoReflect.Descriptor instead.
func (*CreateAccountRequest) Descriptor() ([]byte, []int) {
	return file_protos_organization_proto_rawDescGZIP(), []int{4}
}

func (x *CreateAccountRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateAccountRequest) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *CreateAccountRequest) GetRoleName() string {
	if x != nil && x.RoleName != nil {
		return *x.RoleName
	}
	return ""
}

func (x *CreateAccountRequest) GetIamUserAccessToBilling() string {
	if x != nil {
		return x.IamUserAccessToBilling
	}
	return ""
}

type CreateAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                 string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountName        string `protobuf:"bytes,2,opt,name=account_name,json=accountName,proto3" json:"account_name,omitempty"`
	AccountId          string `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	State              string `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	RequestedTimestamp string `protobuf:"bytes,5,opt,name=requested_timestamp,json=requestedTimestamp,proto3" json:"requested_timestamp,omitempty"`
	CompletedTimestamp string `protobuf:"bytes,6,opt,name=completed_timestamp,json=completedTimestamp,proto3" json:"completed_timestamp,omitempty"`
	FailureReason      string `protobuf:"bytes,7,opt,name=failure_reason,json=failureReason,proto3" json:"failure_reason,omitempty"`
}

func (x *CreateAccountResponse) Reset() {
	*x = CreateAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_organization_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateAccountResponse) ProtoMessage() {}

func (x *CreateAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_organization_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateAccountResponse.ProtoReflect.Descriptor instead.
func (*CreateAccountResponse) Descriptor() ([]byte, []int) {
	return file_protos_organization_proto_rawDescGZIP(), []int{5}
}

func (x *CreateAccountResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateAccountResponse) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

func (x *CreateAccountResponse) GetAccountId() string {
	if x != nil {
		return x.AccountId
	}
	return ""
}

func (x *CreateAccountResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CreateAccountResponse) GetRequestedTimestamp() string {
	if x != nil {
		return x.RequestedTimestamp
	}
	return ""
}

func (x *CreateAccountResponse) GetCompletedTimestamp() string {
	if x != nil {
		return x.CompletedTimestamp
	}
	return ""
}

func (x *CreateAccountResponse) GetFailureReason() string {
	if x != nil {
		return x.FailureReason
	}
	return ""
}

type InviteAccountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountId *string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3,oneof" json:"account_id,omitempty"`
	EmailId   *string `protobuf:"bytes,2,opt,name=email_id,json=emailId,proto3,oneof" json:"email_id,omitempty"`
}

func (x *InviteAccountRequest) Reset() {
	*x = InviteAccountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_organization_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteAccountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteAccountRequest) ProtoMessage() {}

func (x *InviteAccountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_organization_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteAccountRequest.ProtoReflect.Descriptor instead.
func (*InviteAccountRequest) Descriptor() ([]byte, []int) {
	return file_protos_organization_proto_rawDescGZIP(), []int{6}
}

func (x *InviteAccountRequest) GetAccountId() string {
	if x != nil && x.AccountId != nil {
		return *x.AccountId
	}
	return ""
}

func (x *InviteAccountRequest) GetEmailId() string {
	if x != nil && x.EmailId != nil {
		return *x.EmailId
	}
	return ""
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Type string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_organization_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_protos_organization_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_protos_organization_proto_rawDescGZIP(), []int{7}
}

func (x *Pair) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Pair) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type InviteAccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                  string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Arn                 string  `protobuf:"bytes,2,opt,name=arn,proto3" json:"arn,omitempty"`
	Parties             []*Pair `protobuf:"bytes,3,rep,name=parties,proto3" json:"parties,omitempty"`
	State               string  `protobuf:"bytes,4,opt,name=state,proto3" json:"state,omitempty"`
	RequestedTimestamp  string  `protobuf:"bytes,5,opt,name=requested_timestamp,json=requestedTimestamp,proto3" json:"requested_timestamp,omitempty"`
	ExpirationTimestamp string  `protobuf:"bytes,6,opt,name=expiration_timestamp,json=expirationTimestamp,proto3" json:"expiration_timestamp,omitempty"`
	Action              string  `protobuf:"bytes,7,opt,name=action,proto3" json:"action,omitempty"`
}

func (x *InviteAccountResponse) Reset() {
	*x = InviteAccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_organization_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InviteAccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InviteAccountResponse) ProtoMessage() {}

func (x *InviteAccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_organization_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InviteAccountResponse.ProtoReflect.Descriptor instead.
func (*InviteAccountResponse) Descriptor() ([]byte, []int) {
	return file_protos_organization_proto_rawDescGZIP(), []int{8}
}

func (x *InviteAccountResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *InviteAccountResponse) GetArn() string {
	if x != nil {
		return x.Arn
	}
	return ""
}

func (x *InviteAccountResponse) GetParties() []*Pair {
	if x != nil {
		return x.Parties
	}
	return nil
}

func (x *InviteAccountResponse) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *InviteAccountResponse) GetRequestedTimestamp() string {
	if x != nil {
		return x.RequestedTimestamp
	}
	return ""
}

func (x *InviteAccountResponse) GetExpirationTimestamp() string {
	if x != nil {
		return x.ExpirationTimestamp
	}
	return ""
}

func (x *InviteAccountResponse) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

var File_protos_organization_proto protoreflect.FileDescriptor

var file_protos_organization_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x6f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x0c, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x6e,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x72, 0x6e, 0x12, 0x2c, 0x0a, 0x12, 0x6d,
	0x61, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x61, 0x72,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x41, 0x72, 0x6e, 0x12, 0x2a, 0x0a, 0x11, 0x6d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xbd, 0x01, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x61, 0x72, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x23, 0x0a, 0x0d, 0x6a, 0x6f, 0x69, 0x6e, 0x65,
	0x64, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x29, 0x0a, 0x10,
	0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x36, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x4f,
	0x66, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x24, 0x0a, 0x08, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x22,
	0x2a, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xbb, 0x01, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x09, 0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x08, 0x72, 0x6f, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x3a, 0x0a, 0x1a, 0x69, 0x61, 0x6d, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x5f, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x16, 0x69, 0x61, 0x6d, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x54, 0x6f, 0x42, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x42, 0x0c, 0x0a, 0x0a, 0x5f,
	0x72, 0x6f, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x88, 0x02, 0x0a, 0x15, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x2f, 0x0a, 0x13,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x5f, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x65,
	0x61, 0x73, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x14, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x12, 0x1e, 0x0a, 0x08, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x07, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x49, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x42,
	0x0b, 0x0a, 0x09, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x04,
	0x50, 0x61, 0x69, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0xec, 0x01, 0x0a, 0x15, 0x49, 0x6e, 0x76,
	0x69, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x61, 0x72, 0x6e, 0x12, 0x1f, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x07, 0x70, 0x61,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x31, 0x0a, 0x14,
	0x65, 0x78, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x65, 0x78, 0x70, 0x69,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xa6, 0x02, 0x0a, 0x13, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x3a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x4f, 0x72, 0x67,
	0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x12, 0x46, 0x0a, 0x19, 0x47,
	0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x49, 0x6e, 0x4f, 0x72, 0x67, 0x61,
	0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x0f, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x49, 0x6e, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42,
	0x61, 0x73, 0x65, 0x64, 0x4f, 0x6e, 0x49, 0x64, 0x12, 0x0a, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x64, 0x1a, 0x08, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x00,
	0x12, 0x4e, 0x0a, 0x1b, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x54, 0x6f, 0x4f, 0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x15, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x49, 0x6e, 0x76, 0x69, 0x74, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x42, 0x5a, 0x40, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69,
	0x69, 0x2d, 0x76, 0x61, 0x72, 0x73, 0x68, 0x61, 0x2d, 0x69, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75,
	0x64, 0x2d, 0x6e, 0x73, 0x70, 0x2d, 0x63, 0x72, 0x6f, 0x73, 0x73, 0x2d, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_organization_proto_rawDescOnce sync.Once
	file_protos_organization_proto_rawDescData = file_protos_organization_proto_rawDesc
)

func file_protos_organization_proto_rawDescGZIP() []byte {
	file_protos_organization_proto_rawDescOnce.Do(func() {
		file_protos_organization_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_organization_proto_rawDescData)
	})
	return file_protos_organization_proto_rawDescData
}

var file_protos_organization_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_organization_proto_goTypes = []interface{}{
	(*Organization)(nil),          // 0: Organization
	(*Account)(nil),               // 1: Account
	(*ListOfAccounts)(nil),        // 2: ListOfAccounts
	(*AccountId)(nil),             // 3: AccountId
	(*CreateAccountRequest)(nil),  // 4: CreateAccountRequest
	(*CreateAccountResponse)(nil), // 5: CreateAccountResponse
	(*InviteAccountRequest)(nil),  // 6: InviteAccountRequest
	(*Pair)(nil),                  // 7: Pair
	(*InviteAccountResponse)(nil), // 8: InviteAccountResponse
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_protos_organization_proto_depIdxs = []int32{
	1, // 0: ListOfAccounts.accounts:type_name -> Account
	7, // 1: InviteAccountResponse.parties:type_name -> Pair
	9, // 2: OrganizationService.GetOrganization:input_type -> google.protobuf.Empty
	9, // 3: OrganizationService.GetAccountsInOrganization:input_type -> google.protobuf.Empty
	3, // 4: OrganizationService.GetAccountInOrganizationBasedOnId:input_type -> AccountId
	6, // 5: OrganizationService.InviteAccountToOrganization:input_type -> InviteAccountRequest
	0, // 6: OrganizationService.GetOrganization:output_type -> Organization
	2, // 7: OrganizationService.GetAccountsInOrganization:output_type -> ListOfAccounts
	1, // 8: OrganizationService.GetAccountInOrganizationBasedOnId:output_type -> Account
	8, // 9: OrganizationService.InviteAccountToOrganization:output_type -> InviteAccountResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_organization_proto_init() }
func file_protos_organization_proto_init() {
	if File_protos_organization_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_organization_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Organization); i {
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
		file_protos_organization_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Account); i {
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
		file_protos_organization_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfAccounts); i {
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
		file_protos_organization_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountId); i {
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
		file_protos_organization_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountRequest); i {
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
		file_protos_organization_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateAccountResponse); i {
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
		file_protos_organization_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteAccountRequest); i {
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
		file_protos_organization_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_protos_organization_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InviteAccountResponse); i {
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
	file_protos_organization_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_protos_organization_proto_msgTypes[6].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_organization_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_organization_proto_goTypes,
		DependencyIndexes: file_protos_organization_proto_depIdxs,
		MessageInfos:      file_protos_organization_proto_msgTypes,
	}.Build()
	File_protos_organization_proto = out.File
	file_protos_organization_proto_rawDesc = nil
	file_protos_organization_proto_goTypes = nil
	file_protos_organization_proto_depIdxs = nil
}
