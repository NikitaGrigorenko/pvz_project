// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.1
// source: api/devtask/pvz/v1/pvz.proto

package devtask_pvz_v1

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

type PVZInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Contact string `protobuf:"bytes,4,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *PVZInfo) Reset() {
	*x = PVZInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PVZInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PVZInfo) ProtoMessage() {}

func (x *PVZInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PVZInfo.ProtoReflect.Descriptor instead.
func (*PVZInfo) Descriptor() ([]byte, []int) {
	return file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP(), []int{0}
}

func (x *PVZInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PVZInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PVZInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PVZInfo) GetContact() string {
	if x != nil {
		return x.Contact
	}
	return ""
}

type AddInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pvz *PVZInfo `protobuf:"bytes,1,opt,name=pvz,proto3" json:"pvz,omitempty"`
}

func (x *AddInfoRequest) Reset() {
	*x = AddInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInfoRequest) ProtoMessage() {}

func (x *AddInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInfoRequest.ProtoReflect.Descriptor instead.
func (*AddInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP(), []int{1}
}

func (x *AddInfoRequest) GetPvz() *PVZInfo {
	if x != nil {
		return x.Pvz
	}
	return nil
}

type AddInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pvz *PVZInfo `protobuf:"bytes,1,opt,name=pvz,proto3" json:"pvz,omitempty"`
}

func (x *AddInfoResponse) Reset() {
	*x = AddInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddInfoResponse) ProtoMessage() {}

func (x *AddInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddInfoResponse.ProtoReflect.Descriptor instead.
func (*AddInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP(), []int{2}
}

func (x *AddInfoResponse) GetPvz() *PVZInfo {
	if x != nil {
		return x.Pvz
	}
	return nil
}

type DeleteInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteInfoRequest) Reset() {
	*x = DeleteInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInfoRequest) ProtoMessage() {}

func (x *DeleteInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInfoRequest.ProtoReflect.Descriptor instead.
func (*DeleteInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteInfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteInfoResponse) Reset() {
	*x = DeleteInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteInfoResponse) ProtoMessage() {}

func (x *DeleteInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteInfoResponse.ProtoReflect.Descriptor instead.
func (*DeleteInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP(), []int{4}
}

type GetInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetInfoRequest) Reset() {
	*x = GetInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoRequest) ProtoMessage() {}

func (x *GetInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoRequest.ProtoReflect.Descriptor instead.
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP(), []int{5}
}

func (x *GetInfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pvz *PVZInfo `protobuf:"bytes,1,opt,name=pvz,proto3" json:"pvz,omitempty"`
}

func (x *GetInfoResponse) Reset() {
	*x = GetInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInfoResponse) ProtoMessage() {}

func (x *GetInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInfoResponse.ProtoReflect.Descriptor instead.
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP(), []int{6}
}

func (x *GetInfoResponse) GetPvz() *PVZInfo {
	if x != nil {
		return x.Pvz
	}
	return nil
}

type ListInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListInfoRequest) Reset() {
	*x = ListInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInfoRequest) ProtoMessage() {}

func (x *ListInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInfoRequest.ProtoReflect.Descriptor instead.
func (*ListInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP(), []int{7}
}

type ListInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pvz []*PVZInfo `protobuf:"bytes,1,rep,name=pvz,proto3" json:"pvz,omitempty"`
}

func (x *ListInfoResponse) Reset() {
	*x = ListInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListInfoResponse) ProtoMessage() {}

func (x *ListInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListInfoResponse.ProtoReflect.Descriptor instead.
func (*ListInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP(), []int{8}
}

func (x *ListInfoResponse) GetPvz() []*PVZInfo {
	if x != nil {
		return x.Pvz
	}
	return nil
}

type UpdateInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pvz *PVZInfo `protobuf:"bytes,1,opt,name=pvz,proto3" json:"pvz,omitempty"`
}

func (x *UpdateInfoRequest) Reset() {
	*x = UpdateInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfoRequest) ProtoMessage() {}

func (x *UpdateInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfoRequest.ProtoReflect.Descriptor instead.
func (*UpdateInfoRequest) Descriptor() ([]byte, []int) {
	return file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP(), []int{9}
}

func (x *UpdateInfoRequest) GetPvz() *PVZInfo {
	if x != nil {
		return x.Pvz
	}
	return nil
}

type UpdateInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pvz *PVZInfo `protobuf:"bytes,1,opt,name=pvz,proto3" json:"pvz,omitempty"`
}

func (x *UpdateInfoResponse) Reset() {
	*x = UpdateInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateInfoResponse) ProtoMessage() {}

func (x *UpdateInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_devtask_pvz_v1_pvz_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateInfoResponse.ProtoReflect.Descriptor instead.
func (*UpdateInfoResponse) Descriptor() ([]byte, []int) {
	return file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP(), []int{10}
}

func (x *UpdateInfoResponse) GetPvz() *PVZInfo {
	if x != nil {
		return x.Pvz
	}
	return nil
}

var File_api_devtask_pvz_v1_pvz_proto protoreflect.FileDescriptor

var file_api_devtask_pvz_v1_pvz_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x76,
	0x7a, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x76, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e,
	0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e, 0x76, 0x31, 0x22, 0x61,
	0x0a, 0x07, 0x50, 0x56, 0x5a, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x22, 0x3b, 0x0a, 0x0e, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x03, 0x70, 0x76, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x56, 0x5a, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x70, 0x76, 0x7a, 0x22, 0x3c,
	0x0a, 0x0f, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x29, 0x0a, 0x03, 0x70, 0x76, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x56, 0x5a, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x70, 0x76, 0x7a, 0x22, 0x23, 0x0a, 0x11,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3c, 0x0a, 0x0f, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29, 0x0a, 0x03,
	0x70, 0x76, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x65, 0x76, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x56, 0x5a, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x03, 0x70, 0x76, 0x7a, 0x22, 0x11, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3d, 0x0a, 0x10, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x29,
	0x0a, 0x03, 0x70, 0x76, 0x7a, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x65,
	0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x56, 0x5a,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x70, 0x76, 0x7a, 0x22, 0x3e, 0x0a, 0x11, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x29,
	0x0a, 0x03, 0x70, 0x76, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x65,
	0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x56, 0x5a,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x70, 0x76, 0x7a, 0x22, 0x3f, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x29, 0x0a, 0x03, 0x70, 0x76, 0x7a, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64,
	0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x56,
	0x5a, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x70, 0x76, 0x7a, 0x32, 0xa0, 0x03, 0x0a, 0x03, 0x50,
	0x56, 0x5a, 0x12, 0x4c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x41,
	0x64, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x55, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21,
	0x2e, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1e, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1f, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x76, 0x7a,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x55, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70,
	0x76, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73,
	0x6b, 0x2e, 0x70, 0x76, 0x7a, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x5a,
	0x1d, 0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x2f, 0x70, 0x76, 0x7a, 0x2f, 0x76, 0x31, 0x3b,
	0x64, 0x65, 0x76, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x70, 0x76, 0x7a, 0x5f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_devtask_pvz_v1_pvz_proto_rawDescOnce sync.Once
	file_api_devtask_pvz_v1_pvz_proto_rawDescData = file_api_devtask_pvz_v1_pvz_proto_rawDesc
)

func file_api_devtask_pvz_v1_pvz_proto_rawDescGZIP() []byte {
	file_api_devtask_pvz_v1_pvz_proto_rawDescOnce.Do(func() {
		file_api_devtask_pvz_v1_pvz_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_devtask_pvz_v1_pvz_proto_rawDescData)
	})
	return file_api_devtask_pvz_v1_pvz_proto_rawDescData
}

var file_api_devtask_pvz_v1_pvz_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_devtask_pvz_v1_pvz_proto_goTypes = []interface{}{
	(*PVZInfo)(nil),            // 0: devtask.pvz.v1.PVZInfo
	(*AddInfoRequest)(nil),     // 1: devtask.pvz.v1.AddInfoRequest
	(*AddInfoResponse)(nil),    // 2: devtask.pvz.v1.AddInfoResponse
	(*DeleteInfoRequest)(nil),  // 3: devtask.pvz.v1.DeleteInfoRequest
	(*DeleteInfoResponse)(nil), // 4: devtask.pvz.v1.DeleteInfoResponse
	(*GetInfoRequest)(nil),     // 5: devtask.pvz.v1.GetInfoRequest
	(*GetInfoResponse)(nil),    // 6: devtask.pvz.v1.GetInfoResponse
	(*ListInfoRequest)(nil),    // 7: devtask.pvz.v1.ListInfoRequest
	(*ListInfoResponse)(nil),   // 8: devtask.pvz.v1.ListInfoResponse
	(*UpdateInfoRequest)(nil),  // 9: devtask.pvz.v1.UpdateInfoRequest
	(*UpdateInfoResponse)(nil), // 10: devtask.pvz.v1.UpdateInfoResponse
}
var file_api_devtask_pvz_v1_pvz_proto_depIdxs = []int32{
	0,  // 0: devtask.pvz.v1.AddInfoRequest.pvz:type_name -> devtask.pvz.v1.PVZInfo
	0,  // 1: devtask.pvz.v1.AddInfoResponse.pvz:type_name -> devtask.pvz.v1.PVZInfo
	0,  // 2: devtask.pvz.v1.GetInfoResponse.pvz:type_name -> devtask.pvz.v1.PVZInfo
	0,  // 3: devtask.pvz.v1.ListInfoResponse.pvz:type_name -> devtask.pvz.v1.PVZInfo
	0,  // 4: devtask.pvz.v1.UpdateInfoRequest.pvz:type_name -> devtask.pvz.v1.PVZInfo
	0,  // 5: devtask.pvz.v1.UpdateInfoResponse.pvz:type_name -> devtask.pvz.v1.PVZInfo
	1,  // 6: devtask.pvz.v1.PVZ.AddInfo:input_type -> devtask.pvz.v1.AddInfoRequest
	3,  // 7: devtask.pvz.v1.PVZ.DeleteInfo:input_type -> devtask.pvz.v1.DeleteInfoRequest
	5,  // 8: devtask.pvz.v1.PVZ.GetInfo:input_type -> devtask.pvz.v1.GetInfoRequest
	7,  // 9: devtask.pvz.v1.PVZ.ListInfo:input_type -> devtask.pvz.v1.ListInfoRequest
	9,  // 10: devtask.pvz.v1.PVZ.UpdateInfo:input_type -> devtask.pvz.v1.UpdateInfoRequest
	2,  // 11: devtask.pvz.v1.PVZ.AddInfo:output_type -> devtask.pvz.v1.AddInfoResponse
	4,  // 12: devtask.pvz.v1.PVZ.DeleteInfo:output_type -> devtask.pvz.v1.DeleteInfoResponse
	6,  // 13: devtask.pvz.v1.PVZ.GetInfo:output_type -> devtask.pvz.v1.GetInfoResponse
	8,  // 14: devtask.pvz.v1.PVZ.ListInfo:output_type -> devtask.pvz.v1.ListInfoResponse
	10, // 15: devtask.pvz.v1.PVZ.UpdateInfo:output_type -> devtask.pvz.v1.UpdateInfoResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_api_devtask_pvz_v1_pvz_proto_init() }
func file_api_devtask_pvz_v1_pvz_proto_init() {
	if File_api_devtask_pvz_v1_pvz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_devtask_pvz_v1_pvz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PVZInfo); i {
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
		file_api_devtask_pvz_v1_pvz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInfoRequest); i {
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
		file_api_devtask_pvz_v1_pvz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddInfoResponse); i {
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
		file_api_devtask_pvz_v1_pvz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInfoRequest); i {
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
		file_api_devtask_pvz_v1_pvz_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteInfoResponse); i {
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
		file_api_devtask_pvz_v1_pvz_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoRequest); i {
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
		file_api_devtask_pvz_v1_pvz_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInfoResponse); i {
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
		file_api_devtask_pvz_v1_pvz_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInfoRequest); i {
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
		file_api_devtask_pvz_v1_pvz_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListInfoResponse); i {
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
		file_api_devtask_pvz_v1_pvz_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfoRequest); i {
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
		file_api_devtask_pvz_v1_pvz_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateInfoResponse); i {
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
			RawDescriptor: file_api_devtask_pvz_v1_pvz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_devtask_pvz_v1_pvz_proto_goTypes,
		DependencyIndexes: file_api_devtask_pvz_v1_pvz_proto_depIdxs,
		MessageInfos:      file_api_devtask_pvz_v1_pvz_proto_msgTypes,
	}.Build()
	File_api_devtask_pvz_v1_pvz_proto = out.File
	file_api_devtask_pvz_v1_pvz_proto_rawDesc = nil
	file_api_devtask_pvz_v1_pvz_proto_goTypes = nil
	file_api_devtask_pvz_v1_pvz_proto_depIdxs = nil
}
