// Code generated by protoc-gen-go. DO NOT EDIT.
// source: public/messages.proto

package main

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

type FromServer struct {
	Epoch uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// Types that are valid to be assigned to Msg:
	//	*FromServer_Hello
	//	*FromServer_Ping
	//	*FromServer_Pong
	Msg                  isFromServer_Msg `protobuf_oneof:"msg"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FromServer) Reset()         { *m = FromServer{} }
func (m *FromServer) String() string { return proto.CompactTextString(m) }
func (*FromServer) ProtoMessage()    {}
func (*FromServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2329d7f415054398, []int{0}
}

func (m *FromServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FromServer.Unmarshal(m, b)
}
func (m *FromServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FromServer.Marshal(b, m, deterministic)
}
func (m *FromServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FromServer.Merge(m, src)
}
func (m *FromServer) XXX_Size() int {
	return xxx_messageInfo_FromServer.Size(m)
}
func (m *FromServer) XXX_DiscardUnknown() {
	xxx_messageInfo_FromServer.DiscardUnknown(m)
}

var xxx_messageInfo_FromServer proto.InternalMessageInfo

func (m *FromServer) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

type isFromServer_Msg interface {
	isFromServer_Msg()
}

type FromServer_Hello struct {
	Hello *Hello `protobuf:"bytes,2,opt,name=hello,proto3,oneof"`
}

type FromServer_Ping struct {
	Ping *Ping `protobuf:"bytes,3,opt,name=ping,proto3,oneof"`
}

type FromServer_Pong struct {
	Pong *Pong `protobuf:"bytes,4,opt,name=pong,proto3,oneof"`
}

func (*FromServer_Hello) isFromServer_Msg() {}

func (*FromServer_Ping) isFromServer_Msg() {}

func (*FromServer_Pong) isFromServer_Msg() {}

func (m *FromServer) GetMsg() isFromServer_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *FromServer) GetHello() *Hello {
	if x, ok := m.GetMsg().(*FromServer_Hello); ok {
		return x.Hello
	}
	return nil
}

func (m *FromServer) GetPing() *Ping {
	if x, ok := m.GetMsg().(*FromServer_Ping); ok {
		return x.Ping
	}
	return nil
}

func (m *FromServer) GetPong() *Pong {
	if x, ok := m.GetMsg().(*FromServer_Pong); ok {
		return x.Pong
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FromServer) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FromServer_Hello)(nil),
		(*FromServer_Ping)(nil),
		(*FromServer_Pong)(nil),
	}
}

type Ping struct {
	SentAt               uint64   `protobuf:"varint,1,opt,name=sentAt,proto3" json:"sentAt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ping) Reset()         { *m = Ping{} }
func (m *Ping) String() string { return proto.CompactTextString(m) }
func (*Ping) ProtoMessage()    {}
func (*Ping) Descriptor() ([]byte, []int) {
	return fileDescriptor_2329d7f415054398, []int{1}
}

func (m *Ping) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ping.Unmarshal(m, b)
}
func (m *Ping) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ping.Marshal(b, m, deterministic)
}
func (m *Ping) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ping.Merge(m, src)
}
func (m *Ping) XXX_Size() int {
	return xxx_messageInfo_Ping.Size(m)
}
func (m *Ping) XXX_DiscardUnknown() {
	xxx_messageInfo_Ping.DiscardUnknown(m)
}

var xxx_messageInfo_Ping proto.InternalMessageInfo

func (m *Ping) GetSentAt() uint64 {
	if m != nil {
		return m.SentAt
	}
	return 0
}

type Pong struct {
	Replyfrom            uint64   `protobuf:"varint,1,opt,name=replyfrom,proto3" json:"replyfrom,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Pong) Reset()         { *m = Pong{} }
func (m *Pong) String() string { return proto.CompactTextString(m) }
func (*Pong) ProtoMessage()    {}
func (*Pong) Descriptor() ([]byte, []int) {
	return fileDescriptor_2329d7f415054398, []int{2}
}

func (m *Pong) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Pong.Unmarshal(m, b)
}
func (m *Pong) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Pong.Marshal(b, m, deterministic)
}
func (m *Pong) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pong.Merge(m, src)
}
func (m *Pong) XXX_Size() int {
	return xxx_messageInfo_Pong.Size(m)
}
func (m *Pong) XXX_DiscardUnknown() {
	xxx_messageInfo_Pong.DiscardUnknown(m)
}

var xxx_messageInfo_Pong proto.InternalMessageInfo

func (m *Pong) GetReplyfrom() uint64 {
	if m != nil {
		return m.Replyfrom
	}
	return 0
}

type Hello struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hello) Reset()         { *m = Hello{} }
func (m *Hello) String() string { return proto.CompactTextString(m) }
func (*Hello) ProtoMessage()    {}
func (*Hello) Descriptor() ([]byte, []int) {
	return fileDescriptor_2329d7f415054398, []int{3}
}

func (m *Hello) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hello.Unmarshal(m, b)
}
func (m *Hello) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hello.Marshal(b, m, deterministic)
}
func (m *Hello) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hello.Merge(m, src)
}
func (m *Hello) XXX_Size() int {
	return xxx_messageInfo_Hello.Size(m)
}
func (m *Hello) XXX_DiscardUnknown() {
	xxx_messageInfo_Hello.DiscardUnknown(m)
}

var xxx_messageInfo_Hello proto.InternalMessageInfo

func (m *Hello) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

type FromClient struct {
	Epoch uint64 `protobuf:"varint,1,opt,name=epoch,proto3" json:"epoch,omitempty"`
	// Types that are valid to be assigned to Msg:
	//	*FromClient_Hello
	//	*FromClient_Ping
	//	*FromClient_Pong
	Msg                  isFromClient_Msg `protobuf_oneof:"msg"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *FromClient) Reset()         { *m = FromClient{} }
func (m *FromClient) String() string { return proto.CompactTextString(m) }
func (*FromClient) ProtoMessage()    {}
func (*FromClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_2329d7f415054398, []int{4}
}

func (m *FromClient) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FromClient.Unmarshal(m, b)
}
func (m *FromClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FromClient.Marshal(b, m, deterministic)
}
func (m *FromClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FromClient.Merge(m, src)
}
func (m *FromClient) XXX_Size() int {
	return xxx_messageInfo_FromClient.Size(m)
}
func (m *FromClient) XXX_DiscardUnknown() {
	xxx_messageInfo_FromClient.DiscardUnknown(m)
}

var xxx_messageInfo_FromClient proto.InternalMessageInfo

func (m *FromClient) GetEpoch() uint64 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

type isFromClient_Msg interface {
	isFromClient_Msg()
}

type FromClient_Hello struct {
	Hello *Hello `protobuf:"bytes,2,opt,name=hello,proto3,oneof"`
}

type FromClient_Ping struct {
	Ping *Ping `protobuf:"bytes,3,opt,name=ping,proto3,oneof"`
}

type FromClient_Pong struct {
	Pong *Pong `protobuf:"bytes,4,opt,name=pong,proto3,oneof"`
}

func (*FromClient_Hello) isFromClient_Msg() {}

func (*FromClient_Ping) isFromClient_Msg() {}

func (*FromClient_Pong) isFromClient_Msg() {}

func (m *FromClient) GetMsg() isFromClient_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (m *FromClient) GetHello() *Hello {
	if x, ok := m.GetMsg().(*FromClient_Hello); ok {
		return x.Hello
	}
	return nil
}

func (m *FromClient) GetPing() *Ping {
	if x, ok := m.GetMsg().(*FromClient_Ping); ok {
		return x.Ping
	}
	return nil
}

func (m *FromClient) GetPong() *Pong {
	if x, ok := m.GetMsg().(*FromClient_Pong); ok {
		return x.Pong
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*FromClient) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*FromClient_Hello)(nil),
		(*FromClient_Ping)(nil),
		(*FromClient_Pong)(nil),
	}
}

func init() {
	proto.RegisterType((*FromServer)(nil), "rts.FromServer")
	proto.RegisterType((*Ping)(nil), "rts.Ping")
	proto.RegisterType((*Pong)(nil), "rts.Pong")
	proto.RegisterType((*Hello)(nil), "rts.Hello")
	proto.RegisterType((*FromClient)(nil), "rts.FromClient")
}

func init() { proto.RegisterFile("public/messages.proto", fileDescriptor_2329d7f415054398) }

var fileDescriptor_2329d7f415054398 = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xcc, 0x91, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xbb, 0x36, 0xbb, 0xb8, 0xe3, 0x2d, 0xa8, 0x2c, 0x22, 0x5a, 0xa2, 0x87, 0x9e, 0x56,
	0xd0, 0x5f, 0x60, 0x05, 0xd9, 0xa3, 0xac, 0x37, 0x6f, 0x6d, 0x19, 0xd3, 0x40, 0x92, 0x09, 0x93,
	0x28, 0xf8, 0x2b, 0xfc, 0xcb, 0x92, 0x74, 0xa9, 0xfe, 0x04, 0x8f, 0x6f, 0xde, 0xf7, 0xe0, 0x3d,
	0x06, 0xce, 0xc2, 0xc7, 0xc6, 0x9a, 0xed, 0x9d, 0xc3, 0x18, 0xd7, 0x1a, 0x63, 0x1f, 0x98, 0x12,
	0xc9, 0x39, 0xa7, 0xa8, 0xbe, 0x2b, 0x80, 0x67, 0x26, 0xf7, 0x8a, 0xfc, 0x89, 0x2c, 0x4f, 0xa1,
	0xc6, 0x40, 0xdb, 0x5d, 0x57, 0x2d, 0xaa, 0xa5, 0x18, 0xf7, 0x42, 0x2a, 0xa8, 0x77, 0x68, 0x2d,
	0x75, 0x47, 0x8b, 0x6a, 0x79, 0x72, 0x0f, 0x3d, 0xa7, 0xd8, 0x0f, 0xf9, 0x32, 0xcc, 0xc6, 0xbd,
	0x25, 0xaf, 0x41, 0x04, 0xe3, 0x75, 0x37, 0x2f, 0x48, 0x5b, 0x90, 0x17, 0xe3, 0xf5, 0x30, 0x1b,
	0x8b, 0x51, 0x00, 0xf2, 0xba, 0x13, 0x7f, 0x01, 0x9a, 0x00, 0xf2, 0x7a, 0x55, 0xc3, 0xdc, 0x45,
	0xad, 0xae, 0x40, 0xe4, 0x9c, 0x3c, 0x87, 0x26, 0xa2, 0x4f, 0x8f, 0x69, 0xea, 0x32, 0x29, 0x75,
	0x0b, 0x22, 0xc7, 0xe4, 0x25, 0xb4, 0x8c, 0xc1, 0x7e, 0xbd, 0x33, 0xb9, 0x09, 0xf9, 0x3d, 0xa8,
	0x1b, 0xa8, 0x4b, 0x41, 0x79, 0x01, 0xc7, 0x9a, 0x11, 0x53, 0xee, 0x96, 0xa9, 0x76, 0x3c, 0xe8,
	0xc3, 0xf8, 0x27, 0x6b, 0xd0, 0xa7, 0x7f, 0x30, 0x7e, 0xd5, 0xbc, 0x09, 0xb7, 0x36, 0x7e, 0xd3,
	0x94, 0x17, 0x3d, 0xfc, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf7, 0x12, 0xdb, 0x07, 0xbb, 0x01, 0x00,
	0x00,
}
