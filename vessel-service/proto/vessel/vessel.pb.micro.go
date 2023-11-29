// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: vessel.proto

package vessel

import (
	fmt "fmt"
	proto "google.golang.org/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "go-micro.dev/v4/api"
	client "go-micro.dev/v4/client"
	server "go-micro.dev/v4/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for VesselService service

func NewVesselServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for VesselService service

type VesselService interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
	Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error)
}

type vesselService struct {
	c    client.Client
	name string
}

func NewVesselService(name string, c client.Client) VesselService {
	return &vesselService{
		c:    c,
		name: name,
	}
}

func (c *vesselService) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselService) Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "VesselService.Create", in)
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
	Create(context.Context, *Vessel, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) error {
	type vesselService interface {
		FindAvailable(ctx context.Context, in *Specification, out *Response) error
		Create(ctx context.Context, in *Vessel, out *Response) error
	}
	type VesselService struct {
		vesselService
	}
	h := &vesselServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&VesselService{h}, opts...))
}

type vesselServiceHandler struct {
	VesselServiceHandler
}

func (h *vesselServiceHandler) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func (h *vesselServiceHandler) Create(ctx context.Context, in *Vessel, out *Response) error {
	return h.VesselServiceHandler.Create(ctx, in, out)
}
