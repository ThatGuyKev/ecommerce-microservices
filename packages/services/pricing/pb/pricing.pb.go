// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: _proto/pricing.proto

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

type GetPriceInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPriceInfoReq) Reset() {
	*x = GetPriceInfoReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPriceInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPriceInfoReq) ProtoMessage() {}

func (x *GetPriceInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPriceInfoReq.ProtoReflect.Descriptor instead.
func (*GetPriceInfoReq) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{0}
}

type GetPriceInfoRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPriceInfoRes) Reset() {
	*x = GetPriceInfoRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPriceInfoRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPriceInfoRes) ProtoMessage() {}

func (x *GetPriceInfoRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPriceInfoRes.ProtoReflect.Descriptor instead.
func (*GetPriceInfoRes) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{1}
}

type GetPricesListsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPricesListsReq) Reset() {
	*x = GetPricesListsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPricesListsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPricesListsReq) ProtoMessage() {}

func (x *GetPricesListsReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPricesListsReq.ProtoReflect.Descriptor instead.
func (*GetPricesListsReq) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{2}
}

type GetPricesListsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPricesListsRes) Reset() {
	*x = GetPricesListsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPricesListsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPricesListsRes) ProtoMessage() {}

func (x *GetPricesListsRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPricesListsRes.ProtoReflect.Descriptor instead.
func (*GetPricesListsRes) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{3}
}

type CreatePriceListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePriceListReq) Reset() {
	*x = CreatePriceListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePriceListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePriceListReq) ProtoMessage() {}

func (x *CreatePriceListReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePriceListReq.ProtoReflect.Descriptor instead.
func (*CreatePriceListReq) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{4}
}

type CreatePriceListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePriceListRes) Reset() {
	*x = CreatePriceListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePriceListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePriceListRes) ProtoMessage() {}

func (x *CreatePriceListRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePriceListRes.ProtoReflect.Descriptor instead.
func (*CreatePriceListRes) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{5}
}

type GetPricesByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPricesByIdReq) Reset() {
	*x = GetPricesByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPricesByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPricesByIdReq) ProtoMessage() {}

func (x *GetPricesByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPricesByIdReq.ProtoReflect.Descriptor instead.
func (*GetPricesByIdReq) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{6}
}

type GetPricesByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPricesByIdRes) Reset() {
	*x = GetPricesByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPricesByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPricesByIdRes) ProtoMessage() {}

func (x *GetPricesByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPricesByIdRes.ProtoReflect.Descriptor instead.
func (*GetPricesByIdRes) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{7}
}

type GetPriceDataByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPriceDataByIdReq) Reset() {
	*x = GetPriceDataByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPriceDataByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPriceDataByIdReq) ProtoMessage() {}

func (x *GetPriceDataByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPriceDataByIdReq.ProtoReflect.Descriptor instead.
func (*GetPriceDataByIdReq) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{8}
}

type GetPriceDataByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetPriceDataByIdRes) Reset() {
	*x = GetPriceDataByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPriceDataByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPriceDataByIdRes) ProtoMessage() {}

func (x *GetPriceDataByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPriceDataByIdRes.ProtoReflect.Descriptor instead.
func (*GetPriceDataByIdRes) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{9}
}

type UpdatePriceDataByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePriceDataByIdReq) Reset() {
	*x = UpdatePriceDataByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePriceDataByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePriceDataByIdReq) ProtoMessage() {}

func (x *UpdatePriceDataByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePriceDataByIdReq.ProtoReflect.Descriptor instead.
func (*UpdatePriceDataByIdReq) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{10}
}

type UpdatePriceDataByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePriceDataByIdRes) Reset() {
	*x = UpdatePriceDataByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePriceDataByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePriceDataByIdRes) ProtoMessage() {}

func (x *UpdatePriceDataByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePriceDataByIdRes.ProtoReflect.Descriptor instead.
func (*UpdatePriceDataByIdRes) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{11}
}

type DeletePriceDataByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePriceDataByIdReq) Reset() {
	*x = DeletePriceDataByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePriceDataByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePriceDataByIdReq) ProtoMessage() {}

func (x *DeletePriceDataByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePriceDataByIdReq.ProtoReflect.Descriptor instead.
func (*DeletePriceDataByIdReq) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{12}
}

type DeletePriceDataByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePriceDataByIdRes) Reset() {
	*x = DeletePriceDataByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePriceDataByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePriceDataByIdRes) ProtoMessage() {}

func (x *DeletePriceDataByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePriceDataByIdRes.ProtoReflect.Descriptor instead.
func (*DeletePriceDataByIdRes) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{13}
}

type CreatePriceDataByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePriceDataByIdReq) Reset() {
	*x = CreatePriceDataByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[14]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePriceDataByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePriceDataByIdReq) ProtoMessage() {}

func (x *CreatePriceDataByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[14]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePriceDataByIdReq.ProtoReflect.Descriptor instead.
func (*CreatePriceDataByIdReq) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{14}
}

type CreatePriceDataByIdRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreatePriceDataByIdRes) Reset() {
	*x = CreatePriceDataByIdRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[15]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePriceDataByIdRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePriceDataByIdRes) ProtoMessage() {}

func (x *CreatePriceDataByIdRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[15]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePriceDataByIdRes.ProtoReflect.Descriptor instead.
func (*CreatePriceDataByIdRes) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{15}
}

type UpdatePriceListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePriceListReq) Reset() {
	*x = UpdatePriceListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[16]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePriceListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePriceListReq) ProtoMessage() {}

func (x *UpdatePriceListReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[16]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePriceListReq.ProtoReflect.Descriptor instead.
func (*UpdatePriceListReq) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{16}
}

type UpdatePriceListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdatePriceListRes) Reset() {
	*x = UpdatePriceListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[17]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePriceListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePriceListRes) ProtoMessage() {}

func (x *UpdatePriceListRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[17]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePriceListRes.ProtoReflect.Descriptor instead.
func (*UpdatePriceListRes) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{17}
}

type DeletePriceListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePriceListReq) Reset() {
	*x = DeletePriceListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[18]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePriceListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePriceListReq) ProtoMessage() {}

func (x *DeletePriceListReq) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[18]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePriceListReq.ProtoReflect.Descriptor instead.
func (*DeletePriceListReq) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{18}
}

type DeletePriceListRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeletePriceListRes) Reset() {
	*x = DeletePriceListRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file___proto_pricing_proto_msgTypes[19]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePriceListRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePriceListRes) ProtoMessage() {}

func (x *DeletePriceListRes) ProtoReflect() protoreflect.Message {
	mi := &file___proto_pricing_proto_msgTypes[19]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePriceListRes.ProtoReflect.Descriptor instead.
func (*DeletePriceListRes) Descriptor() ([]byte, []int) {
	return file___proto_pricing_proto_rawDescGZIP(), []int{19}
}

var File___proto_pricing_proto protoreflect.FileDescriptor

var file___proto_pricing_proto_rawDesc = []byte{
	0x0a, 0x14, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x22,
	0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x22, 0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x73, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x22,
	0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x12, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x22, 0x18, 0x0a, 0x16, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x22,
	0x18, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x22, 0x14, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x22, 0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x22,
	0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x32, 0xbb, 0x06, 0x0a, 0x0e, 0x50, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x71, 0x1a, 0x18, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4a,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73,
	0x12, 0x1a, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x70,
	0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e,
	0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0d, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x69,
	0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79,
	0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x73, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73,
	0x22, 0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x4d, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00,
	0x12, 0x59, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e,
	0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69,
	0x6e, 0x67, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79,
	0x49, 0x64, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x59, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1f, 0x2e,
	0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x1f,
	0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x50, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67, 0x2e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x22, 0x00, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x70, 0x72, 0x69, 0x63, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file___proto_pricing_proto_rawDescOnce sync.Once
	file___proto_pricing_proto_rawDescData = file___proto_pricing_proto_rawDesc
)

func file___proto_pricing_proto_rawDescGZIP() []byte {
	file___proto_pricing_proto_rawDescOnce.Do(func() {
		file___proto_pricing_proto_rawDescData = protoimpl.X.CompressGZIP(file___proto_pricing_proto_rawDescData)
	})
	return file___proto_pricing_proto_rawDescData
}

var file___proto_pricing_proto_msgTypes = make([]protoimpl.MessageInfo, 20)
var file___proto_pricing_proto_goTypes = []interface{}{
	(*GetPriceInfoReq)(nil),        // 0: pricing.GetPriceInfoReq
	(*GetPriceInfoRes)(nil),        // 1: pricing.GetPriceInfoRes
	(*GetPricesListsReq)(nil),      // 2: pricing.GetPricesListsReq
	(*GetPricesListsRes)(nil),      // 3: pricing.GetPricesListsRes
	(*CreatePriceListReq)(nil),     // 4: pricing.CreatePriceListReq
	(*CreatePriceListRes)(nil),     // 5: pricing.CreatePriceListRes
	(*GetPricesByIdReq)(nil),       // 6: pricing.GetPricesByIdReq
	(*GetPricesByIdRes)(nil),       // 7: pricing.GetPricesByIdRes
	(*GetPriceDataByIdReq)(nil),    // 8: pricing.GetPriceDataByIdReq
	(*GetPriceDataByIdRes)(nil),    // 9: pricing.GetPriceDataByIdRes
	(*UpdatePriceDataByIdReq)(nil), // 10: pricing.UpdatePriceDataByIdReq
	(*UpdatePriceDataByIdRes)(nil), // 11: pricing.UpdatePriceDataByIdRes
	(*DeletePriceDataByIdReq)(nil), // 12: pricing.DeletePriceDataByIdReq
	(*DeletePriceDataByIdRes)(nil), // 13: pricing.DeletePriceDataByIdRes
	(*CreatePriceDataByIdReq)(nil), // 14: pricing.CreatePriceDataByIdReq
	(*CreatePriceDataByIdRes)(nil), // 15: pricing.CreatePriceDataByIdRes
	(*UpdatePriceListReq)(nil),     // 16: pricing.UpdatePriceListReq
	(*UpdatePriceListRes)(nil),     // 17: pricing.UpdatePriceListRes
	(*DeletePriceListReq)(nil),     // 18: pricing.DeletePriceListReq
	(*DeletePriceListRes)(nil),     // 19: pricing.DeletePriceListRes
}
var file___proto_pricing_proto_depIdxs = []int32{
	0,  // 0: pricing.PricingService.GetPriceInfo:input_type -> pricing.GetPriceInfoReq
	2,  // 1: pricing.PricingService.GetPricesLists:input_type -> pricing.GetPricesListsReq
	4,  // 2: pricing.PricingService.CreatePriceList:input_type -> pricing.CreatePriceListReq
	6,  // 3: pricing.PricingService.GetPricesById:input_type -> pricing.GetPricesByIdReq
	18, // 4: pricing.PricingService.DeletePriceList:input_type -> pricing.DeletePriceListReq
	16, // 5: pricing.PricingService.UpdatePriceList:input_type -> pricing.UpdatePriceListReq
	14, // 6: pricing.PricingService.CreatePriceDataById:input_type -> pricing.CreatePriceDataByIdReq
	12, // 7: pricing.PricingService.DeletePriceDataById:input_type -> pricing.DeletePriceDataByIdReq
	10, // 8: pricing.PricingService.UpdatePriceDataById:input_type -> pricing.UpdatePriceDataByIdReq
	8,  // 9: pricing.PricingService.GetPriceDataById:input_type -> pricing.GetPriceDataByIdReq
	1,  // 10: pricing.PricingService.GetPriceInfo:output_type -> pricing.GetPriceInfoRes
	3,  // 11: pricing.PricingService.GetPricesLists:output_type -> pricing.GetPricesListsRes
	5,  // 12: pricing.PricingService.CreatePriceList:output_type -> pricing.CreatePriceListRes
	7,  // 13: pricing.PricingService.GetPricesById:output_type -> pricing.GetPricesByIdRes
	19, // 14: pricing.PricingService.DeletePriceList:output_type -> pricing.DeletePriceListRes
	17, // 15: pricing.PricingService.UpdatePriceList:output_type -> pricing.UpdatePriceListRes
	15, // 16: pricing.PricingService.CreatePriceDataById:output_type -> pricing.CreatePriceDataByIdRes
	13, // 17: pricing.PricingService.DeletePriceDataById:output_type -> pricing.DeletePriceDataByIdRes
	11, // 18: pricing.PricingService.UpdatePriceDataById:output_type -> pricing.UpdatePriceDataByIdRes
	9,  // 19: pricing.PricingService.GetPriceDataById:output_type -> pricing.GetPriceDataByIdRes
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file___proto_pricing_proto_init() }
func file___proto_pricing_proto_init() {
	if File___proto_pricing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file___proto_pricing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPriceInfoReq); i {
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
		file___proto_pricing_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPriceInfoRes); i {
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
		file___proto_pricing_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPricesListsReq); i {
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
		file___proto_pricing_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPricesListsRes); i {
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
		file___proto_pricing_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePriceListReq); i {
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
		file___proto_pricing_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePriceListRes); i {
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
		file___proto_pricing_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPricesByIdReq); i {
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
		file___proto_pricing_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPricesByIdRes); i {
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
		file___proto_pricing_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPriceDataByIdReq); i {
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
		file___proto_pricing_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPriceDataByIdRes); i {
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
		file___proto_pricing_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePriceDataByIdReq); i {
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
		file___proto_pricing_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePriceDataByIdRes); i {
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
		file___proto_pricing_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePriceDataByIdReq); i {
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
		file___proto_pricing_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePriceDataByIdRes); i {
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
		file___proto_pricing_proto_msgTypes[14].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePriceDataByIdReq); i {
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
		file___proto_pricing_proto_msgTypes[15].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePriceDataByIdRes); i {
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
		file___proto_pricing_proto_msgTypes[16].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePriceListReq); i {
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
		file___proto_pricing_proto_msgTypes[17].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePriceListRes); i {
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
		file___proto_pricing_proto_msgTypes[18].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePriceListReq); i {
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
		file___proto_pricing_proto_msgTypes[19].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePriceListRes); i {
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
			RawDescriptor: file___proto_pricing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   20,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file___proto_pricing_proto_goTypes,
		DependencyIndexes: file___proto_pricing_proto_depIdxs,
		MessageInfos:      file___proto_pricing_proto_msgTypes,
	}.Build()
	File___proto_pricing_proto = out.File
	file___proto_pricing_proto_rawDesc = nil
	file___proto_pricing_proto_goTypes = nil
	file___proto_pricing_proto_depIdxs = nil
}