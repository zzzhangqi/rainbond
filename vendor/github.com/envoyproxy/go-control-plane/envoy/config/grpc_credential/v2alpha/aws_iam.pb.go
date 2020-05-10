// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/grpc_credential/v2alpha/aws_iam.proto

package envoy_config_grpc_credential_v2alpha

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type AwsIamConfig struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	Region               string   `protobuf:"bytes,2,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsIamConfig) Reset()         { *m = AwsIamConfig{} }
func (m *AwsIamConfig) String() string { return proto.CompactTextString(m) }
func (*AwsIamConfig) ProtoMessage()    {}
func (*AwsIamConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb99e742622a8430, []int{0}
}

func (m *AwsIamConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsIamConfig.Unmarshal(m, b)
}
func (m *AwsIamConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsIamConfig.Marshal(b, m, deterministic)
}
func (m *AwsIamConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsIamConfig.Merge(m, src)
}
func (m *AwsIamConfig) XXX_Size() int {
	return xxx_messageInfo_AwsIamConfig.Size(m)
}
func (m *AwsIamConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsIamConfig.DiscardUnknown(m)
}

var xxx_messageInfo_AwsIamConfig proto.InternalMessageInfo

func (m *AwsIamConfig) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *AwsIamConfig) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

func init() {
	proto.RegisterType((*AwsIamConfig)(nil), "envoy.config.grpc_credential.v2alpha.AwsIamConfig")
}

func init() {
	proto.RegisterFile("envoy/config/grpc_credential/v2alpha/aws_iam.proto", fileDescriptor_fb99e742622a8430)
}

var fileDescriptor_fb99e742622a8430 = []byte{
	// 204 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xce, 0xbf, 0x4e, 0x87, 0x30,
	0x10, 0xc0, 0xf1, 0x94, 0x18, 0x8c, 0x85, 0x89, 0x41, 0x89, 0x13, 0x31, 0x0e, 0xc6, 0xa1, 0x4d,
	0xf0, 0x09, 0x2c, 0x93, 0x8b, 0x21, 0xbc, 0x00, 0x39, 0xcb, 0x89, 0x4d, 0xe8, 0x9f, 0x14, 0x52,
	0xe4, 0xd5, 0x9d, 0x0c, 0x05, 0x17, 0xa7, 0xdf, 0xd6, 0xe6, 0xee, 0x7b, 0xf9, 0xd0, 0x1a, 0x4d,
	0xb0, 0x1b, 0x97, 0xd6, 0x7c, 0xaa, 0x91, 0x8f, 0xde, 0xc9, 0x5e, 0x7a, 0x1c, 0xd0, 0x2c, 0x0a,
	0x26, 0x1e, 0x6a, 0x98, 0xdc, 0x17, 0x70, 0x58, 0xe7, 0x5e, 0x81, 0x66, 0xce, 0xdb, 0xc5, 0x16,
	0x8f, 0xb1, 0x61, 0x47, 0xc3, 0xfe, 0x35, 0xec, 0x6c, 0xee, 0xef, 0x02, 0x4c, 0x6a, 0x80, 0x05,
	0xf9, 0xdf, 0xe3, 0xc8, 0x1f, 0x3a, 0x9a, 0xbf, 0xae, 0xf3, 0x1b, 0xe8, 0x26, 0x1e, 0x28, 0x9e,
	0x69, 0x3e, 0xa3, 0x0f, 0x4a, 0x62, 0x6f, 0x40, 0x63, 0x49, 0x2a, 0xf2, 0x74, 0x23, 0xae, 0x7f,
	0xc4, 0x95, 0x4f, 0x2a, 0xd2, 0x65, 0xe7, 0xf0, 0x1d, 0x34, 0x16, 0xb7, 0x34, 0xf5, 0x38, 0x2a,
	0x6b, 0xca, 0x64, 0xdf, 0xea, 0xce, 0x9f, 0x68, 0x68, 0xad, 0x2c, 0x8b, 0x2e, 0xe7, 0xed, 0xf7,
	0xc6, 0x2e, 0x21, 0x8a, 0xec, 0x70, 0xb4, 0x3b, 0xab, 0x25, 0x1f, 0x69, 0xf4, 0xbd, 0xfc, 0x06,
	0x00, 0x00, 0xff, 0xff, 0x3a, 0xda, 0xe6, 0x69, 0x14, 0x01, 0x00, 0x00,
}