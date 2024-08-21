// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: order.proto

package order

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderClient interface {
	MakeOrder(ctx context.Context, in *NewOrder, opts ...grpc.CallOption) (*NewOrderResp, error)
	ChangeStatus(ctx context.Context, in *Status, opts ...grpc.CallOption) (*UpdatedOrder, error)
	GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OrderInfo, error)
	FetchOrdersForCustomer(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*OrdersCustomer, error)
	FetchOrdersForKitchen(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*OrdersKitchen, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) MakeOrder(ctx context.Context, in *NewOrder, opts ...grpc.CallOption) (*NewOrderResp, error) {
	out := new(NewOrderResp)
	err := c.cc.Invoke(ctx, "/order.Order/MakeOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) ChangeStatus(ctx context.Context, in *Status, opts ...grpc.CallOption) (*UpdatedOrder, error) {
	out := new(UpdatedOrder)
	err := c.cc.Invoke(ctx, "/order.Order/ChangeStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) GetOrderByID(ctx context.Context, in *ID, opts ...grpc.CallOption) (*OrderInfo, error) {
	out := new(OrderInfo)
	err := c.cc.Invoke(ctx, "/order.Order/GetOrderByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) FetchOrdersForCustomer(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*OrdersCustomer, error) {
	out := new(OrdersCustomer)
	err := c.cc.Invoke(ctx, "/order.Order/FetchOrdersForCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderClient) FetchOrdersForKitchen(ctx context.Context, in *Filter, opts ...grpc.CallOption) (*OrdersKitchen, error) {
	out := new(OrdersKitchen)
	err := c.cc.Invoke(ctx, "/order.Order/FetchOrdersForKitchen", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
// All implementations must embed UnimplementedOrderServer
// for forward compatibility
type OrderServer interface {
	MakeOrder(context.Context, *NewOrder) (*NewOrderResp, error)
	ChangeStatus(context.Context, *Status) (*UpdatedOrder, error)
	GetOrderByID(context.Context, *ID) (*OrderInfo, error)
	FetchOrdersForCustomer(context.Context, *Pagination) (*OrdersCustomer, error)
	FetchOrdersForKitchen(context.Context, *Filter) (*OrdersKitchen, error)
	mustEmbedUnimplementedOrderServer()
}

// UnimplementedOrderServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (UnimplementedOrderServer) MakeOrder(context.Context, *NewOrder) (*NewOrderResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeOrder not implemented")
}
func (UnimplementedOrderServer) ChangeStatus(context.Context, *Status) (*UpdatedOrder, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatus not implemented")
}
func (UnimplementedOrderServer) GetOrderByID(context.Context, *ID) (*OrderInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderByID not implemented")
}
func (UnimplementedOrderServer) FetchOrdersForCustomer(context.Context, *Pagination) (*OrdersCustomer, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchOrdersForCustomer not implemented")
}
func (UnimplementedOrderServer) FetchOrdersForKitchen(context.Context, *Filter) (*OrdersKitchen, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchOrdersForKitchen not implemented")
}
func (UnimplementedOrderServer) mustEmbedUnimplementedOrderServer() {}

// UnsafeOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServer will
// result in compilation errors.
type UnsafeOrderServer interface {
	mustEmbedUnimplementedOrderServer()
}

func RegisterOrderServer(s grpc.ServiceRegistrar, srv OrderServer) {
	s.RegisterService(&Order_ServiceDesc, srv)
}

func _Order_MakeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).MakeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/MakeOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).MakeOrder(ctx, req.(*NewOrder))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_ChangeStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Status)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).ChangeStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/ChangeStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).ChangeStatus(ctx, req.(*Status))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_GetOrderByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrderByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/GetOrderByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrderByID(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_FetchOrdersForCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pagination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).FetchOrdersForCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/FetchOrdersForCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).FetchOrdersForCustomer(ctx, req.(*Pagination))
	}
	return interceptor(ctx, in, info, handler)
}

func _Order_FetchOrdersForKitchen_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Filter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).FetchOrdersForKitchen(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/FetchOrdersForKitchen",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).FetchOrdersForKitchen(ctx, req.(*Filter))
	}
	return interceptor(ctx, in, info, handler)
}

// Order_ServiceDesc is the grpc.ServiceDesc for Order service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Order_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeOrder",
			Handler:    _Order_MakeOrder_Handler,
		},
		{
			MethodName: "ChangeStatus",
			Handler:    _Order_ChangeStatus_Handler,
		},
		{
			MethodName: "GetOrderByID",
			Handler:    _Order_GetOrderByID_Handler,
		},
		{
			MethodName: "FetchOrdersForCustomer",
			Handler:    _Order_FetchOrdersForCustomer_Handler,
		},
		{
			MethodName: "FetchOrdersForKitchen",
			Handler:    _Order_FetchOrdersForKitchen_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}