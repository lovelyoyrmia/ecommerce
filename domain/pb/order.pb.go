// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: order.proto

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

type CartProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid         string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Stock       int32  `protobuf:"varint,4,opt,name=stock,proto3" json:"stock,omitempty"`
	Price       int64  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Category    string `protobuf:"bytes,6,opt,name=category,proto3" json:"category,omitempty"`
	Quantity    int32  `protobuf:"varint,7,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Amount      int64  `protobuf:"varint,8,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CartProduct) Reset() {
	*x = CartProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartProduct) ProtoMessage() {}

func (x *CartProduct) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartProduct.ProtoReflect.Descriptor instead.
func (*CartProduct) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{0}
}

func (x *CartProduct) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *CartProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CartProduct) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CartProduct) GetStock() int32 {
	if x != nil {
		return x.Stock
	}
	return 0
}

func (x *CartProduct) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *CartProduct) GetCategory() string {
	if x != nil {
		return x.Category
	}
	return ""
}

func (x *CartProduct) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartProduct) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CheckoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid         string `protobuf:"bytes,1,opt,name=oid,proto3" json:"oid,omitempty"`
	OrderStatus string `protobuf:"bytes,2,opt,name=order_status,json=orderStatus,proto3" json:"order_status,omitempty"`
}

func (x *CheckoutResponse) Reset() {
	*x = CheckoutResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutResponse) ProtoMessage() {}

func (x *CheckoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutResponse.ProtoReflect.Descriptor instead.
func (*CheckoutResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{1}
}

func (x *CheckoutResponse) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

func (x *CheckoutResponse) GetOrderStatus() string {
	if x != nil {
		return x.OrderStatus
	}
	return ""
}

type CreateCartParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid      string `protobuf:"bytes,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Uid      string `protobuf:"bytes,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Quantity int32  `protobuf:"varint,3,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Amount   int64  `protobuf:"varint,4,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CreateCartParams) Reset() {
	*x = CreateCartParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartParams) ProtoMessage() {}

func (x *CreateCartParams) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartParams.ProtoReflect.Descriptor instead.
func (*CreateCartParams) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCartParams) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

func (x *CreateCartParams) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *CreateCartParams) GetQuantity() int32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CreateCartParams) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CreateCartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateCartResponse) Reset() {
	*x = CreateCartResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCartResponse) ProtoMessage() {}

func (x *CreateCartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCartResponse.ProtoReflect.Descriptor instead.
func (*CreateCartResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCartResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetCartUserParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid string `protobuf:"bytes,1,opt,name=oid,proto3" json:"oid,omitempty"`
}

func (x *GetCartUserParams) Reset() {
	*x = GetCartUserParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartUserParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartUserParams) ProtoMessage() {}

func (x *GetCartUserParams) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartUserParams.ProtoReflect.Descriptor instead.
func (*GetCartUserParams) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{4}
}

func (x *GetCartUserParams) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

type GetCartUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid      string         `protobuf:"bytes,1,opt,name=oid,proto3" json:"oid,omitempty"`
	Products []*CartProduct `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GetCartUserResponse) Reset() {
	*x = GetCartUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCartUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCartUserResponse) ProtoMessage() {}

func (x *GetCartUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCartUserResponse.ProtoReflect.Descriptor instead.
func (*GetCartUserResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{5}
}

func (x *GetCartUserResponse) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

func (x *GetCartUserResponse) GetProducts() []*CartProduct {
	if x != nil {
		return x.Products
	}
	return nil
}

type DeleteCartProductParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Oid string `protobuf:"bytes,1,opt,name=oid,proto3" json:"oid,omitempty"`
	Pid string `protobuf:"bytes,2,opt,name=pid,proto3" json:"pid,omitempty"`
}

func (x *DeleteCartProductParams) Reset() {
	*x = DeleteCartProductParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartProductParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartProductParams) ProtoMessage() {}

func (x *DeleteCartProductParams) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartProductParams.ProtoReflect.Descriptor instead.
func (*DeleteCartProductParams) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCartProductParams) GetOid() string {
	if x != nil {
		return x.Oid
	}
	return ""
}

func (x *DeleteCartProductParams) GetPid() string {
	if x != nil {
		return x.Pid
	}
	return ""
}

type DeleteCartProductResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteCartProductResponse) Reset() {
	*x = DeleteCartProductResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCartProductResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCartProductResponse) ProtoMessage() {}

func (x *DeleteCartProductResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCartProductResponse.ProtoReflect.Descriptor instead.
func (*DeleteCartProductResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCartProductResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_order_proto protoreflect.FileDescriptor

var file_order_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70,
	0x62, 0x22, 0xd1, 0x01, 0x0a, 0x0b, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x70, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x6f,
	0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x10, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x6a,
	0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x70, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x2e, 0x0a, 0x12, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x25, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x43, 0x61, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x69,
	0x64, 0x22, 0x54, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6f, 0x69, 0x64, 0x12, 0x2b, 0x0a, 0x08, 0x70, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x62, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x3d, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6f, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x35, 0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x2d, 0x5a,
	0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x6f, 0x76, 0x65,
	0x6c, 0x79, 0x6f, 0x79, 0x72, 0x6d, 0x69, 0x61, 0x2f, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x72,
	0x63, 0x65, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_proto_rawDescOnce sync.Once
	file_order_proto_rawDescData = file_order_proto_rawDesc
)

func file_order_proto_rawDescGZIP() []byte {
	file_order_proto_rawDescOnce.Do(func() {
		file_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_rawDescData)
	})
	return file_order_proto_rawDescData
}

var file_order_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_order_proto_goTypes = []interface{}{
	(*CartProduct)(nil),               // 0: pb.CartProduct
	(*CheckoutResponse)(nil),          // 1: pb.CheckoutResponse
	(*CreateCartParams)(nil),          // 2: pb.CreateCartParams
	(*CreateCartResponse)(nil),        // 3: pb.CreateCartResponse
	(*GetCartUserParams)(nil),         // 4: pb.GetCartUserParams
	(*GetCartUserResponse)(nil),       // 5: pb.GetCartUserResponse
	(*DeleteCartProductParams)(nil),   // 6: pb.DeleteCartProductParams
	(*DeleteCartProductResponse)(nil), // 7: pb.DeleteCartProductResponse
}
var file_order_proto_depIdxs = []int32{
	0, // 0: pb.GetCartUserResponse.products:type_name -> pb.CartProduct
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_order_proto_init() }
func file_order_proto_init() {
	if File_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartProduct); i {
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
		file_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutResponse); i {
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
		file_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartParams); i {
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
		file_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCartResponse); i {
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
		file_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartUserParams); i {
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
		file_order_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCartUserResponse); i {
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
		file_order_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartProductParams); i {
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
		file_order_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCartProductResponse); i {
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
			RawDescriptor: file_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_proto_goTypes,
		DependencyIndexes: file_order_proto_depIdxs,
		MessageInfos:      file_order_proto_msgTypes,
	}.Build()
	File_order_proto = out.File
	file_order_proto_rawDesc = nil
	file_order_proto_goTypes = nil
	file_order_proto_depIdxs = nil
}
