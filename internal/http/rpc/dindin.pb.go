// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.17.3
// source: pb/dindin.proto

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

type EatingTomorrowRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SlackUID string `protobuf:"bytes,1,opt,name=SlackUID,proto3" json:"SlackUID,omitempty"`
}

func (x *EatingTomorrowRequest) Reset() {
	*x = EatingTomorrowRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_dindin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EatingTomorrowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EatingTomorrowRequest) ProtoMessage() {}

func (x *EatingTomorrowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_dindin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EatingTomorrowRequest.ProtoReflect.Descriptor instead.
func (*EatingTomorrowRequest) Descriptor() ([]byte, []int) {
	return file_pb_dindin_proto_rawDescGZIP(), []int{0}
}

func (x *EatingTomorrowRequest) GetSlackUID() string {
	if x != nil {
		return x.SlackUID
	}
	return ""
}

type EatingTomorrowResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EatingTomorrowResponse) Reset() {
	*x = EatingTomorrowResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_dindin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EatingTomorrowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EatingTomorrowResponse) ProtoMessage() {}

func (x *EatingTomorrowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_dindin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EatingTomorrowResponse.ProtoReflect.Descriptor instead.
func (*EatingTomorrowResponse) Descriptor() ([]byte, []int) {
	return file_pb_dindin_proto_rawDescGZIP(), []int{1}
}

type PingMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *PingMessage) Reset() {
	*x = PingMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_dindin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingMessage) ProtoMessage() {}

func (x *PingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_pb_dindin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingMessage.ProtoReflect.Descriptor instead.
func (*PingMessage) Descriptor() ([]byte, []int) {
	return file_pb_dindin_proto_rawDescGZIP(), []int{2}
}

func (x *PingMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_pb_dindin_proto protoreflect.FileDescriptor

var file_pb_dindin_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x64, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x33, 0x0a, 0x15, 0x45, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54,
	0x6f, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x49, 0x44, 0x22, 0x18, 0x0a, 0x16, 0x45, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x85, 0x01,
	0x0a, 0x0c, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x49,
	0x0a, 0x0e, 0x45, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x12, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6d, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x45, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x04, 0x50, 0x69, 0x6e,
	0x67, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x64, 0x72, 0x69, 0x74, 0x7a, 0x65, 0x6e, 0x68, 0x6f, 0x66, 0x66,
	0x2f, 0x64, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pb_dindin_proto_rawDescOnce sync.Once
	file_pb_dindin_proto_rawDescData = file_pb_dindin_proto_rawDesc
)

func file_pb_dindin_proto_rawDescGZIP() []byte {
	file_pb_dindin_proto_rawDescOnce.Do(func() {
		file_pb_dindin_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_dindin_proto_rawDescData)
	})
	return file_pb_dindin_proto_rawDescData
}

var file_pb_dindin_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_dindin_proto_goTypes = []interface{}{
	(*EatingTomorrowRequest)(nil),  // 0: pb.EatingTomorrowRequest
	(*EatingTomorrowResponse)(nil), // 1: pb.EatingTomorrowResponse
	(*PingMessage)(nil),            // 2: pb.PingMessage
}
var file_pb_dindin_proto_depIdxs = []int32{
	0, // 0: pb.SlackActions.EatingTomorrow:input_type -> pb.EatingTomorrowRequest
	2, // 1: pb.SlackActions.Ping:input_type -> pb.PingMessage
	1, // 2: pb.SlackActions.EatingTomorrow:output_type -> pb.EatingTomorrowResponse
	2, // 3: pb.SlackActions.Ping:output_type -> pb.PingMessage
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_dindin_proto_init() }
func file_pb_dindin_proto_init() {
	if File_pb_dindin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_dindin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EatingTomorrowRequest); i {
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
		file_pb_dindin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EatingTomorrowResponse); i {
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
		file_pb_dindin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingMessage); i {
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
			RawDescriptor: file_pb_dindin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_dindin_proto_goTypes,
		DependencyIndexes: file_pb_dindin_proto_depIdxs,
		MessageInfos:      file_pb_dindin_proto_msgTypes,
	}.Build()
	File_pb_dindin_proto = out.File
	file_pb_dindin_proto_rawDesc = nil
	file_pb_dindin_proto_goTypes = nil
	file_pb_dindin_proto_depIdxs = nil
}