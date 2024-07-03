// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v4.25.3
// source: proto/order.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	OrderService_CreateOrder_FullMethodName  = "/OrderService/CreateOrder"
	OrderService_CreateOrders_FullMethodName = "/OrderService/CreateOrders"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	CreateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderService_CreateOrdersClient, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrders(ctx context.Context, opts ...grpc.CallOption) (OrderService_CreateOrdersClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OrderService_ServiceDesc.Streams[0], OrderService_CreateOrders_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &orderServiceCreateOrdersClient{ClientStream: stream}
	return x, nil
}

type OrderService_CreateOrdersClient interface {
	Send(*CreateOrderRequest) error
	Recv() (*CreateOrderResponse, error)
	grpc.ClientStream
}

type orderServiceCreateOrdersClient struct {
	grpc.ClientStream
}

func (x *orderServiceCreateOrdersClient) Send(m *CreateOrderRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *orderServiceCreateOrdersClient) Recv() (*CreateOrderResponse, error) {
	m := new(CreateOrderResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	CreateOrders(OrderService_CreateOrdersServer) error
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrders(OrderService_CreateOrdersServer) error {
	return status.Errorf(codes.Unimplemented, "method CreateOrders not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrders_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(OrderServiceServer).CreateOrders(&orderServiceCreateOrdersServer{ServerStream: stream})
}

type OrderService_CreateOrdersServer interface {
	Send(*CreateOrderResponse) error
	Recv() (*CreateOrderRequest, error)
	grpc.ServerStream
}

type orderServiceCreateOrdersServer struct {
	grpc.ServerStream
}

func (x *orderServiceCreateOrdersServer) Send(m *CreateOrderResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *orderServiceCreateOrdersServer) Recv() (*CreateOrderRequest, error) {
	m := new(CreateOrderRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CreateOrders",
			Handler:       _OrderService_CreateOrders_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/order.proto",
}
