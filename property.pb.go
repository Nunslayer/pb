// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.1
// 	protoc        v3.6.1
// source: property.proto

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

type Status int32

const (
	Status_STATUS_OK        Status = 0
	Status_STATUS_NOT_FOUND Status = 1
)

// Enum value maps for Status.
var (
	Status_name = map[int32]string{
		0: "STATUS_OK",
		1: "STATUS_NOT_FOUND",
	}
	Status_value = map[string]int32{
		"STATUS_OK":        0,
		"STATUS_NOT_FOUND": 1,
	}
)

func (x Status) Enum() *Status {
	p := new(Status)
	*p = x
	return p
}

func (x Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Status) Descriptor() protoreflect.EnumDescriptor {
	return file_property_proto_enumTypes[0].Descriptor()
}

func (Status) Type() protoreflect.EnumType {
	return &file_property_proto_enumTypes[0]
}

func (x Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Status.Descriptor instead.
func (Status) EnumDescriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{0}
}

type CheckPropertysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CartItems []*CartItem `protobuf:"bytes,1,rep,name=cart_items,json=cartItems,proto3" json:"cart_items,omitempty"`
}

func (x *CheckPropertysRequest) Reset() {
	*x = CheckPropertysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPropertysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPropertysRequest) ProtoMessage() {}

func (x *CheckPropertysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPropertysRequest.ProtoReflect.Descriptor instead.
func (*CheckPropertysRequest) Descriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{0}
}

func (x *CheckPropertysRequest) GetCartItems() []*CartItem {
	if x != nil {
		return x.CartItems
	}
	return nil
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyId uint64 `protobuf:"varint,1,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	Amount     int64  `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_property_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{1}
}

func (x *CartItem) GetPropertyId() uint64 {
	if x != nil {
		return x.PropertyId
	}
	return 0
}

func (x *CartItem) GetAmount() int64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type CheckPropertysResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyStatuses []*PropertyStatus `protobuf:"bytes,1,rep,name=property_statuses,json=propertyStatuses,proto3" json:"property_statuses,omitempty"`
}

func (x *CheckPropertysResponse) Reset() {
	*x = CheckPropertysResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckPropertysResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckPropertysResponse) ProtoMessage() {}

func (x *CheckPropertysResponse) ProtoReflect() protoreflect.Message {
	mi := &file_property_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckPropertysResponse.ProtoReflect.Descriptor instead.
func (*CheckPropertysResponse) Descriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{2}
}

func (x *CheckPropertysResponse) GetPropertyStatuses() []*PropertyStatus {
	if x != nil {
		return x.PropertyStatuses
	}
	return nil
}

type PropertyStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyId uint64 `protobuf:"varint,1,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	Price      int64  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Status     Status `protobuf:"varint,3,opt,name=status,proto3,enum=property.Status" json:"status,omitempty"`
}

func (x *PropertyStatus) Reset() {
	*x = PropertyStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PropertyStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PropertyStatus) ProtoMessage() {}

func (x *PropertyStatus) ProtoReflect() protoreflect.Message {
	mi := &file_property_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PropertyStatus.ProtoReflect.Descriptor instead.
func (*PropertyStatus) Descriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{3}
}

func (x *PropertyStatus) GetPropertyId() uint64 {
	if x != nil {
		return x.PropertyId
	}
	return 0
}

func (x *PropertyStatus) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *PropertyStatus) GetStatus() Status {
	if x != nil {
		return x.Status
	}
	return Status_STATUS_OK
}

type GetPropertysRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyIds []uint64 `protobuf:"varint,1,rep,packed,name=property_ids,json=propertyIds,proto3" json:"property_ids,omitempty"`
}

func (x *GetPropertysRequest) Reset() {
	*x = GetPropertysRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPropertysRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPropertysRequest) ProtoMessage() {}

func (x *GetPropertysRequest) ProtoReflect() protoreflect.Message {
	mi := &file_property_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPropertysRequest.ProtoReflect.Descriptor instead.
func (*GetPropertysRequest) Descriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{4}
}

func (x *GetPropertysRequest) GetPropertyIds() []uint64 {
	if x != nil {
		return x.PropertyIds
	}
	return nil
}

type Propertys struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Propertys []*Property `protobuf:"bytes,1,rep,name=propertys,proto3" json:"propertys,omitempty"`
}

func (x *Propertys) Reset() {
	*x = Propertys{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Propertys) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Propertys) ProtoMessage() {}

func (x *Propertys) ProtoReflect() protoreflect.Message {
	mi := &file_property_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Propertys.ProtoReflect.Descriptor instead.
func (*Propertys) Descriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{5}
}

func (x *Propertys) GetPropertys() []*Property {
	if x != nil {
		return x.Propertys
	}
	return nil
}

type Property struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PropertyId   uint64 `protobuf:"varint,1,opt,name=property_id,json=propertyId,proto3" json:"property_id,omitempty"`
	PropertyName string `protobuf:"bytes,2,opt,name=property_name,json=propertyName,proto3" json:"property_name,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	BrandName    string `protobuf:"bytes,4,opt,name=brand_name,json=brandName,proto3" json:"brand_name,omitempty"`
	Inventory    int64  `protobuf:"varint,5,opt,name=inventory,proto3" json:"inventory,omitempty"`
	Price        int64  `protobuf:"varint,6,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *Property) Reset() {
	*x = Property{}
	if protoimpl.UnsafeEnabled {
		mi := &file_property_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Property) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Property) ProtoMessage() {}

func (x *Property) ProtoReflect() protoreflect.Message {
	mi := &file_property_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Property.ProtoReflect.Descriptor instead.
func (*Property) Descriptor() ([]byte, []int) {
	return file_property_proto_rawDescGZIP(), []int{6}
}

func (x *Property) GetPropertyId() uint64 {
	if x != nil {
		return x.PropertyId
	}
	return 0
}

func (x *Property) GetPropertyName() string {
	if x != nil {
		return x.PropertyName
	}
	return ""
}

func (x *Property) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Property) GetBrandName() string {
	if x != nil {
		return x.BrandName
	}
	return ""
}

func (x *Property) GetInventory() int64 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

func (x *Property) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

var File_property_proto protoreflect.FileDescriptor

var file_property_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x22, 0x4a, 0x0a, 0x15, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0a, 0x63, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x63, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x43, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5f, 0x0a, 0x16, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x11, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x22, 0x71, 0x0a, 0x0e,
	0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22,
	0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0b, 0x70, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x73, 0x22, 0x3d, 0x0a, 0x09, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x73, 0x12, 0x30, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x09, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x73, 0x22, 0xc5, 0x01, 0x0a, 0x08, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70,
	0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a,
	0x0a, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x2a, 0x2d, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x4b, 0x10, 0x00, 0x12, 0x14, 0x0a, 0x10, 0x53, 0x54, 0x41,
	0x54, 0x55, 0x53, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x01, 0x32,
	0xae, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x73, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74,
	0x79, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x70,
	0x65, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x73, 0x22, 0x00,
	0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_property_proto_rawDescOnce sync.Once
	file_property_proto_rawDescData = file_property_proto_rawDesc
)

func file_property_proto_rawDescGZIP() []byte {
	file_property_proto_rawDescOnce.Do(func() {
		file_property_proto_rawDescData = protoimpl.X.CompressGZIP(file_property_proto_rawDescData)
	})
	return file_property_proto_rawDescData
}

var file_property_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_property_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_property_proto_goTypes = []interface{}{
	(Status)(0),                    // 0: property.Status
	(*CheckPropertysRequest)(nil),  // 1: property.CheckPropertysRequest
	(*CartItem)(nil),               // 2: property.CartItem
	(*CheckPropertysResponse)(nil), // 3: property.CheckPropertysResponse
	(*PropertyStatus)(nil),         // 4: property.PropertyStatus
	(*GetPropertysRequest)(nil),    // 5: property.GetPropertysRequest
	(*Propertys)(nil),              // 6: property.Propertys
	(*Property)(nil),               // 7: property.Property
}
var file_property_proto_depIdxs = []int32{
	2, // 0: property.CheckPropertysRequest.cart_items:type_name -> property.CartItem
	4, // 1: property.CheckPropertysResponse.property_statuses:type_name -> property.PropertyStatus
	0, // 2: property.PropertyStatus.status:type_name -> property.Status
	7, // 3: property.Propertys.propertys:type_name -> property.Property
	1, // 4: property.PropertyService.CheckPropertys:input_type -> property.CheckPropertysRequest
	5, // 5: property.PropertyService.GetPropertys:input_type -> property.GetPropertysRequest
	3, // 6: property.PropertyService.CheckPropertys:output_type -> property.CheckPropertysResponse
	6, // 7: property.PropertyService.GetPropertys:output_type -> property.Propertys
	6, // [6:8] is the sub-list for method output_type
	4, // [4:6] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_property_proto_init() }
func file_property_proto_init() {
	if File_property_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_property_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPropertysRequest); i {
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
		file_property_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_property_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckPropertysResponse); i {
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
		file_property_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PropertyStatus); i {
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
		file_property_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPropertysRequest); i {
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
		file_property_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Propertys); i {
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
		file_property_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Property); i {
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
			RawDescriptor: file_property_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_property_proto_goTypes,
		DependencyIndexes: file_property_proto_depIdxs,
		EnumInfos:         file_property_proto_enumTypes,
		MessageInfos:      file_property_proto_msgTypes,
	}.Build()
	File_property_proto = out.File
	file_property_proto_rawDesc = nil
	file_property_proto_goTypes = nil
	file_property_proto_depIdxs = nil
}
