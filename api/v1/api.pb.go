// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/v1/api.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import _ "github.com/lyft/protoc-gen-validate/validate"

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

// Passenger for the journey
type Passenger struct {
	// name of Passenger (min.length = 3)
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Passenger) Reset()         { *m = Passenger{} }
func (m *Passenger) String() string { return proto.CompactTextString(m) }
func (*Passenger) ProtoMessage()    {}
func (*Passenger) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_a6fd6faeb012fef5, []int{0}
}
func (m *Passenger) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Passenger.Unmarshal(m, b)
}
func (m *Passenger) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Passenger.Marshal(b, m, deterministic)
}
func (dst *Passenger) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Passenger.Merge(dst, src)
}
func (m *Passenger) XXX_Size() int {
	return xxx_messageInfo_Passenger.Size(m)
}
func (m *Passenger) XXX_DiscardUnknown() {
	xxx_messageInfo_Passenger.DiscardUnknown(m)
}

var xxx_messageInfo_Passenger proto.InternalMessageInfo

func (m *Passenger) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// JumpRequest is a request for the jump throuth time
type JumpRequest struct {
	To *timestamp.Timestamp `protobuf:"bytes,1,opt,name=to,proto3" json:"to,omitempty"`
	// Passengers for the trip (min=2 max=5)
	Passenger            []*Passenger `protobuf:"bytes,2,rep,name=passenger,proto3" json:"passenger,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *JumpRequest) Reset()         { *m = JumpRequest{} }
func (m *JumpRequest) String() string { return proto.CompactTextString(m) }
func (*JumpRequest) ProtoMessage()    {}
func (*JumpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_a6fd6faeb012fef5, []int{1}
}
func (m *JumpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JumpRequest.Unmarshal(m, b)
}
func (m *JumpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JumpRequest.Marshal(b, m, deterministic)
}
func (dst *JumpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JumpRequest.Merge(dst, src)
}
func (m *JumpRequest) XXX_Size() int {
	return xxx_messageInfo_JumpRequest.Size(m)
}
func (m *JumpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_JumpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_JumpRequest proto.InternalMessageInfo

func (m *JumpRequest) GetTo() *timestamp.Timestamp {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *JumpRequest) GetPassenger() []*Passenger {
	if m != nil {
		return m.Passenger
	}
	return nil
}

func init() {
	proto.RegisterType((*Passenger)(nil), "mbobakov.timemachine.api.v1.Passenger")
	proto.RegisterType((*JumpRequest)(nil), "mbobakov.timemachine.api.v1.JumpRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimeMachineClient is the client API for TimeMachine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimeMachineClient interface {
	// Jump to specific moment in time
	Jump(ctx context.Context, in *JumpRequest, opts ...grpc.CallOption) (*duration.Duration, error)
}

type timeMachineClient struct {
	cc *grpc.ClientConn
}

func NewTimeMachineClient(cc *grpc.ClientConn) TimeMachineClient {
	return &timeMachineClient{cc}
}

func (c *timeMachineClient) Jump(ctx context.Context, in *JumpRequest, opts ...grpc.CallOption) (*duration.Duration, error) {
	out := new(duration.Duration)
	err := c.cc.Invoke(ctx, "/mbobakov.timemachine.api.v1.TimeMachine/Jump", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TimeMachineServer is the server API for TimeMachine service.
type TimeMachineServer interface {
	// Jump to specific moment in time
	Jump(context.Context, *JumpRequest) (*duration.Duration, error)
}

func RegisterTimeMachineServer(s *grpc.Server, srv TimeMachineServer) {
	s.RegisterService(&_TimeMachine_serviceDesc, srv)
}

func _TimeMachine_Jump_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JumpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeMachineServer).Jump(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mbobakov.timemachine.api.v1.TimeMachine/Jump",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeMachineServer).Jump(ctx, req.(*JumpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TimeMachine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "mbobakov.timemachine.api.v1.TimeMachine",
	HandlerType: (*TimeMachineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Jump",
			Handler:    _TimeMachine_Jump_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api.proto",
}

func init() { proto.RegisterFile("api/v1/api.proto", fileDescriptor_api_a6fd6faeb012fef5) }

var fileDescriptor_api_a6fd6faeb012fef5 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4e, 0xb4, 0x30,
	0x14, 0x85, 0x43, 0xe1, 0xff, 0x95, 0xe2, 0x82, 0x74, 0xe3, 0x88, 0x51, 0x27, 0xb3, 0x30, 0x64,
	0x16, 0x25, 0xe0, 0x1b, 0x10, 0x57, 0x1a, 0x13, 0xd3, 0xb8, 0x9a, 0xdd, 0xc5, 0xa9, 0xd8, 0x38,
	0xa5, 0x15, 0x4a, 0xdf, 0xc2, 0x8d, 0x8f, 0xe3, 0xca, 0xd7, 0xf1, 0x2d, 0x4c, 0x61, 0x50, 0xa3,
	0xc9, 0xec, 0x9a, 0x9e, 0x73, 0xcf, 0xfd, 0xce, 0xc5, 0x31, 0x68, 0x91, 0xd9, 0x3c, 0x03, 0x2d,
	0xa8, 0x6e, 0x95, 0x51, 0xe4, 0x58, 0x56, 0xaa, 0x82, 0x27, 0x65, 0xa9, 0x11, 0x92, 0x4b, 0xb8,
	0x7f, 0x14, 0x0d, 0xa7, 0x4e, 0xb7, 0x79, 0x72, 0x5a, 0x2b, 0x55, 0x6f, 0x78, 0x36, 0x58, 0xab,
	0xfe, 0x21, 0x5b, 0xf7, 0x2d, 0x18, 0xa1, 0x9a, 0x71, 0x38, 0x39, 0xfb, 0xad, 0xbb, 0x8c, 0xce,
	0x80, 0xd4, 0x5b, 0xc3, 0xa1, 0x85, 0x8d, 0x58, 0x83, 0xe1, 0xd9, 0xf4, 0x18, 0x85, 0xc5, 0x12,
	0x87, 0xb7, 0xd0, 0x75, 0xbc, 0xa9, 0x79, 0x4b, 0x4e, 0x70, 0xd0, 0x80, 0xe4, 0x33, 0x6f, 0xee,
	0xa5, 0x61, 0x19, 0xbe, 0x7d, 0xbc, 0xfb, 0x41, 0x8b, 0x62, 0x9f, 0x0d, 0xdf, 0x8b, 0x17, 0x0f,
	0x47, 0x57, 0xbd, 0xd4, 0x8c, 0x3f, 0xf7, 0xbc, 0x33, 0x64, 0x89, 0x91, 0x51, 0x83, 0x39, 0x2a,
	0x12, 0x3a, 0x22, 0xd0, 0x09, 0x81, 0xde, 0x4d, 0x08, 0x0c, 0x19, 0x45, 0x18, 0x0e, 0xf5, 0xb4,
	0x67, 0x86, 0xe6, 0x7e, 0x1a, 0x15, 0xe7, 0x74, 0x47, 0x65, 0xfa, 0x45, 0x55, 0x1e, 0x38, 0x8e,
	0xbd, 0x57, 0x2f, 0xd8, 0x47, 0xf1, 0x3f, 0xf6, 0x1d, 0x53, 0xac, 0x70, 0xe4, 0x96, 0xdc, 0x8c,
	0x83, 0xe4, 0x1a, 0x07, 0x8e, 0x8e, 0xa4, 0x3b, 0x73, 0x7f, 0x14, 0x48, 0x8e, 0xfe, 0x40, 0x5f,
	0x6e, 0xef, 0x5a, 0x06, 0x2b, 0x64, 0xf3, 0xea, 0xff, 0x20, 0x5c, 0x7c, 0x06, 0x00, 0x00, 0xff,
	0xff, 0x84, 0xb0, 0xe6, 0x93, 0xaf, 0x01, 0x00, 0x00,
}