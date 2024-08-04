// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: api/orders_management_system/messages.proto

package orders_management_system

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// CreateOrderRequest - запрос CreateOrder
type CreateOrderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// user_id - id пользователя
	UserId uint64 `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	// items - товары в заказе
	Items []*CreateOrderRequest_SKU `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
	// delivery_info - информация о доставке
	DeliveryInfo *CreateOrderRequest_DeliveryInfo `protobuf:"bytes,3,opt,name=delivery_info,proto3" json:"delivery_info,omitempty"`
}

func (x *CreateOrderRequest) Reset() {
	*x = CreateOrderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_orders_management_system_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest) ProtoMessage() {}

func (x *CreateOrderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_orders_management_system_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest) Descriptor() ([]byte, []int) {
	return file_api_orders_management_system_messages_proto_rawDescGZIP(), []int{0}
}

func (x *CreateOrderRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CreateOrderRequest) GetItems() []*CreateOrderRequest_SKU {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CreateOrderRequest) GetDeliveryInfo() *CreateOrderRequest_DeliveryInfo {
	if x != nil {
		return x.DeliveryInfo
	}
	return nil
}

// CreateOrderResponse - ответ CreateOrder
type CreateOrderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// order_id - id созданного заказа
	OrderId string `protobuf:"bytes,1,opt,name=order_id,proto3" json:"order_id,omitempty"`
}

func (x *CreateOrderResponse) Reset() {
	*x = CreateOrderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_orders_management_system_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderResponse) ProtoMessage() {}

func (x *CreateOrderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_orders_management_system_messages_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderResponse.ProtoReflect.Descriptor instead.
func (*CreateOrderResponse) Descriptor() ([]byte, []int) {
	return file_api_orders_management_system_messages_proto_rawDescGZIP(), []int{1}
}

func (x *CreateOrderResponse) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

// SKU - товарная единица
type CreateOrderRequest_SKU struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id - id SKU
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// quantity - количество
	Quantity uint32 `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	// warehouse_id - id склада, на котором лежит данный SKU
	WarehouseId uint64 `protobuf:"varint,3,opt,name=warehouse_id,proto3" json:"warehouse_id,omitempty"`
}

func (x *CreateOrderRequest_SKU) Reset() {
	*x = CreateOrderRequest_SKU{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_orders_management_system_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest_SKU) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest_SKU) ProtoMessage() {}

func (x *CreateOrderRequest_SKU) ProtoReflect() protoreflect.Message {
	mi := &file_api_orders_management_system_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest_SKU.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest_SKU) Descriptor() ([]byte, []int) {
	return file_api_orders_management_system_messages_proto_rawDescGZIP(), []int{0, 0}
}

func (x *CreateOrderRequest_SKU) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *CreateOrderRequest_SKU) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CreateOrderRequest_SKU) GetWarehouseId() uint64 {
	if x != nil {
		return x.WarehouseId
	}
	return 0
}

// DeliveryInfo - информация о доставке
type CreateOrderRequest_DeliveryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// delivery_variant_id - id способа доставки
	DeliveryVariantId uint64 `protobuf:"varint,1,opt,name=delivery_variant_id,proto3" json:"delivery_variant_id,omitempty"`
	// delivery_date - срок доставки
	DeliveryDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=delivery_date,proto3" json:"delivery_date,omitempty"`
}

func (x *CreateOrderRequest_DeliveryInfo) Reset() {
	*x = CreateOrderRequest_DeliveryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_orders_management_system_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateOrderRequest_DeliveryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateOrderRequest_DeliveryInfo) ProtoMessage() {}

func (x *CreateOrderRequest_DeliveryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_orders_management_system_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateOrderRequest_DeliveryInfo.ProtoReflect.Descriptor instead.
func (*CreateOrderRequest_DeliveryInfo) Descriptor() ([]byte, []int) {
	return file_api_orders_management_system_messages_proto_rawDescGZIP(), []int{0, 1}
}

func (x *CreateOrderRequest_DeliveryInfo) GetDeliveryVariantId() uint64 {
	if x != nil {
		return x.DeliveryVariantId
	}
	return 0
}

func (x *CreateOrderRequest_DeliveryInfo) GetDeliveryDate() *timestamppb.Timestamp {
	if x != nil {
		return x.DeliveryDate
	}
	return nil
}

var File_api_orders_management_system_messages_proto protoreflect.FileDescriptor

var file_api_orders_management_system_messages_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x6e,
	0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3a, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x67, 0x75, 0x63, 0x68,
	0x65, 0x76, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9f, 0x06, 0x0a, 0x12, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x24, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04,
	0x42, 0x0a, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x75, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x52, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x6f, 0x67, 0x75, 0x63, 0x68, 0x65, 0x76, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x4b, 0x55, 0x42, 0x0b, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x05,
	0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x8c, 0x01, 0x0a,
	0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x5b, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2e, 0x6d, 0x6f, 0x67, 0x75, 0x63, 0x68, 0x65, 0x76, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x42, 0x09, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0d, 0x64, 0x65,
	0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x1a, 0x79, 0x0a, 0x03, 0x53,
	0x4b, 0x55, 0x12, 0x1a, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0a,
	0xe0, 0x41, 0x02, 0xba, 0x48, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26,
	0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x42, 0x0a, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2e, 0x0a, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x0a, 0xe0, 0x41,
	0x02, 0xba, 0x48, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0c, 0x77, 0x61, 0x72, 0x65, 0x68, 0x6f,
	0x75, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x1a, 0x9b, 0x01, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x13, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x42, 0x0a, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x04, 0x32, 0x02, 0x20, 0x00,
	0x52, 0x13, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x76, 0x61, 0x72, 0x69, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x4d, 0x0a, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72,
	0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0b, 0xe0, 0x41, 0x02, 0xba, 0x48, 0x05,
	0xb2, 0x01, 0x02, 0x40, 0x01, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x3a, 0xc3, 0x01, 0x92, 0x41, 0xbf, 0x01, 0x0a, 0x65, 0x2a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x32, 0x2d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x2d, 0x20, 0xd0, 0xb7, 0xd0, 0xb0, 0xd0, 0xbf, 0xd1, 0x80,
	0xd0, 0xbe, 0xd1, 0x81, 0x20, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0xd2, 0x01, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0xd2, 0x01, 0x05, 0x69, 0x74, 0x65,
	0x6d, 0x73, 0xd2, 0x01, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x2a, 0x56, 0x0a, 0x24, 0x46, 0x69, 0x6e, 0x64, 0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d,
	0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20, 0x41, 0x42, 0x69, 0x74, 0x4f, 0x66,
	0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x68, 0x74, 0x74, 0x70,
	0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0x83, 0x03, 0x0a, 0x13, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0xc8, 0x01, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0xab, 0x01, 0x92, 0x41, 0xa7, 0x01, 0x2a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x32, 0x24, 0x69, 0x64, 0x20, 0xd1, 0x81, 0xd0, 0xbe, 0xd0,
	0xb7, 0xd0, 0xb4, 0xd0, 0xb0, 0xd0, 0xbd, 0xd0, 0xbd, 0xd0, 0xbe, 0xd0, 0xb3, 0xd0, 0xbe, 0x20,
	0xd0, 0xb7, 0xd0, 0xb0, 0xd0, 0xba, 0xd0, 0xb0, 0xd0, 0xb7, 0xd0, 0xb0, 0x4a, 0x26, 0x22, 0x32,
	0x34, 0x33, 0x38, 0x61, 0x63, 0x33, 0x63, 0x2d, 0x33, 0x37, 0x65, 0x62, 0x2d, 0x34, 0x39, 0x30,
	0x32, 0x2d, 0x61, 0x64, 0x65, 0x66, 0x2d, 0x65, 0x64, 0x31, 0x36, 0x62, 0x34, 0x34, 0x33, 0x31,
	0x30, 0x33, 0x30, 0x22, 0x8a, 0x01, 0x45, 0x5e, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d,
	0x7b, 0x38, 0x7d, 0x2d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x34, 0x7d, 0x2d,
	0x34, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b, 0x38, 0x39,
	0x41, 0x42, 0x5d, 0x5b, 0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x33, 0x7d, 0x2d, 0x5b,
	0x30, 0x2d, 0x39, 0x41, 0x2d, 0x46, 0x5d, 0x7b, 0x31, 0x32, 0x7d, 0x24, 0xa2, 0x02, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x52, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x3a, 0xa0, 0x01,
	0x92, 0x41, 0x9c, 0x01, 0x0a, 0x42, 0x2a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x2b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20,
	0x2d, 0x20, 0xd0, 0xbe, 0xd1, 0x82, 0xd0, 0xb2, 0xd0, 0xb5, 0xd1, 0x82, 0x20, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x2a, 0x56, 0x0a, 0x24, 0x46, 0x69, 0x6e, 0x64,
	0x20, 0x6f, 0x75, 0x74, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x20, 0x61, 0x62, 0x6f, 0x75, 0x74, 0x20,
	0x41, 0x42, 0x69, 0x74, 0x4f, 0x66, 0x45, 0x76, 0x65, 0x72, 0x79, 0x74, 0x68, 0x69, 0x6e, 0x67,
	0x12, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x42, 0x96, 0x01, 0x5a, 0x93, 0x01, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x42, 0x61, 0x6c, 0x75, 0x6e, 0x2d, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x73, 0x2f, 0x6d,
	0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x6c, 0x69, 0x6b,
	0x65, 0x5f, 0x69, 0x6e, 0x5f, 0x62, 0x69, 0x67, 0x74, 0x65, 0x63, 0x68, 0x2f, 0x6c, 0x65, 0x63,
	0x74, 0x75, 0x72, 0x65, 0x5f, 0x36, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61,
	0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x3b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x5f, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_orders_management_system_messages_proto_rawDescOnce sync.Once
	file_api_orders_management_system_messages_proto_rawDescData = file_api_orders_management_system_messages_proto_rawDesc
)

func file_api_orders_management_system_messages_proto_rawDescGZIP() []byte {
	file_api_orders_management_system_messages_proto_rawDescOnce.Do(func() {
		file_api_orders_management_system_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_orders_management_system_messages_proto_rawDescData)
	})
	return file_api_orders_management_system_messages_proto_rawDescData
}

var file_api_orders_management_system_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_orders_management_system_messages_proto_goTypes = []any{
	(*CreateOrderRequest)(nil),              // 0: github.com.moguchev.microservices.orders_management_system.CreateOrderRequest
	(*CreateOrderResponse)(nil),             // 1: github.com.moguchev.microservices.orders_management_system.CreateOrderResponse
	(*CreateOrderRequest_SKU)(nil),          // 2: github.com.moguchev.microservices.orders_management_system.CreateOrderRequest.SKU
	(*CreateOrderRequest_DeliveryInfo)(nil), // 3: github.com.moguchev.microservices.orders_management_system.CreateOrderRequest.DeliveryInfo
	(*timestamppb.Timestamp)(nil),           // 4: google.protobuf.Timestamp
}
var file_api_orders_management_system_messages_proto_depIdxs = []int32{
	2, // 0: github.com.moguchev.microservices.orders_management_system.CreateOrderRequest.items:type_name -> github.com.moguchev.microservices.orders_management_system.CreateOrderRequest.SKU
	3, // 1: github.com.moguchev.microservices.orders_management_system.CreateOrderRequest.delivery_info:type_name -> github.com.moguchev.microservices.orders_management_system.CreateOrderRequest.DeliveryInfo
	4, // 2: github.com.moguchev.microservices.orders_management_system.CreateOrderRequest.DeliveryInfo.delivery_date:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_orders_management_system_messages_proto_init() }
func file_api_orders_management_system_messages_proto_init() {
	if File_api_orders_management_system_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_orders_management_system_messages_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOrderRequest); i {
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
		file_api_orders_management_system_messages_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOrderResponse); i {
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
		file_api_orders_management_system_messages_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOrderRequest_SKU); i {
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
		file_api_orders_management_system_messages_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateOrderRequest_DeliveryInfo); i {
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
			RawDescriptor: file_api_orders_management_system_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_orders_management_system_messages_proto_goTypes,
		DependencyIndexes: file_api_orders_management_system_messages_proto_depIdxs,
		MessageInfos:      file_api_orders_management_system_messages_proto_msgTypes,
	}.Build()
	File_api_orders_management_system_messages_proto = out.File
	file_api_orders_management_system_messages_proto_rawDesc = nil
	file_api_orders_management_system_messages_proto_goTypes = nil
	file_api_orders_management_system_messages_proto_depIdxs = nil
}
