// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/type/http.proto

package envoy_type

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type CodecClientType int32

const (
	CodecClientType_HTTP1 CodecClientType = 0
	CodecClientType_HTTP2 CodecClientType = 1
	CodecClientType_HTTP3 CodecClientType = 2
)

var CodecClientType_name = map[int32]string{
	0: "HTTP1",
	1: "HTTP2",
	2: "HTTP3",
}

var CodecClientType_value = map[string]int32{
	"HTTP1": 0,
	"HTTP2": 1,
	"HTTP3": 2,
}

func (x CodecClientType) String() string {
	return proto.EnumName(CodecClientType_name, int32(x))
}

func (CodecClientType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e223f71709aae71c, []int{0}
}

func init() {
	proto.RegisterEnum("envoy.type.CodecClientType", CodecClientType_name, CodecClientType_value)
}

func init() { proto.RegisterFile("envoy/type/http.proto", fileDescriptor_e223f71709aae71c) }

var fileDescriptor_e223f71709aae71c = []byte{
	// 162 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4d, 0xcd, 0x2b, 0xcb,
	0xaf, 0xd4, 0x2f, 0xa9, 0x2c, 0x48, 0xd5, 0xcf, 0x28, 0x29, 0x29, 0xd0, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0xe2, 0x02, 0x0b, 0xeb, 0x81, 0x84, 0xa5, 0x64, 0x4b, 0x53, 0x0a, 0x12, 0xf5, 0x13,
	0xf3, 0xf2, 0xf2, 0x4b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x8a, 0xf5, 0x8b, 0x4b, 0x12, 0x4b, 0x4a,
	0x8b, 0x21, 0x4a, 0xb5, 0x8c, 0xb8, 0xf8, 0x9d, 0xf3, 0x53, 0x52, 0x93, 0x9d, 0x73, 0x32, 0x53,
	0xf3, 0x4a, 0x42, 0x2a, 0x0b, 0x52, 0x85, 0x38, 0xb9, 0x58, 0x3d, 0x42, 0x42, 0x02, 0x0c, 0x05,
	0x18, 0x60, 0x4c, 0x23, 0x01, 0x46, 0x18, 0xd3, 0x58, 0x80, 0xc9, 0x49, 0x7f, 0x57, 0xc3, 0x89,
	0x8b, 0x6c, 0x4c, 0x02, 0x8c, 0x5c, 0x12, 0x99, 0xf9, 0x7a, 0x60, 0xbb, 0x0a, 0x8a, 0xf2, 0x2b,
	0x2a, 0xf5, 0x10, 0xd6, 0x3a, 0x71, 0x7a, 0x94, 0x94, 0x14, 0x04, 0x80, 0xac, 0x08, 0x60, 0x4c,
	0x62, 0x03, 0xdb, 0x65, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xaf, 0xbe, 0x34, 0x37, 0xaf, 0x00,
	0x00, 0x00,
}
