// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

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

type PingMessage struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8341cd3494390082, []int{0}
}
func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (dst *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(dst, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type GetElevatorStatusRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetElevatorStatusRequest) Reset()         { *m = GetElevatorStatusRequest{} }
func (m *GetElevatorStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetElevatorStatusRequest) ProtoMessage()    {}
func (*GetElevatorStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8341cd3494390082, []int{1}
}
func (m *GetElevatorStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetElevatorStatusRequest.Unmarshal(m, b)
}
func (m *GetElevatorStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetElevatorStatusRequest.Marshal(b, m, deterministic)
}
func (dst *GetElevatorStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetElevatorStatusRequest.Merge(dst, src)
}
func (m *GetElevatorStatusRequest) XXX_Size() int {
	return xxx_messageInfo_GetElevatorStatusRequest.Size(m)
}
func (m *GetElevatorStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetElevatorStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetElevatorStatusRequest proto.InternalMessageInfo

type GetElevatorStatusResponse struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Direction            string   `protobuf:"bytes,2,opt,name=direction,proto3" json:"direction,omitempty"`
	Floor                int32    `protobuf:"varint,3,opt,name=floor,proto3" json:"floor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetElevatorStatusResponse) Reset()         { *m = GetElevatorStatusResponse{} }
func (m *GetElevatorStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetElevatorStatusResponse) ProtoMessage()    {}
func (*GetElevatorStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8341cd3494390082, []int{2}
}
func (m *GetElevatorStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetElevatorStatusResponse.Unmarshal(m, b)
}
func (m *GetElevatorStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetElevatorStatusResponse.Marshal(b, m, deterministic)
}
func (dst *GetElevatorStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetElevatorStatusResponse.Merge(dst, src)
}
func (m *GetElevatorStatusResponse) XXX_Size() int {
	return xxx_messageInfo_GetElevatorStatusResponse.Size(m)
}
func (m *GetElevatorStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetElevatorStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetElevatorStatusResponse proto.InternalMessageInfo

func (m *GetElevatorStatusResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetElevatorStatusResponse) GetDirection() string {
	if m != nil {
		return m.Direction
	}
	return ""
}

func (m *GetElevatorStatusResponse) GetFloor() int32 {
	if m != nil {
		return m.Floor
	}
	return 0
}

type ElevatorUpRequest struct {
	Destination          int32    `protobuf:"varint,1,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElevatorUpRequest) Reset()         { *m = ElevatorUpRequest{} }
func (m *ElevatorUpRequest) String() string { return proto.CompactTextString(m) }
func (*ElevatorUpRequest) ProtoMessage()    {}
func (*ElevatorUpRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8341cd3494390082, []int{3}
}
func (m *ElevatorUpRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElevatorUpRequest.Unmarshal(m, b)
}
func (m *ElevatorUpRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElevatorUpRequest.Marshal(b, m, deterministic)
}
func (dst *ElevatorUpRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElevatorUpRequest.Merge(dst, src)
}
func (m *ElevatorUpRequest) XXX_Size() int {
	return xxx_messageInfo_ElevatorUpRequest.Size(m)
}
func (m *ElevatorUpRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ElevatorUpRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ElevatorUpRequest proto.InternalMessageInfo

func (m *ElevatorUpRequest) GetDestination() int32 {
	if m != nil {
		return m.Destination
	}
	return 0
}

type ElevatorUpResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElevatorUpResponse) Reset()         { *m = ElevatorUpResponse{} }
func (m *ElevatorUpResponse) String() string { return proto.CompactTextString(m) }
func (*ElevatorUpResponse) ProtoMessage()    {}
func (*ElevatorUpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8341cd3494390082, []int{4}
}
func (m *ElevatorUpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElevatorUpResponse.Unmarshal(m, b)
}
func (m *ElevatorUpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElevatorUpResponse.Marshal(b, m, deterministic)
}
func (dst *ElevatorUpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElevatorUpResponse.Merge(dst, src)
}
func (m *ElevatorUpResponse) XXX_Size() int {
	return xxx_messageInfo_ElevatorUpResponse.Size(m)
}
func (m *ElevatorUpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ElevatorUpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ElevatorUpResponse proto.InternalMessageInfo

type ElevatorDownRequest struct {
	Destination          int32    `protobuf:"varint,1,opt,name=destination,proto3" json:"destination,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElevatorDownRequest) Reset()         { *m = ElevatorDownRequest{} }
func (m *ElevatorDownRequest) String() string { return proto.CompactTextString(m) }
func (*ElevatorDownRequest) ProtoMessage()    {}
func (*ElevatorDownRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8341cd3494390082, []int{5}
}
func (m *ElevatorDownRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElevatorDownRequest.Unmarshal(m, b)
}
func (m *ElevatorDownRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElevatorDownRequest.Marshal(b, m, deterministic)
}
func (dst *ElevatorDownRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElevatorDownRequest.Merge(dst, src)
}
func (m *ElevatorDownRequest) XXX_Size() int {
	return xxx_messageInfo_ElevatorDownRequest.Size(m)
}
func (m *ElevatorDownRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ElevatorDownRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ElevatorDownRequest proto.InternalMessageInfo

func (m *ElevatorDownRequest) GetDestination() int32 {
	if m != nil {
		return m.Destination
	}
	return 0
}

type ElevatorDownResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElevatorDownResponse) Reset()         { *m = ElevatorDownResponse{} }
func (m *ElevatorDownResponse) String() string { return proto.CompactTextString(m) }
func (*ElevatorDownResponse) ProtoMessage()    {}
func (*ElevatorDownResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8341cd3494390082, []int{6}
}
func (m *ElevatorDownResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElevatorDownResponse.Unmarshal(m, b)
}
func (m *ElevatorDownResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElevatorDownResponse.Marshal(b, m, deterministic)
}
func (dst *ElevatorDownResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElevatorDownResponse.Merge(dst, src)
}
func (m *ElevatorDownResponse) XXX_Size() int {
	return xxx_messageInfo_ElevatorDownResponse.Size(m)
}
func (m *ElevatorDownResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ElevatorDownResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ElevatorDownResponse proto.InternalMessageInfo

type ElevatorStopRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElevatorStopRequest) Reset()         { *m = ElevatorStopRequest{} }
func (m *ElevatorStopRequest) String() string { return proto.CompactTextString(m) }
func (*ElevatorStopRequest) ProtoMessage()    {}
func (*ElevatorStopRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8341cd3494390082, []int{7}
}
func (m *ElevatorStopRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElevatorStopRequest.Unmarshal(m, b)
}
func (m *ElevatorStopRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElevatorStopRequest.Marshal(b, m, deterministic)
}
func (dst *ElevatorStopRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElevatorStopRequest.Merge(dst, src)
}
func (m *ElevatorStopRequest) XXX_Size() int {
	return xxx_messageInfo_ElevatorStopRequest.Size(m)
}
func (m *ElevatorStopRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ElevatorStopRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ElevatorStopRequest proto.InternalMessageInfo

type ElevatorStopResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ElevatorStopResponse) Reset()         { *m = ElevatorStopResponse{} }
func (m *ElevatorStopResponse) String() string { return proto.CompactTextString(m) }
func (*ElevatorStopResponse) ProtoMessage()    {}
func (*ElevatorStopResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_api_8341cd3494390082, []int{8}
}
func (m *ElevatorStopResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ElevatorStopResponse.Unmarshal(m, b)
}
func (m *ElevatorStopResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ElevatorStopResponse.Marshal(b, m, deterministic)
}
func (dst *ElevatorStopResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ElevatorStopResponse.Merge(dst, src)
}
func (m *ElevatorStopResponse) XXX_Size() int {
	return xxx_messageInfo_ElevatorStopResponse.Size(m)
}
func (m *ElevatorStopResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ElevatorStopResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ElevatorStopResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PingMessage)(nil), "api.PingMessage")
	proto.RegisterType((*GetElevatorStatusRequest)(nil), "api.getElevatorStatusRequest")
	proto.RegisterType((*GetElevatorStatusResponse)(nil), "api.getElevatorStatusResponse")
	proto.RegisterType((*ElevatorUpRequest)(nil), "api.elevatorUpRequest")
	proto.RegisterType((*ElevatorUpResponse)(nil), "api.elevatorUpResponse")
	proto.RegisterType((*ElevatorDownRequest)(nil), "api.elevatorDownRequest")
	proto.RegisterType((*ElevatorDownResponse)(nil), "api.elevatorDownResponse")
	proto.RegisterType((*ElevatorStopRequest)(nil), "api.elevatorStopRequest")
	proto.RegisterType((*ElevatorStopResponse)(nil), "api.elevatorStopResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PingClient is the client API for Ping service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PingClient interface {
	SayHello(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessage, error)
}

type pingClient struct {
	cc *grpc.ClientConn
}

func NewPingClient(cc *grpc.ClientConn) PingClient {
	return &pingClient{cc}
}

func (c *pingClient) SayHello(ctx context.Context, in *PingMessage, opts ...grpc.CallOption) (*PingMessage, error) {
	out := new(PingMessage)
	err := c.cc.Invoke(ctx, "/api.Ping/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServer is the server API for Ping service.
type PingServer interface {
	SayHello(context.Context, *PingMessage) (*PingMessage, error)
}

func RegisterPingServer(s *grpc.Server, srv PingServer) {
	s.RegisterService(&_Ping_serviceDesc, srv)
}

func _Ping_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Ping/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).SayHello(ctx, req.(*PingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ping_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Ping",
	HandlerType: (*PingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Ping_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// ElevatorServiceClient is the client API for ElevatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ElevatorServiceClient interface {
	GetElevatorStatus(ctx context.Context, in *GetElevatorStatusRequest, opts ...grpc.CallOption) (*GetElevatorStatusResponse, error)
	ElevatorUp(ctx context.Context, in *ElevatorUpRequest, opts ...grpc.CallOption) (*ElevatorUpResponse, error)
	ElevatorDown(ctx context.Context, in *ElevatorDownRequest, opts ...grpc.CallOption) (*ElevatorDownResponse, error)
	ElevatorStop(ctx context.Context, in *ElevatorStopRequest, opts ...grpc.CallOption) (*ElevatorStopResponse, error)
}

type elevatorServiceClient struct {
	cc *grpc.ClientConn
}

func NewElevatorServiceClient(cc *grpc.ClientConn) ElevatorServiceClient {
	return &elevatorServiceClient{cc}
}

func (c *elevatorServiceClient) GetElevatorStatus(ctx context.Context, in *GetElevatorStatusRequest, opts ...grpc.CallOption) (*GetElevatorStatusResponse, error) {
	out := new(GetElevatorStatusResponse)
	err := c.cc.Invoke(ctx, "/api.ElevatorService/GetElevatorStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elevatorServiceClient) ElevatorUp(ctx context.Context, in *ElevatorUpRequest, opts ...grpc.CallOption) (*ElevatorUpResponse, error) {
	out := new(ElevatorUpResponse)
	err := c.cc.Invoke(ctx, "/api.ElevatorService/ElevatorUp", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elevatorServiceClient) ElevatorDown(ctx context.Context, in *ElevatorDownRequest, opts ...grpc.CallOption) (*ElevatorDownResponse, error) {
	out := new(ElevatorDownResponse)
	err := c.cc.Invoke(ctx, "/api.ElevatorService/ElevatorDown", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elevatorServiceClient) ElevatorStop(ctx context.Context, in *ElevatorStopRequest, opts ...grpc.CallOption) (*ElevatorStopResponse, error) {
	out := new(ElevatorStopResponse)
	err := c.cc.Invoke(ctx, "/api.ElevatorService/ElevatorStop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ElevatorServiceServer is the server API for ElevatorService service.
type ElevatorServiceServer interface {
	GetElevatorStatus(context.Context, *GetElevatorStatusRequest) (*GetElevatorStatusResponse, error)
	ElevatorUp(context.Context, *ElevatorUpRequest) (*ElevatorUpResponse, error)
	ElevatorDown(context.Context, *ElevatorDownRequest) (*ElevatorDownResponse, error)
	ElevatorStop(context.Context, *ElevatorStopRequest) (*ElevatorStopResponse, error)
}

func RegisterElevatorServiceServer(s *grpc.Server, srv ElevatorServiceServer) {
	s.RegisterService(&_ElevatorService_serviceDesc, srv)
}

func _ElevatorService_GetElevatorStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetElevatorStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElevatorServiceServer).GetElevatorStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ElevatorService/GetElevatorStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElevatorServiceServer).GetElevatorStatus(ctx, req.(*GetElevatorStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElevatorService_ElevatorUp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElevatorUpRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElevatorServiceServer).ElevatorUp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ElevatorService/ElevatorUp",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElevatorServiceServer).ElevatorUp(ctx, req.(*ElevatorUpRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElevatorService_ElevatorDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElevatorDownRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElevatorServiceServer).ElevatorDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ElevatorService/ElevatorDown",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElevatorServiceServer).ElevatorDown(ctx, req.(*ElevatorDownRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ElevatorService_ElevatorStop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ElevatorStopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ElevatorServiceServer).ElevatorStop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ElevatorService/ElevatorStop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ElevatorServiceServer).ElevatorStop(ctx, req.(*ElevatorStopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ElevatorService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ElevatorService",
	HandlerType: (*ElevatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetElevatorStatus",
			Handler:    _ElevatorService_GetElevatorStatus_Handler,
		},
		{
			MethodName: "ElevatorUp",
			Handler:    _ElevatorService_ElevatorUp_Handler,
		},
		{
			MethodName: "ElevatorDown",
			Handler:    _ElevatorService_ElevatorDown_Handler,
		},
		{
			MethodName: "ElevatorStop",
			Handler:    _ElevatorService_ElevatorStop_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor_api_8341cd3494390082) }

var fileDescriptor_api_8341cd3494390082 = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x4d, 0x4f, 0xf2, 0x40,
	0x10, 0xc7, 0x29, 0x2f, 0x4f, 0x60, 0x78, 0x12, 0x65, 0x44, 0x2c, 0x8d, 0x1a, 0xb2, 0x27, 0xbc,
	0x10, 0x83, 0x31, 0x7a, 0xf3, 0x62, 0xa3, 0x17, 0x13, 0x03, 0xfa, 0x01, 0x56, 0x18, 0x9b, 0x4d,
	0xea, 0x6e, 0xed, 0x2e, 0x18, 0xbf, 0x98, 0x9f, 0xcf, 0xb0, 0x2d, 0x65, 0x79, 0x4b, 0xbc, 0xb1,
	0xff, 0x99, 0xf9, 0xed, 0xf2, 0x9b, 0x42, 0x83, 0x27, 0x62, 0x90, 0xa4, 0xca, 0x28, 0xac, 0xf0,
	0x44, 0xb0, 0x0b, 0x68, 0x3e, 0x0b, 0x19, 0x3d, 0x91, 0xd6, 0x3c, 0x22, 0x0c, 0xa0, 0x1e, 0xa5,
	0x44, 0x46, 0xc8, 0xc8, 0xf7, 0x7a, 0x5e, 0xbf, 0x31, 0x2a, 0xce, 0x2c, 0x00, 0x3f, 0x22, 0x13,
	0xc6, 0x34, 0xe7, 0x46, 0xa5, 0x63, 0xc3, 0xcd, 0x4c, 0x8f, 0xe8, 0x73, 0x46, 0xda, 0xb0, 0x09,
	0x74, 0x77, 0xd4, 0x74, 0xa2, 0xa4, 0x26, 0x44, 0xa8, 0x4a, 0xfe, 0x41, 0x39, 0xd0, 0xfe, 0xc6,
	0x53, 0x68, 0x4c, 0x45, 0x4a, 0x13, 0x23, 0x94, 0xf4, 0xcb, 0xb6, 0xb0, 0x0a, 0xb0, 0x0d, 0xb5,
	0xf7, 0x58, 0xa9, 0xd4, 0xaf, 0xf4, 0xbc, 0x7e, 0x6d, 0x94, 0x1d, 0xd8, 0x35, 0xb4, 0x28, 0xbf,
	0xe1, 0x35, 0xc9, 0x6f, 0xc6, 0x1e, 0x34, 0xa7, 0xa4, 0x8d, 0x90, 0xdc, 0xa2, 0x3c, 0x3b, 0xe0,
	0x46, 0xac, 0x0d, 0xe8, 0x8e, 0x65, 0x8f, 0x62, 0x37, 0x70, 0xb4, 0x4c, 0xef, 0xd5, 0x97, 0xfc,
	0x3b, 0xae, 0x03, 0xed, 0xf5, 0xc1, 0x1c, 0x78, 0xbc, 0x02, 0x8e, 0x8d, 0x5a, 0xbe, 0xcf, 0x6d,
	0xcf, 0xe2, 0xac, 0x7d, 0x78, 0x0b, 0xd5, 0x85, 0x78, 0xbc, 0x84, 0xfa, 0x98, 0x7f, 0x3f, 0x52,
	0x1c, 0x2b, 0x3c, 0x1c, 0x2c, 0xb6, 0xe3, 0xec, 0x23, 0xd8, 0x4a, 0x58, 0x69, 0xf8, 0x53, 0x86,
	0x83, 0xc2, 0x34, 0xa5, 0x73, 0x31, 0x21, 0x7c, 0x81, 0xd6, 0xc3, 0xa6, 0x7f, 0x3c, 0xb3, 0xc3,
	0xfb, 0x76, 0x16, 0x9c, 0xef, 0x2b, 0xe7, 0x7f, 0xa8, 0x84, 0x77, 0x00, 0x61, 0x61, 0x0e, 0x3b,
	0xb6, 0x7f, 0x6b, 0x03, 0xc1, 0xc9, 0x56, 0x5e, 0x00, 0x42, 0xf8, 0x1f, 0x3a, 0xae, 0xd0, 0x5f,
	0x6b, 0x75, 0xbc, 0x07, 0xdd, 0x1d, 0x95, 0x5d, 0x98, 0x85, 0xc3, 0x0d, 0x8c, 0x63, 0x7b, 0x03,
	0xe3, 0x0a, 0x67, 0xa5, 0xb7, 0x7f, 0xf6, 0xbb, 0xbf, 0xfa, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xcf,
	0x80, 0x63, 0xdc, 0x04, 0x03, 0x00, 0x00,
}
