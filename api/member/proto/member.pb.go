// Code generated by protoc-gen-go. DO NOT EDIT.
// source: member.proto

package member

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

type GetByIDRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetByIDRequest) Reset()         { *m = GetByIDRequest{} }
func (m *GetByIDRequest) String() string { return proto.CompactTextString(m) }
func (*GetByIDRequest) ProtoMessage()    {}
func (*GetByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{0}
}

func (m *GetByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetByIDRequest.Unmarshal(m, b)
}
func (m *GetByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetByIDRequest.Marshal(b, m, deterministic)
}
func (m *GetByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetByIDRequest.Merge(m, src)
}
func (m *GetByIDRequest) XXX_Size() int {
	return xxx_messageInfo_GetByIDRequest.Size(m)
}
func (m *GetByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetByIDRequest proto.InternalMessageInfo

func (m *GetByIDRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MemberResponse struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Phone                string   `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,4,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberResponse) Reset()         { *m = MemberResponse{} }
func (m *MemberResponse) String() string { return proto.CompactTextString(m) }
func (*MemberResponse) ProtoMessage()    {}
func (*MemberResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_9b9836b7e13de206, []int{1}
}

func (m *MemberResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberResponse.Unmarshal(m, b)
}
func (m *MemberResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberResponse.Marshal(b, m, deterministic)
}
func (m *MemberResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberResponse.Merge(m, src)
}
func (m *MemberResponse) XXX_Size() int {
	return xxx_messageInfo_MemberResponse.Size(m)
}
func (m *MemberResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberResponse proto.InternalMessageInfo

func (m *MemberResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *MemberResponse) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

func (m *MemberResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MemberResponse) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

func init() {
	proto.RegisterType((*GetByIDRequest)(nil), "GetByIDRequest")
	proto.RegisterType((*MemberResponse)(nil), "MemberResponse")
}

func init() { proto.RegisterFile("member.proto", fileDescriptor_9b9836b7e13de206) }

var fileDescriptor_9b9836b7e13de206 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc9, 0x4d, 0xcd, 0x4d,
	0x4a, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe0, 0xe2, 0x73, 0x4f, 0x2d, 0x71,
	0xaa, 0xf4, 0x74, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0e, 0x62, 0xca, 0x4c, 0x51, 0x8a, 0xe1, 0xe2, 0xf3, 0x05,
	0xeb, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x45, 0x57, 0x21, 0x24, 0xc2, 0xc5, 0x5a,
	0x90, 0x91, 0x9f, 0x97, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x08, 0x09, 0x71,
	0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x83, 0x05, 0xc1, 0x6c, 0x21, 0x01, 0x2e, 0xe6, 0xc4,
	0xf4, 0x54, 0x09, 0x16, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x10, 0xd3, 0xc8, 0x94, 0x8b, 0x0d, 0x62,
	0xba, 0x90, 0x36, 0x17, 0x3b, 0xd4, 0x25, 0x42, 0xfc, 0x7a, 0xa8, 0x6e, 0x92, 0xe2, 0xd7, 0x43,
	0x75, 0x42, 0x12, 0x1b, 0xd8, 0xf5, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb2, 0xa7, 0xbc,
	0xca, 0xcd, 0x00, 0x00, 0x00,
}
