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

type ConnectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *ConnectRequest) Reset() {
	*x = ConnectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectRequest) ProtoMessage() {}

func (x *ConnectRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ConnectRequest.ProtoReflect.Descriptor instead.
func (*ConnectRequest) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{0}
}

func (x *ConnectRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ConnectReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Correct bool `protobuf:"varint,1,opt,name=correct,proto3" json:"correct,omitempty"`
}

func (x *ConnectReply) Reset() {
	*x = ConnectReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graytorrent_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectReply) ProtoMessage() {}

func (x *ConnectReply) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ConnectReply.ProtoReflect.Descriptor instead.
func (*ConnectReply) Descriptor() ([]byte, []int) {
	return file_graytorrent_proto_rawDescGZIP(), []int{1}
}

func (x *ConnectReply) GetCorrect() bool {
	if x != nil {
		return x.Correct
	}
	return false
}

var File_graytorrent_proto protoreflect.FileDescriptor

var file_graytorrent_proto_rawDesc = []byte{
	0x0a, 0x11, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x22, 0x22, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x22, 0x28, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x32, 0x4e,
	0x0a, 0x07, 0x54, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x07, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x2e,
	0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x25,
	0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x79, 0x6c,
	0x65, 0x63, 0x37, 0x32, 0x35, 0x2f, 0x67, 0x72, 0x61, 0x79, 0x74, 0x6f, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_graytorrent_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_graytorrent_proto_goTypes = []interface{}{
	(*ConnectRequest)(nil), // 0: graytorrent.ConnectRequest
	(*ConnectReply)(nil),   // 1: graytorrent.ConnectReply
}
var file_graytorrent_proto_depIdxs = []int32{
	0, // 0: graytorrent.Torrent.Connect:input_type -> graytorrent.ConnectRequest
	1, // 1: graytorrent.Torrent.Connect:output_type -> graytorrent.ConnectReply
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_graytorrent_proto_init() }
func file_graytorrent_proto_init() {
	if File_graytorrent_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graytorrent_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectRequest); i {
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
			switch v := v.(*ConnectReply); i {
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
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_graytorrent_proto_goTypes,
		DependencyIndexes: file_graytorrent_proto_depIdxs,
		MessageInfos:      file_graytorrent_proto_msgTypes,
	}.Build()
	File_graytorrent_proto = out.File
	file_graytorrent_proto_rawDesc = nil
	file_graytorrent_proto_goTypes = nil
	file_graytorrent_proto_depIdxs = nil
}
