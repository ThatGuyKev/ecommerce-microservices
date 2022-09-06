// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: _proto/cart.proto

package pb

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

// CartServiceClient is the client API for CartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CartServiceClient interface {
	CreateCart(ctx context.Context, in *CreateCartReq, opts ...grpc.CallOption) (*CreateCartRes, error)
	AddItemToCart(ctx context.Context, in *AddItemToCartReq, opts ...grpc.CallOption) (*AddItemToCartRes, error)
	AddBulkToCart(ctx context.Context, in *AddBulkToCartReq, opts ...grpc.CallOption) (*AddBulkToCartRes, error)
	GetCartItemsByCartId(ctx context.Context, in *GetCartItemsByCartIdReq, opts ...grpc.CallOption) (*GetCartItemsByCartIdRes, error)
	GetCartItemById(ctx context.Context, in *GetCartItemByIdReq, opts ...grpc.CallOption) (*GetCartItemByIdRes, error)
	UpdateCartItem(ctx context.Context, in *UpdateCartItemReq, opts ...grpc.CallOption) (*UpdateCartItemRes, error)
	DeleteCartItem(ctx context.Context, in *DeleteCartItemReq, opts ...grpc.CallOption) (*DeleteCartItemRes, error)
	GetCartById(ctx context.Context, in *GetCartByIdReq, opts ...grpc.CallOption) (*GetCartByIdRes, error)
	UpdateCart(ctx context.Context, in *UpdateCartReq, opts ...grpc.CallOption) (*UpdateCartRes, error)
	DeleteCart(ctx context.Context, in *DeleteCartReq, opts ...grpc.CallOption) (*DeleteCartRes, error)
	PriceCart(ctx context.Context, in *PriceCartReq, opts ...grpc.CallOption) (*PriceCartRes, error)
	ProcessCart(ctx context.Context, in *ProcessCartReq, opts ...grpc.CallOption) (*ProcessCartRes, error)
	GetCustomerCart(ctx context.Context, in *GetCustomersCartReq, opts ...grpc.CallOption) (*GetCustomersCartRes, error)
	// TODO: anonymous cart and connect offer service
	CreateGuestToken(ctx context.Context, in *CreateGuestTokenReq, opts ...grpc.CallOption) (*CreateGuestTokenRes, error)
	GetGuestToken(ctx context.Context, in *GetGuestTokenReq, opts ...grpc.CallOption) (*GetGuestTokenRes, error)
	DeleteGuestToken(ctx context.Context, in *DeleteGuestTokenReq, opts ...grpc.CallOption) (*DeleteGuestTokenRes, error)
	ValidateGuestToken(ctx context.Context, in *ValidateGuestTokenReq, opts ...grpc.CallOption) (*ValidateGuestTokenRes, error)
	AddOfferCode(ctx context.Context, in *AddOfferCodeReq, opts ...grpc.CallOption) (*AddOfferCodeRes, error)
}

type cartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCartServiceClient(cc grpc.ClientConnInterface) CartServiceClient {
	return &cartServiceClient{cc}
}

func (c *cartServiceClient) CreateCart(ctx context.Context, in *CreateCartReq, opts ...grpc.CallOption) (*CreateCartRes, error) {
	out := new(CreateCartRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/CreateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddItemToCart(ctx context.Context, in *AddItemToCartReq, opts ...grpc.CallOption) (*AddItemToCartRes, error) {
	out := new(AddItemToCartRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/AddItemToCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddBulkToCart(ctx context.Context, in *AddBulkToCartReq, opts ...grpc.CallOption) (*AddBulkToCartRes, error) {
	out := new(AddBulkToCartRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/AddBulkToCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCartItemsByCartId(ctx context.Context, in *GetCartItemsByCartIdReq, opts ...grpc.CallOption) (*GetCartItemsByCartIdRes, error) {
	out := new(GetCartItemsByCartIdRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/GetCartItemsByCartId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCartItemById(ctx context.Context, in *GetCartItemByIdReq, opts ...grpc.CallOption) (*GetCartItemByIdRes, error) {
	out := new(GetCartItemByIdRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/GetCartItemById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) UpdateCartItem(ctx context.Context, in *UpdateCartItemReq, opts ...grpc.CallOption) (*UpdateCartItemRes, error) {
	out := new(UpdateCartItemRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/UpdateCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) DeleteCartItem(ctx context.Context, in *DeleteCartItemReq, opts ...grpc.CallOption) (*DeleteCartItemRes, error) {
	out := new(DeleteCartItemRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/DeleteCartItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCartById(ctx context.Context, in *GetCartByIdReq, opts ...grpc.CallOption) (*GetCartByIdRes, error) {
	out := new(GetCartByIdRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/GetCartById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) UpdateCart(ctx context.Context, in *UpdateCartReq, opts ...grpc.CallOption) (*UpdateCartRes, error) {
	out := new(UpdateCartRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/UpdateCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) DeleteCart(ctx context.Context, in *DeleteCartReq, opts ...grpc.CallOption) (*DeleteCartRes, error) {
	out := new(DeleteCartRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/DeleteCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) PriceCart(ctx context.Context, in *PriceCartReq, opts ...grpc.CallOption) (*PriceCartRes, error) {
	out := new(PriceCartRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/PriceCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) ProcessCart(ctx context.Context, in *ProcessCartReq, opts ...grpc.CallOption) (*ProcessCartRes, error) {
	out := new(ProcessCartRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/ProcessCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetCustomerCart(ctx context.Context, in *GetCustomersCartReq, opts ...grpc.CallOption) (*GetCustomersCartRes, error) {
	out := new(GetCustomersCartRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/GetCustomerCart", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) CreateGuestToken(ctx context.Context, in *CreateGuestTokenReq, opts ...grpc.CallOption) (*CreateGuestTokenRes, error) {
	out := new(CreateGuestTokenRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/CreateGuestToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) GetGuestToken(ctx context.Context, in *GetGuestTokenReq, opts ...grpc.CallOption) (*GetGuestTokenRes, error) {
	out := new(GetGuestTokenRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/GetGuestToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) DeleteGuestToken(ctx context.Context, in *DeleteGuestTokenReq, opts ...grpc.CallOption) (*DeleteGuestTokenRes, error) {
	out := new(DeleteGuestTokenRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/DeleteGuestToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) ValidateGuestToken(ctx context.Context, in *ValidateGuestTokenReq, opts ...grpc.CallOption) (*ValidateGuestTokenRes, error) {
	out := new(ValidateGuestTokenRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/ValidateGuestToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cartServiceClient) AddOfferCode(ctx context.Context, in *AddOfferCodeReq, opts ...grpc.CallOption) (*AddOfferCodeRes, error) {
	out := new(AddOfferCodeRes)
	err := c.cc.Invoke(ctx, "/cart.CartService/AddOfferCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CartServiceServer is the server API for CartService service.
// All implementations must embed UnimplementedCartServiceServer
// for forward compatibility
type CartServiceServer interface {
	CreateCart(context.Context, *CreateCartReq) (*CreateCartRes, error)
	AddItemToCart(context.Context, *AddItemToCartReq) (*AddItemToCartRes, error)
	AddBulkToCart(context.Context, *AddBulkToCartReq) (*AddBulkToCartRes, error)
	GetCartItemsByCartId(context.Context, *GetCartItemsByCartIdReq) (*GetCartItemsByCartIdRes, error)
	GetCartItemById(context.Context, *GetCartItemByIdReq) (*GetCartItemByIdRes, error)
	UpdateCartItem(context.Context, *UpdateCartItemReq) (*UpdateCartItemRes, error)
	DeleteCartItem(context.Context, *DeleteCartItemReq) (*DeleteCartItemRes, error)
	GetCartById(context.Context, *GetCartByIdReq) (*GetCartByIdRes, error)
	UpdateCart(context.Context, *UpdateCartReq) (*UpdateCartRes, error)
	DeleteCart(context.Context, *DeleteCartReq) (*DeleteCartRes, error)
	PriceCart(context.Context, *PriceCartReq) (*PriceCartRes, error)
	ProcessCart(context.Context, *ProcessCartReq) (*ProcessCartRes, error)
	GetCustomerCart(context.Context, *GetCustomersCartReq) (*GetCustomersCartRes, error)
	// TODO: anonymous cart and connect offer service
	CreateGuestToken(context.Context, *CreateGuestTokenReq) (*CreateGuestTokenRes, error)
	GetGuestToken(context.Context, *GetGuestTokenReq) (*GetGuestTokenRes, error)
	DeleteGuestToken(context.Context, *DeleteGuestTokenReq) (*DeleteGuestTokenRes, error)
	ValidateGuestToken(context.Context, *ValidateGuestTokenReq) (*ValidateGuestTokenRes, error)
	AddOfferCode(context.Context, *AddOfferCodeReq) (*AddOfferCodeRes, error)
	mustEmbedUnimplementedCartServiceServer()
}

// UnimplementedCartServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCartServiceServer struct {
}

func (UnimplementedCartServiceServer) CreateCart(context.Context, *CreateCartReq) (*CreateCartRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCart not implemented")
}
func (UnimplementedCartServiceServer) AddItemToCart(context.Context, *AddItemToCartReq) (*AddItemToCartRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddItemToCart not implemented")
}
func (UnimplementedCartServiceServer) AddBulkToCart(context.Context, *AddBulkToCartReq) (*AddBulkToCartRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBulkToCart not implemented")
}
func (UnimplementedCartServiceServer) GetCartItemsByCartId(context.Context, *GetCartItemsByCartIdReq) (*GetCartItemsByCartIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartItemsByCartId not implemented")
}
func (UnimplementedCartServiceServer) GetCartItemById(context.Context, *GetCartItemByIdReq) (*GetCartItemByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartItemById not implemented")
}
func (UnimplementedCartServiceServer) UpdateCartItem(context.Context, *UpdateCartItemReq) (*UpdateCartItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCartItem not implemented")
}
func (UnimplementedCartServiceServer) DeleteCartItem(context.Context, *DeleteCartItemReq) (*DeleteCartItemRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCartItem not implemented")
}
func (UnimplementedCartServiceServer) GetCartById(context.Context, *GetCartByIdReq) (*GetCartByIdRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartById not implemented")
}
func (UnimplementedCartServiceServer) UpdateCart(context.Context, *UpdateCartReq) (*UpdateCartRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCart not implemented")
}
func (UnimplementedCartServiceServer) DeleteCart(context.Context, *DeleteCartReq) (*DeleteCartRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCart not implemented")
}
func (UnimplementedCartServiceServer) PriceCart(context.Context, *PriceCartReq) (*PriceCartRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PriceCart not implemented")
}
func (UnimplementedCartServiceServer) ProcessCart(context.Context, *ProcessCartReq) (*ProcessCartRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessCart not implemented")
}
func (UnimplementedCartServiceServer) GetCustomerCart(context.Context, *GetCustomersCartReq) (*GetCustomersCartRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCustomerCart not implemented")
}
func (UnimplementedCartServiceServer) CreateGuestToken(context.Context, *CreateGuestTokenReq) (*CreateGuestTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGuestToken not implemented")
}
func (UnimplementedCartServiceServer) GetGuestToken(context.Context, *GetGuestTokenReq) (*GetGuestTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGuestToken not implemented")
}
func (UnimplementedCartServiceServer) DeleteGuestToken(context.Context, *DeleteGuestTokenReq) (*DeleteGuestTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGuestToken not implemented")
}
func (UnimplementedCartServiceServer) ValidateGuestToken(context.Context, *ValidateGuestTokenReq) (*ValidateGuestTokenRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateGuestToken not implemented")
}
func (UnimplementedCartServiceServer) AddOfferCode(context.Context, *AddOfferCodeReq) (*AddOfferCodeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOfferCode not implemented")
}
func (UnimplementedCartServiceServer) mustEmbedUnimplementedCartServiceServer() {}

// UnsafeCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CartServiceServer will
// result in compilation errors.
type UnsafeCartServiceServer interface {
	mustEmbedUnimplementedCartServiceServer()
}

func RegisterCartServiceServer(s grpc.ServiceRegistrar, srv CartServiceServer) {
	s.RegisterService(&CartService_ServiceDesc, srv)
}

func _CartService_CreateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CreateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/CreateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CreateCart(ctx, req.(*CreateCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddItemToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddItemToCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddItemToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/AddItemToCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddItemToCart(ctx, req.(*AddItemToCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddBulkToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBulkToCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddBulkToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/AddBulkToCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddBulkToCart(ctx, req.(*AddBulkToCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCartItemsByCartId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartItemsByCartIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCartItemsByCartId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/GetCartItemsByCartId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCartItemsByCartId(ctx, req.(*GetCartItemsByCartIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCartItemById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartItemByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCartItemById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/GetCartItemById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCartItemById(ctx, req.(*GetCartItemByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_UpdateCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).UpdateCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/UpdateCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).UpdateCartItem(ctx, req.(*UpdateCartItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_DeleteCartItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCartItemReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).DeleteCartItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/DeleteCartItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).DeleteCartItem(ctx, req.(*DeleteCartItemReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCartById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartByIdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCartById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/GetCartById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCartById(ctx, req.(*GetCartByIdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_UpdateCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).UpdateCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/UpdateCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).UpdateCart(ctx, req.(*UpdateCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_DeleteCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).DeleteCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/DeleteCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).DeleteCart(ctx, req.(*DeleteCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_PriceCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).PriceCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/PriceCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).PriceCart(ctx, req.(*PriceCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_ProcessCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).ProcessCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/ProcessCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).ProcessCart(ctx, req.(*ProcessCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetCustomerCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCustomersCartReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetCustomerCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/GetCustomerCart",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetCustomerCart(ctx, req.(*GetCustomersCartReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_CreateGuestToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGuestTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).CreateGuestToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/CreateGuestToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).CreateGuestToken(ctx, req.(*CreateGuestTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_GetGuestToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGuestTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).GetGuestToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/GetGuestToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).GetGuestToken(ctx, req.(*GetGuestTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_DeleteGuestToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGuestTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).DeleteGuestToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/DeleteGuestToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).DeleteGuestToken(ctx, req.(*DeleteGuestTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_ValidateGuestToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateGuestTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).ValidateGuestToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/ValidateGuestToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).ValidateGuestToken(ctx, req.(*ValidateGuestTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _CartService_AddOfferCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOfferCodeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CartServiceServer).AddOfferCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cart.CartService/AddOfferCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CartServiceServer).AddOfferCode(ctx, req.(*AddOfferCodeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// CartService_ServiceDesc is the grpc.ServiceDesc for CartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cart.CartService",
	HandlerType: (*CartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCart",
			Handler:    _CartService_CreateCart_Handler,
		},
		{
			MethodName: "AddItemToCart",
			Handler:    _CartService_AddItemToCart_Handler,
		},
		{
			MethodName: "AddBulkToCart",
			Handler:    _CartService_AddBulkToCart_Handler,
		},
		{
			MethodName: "GetCartItemsByCartId",
			Handler:    _CartService_GetCartItemsByCartId_Handler,
		},
		{
			MethodName: "GetCartItemById",
			Handler:    _CartService_GetCartItemById_Handler,
		},
		{
			MethodName: "UpdateCartItem",
			Handler:    _CartService_UpdateCartItem_Handler,
		},
		{
			MethodName: "DeleteCartItem",
			Handler:    _CartService_DeleteCartItem_Handler,
		},
		{
			MethodName: "GetCartById",
			Handler:    _CartService_GetCartById_Handler,
		},
		{
			MethodName: "UpdateCart",
			Handler:    _CartService_UpdateCart_Handler,
		},
		{
			MethodName: "DeleteCart",
			Handler:    _CartService_DeleteCart_Handler,
		},
		{
			MethodName: "PriceCart",
			Handler:    _CartService_PriceCart_Handler,
		},
		{
			MethodName: "ProcessCart",
			Handler:    _CartService_ProcessCart_Handler,
		},
		{
			MethodName: "GetCustomerCart",
			Handler:    _CartService_GetCustomerCart_Handler,
		},
		{
			MethodName: "CreateGuestToken",
			Handler:    _CartService_CreateGuestToken_Handler,
		},
		{
			MethodName: "GetGuestToken",
			Handler:    _CartService_GetGuestToken_Handler,
		},
		{
			MethodName: "DeleteGuestToken",
			Handler:    _CartService_DeleteGuestToken_Handler,
		},
		{
			MethodName: "ValidateGuestToken",
			Handler:    _CartService_ValidateGuestToken_Handler,
		},
		{
			MethodName: "AddOfferCode",
			Handler:    _CartService_AddOfferCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "_proto/cart.proto",
}