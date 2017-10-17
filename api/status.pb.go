// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/tomogoma/seedms/api/status.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	github.com/tomogoma/seedms/api/status.proto

It has these top-level messages:
	Request
	Response
*/
package api

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

type Request struct {
	APIKey string `protobuf:"bytes,1,opt,name=APIKey" json:"APIKey,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetAPIKey() string {
	if m != nil {
		return m.APIKey
	}
	return ""
}

type Response struct {
	Name          string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Version       string `protobuf:"bytes,2,opt,name=version" json:"version,omitempty"`
	Description   string `protobuf:"bytes,3,opt,name=description" json:"description,omitempty"`
	CanonicalName string `protobuf:"bytes,4,opt,name=canonicalName" json:"canonicalName,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Response) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Response) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Response) GetCanonicalName() string {
	if m != nil {
		return m.CanonicalName
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "api.Request")
	proto.RegisterType((*Response)(nil), "api.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Status service

type StatusClient interface {
	Check(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type statusClient struct {
	c           client.Client
	serviceName string
}

func NewStatusClient(serviceName string, c client.Client) StatusClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "api"
	}
	return &statusClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *statusClient) Check(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Status.Check", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Status service

type StatusHandler interface {
	Check(context.Context, *Request, *Response) error
}

func RegisterStatusHandler(s server.Server, hdlr StatusHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Status{hdlr}, opts...))
}

type Status struct {
	StatusHandler
}

func (h *Status) Check(ctx context.Context, in *Request, out *Response) error {
	return h.StatusHandler.Check(ctx, in, out)
}

func init() { proto.RegisterFile("github.com/tomogoma/seedms/api/status.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0x5d, 0x5b, 0xb7, 0x3a, 0xda, 0xcb, 0x1c, 0x64, 0xf1, 0x54, 0x17, 0x11, 0x41, 0xd8,
	0x88, 0x3e, 0x81, 0x78, 0x12, 0x41, 0x64, 0x7d, 0x82, 0x69, 0x3a, 0xb4, 0x41, 0x93, 0x89, 0x3b,
	0x59, 0xc1, 0x9b, 0x8f, 0x2e, 0x8d, 0x11, 0xec, 0x2d, 0xff, 0xf7, 0xff, 0x84, 0x6f, 0xe0, 0x7a,
	0xed, 0xd2, 0x66, 0x5c, 0x76, 0x56, 0xbc, 0x49, 0xe2, 0x65, 0x2d, 0x9e, 0x8c, 0x32, 0xaf, 0xbc,
	0x1a, 0x8a, 0xce, 0x68, 0xa2, 0x34, 0x6a, 0x17, 0x07, 0x49, 0x82, 0x13, 0x8a, 0xae, 0x3d, 0x87,
	0x59, 0xcf, 0x1f, 0x23, 0x6b, 0xc2, 0x53, 0xa8, 0xef, 0x5f, 0x1e, 0x9f, 0xf8, 0xab, 0xa9, 0x16,
	0xd5, 0xd5, 0x51, 0x5f, 0x52, 0xfb, 0x5d, 0xc1, 0x61, 0xcf, 0x1a, 0x25, 0x28, 0x23, 0xc2, 0x34,
	0x90, 0xe7, 0x32, 0xc9, 0x6f, 0x6c, 0x60, 0xf6, 0xc9, 0x83, 0x3a, 0x09, 0xcd, 0x7e, 0xc6, 0x7f,
	0x11, 0x17, 0x70, 0xbc, 0x62, 0xb5, 0x83, 0x8b, 0x69, 0xdb, 0x4e, 0x72, 0xfb, 0x1f, 0xe1, 0x05,
	0xcc, 0x2d, 0x05, 0x09, 0xce, 0xd2, 0xfb, 0xf3, 0xf6, 0xe3, 0x69, 0xde, 0xec, 0xc2, 0xdb, 0x1b,
	0xa8, 0x5f, 0xb3, 0x3a, 0x5e, 0xc2, 0xc1, 0xc3, 0x86, 0xed, 0x1b, 0x9e, 0x74, 0x14, 0x5d, 0x57,
	0xdc, 0xcf, 0xe6, 0x25, 0xfd, 0x5a, 0xb6, 0x7b, 0xcb, 0x3a, 0xdf, 0x78, 0xf7, 0x13, 0x00, 0x00,
	0xff, 0xff, 0x08, 0x3f, 0xa2, 0x74, 0x12, 0x01, 0x00, 0x00,
}