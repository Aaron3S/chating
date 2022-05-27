// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: message.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type SendMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Channel string   `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{0}
}

func (x *SendMessageRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *SendMessageRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

type SendMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ReceiveMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Channel  string `protobuf:"bytes,1,opt,name=channel,proto3" json:"channel,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
}

func (x *ReceiveMessageRequest) Reset() {
	*x = ReceiveMessageRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveMessageRequest) ProtoMessage() {}

func (x *ReceiveMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveMessageRequest.ProtoReflect.Descriptor instead.
func (*ReceiveMessageRequest) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{2}
}

func (x *ReceiveMessageRequest) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *ReceiveMessageRequest) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

type ReceiveMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Messages []*Message `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	Ok       bool       `protobuf:"varint,2,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *ReceiveMessageResponse) Reset() {
	*x = ReceiveMessageResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReceiveMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReceiveMessageResponse) ProtoMessage() {}

func (x *ReceiveMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReceiveMessageResponse.ProtoReflect.Descriptor instead.
func (*ReceiveMessageResponse) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{3}
}

func (x *ReceiveMessageResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *ReceiveMessageResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type StreamMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message   *Message `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Channel   string   `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Handshake bool     `protobuf:"varint,3,opt,name=handshake,proto3" json:"handshake,omitempty"`
}

func (x *StreamMessage) Reset() {
	*x = StreamMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StreamMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StreamMessage) ProtoMessage() {}

func (x *StreamMessage) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StreamMessage.ProtoReflect.Descriptor instead.
func (*StreamMessage) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{4}
}

func (x *StreamMessage) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *StreamMessage) GetChannel() string {
	if x != nil {
		return x.Channel
	}
	return ""
}

func (x *StreamMessage) GetHandshake() bool {
	if x != nil {
		return x.Handshake
	}
	return false
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content  []byte `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Id       int32  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_message_proto_rawDescGZIP(), []int{5}
}

func (x *Message) GetContent() []byte {
	if x != nil {
		return x.Content
	}
	return nil
}

func (x *Message) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Message) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_message_proto protoreflect.FileDescriptor

var file_message_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5a, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61,
	0x6e, 0x6e, 0x65, 0x6c, 0x22, 0x25, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f,
	0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x4d, 0x0a, 0x15, 0x52,
	0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x16, 0x52, 0x65,
	0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02,
	0x6f, 0x6b, 0x22, 0x73, 0x0a, 0x0d, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x2a, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x61, 0x6e,
	0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x68, 0x61,
	0x6e, 0x64, 0x73, 0x68, 0x61, 0x6b, 0x65, 0x22, 0x4f, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x32, 0xe4, 0x01, 0x0a, 0x0e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x04, 0x53,
	0x65, 0x6e, 0x64, 0x12, 0x1b, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65,
	0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x4c, 0x0a, 0x07, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x12, 0x1e, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x52, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f,
	0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x1a, 0x16, 0x2e, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42,
	0x07, 0x5a, 0x05, 0x2e, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_message_proto_rawDescOnce sync.Once
	file_message_proto_rawDescData = file_message_proto_rawDesc
)

func file_message_proto_rawDescGZIP() []byte {
	file_message_proto_rawDescOnce.Do(func() {
		file_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_message_proto_rawDescData)
	})
	return file_message_proto_rawDescData
}

var file_message_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_message_proto_goTypes = []interface{}{
	(*SendMessageRequest)(nil),     // 0: message.SendMessageRequest
	(*SendMessageResponse)(nil),    // 1: message.SendMessageResponse
	(*ReceiveMessageRequest)(nil),  // 2: message.ReceiveMessageRequest
	(*ReceiveMessageResponse)(nil), // 3: message.ReceiveMessageResponse
	(*StreamMessage)(nil),          // 4: message.StreamMessage
	(*Message)(nil),                // 5: message.Message
}
var file_message_proto_depIdxs = []int32{
	5, // 0: message.SendMessageRequest.message:type_name -> message.Message
	5, // 1: message.ReceiveMessageResponse.messages:type_name -> message.Message
	5, // 2: message.StreamMessage.message:type_name -> message.Message
	0, // 3: message.MessageService.Send:input_type -> message.SendMessageRequest
	2, // 4: message.MessageService.Receive:input_type -> message.ReceiveMessageRequest
	4, // 5: message.MessageService.Connect:input_type -> message.StreamMessage
	1, // 6: message.MessageService.Send:output_type -> message.SendMessageResponse
	3, // 7: message.MessageService.Receive:output_type -> message.ReceiveMessageResponse
	4, // 8: message.MessageService.Connect:output_type -> message.StreamMessage
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_message_proto_init() }
func file_message_proto_init() {
	if File_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageRequest); i {
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
		file_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMessageResponse); i {
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
		file_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveMessageRequest); i {
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
		file_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReceiveMessageResponse); i {
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
		file_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StreamMessage); i {
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
		file_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
			RawDescriptor: file_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_message_proto_goTypes,
		DependencyIndexes: file_message_proto_depIdxs,
		MessageInfos:      file_message_proto_msgTypes,
	}.Build()
	File_message_proto = out.File
	file_message_proto_rawDesc = nil
	file_message_proto_goTypes = nil
	file_message_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	Send(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error)
	Receive(ctx context.Context, in *ReceiveMessageRequest, opts ...grpc.CallOption) (*ReceiveMessageResponse, error)
	Connect(ctx context.Context, opts ...grpc.CallOption) (MessageService_ConnectClient, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) Send(ctx context.Context, in *SendMessageRequest, opts ...grpc.CallOption) (*SendMessageResponse, error) {
	out := new(SendMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) Receive(ctx context.Context, in *ReceiveMessageRequest, opts ...grpc.CallOption) (*ReceiveMessageResponse, error) {
	out := new(ReceiveMessageResponse)
	err := c.cc.Invoke(ctx, "/message.MessageService/Receive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *messageServiceClient) Connect(ctx context.Context, opts ...grpc.CallOption) (MessageService_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageService_serviceDesc.Streams[0], "/message.MessageService/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceConnectClient{stream}
	return x, nil
}

type MessageService_ConnectClient interface {
	Send(*StreamMessage) error
	Recv() (*StreamMessage, error)
	grpc.ClientStream
}

type messageServiceConnectClient struct {
	grpc.ClientStream
}

func (x *messageServiceConnectClient) Send(m *StreamMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceConnectClient) Recv() (*StreamMessage, error) {
	m := new(StreamMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	Send(context.Context, *SendMessageRequest) (*SendMessageResponse, error)
	Receive(context.Context, *ReceiveMessageRequest) (*ReceiveMessageResponse, error)
	Connect(MessageService_ConnectServer) error
}

// UnimplementedMessageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (*UnimplementedMessageServiceServer) Send(context.Context, *SendMessageRequest) (*SendMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (*UnimplementedMessageServiceServer) Receive(context.Context, *ReceiveMessageRequest) (*ReceiveMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Receive not implemented")
}
func (*UnimplementedMessageServiceServer) Connect(MessageService_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Send(ctx, req.(*SendMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_Receive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReceiveMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServiceServer).Receive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.MessageService/Receive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServiceServer).Receive(ctx, req.(*ReceiveMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MessageService_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).Connect(&messageServiceConnectServer{stream})
}

type MessageService_ConnectServer interface {
	Send(*StreamMessage) error
	Recv() (*StreamMessage, error)
	grpc.ServerStream
}

type messageServiceConnectServer struct {
	grpc.ServerStream
}

func (x *messageServiceConnectServer) Send(m *StreamMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceConnectServer) Recv() (*StreamMessage, error) {
	m := new(StreamMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _MessageService_Send_Handler,
		},
		{
			MethodName: "Receive",
			Handler:    _MessageService_Receive_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _MessageService_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}
