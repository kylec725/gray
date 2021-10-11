// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: graytorrent.proto

package rpc

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

type TorrentInfo_State int32

const (
	TorrentInfo_DOWNLOADING TorrentInfo_State = 0
	TorrentInfo_STOPPED     TorrentInfo_State = 1
	TorrentInfo_STALLED     TorrentInfo_State = 2
	TorrentInfo_SEEDING     TorrentInfo_State = 3
	TorrentInfo_COMPLETE    TorrentInfo_State = 4
)

// Enum value maps for TorrentInfo_State.
var (
	TorrentInfo_State_name = map[int32]string{
		0: "DOWNLOADING",
		1: "STOPPED",
		2: "STALLED",
		3: "SEEDING",
		4: "COMPLETE",
	}
	TorrentInfo_State_value = map[string]int32{
		"DOWNLOADING": 0,
		"STOPPED":     1,
		"STALLED":     2,
		"SEEDING":     3,
		"COMPLETE":    4,
	}
)

func (x TorrentInfo_State) Enum() *TorrentInfo_State {
	p := new(TorrentInfo_State)
	*p = x
	return p
}

func (x TorrentInfo_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TorrentInfo_State) Descriptor() protoreflect.EnumDescriptor {
	return file_graytorrent_proto_enumTypes[0].Descriptor()
}

func (TorrentInfo_State) Type() protoreflect.EnumType {
	return &file_graytorrent_proto_enumTypes[0]
}

func (x TorrentInfo_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TorrentInfo_State.Descriptor instead.
func (TorrentInfo_State) EnumDescriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{1, 0}
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_graytorrent_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{0}
}

func (x *ListRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type TorrentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	InfoHash    []byte            `protobuf:"bytes,2,opt,name=infoHash,proto3" json:"infoHash,omitempty"`
	TotalLength uint32            `protobuf:"varint,3,opt,name=totalLength,proto3" json:"totalLength,omitempty"`
	Left        uint32            `protobuf:"varint,4,opt,name=left,proto3" json:"left,omitempty"`
	DownRate    uint32            `protobuf:"varint,5,opt,name=downRate,proto3" json:"downRate,omitempty"`
	UpRate      uint32            `protobuf:"varint,6,opt,name=upRate,proto3" json:"upRate,omitempty"`
	State       TorrentInfo_State `protobuf:"varint,7,opt,name=state,proto3,enum=graytorrent.TorrentInfo_State" json:"state,omitempty"`
}

func (x *TorrentInfo) Reset() {
	*x = TorrentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TorrentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TorrentInfo) ProtoMessage() {}

func (x *TorrentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_graytorrent_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TorrentInfo.ProtoReflect.Descriptor instead.
func (*TorrentInfo) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{1}
}

func (x *TorrentInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TorrentInfo) GetInfoHash() []byte {
	if x != nil {
		return x.InfoHash
	}
	return nil
}

func (x *TorrentInfo) GetTotalLength() uint32 {
	if x != nil {
		return x.TotalLength
	}
	return 0
}

func (x *TorrentInfo) GetLeft() uint32 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *TorrentInfo) GetDownRate() uint32 {
	if x != nil {
		return x.DownRate
	}
	return 0
}

func (x *TorrentInfo) GetUpRate() uint32 {
	if x != nil {
		return x.UpRate
	}
	return 0
}

func (x *TorrentInfo) GetState() TorrentInfo_State {
	if x != nil {
		return x.State
	}
	return TorrentInfo_DOWNLOADING
}

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_graytorrent_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{2}
}

type AddReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *AddReply) Reset() {
	*x = AddReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddReply) ProtoMessage() {}

func (x *AddReply) ProtoReflect() protoreflect.Message {
	mi := &file_graytorrent_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddReply.ProtoReflect.Descriptor instead.
func (*AddReply) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{3}
}

func (x *AddReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type RemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveRequest) Reset() {
	*x = RemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveRequest) ProtoMessage() {}

func (x *RemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_graytorrent_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveRequest.ProtoReflect.Descriptor instead.
func (*RemoveRequest) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{4}
}

type RemoveReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveReply) Reset() {
	*x = RemoveReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveReply) ProtoMessage() {}

func (x *RemoveReply) ProtoReflect() protoreflect.Message {
	mi := &file_graytorrent_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveReply.ProtoReflect.Descriptor instead.
func (*RemoveReply) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{5}
}

type StartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartRequest) Reset() {
	*x = StartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartRequest) ProtoMessage() {}

func (x *StartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_graytorrent_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartRequest.ProtoReflect.Descriptor instead.
func (*StartRequest) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{6}
}

type StartReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StartReply) Reset() {
	*x = StartReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartReply) ProtoMessage() {}

func (x *StartReply) ProtoReflect() protoreflect.Message {
	mi := &file_graytorrent_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartReply.ProtoReflect.Descriptor instead.
func (*StartReply) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{7}
}

type StopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopRequest) Reset() {
	*x = StopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopRequest) ProtoMessage() {}

func (x *StopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_graytorrent_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopRequest.ProtoReflect.Descriptor instead.
func (*StopRequest) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{8}
}

type StopReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StopReply) Reset() {
	*x = StopReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StopReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StopReply) ProtoMessage() {}

func (x *StopReply) ProtoReflect() protoreflect.Message {
	mi := &file_graytorrent_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StopReply.ProtoReflect.Descriptor instead.
func (*StopReply) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{9}
}

var File_graytorrent_proto protoreflect.FileDescriptor

var file_graytorrent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x22, 0x1f, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0xac, 0x02, 0x0a, 0x0b, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x69, 0x6e, 0x66, 0x6f, 0x48, 0x61, 0x73,
	0x68, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e, 0x67, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x4c, 0x65, 0x6e,
	0x67, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x52,
	0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x52,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70, 0x52, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x06, 0x75, 0x70, 0x52, 0x61, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x67, 0x72, 0x61,
	0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x22, 0x4d, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x44, 0x4f,
	0x57, 0x4e, 0x4c, 0x4f, 0x41, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53,
	0x54, 0x4f, 0x50, 0x50, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x54, 0x41, 0x4c,
	0x4c, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x45, 0x44, 0x49, 0x4e, 0x47,
	0x10, 0x03, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04,
	0x22, 0x0c, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24,
	0x0a, 0x08, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x0d, 0x0a, 0x0b, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x0e, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x0c, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x0d, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x0b, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xbf,
	0x02, 0x0a, 0x07, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x3e, 0x0a, 0x04, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67,
	0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x6f, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x37, 0x0a, 0x03, 0x41, 0x64,
	0x64, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x61,
	0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x1a, 0x2e,
	0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x61, 0x79,
	0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x19,
	0x2e, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61,
	0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x72, 0x61, 0x79,
	0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x12, 0x18, 0x2e, 0x67,
	0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x2e, 0x53, 0x74, 0x6f, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x25, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b,
	0x79, 0x6c, 0x65, 0x63, 0x37, 0x32, 0x35, 0x2f, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_graytorrent_proto_rawDescOnce sync.Once
	file_graytorrent_proto_rawDescData = file_graytorrent_proto_rawDesc
)

func file_graytorrent_proto_rawDescGZIP() []byte {
	file_graytorrent_proto_rawDescOnce.Do(func() {
		file_graytorrent_proto_rawDescData = protoimpl.X.CompressGZIP(file_graytorrent_proto_rawDescData)
	})
	return file_graytorrent_proto_rawDescData
}

var file_graytorrent_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_graytorrent_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_graytorrent_proto_goTypes = []interface{}{
	(TorrentInfo_State)(0), // 0: graytorrent.TorrentInfo.State
	(*ListRequest)(nil),    // 1: graytorrent.ListRequest
	(*TorrentInfo)(nil),    // 2: graytorrent.TorrentInfo
	(*AddRequest)(nil),     // 3: graytorrent.AddRequest
	(*AddReply)(nil),       // 4: graytorrent.AddReply
	(*RemoveRequest)(nil),  // 5: graytorrent.RemoveRequest
	(*RemoveReply)(nil),    // 6: graytorrent.RemoveReply
	(*StartRequest)(nil),   // 7: graytorrent.StartRequest
	(*StartReply)(nil),     // 8: graytorrent.StartReply
	(*StopRequest)(nil),    // 9: graytorrent.StopRequest
	(*StopReply)(nil),      // 10: graytorrent.StopReply
}
var file_graytorrent_proto_depIdxs = []int32{
	0,  // 0: graytorrent.TorrentInfo.state:type_name -> graytorrent.TorrentInfo.State
	1,  // 1: graytorrent.Torrent.List:input_type -> graytorrent.ListRequest
	3,  // 2: graytorrent.Torrent.Add:input_type -> graytorrent.AddRequest
	5,  // 3: graytorrent.Torrent.Remove:input_type -> graytorrent.RemoveRequest
	7,  // 4: graytorrent.Torrent.Start:input_type -> graytorrent.StartRequest
	9,  // 5: graytorrent.Torrent.Stop:input_type -> graytorrent.StopRequest
	2,  // 6: graytorrent.Torrent.List:output_type -> graytorrent.TorrentInfo
	4,  // 7: graytorrent.Torrent.Add:output_type -> graytorrent.AddReply
	6,  // 8: graytorrent.Torrent.Remove:output_type -> graytorrent.RemoveReply
	8,  // 9: graytorrent.Torrent.Start:output_type -> graytorrent.StartReply
	10, // 10: graytorrent.Torrent.Stop:output_type -> graytorrent.StopReply
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_graytorrent_proto_init() }
func file_graytorrent_proto_init() {
	if File_graytorrent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graytorrent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_graytorrent_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TorrentInfo); i {
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
		file_graytorrent_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_graytorrent_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddReply); i {
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
		file_graytorrent_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveRequest); i {
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
		file_graytorrent_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveReply); i {
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
		file_graytorrent_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartRequest); i {
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
		file_graytorrent_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartReply); i {
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
		file_graytorrent_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopRequest); i {
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
		file_graytorrent_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StopReply); i {
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
			RawDescriptor: file_graytorrent_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_graytorrent_proto_goTypes,
		DependencyIndexes: file_graytorrent_proto_depIdxs,
		EnumInfos:         file_graytorrent_proto_enumTypes,
		MessageInfos:      file_graytorrent_proto_msgTypes,
	}.Build()
	File_graytorrent_proto = out.File
	file_graytorrent_proto_rawDesc = nil
	file_graytorrent_proto_goTypes = nil
	file_graytorrent_proto_depIdxs = nil
}
