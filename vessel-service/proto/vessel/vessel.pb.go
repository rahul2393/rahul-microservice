// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/vessel/vessel.proto

/*
Package go_micor_srv_vessel is a generated protocol buffer package.

It is generated from these files:
	proto/vessel/vessel.proto

It has these top-level messages:
	Vessel
	Specification
	Response
*/
package go_micor_srv_vessel

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type Vessel struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Capacity  int32  `protobuf:"varint,2,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight int32  `protobuf:"varint,3,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Available bool   `protobuf:"varint,5,opt,name=available" json:"available,omitempty"`
	OwnerId   string `protobuf:"bytes,6,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
}

func (m *Vessel) Reset()                    { *m = Vessel{} }
func (m *Vessel) String() string            { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()               {}
func (*Vessel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity  int32 `protobuf:"varint,1,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight int32 `protobuf:"varint,2,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
}

func (m *Specification) Reset()                    { *m = Specification{} }
func (m *Specification) String() string            { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()               {}
func (*Specification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel  *Vessel   `protobuf:"bytes,1,opt,name=vessel" json:"vessel,omitempty"`
	Vessels []*Vessel `protobuf:"bytes,2,rep,name=vessels" json:"vessels,omitempty"`
	Created bool      `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "go.micor.srv.vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "go.micor.srv.vessel.Specification")
	proto.RegisterType((*Response)(nil), "go.micor.srv.vessel.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micor.srv.vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func init() { proto.RegisterFile("proto/vessel/vessel.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xcf, 0x4e, 0xfa, 0x40,
	0x10, 0xc7, 0x7f, 0x5b, 0xa0, 0x94, 0xf9, 0x05, 0x0f, 0xe3, 0x65, 0x41, 0x49, 0x9a, 0x9e, 0x38,
	0xd5, 0x04, 0xe2, 0x03, 0x78, 0x31, 0xd1, 0x63, 0x31, 0x7a, 0x24, 0xcb, 0x76, 0xc4, 0x49, 0x68,
	0xb7, 0xd9, 0x6d, 0x0a, 0xbe, 0x85, 0x8f, 0xe0, 0xa3, 0x9a, 0x6c, 0x01, 0x83, 0x21, 0x7a, 0xda,
	0xf9, 0xf3, 0x9d, 0xd9, 0xef, 0x7e, 0xb2, 0x30, 0xaa, 0xac, 0xa9, 0xcd, 0x4d, 0x43, 0xce, 0xd1,
	0x66, 0x7f, 0xa4, 0xbe, 0x86, 0x97, 0x6b, 0x93, 0x16, 0xac, 0x8d, 0x4d, 0x9d, 0x6d, 0xd2, 0xb6,
	0x95, 0x7c, 0x0a, 0x08, 0x9f, 0x7d, 0x88, 0x17, 0x10, 0x70, 0x2e, 0x45, 0x2c, 0xa6, 0x83, 0x2c,
	0xe0, 0x1c, 0xc7, 0x10, 0x69, 0x55, 0x29, 0xcd, 0xf5, 0xbb, 0x0c, 0x62, 0x31, 0xed, 0x65, 0xc7,
	0x1c, 0x27, 0x00, 0x85, 0xda, 0x2d, 0xb7, 0xc4, 0xeb, 0xb7, 0x5a, 0x76, 0x7c, 0x77, 0x50, 0xa8,
	0xdd, 0x8b, 0x2f, 0x20, 0x42, 0xb7, 0x54, 0x05, 0xc9, 0xae, 0x5f, 0xe6, 0x63, 0xbc, 0x86, 0x81,
	0x6a, 0x14, 0x6f, 0xd4, 0x6a, 0x43, 0xb2, 0x17, 0x8b, 0x69, 0x94, 0x7d, 0x17, 0x70, 0x04, 0x91,
	0xd9, 0x96, 0x64, 0x97, 0x9c, 0xcb, 0xd0, 0x4f, 0xf5, 0x7d, 0xfe, 0x90, 0x27, 0x8f, 0x30, 0x5c,
	0x54, 0xa4, 0xf9, 0x95, 0xb5, 0xaa, 0xd9, 0x94, 0x27, 0xc6, 0xc4, 0xaf, 0xc6, 0x82, 0x1f, 0xc6,
	0x92, 0x0f, 0x01, 0x51, 0x46, 0xae, 0x32, 0xa5, 0x23, 0x9c, 0x43, 0xd8, 0x52, 0xf0, 0x5b, 0xfe,
	0xcf, 0xae, 0xd2, 0x33, 0x84, 0xd2, 0x96, 0x4e, 0xb6, 0x97, 0xe2, 0x2d, 0xf4, 0xdb, 0xc8, 0xc9,
	0x20, 0xee, 0xfc, 0x35, 0x75, 0xd0, 0xa2, 0x84, 0xbe, 0xb6, 0xa4, 0x6a, 0xca, 0x3d, 0xad, 0x28,
	0x3b, 0xa4, 0x33, 0x82, 0x61, 0x2b, 0x5e, 0x90, 0x6d, 0x58, 0x13, 0x3e, 0xc1, 0xf0, 0x9e, 0xcb,
	0xfc, 0xee, 0xc8, 0x26, 0x39, 0x7b, 0xc3, 0x09, 0x93, 0xf1, 0xe4, 0xac, 0xe6, 0xf0, 0xd4, 0xe4,
	0xdf, 0x2a, 0xf4, 0x9f, 0x60, 0xfe, 0x15, 0x00, 0x00, 0xff, 0xff, 0xf3, 0xee, 0x33, 0xbb, 0x21,
	0x02, 0x00, 0x00,
}
