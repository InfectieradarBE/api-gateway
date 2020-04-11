// Code generated by protoc-gen-go. DO NOT EDIT.
// source: global-types.proto

package api

import (
	fmt "fmt"
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

type Status_StatusValue int32

const (
	Status_NORMAL  Status_StatusValue = 0
	Status_PROBLEM Status_StatusValue = 1
)

var Status_StatusValue_name = map[int32]string{
	0: "NORMAL",
	1: "PROBLEM",
}

var Status_StatusValue_value = map[string]int32{
	"NORMAL":  0,
	"PROBLEM": 1,
}

func (x Status_StatusValue) String() string {
	return proto.EnumName(Status_StatusValue_name, int32(x))
}

func (Status_StatusValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3fa09321e03748b2, []int{2, 0}
}

//
// Authentication types used in multiple services
type UserCredentials struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	InstanceId           string   `protobuf:"bytes,3,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserCredentials) Reset()         { *m = UserCredentials{} }
func (m *UserCredentials) String() string { return proto.CompactTextString(m) }
func (*UserCredentials) ProtoMessage()    {}
func (*UserCredentials) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fa09321e03748b2, []int{0}
}

func (m *UserCredentials) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCredentials.Unmarshal(m, b)
}
func (m *UserCredentials) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCredentials.Marshal(b, m, deterministic)
}
func (m *UserCredentials) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCredentials.Merge(m, src)
}
func (m *UserCredentials) XXX_Size() int {
	return xxx_messageInfo_UserCredentials.Size(m)
}
func (m *UserCredentials) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCredentials.DiscardUnknown(m)
}

var xxx_messageInfo_UserCredentials proto.InternalMessageInfo

func (m *UserCredentials) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserCredentials) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserCredentials) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

type TokenInfos struct {
	Id                   string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	InstanceId           string            `protobuf:"bytes,2,opt,name=instance_id,json=instanceId,proto3" json:"instance_id,omitempty"`
	IssuedAt             int64             `protobuf:"varint,3,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	Payload              map[string]string `protobuf:"bytes,4,rep,name=payload,proto3" json:"payload,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ProfilId             string            `protobuf:"bytes,5,opt,name=profil_id,json=profilId,proto3" json:"profil_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TokenInfos) Reset()         { *m = TokenInfos{} }
func (m *TokenInfos) String() string { return proto.CompactTextString(m) }
func (*TokenInfos) ProtoMessage()    {}
func (*TokenInfos) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fa09321e03748b2, []int{1}
}

func (m *TokenInfos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TokenInfos.Unmarshal(m, b)
}
func (m *TokenInfos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TokenInfos.Marshal(b, m, deterministic)
}
func (m *TokenInfos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenInfos.Merge(m, src)
}
func (m *TokenInfos) XXX_Size() int {
	return xxx_messageInfo_TokenInfos.Size(m)
}
func (m *TokenInfos) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenInfos.DiscardUnknown(m)
}

var xxx_messageInfo_TokenInfos proto.InternalMessageInfo

func (m *TokenInfos) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TokenInfos) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

func (m *TokenInfos) GetIssuedAt() int64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *TokenInfos) GetPayload() map[string]string {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *TokenInfos) GetProfilId() string {
	if m != nil {
		return m.ProfilId
	}
	return ""
}

//
// Status is typically used as a return value indicating if the method was
// performed normally, or the system has any internal error e.g. checking system
// status of a service
type Status struct {
	Status               Status_StatusValue `protobuf:"varint,1,opt,name=status,proto3,enum=inf.Status_StatusValue" json:"status,omitempty"`
	Msg                  string             `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_3fa09321e03748b2, []int{2}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetStatus() Status_StatusValue {
	if m != nil {
		return m.Status
	}
	return Status_NORMAL
}

func (m *Status) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterEnum("inf.Status_StatusValue", Status_StatusValue_name, Status_StatusValue_value)
	proto.RegisterType((*UserCredentials)(nil), "inf.UserCredentials")
	proto.RegisterType((*TokenInfos)(nil), "inf.TokenInfos")
	proto.RegisterMapType((map[string]string)(nil), "inf.TokenInfos.PayloadEntry")
	proto.RegisterType((*Status)(nil), "inf.Status")
}

func init() {
	proto.RegisterFile("global-types.proto", fileDescriptor_3fa09321e03748b2)
}

var fileDescriptor_3fa09321e03748b2 = []byte{
	// 335 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x5d, 0x6b, 0xf2, 0x30,
	0x14, 0x7e, 0xdb, 0xbe, 0x56, 0x3d, 0x1d, 0x4e, 0xc2, 0x60, 0xc5, 0x0d, 0x26, 0xbd, 0x18, 0xde,
	0xac, 0x03, 0x07, 0x63, 0x78, 0xa7, 0xc3, 0x0b, 0x41, 0xa7, 0x74, 0x1f, 0x17, 0xbb, 0x91, 0xb8,
	0xa4, 0x12, 0xac, 0x49, 0x69, 0xe2, 0x46, 0x7f, 0xf3, 0xfe, 0xc4, 0x48, 0x1a, 0x37, 0xd9, 0xae,
	0x72, 0xce, 0x79, 0xce, 0xf3, 0x91, 0x04, 0xd0, 0x3a, 0x13, 0x2b, 0x9c, 0x5d, 0xa9, 0x32, 0xa7,
	0x32, 0xce, 0x0b, 0xa1, 0x04, 0xf2, 0x18, 0x4f, 0x23, 0x02, 0xc7, 0xcf, 0x92, 0x16, 0xf7, 0x05,
	0x25, 0x94, 0x2b, 0x86, 0x33, 0x89, 0x4e, 0xa0, 0x46, 0xb7, 0x98, 0x65, 0xa1, 0xd3, 0x75, 0x7a,
	0xcd, 0xa4, 0x6a, 0x50, 0x07, 0x1a, 0x39, 0x96, 0xf2, 0x43, 0x14, 0x24, 0x74, 0x0d, 0xf0, 0xdd,
	0xa3, 0x0b, 0x08, 0x18, 0x97, 0x0a, 0xf3, 0x37, 0xba, 0x64, 0x24, 0xf4, 0x0c, 0x0c, 0xfb, 0xd1,
	0x84, 0x44, 0x9f, 0x0e, 0xc0, 0x93, 0xd8, 0x50, 0x3e, 0xe1, 0xa9, 0x90, 0xa8, 0x05, 0x2e, 0x23,
	0x56, 0xde, 0x65, 0x7f, 0xf8, 0xee, 0x6f, 0x3e, 0x3a, 0x83, 0x26, 0x93, 0x72, 0x47, 0xc9, 0x12,
	0x2b, 0x23, 0xef, 0x25, 0x8d, 0x6a, 0x30, 0x54, 0xe8, 0x16, 0xea, 0x39, 0x2e, 0x33, 0x81, 0x49,
	0xf8, 0xbf, 0xeb, 0xf5, 0x82, 0xfe, 0x79, 0xcc, 0x78, 0x1a, 0xff, 0xf8, 0xc5, 0x8b, 0x0a, 0x1e,
	0x73, 0x55, 0x94, 0xc9, 0x7e, 0x59, 0x8b, 0xe6, 0x85, 0x48, 0x59, 0xa6, 0x3d, 0x6b, 0xf6, 0x4a,
	0x66, 0x30, 0x21, 0x9d, 0x01, 0x1c, 0x1d, 0xb2, 0x50, 0x1b, 0xbc, 0x0d, 0x2d, 0x6d, 0x66, 0x5d,
	0xea, 0x67, 0x7a, 0xc7, 0xd9, 0x8e, 0xda, 0xb8, 0x55, 0x33, 0x70, 0xef, 0x9c, 0x48, 0x82, 0xff,
	0xa8, 0xb0, 0xda, 0x49, 0x74, 0x0d, 0xbe, 0x34, 0x95, 0x21, 0xb6, 0xfa, 0xa7, 0x26, 0x59, 0x05,
	0xda, 0xe3, 0x45, 0xd3, 0x12, 0xbb, 0xa6, 0x6d, 0xb6, 0x72, 0x6d, 0x25, 0x75, 0x19, 0x5d, 0x42,
	0x70, 0xb0, 0x88, 0x00, 0xfc, 0x87, 0x79, 0x32, 0x1b, 0x4e, 0xdb, 0xff, 0x50, 0x00, 0xf5, 0x45,
	0x32, 0x1f, 0x4d, 0xc7, 0xb3, 0xb6, 0x33, 0xaa, 0xbd, 0x7a, 0x38, 0x67, 0x2b, 0xdf, 0xfc, 0xed,
	0xcd, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x5a, 0x86, 0xd1, 0x02, 0xf1, 0x01, 0x00, 0x00,
}