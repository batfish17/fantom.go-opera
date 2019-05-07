// Code generated by protoc-gen-go. DO NOT EDIT.
// source: wire/ctrl.proto

package wire

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type IDResponse struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDResponse) Reset()         { *m = IDResponse{} }
func (m *IDResponse) String() string { return proto.CompactTextString(m) }
func (*IDResponse) ProtoMessage()    {}
func (*IDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9e6d567004e21b4, []int{0}
}

func (m *IDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDResponse.Unmarshal(m, b)
}
func (m *IDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDResponse.Marshal(b, m, deterministic)
}
func (m *IDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDResponse.Merge(m, src)
}
func (m *IDResponse) XXX_Size() int {
	return xxx_messageInfo_IDResponse.Size(m)
}
func (m *IDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IDResponse proto.InternalMessageInfo

func (m *IDResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type StakeResponse struct {
	Value                float64  `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StakeResponse) Reset()         { *m = StakeResponse{} }
func (m *StakeResponse) String() string { return proto.CompactTextString(m) }
func (*StakeResponse) ProtoMessage()    {}
func (*StakeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9e6d567004e21b4, []int{1}
}

func (m *StakeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StakeResponse.Unmarshal(m, b)
}
func (m *StakeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StakeResponse.Marshal(b, m, deterministic)
}
func (m *StakeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StakeResponse.Merge(m, src)
}
func (m *StakeResponse) XXX_Size() int {
	return xxx_messageInfo_StakeResponse.Size(m)
}
func (m *StakeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StakeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StakeResponse proto.InternalMessageInfo

func (m *StakeResponse) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type InternalTxnRequest struct {
	Amount               uint64   `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	Receiver             string   `protobuf:"bytes,2,opt,name=receiver,proto3" json:"receiver,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InternalTxnRequest) Reset()         { *m = InternalTxnRequest{} }
func (m *InternalTxnRequest) String() string { return proto.CompactTextString(m) }
func (*InternalTxnRequest) ProtoMessage()    {}
func (*InternalTxnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c9e6d567004e21b4, []int{2}
}

func (m *InternalTxnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InternalTxnRequest.Unmarshal(m, b)
}
func (m *InternalTxnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InternalTxnRequest.Marshal(b, m, deterministic)
}
func (m *InternalTxnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InternalTxnRequest.Merge(m, src)
}
func (m *InternalTxnRequest) XXX_Size() int {
	return xxx_messageInfo_InternalTxnRequest.Size(m)
}
func (m *InternalTxnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InternalTxnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InternalTxnRequest proto.InternalMessageInfo

func (m *InternalTxnRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *InternalTxnRequest) GetReceiver() string {
	if m != nil {
		return m.Receiver
	}
	return ""
}

func init() {
	proto.RegisterType((*IDResponse)(nil), "wire.IDResponse")
	proto.RegisterType((*StakeResponse)(nil), "wire.StakeResponse")
	proto.RegisterType((*InternalTxnRequest)(nil), "wire.InternalTxnRequest")
}

func init() { proto.RegisterFile("wire/ctrl.proto", fileDescriptor_c9e6d567004e21b4) }

var fileDescriptor_c9e6d567004e21b4 = []byte{
	// 249 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x93, 0x90, 0x16, 0x1d, 0xf1, 0x0f, 0xa3, 0x94, 0x10, 0x3d, 0xc8, 0x82, 0xe0, 0x69,
	0x23, 0x0a, 0xde, 0xc5, 0x0a, 0xe6, 0xba, 0xfa, 0x05, 0xd2, 0x76, 0x2c, 0xc1, 0xed, 0x6e, 0xdc,
	0x4c, 0xaa, 0x7e, 0x2e, 0xbf, 0xa0, 0x64, 0xb7, 0x56, 0xa5, 0xf4, 0xf8, 0x1e, 0x6f, 0xde, 0xfc,
	0x1e, 0x1c, 0xbe, 0xd7, 0x8e, 0x8a, 0x29, 0x3b, 0x2d, 0x1b, 0x67, 0xd9, 0x62, 0xda, 0x1b, 0xf9,
	0xe9, 0xdc, 0xda, 0xb9, 0xa6, 0xc2, 0x7b, 0x93, 0xee, 0xa5, 0xa0, 0x45, 0xc3, 0x9f, 0x21, 0x22,
	0xce, 0x00, 0xca, 0xb1, 0xa2, 0xb6, 0xb1, 0xa6, 0x25, 0x3c, 0x80, 0xa4, 0x9e, 0x65, 0xf1, 0x79,
	0x7c, 0xb9, 0xab, 0x92, 0x7a, 0x26, 0x2e, 0x60, 0xff, 0x89, 0xab, 0x57, 0x5a, 0x07, 0x4e, 0x60,
	0xb0, 0xac, 0x74, 0x47, 0x3e, 0x13, 0xab, 0x20, 0xc4, 0x23, 0x60, 0x69, 0x98, 0x9c, 0xa9, 0xf4,
	0xf3, 0x87, 0x51, 0xf4, 0xd6, 0x51, 0xcb, 0x38, 0x82, 0x61, 0xb5, 0xb0, 0x9d, 0x61, 0x1f, 0x4e,
	0xd5, 0x4a, 0x61, 0x0e, 0x3b, 0x8e, 0xa6, 0x54, 0x2f, 0xc9, 0x65, 0x89, 0x7f, 0xb5, 0xd6, 0xd7,
	0x5f, 0x31, 0xa4, 0xf7, 0xec, 0x34, 0x5e, 0x41, 0x52, 0x8e, 0x71, 0x24, 0x03, 0xbb, 0xfc, 0x61,
	0x97, 0x0f, 0x3d, 0x7b, 0x7e, 0x24, 0xfb, 0x65, 0xf2, 0x97, 0x5c, 0x44, 0x78, 0x0b, 0x03, 0xcf,
	0xba, 0xf5, 0xe8, 0x38, 0x1c, 0xfd, 0x1b, 0x24, 0x22, 0xbc, 0x83, 0xbd, 0x3f, 0xf0, 0x98, 0xad,
	0xaa, 0x37, 0xf6, 0xe4, 0x5b, 0x7a, 0x45, 0x34, 0x19, 0x7a, 0xe7, 0xe6, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xb8, 0x2c, 0xd0, 0xf4, 0x81, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CtrlClient is the client API for Ctrl service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CtrlClient interface {
	ID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IDResponse, error)
	Stake(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StakeResponse, error)
	InternalTxn(ctx context.Context, in *InternalTxnRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type ctrlClient struct {
	cc *grpc.ClientConn
}

func NewCtrlClient(cc *grpc.ClientConn) CtrlClient {
	return &ctrlClient{cc}
}

func (c *ctrlClient) ID(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*IDResponse, error) {
	out := new(IDResponse)
	err := c.cc.Invoke(ctx, "/wire.Ctrl/ID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctrlClient) Stake(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*StakeResponse, error) {
	out := new(StakeResponse)
	err := c.cc.Invoke(ctx, "/wire.Ctrl/Stake", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ctrlClient) InternalTxn(ctx context.Context, in *InternalTxnRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/wire.Ctrl/InternalTxn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CtrlServer is the server API for Ctrl service.
type CtrlServer interface {
	ID(context.Context, *empty.Empty) (*IDResponse, error)
	Stake(context.Context, *empty.Empty) (*StakeResponse, error)
	InternalTxn(context.Context, *InternalTxnRequest) (*empty.Empty, error)
}

// UnimplementedCtrlServer can be embedded to have forward compatible implementations.
type UnimplementedCtrlServer struct {
}

func (*UnimplementedCtrlServer) ID(ctx context.Context, req *empty.Empty) (*IDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ID not implemented")
}
func (*UnimplementedCtrlServer) Stake(ctx context.Context, req *empty.Empty) (*StakeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stake not implemented")
}
func (*UnimplementedCtrlServer) InternalTxn(ctx context.Context, req *InternalTxnRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InternalTxn not implemented")
}

func RegisterCtrlServer(s *grpc.Server, srv CtrlServer) {
	s.RegisterService(&_Ctrl_serviceDesc, srv)
}

func _Ctrl_ID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtrlServer).ID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wire.Ctrl/ID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtrlServer).ID(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ctrl_Stake_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtrlServer).Stake(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wire.Ctrl/Stake",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtrlServer).Stake(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ctrl_InternalTxn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InternalTxnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CtrlServer).InternalTxn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/wire.Ctrl/InternalTxn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CtrlServer).InternalTxn(ctx, req.(*InternalTxnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ctrl_serviceDesc = grpc.ServiceDesc{
	ServiceName: "wire.Ctrl",
	HandlerType: (*CtrlServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ID",
			Handler:    _Ctrl_ID_Handler,
		},
		{
			MethodName: "Stake",
			Handler:    _Ctrl_Stake_Handler,
		},
		{
			MethodName: "InternalTxn",
			Handler:    _Ctrl_InternalTxn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wire/ctrl.proto",
}
