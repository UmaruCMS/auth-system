// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/auth/main.proto

package auth

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

type Token struct {
	Raw                  string   `protobuf:"bytes,1,opt,name=raw" json:"raw,omitempty"`
	UserId               uint32   `protobuf:"varint,2,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	Valid                bool     `protobuf:"varint,3,opt,name=valid" json:"valid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_main_c8196143926e600c, []int{0}
}
func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (dst *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(dst, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetRaw() string {
	if m != nil {
		return m.Raw
	}
	return ""
}

func (m *Token) GetUserId() uint32 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func init() {
	proto.RegisterType((*Token)(nil), "auth.Token")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Auth service

type AuthClient interface {
	VerifyToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) VerifyToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := grpc.Invoke(ctx, "/auth.Auth/VerifyToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthServer interface {
	VerifyToken(context.Context, *Token) (*Token, error)
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifyToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyToken",
			Handler:    _Auth_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "base/auth/main.proto",
}

func init() { proto.RegisterFile("base/auth/main.proto", fileDescriptor_main_c8196143926e600c) }

var fileDescriptor_main_c8196143926e600c = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x49, 0x4a, 0x2c, 0x4e,
	0xd5, 0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0xcf, 0x4d, 0xcc, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9,
	0x17, 0x62, 0x01, 0x09, 0x28, 0x79, 0x70, 0xb1, 0x86, 0xe4, 0x67, 0xa7, 0xe6, 0x09, 0x09, 0x70,
	0x31, 0x17, 0x25, 0x96, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0x98, 0x42, 0xe2, 0x5c,
	0xec, 0xa5, 0xc5, 0xa9, 0x45, 0xf1, 0x99, 0x29, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xbc, 0x41, 0x6c,
	0x20, 0xae, 0x67, 0x8a, 0x90, 0x08, 0x17, 0x6b, 0x59, 0x62, 0x4e, 0x66, 0x8a, 0x04, 0xb3, 0x02,
	0xa3, 0x06, 0x47, 0x10, 0x84, 0x63, 0x64, 0xc8, 0xc5, 0xe2, 0x58, 0x5a, 0x92, 0x21, 0xa4, 0xc9,
	0xc5, 0x1d, 0x96, 0x5a, 0x94, 0x99, 0x56, 0x09, 0x31, 0x97, 0x5b, 0x0f, 0x64, 0x8f, 0x1e, 0x98,
	0x23, 0x85, 0xcc, 0x51, 0x62, 0x48, 0x62, 0x03, 0xbb, 0xc4, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff,
	0xfb, 0x23, 0xb4, 0xfe, 0xa1, 0x00, 0x00, 0x00,
}
