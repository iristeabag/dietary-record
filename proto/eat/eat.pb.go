// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.12.4
// source: eat.proto

package eat

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

type Eat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eatid   string  `protobuf:"bytes,1,opt,name=Eatid,proto3" json:"Eatid,omitempty"`
	Foodid  string  `protobuf:"bytes,2,opt,name=Foodid,proto3" json:"Foodid,omitempty"`
	Name    string  `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	Amount  float32 `protobuf:"fixed32,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Unit    string  `protobuf:"bytes,5,opt,name=Unit,proto3" json:"Unit,omitempty"`
	Carb    float32 `protobuf:"fixed32,6,opt,name=Carb,proto3" json:"Carb,omitempty"`
	Portein float32 `protobuf:"fixed32,7,opt,name=Portein,proto3" json:"Portein,omitempty"`
	Fat     float32 `protobuf:"fixed32,8,opt,name=Fat,proto3" json:"Fat,omitempty"`
	Cal     float32 `protobuf:"fixed32,9,opt,name=Cal,proto3" json:"Cal,omitempty"`
}

func (x *Eat) Reset() {
	*x = Eat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Eat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Eat) ProtoMessage() {}

func (x *Eat) ProtoReflect() protoreflect.Message {
	mi := &file_eat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Eat.ProtoReflect.Descriptor instead.
func (*Eat) Descriptor() ([]byte, []int) {
	return file_eat_proto_rawDescGZIP(), []int{0}
}

func (x *Eat) GetEatid() string {
	if x != nil {
		return x.Eatid
	}
	return ""
}

func (x *Eat) GetFoodid() string {
	if x != nil {
		return x.Foodid
	}
	return ""
}

func (x *Eat) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Eat) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Eat) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *Eat) GetCarb() float32 {
	if x != nil {
		return x.Carb
	}
	return 0
}

func (x *Eat) GetPortein() float32 {
	if x != nil {
		return x.Portein
	}
	return 0
}

func (x *Eat) GetFat() float32 {
	if x != nil {
		return x.Fat
	}
	return 0
}

func (x *Eat) GetCal() float32 {
	if x != nil {
		return x.Cal
	}
	return 0
}

type GetByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
}

func (x *GetByIdRequest) Reset() {
	*x = GetByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdRequest) ProtoMessage() {}

func (x *GetByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdRequest.ProtoReflect.Descriptor instead.
func (*GetByIdRequest) Descriptor() ([]byte, []int) {
	return file_eat_proto_rawDescGZIP(), []int{1}
}

func (x *GetByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetEatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetEatsRequest) Reset() {
	*x = GetEatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEatsRequest) ProtoMessage() {}

func (x *GetEatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEatsRequest.ProtoReflect.Descriptor instead.
func (*GetEatsRequest) Descriptor() ([]byte, []int) {
	return file_eat_proto_rawDescGZIP(), []int{2}
}

type GetEatByIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eat *Eat `protobuf:"bytes,1,opt,name=eat,proto3" json:"eat,omitempty"`
}

func (x *GetEatByIdResponse) Reset() {
	*x = GetEatByIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eat_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEatByIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEatByIdResponse) ProtoMessage() {}

func (x *GetEatByIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eat_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEatByIdResponse.ProtoReflect.Descriptor instead.
func (*GetEatByIdResponse) Descriptor() ([]byte, []int) {
	return file_eat_proto_rawDescGZIP(), []int{3}
}

func (x *GetEatByIdResponse) GetEat() *Eat {
	if x != nil {
		return x.Eat
	}
	return nil
}

type GetEatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eat []*Eat `protobuf:"bytes,1,rep,name=eat,proto3" json:"eat,omitempty"`
}

func (x *GetEatsResponse) Reset() {
	*x = GetEatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eat_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEatsResponse) ProtoMessage() {}

func (x *GetEatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eat_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEatsResponse.ProtoReflect.Descriptor instead.
func (*GetEatsResponse) Descriptor() ([]byte, []int) {
	return file_eat_proto_rawDescGZIP(), []int{4}
}

func (x *GetEatsResponse) GetEat() []*Eat {
	if x != nil {
		return x.Eat
	}
	return nil
}

type EatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eat *Eat `protobuf:"bytes,1,opt,name=eat,proto3" json:"eat,omitempty"`
}

func (x *EatRequest) Reset() {
	*x = EatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_eat_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EatRequest) ProtoMessage() {}

func (x *EatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_eat_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EatRequest.ProtoReflect.Descriptor instead.
func (*EatRequest) Descriptor() ([]byte, []int) {
	return file_eat_proto_rawDescGZIP(), []int{5}
}

func (x *EatRequest) GetEat() *Eat {
	if x != nil {
		return x.Eat
	}
	return nil
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
		mi := &file_eat_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DefaultResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DefaultResponse) ProtoMessage() {}

func (x *DefaultResponse) ProtoReflect() protoreflect.Message {
	mi := &file_eat_proto_msgTypes[6]
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
	return file_eat_proto_rawDescGZIP(), []int{6}
}

func (x *DefaultResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_eat_proto protoreflect.FileDescriptor

var file_eat_proto_rawDesc = []byte{
	0x0a, 0x09, 0x65, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x65, 0x61, 0x74,
	0x22, 0xc5, 0x01, 0x0a, 0x03, 0x45, 0x61, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x61, 0x74, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x61, 0x74, 0x69, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x46, 0x6f, 0x6f, 0x64, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x46, 0x6f, 0x6f, 0x64, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x62, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x43, 0x61, 0x72, 0x62, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x6f,
	0x72, 0x74, 0x65, 0x69, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x02, 0x52, 0x07, 0x50, 0x6f, 0x72,
	0x74, 0x65, 0x69, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x46, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x46, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x61, 0x6c, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x03, 0x43, 0x61, 0x6c, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42,
	0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x45, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x30, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x45, 0x61, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x65, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x08, 0x2e, 0x65, 0x61, 0x74, 0x2e, 0x45, 0x61, 0x74, 0x52, 0x03, 0x65, 0x61, 0x74, 0x22, 0x2d,
	0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1a, 0x0a, 0x03, 0x65, 0x61, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08,
	0x2e, 0x65, 0x61, 0x74, 0x2e, 0x45, 0x61, 0x74, 0x52, 0x03, 0x65, 0x61, 0x74, 0x22, 0x28, 0x0a,
	0x0a, 0x45, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x65,
	0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x65, 0x61, 0x74, 0x2e, 0x45,
	0x61, 0x74, 0x52, 0x03, 0x65, 0x61, 0x74, 0x22, 0x29, 0x0a, 0x0f, 0x44, 0x65, 0x66, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x32, 0xa8, 0x02, 0x0a, 0x0a, 0x45, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x45, 0x61, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12,
	0x13, 0x2e, 0x65, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x65, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x61,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x36, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x45, 0x61, 0x74, 0x73, 0x12, 0x13, 0x2e, 0x65, 0x61, 0x74,
	0x2e, 0x47, 0x65, 0x74, 0x45, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x14, 0x2e, 0x65, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x45, 0x61, 0x74, 0x12, 0x0f, 0x2e, 0x65, 0x61, 0x74, 0x2e, 0x45, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x34, 0x0a,
	0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x61, 0x74, 0x12, 0x0f, 0x2e, 0x65, 0x61, 0x74,
	0x2e, 0x45, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65, 0x61,
	0x74, 0x2e, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x61, 0x74,
	0x12, 0x13, 0x2e, 0x65, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65, 0x61, 0x74, 0x2e, 0x44, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a,
	0x09, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_eat_proto_rawDescOnce sync.Once
	file_eat_proto_rawDescData = file_eat_proto_rawDesc
)

func file_eat_proto_rawDescGZIP() []byte {
	file_eat_proto_rawDescOnce.Do(func() {
		file_eat_proto_rawDescData = protoimpl.X.CompressGZIP(file_eat_proto_rawDescData)
	})
	return file_eat_proto_rawDescData
}

var file_eat_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_eat_proto_goTypes = []interface{}{
	(*Eat)(nil),                // 0: eat.Eat
	(*GetByIdRequest)(nil),     // 1: eat.GetByIdRequest
	(*GetEatsRequest)(nil),     // 2: eat.GetEatsRequest
	(*GetEatByIdResponse)(nil), // 3: eat.GetEatByIdResponse
	(*GetEatsResponse)(nil),    // 4: eat.GetEatsResponse
	(*EatRequest)(nil),         // 5: eat.EatRequest
	(*DefaultResponse)(nil),    // 6: eat.DefaultResponse
}
var file_eat_proto_depIdxs = []int32{
	0, // 0: eat.GetEatByIdResponse.eat:type_name -> eat.Eat
	0, // 1: eat.GetEatsResponse.eat:type_name -> eat.Eat
	0, // 2: eat.EatRequest.eat:type_name -> eat.Eat
	1, // 3: eat.EatService.GetEatById:input_type -> eat.GetByIdRequest
	2, // 4: eat.EatService.GetEats:input_type -> eat.GetEatsRequest
	5, // 5: eat.EatService.CreateEat:input_type -> eat.EatRequest
	5, // 6: eat.EatService.UpdateEat:input_type -> eat.EatRequest
	1, // 7: eat.EatService.DeleteEat:input_type -> eat.GetByIdRequest
	3, // 8: eat.EatService.GetEatById:output_type -> eat.GetEatByIdResponse
	4, // 9: eat.EatService.GetEats:output_type -> eat.GetEatsResponse
	6, // 10: eat.EatService.CreateEat:output_type -> eat.DefaultResponse
	6, // 11: eat.EatService.UpdateEat:output_type -> eat.DefaultResponse
	6, // 12: eat.EatService.DeleteEat:output_type -> eat.DefaultResponse
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_eat_proto_init() }
func file_eat_proto_init() {
	if File_eat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_eat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Eat); i {
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
		file_eat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdRequest); i {
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
		file_eat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEatsRequest); i {
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
		file_eat_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEatByIdResponse); i {
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
		file_eat_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEatsResponse); i {
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
		file_eat_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EatRequest); i {
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
		file_eat_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_eat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_eat_proto_goTypes,
		DependencyIndexes: file_eat_proto_depIdxs,
		MessageInfos:      file_eat_proto_msgTypes,
	}.Build()
	File_eat_proto = out.File
	file_eat_proto_rawDesc = nil
	file_eat_proto_goTypes = nil
	file_eat_proto_depIdxs = nil
}
