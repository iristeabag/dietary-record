// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: food.proto

package food

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetFoodByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetFoodByIdRequest) Reset() {
	*x = GetFoodByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFoodByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFoodByIdRequest) ProtoMessage() {}

func (x *GetFoodByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFoodByIdRequest.ProtoReflect.Descriptor instead.
func (*GetFoodByIdRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{0}
}

func (x *GetFoodByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetFoodByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foodid  string  `protobuf:"bytes,1,opt,name=Foodid,proto3" json:"Foodid,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Brand   string  `protobuf:"bytes,3,opt,name=Brand,proto3" json:"Brand,omitempty"`
	Amount  float32 `protobuf:"fixed32,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Unit    string  `protobuf:"bytes,5,opt,name=Unit,proto3" json:"Unit,omitempty"`
	Carb    float32 `protobuf:"fixed32,6,opt,name=Carb,proto3" json:"Carb,omitempty"`
	Portein float32 `protobuf:"fixed32,7,opt,name=Portein,proto3" json:"Portein,omitempty"`
	Fat     float32 `protobuf:"fixed32,8,opt,name=Fat,proto3" json:"Fat,omitempty"`
	Cal     float32 `protobuf:"fixed32,9,opt,name=Cal,proto3" json:"Cal,omitempty"`
}

func (x *GetFoodByIdResponse) Reset() {
	*x = GetFoodByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFoodByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFoodByIdResponse) ProtoMessage() {}

func (x *GetFoodByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFoodByIdResponse.ProtoReflect.Descriptor instead.
func (*GetFoodByIdResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{1}
}

func (x *GetFoodByIdResponse) GetFoodid() string {
	if x != nil {
		return x.Foodid
	}
	return ""
}

func (x *GetFoodByIdResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetFoodByIdResponse) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *GetFoodByIdResponse) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *GetFoodByIdResponse) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *GetFoodByIdResponse) GetCarb() float32 {
	if x != nil {
		return x.Carb
	}
	return 0
}

func (x *GetFoodByIdResponse) GetPortein() float32 {
	if x != nil {
		return x.Portein
	}
	return 0
}

func (x *GetFoodByIdResponse) GetFat() float32 {
	if x != nil {
		return x.Fat
	}
	return 0
}

func (x *GetFoodByIdResponse) GetCal() float32 {
	if x != nil {
		return x.Cal
	}
	return 0
}

type CreateFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string  `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Brand   string  `protobuf:"bytes,2,opt,name=Brand,proto3" json:"Brand,omitempty"`
	Amount  float32 `protobuf:"fixed32,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Unit    string  `protobuf:"bytes,4,opt,name=Unit,proto3" json:"Unit,omitempty"`
	Carb    float32 `protobuf:"fixed32,5,opt,name=Carb,proto3" json:"Carb,omitempty"`
	Portein float32 `protobuf:"fixed32,6,opt,name=Portein,proto3" json:"Portein,omitempty"`
	Fat     float32 `protobuf:"fixed32,7,opt,name=Fat,proto3" json:"Fat,omitempty"`
	Cal     float32 `protobuf:"fixed32,8,opt,name=Cal,proto3" json:"Cal,omitempty"`
}

func (x *CreateFoodRequest) Reset() {
	*x = CreateFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateFoodRequest) ProtoMessage() {}

func (x *CreateFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateFoodRequest.ProtoReflect.Descriptor instead.
func (*CreateFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{2}
}

func (x *CreateFoodRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateFoodRequest) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *CreateFoodRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *CreateFoodRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *CreateFoodRequest) GetCarb() float32 {
	if x != nil {
		return x.Carb
	}
	return 0
}

func (x *CreateFoodRequest) GetPortein() float32 {
	if x != nil {
		return x.Portein
	}
	return 0
}

func (x *CreateFoodRequest) GetFat() float32 {
	if x != nil {
		return x.Fat
	}
	return 0
}

func (x *CreateFoodRequest) GetCal() float32 {
	if x != nil {
		return x.Cal
	}
	return 0
}

type UpdateFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Foodid  string  `protobuf:"bytes,1,opt,name=Foodid,proto3" json:"Foodid,omitempty"`
	Name    string  `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Brand   string  `protobuf:"bytes,3,opt,name=Brand,proto3" json:"Brand,omitempty"`
	Amount  float32 `protobuf:"fixed32,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Unit    string  `protobuf:"bytes,5,opt,name=Unit,proto3" json:"Unit,omitempty"`
	Carb    float32 `protobuf:"fixed32,6,opt,name=Carb,proto3" json:"Carb,omitempty"`
	Portein float32 `protobuf:"fixed32,7,opt,name=Portein,proto3" json:"Portein,omitempty"`
	Fat     float32 `protobuf:"fixed32,8,opt,name=Fat,proto3" json:"Fat,omitempty"`
	Cal     float32 `protobuf:"fixed32,9,opt,name=Cal,proto3" json:"Cal,omitempty"`
}

func (x *UpdateFoodRequest) Reset() {
	*x = UpdateFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFoodRequest) ProtoMessage() {}

func (x *UpdateFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFoodRequest.ProtoReflect.Descriptor instead.
func (*UpdateFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateFoodRequest) GetFoodid() string {
	if x != nil {
		return x.Foodid
	}
	return ""
}

func (x *UpdateFoodRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateFoodRequest) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *UpdateFoodRequest) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *UpdateFoodRequest) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *UpdateFoodRequest) GetCarb() float32 {
	if x != nil {
		return x.Carb
	}
	return 0
}

func (x *UpdateFoodRequest) GetPortein() float32 {
	if x != nil {
		return x.Portein
	}
	return 0
}

func (x *UpdateFoodRequest) GetFat() float32 {
	if x != nil {
		return x.Fat
	}
	return 0
}

func (x *UpdateFoodRequest) GetCal() float32 {
	if x != nil {
		return x.Cal
	}
	return 0
}

type DefaultResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *DefaultResponse) Reset() {
	*x = DefaultResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultResponse) ProtoMessage() {}

func (x *DefaultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DefaultResponse.ProtoReflect.Descriptor instead.
func (*DefaultResponse) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{4}
}

func (x *DefaultResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type DeleteFoodRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *DeleteFoodRequest) Reset() {
	*x = DeleteFoodRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_food_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFoodRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFoodRequest) ProtoMessage() {}

func (x *DeleteFoodRequest) ProtoReflect() protoreflect.Message {
	mi := &file_food_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFoodRequest.ProtoReflect.Descriptor instead.
func (*DeleteFoodRequest) Descriptor() ([]byte, []int) {
	return file_food_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteFoodRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_food_proto protoreflect.FileDescriptor

var file_food_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6f,
	0x6f, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x49,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0xd5, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x46, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x46, 0x6f, 0x6f, 0x64, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x46, 0x6f, 0x6f, 0x64, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x42, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x42, 0x72, 0x61,
	0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x02, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e,
	0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x43, 0x61, 0x72, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x43, 0x61,
	0x72, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6e, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x07, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03,
	0x46, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x46, 0x61, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x43, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x43, 0x61, 0x6c,
	0x22, 0xbb, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x42, 0x72,
	0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64,
	0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x43, 0x61, 0x72, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x43, 0x61, 0x72, 0x62,
	0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x07, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x46, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x46, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03,
	0x43, 0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x43, 0x61, 0x6c, 0x22, 0xd3,
	0x01, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x6f, 0x6f, 0x64, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x6f, 0x6f, 0x64, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x6e,
	0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x62, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x43, 0x61, 0x72, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x69,
	0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x69, 0x6e,
	0x12, 0x10, 0x0a, 0x03, 0x46, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x46,
	0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x03, 0x43, 0x61, 0x6c, 0x22, 0x29, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x49, 0x64, 0x32, 0x93, 0x02, 0x0a, 0x0b, 0x46, 0x6f, 0x6f, 0x64, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x42,
	0x79, 0x49, 0x64, 0x12, 0x18, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f,
	0x6f, 0x64, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e,
	0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x47, 0x65, 0x74, 0x46, 0x6f, 0x6f, 0x64, 0x42, 0x79, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x12, 0x17, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x6f, 0x6f, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x15, 0x2e, 0x66, 0x6f, 0x6f, 0x64, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6f, 0x6f, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_food_proto_rawDescOnce sync.Once
	file_food_proto_rawDescData = file_food_proto_rawDesc
)

func file_food_proto_rawDescGZIP() []byte {
	file_food_proto_rawDescOnce.Do(func() {
		file_food_proto_rawDescData = protoimpl.X.CompressGZIP(file_food_proto_rawDescData)
	})
	return file_food_proto_rawDescData
}

var file_food_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_food_proto_goTypes = []interface{}{
	(*GetFoodByIdRequest)(nil),  // 0: food.GetFoodByIdRequest
	(*GetFoodByIdResponse)(nil), // 1: food.GetFoodByIdResponse
	(*CreateFoodRequest)(nil),   // 2: food.CreateFoodRequest
	(*UpdateFoodRequest)(nil),   // 3: food.UpdateFoodRequest
	(*DefaultResponse)(nil),     // 4: food.DefaultResponse
	(*DeleteFoodRequest)(nil),   // 5: food.DeleteFoodRequest
}
var file_food_proto_depIdxs = []int32{
	0, // 0: food.FoodService.GetFoodById:input_type -> food.GetFoodByIdRequest
	2, // 1: food.FoodService.CreateFood:input_type -> food.CreateFoodRequest
	3, // 2: food.FoodService.UpdateFood:input_type -> food.UpdateFoodRequest
	5, // 3: food.FoodService.DeleteFood:input_type -> food.DeleteFoodRequest
	1, // 4: food.FoodService.GetFoodById:output_type -> food.GetFoodByIdResponse
	4, // 5: food.FoodService.CreateFood:output_type -> food.DefaultResponse
	4, // 6: food.FoodService.UpdateFood:output_type -> food.DefaultResponse
	4, // 7: food.FoodService.DeleteFood:output_type -> food.DefaultResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_food_proto_init() }
func file_food_proto_init() {
	if File_food_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_food_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFoodByIdRequest); i {
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
		file_food_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFoodByIdResponse); i {
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
		file_food_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateFoodRequest); i {
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
		file_food_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFoodRequest); i {
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
		file_food_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DefaultResponse); i {
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
		file_food_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFoodRequest); i {
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
			RawDescriptor: file_food_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_food_proto_goTypes,
		DependencyIndexes: file_food_proto_depIdxs,
		MessageInfos:      file_food_proto_msgTypes,
	}.Build()
	File_food_proto = out.File
	file_food_proto_rawDesc = nil
	file_food_proto_goTypes = nil
	file_food_proto_depIdxs = nil
}