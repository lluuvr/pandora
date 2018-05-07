// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tracker.proto

package tracker

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_9ef0c7a43135abbf, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type BrokerOpts struct {
	Endpoint             string   `protobuf:"bytes,1,opt,name=endpoint" json:"endpoint,omitempty"`
	User                 string   `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Password             string   `protobuf:"bytes,3,opt,name=password" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BrokerOpts) Reset()         { *m = BrokerOpts{} }
func (m *BrokerOpts) String() string { return proto.CompactTextString(m) }
func (*BrokerOpts) ProtoMessage()    {}
func (*BrokerOpts) Descriptor() ([]byte, []int) {
	return fileDescriptor_tracker_9ef0c7a43135abbf, []int{1}
}
func (m *BrokerOpts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BrokerOpts.Unmarshal(m, b)
}
func (m *BrokerOpts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BrokerOpts.Marshal(b, m, deterministic)
}
func (dst *BrokerOpts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BrokerOpts.Merge(dst, src)
}
func (m *BrokerOpts) XXX_Size() int {
	return xxx_messageInfo_BrokerOpts.Size(m)
}
func (m *BrokerOpts) XXX_DiscardUnknown() {
	xxx_messageInfo_BrokerOpts.DiscardUnknown(m)
}

var xxx_messageInfo_BrokerOpts proto.InternalMessageInfo

func (m *BrokerOpts) GetEndpoint() string {
	if m != nil {
		return m.Endpoint
	}
	return ""
}

func (m *BrokerOpts) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *BrokerOpts) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func init() {
	proto.RegisterType((*Empty)(nil), "tracker.Empty")
	proto.RegisterType((*BrokerOpts)(nil), "tracker.BrokerOpts")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Tracker service

type TrackerClient interface {
	GetBrokerOpts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BrokerOpts, error)
}

type trackerClient struct {
	cc *grpc.ClientConn
}

func NewTrackerClient(cc *grpc.ClientConn) TrackerClient {
	return &trackerClient{cc}
}

func (c *trackerClient) GetBrokerOpts(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BrokerOpts, error) {
	out := new(BrokerOpts)
	err := grpc.Invoke(ctx, "/tracker.Tracker/GetBrokerOpts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Tracker service

type TrackerServer interface {
	GetBrokerOpts(context.Context, *Empty) (*BrokerOpts, error)
}

func RegisterTrackerServer(s *grpc.Server, srv TrackerServer) {
	s.RegisterService(&_Tracker_serviceDesc, srv)
}

func _Tracker_GetBrokerOpts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TrackerServer).GetBrokerOpts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tracker.Tracker/GetBrokerOpts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TrackerServer).GetBrokerOpts(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tracker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tracker.Tracker",
	HandlerType: (*TrackerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBrokerOpts",
			Handler:    _Tracker_GetBrokerOpts_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracker.proto",
}

func init() { proto.RegisterFile("tracker.proto", fileDescriptor_tracker_9ef0c7a43135abbf) }

var fileDescriptor_tracker_9ef0c7a43135abbf = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x29, 0x4a, 0x4c,
	0xce, 0x4e, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xd8, 0xb9,
	0x58, 0x5d, 0x73, 0x0b, 0x4a, 0x2a, 0x95, 0x22, 0xb8, 0xb8, 0x9c, 0x8a, 0xf2, 0xb3, 0x53, 0x8b,
	0xfc, 0x0b, 0x4a, 0x8a, 0x85, 0xa4, 0xb8, 0x38, 0x52, 0xf3, 0x52, 0x0a, 0xf2, 0x33, 0xf3, 0x4a,
	0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x21, 0x21, 0x2e, 0x96, 0xd2, 0xe2, 0xd4,
	0x22, 0x09, 0x26, 0xb0, 0x38, 0x98, 0x0d, 0x52, 0x5f, 0x90, 0x58, 0x5c, 0x5c, 0x9e, 0x5f, 0x94,
	0x22, 0xc1, 0x0c, 0x51, 0x0f, 0xe3, 0x1b, 0x39, 0x72, 0xb1, 0x87, 0x40, 0x6c, 0x13, 0x32, 0xe3,
	0xe2, 0x75, 0x4f, 0x2d, 0x41, 0xb2, 0x87, 0x4f, 0x0f, 0xe6, 0x2e, 0xb0, 0x2b, 0xa4, 0x84, 0xe1,
	0x7c, 0x84, 0x22, 0x25, 0x86, 0x24, 0x36, 0xb0, 0xab, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff,
	0xfa, 0x75, 0x92, 0xaf, 0xc6, 0x00, 0x00, 0x00,
}
