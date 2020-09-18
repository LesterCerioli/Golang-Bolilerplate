// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: src/transport/grpc/proto/account.proto

package pb

import (
	context "context"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Account Update Request
type AccountUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID   uint32 `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	Name string `protobuf:"bytes,2,opt,name=Name,proto3" json:"name"`
}

func (x *AccountUpdateRequest) Reset() {
	*x = AccountUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_grpc_proto_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountUpdateRequest) ProtoMessage() {}

func (x *AccountUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_grpc_proto_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountUpdateRequest.ProtoReflect.Descriptor instead.
func (*AccountUpdateRequest) Descriptor() ([]byte, []int) {
	return file_src_transport_grpc_proto_account_proto_rawDescGZIP(), []int{0}
}

func (x *AccountUpdateRequest) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AccountUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// Account Get/Update Response
type AccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32                `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	Name      string                `protobuf:"bytes,2,opt,name=Name,proto3" json:"name"`
	Website   *wrappers.StringValue `protobuf:"bytes,3,opt,name=Website,proto3" json:"website"`
	CreatedAt *timestamp.Timestamp  `protobuf:"bytes,4,opt,name=CreatedAt,proto3" json:"created_at"`
	UpdatedAt *timestamp.Timestamp  `protobuf:"bytes,5,opt,name=UpdatedAt,proto3" json:"updated_at"`
}

func (x *AccountResponse) Reset() {
	*x = AccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_grpc_proto_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountResponse) ProtoMessage() {}

func (x *AccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_grpc_proto_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountResponse.ProtoReflect.Descriptor instead.
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return file_src_transport_grpc_proto_account_proto_rawDescGZIP(), []int{1}
}

func (x *AccountResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AccountResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountResponse) GetWebsite() *wrappers.StringValue {
	if x != nil {
		return x.Website
	}
	return nil
}

func (x *AccountResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AccountResponse) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Account Create Request
type AccountCreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string         `protobuf:"bytes,1,opt,name=Name,proto3" json:"name"`
	Employees []*UserRequest `protobuf:"bytes,4,rep,name=Employees,proto3" json:"employees"`
}

func (x *AccountCreateRequest) Reset() {
	*x = AccountCreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_grpc_proto_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountCreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountCreateRequest) ProtoMessage() {}

func (x *AccountCreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_grpc_proto_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountCreateRequest.ProtoReflect.Descriptor instead.
func (*AccountCreateRequest) Descriptor() ([]byte, []int) {
	return file_src_transport_grpc_proto_account_proto_rawDescGZIP(), []int{2}
}

func (x *AccountCreateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountCreateRequest) GetEmployees() []*UserRequest {
	if x != nil {
		return x.Employees
	}
	return nil
}

// Account Create Response
type AccountCreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32               `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	Name      string               `protobuf:"bytes,2,opt,name=Name,proto3" json:"name"`
	Users     []*UserResponse      `protobuf:"bytes,5,rep,name=Users,proto3" json:"users"`
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"created_at"`
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"updated_at"`
}

func (x *AccountCreateResponse) Reset() {
	*x = AccountCreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_grpc_proto_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountCreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountCreateResponse) ProtoMessage() {}

func (x *AccountCreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_grpc_proto_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountCreateResponse.ProtoReflect.Descriptor instead.
func (*AccountCreateResponse) Descriptor() ([]byte, []int) {
	return file_src_transport_grpc_proto_account_proto_rawDescGZIP(), []int{3}
}

func (x *AccountCreateResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *AccountCreateResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *AccountCreateResponse) GetUsers() []*UserResponse {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *AccountCreateResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *AccountCreateResponse) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

// Employee Create Request
type UserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string                `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"first_name"`
	LastName  *wrappers.StringValue `protobuf:"bytes,3,opt,name=LastName,proto3" json:"last_name"`
	Email     string                `protobuf:"bytes,4,opt,name=Email,proto3" json:"email"`
	Password  string                `protobuf:"bytes,5,opt,name=Password,proto3" json:"password"`
}

func (x *UserRequest) Reset() {
	*x = UserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_grpc_proto_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRequest) ProtoMessage() {}

func (x *UserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_grpc_proto_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRequest.ProtoReflect.Descriptor instead.
func (*UserRequest) Descriptor() ([]byte, []int) {
	return file_src_transport_grpc_proto_account_proto_rawDescGZIP(), []int{4}
}

func (x *UserRequest) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserRequest) GetLastName() *wrappers.StringValue {
	if x != nil {
		return x.LastName
	}
	return nil
}

func (x *UserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

// Employee Create Response, Update Request + Response
type UserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID        uint32                `protobuf:"varint,1,opt,name=ID,proto3" json:"id"`
	FirstName string                `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"first_name"`
	LastName  *wrappers.StringValue `protobuf:"bytes,3,opt,name=LastName,proto3" json:"last_name"`
	Email     string                `protobuf:"bytes,4,opt,name=Email,proto3" json:"email"`
	Active    bool                  `protobuf:"varint,5,opt,name=Active,proto3" json:"active"`
	CreatedAt *timestamp.Timestamp  `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"created_at"`
	UpdatedAt *timestamp.Timestamp  `protobuf:"bytes,7,opt,name=UpdatedAt,proto3" json:"updated_at"`
}

func (x *UserResponse) Reset() {
	*x = UserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_transport_grpc_proto_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserResponse) ProtoMessage() {}

func (x *UserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_transport_grpc_proto_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserResponse.ProtoReflect.Descriptor instead.
func (*UserResponse) Descriptor() ([]byte, []int) {
	return file_src_transport_grpc_proto_account_proto_rawDescGZIP(), []int{5}
}

func (x *UserResponse) GetID() uint32 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *UserResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *UserResponse) GetLastName() *wrappers.StringValue {
	if x != nil {
		return x.LastName
	}
	return nil
}

func (x *UserResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserResponse) GetActive() bool {
	if x != nil {
		return x.Active
	}
	return false
}

func (x *UserResponse) GetCreatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *UserResponse) GetUpdatedAt() *timestamp.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_src_transport_grpc_proto_account_proto protoreflect.FileDescriptor

var file_src_transport_grpc_proto_account_proto_rawDesc = []byte{
	0x0a, 0x26, 0x73, 0x72, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x41, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f,
	0x78, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27,
	0x73, 0x72, 0x63, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x72, 0x69, 0x6b, 0x72, 0x73, 0x6e, 0x61, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x74, 0x61, 0x67, 0x2f, 0x74, 0x61,
	0x67, 0x67, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6a, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x15, 0x9a, 0x84, 0x9e, 0x03, 0x09, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x69, 0x64, 0x22, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x02, 0x49, 0x44,
	0x12, 0x2b, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17,
	0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xc8, 0x02,
	0x0a, 0x0f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0e, 0x9a,
	0x84, 0x9e, 0x03, 0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x49,
	0x44, 0x12, 0x24, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x10, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x73, 0x69,
	0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x13, 0x9a, 0x84, 0x9e, 0x03, 0x0e, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x77, 0x65, 0x62, 0x73, 0x69, 0x74, 0x65, 0x22, 0x52, 0x07, 0x57, 0x65, 0x62,
	0x73, 0x69, 0x74, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x42, 0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x09, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x50, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x09, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x99, 0x01, 0x0a, 0x14, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x2b, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x17, 0x9a, 0x84, 0x9e, 0x03, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x54,
	0x0a, 0x09, 0x45, 0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x42, 0x25, 0x9a, 0x84, 0x9e, 0x03, 0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65,
	0x6d, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x65, 0x73, 0x22, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08,
	0x01, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x10, 0x02, 0x52, 0x09, 0x45, 0x6d, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x65, 0x73, 0x22, 0xbc, 0x02, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e,
	0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0e, 0x9a, 0x84, 0x9e, 0x03,
	0x09, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x49, 0x44, 0x12, 0x24,
	0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x10, 0x9a, 0x84,
	0x9e, 0x03, 0x0b, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x11, 0x9a, 0x84, 0x9e, 0x03, 0x0c, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x52, 0x05, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12,
	0x50, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x16,
	0x9a, 0x84, 0x9e, 0x03, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x50, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x42, 0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x92, 0x02, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x3b, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1d, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x56, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x42, 0x1c, 0x9a, 0x84, 0x9e, 0x03, 0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x73,
	0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x02, 0x52, 0x08,
	0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0x9a, 0x84, 0x9e, 0x03, 0x0c, 0x6a, 0x73,
	0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60,
	0x01, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x3e, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73,
	0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x22, 0x9a, 0x84, 0x9e, 0x03,
	0x0f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x08, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x52, 0x08,
	0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0xae, 0x03, 0x0a, 0x0c, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x02, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x0e, 0x9a, 0x84, 0x9e, 0x03, 0x09, 0x6a, 0x73, 0x6f, 0x6e,
	0x3a, 0x22, 0x69, 0x64, 0x22, 0x52, 0x02, 0x49, 0x44, 0x12, 0x34, 0x0a, 0x09, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x16, 0x9a, 0x84,
	0x9e, 0x03, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x52, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x4f, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42,
	0x15, 0x9a, 0x84, 0x9e, 0x03, 0x10, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x27, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x11, 0x9a, 0x84, 0x9e, 0x03, 0x0c, 0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x2a, 0x0a, 0x06, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d,
	0x6a, 0x73, 0x6f, 0x6e, 0x3a, 0x22, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x52, 0x06, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x50, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x42, 0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x22, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x09, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x50, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x16, 0x9a, 0x84, 0x9e, 0x03, 0x11, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x22, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x52, 0x09,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x8f, 0x02, 0x0a, 0x07, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x06, 0x2e, 0x70,
	0x62, 0x2e, 0x49, 0x44, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1b, 0x12, 0x19, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x12, 0x5e, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x01, 0x2a, 0x12, 0x5d, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1e, 0x1a, 0x19, 0x2f,
	0x76, 0x31, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2f, 0x7b, 0x49, 0x44, 0x7d, 0x3a, 0x01, 0x2a, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_src_transport_grpc_proto_account_proto_rawDescOnce sync.Once
	file_src_transport_grpc_proto_account_proto_rawDescData = file_src_transport_grpc_proto_account_proto_rawDesc
)

func file_src_transport_grpc_proto_account_proto_rawDescGZIP() []byte {
	file_src_transport_grpc_proto_account_proto_rawDescOnce.Do(func() {
		file_src_transport_grpc_proto_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_transport_grpc_proto_account_proto_rawDescData)
	})
	return file_src_transport_grpc_proto_account_proto_rawDescData
}

var file_src_transport_grpc_proto_account_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_src_transport_grpc_proto_account_proto_goTypes = []interface{}{
	(*AccountUpdateRequest)(nil),  // 0: pb.AccountUpdateRequest
	(*AccountResponse)(nil),       // 1: pb.AccountResponse
	(*AccountCreateRequest)(nil),  // 2: pb.AccountCreateRequest
	(*AccountCreateResponse)(nil), // 3: pb.AccountCreateResponse
	(*UserRequest)(nil),           // 4: pb.UserRequest
	(*UserResponse)(nil),          // 5: pb.UserResponse
	(*wrappers.StringValue)(nil),  // 6: google.protobuf.StringValue
	(*timestamp.Timestamp)(nil),   // 7: google.protobuf.Timestamp
	(*ID)(nil),                    // 8: pb.ID
}
var file_src_transport_grpc_proto_account_proto_depIdxs = []int32{
	6,  // 0: pb.AccountResponse.Website:type_name -> google.protobuf.StringValue
	7,  // 1: pb.AccountResponse.CreatedAt:type_name -> google.protobuf.Timestamp
	7,  // 2: pb.AccountResponse.UpdatedAt:type_name -> google.protobuf.Timestamp
	4,  // 3: pb.AccountCreateRequest.Employees:type_name -> pb.UserRequest
	5,  // 4: pb.AccountCreateResponse.Users:type_name -> pb.UserResponse
	7,  // 5: pb.AccountCreateResponse.CreatedAt:type_name -> google.protobuf.Timestamp
	7,  // 6: pb.AccountCreateResponse.UpdatedAt:type_name -> google.protobuf.Timestamp
	6,  // 7: pb.UserRequest.LastName:type_name -> google.protobuf.StringValue
	6,  // 8: pb.UserResponse.LastName:type_name -> google.protobuf.StringValue
	7,  // 9: pb.UserResponse.CreatedAt:type_name -> google.protobuf.Timestamp
	7,  // 10: pb.UserResponse.UpdatedAt:type_name -> google.protobuf.Timestamp
	8,  // 11: pb.Account.Get:input_type -> pb.ID
	2,  // 12: pb.Account.Create:input_type -> pb.AccountCreateRequest
	0,  // 13: pb.Account.Update:input_type -> pb.AccountUpdateRequest
	1,  // 14: pb.Account.Get:output_type -> pb.AccountResponse
	3,  // 15: pb.Account.Create:output_type -> pb.AccountCreateResponse
	1,  // 16: pb.Account.Update:output_type -> pb.AccountResponse
	14, // [14:17] is the sub-list for method output_type
	11, // [11:14] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_src_transport_grpc_proto_account_proto_init() }
func file_src_transport_grpc_proto_account_proto_init() {
	if File_src_transport_grpc_proto_account_proto != nil {
		return
	}
	file_src_transport_grpc_proto_generics_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_src_transport_grpc_proto_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountUpdateRequest); i {
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
		file_src_transport_grpc_proto_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountResponse); i {
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
		file_src_transport_grpc_proto_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountCreateRequest); i {
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
		file_src_transport_grpc_proto_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountCreateResponse); i {
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
		file_src_transport_grpc_proto_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRequest); i {
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
		file_src_transport_grpc_proto_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserResponse); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_transport_grpc_proto_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_transport_grpc_proto_account_proto_goTypes,
		DependencyIndexes: file_src_transport_grpc_proto_account_proto_depIdxs,
		MessageInfos:      file_src_transport_grpc_proto_account_proto_msgTypes,
	}.Build()
	File_src_transport_grpc_proto_account_proto = out.File
	file_src_transport_grpc_proto_account_proto_rawDesc = nil
	file_src_transport_grpc_proto_account_proto_goTypes = nil
	file_src_transport_grpc_proto_account_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AccountClient is the client API for Account service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AccountClient interface {
	// Simple return the business by id
	Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AccountResponse, error)
	// Create Account & Offices & Main Employee
	Create(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*AccountCreateResponse, error)
	// Update Account
	Update(ctx context.Context, in *AccountUpdateRequest, opts ...grpc.CallOption) (*AccountResponse, error)
}

type accountClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountClient(cc grpc.ClientConnInterface) AccountClient {
	return &accountClient{cc}
}

func (c *accountClient) Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/pb.Account/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Create(ctx context.Context, in *AccountCreateRequest, opts ...grpc.CallOption) (*AccountCreateResponse, error) {
	out := new(AccountCreateResponse)
	err := c.cc.Invoke(ctx, "/pb.Account/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountClient) Update(ctx context.Context, in *AccountUpdateRequest, opts ...grpc.CallOption) (*AccountResponse, error) {
	out := new(AccountResponse)
	err := c.cc.Invoke(ctx, "/pb.Account/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServer is the server API for Account service.
type AccountServer interface {
	// Simple return the business by id
	Get(context.Context, *ID) (*AccountResponse, error)
	// Create Account & Offices & Main Employee
	Create(context.Context, *AccountCreateRequest) (*AccountCreateResponse, error)
	// Update Account
	Update(context.Context, *AccountUpdateRequest) (*AccountResponse, error)
}

// UnimplementedAccountServer can be embedded to have forward compatible implementations.
type UnimplementedAccountServer struct {
}

func (*UnimplementedAccountServer) Get(context.Context, *ID) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedAccountServer) Create(context.Context, *AccountCreateRequest) (*AccountCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (*UnimplementedAccountServer) Update(context.Context, *AccountUpdateRequest) (*AccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func RegisterAccountServer(s *grpc.Server, srv AccountServer) {
	s.RegisterService(&_Account_serviceDesc, srv)
}

func _Account_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Get(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Create(ctx, req.(*AccountCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Account_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccountUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Account/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServer).Update(ctx, req.(*AccountUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Account_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Account",
	HandlerType: (*AccountServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Account_Get_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _Account_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Account_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/transport/grpc/proto/account.proto",
}