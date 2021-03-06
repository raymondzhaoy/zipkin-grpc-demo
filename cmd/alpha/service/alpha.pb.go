// Code generated by protoc-gen-go.
// source: service/alpha.proto
// DO NOT EDIT!

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	service/alpha.proto

It has these top-level messages:
	GetRequest
	GetResponse
*/
package service

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

type GetRequest struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetResponse struct {
	Data string `protobuf:"bytes,1,opt,name=data" json:"data,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*GetRequest)(nil), "alpha.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "alpha.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for AlphaService service

type AlphaServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type alphaServiceClient struct {
	cc *grpc.ClientConn
}

func NewAlphaServiceClient(cc *grpc.ClientConn) AlphaServiceClient {
	return &alphaServiceClient{cc}
}

func (c *alphaServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := grpc.Invoke(ctx, "/alpha.AlphaService/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AlphaService service

type AlphaServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
}

func RegisterAlphaServiceServer(s *grpc.Server, srv AlphaServiceServer) {
	s.RegisterService(&_AlphaService_serviceDesc, srv)
}

func _AlphaService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AlphaServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpha.AlphaService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AlphaServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AlphaService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alpha.AlphaService",
	HandlerType: (*AlphaServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AlphaService_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("service/alpha.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 136 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0xcc, 0x29, 0xc8, 0x48, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x62, 0x05, 0x73, 0x94, 0x14, 0xb8, 0xb8, 0xdc, 0x53, 0x4b, 0x82, 0x52, 0x0b, 0x4b, 0x53, 0x8b,
	0x4b, 0x84, 0x84, 0xb8, 0x58, 0x52, 0x12, 0x4b, 0x12, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xc0, 0x6c, 0x25, 0x45, 0x2e, 0x6e, 0xb0, 0x8a, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x6c, 0x4a,
	0x8c, 0x6c, 0xb8, 0x78, 0x1c, 0x41, 0xa6, 0x05, 0x43, 0xec, 0x11, 0xd2, 0xe1, 0x62, 0x76, 0x4f,
	0x2d, 0x11, 0x12, 0xd4, 0x83, 0x58, 0x88, 0xb0, 0x40, 0x4a, 0x08, 0x59, 0x08, 0x62, 0xa2, 0x13,
	0x67, 0x14, 0x3b, 0xd4, 0x81, 0x49, 0x6c, 0x60, 0xb7, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xd9, 0x16, 0xc2, 0x54, 0xb2, 0x00, 0x00, 0x00,
}
