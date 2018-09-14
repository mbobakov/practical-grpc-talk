// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/api.proto

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import empty "github.com/golang/protobuf/ptypes/empty"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimeClient is the client API for Time service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimeClient interface {
	// Returns current day length. Server time will be used
	CurrentDayLength(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*duration.Duration, error)
	// Clock is a ticker with second intervale. Server time will be used
	Clock(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Time_ClockClient, error)
}

type timeClient struct {
	cc *grpc.ClientConn
}

func NewTimeClient(cc *grpc.ClientConn) TimeClient {
	return &timeClient{cc}
}

func (c *timeClient) CurrentDayLength(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*duration.Duration, error) {
	out := new(duration.Duration)
	err := c.cc.Invoke(ctx, "/api.Time/CurrentDayLength", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *timeClient) Clock(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (Time_ClockClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Time_serviceDesc.Streams[0], "/api.Time/Clock", opts...)
	if err != nil {
		return nil, err
	}
	x := &timeClockClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Time_ClockClient interface {
	Recv() (*timestamp.Timestamp, error)
	grpc.ClientStream
}

type timeClockClient struct {
	grpc.ClientStream
}

func (x *timeClockClient) Recv() (*timestamp.Timestamp, error) {
	m := new(timestamp.Timestamp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TimeServer is the server API for Time service.
type TimeServer interface {
	// Returns current day length. Server time will be used
	CurrentDayLength(context.Context, *empty.Empty) (*duration.Duration, error)
	// Clock is a ticker with second intervale. Server time will be used
	Clock(*empty.Empty, Time_ClockServer) error
}

func RegisterTimeServer(s *grpc.Server, srv TimeServer) {
	s.RegisterService(&_Time_serviceDesc, srv)
}

func _Time_CurrentDayLength_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServer).CurrentDayLength(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Time/CurrentDayLength",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServer).CurrentDayLength(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Time_Clock_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(empty.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TimeServer).Clock(m, &timeClockServer{stream})
}

type Time_ClockServer interface {
	Send(*timestamp.Timestamp) error
	grpc.ServerStream
}

type timeClockServer struct {
	grpc.ServerStream
}

func (x *timeClockServer) Send(m *timestamp.Timestamp) error {
	return x.ServerStream.SendMsg(m)
}

var _Time_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Time",
	HandlerType: (*TimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CurrentDayLength",
			Handler:    _Time_CurrentDayLength_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Clock",
			Handler:       _Time_Clock_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api/api.proto",
}

func init() { proto.RegisterFile("api/api.proto", fileDescriptor_api_0685bc079bc46d62) }

var fileDescriptor_api_0685bc079bc46d62 = []byte{
	// 172 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0xc8, 0xd4,
	0x4f, 0x2c, 0xc8, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x94, 0x92,
	0x4b, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x07, 0x0b, 0x25, 0x95, 0xa6, 0xe9, 0xa7, 0x94, 0x16,
	0x25, 0x96, 0x64, 0xe6, 0xe7, 0x41, 0x14, 0x49, 0x49, 0xa3, 0xcb, 0xa7, 0xe6, 0x16, 0x94, 0x54,
	0x42, 0x25, 0xe5, 0xd1, 0x25, 0x4b, 0x32, 0x73, 0x53, 0x8b, 0x4b, 0x12, 0x73, 0x0b, 0x20, 0x0a,
	0x8c, 0x7a, 0x18, 0xb9, 0x58, 0x42, 0x32, 0x73, 0x53, 0x85, 0x5c, 0xb9, 0x04, 0x9c, 0x4b, 0x8b,
	0x8a, 0x52, 0xf3, 0x4a, 0x5c, 0x12, 0x2b, 0x7d, 0x52, 0xf3, 0xd2, 0x4b, 0x32, 0x84, 0xc4, 0xf4,
	0x20, 0xda, 0xf5, 0x60, 0xda, 0xf5, 0x5c, 0x41, 0x66, 0x4b, 0x49, 0x62, 0x88, 0xbb, 0x40, 0xdd,
	0x24, 0x64, 0xcb, 0xc5, 0xea, 0x9c, 0x93, 0x9f, 0x9c, 0x8d, 0x53, 0xaf, 0x14, 0x86, 0x78, 0x08,
	0xcc, 0x49, 0x06, 0x8c, 0x4e, 0xac, 0x51, 0x20, 0x3f, 0x27, 0xb1, 0x81, 0x25, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xf6, 0x37, 0x09, 0xa4, 0x10, 0x01, 0x00, 0x00,
}
