// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: dish.proto

package dish

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

// DishClient is the client API for Dish service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DishClient interface {
	Add(ctx context.Context, in *NewDish, opts ...grpc.CallOption) (*NewDishResp, error)
	Read(ctx context.Context, in *ID, opts ...grpc.CallOption) (*DishInfo, error)
	Update(ctx context.Context, in *NewData, opts ...grpc.CallOption) (*UpdatedData, error)
	Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Void, error)
	Fetch(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*Dishes, error)
}

type dishClient struct {
	cc grpc.ClientConnInterface
}

func NewDishClient(cc grpc.ClientConnInterface) DishClient {
	return &dishClient{cc}
}

func (c *dishClient) Add(ctx context.Context, in *NewDish, opts ...grpc.CallOption) (*NewDishResp, error) {
	out := new(NewDishResp)
	err := c.cc.Invoke(ctx, "/dish.Dish/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dishClient) Read(ctx context.Context, in *ID, opts ...grpc.CallOption) (*DishInfo, error) {
	out := new(DishInfo)
	err := c.cc.Invoke(ctx, "/dish.Dish/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dishClient) Update(ctx context.Context, in *NewData, opts ...grpc.CallOption) (*UpdatedData, error) {
	out := new(UpdatedData)
	err := c.cc.Invoke(ctx, "/dish.Dish/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dishClient) Delete(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Void, error) {
	out := new(Void)
	err := c.cc.Invoke(ctx, "/dish.Dish/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dishClient) Fetch(ctx context.Context, in *Pagination, opts ...grpc.CallOption) (*Dishes, error) {
	out := new(Dishes)
	err := c.cc.Invoke(ctx, "/dish.Dish/Fetch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DishServer is the server API for Dish service.
// All implementations must embed UnimplementedDishServer
// for forward compatibility
type DishServer interface {
	Add(context.Context, *NewDish) (*NewDishResp, error)
	Read(context.Context, *ID) (*DishInfo, error)
	Update(context.Context, *NewData) (*UpdatedData, error)
	Delete(context.Context, *ID) (*Void, error)
	Fetch(context.Context, *Pagination) (*Dishes, error)
	mustEmbedUnimplementedDishServer()
}

// UnimplementedDishServer must be embedded to have forward compatible implementations.
type UnimplementedDishServer struct {
}

func (UnimplementedDishServer) Add(context.Context, *NewDish) (*NewDishResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedDishServer) Read(context.Context, *ID) (*DishInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedDishServer) Update(context.Context, *NewData) (*UpdatedData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDishServer) Delete(context.Context, *ID) (*Void, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedDishServer) Fetch(context.Context, *Pagination) (*Dishes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Fetch not implemented")
}
func (UnimplementedDishServer) mustEmbedUnimplementedDishServer() {}

// UnsafeDishServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DishServer will
// result in compilation errors.
type UnsafeDishServer interface {
	mustEmbedUnimplementedDishServer()
}

func RegisterDishServer(s grpc.ServiceRegistrar, srv DishServer) {
	s.RegisterService(&Dish_ServiceDesc, srv)
}

func _Dish_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewDish)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DishServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dish.Dish/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DishServer).Add(ctx, req.(*NewDish))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dish_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DishServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dish.Dish/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DishServer).Read(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dish_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewData)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DishServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dish.Dish/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DishServer).Update(ctx, req.(*NewData))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dish_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DishServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dish.Dish/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DishServer).Delete(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _Dish_Fetch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pagination)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DishServer).Fetch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dish.Dish/Fetch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DishServer).Fetch(ctx, req.(*Pagination))
	}
	return interceptor(ctx, in, info, handler)
}

// Dish_ServiceDesc is the grpc.ServiceDesc for Dish service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Dish_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "dish.Dish",
	HandlerType: (*DishServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _Dish_Add_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Dish_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Dish_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Dish_Delete_Handler,
		},
		{
			MethodName: "Fetch",
			Handler:    _Dish_Fetch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dish.proto",
}
