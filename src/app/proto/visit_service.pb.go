// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/app/proto/visit_service.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ID struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ID) Reset()         { *m = ID{} }
func (m *ID) String() string { return proto.CompactTextString(m) }
func (*ID) ProtoMessage()    {}
func (*ID) Descriptor() ([]byte, []int) {
	return fileDescriptor_visit_service_63e4c4ba5e3c46d5, []int{0}
}
func (m *ID) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ID.Unmarshal(m, b)
}
func (m *ID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ID.Marshal(b, m, deterministic)
}
func (dst *ID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ID.Merge(dst, src)
}
func (m *ID) XXX_Size() int {
	return xxx_messageInfo_ID.Size(m)
}
func (m *ID) XXX_DiscardUnknown() {
	xxx_messageInfo_ID.DiscardUnknown(m)
}

var xxx_messageInfo_ID proto.InternalMessageInfo

func (m *ID) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type Visit struct {
	Id                   int32                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt            *timestamp.Timestamp `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Visit) Reset()         { *m = Visit{} }
func (m *Visit) String() string { return proto.CompactTextString(m) }
func (*Visit) ProtoMessage()    {}
func (*Visit) Descriptor() ([]byte, []int) {
	return fileDescriptor_visit_service_63e4c4ba5e3c46d5, []int{1}
}
func (m *Visit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Visit.Unmarshal(m, b)
}
func (m *Visit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Visit.Marshal(b, m, deterministic)
}
func (dst *Visit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Visit.Merge(dst, src)
}
func (m *Visit) XXX_Size() int {
	return xxx_messageInfo_Visit.Size(m)
}
func (m *Visit) XXX_DiscardUnknown() {
	xxx_messageInfo_Visit.DiscardUnknown(m)
}

var xxx_messageInfo_Visit proto.InternalMessageInfo

func (m *Visit) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Visit) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func init() {
	proto.RegisterType((*ID)(nil), "pb.ID")
	proto.RegisterType((*Visit)(nil), "pb.Visit")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VisitServiceClient is the client API for VisitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VisitServiceClient interface {
	// Simple return the visit id
	Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Visit, error)
	// Update/Create a device
	Set(ctx context.Context, in *Visit, opts ...grpc.CallOption) (*Visit, error)
}

type visitServiceClient struct {
	cc *grpc.ClientConn
}

func NewVisitServiceClient(cc *grpc.ClientConn) VisitServiceClient {
	return &visitServiceClient{cc}
}

func (c *visitServiceClient) Get(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Visit, error) {
	out := new(Visit)
	err := c.cc.Invoke(ctx, "/pb.VisitService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *visitServiceClient) Set(ctx context.Context, in *Visit, opts ...grpc.CallOption) (*Visit, error) {
	out := new(Visit)
	err := c.cc.Invoke(ctx, "/pb.VisitService/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VisitServiceServer is the server API for VisitService service.
type VisitServiceServer interface {
	// Simple return the visit id
	Get(context.Context, *ID) (*Visit, error)
	// Update/Create a device
	Set(context.Context, *Visit) (*Visit, error)
}

func RegisterVisitServiceServer(s *grpc.Server, srv VisitServiceServer) {
	s.RegisterService(&_VisitService_serviceDesc, srv)
}

func _VisitService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisitServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VisitService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisitServiceServer).Get(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VisitService_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Visit)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VisitServiceServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.VisitService/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VisitServiceServer).Set(ctx, req.(*Visit))
	}
	return interceptor(ctx, in, info, handler)
}

var _VisitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.VisitService",
	HandlerType: (*VisitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _VisitService_Get_Handler,
		},
		{
			MethodName: "Set",
			Handler:    _VisitService_Set_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/app/proto/visit_service.proto",
}

func init() {
	proto.RegisterFile("src/app/proto/visit_service.proto", fileDescriptor_visit_service_63e4c4ba5e3c46d5)
}

var fileDescriptor_visit_service_63e4c4ba5e3c46d5 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x2e, 0x4a, 0xd6,
	0x4f, 0x2c, 0x28, 0xd0, 0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x2f, 0xcb, 0x2c, 0xce, 0x2c, 0x89,
	0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x03, 0x8b, 0x09, 0x31, 0x15, 0x24, 0x49, 0xc9,
	0xa7, 0xe7, 0xe7, 0xa7, 0xe7, 0xa4, 0x42, 0x54, 0x25, 0x95, 0xa6, 0xe9, 0x97, 0x64, 0xe6, 0xa6,
	0x16, 0x97, 0x24, 0xe6, 0x16, 0x40, 0x14, 0x49, 0xc9, 0x40, 0x15, 0x24, 0x16, 0x64, 0xea, 0x27,
	0xe6, 0xe5, 0xe5, 0x97, 0x24, 0x96, 0x64, 0xe6, 0xe7, 0x15, 0x43, 0x64, 0x95, 0x44, 0xb8, 0x98,
	0x3c, 0x5d, 0x84, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x98,
	0x32, 0x53, 0x94, 0x82, 0xb8, 0x58, 0xc3, 0x40, 0xf6, 0xa1, 0x4b, 0x08, 0x59, 0x72, 0x71, 0x25,
	0x17, 0xa5, 0x26, 0x96, 0xa4, 0xa6, 0xc4, 0x27, 0x96, 0x48, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x1b,
	0x49, 0xe9, 0x41, 0x6c, 0xd0, 0x83, 0x39, 0x41, 0x2f, 0x04, 0xe6, 0x84, 0x20, 0x4e, 0xa8, 0x6a,
	0xc7, 0x12, 0xa3, 0x78, 0x2e, 0x1e, 0xb0, 0x99, 0xc1, 0x10, 0x2f, 0x08, 0x19, 0x70, 0x31, 0xbb,
	0xa7, 0x96, 0x08, 0xb1, 0xe9, 0x15, 0x24, 0xe9, 0x79, 0xba, 0x48, 0x71, 0x82, 0x68, 0xb0, 0x02,
	0x25, 0xb1, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x09, 0x08, 0xf1, 0xe9, 0x97, 0x19, 0x42, 0xfc, 0xad,
	0x5f, 0x9d, 0x99, 0x52, 0x2b, 0x24, 0xcb, 0xc5, 0x1c, 0x9c, 0x5a, 0x22, 0x84, 0x50, 0x89, 0xac,
	0x89, 0x21, 0x89, 0x0d, 0x6c, 0xbf, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xe2, 0xfc, 0x96, 0xcd,
	0x39, 0x01, 0x00, 0x00,
}
