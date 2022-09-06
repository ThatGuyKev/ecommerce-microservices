// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: _proto/order.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetOrdersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrdersReq) Reset() {
	*x = GetOrdersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersReq) ProtoMessage() {}

func (x *GetOrdersReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersReq.ProtoReflect.Descriptor instead.
func (*GetOrdersReq) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{0}
}

type GetOrdersRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrdersRes) Reset() {
	*x = GetOrdersRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersRes) ProtoMessage() {}

func (x *GetOrdersRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersRes.ProtoReflect.Descriptor instead.
func (*GetOrdersRes) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{1}
}

type CreateOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOrderReq) Reset() {
	*x = CreateOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderReq) ProtoMessage() {}

func (x *CreateOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderReq.ProtoReflect.Descriptor instead.
func (*CreateOrderReq) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{2}
}

type CreateOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateOrderRes) Reset() {
	*x = CreateOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRes) ProtoMessage() {}

func (x *CreateOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRes.ProtoReflect.Descriptor instead.
func (*CreateOrderRes) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{3}
}

type GetOrderByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrderByIdReq) Reset() {
	*x = GetOrderByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderByIdReq) ProtoMessage() {}

func (x *GetOrderByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderByIdReq.ProtoReflect.Descriptor instead.
func (*GetOrderByIdReq) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{4}
}

type GetOrderByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrderByIdRes) Reset() {
	*x = GetOrderByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderByIdRes) ProtoMessage() {}

func (x *GetOrderByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderByIdRes.ProtoReflect.Descriptor instead.
func (*GetOrderByIdRes) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{5}
}

type GetOrderByOrderNumReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrderByOrderNumReq) Reset() {
	*x = GetOrderByOrderNumReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderByOrderNumReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderByOrderNumReq) ProtoMessage() {}

func (x *GetOrderByOrderNumReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderByOrderNumReq.ProtoReflect.Descriptor instead.
func (*GetOrderByOrderNumReq) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{6}
}

type GetOrderByOrderNumRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrderByOrderNumRes) Reset() {
	*x = GetOrderByOrderNumRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderByOrderNumRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderByOrderNumRes) ProtoMessage() {}

func (x *GetOrderByOrderNumRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderByOrderNumRes.ProtoReflect.Descriptor instead.
func (*GetOrderByOrderNumRes) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{7}
}

type UpdateOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateOrderReq) Reset() {
	*x = UpdateOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderReq) ProtoMessage() {}

func (x *UpdateOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderReq.ProtoReflect.Descriptor instead.
func (*UpdateOrderReq) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{8}
}

type UpdateOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateOrderRes) Reset() {
	*x = UpdateOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateOrderRes) ProtoMessage() {}

func (x *UpdateOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateOrderRes.ProtoReflect.Descriptor instead.
func (*UpdateOrderRes) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{9}
}

type DeleteOrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteOrderReq) Reset() {
	*x = DeleteOrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderReq) ProtoMessage() {}

func (x *DeleteOrderReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderReq.ProtoReflect.Descriptor instead.
func (*DeleteOrderReq) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{10}
}

type DeleteOrderRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteOrderRes) Reset() {
	*x = DeleteOrderRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteOrderRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteOrderRes) ProtoMessage() {}

func (x *DeleteOrderRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteOrderRes.ProtoReflect.Descriptor instead.
func (*DeleteOrderRes) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{11}
}

type GetOrdersByCartIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrdersByCartIdReq) Reset() {
	*x = GetOrdersByCartIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersByCartIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersByCartIdReq) ProtoMessage() {}

func (x *GetOrdersByCartIdReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersByCartIdReq.ProtoReflect.Descriptor instead.
func (*GetOrdersByCartIdReq) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{12}
}

type GetOrdersByCartIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrdersByCartIdRes) Reset() {
	*x = GetOrdersByCartIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrdersByCartIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrdersByCartIdRes) ProtoMessage() {}

func (x *GetOrdersByCartIdRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrdersByCartIdRes.ProtoReflect.Descriptor instead.
func (*GetOrdersByCartIdRes) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{13}
}

type GetOrderHistoryReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrderHistoryReq) Reset() {
	*x = GetOrderHistoryReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderHistoryReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderHistoryReq) ProtoMessage() {}

func (x *GetOrderHistoryReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderHistoryReq.ProtoReflect.Descriptor instead.
func (*GetOrderHistoryReq) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{14}
}

type GetOrderHistoryRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetOrderHistoryRes) Reset() {
	*x = GetOrderHistoryRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderHistoryRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderHistoryRes) ProtoMessage() {}

func (x *GetOrderHistoryRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderHistoryRes.ProtoReflect.Descriptor instead.
func (*GetOrderHistoryRes) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{15}
}

type ChangeOrderStatusReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeOrderStatusReq) Reset() {
	*x = ChangeOrderStatusReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeOrderStatusReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeOrderStatusReq) ProtoMessage() {}

func (x *ChangeOrderStatusReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeOrderStatusReq.ProtoReflect.Descriptor instead.
func (*ChangeOrderStatusReq) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{16}
}

type ChangeOrderStatusRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeOrderStatusRes) Reset() {
	*x = ChangeOrderStatusRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_order_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeOrderStatusRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeOrderStatusRes) ProtoMessage() {}

func (x *ChangeOrderStatusRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_order_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeOrderStatusRes.ProtoReflect.Descriptor instead.
func (*ChangeOrderStatusRes) Descriptor() ([]byte, []int) {
	return file___proto_order_proto_rawDescGZIP(), []int{17}
}

var File___proto_order_proto protoreflect.FileDescriptor

var file___proto_order_proto_rawDesc = []byte{
	0x0a, 0x12, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x0e, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x22, 0x0e, 0x0a, 0x0c, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x10, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22,
	0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52,
	0x65, 0x71, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x22, 0x17,
	0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x10, 0x0a, 0x0e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x10, 0x0a, 0x0e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x22, 0x10, 0x0a,
	0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22,
	0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x22, 0x16, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x22,
	0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x22, 0x16, 0x0a, 0x14, 0x43,
	0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x71, 0x22, 0x16, 0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x32, 0x81, 0x05, 0x0a, 0x0c,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x12, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x42, 0x79, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x12, 0x1c, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65,
	0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x0b, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a,
	0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x2e,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42,
	0x79, 0x43, 0x61, 0x72, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x42, 0x79, 0x43, 0x61,
	0x72, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4f,
	0x0a, 0x11, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71,
	0x1a, 0x1b, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42,
	0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file___proto_order_proto_rawDescOnce sync.Once
	file___proto_order_proto_rawDescData = file___proto_order_proto_rawDesc
)

func file___proto_order_proto_rawDescGZIP() []byte {
	file___proto_order_proto_rawDescOnce.Do(func() {
		file___proto_order_proto_rawDescData = protoimpl.X.CompressGZIP(file___proto_order_proto_rawDescData)
	})
	return file___proto_order_proto_rawDescData
}

var file___proto_order_proto_msgTypes = make([]protoimpl.MessageInfo, 18)
var file___proto_order_proto_goTypes = []interface{}{
	(*GetOrdersReq)(nil),          // 0: order.GetOrdersReq
	(*GetOrdersRes)(nil),          // 1: order.GetOrdersRes
	(*CreateOrderReq)(nil),        // 2: order.CreateOrderReq
	(*CreateOrderRes)(nil),        // 3: order.CreateOrderRes
	(*GetOrderByIdReq)(nil),       // 4: order.GetOrderByIdReq
	(*GetOrderByIdRes)(nil),       // 5: order.GetOrderByIdRes
	(*GetOrderByOrderNumReq)(nil), // 6: order.GetOrderByOrderNumReq
	(*GetOrderByOrderNumRes)(nil), // 7: order.GetOrderByOrderNumRes
	(*UpdateOrderReq)(nil),        // 8: order.UpdateOrderReq
	(*UpdateOrderRes)(nil),        // 9: order.UpdateOrderRes
	(*DeleteOrderReq)(nil),        // 10: order.DeleteOrderReq
	(*DeleteOrderRes)(nil),        // 11: order.DeleteOrderRes
	(*GetOrdersByCartIdReq)(nil),  // 12: order.GetOrdersByCartIdReq
	(*GetOrdersByCartIdRes)(nil),  // 13: order.GetOrdersByCartIdRes
	(*GetOrderHistoryReq)(nil),    // 14: order.GetOrderHistoryReq
	(*GetOrderHistoryRes)(nil),    // 15: order.GetOrderHistoryRes
	(*ChangeOrderStatusReq)(nil),  // 16: order.ChangeOrderStatusReq
	(*ChangeOrderStatusRes)(nil),  // 17: order.ChangeOrderStatusRes
}
var file___proto_order_proto_depIdxs = []int32{
	2,  // 0: order.OrderService.CreateOrder:input_type -> order.CreateOrderReq
	0,  // 1: order.OrderService.GetOrders:input_type -> order.GetOrdersReq
	4,  // 2: order.OrderService.GetOrderById:input_type -> order.GetOrderByIdReq
	6,  // 3: order.OrderService.GetOrderByOrderNum:input_type -> order.GetOrderByOrderNumReq
	8,  // 4: order.OrderService.UpdateOrder:input_type -> order.UpdateOrderReq
	10, // 5: order.OrderService.DeleteOrder:input_type -> order.DeleteOrderReq
	12, // 6: order.OrderService.GetOrdersByCartId:input_type -> order.GetOrdersByCartIdReq
	4,  // 7: order.OrderService.GetOrderHistory:input_type -> order.GetOrderByIdReq
	16, // 8: order.OrderService.ChangeOrderStatus:input_type -> order.ChangeOrderStatusReq
	3,  // 9: order.OrderService.CreateOrder:output_type -> order.CreateOrderRes
	1,  // 10: order.OrderService.GetOrders:output_type -> order.GetOrdersRes
	5,  // 11: order.OrderService.GetOrderById:output_type -> order.GetOrderByIdRes
	7,  // 12: order.OrderService.GetOrderByOrderNum:output_type -> order.GetOrderByOrderNumRes
	9,  // 13: order.OrderService.UpdateOrder:output_type -> order.UpdateOrderRes
	11, // 14: order.OrderService.DeleteOrder:output_type -> order.DeleteOrderRes
	13, // 15: order.OrderService.GetOrdersByCartId:output_type -> order.GetOrdersByCartIdRes
	5,  // 16: order.OrderService.GetOrderHistory:output_type -> order.GetOrderByIdRes
	17, // 17: order.OrderService.ChangeOrderStatus:output_type -> order.ChangeOrderStatusRes
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file___proto_order_proto_init() }
func file___proto_order_proto_init() {
	if File___proto_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file___proto_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateOrderRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderByIdReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderByIdRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderByOrderNumReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderByOrderNumRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateOrderRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOrderReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteOrderRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersByCartIdReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrdersByCartIdRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderHistoryReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOrderHistoryRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeOrderStatusReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file___proto_order_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeOrderStatusRes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file___proto_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   18,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file___proto_order_proto_goTypes,
		DependencyIndexes: file___proto_order_proto_depIdxs,
		MessageInfos:      file___proto_order_proto_msgTypes,
	}.Build()
	File___proto_order_proto = out.File
	file___proto_order_proto_rawDesc = nil
	file___proto_order_proto_goTypes = nil
	file___proto_order_proto_depIdxs = nil
}