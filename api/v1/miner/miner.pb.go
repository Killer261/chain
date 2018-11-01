// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1/miner/miner.proto

package miner

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "bitbucket.org/cpchain/chain/api/v1/common"
import empty "github.com/golang/protobuf/ptypes/empty"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

// Threads represents the number of threads opened.
type Threads struct {
	Threads              int32    `protobuf:"varint,1,opt,name=threads,proto3" json:"threads,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Threads) Reset()         { *m = Threads{} }
func (m *Threads) String() string { return proto.CompactTextString(m) }
func (*Threads) ProtoMessage()    {}
func (*Threads) Descriptor() ([]byte, []int) {
	return fileDescriptor_miner_5a07a9ff93708cdc, []int{0}
}
func (m *Threads) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Threads.Unmarshal(m, b)
}
func (m *Threads) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Threads.Marshal(b, m, deterministic)
}
func (dst *Threads) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Threads.Merge(dst, src)
}
func (m *Threads) XXX_Size() int {
	return xxx_messageInfo_Threads.Size(m)
}
func (m *Threads) XXX_DiscardUnknown() {
	xxx_messageInfo_Threads.DiscardUnknown(m)
}

var xxx_messageInfo_Threads proto.InternalMessageInfo

func (m *Threads) GetThreads() int32 {
	if m != nil {
		return m.Threads
	}
	return 0
}

// ExTra the extra data string that is included when this miner mines a block.
type Extra struct {
	Extra                string   `protobuf:"bytes,1,opt,name=extra,proto3" json:"extra,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Extra) Reset()         { *m = Extra{} }
func (m *Extra) String() string { return proto.CompactTextString(m) }
func (*Extra) ProtoMessage()    {}
func (*Extra) Descriptor() ([]byte, []int) {
	return fileDescriptor_miner_5a07a9ff93708cdc, []int{1}
}
func (m *Extra) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extra.Unmarshal(m, b)
}
func (m *Extra) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extra.Marshal(b, m, deterministic)
}
func (dst *Extra) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extra.Merge(dst, src)
}
func (m *Extra) XXX_Size() int {
	return xxx_messageInfo_Extra.Size(m)
}
func (m *Extra) XXX_DiscardUnknown() {
	xxx_messageInfo_Extra.DiscardUnknown(m)
}

var xxx_messageInfo_Extra proto.InternalMessageInfo

func (m *Extra) GetExtra() string {
	if m != nil {
		return m.Extra
	}
	return ""
}

func init() {
	proto.RegisterType((*Threads)(nil), "miner.Threads")
	proto.RegisterType((*Extra)(nil), "miner.Extra")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MinerManagerClient is the client API for MinerManager service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MinerManagerClient interface {
	// Start the miner with the given number of threads. If threads is nil the number
	// of workers started is equal to the number of logical CPUs that are usable by
	// this process. If mining is already running, this method adjust the number of
	// threads allowed to use.
	Start(ctx context.Context, in *Threads, opts ...grpc.CallOption) (*empty.Empty, error)
	//    Stop the miner
	Stop(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.IsOk, error)
	// SetExtra sets the extra data string that is included when this miner mines a block.
	SetExtra(ctx context.Context, in *Extra, opts ...grpc.CallOption) (*common.IsOk, error)
	// SetGasPrice sets the minimum accepted gas price for the miner.
	SetGasPrice(ctx context.Context, in *common.GasPrice, opts ...grpc.CallOption) (*common.IsOk, error)
	// SetCoinbase sets the coinbase of the miner
	SetCoinbase(ctx context.Context, in *common.Address, opts ...grpc.CallOption) (*common.IsOk, error)
}

type minerManagerClient struct {
	cc *grpc.ClientConn
}

func NewMinerManagerClient(cc *grpc.ClientConn) MinerManagerClient {
	return &minerManagerClient{cc}
}

func (c *minerManagerClient) Start(ctx context.Context, in *Threads, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/miner.MinerManager/Start", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerManagerClient) Stop(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*common.IsOk, error) {
	out := new(common.IsOk)
	err := c.cc.Invoke(ctx, "/miner.MinerManager/Stop", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerManagerClient) SetExtra(ctx context.Context, in *Extra, opts ...grpc.CallOption) (*common.IsOk, error) {
	out := new(common.IsOk)
	err := c.cc.Invoke(ctx, "/miner.MinerManager/SetExtra", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerManagerClient) SetGasPrice(ctx context.Context, in *common.GasPrice, opts ...grpc.CallOption) (*common.IsOk, error) {
	out := new(common.IsOk)
	err := c.cc.Invoke(ctx, "/miner.MinerManager/SetGasPrice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *minerManagerClient) SetCoinbase(ctx context.Context, in *common.Address, opts ...grpc.CallOption) (*common.IsOk, error) {
	out := new(common.IsOk)
	err := c.cc.Invoke(ctx, "/miner.MinerManager/SetCoinbase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MinerManagerServer is the server API for MinerManager service.
type MinerManagerServer interface {
	// Start the miner with the given number of threads. If threads is nil the number
	// of workers started is equal to the number of logical CPUs that are usable by
	// this process. If mining is already running, this method adjust the number of
	// threads allowed to use.
	Start(context.Context, *Threads) (*empty.Empty, error)
	//    Stop the miner
	Stop(context.Context, *empty.Empty) (*common.IsOk, error)
	// SetExtra sets the extra data string that is included when this miner mines a block.
	SetExtra(context.Context, *Extra) (*common.IsOk, error)
	// SetGasPrice sets the minimum accepted gas price for the miner.
	SetGasPrice(context.Context, *common.GasPrice) (*common.IsOk, error)
	// SetCoinbase sets the coinbase of the miner
	SetCoinbase(context.Context, *common.Address) (*common.IsOk, error)
}

func RegisterMinerManagerServer(s *grpc.Server, srv MinerManagerServer) {
	s.RegisterService(&_MinerManager_serviceDesc, srv)
}

func _MinerManager_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Threads)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerManagerServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miner.MinerManager/Start",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerManagerServer).Start(ctx, req.(*Threads))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinerManager_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerManagerServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miner.MinerManager/Stop",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerManagerServer).Stop(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinerManager_SetExtra_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Extra)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerManagerServer).SetExtra(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miner.MinerManager/SetExtra",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerManagerServer).SetExtra(ctx, req.(*Extra))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinerManager_SetGasPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.GasPrice)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerManagerServer).SetGasPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miner.MinerManager/SetGasPrice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerManagerServer).SetGasPrice(ctx, req.(*common.GasPrice))
	}
	return interceptor(ctx, in, info, handler)
}

func _MinerManager_SetCoinbase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Address)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MinerManagerServer).SetCoinbase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/miner.MinerManager/SetCoinbase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MinerManagerServer).SetCoinbase(ctx, req.(*common.Address))
	}
	return interceptor(ctx, in, info, handler)
}

var _MinerManager_serviceDesc = grpc.ServiceDesc{
	ServiceName: "miner.MinerManager",
	HandlerType: (*MinerManagerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Start",
			Handler:    _MinerManager_Start_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _MinerManager_Stop_Handler,
		},
		{
			MethodName: "SetExtra",
			Handler:    _MinerManager_SetExtra_Handler,
		},
		{
			MethodName: "SetGasPrice",
			Handler:    _MinerManager_SetGasPrice_Handler,
		},
		{
			MethodName: "SetCoinbase",
			Handler:    _MinerManager_SetCoinbase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "v1/miner/miner.proto",
}

func init() { proto.RegisterFile("v1/miner/miner.proto", fileDescriptor_miner_5a07a9ff93708cdc) }

var fileDescriptor_miner_5a07a9ff93708cdc = []byte{
	// 343 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x18, 0xc5, 0xc9, 0xbd, 0xc6, 0xea, 0x58, 0x54, 0xa6, 0xa5, 0xd6, 0xd4, 0x16, 0x1d, 0x5d, 0x48,
	0x17, 0x13, 0xd4, 0x5d, 0x77, 0x2a, 0x45, 0x8a, 0x94, 0x4a, 0xe3, 0xca, 0xdd, 0xa4, 0x19, 0x63,
	0xd0, 0xcc, 0x84, 0x99, 0x51, 0x74, 0xeb, 0x2b, 0xf8, 0x02, 0xbe, 0x93, 0xaf, 0xe0, 0x83, 0xc8,
	0xfc, 0x73, 0x11, 0x11, 0xdc, 0x24, 0x73, 0x72, 0x72, 0x7e, 0x39, 0xdf, 0x47, 0x40, 0xfb, 0xe9,
	0x28, 0x2e, 0x0b, 0x46, 0x85, 0xbd, 0xe2, 0x4a, 0x70, 0xc5, 0x61, 0x68, 0x44, 0xd4, 0x5a, 0xf0,
	0xb2, 0xe4, 0x2c, 0xb6, 0x37, 0xeb, 0x45, 0xbd, 0x9c, 0xf3, 0xfc, 0x81, 0xc6, 0x46, 0xa5, 0x8f,
	0xb7, 0x31, 0x2d, 0x2b, 0xf5, 0xe2, 0xcc, 0x1d, 0x67, 0x92, 0xaa, 0x88, 0x09, 0x63, 0x5c, 0x11,
	0x55, 0x70, 0x26, 0xad, 0x8b, 0xf6, 0x41, 0xe3, 0xfa, 0x4e, 0x50, 0x92, 0x49, 0xd8, 0x05, 0x0d,
	0x65, 0x8f, 0xdd, 0x60, 0x37, 0x38, 0x0c, 0xe7, 0x5e, 0xa2, 0x3e, 0x08, 0xc7, 0xcf, 0x4a, 0x10,
	0xd8, 0x06, 0x21, 0xd5, 0x07, 0xf3, 0xc2, 0xea, 0xdc, 0x8a, 0xe3, 0xf7, 0xff, 0xa0, 0x39, 0xd5,
	0xed, 0xa6, 0x84, 0x91, 0x9c, 0x0a, 0x38, 0x03, 0x61, 0xa2, 0x88, 0x50, 0x70, 0x1d, 0xdb, 0x11,
	0xdc, 0x27, 0xa2, 0x0e, 0xb6, 0x65, 0xb0, 0x6f, 0x8a, 0xc7, 0xba, 0x29, 0x1a, 0xbc, 0x7e, 0x7c,
	0xbe, 0xfd, 0xeb, 0xa2, 0x96, 0x69, 0xf9, 0x3d, 0xbf, 0xd4, 0x90, 0x51, 0x30, 0x84, 0x97, 0x60,
	0x29, 0x51, 0xbc, 0x82, 0xbf, 0xe4, 0xa3, 0x26, 0x76, 0xfb, 0x98, 0xc8, 0xd9, 0x3d, 0xea, 0x1b,
	0xda, 0x16, 0x82, 0x75, 0x1a, 0xaf, 0x34, 0x6c, 0x02, 0x56, 0x12, 0xaa, 0xec, 0x40, 0x4d, 0x57,
	0xd0, 0xa8, 0x1a, 0x66, 0xcf, 0x60, 0x7a, 0xa8, 0x53, 0xc3, 0xb8, 0xac, 0x46, 0x25, 0x60, 0x2d,
	0xa1, 0xea, 0x82, 0xc8, 0x2b, 0x51, 0x2c, 0x28, 0xdc, 0xf4, 0x79, 0xff, 0xa4, 0x46, 0x3c, 0x30,
	0xc4, 0x01, 0xda, 0xfe, 0x41, 0xf4, 0x01, 0x0d, 0x9d, 0x1b, 0xe8, 0x39, 0x2f, 0x58, 0x4a, 0x24,
	0x85, 0x1b, 0x1e, 0x71, 0x9a, 0x65, 0x82, 0x4a, 0xf9, 0x67, 0xa6, 0x27, 0x8c, 0x82, 0xe1, 0x59,
	0xe3, 0xc6, 0xfe, 0x3f, 0xe9, 0xb2, 0xd9, 0xdc, 0xc9, 0x57, 0x00, 0x00, 0x00, 0xff, 0xff, 0x20,
	0xe1, 0xce, 0xa8, 0x65, 0x02, 0x00, 0x00,
}
