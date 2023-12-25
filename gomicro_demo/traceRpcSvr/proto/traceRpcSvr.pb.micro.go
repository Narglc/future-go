// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/traceRpcSvr.proto

package traceRpcSvr

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

// Api Endpoints for TraceRpcSvr service

func NewTraceRpcSvrEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for TraceRpcSvr service

type TraceRpcSvrService interface {
	Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error)
	ClientStream(ctx context.Context, opts ...client.CallOption) (TraceRpcSvr_ClientStreamService, error)
	ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (TraceRpcSvr_ServerStreamService, error)
	BidiStream(ctx context.Context, opts ...client.CallOption) (TraceRpcSvr_BidiStreamService, error)
}

type traceRpcSvrService struct {
	c    client.Client
	name string
}

func NewTraceRpcSvrService(name string, c client.Client) TraceRpcSvrService {
	return &traceRpcSvrService{
		c:    c,
		name: name,
	}
}

func (c *traceRpcSvrService) Call(ctx context.Context, in *CallRequest, opts ...client.CallOption) (*CallResponse, error) {
	req := c.c.NewRequest(c.name, "TraceRpcSvr.Call", in)
	out := new(CallResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceRpcSvrService) ClientStream(ctx context.Context, opts ...client.CallOption) (TraceRpcSvr_ClientStreamService, error) {
	req := c.c.NewRequest(c.name, "TraceRpcSvr.ClientStream", &ClientStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &traceRpcSvrServiceClientStream{stream}, nil
}

type TraceRpcSvr_ClientStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*ClientStreamRequest) error
}

type traceRpcSvrServiceClientStream struct {
	stream client.Stream
}

func (x *traceRpcSvrServiceClientStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *traceRpcSvrServiceClientStream) Close() error {
	return x.stream.Close()
}

func (x *traceRpcSvrServiceClientStream) Context() context.Context {
	return x.stream.Context()
}

func (x *traceRpcSvrServiceClientStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *traceRpcSvrServiceClientStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *traceRpcSvrServiceClientStream) Send(m *ClientStreamRequest) error {
	return x.stream.Send(m)
}

func (c *traceRpcSvrService) ServerStream(ctx context.Context, in *ServerStreamRequest, opts ...client.CallOption) (TraceRpcSvr_ServerStreamService, error) {
	req := c.c.NewRequest(c.name, "TraceRpcSvr.ServerStream", &ServerStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &traceRpcSvrServiceServerStream{stream}, nil
}

type TraceRpcSvr_ServerStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Recv() (*ServerStreamResponse, error)
}

type traceRpcSvrServiceServerStream struct {
	stream client.Stream
}

func (x *traceRpcSvrServiceServerStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *traceRpcSvrServiceServerStream) Close() error {
	return x.stream.Close()
}

func (x *traceRpcSvrServiceServerStream) Context() context.Context {
	return x.stream.Context()
}

func (x *traceRpcSvrServiceServerStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *traceRpcSvrServiceServerStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *traceRpcSvrServiceServerStream) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *traceRpcSvrService) BidiStream(ctx context.Context, opts ...client.CallOption) (TraceRpcSvr_BidiStreamService, error) {
	req := c.c.NewRequest(c.name, "TraceRpcSvr.BidiStream", &BidiStreamRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &traceRpcSvrServiceBidiStream{stream}, nil
}

type TraceRpcSvr_BidiStreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	CloseSend() error
	Close() error
	Send(*BidiStreamRequest) error
	Recv() (*BidiStreamResponse, error)
}

type traceRpcSvrServiceBidiStream struct {
	stream client.Stream
}

func (x *traceRpcSvrServiceBidiStream) CloseSend() error {
	return x.stream.CloseSend()
}

func (x *traceRpcSvrServiceBidiStream) Close() error {
	return x.stream.Close()
}

func (x *traceRpcSvrServiceBidiStream) Context() context.Context {
	return x.stream.Context()
}

func (x *traceRpcSvrServiceBidiStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *traceRpcSvrServiceBidiStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *traceRpcSvrServiceBidiStream) Send(m *BidiStreamRequest) error {
	return x.stream.Send(m)
}

func (x *traceRpcSvrServiceBidiStream) Recv() (*BidiStreamResponse, error) {
	m := new(BidiStreamResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for TraceRpcSvr service

type TraceRpcSvrHandler interface {
	Call(context.Context, *CallRequest, *CallResponse) error
	ClientStream(context.Context, TraceRpcSvr_ClientStreamStream) error
	ServerStream(context.Context, *ServerStreamRequest, TraceRpcSvr_ServerStreamStream) error
	BidiStream(context.Context, TraceRpcSvr_BidiStreamStream) error
}

func RegisterTraceRpcSvrHandler(s server.Server, hdlr TraceRpcSvrHandler, opts ...server.HandlerOption) error {
	type traceRpcSvr interface {
		Call(ctx context.Context, in *CallRequest, out *CallResponse) error
		ClientStream(ctx context.Context, stream server.Stream) error
		ServerStream(ctx context.Context, stream server.Stream) error
		BidiStream(ctx context.Context, stream server.Stream) error
	}
	type TraceRpcSvr struct {
		traceRpcSvr
	}
	h := &traceRpcSvrHandler{hdlr}
	return s.Handle(s.NewHandler(&TraceRpcSvr{h}, opts...))
}

type traceRpcSvrHandler struct {
	TraceRpcSvrHandler
}

func (h *traceRpcSvrHandler) Call(ctx context.Context, in *CallRequest, out *CallResponse) error {
	return h.TraceRpcSvrHandler.Call(ctx, in, out)
}

func (h *traceRpcSvrHandler) ClientStream(ctx context.Context, stream server.Stream) error {
	return h.TraceRpcSvrHandler.ClientStream(ctx, &traceRpcSvrClientStreamStream{stream})
}

type TraceRpcSvr_ClientStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*ClientStreamRequest, error)
}

type traceRpcSvrClientStreamStream struct {
	stream server.Stream
}

func (x *traceRpcSvrClientStreamStream) Close() error {
	return x.stream.Close()
}

func (x *traceRpcSvrClientStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *traceRpcSvrClientStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *traceRpcSvrClientStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *traceRpcSvrClientStreamStream) Recv() (*ClientStreamRequest, error) {
	m := new(ClientStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (h *traceRpcSvrHandler) ServerStream(ctx context.Context, stream server.Stream) error {
	m := new(ServerStreamRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.TraceRpcSvrHandler.ServerStream(ctx, m, &traceRpcSvrServerStreamStream{stream})
}

type TraceRpcSvr_ServerStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*ServerStreamResponse) error
}

type traceRpcSvrServerStreamStream struct {
	stream server.Stream
}

func (x *traceRpcSvrServerStreamStream) Close() error {
	return x.stream.Close()
}

func (x *traceRpcSvrServerStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *traceRpcSvrServerStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *traceRpcSvrServerStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *traceRpcSvrServerStreamStream) Send(m *ServerStreamResponse) error {
	return x.stream.Send(m)
}

func (h *traceRpcSvrHandler) BidiStream(ctx context.Context, stream server.Stream) error {
	return h.TraceRpcSvrHandler.BidiStream(ctx, &traceRpcSvrBidiStreamStream{stream})
}

type TraceRpcSvr_BidiStreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*BidiStreamResponse) error
	Recv() (*BidiStreamRequest, error)
}

type traceRpcSvrBidiStreamStream struct {
	stream server.Stream
}

func (x *traceRpcSvrBidiStreamStream) Close() error {
	return x.stream.Close()
}

func (x *traceRpcSvrBidiStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *traceRpcSvrBidiStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *traceRpcSvrBidiStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *traceRpcSvrBidiStreamStream) Send(m *BidiStreamResponse) error {
	return x.stream.Send(m)
}

func (x *traceRpcSvrBidiStreamStream) Recv() (*BidiStreamRequest, error) {
	m := new(BidiStreamRequest)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
