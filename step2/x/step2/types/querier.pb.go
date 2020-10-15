// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: step2/v1beta/querier.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("step2/v1beta/querier.proto", fileDescriptor_51134fcfc5a6b20c) }

var fileDescriptor_51134fcfc5a6b20c = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x2e, 0x49, 0x2d,
	0x30, 0xd4, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0xd4, 0x2f, 0x2c, 0x4d, 0x2d, 0xca, 0x4c, 0x2d,
	0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x06, 0xcb, 0xe9, 0x41, 0x48, 0x88, 0x0a, 0x43,
	0x29, 0xad, 0xe4, 0xfc, 0xe2, 0xdc, 0xfc, 0x62, 0xfd, 0xa4, 0xc4, 0xe2, 0x54, 0xb0, 0xfa, 0x4a,
	0xa8, 0x66, 0x43, 0xfd, 0x82, 0xc4, 0xf4, 0xcc, 0xbc, 0xc4, 0x92, 0xcc, 0xfc, 0x3c, 0x88, 0x01,
	0x46, 0xec, 0x5c, 0xac, 0x81, 0x20, 0x15, 0x4e, 0x6e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24,
	0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78,
	0x2c, 0xc7, 0x10, 0xa5, 0x93, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f,
	0x98, 0x9b, 0x58, 0x5a, 0x54, 0x99, 0x9b, 0x58, 0x54, 0x92, 0x99, 0x57, 0xa9, 0x0f, 0x71, 0x58,
	0x05, 0x94, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x9b, 0x6b, 0x0c, 0x08, 0x00, 0x00,
	0xff, 0xff, 0xa4, 0x3a, 0x37, 0xdd, 0xb6, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "step2.step2.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "step2/v1beta/querier.proto",
}
