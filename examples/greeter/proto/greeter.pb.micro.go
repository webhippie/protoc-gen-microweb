// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: greeter.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Greeter service

type GreeterService interface {
	Say(ctx context.Context, in *SayRequest, opts ...client.CallOption) (*SayResponse, error)
	SayAnything(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*SayResponse, error)
}

type greeterService struct {
	c    client.Client
	name string
}

func NewGreeterService(name string, c client.Client) GreeterService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "greeter"
	}
	return &greeterService{
		c:    c,
		name: name,
	}
}

func (c *greeterService) Say(ctx context.Context, in *SayRequest, opts ...client.CallOption) (*SayResponse, error) {
	req := c.c.NewRequest(c.name, "Greeter.Say", in)
	out := new(SayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *greeterService) SayAnything(ctx context.Context, in *empty.Empty, opts ...client.CallOption) (*SayResponse, error) {
	req := c.c.NewRequest(c.name, "Greeter.SayAnything", in)
	out := new(SayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Greeter service

type GreeterHandler interface {
	Say(context.Context, *SayRequest, *SayResponse) error
	SayAnything(context.Context, *empty.Empty, *SayResponse) error
}

func RegisterGreeterHandler(s server.Server, hdlr GreeterHandler, opts ...server.HandlerOption) error {
	type greeter interface {
		Say(ctx context.Context, in *SayRequest, out *SayResponse) error
		SayAnything(ctx context.Context, in *empty.Empty, out *SayResponse) error
	}
	type Greeter struct {
		greeter
	}
	h := &greeterHandler{hdlr}
	return s.Handle(s.NewHandler(&Greeter{h}, opts...))
}

type greeterHandler struct {
	GreeterHandler
}

func (h *greeterHandler) Say(ctx context.Context, in *SayRequest, out *SayResponse) error {
	return h.GreeterHandler.Say(ctx, in, out)
}

func (h *greeterHandler) SayAnything(ctx context.Context, in *empty.Empty, out *SayResponse) error {
	return h.GreeterHandler.SayAnything(ctx, in, out)
}
