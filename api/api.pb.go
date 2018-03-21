// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	GetArtistInfoRequest
	ListArtistsRequest
	ListArtistsResponse
	ArtistInfo
	GetVenueInfoRequest
	ListVenuesRequest
	ListVenuesResponse
	VenueInfo
*/
package api

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

type GetArtistInfoRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetArtistInfoRequest) Reset()                    { *m = GetArtistInfoRequest{} }
func (m *GetArtistInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetArtistInfoRequest) ProtoMessage()               {}
func (*GetArtistInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetArtistInfoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListArtistsRequest struct {
}

func (m *ListArtistsRequest) Reset()                    { *m = ListArtistsRequest{} }
func (m *ListArtistsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListArtistsRequest) ProtoMessage()               {}
func (*ListArtistsRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ListArtistsResponse struct {
	Artists []*ArtistInfo `protobuf:"bytes,1,rep,name=artists" json:"artists,omitempty"`
}

func (m *ListArtistsResponse) Reset()                    { *m = ListArtistsResponse{} }
func (m *ListArtistsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListArtistsResponse) ProtoMessage()               {}
func (*ListArtistsResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ListArtistsResponse) GetArtists() []*ArtistInfo {
	if m != nil {
		return m.Artists
	}
	return nil
}

type ArtistInfo struct {
	Id   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ArtistInfo) Reset()                    { *m = ArtistInfo{} }
func (m *ArtistInfo) String() string            { return proto.CompactTextString(m) }
func (*ArtistInfo) ProtoMessage()               {}
func (*ArtistInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ArtistInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ArtistInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type GetVenueInfoRequest struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *GetVenueInfoRequest) Reset()                    { *m = GetVenueInfoRequest{} }
func (m *GetVenueInfoRequest) String() string            { return proto.CompactTextString(m) }
func (*GetVenueInfoRequest) ProtoMessage()               {}
func (*GetVenueInfoRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *GetVenueInfoRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ListVenuesRequest struct {
}

func (m *ListVenuesRequest) Reset()                    { *m = ListVenuesRequest{} }
func (m *ListVenuesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListVenuesRequest) ProtoMessage()               {}
func (*ListVenuesRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListVenuesResponse struct {
	Venues []*VenueInfo `protobuf:"bytes,1,rep,name=venues" json:"venues,omitempty"`
}

func (m *ListVenuesResponse) Reset()                    { *m = ListVenuesResponse{} }
func (m *ListVenuesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListVenuesResponse) ProtoMessage()               {}
func (*ListVenuesResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListVenuesResponse) GetVenues() []*VenueInfo {
	if m != nil {
		return m.Venues
	}
	return nil
}

type VenueInfo struct {
	Id       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	City     string `protobuf:"bytes,3,opt,name=city" json:"city,omitempty"`
	State    string `protobuf:"bytes,4,opt,name=state" json:"state,omitempty"`
	Capacity int32  `protobuf:"varint,5,opt,name=capacity" json:"capacity,omitempty"`
}

func (m *VenueInfo) Reset()                    { *m = VenueInfo{} }
func (m *VenueInfo) String() string            { return proto.CompactTextString(m) }
func (*VenueInfo) ProtoMessage()               {}
func (*VenueInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *VenueInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VenueInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *VenueInfo) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *VenueInfo) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *VenueInfo) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func init() {
	proto.RegisterType((*GetArtistInfoRequest)(nil), "api.GetArtistInfoRequest")
	proto.RegisterType((*ListArtistsRequest)(nil), "api.ListArtistsRequest")
	proto.RegisterType((*ListArtistsResponse)(nil), "api.ListArtistsResponse")
	proto.RegisterType((*ArtistInfo)(nil), "api.ArtistInfo")
	proto.RegisterType((*GetVenueInfoRequest)(nil), "api.GetVenueInfoRequest")
	proto.RegisterType((*ListVenuesRequest)(nil), "api.ListVenuesRequest")
	proto.RegisterType((*ListVenuesResponse)(nil), "api.ListVenuesResponse")
	proto.RegisterType((*VenueInfo)(nil), "api.VenueInfo")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for ArtistService service

type ArtistServiceClient interface {
	GetArtistInfo(ctx context.Context, in *GetArtistInfoRequest, opts ...grpc.CallOption) (*ArtistInfo, error)
	ListArtists(ctx context.Context, in *ListArtistsRequest, opts ...grpc.CallOption) (*ListArtistsResponse, error)
}

type artistServiceClient struct {
	cc *grpc.ClientConn
}

func NewArtistServiceClient(cc *grpc.ClientConn) ArtistServiceClient {
	return &artistServiceClient{cc}
}

func (c *artistServiceClient) GetArtistInfo(ctx context.Context, in *GetArtistInfoRequest, opts ...grpc.CallOption) (*ArtistInfo, error) {
	out := new(ArtistInfo)
	err := grpc.Invoke(ctx, "/api.ArtistService/GetArtistInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *artistServiceClient) ListArtists(ctx context.Context, in *ListArtistsRequest, opts ...grpc.CallOption) (*ListArtistsResponse, error) {
	out := new(ListArtistsResponse)
	err := grpc.Invoke(ctx, "/api.ArtistService/ListArtists", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ArtistService service

type ArtistServiceServer interface {
	GetArtistInfo(context.Context, *GetArtistInfoRequest) (*ArtistInfo, error)
	ListArtists(context.Context, *ListArtistsRequest) (*ListArtistsResponse, error)
}

func RegisterArtistServiceServer(s *grpc.Server, srv ArtistServiceServer) {
	s.RegisterService(&_ArtistService_serviceDesc, srv)
}

func _ArtistService_GetArtistInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetArtistInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).GetArtistInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArtistService/GetArtistInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).GetArtistInfo(ctx, req.(*GetArtistInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ArtistService_ListArtists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListArtistsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArtistServiceServer).ListArtists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ArtistService/ListArtists",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArtistServiceServer).ListArtists(ctx, req.(*ListArtistsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ArtistService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ArtistService",
	HandlerType: (*ArtistServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetArtistInfo",
			Handler:    _ArtistService_GetArtistInfo_Handler,
		},
		{
			MethodName: "ListArtists",
			Handler:    _ArtistService_ListArtists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// Client API for VenueService service

type VenueServiceClient interface {
	GetVenueInfo(ctx context.Context, in *GetVenueInfoRequest, opts ...grpc.CallOption) (*VenueInfo, error)
	ListVenues(ctx context.Context, in *ListVenuesRequest, opts ...grpc.CallOption) (*ListVenuesResponse, error)
}

type venueServiceClient struct {
	cc *grpc.ClientConn
}

func NewVenueServiceClient(cc *grpc.ClientConn) VenueServiceClient {
	return &venueServiceClient{cc}
}

func (c *venueServiceClient) GetVenueInfo(ctx context.Context, in *GetVenueInfoRequest, opts ...grpc.CallOption) (*VenueInfo, error) {
	out := new(VenueInfo)
	err := grpc.Invoke(ctx, "/api.VenueService/GetVenueInfo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *venueServiceClient) ListVenues(ctx context.Context, in *ListVenuesRequest, opts ...grpc.CallOption) (*ListVenuesResponse, error) {
	out := new(ListVenuesResponse)
	err := grpc.Invoke(ctx, "/api.VenueService/ListVenues", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VenueService service

type VenueServiceServer interface {
	GetVenueInfo(context.Context, *GetVenueInfoRequest) (*VenueInfo, error)
	ListVenues(context.Context, *ListVenuesRequest) (*ListVenuesResponse, error)
}

func RegisterVenueServiceServer(s *grpc.Server, srv VenueServiceServer) {
	s.RegisterService(&_VenueService_serviceDesc, srv)
}

func _VenueService_GetVenueInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVenueInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VenueServiceServer).GetVenueInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VenueService/GetVenueInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VenueServiceServer).GetVenueInfo(ctx, req.(*GetVenueInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VenueService_ListVenues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVenuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VenueServiceServer).ListVenues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.VenueService/ListVenues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VenueServiceServer).ListVenues(ctx, req.(*ListVenuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VenueService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.VenueService",
	HandlerType: (*VenueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVenueInfo",
			Handler:    _VenueService_GetVenueInfo_Handler,
		},
		{
			MethodName: "ListVenues",
			Handler:    _VenueService_ListVenues_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

func init() { proto.RegisterFile("api.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x86, 0x3b, 0xbd, 0x69, 0x4f, 0x2f, 0xe2, 0x69, 0xb1, 0x63, 0x56, 0x65, 0xc0, 0x52, 0x37,
	0x45, 0xea, 0x4e, 0x14, 0x2f, 0x9b, 0x22, 0xb8, 0x8a, 0xe0, 0x7e, 0x4c, 0x47, 0x98, 0x85, 0x49,
	0xcc, 0x4c, 0x0a, 0x3e, 0x83, 0x3b, 0x9f, 0x58, 0x32, 0x93, 0x7b, 0x44, 0xdc, 0xcd, 0x7c, 0x73,
	0xe0, 0x7c, 0xf9, 0xff, 0xc0, 0x80, 0x87, 0x72, 0x1d, 0x46, 0x81, 0x0e, 0xb0, 0xc3, 0x43, 0xc9,
	0x96, 0x30, 0xdb, 0x0a, 0x7d, 0x1f, 0x69, 0xa9, 0xf4, 0xa3, 0xff, 0x16, 0xb8, 0xe2, 0x23, 0x16,
	0x4a, 0xe3, 0x04, 0xda, 0x72, 0x47, 0xc9, 0x82, 0xac, 0x06, 0x6e, 0x5b, 0xee, 0xd8, 0x0c, 0xf0,
	0x49, 0xaa, 0x74, 0x50, 0xa5, 0x53, 0xec, 0x0e, 0xa6, 0x15, 0xaa, 0xc2, 0xc0, 0x57, 0x02, 0xcf,
	0xe1, 0x80, 0x5b, 0x44, 0xc9, 0xa2, 0xb3, 0x1a, 0x6e, 0x8e, 0xd6, 0xc9, 0xda, 0xd2, 0x96, 0xec,
	0x9d, 0x5d, 0x00, 0x14, 0xb8, 0xbe, 0x15, 0x11, 0xba, 0x3e, 0x7f, 0x17, 0xb4, 0x6d, 0x88, 0x39,
	0xb3, 0x33, 0x98, 0x6e, 0x85, 0x7e, 0x11, 0x7e, 0x2c, 0xfe, 0x12, 0x9e, 0xc2, 0x71, 0xa2, 0x66,
	0xe6, 0x72, 0xdf, 0x6b, 0xfb, 0x15, 0x19, 0x4c, 0x75, 0x97, 0xd0, 0xdf, 0x1b, 0x92, 0xda, 0x4e,
	0x8c, 0x6d, 0xb1, 0x21, 0x7d, 0x65, 0x31, 0x0c, 0x72, 0xf8, 0x1f, 0xd5, 0x84, 0x79, 0x52, 0x7f,
	0xd2, 0x8e, 0x65, 0xc9, 0x19, 0x67, 0xd0, 0x53, 0x9a, 0x6b, 0x41, 0xbb, 0x06, 0xda, 0x0b, 0x3a,
	0x70, 0xe8, 0xf1, 0x90, 0x9b, 0xe9, 0xde, 0x82, 0xac, 0x7a, 0x6e, 0x7e, 0xdf, 0x7c, 0x13, 0x18,
	0xdb, 0x8c, 0x9e, 0x45, 0xb4, 0x97, 0x9e, 0xc0, 0x1b, 0x18, 0x57, 0x4a, 0xc3, 0x53, 0x63, 0xfc,
	0x5b, 0x91, 0x4e, 0x3d, 0x7a, 0xd6, 0xc2, 0x07, 0x18, 0x96, 0x5a, 0xc3, 0xb9, 0x99, 0x68, 0xb6,
	0xeb, 0xd0, 0xe6, 0x83, 0x4d, 0x8c, 0xb5, 0x36, 0x5f, 0x04, 0x46, 0x26, 0x8c, 0xcc, 0xe9, 0x0a,
	0x46, 0xe5, 0x5a, 0x90, 0x66, 0x4a, 0xf5, 0xa6, 0x9c, 0x5a, 0xbc, 0xac, 0x85, 0xb7, 0x00, 0x45,
	0x2d, 0x78, 0x92, 0xaf, 0xad, 0x94, 0xe7, 0xcc, 0x1b, 0x3c, 0xb3, 0x79, 0xed, 0x9b, 0x3f, 0xfa,
	0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x14, 0x80, 0x24, 0x9e, 0xde, 0x02, 0x00, 0x00,
}
