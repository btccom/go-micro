// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: transport/grpc/proto/transport.proto

package go_micro_transport_grpc

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/btccom/go-micro/v2/api"
	client "github.com/btccom/go-micro/v2/client"
	server "github.com/btccom/go-micro/v2/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Transport service

func NewTransportEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Transport service

type TransportService interface {
	Stream(ctx context.Context, opts ...client.CallOption) (Transport_StreamService, error)
}

type transportService struct {
	c    client.Client
	name string
}

func NewTransportService(name string, c client.Client) TransportService {
	return &transportService{
		c:    c,
		name: name,
	}
}

func (c *transportService) Stream(ctx context.Context, opts ...client.CallOption) (Transport_StreamService, error) {
	req := c.c.NewRequest(c.name, "Transport.Stream", &Message{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &transportServiceStream{stream}, nil
}

type Transport_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Message) error
	Recv() (*Message, error)
}

type transportServiceStream struct {
	stream client.Stream
}

func (x *transportServiceStream) Close() error {
	return x.stream.Close()
}

func (x *transportServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *transportServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *transportServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *transportServiceStream) Send(m *Message) error {
	return x.stream.Send(m)
}

func (x *transportServiceStream) Recv() (*Message, error) {
	m := new(Message)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Transport service

type TransportHandler interface {
	Stream(context.Context, Transport_StreamStream) error
}

func RegisterTransportHandler(s server.Server, hdlr TransportHandler, opts ...server.HandlerOption) error {
	type transport interface {
		Stream(ctx context.Context, stream server.Stream) error
	}
	type Transport struct {
		transport
	}
	h := &transportHandler{hdlr}
	return s.Handle(s.NewHandler(&Transport{h}, opts...))
}

type transportHandler struct {
	TransportHandler
}

func (h *transportHandler) Stream(ctx context.Context, stream server.Stream) error {
	return h.TransportHandler.Stream(ctx, &transportStreamStream{stream})
}

type Transport_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Message) error
	Recv() (*Message, error)
}

type transportStreamStream struct {
	stream server.Stream
}

func (x *transportStreamStream) Close() error {
	return x.stream.Close()
}

func (x *transportStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *transportStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *transportStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *transportStreamStream) Send(m *Message) error {
	return x.stream.Send(m)
}

func (x *transportStreamStream) Recv() (*Message, error) {
	m := new(Message)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
