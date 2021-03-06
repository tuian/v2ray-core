// Code generated by protoc-gen-go.
// source: v2ray.com/core/transport/internet/authenticator.proto
// DO NOT EDIT!

/*
Package internet is a generated protocol buffer package.

It is generated from these files:
	v2ray.com/core/transport/internet/authenticator.proto
	v2ray.com/core/transport/internet/config.proto

It has these top-level messages:
	AuthenticatorConfig
	SecuritySettings
	NetworkSettings
	StreamConfig
*/
package internet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AuthenticatorConfig struct {
	Name     string               `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Settings *google_protobuf.Any `protobuf:"bytes,2,opt,name=settings" json:"settings,omitempty"`
}

func (m *AuthenticatorConfig) Reset()                    { *m = AuthenticatorConfig{} }
func (m *AuthenticatorConfig) String() string            { return proto.CompactTextString(m) }
func (*AuthenticatorConfig) ProtoMessage()               {}
func (*AuthenticatorConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *AuthenticatorConfig) GetSettings() *google_protobuf.Any {
	if m != nil {
		return m.Settings
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthenticatorConfig)(nil), "v2ray.core.transport.internet.AuthenticatorConfig")
}

func init() {
	proto.RegisterFile("v2ray.com/core/transport/internet/authenticator.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x8e, 0xc1, 0x4a, 0xc4, 0x40,
	0x0c, 0x40, 0xa9, 0x88, 0xac, 0xe3, 0x6d, 0xf4, 0xb0, 0x0a, 0xc2, 0xea, 0x69, 0x4f, 0x19, 0xa9,
	0xf8, 0x01, 0xad, 0x3f, 0x20, 0x3d, 0xea, 0x69, 0x3a, 0xa4, 0xe3, 0x80, 0x4d, 0x4a, 0x9a, 0x0a,
	0xfd, 0x7b, 0xb1, 0x65, 0x8a, 0x5e, 0xf6, 0x16, 0xc8, 0x7b, 0x2f, 0x31, 0x2f, 0xdf, 0xa5, 0xf8,
	0x19, 0x02, 0xf7, 0x2e, 0xb0, 0xa0, 0x53, 0xf1, 0x34, 0x0e, 0x2c, 0xea, 0x12, 0x29, 0x0a, 0xa1,
	0x3a, 0x3f, 0xe9, 0x27, 0x92, 0xa6, 0xe0, 0x95, 0x05, 0x06, 0x61, 0x65, 0x7b, 0x9f, 0x35, 0x41,
	0xd8, 0x14, 0xc8, 0xca, 0xdd, 0x6d, 0x64, 0x8e, 0x5f, 0xe8, 0x16, 0xb8, 0x9d, 0x3a, 0xe7, 0x69,
	0x5e, 0xcd, 0xc7, 0x0f, 0x73, 0x5d, 0xfd, 0x0d, 0xbe, 0x32, 0x75, 0x29, 0x5a, 0x6b, 0xce, 0xc9,
	0xf7, 0xb8, 0x2f, 0x0e, 0xc5, 0xf1, 0xb2, 0x59, 0x66, 0xfb, 0x64, 0x76, 0x23, 0xaa, 0x26, 0x8a,
	0xe3, 0xfe, 0xec, 0x50, 0x1c, 0xaf, 0xca, 0x1b, 0x58, 0xc3, 0x90, 0xc3, 0x50, 0xd1, 0xdc, 0x6c,
	0x54, 0x5d, 0x99, 0x87, 0xc0, 0x3d, 0x9c, 0x7c, 0xae, 0xb6, 0xff, 0xee, 0xbf, 0xfd, 0x96, 0xde,
	0x77, 0x79, 0xdb, 0x5e, 0x2c, 0xe9, 0xe7, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x59, 0x3f, 0x6f,
	0x4b, 0x19, 0x01, 0x00, 0x00,
}
