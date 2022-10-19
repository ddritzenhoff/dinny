// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: internal/http/rpc/dindin.proto

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
		mi := &file_internal_http_rpc_dindin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EatingTomorrowRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EatingTomorrowRequest) ProtoMessage() {}

func (x *EatingTomorrowRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_http_rpc_dindin_proto_msgTypes[0]
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
	return file_internal_http_rpc_dindin_proto_rawDescGZIP(), []int{0}
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
		mi := &file_internal_http_rpc_dindin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EatingTomorrowResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EatingTomorrowResponse) ProtoMessage() {}

func (x *EatingTomorrowResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_http_rpc_dindin_proto_msgTypes[1]
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
	return file_internal_http_rpc_dindin_proto_rawDescGZIP(), []int{1}
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
		mi := &file_internal_http_rpc_dindin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingMessage) ProtoMessage() {}

func (x *PingMessage) ProtoReflect() protoreflect.Message {
	mi := &file_internal_http_rpc_dindin_proto_msgTypes[2]
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
	return file_internal_http_rpc_dindin_proto_rawDescGZIP(), []int{2}
}

func (x *PingMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetMembersRequest) Reset() {
	*x = GetMembersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_http_rpc_dindin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMembersRequest) ProtoMessage() {}

func (x *GetMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_http_rpc_dindin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMembersRequest.ProtoReflect.Descriptor instead.
func (*GetMembersRequest) Descriptor() ([]byte, []int) {
	return file_internal_http_rpc_dindin_proto_rawDescGZIP(), []int{3}
}

type GetMembersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName   string `protobuf:"bytes,1,opt,name=firstName,proto3" json:"firstName,omitempty"`
	LastName    string `protobuf:"bytes,2,opt,name=lastName,proto3" json:"lastName,omitempty"`
	RealName    string `protobuf:"bytes,3,opt,name=realName,proto3" json:"realName,omitempty"`
	DisplayName string `protobuf:"bytes,4,opt,name=displayName,proto3" json:"displayName,omitempty"`
	SlackUID    string `protobuf:"bytes,5,opt,name=slackUID,proto3" json:"slackUID,omitempty"`
}

func (x *GetMembersResponse) Reset() {
	*x = GetMembersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_http_rpc_dindin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMembersResponse) ProtoMessage() {}

func (x *GetMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_http_rpc_dindin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMembersResponse.ProtoReflect.Descriptor instead.
func (*GetMembersResponse) Descriptor() ([]byte, []int) {
	return file_internal_http_rpc_dindin_proto_rawDescGZIP(), []int{4}
}

func (x *GetMembersResponse) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *GetMembersResponse) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *GetMembersResponse) GetRealName() string {
	if x != nil {
		return x.RealName
	}
	return ""
}

func (x *GetMembersResponse) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *GetMembersResponse) GetSlackUID() string {
	if x != nil {
		return x.SlackUID
	}
	return ""
}

type WeeklyUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WeeklyUpdateRequest) Reset() {
	*x = WeeklyUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_http_rpc_dindin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeeklyUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeeklyUpdateRequest) ProtoMessage() {}

func (x *WeeklyUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_http_rpc_dindin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeeklyUpdateRequest.ProtoReflect.Descriptor instead.
func (*WeeklyUpdateRequest) Descriptor() ([]byte, []int) {
	return file_internal_http_rpc_dindin_proto_rawDescGZIP(), []int{5}
}

type WeeklyUpdateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WeeklyUpdateResponse) Reset() {
	*x = WeeklyUpdateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_http_rpc_dindin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WeeklyUpdateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeeklyUpdateResponse) ProtoMessage() {}

func (x *WeeklyUpdateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_http_rpc_dindin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeeklyUpdateResponse.ProtoReflect.Descriptor instead.
func (*WeeklyUpdateResponse) Descriptor() ([]byte, []int) {
	return file_internal_http_rpc_dindin_proto_rawDescGZIP(), []int{6}
}

var File_internal_http_rpc_dindin_proto protoreflect.FileDescriptor

var file_internal_http_rpc_dindin_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x64, 0x69, 0x6e, 0x64, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x72, 0x70, 0x63, 0x22, 0x33, 0x0a, 0x15, 0x45, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54,
	0x6f, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x08, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x49, 0x44, 0x22, 0x18, 0x0a, 0x16, 0x45, 0x61,
	0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x27, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x13, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xa8, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69,
	0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x49, 0x44, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6c, 0x61, 0x63, 0x6b, 0x55, 0x49, 0x44, 0x22, 0x15, 0x0a,
	0x13, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x16, 0x0a, 0x14, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x93, 0x02, 0x0a,
	0x0c, 0x53, 0x6c, 0x61, 0x63, 0x6b, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x4b, 0x0a,
	0x0e, 0x45, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x77, 0x12,
	0x1a, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6d, 0x6f,
	0x72, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6d, 0x6f, 0x72, 0x72, 0x6f, 0x77,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2c, 0x0a, 0x04, 0x50, 0x69,
	0x6e, 0x67, 0x12, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x10, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01, 0x12, 0x45, 0x0a, 0x0c, 0x57,
	0x65, 0x65, 0x6b, 0x6c, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x18, 0x2e, 0x72, 0x70,
	0x63, 0x2e, 0x57, 0x65, 0x65, 0x6b, 0x6c, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x72, 0x70, 0x63, 0x2e, 0x57, 0x65, 0x65, 0x6b,
	0x6c, 0x79, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x64, 0x72, 0x69, 0x74, 0x7a, 0x65, 0x6e, 0x68, 0x6f, 0x66, 0x66, 0x2f, 0x64, 0x69,
	0x6e, 0x64, 0x69, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x68, 0x74,
	0x74, 0x70, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_http_rpc_dindin_proto_rawDescOnce sync.Once
	file_internal_http_rpc_dindin_proto_rawDescData = file_internal_http_rpc_dindin_proto_rawDesc
)

func file_internal_http_rpc_dindin_proto_rawDescGZIP() []byte {
	file_internal_http_rpc_dindin_proto_rawDescOnce.Do(func() {
		file_internal_http_rpc_dindin_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_http_rpc_dindin_proto_rawDescData)
	})
	return file_internal_http_rpc_dindin_proto_rawDescData
}

var file_internal_http_rpc_dindin_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_internal_http_rpc_dindin_proto_goTypes = []interface{}{
	(*EatingTomorrowRequest)(nil),  // 0: rpc.EatingTomorrowRequest
	(*EatingTomorrowResponse)(nil), // 1: rpc.EatingTomorrowResponse
	(*PingMessage)(nil),            // 2: rpc.PingMessage
	(*GetMembersRequest)(nil),      // 3: rpc.GetMembersRequest
	(*GetMembersResponse)(nil),     // 4: rpc.GetMembersResponse
	(*WeeklyUpdateRequest)(nil),    // 5: rpc.WeeklyUpdateRequest
	(*WeeklyUpdateResponse)(nil),   // 6: rpc.WeeklyUpdateResponse
}
var file_internal_http_rpc_dindin_proto_depIdxs = []int32{
	0, // 0: rpc.SlackActions.EatingTomorrow:input_type -> rpc.EatingTomorrowRequest
	2, // 1: rpc.SlackActions.Ping:input_type -> rpc.PingMessage
	3, // 2: rpc.SlackActions.GetMembers:input_type -> rpc.GetMembersRequest
	5, // 3: rpc.SlackActions.WeeklyUpdate:input_type -> rpc.WeeklyUpdateRequest
	1, // 4: rpc.SlackActions.EatingTomorrow:output_type -> rpc.EatingTomorrowResponse
	2, // 5: rpc.SlackActions.Ping:output_type -> rpc.PingMessage
	4, // 6: rpc.SlackActions.GetMembers:output_type -> rpc.GetMembersResponse
	6, // 7: rpc.SlackActions.WeeklyUpdate:output_type -> rpc.WeeklyUpdateResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_http_rpc_dindin_proto_init() }
func file_internal_http_rpc_dindin_proto_init() {
	if File_internal_http_rpc_dindin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_http_rpc_dindin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_http_rpc_dindin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_http_rpc_dindin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_http_rpc_dindin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMembersRequest); i {
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
		file_internal_http_rpc_dindin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMembersResponse); i {
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
		file_internal_http_rpc_dindin_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeeklyUpdateRequest); i {
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
		file_internal_http_rpc_dindin_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WeeklyUpdateResponse); i {
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
			RawDescriptor: file_internal_http_rpc_dindin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_http_rpc_dindin_proto_goTypes,
		DependencyIndexes: file_internal_http_rpc_dindin_proto_depIdxs,
		MessageInfos:      file_internal_http_rpc_dindin_proto_msgTypes,
	}.Build()
	File_internal_http_rpc_dindin_proto = out.File
	file_internal_http_rpc_dindin_proto_rawDesc = nil
	file_internal_http_rpc_dindin_proto_goTypes = nil
	file_internal_http_rpc_dindin_proto_depIdxs = nil
}
