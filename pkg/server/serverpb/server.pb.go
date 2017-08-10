// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server/serverpb/server.proto

/*
Package serverpb is a generated protocol buffer package.

It is generated from these files:
	server/serverpb/server.proto

It has these top-level messages:
	DrawIPRequest
	DrawIPResponse
	GetNetworkIncludingIPRequest
	GetNetworkIncludingIPResponse
	ActivateIPRequest
	ActivateIPResponse
	GetNetworkRequest
	GetNetworkResponse
	CreateNetworkRequest
	CreateNetworkResponse
	CreatePoolRequest
	CreatePoolResponse
	Tag
	Pool
*/
package serverpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/mwitkow/go-proto-validators"

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

type DrawIPRequest struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Mask int32  `protobuf:"varint,2,opt,name=mask" json:"mask,omitempty"`
	Tags []*Tag `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty"`
}

func (m *DrawIPRequest) Reset()                    { *m = DrawIPRequest{} }
func (m *DrawIPRequest) String() string            { return proto.CompactTextString(m) }
func (*DrawIPRequest) ProtoMessage()               {}
func (*DrawIPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DrawIPRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *DrawIPRequest) GetMask() int32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

func (m *DrawIPRequest) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type DrawIPResponse struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *DrawIPResponse) Reset()                    { *m = DrawIPResponse{} }
func (m *DrawIPResponse) String() string            { return proto.CompactTextString(m) }
func (*DrawIPResponse) ProtoMessage()               {}
func (*DrawIPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DrawIPResponse) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type GetNetworkIncludingIPRequest struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *GetNetworkIncludingIPRequest) Reset()                    { *m = GetNetworkIncludingIPRequest{} }
func (m *GetNetworkIncludingIPRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNetworkIncludingIPRequest) ProtoMessage()               {}
func (*GetNetworkIncludingIPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetNetworkIncludingIPRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type GetNetworkIncludingIPResponse struct {
}

func (m *GetNetworkIncludingIPResponse) Reset()                    { *m = GetNetworkIncludingIPResponse{} }
func (m *GetNetworkIncludingIPResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNetworkIncludingIPResponse) ProtoMessage()               {}
func (*GetNetworkIncludingIPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ActivateIPRequest struct {
	Ip string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
}

func (m *ActivateIPRequest) Reset()                    { *m = ActivateIPRequest{} }
func (m *ActivateIPRequest) String() string            { return proto.CompactTextString(m) }
func (*ActivateIPRequest) ProtoMessage()               {}
func (*ActivateIPRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ActivateIPRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

type ActivateIPResponse struct {
}

func (m *ActivateIPResponse) Reset()                    { *m = ActivateIPResponse{} }
func (m *ActivateIPResponse) String() string            { return proto.CompactTextString(m) }
func (*ActivateIPResponse) ProtoMessage()               {}
func (*ActivateIPResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type GetNetworkRequest struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Mask int32  `protobuf:"varint,2,opt,name=mask" json:"mask,omitempty"`
}

func (m *GetNetworkRequest) Reset()                    { *m = GetNetworkRequest{} }
func (m *GetNetworkRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNetworkRequest) ProtoMessage()               {}
func (*GetNetworkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetNetworkRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *GetNetworkRequest) GetMask() int32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

type GetNetworkResponse struct {
	Ipnet           string   `protobuf:"bytes,1,opt,name=ipnet" json:"ipnet,omitempty"`
	DefaultGateways []string `protobuf:"bytes,2,rep,name=default_gateways,json=defaultGateways" json:"default_gateways,omitempty"`
	Broadcast       string   `protobuf:"bytes,3,opt,name=broadcast" json:"broadcast,omitempty"`
	Netmask         string   `protobuf:"bytes,4,opt,name=netmask" json:"netmask,omitempty"`
	Tags            []*Tag   `protobuf:"bytes,5,rep,name=tags" json:"tags,omitempty"`
}

func (m *GetNetworkResponse) Reset()                    { *m = GetNetworkResponse{} }
func (m *GetNetworkResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNetworkResponse) ProtoMessage()               {}
func (*GetNetworkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *GetNetworkResponse) GetIpnet() string {
	if m != nil {
		return m.Ipnet
	}
	return ""
}

func (m *GetNetworkResponse) GetDefaultGateways() []string {
	if m != nil {
		return m.DefaultGateways
	}
	return nil
}

func (m *GetNetworkResponse) GetBroadcast() string {
	if m != nil {
		return m.Broadcast
	}
	return ""
}

func (m *GetNetworkResponse) GetNetmask() string {
	if m != nil {
		return m.Netmask
	}
	return ""
}

func (m *GetNetworkResponse) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CreateNetworkRequest struct {
	Ip              string   `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Mask            int32    `protobuf:"varint,2,opt,name=mask" json:"mask,omitempty"`
	DefaultGateways []string `protobuf:"bytes,3,rep,name=default_gateways,json=defaultGateways" json:"default_gateways,omitempty"`
	Broadcast       string   `protobuf:"bytes,4,opt,name=broadcast" json:"broadcast,omitempty"`
	Netmask         string   `protobuf:"bytes,5,opt,name=netmask" json:"netmask,omitempty"`
	Tags            []*Tag   `protobuf:"bytes,6,rep,name=tags" json:"tags,omitempty"`
}

func (m *CreateNetworkRequest) Reset()                    { *m = CreateNetworkRequest{} }
func (m *CreateNetworkRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNetworkRequest) ProtoMessage()               {}
func (*CreateNetworkRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CreateNetworkRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CreateNetworkRequest) GetMask() int32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

func (m *CreateNetworkRequest) GetDefaultGateways() []string {
	if m != nil {
		return m.DefaultGateways
	}
	return nil
}

func (m *CreateNetworkRequest) GetBroadcast() string {
	if m != nil {
		return m.Broadcast
	}
	return ""
}

func (m *CreateNetworkRequest) GetNetmask() string {
	if m != nil {
		return m.Netmask
	}
	return ""
}

func (m *CreateNetworkRequest) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type CreateNetworkResponse struct {
}

func (m *CreateNetworkResponse) Reset()                    { *m = CreateNetworkResponse{} }
func (m *CreateNetworkResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateNetworkResponse) ProtoMessage()               {}
func (*CreateNetworkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type CreatePoolRequest struct {
	Ip   string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Mask int32  `protobuf:"varint,2,opt,name=mask" json:"mask,omitempty"`
	Pool *Pool  `protobuf:"bytes,3,opt,name=pool" json:"pool,omitempty"`
}

func (m *CreatePoolRequest) Reset()                    { *m = CreatePoolRequest{} }
func (m *CreatePoolRequest) String() string            { return proto.CompactTextString(m) }
func (*CreatePoolRequest) ProtoMessage()               {}
func (*CreatePoolRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CreatePoolRequest) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *CreatePoolRequest) GetMask() int32 {
	if m != nil {
		return m.Mask
	}
	return 0
}

func (m *CreatePoolRequest) GetPool() *Pool {
	if m != nil {
		return m.Pool
	}
	return nil
}

type CreatePoolResponse struct {
}

func (m *CreatePoolResponse) Reset()                    { *m = CreatePoolResponse{} }
func (m *CreatePoolResponse) String() string            { return proto.CompactTextString(m) }
func (*CreatePoolResponse) ProtoMessage()               {}
func (*CreatePoolResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type Tag struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *Tag) Reset()                    { *m = Tag{} }
func (m *Tag) String() string            { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()               {}
func (*Tag) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *Tag) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *Tag) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Pool struct {
	Start  string `protobuf:"bytes,1,opt,name=start" json:"start,omitempty"`
	End    string `protobuf:"bytes,2,opt,name=end" json:"end,omitempty"`
	Status int32  `protobuf:"varint,3,opt,name=status" json:"status,omitempty"`
	Tags   []*Tag `protobuf:"bytes,4,rep,name=tags" json:"tags,omitempty"`
}

func (m *Pool) Reset()                    { *m = Pool{} }
func (m *Pool) String() string            { return proto.CompactTextString(m) }
func (*Pool) ProtoMessage()               {}
func (*Pool) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *Pool) GetStart() string {
	if m != nil {
		return m.Start
	}
	return ""
}

func (m *Pool) GetEnd() string {
	if m != nil {
		return m.End
	}
	return ""
}

func (m *Pool) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Pool) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*DrawIPRequest)(nil), "serverpb.DrawIPRequest")
	proto.RegisterType((*DrawIPResponse)(nil), "serverpb.DrawIPResponse")
	proto.RegisterType((*GetNetworkIncludingIPRequest)(nil), "serverpb.GetNetworkIncludingIPRequest")
	proto.RegisterType((*GetNetworkIncludingIPResponse)(nil), "serverpb.GetNetworkIncludingIPResponse")
	proto.RegisterType((*ActivateIPRequest)(nil), "serverpb.ActivateIPRequest")
	proto.RegisterType((*ActivateIPResponse)(nil), "serverpb.ActivateIPResponse")
	proto.RegisterType((*GetNetworkRequest)(nil), "serverpb.GetNetworkRequest")
	proto.RegisterType((*GetNetworkResponse)(nil), "serverpb.GetNetworkResponse")
	proto.RegisterType((*CreateNetworkRequest)(nil), "serverpb.CreateNetworkRequest")
	proto.RegisterType((*CreateNetworkResponse)(nil), "serverpb.CreateNetworkResponse")
	proto.RegisterType((*CreatePoolRequest)(nil), "serverpb.CreatePoolRequest")
	proto.RegisterType((*CreatePoolResponse)(nil), "serverpb.CreatePoolResponse")
	proto.RegisterType((*Tag)(nil), "serverpb.Tag")
	proto.RegisterType((*Pool)(nil), "serverpb.Pool")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for NetworkServiceV0 service

type NetworkServiceV0Client interface {
	DrawIP(ctx context.Context, in *DrawIPRequest, opts ...grpc.CallOption) (*DrawIPResponse, error)
	GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*GetNetworkResponse, error)
	CreateNetwork(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*CreateNetworkResponse, error)
	CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*CreatePoolResponse, error)
}

type networkServiceV0Client struct {
	cc *grpc.ClientConn
}

func NewNetworkServiceV0Client(cc *grpc.ClientConn) NetworkServiceV0Client {
	return &networkServiceV0Client{cc}
}

func (c *networkServiceV0Client) DrawIP(ctx context.Context, in *DrawIPRequest, opts ...grpc.CallOption) (*DrawIPResponse, error) {
	out := new(DrawIPResponse)
	err := grpc.Invoke(ctx, "/serverpb.NetworkServiceV0/DrawIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceV0Client) GetNetwork(ctx context.Context, in *GetNetworkRequest, opts ...grpc.CallOption) (*GetNetworkResponse, error) {
	out := new(GetNetworkResponse)
	err := grpc.Invoke(ctx, "/serverpb.NetworkServiceV0/GetNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceV0Client) CreateNetwork(ctx context.Context, in *CreateNetworkRequest, opts ...grpc.CallOption) (*CreateNetworkResponse, error) {
	out := new(CreateNetworkResponse)
	err := grpc.Invoke(ctx, "/serverpb.NetworkServiceV0/CreateNetwork", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *networkServiceV0Client) CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*CreatePoolResponse, error) {
	out := new(CreatePoolResponse)
	err := grpc.Invoke(ctx, "/serverpb.NetworkServiceV0/CreatePool", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for NetworkServiceV0 service

type NetworkServiceV0Server interface {
	DrawIP(context.Context, *DrawIPRequest) (*DrawIPResponse, error)
	GetNetwork(context.Context, *GetNetworkRequest) (*GetNetworkResponse, error)
	CreateNetwork(context.Context, *CreateNetworkRequest) (*CreateNetworkResponse, error)
	CreatePool(context.Context, *CreatePoolRequest) (*CreatePoolResponse, error)
}

func RegisterNetworkServiceV0Server(s *grpc.Server, srv NetworkServiceV0Server) {
	s.RegisterService(&_NetworkServiceV0_serviceDesc, srv)
}

func _NetworkServiceV0_DrawIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DrawIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceV0Server).DrawIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.NetworkServiceV0/DrawIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceV0Server).DrawIP(ctx, req.(*DrawIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServiceV0_GetNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceV0Server).GetNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.NetworkServiceV0/GetNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceV0Server).GetNetwork(ctx, req.(*GetNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServiceV0_CreateNetwork_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNetworkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceV0Server).CreateNetwork(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.NetworkServiceV0/CreateNetwork",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceV0Server).CreateNetwork(ctx, req.(*CreateNetworkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NetworkServiceV0_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NetworkServiceV0Server).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.NetworkServiceV0/CreatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NetworkServiceV0Server).CreatePool(ctx, req.(*CreatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _NetworkServiceV0_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.NetworkServiceV0",
	HandlerType: (*NetworkServiceV0Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DrawIP",
			Handler:    _NetworkServiceV0_DrawIP_Handler,
		},
		{
			MethodName: "GetNetwork",
			Handler:    _NetworkServiceV0_GetNetwork_Handler,
		},
		{
			MethodName: "CreateNetwork",
			Handler:    _NetworkServiceV0_CreateNetwork_Handler,
		},
		{
			MethodName: "CreatePool",
			Handler:    _NetworkServiceV0_CreatePool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/serverpb/server.proto",
}

// Client API for IPServiceV0 service

type IPServiceV0Client interface {
	GetNetworkIncludingIP(ctx context.Context, in *GetNetworkIncludingIPRequest, opts ...grpc.CallOption) (*GetNetworkIncludingIPResponse, error)
	ActivateIP(ctx context.Context, in *ActivateIPRequest, opts ...grpc.CallOption) (*ActivateIPResponse, error)
}

type iPServiceV0Client struct {
	cc *grpc.ClientConn
}

func NewIPServiceV0Client(cc *grpc.ClientConn) IPServiceV0Client {
	return &iPServiceV0Client{cc}
}

func (c *iPServiceV0Client) GetNetworkIncludingIP(ctx context.Context, in *GetNetworkIncludingIPRequest, opts ...grpc.CallOption) (*GetNetworkIncludingIPResponse, error) {
	out := new(GetNetworkIncludingIPResponse)
	err := grpc.Invoke(ctx, "/serverpb.IPServiceV0/GetNetworkIncludingIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iPServiceV0Client) ActivateIP(ctx context.Context, in *ActivateIPRequest, opts ...grpc.CallOption) (*ActivateIPResponse, error) {
	out := new(ActivateIPResponse)
	err := grpc.Invoke(ctx, "/serverpb.IPServiceV0/ActivateIP", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IPServiceV0 service

type IPServiceV0Server interface {
	GetNetworkIncludingIP(context.Context, *GetNetworkIncludingIPRequest) (*GetNetworkIncludingIPResponse, error)
	ActivateIP(context.Context, *ActivateIPRequest) (*ActivateIPResponse, error)
}

func RegisterIPServiceV0Server(s *grpc.Server, srv IPServiceV0Server) {
	s.RegisterService(&_IPServiceV0_serviceDesc, srv)
}

func _IPServiceV0_GetNetworkIncludingIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNetworkIncludingIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPServiceV0Server).GetNetworkIncludingIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.IPServiceV0/GetNetworkIncludingIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPServiceV0Server).GetNetworkIncludingIP(ctx, req.(*GetNetworkIncludingIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IPServiceV0_ActivateIP_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActivateIPRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IPServiceV0Server).ActivateIP(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/serverpb.IPServiceV0/ActivateIP",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IPServiceV0Server).ActivateIP(ctx, req.(*ActivateIPRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _IPServiceV0_serviceDesc = grpc.ServiceDesc{
	ServiceName: "serverpb.IPServiceV0",
	HandlerType: (*IPServiceV0Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetNetworkIncludingIP",
			Handler:    _IPServiceV0_GetNetworkIncludingIP_Handler,
		},
		{
			MethodName: "ActivateIP",
			Handler:    _IPServiceV0_ActivateIP_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/serverpb/server.proto",
}

func init() { proto.RegisterFile("server/serverpb/server.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 846 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x56, 0x41, 0x8f, 0xdb, 0x44,
	0x14, 0xd6, 0xd8, 0x49, 0x60, 0xdf, 0x2a, 0xcb, 0x66, 0xd4, 0xb2, 0x26, 0x4d, 0x49, 0x6a, 0x0a,
	0xcd, 0x56, 0x24, 0x4e, 0xbd, 0x14, 0x69, 0x91, 0x38, 0x50, 0x90, 0xaa, 0xbd, 0xa0, 0x55, 0xa8,
	0x10, 0x52, 0x08, 0x68, 0x92, 0x0c, 0x66, 0x14, 0xaf, 0xc7, 0xd8, 0x93, 0x44, 0x55, 0x12, 0x21,
	0x10, 0x20, 0xf5, 0xc2, 0x05, 0xee, 0x9c, 0xf8, 0x15, 0xdc, 0x10, 0x57, 0x4e, 0xdc, 0xb8, 0x44,
	0x8a, 0xfa, 0x3f, 0x40, 0x1e, 0x4f, 0xd6, 0xde, 0xc6, 0x09, 0x70, 0x29, 0x5d, 0x5f, 0x3c, 0x7e,
	0x6f, 0x66, 0xde, 0xf7, 0x7d, 0xf3, 0xde, 0xf8, 0x41, 0x25, 0xa4, 0xc1, 0x98, 0x06, 0x56, 0xfc,
	0xf2, 0x7b, 0x6a, 0xd0, 0xf4, 0x03, 0x2e, 0x38, 0x7e, 0x7e, 0x65, 0x2e, 0x57, 0x1c, 0xce, 0x1d,
	0x97, 0x5a, 0xc4, 0x67, 0x16, 0xf1, 0x3c, 0x2e, 0x88, 0x60, 0xdc, 0x0b, 0xe3, 0x79, 0xe5, 0x37,
	0x1d, 0x26, 0x3e, 0x1f, 0xf5, 0x9a, 0x7d, 0x7e, 0x66, 0x9d, 0x4d, 0x98, 0x18, 0xf2, 0x89, 0xe5,
	0xf0, 0x86, 0x74, 0x36, 0xc6, 0xc4, 0x65, 0x03, 0x22, 0x78, 0x10, 0x5a, 0xe7, 0xc3, 0x78, 0x9d,
	0xf9, 0x27, 0x82, 0xe2, 0x7b, 0x01, 0x99, 0x9c, 0x9c, 0xb6, 0xe9, 0x17, 0x23, 0x1a, 0x0a, 0xfc,
	0x25, 0x68, 0xcc, 0x37, 0x50, 0x0d, 0xd5, 0x77, 0xee, 0xf1, 0xe5, 0xa2, 0x3a, 0x04, 0xf6, 0x49,
	0xbd, 0xde, 0x69, 0x35, 0x8e, 0xbb, 0xb3, 0xce, 0x9d, 0xc6, 0x71, 0x37, 0x1e, 0xde, 0x91, 0xaf,
	0xa9, 0x3d, 0x9f, 0xd9, 0x9d, 0x56, 0xe3, 0x0d, 0x65, 0xb5, 0xef, 0x76, 0x5a, 0x8d, 0xbb, 0xdd,
	0xc3, 0x8f, 0x9b, 0x87, 0xd3, 0xa3, 0xf9, 0x7f, 0x5d, 0x75, 0xb3, 0xad, 0x31, 0x1f, 0xbf, 0x0a,
	0xb9, 0x33, 0x12, 0x0e, 0x0d, 0xad, 0x86, 0xea, 0xf9, 0x7b, 0xa5, 0xe5, 0xa2, 0x5a, 0xdc, 0xff,
	0x6b, 0xf5, 0x20, 0xe3, 0x46, 0x5b, 0xba, 0xf1, 0x0d, 0xc8, 0x09, 0xe2, 0x84, 0x86, 0x5e, 0xd3,
	0xeb, 0xbb, 0x76, 0xb1, 0xb9, 0x12, 0xaa, 0xf9, 0x80, 0x38, 0x6d, 0xe9, 0x32, 0x6b, 0xb0, 0xb7,
	0xe2, 0x16, 0xfa, 0xdc, 0x0b, 0x29, 0xde, 0x4b, 0xc8, 0x45, 0xb1, 0xcc, 0x9f, 0x10, 0x54, 0xee,
	0x53, 0xf1, 0x3e, 0x15, 0x13, 0x1e, 0x0c, 0x4f, 0xbc, 0xbe, 0x3b, 0x1a, 0x30, 0xcf, 0x79, 0x76,
	0xd4, 0x30, 0xab, 0x70, 0x7d, 0x03, 0xc0, 0x98, 0x92, 0xf9, 0x23, 0x82, 0xd2, 0x3b, 0x7d, 0xc1,
	0xc6, 0x44, 0xd0, 0x67, 0x08, 0xf7, 0x15, 0xc0, 0x69, 0x54, 0x0a, 0xec, 0x2f, 0x08, 0x4a, 0x09,
	0x9d, 0x4b, 0x96, 0x72, 0xe6, 0xaf, 0x1a, 0xe0, 0x34, 0x7a, 0x95, 0x54, 0x3f, 0x23, 0xc8, 0x33,
	0xdf, 0xa3, 0x42, 0x51, 0xf8, 0x1e, 0x2d, 0x17, 0xd5, 0x47, 0x08, 0xbe, 0x43, 0x4f, 0x8b, 0x85,
	0x55, 0x4f, 0xcf, 0x92, 0x53, 0x8e, 0xbb, 0xb3, 0xa3, 0x4e, 0xab, 0x61, 0x47, 0x14, 0x63, 0x74,
	0xf8, 0x10, 0xf6, 0x07, 0xf4, 0x33, 0x32, 0x72, 0xc5, 0xa7, 0x0e, 0x11, 0x74, 0x42, 0x1e, 0x86,
	0x86, 0x56, 0xd3, 0xeb, 0x3b, 0xed, 0x17, 0x94, 0xfd, 0xbe, 0x32, 0xe3, 0x0a, 0xec, 0xf4, 0x02,
	0x4e, 0x06, 0x7d, 0x12, 0x0a, 0x43, 0x97, 0xe5, 0x92, 0x18, 0xb0, 0x01, 0xcf, 0x79, 0x54, 0x48,
	0xc5, 0x72, 0xd2, 0xb7, 0xfa, 0x3c, 0x2f, 0xca, 0xfc, 0xe6, 0xa2, 0xfc, 0x4d, 0x83, 0x2b, 0xef,
	0x06, 0x94, 0x08, 0x7a, 0x39, 0xb3, 0x20, 0x53, 0x46, 0xfd, 0x5f, 0xc8, 0x98, 0xdb, 0x22, 0x63,
	0x3e, 0x5b, 0xc6, 0xc2, 0x66, 0x19, 0x0f, 0xe0, 0xea, 0x13, 0x2a, 0xaa, 0x12, 0x5b, 0x20, 0x28,
	0xc5, 0x9e, 0x53, 0xce, 0xdd, 0xcb, 0x26, 0xae, 0x09, 0x39, 0x9f, 0x73, 0x57, 0xe6, 0xdc, 0xae,
	0xbd, 0x97, 0x30, 0x97, 0x64, 0xa4, 0x2f, 0xba, 0x5a, 0xd2, 0x04, 0x15, 0xef, 0xb7, 0x41, 0x7f,
	0x40, 0x1c, 0x6c, 0x80, 0x3e, 0xa4, 0x0f, 0x15, 0xd3, 0xc2, 0x72, 0x51, 0xd5, 0x3e, 0x42, 0xed,
	0xc8, 0x84, 0x2b, 0x90, 0x1f, 0x13, 0x77, 0x44, 0x25, 0x84, 0xc4, 0x17, 0x1b, 0xcd, 0xdf, 0x35,
	0xc8, 0x45, 0xfb, 0xe1, 0x6f, 0x11, 0xe4, 0x43, 0x41, 0x02, 0xf1, 0x7f, 0xa9, 0x15, 0x47, 0xc7,
	0x5f, 0x21, 0xd0, 0xa9, 0x37, 0x50, 0x68, 0x9f, 0x3a, 0x8a, 0x28, 0x36, 0x7e, 0x11, 0x0a, 0xa1,
	0x20, 0x62, 0x14, 0xca, 0xf3, 0xc8, 0xb7, 0xd5, 0xd7, 0x79, 0x7e, 0xe6, 0x36, 0xe6, 0xa7, 0xfd,
	0x58, 0x87, 0x7d, 0x95, 0x9a, 0x1f, 0xd0, 0x60, 0xcc, 0xfa, 0xf4, 0xc3, 0x16, 0x76, 0xa0, 0x10,
	0xff, 0x90, 0xf1, 0x41, 0xb2, 0xe6, 0x42, 0xfb, 0x51, 0x36, 0xd6, 0x1d, 0xea, 0x80, 0x6f, 0x7f,
	0xfd, 0xc7, 0xe3, 0x1f, 0xb4, 0x9b, 0xd8, 0x94, 0x2d, 0xd0, 0xb8, 0x65, 0x79, 0xf1, 0xf6, 0xd6,
	0x94, 0xf9, 0x73, 0x6b, 0x1a, 0xa5, 0xcf, 0xdc, 0x1a, 0x04, 0x64, 0xc2, 0x7c, 0xec, 0x02, 0x24,
	0x17, 0x35, 0xbe, 0x96, 0xec, 0xb9, 0xf6, 0xf3, 0x29, 0x57, 0xb2, 0x9d, 0x2a, 0xe8, 0x2b, 0x32,
	0xe8, 0x75, 0x7c, 0x6d, 0x4b, 0x50, 0x3c, 0x83, 0xe2, 0x85, 0x5a, 0xc4, 0x2f, 0x27, 0x7b, 0x66,
	0x5d, 0x75, 0xe5, 0xea, 0x46, 0xff, 0x45, 0xae, 0xe6, 0x56, 0xae, 0x7d, 0xb9, 0x14, 0x4f, 0x01,
	0x92, 0x72, 0x48, 0x73, 0x5d, 0xbb, 0x05, 0xd2, 0x5c, 0x33, 0x2a, 0xc8, 0x96, 0x41, 0x5f, 0x37,
	0x6f, 0x6d, 0x0b, 0x1a, 0x55, 0xa0, 0x8a, 0xfc, 0x16, 0xba, 0x6d, 0x7f, 0xa3, 0xc1, 0xee, 0xc9,
	0x69, 0x72, 0xc2, 0x8f, 0x10, 0x5c, 0xcd, 0xec, 0x57, 0xf0, 0x6b, 0x59, 0x3a, 0xaf, 0x77, 0x5c,
	0xe5, 0x5b, 0xff, 0x38, 0x4f, 0xc1, 0xad, 0x4a, 0xb8, 0x2f, 0xe1, 0x83, 0x15, 0x5c, 0xe6, 0xc7,
	0x48, 0x15, 0x6c, 0xcc, 0x00, 0x92, 0x16, 0x24, 0x2d, 0xcc, 0x5a, 0xbb, 0x94, 0x16, 0x26, 0xa3,
	0x6b, 0xa9, 0xc9, 0x48, 0x65, 0xd3, 0x78, 0x32, 0x12, 0x51, 0x73, 0x7b, 0x05, 0xd9, 0x4d, 0x1f,
	0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x71, 0x84, 0x0a, 0x4c, 0xcd, 0x0b, 0x00, 0x00,
}
