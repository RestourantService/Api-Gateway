// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: restaurant.proto

package restaurant

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

type RestaurantDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Address     string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *RestaurantDetails) Reset() {
	*x = RestaurantDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantDetails) ProtoMessage() {}

func (x *RestaurantDetails) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantDetails.ProtoReflect.Descriptor instead.
func (*RestaurantDetails) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{0}
}

func (x *RestaurantDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RestaurantDetails) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RestaurantDetails) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RestaurantDetails) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type RestaurantInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address     string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	PhoneNumber string `protobuf:"bytes,4,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Description string `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *RestaurantInfo) Reset() {
	*x = RestaurantInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RestaurantInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RestaurantInfo) ProtoMessage() {}

func (x *RestaurantInfo) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RestaurantInfo.ProtoReflect.Descriptor instead.
func (*RestaurantInfo) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{1}
}

func (x *RestaurantInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RestaurantInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RestaurantInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RestaurantInfo) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *RestaurantInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{2}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Void struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Void) Reset() {
	*x = Void{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{3}
}

type Pagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *Pagination) Reset() {
	*x = Pagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pagination) ProtoMessage() {}

func (x *Pagination) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pagination.ProtoReflect.Descriptor instead.
func (*Pagination) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{4}
}

func (x *Pagination) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *Pagination) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type Restaurants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Restaurants []*RestaurantInfo `protobuf:"bytes,1,rep,name=restaurants,proto3" json:"restaurants,omitempty"`
}

func (x *Restaurants) Reset() {
	*x = Restaurants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_restaurant_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Restaurants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Restaurants) ProtoMessage() {}

func (x *Restaurants) ProtoReflect() protoreflect.Message {
	mi := &file_restaurant_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Restaurants.ProtoReflect.Descriptor instead.
func (*Restaurants) Descriptor() ([]byte, []int) {
	return file_restaurant_proto_rawDescGZIP(), []int{5}
}

func (x *Restaurants) GetRestaurants() []*RestaurantInfo {
	if x != nil {
		return x.Restaurants
	}
	return nil
}

var File_restaurant_proto protoreflect.FileDescriptor

var file_restaurant_proto_rawDesc = []byte{
	0x0a, 0x10, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x22, 0x86,
	0x01, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x93, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x74,
	0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x14, 0x0a,
	0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x06, 0x0a, 0x04, 0x56, 0x6f, 0x69, 0x64, 0x22, 0x3a, 0x0a, 0x0a, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x4b, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x74, 0x61,
	0x75, 0x72, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x3c, 0x0a, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x73, 0x32, 0xcd, 0x02, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x12, 0x41, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x1a, 0x0e, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x2e, 0x49, 0x44, 0x12, 0x3f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x49, 0x44, 0x1a, 0x1a, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x40, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x1a, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x10, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75,
	0x72, 0x61, 0x6e, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12, 0x34, 0x0a, 0x10, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x49, 0x44, 0x1a, 0x10, 0x2e,
	0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x56, 0x6f, 0x69, 0x64, 0x12,
	0x43, 0x0a, 0x10, 0x46, 0x65, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74,
	0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x17, 0x2e, 0x72, 0x65,
	0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72,
	0x61, 0x6e, 0x74, 0x73, 0x42, 0x15, 0x5a, 0x13, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_restaurant_proto_rawDescOnce sync.Once
	file_restaurant_proto_rawDescData = file_restaurant_proto_rawDesc
)

func file_restaurant_proto_rawDescGZIP() []byte {
	file_restaurant_proto_rawDescOnce.Do(func() {
		file_restaurant_proto_rawDescData = protoimpl.X.CompressGZIP(file_restaurant_proto_rawDescData)
	})
	return file_restaurant_proto_rawDescData
}

var file_restaurant_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_restaurant_proto_goTypes = []interface{}{
	(*RestaurantDetails)(nil), // 0: restaurant.RestaurantDetails
	(*RestaurantInfo)(nil),    // 1: restaurant.RestaurantInfo
	(*ID)(nil),                // 2: restaurant.ID
	(*Void)(nil),              // 3: restaurant.Void
	(*Pagination)(nil),        // 4: restaurant.Pagination
	(*Restaurants)(nil),       // 5: restaurant.Restaurants
}
var file_restaurant_proto_depIdxs = []int32{
	1, // 0: restaurant.Restaurants.restaurants:type_name -> restaurant.RestaurantInfo
	0, // 1: restaurant.Restaurant.CreateRestaurant:input_type -> restaurant.RestaurantDetails
	2, // 2: restaurant.Restaurant.GetRestaurantByID:input_type -> restaurant.ID
	1, // 3: restaurant.Restaurant.UpdateRestaurant:input_type -> restaurant.RestaurantInfo
	2, // 4: restaurant.Restaurant.DeleteRestaurant:input_type -> restaurant.ID
	4, // 5: restaurant.Restaurant.FetchRestaurants:input_type -> restaurant.Pagination
	2, // 6: restaurant.Restaurant.CreateRestaurant:output_type -> restaurant.ID
	1, // 7: restaurant.Restaurant.GetRestaurantByID:output_type -> restaurant.RestaurantInfo
	3, // 8: restaurant.Restaurant.UpdateRestaurant:output_type -> restaurant.Void
	3, // 9: restaurant.Restaurant.DeleteRestaurant:output_type -> restaurant.Void
	5, // 10: restaurant.Restaurant.FetchRestaurants:output_type -> restaurant.Restaurants
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_restaurant_proto_init() }
func file_restaurant_proto_init() {
	if File_restaurant_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_restaurant_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantDetails); i {
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
		file_restaurant_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RestaurantInfo); i {
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
		file_restaurant_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_restaurant_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Void); i {
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
		file_restaurant_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pagination); i {
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
		file_restaurant_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Restaurants); i {
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
			RawDescriptor: file_restaurant_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_restaurant_proto_goTypes,
		DependencyIndexes: file_restaurant_proto_depIdxs,
		MessageInfos:      file_restaurant_proto_msgTypes,
	}.Build()
	File_restaurant_proto = out.File
	file_restaurant_proto_rawDesc = nil
	file_restaurant_proto_goTypes = nil
	file_restaurant_proto_depIdxs = nil
}
