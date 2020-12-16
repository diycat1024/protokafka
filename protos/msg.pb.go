// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.11.4
// source: msg.proto

package GCToLs

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Cmd int32

const (
	Cmd_eMsgToLSFromGC_Unknow Cmd = 0
	//登录相关
	Cmd_eMsgToLSFromGC_AskLogin Cmd = 1
	// 消息相关
	Cmd_eMsgToLSFromGC_SendMsg Cmd = 160
	Cmd_eMsgToLSFromGC_ReadMsg Cmd = 161
	//消息通知相关
	Cmd_eMsgToGCFromLS_NetMsg Cmd = 208
)

// Enum value maps for Cmd.
var (
	Cmd_name = map[int32]string{
		0:   "eMsgToLSFromGC_Unknow",
		1:   "eMsgToLSFromGC_AskLogin",
		160: "eMsgToLSFromGC_SendMsg",
		161: "eMsgToLSFromGC_ReadMsg",
		208: "eMsgToGCFromLS_NetMsg",
	}
	Cmd_value = map[string]int32{
		"eMsgToLSFromGC_Unknow":   0,
		"eMsgToLSFromGC_AskLogin": 1,
		"eMsgToLSFromGC_SendMsg":  160,
		"eMsgToLSFromGC_ReadMsg":  161,
		"eMsgToGCFromLS_NetMsg":   208,
	}
)

func (x Cmd) Enum() *Cmd {
	p := new(Cmd)
	*p = x
	return p
}

func (x Cmd) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Cmd) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (Cmd) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x Cmd) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Cmd.Descriptor instead.
func (Cmd) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

type AskLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  Cmd    `protobuf:"varint,1,opt,name=cmd,proto3,enum=GCToLS.Cmd" json:"cmd,omitempty"`
	Plat string `protobuf:"bytes,2,opt,name=plat,proto3" json:"plat,omitempty"`
	Uid  string `protobuf:"bytes,4,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *AskLogin) Reset() {
	*x = AskLogin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AskLogin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AskLogin) ProtoMessage() {}

func (x *AskLogin) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AskLogin.ProtoReflect.Descriptor instead.
func (*AskLogin) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *AskLogin) GetCmd() Cmd {
	if x != nil {
		return x.Cmd
	}
	return Cmd_eMsgToLSFromGC_Unknow
}

func (x *AskLogin) GetPlat() string {
	if x != nil {
		return x.Plat
	}
	return ""
}

func (x *AskLogin) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type SendMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  Cmd    `protobuf:"varint,1,opt,name=cmd,proto3,enum=GCToLS.Cmd" json:"cmd,omitempty"`
	Send string `protobuf:"bytes,2,opt,name=send,proto3" json:"send,omitempty"`
	Recv string `protobuf:"bytes,3,opt,name=recv,proto3" json:"recv,omitempty"`
	Plat string `protobuf:"bytes,4,opt,name=plat,proto3" json:"plat,omitempty"`
	Data string `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SendMsg) Reset() {
	*x = SendMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMsg) ProtoMessage() {}

func (x *SendMsg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMsg.ProtoReflect.Descriptor instead.
func (*SendMsg) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *SendMsg) GetCmd() Cmd {
	if x != nil {
		return x.Cmd
	}
	return Cmd_eMsgToLSFromGC_Unknow
}

func (x *SendMsg) GetSend() string {
	if x != nil {
		return x.Send
	}
	return ""
}

func (x *SendMsg) GetRecv() string {
	if x != nil {
		return x.Recv
	}
	return ""
}

func (x *SendMsg) GetPlat() string {
	if x != nil {
		return x.Plat
	}
	return ""
}

func (x *SendMsg) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

type NetMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  Cmd    `protobuf:"varint,1,opt,name=cmd,proto3,enum=GCToLS.Cmd" json:"cmd,omitempty"`
	Send string `protobuf:"bytes,2,opt,name=send,proto3" json:"send,omitempty"`
	Recv string `protobuf:"bytes,3,opt,name=recv,proto3" json:"recv,omitempty"`
	Plat string `protobuf:"bytes,4,opt,name=plat,proto3" json:"plat,omitempty"`
	Mid  string `protobuf:"bytes,5,opt,name=mid,proto3" json:"mid,omitempty"`
}

func (x *NetMsg) Reset() {
	*x = NetMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetMsg) ProtoMessage() {}

func (x *NetMsg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetMsg.ProtoReflect.Descriptor instead.
func (*NetMsg) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{2}
}

func (x *NetMsg) GetCmd() Cmd {
	if x != nil {
		return x.Cmd
	}
	return Cmd_eMsgToLSFromGC_Unknow
}

func (x *NetMsg) GetSend() string {
	if x != nil {
		return x.Send
	}
	return ""
}

func (x *NetMsg) GetRecv() string {
	if x != nil {
		return x.Recv
	}
	return ""
}

func (x *NetMsg) GetPlat() string {
	if x != nil {
		return x.Plat
	}
	return ""
}

func (x *NetMsg) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

type ReadMsg struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd  Cmd    `protobuf:"varint,1,opt,name=cmd,proto3,enum=GCToLS.Cmd" json:"cmd,omitempty"`
	Send string `protobuf:"bytes,2,opt,name=send,proto3" json:"send,omitempty"`
	Recv string `protobuf:"bytes,3,opt,name=recv,proto3" json:"recv,omitempty"`
	Mid  string `protobuf:"bytes,4,opt,name=mid,proto3" json:"mid,omitempty"`
}

func (x *ReadMsg) Reset() {
	*x = ReadMsg{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadMsg) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadMsg) ProtoMessage() {}

func (x *ReadMsg) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadMsg.ProtoReflect.Descriptor instead.
func (*ReadMsg) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{3}
}

func (x *ReadMsg) GetCmd() Cmd {
	if x != nil {
		return x.Cmd
	}
	return Cmd_eMsgToLSFromGC_Unknow
}

func (x *ReadMsg) GetSend() string {
	if x != nil {
		return x.Send
	}
	return ""
}

func (x *ReadMsg) GetRecv() string {
	if x != nil {
		return x.Recv
	}
	return ""
}

func (x *ReadMsg) GetMid() string {
	if x != nil {
		return x.Mid
	}
	return ""
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x47, 0x43, 0x54,
	0x6f, 0x4c, 0x53, 0x22, 0x4f, 0x0a, 0x08, 0x41, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x1d, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x47,
	0x43, 0x54, 0x6f, 0x4c, 0x53, 0x2e, 0x43, 0x6d, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6c,
	0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x69, 0x64, 0x22, 0x78, 0x0a, 0x07, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x12,
	0x1d, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x47,
	0x43, 0x54, 0x6f, 0x4c, 0x53, 0x2e, 0x43, 0x6d, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x65,
	0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x63, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x65, 0x63, 0x76, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x75,
	0x0a, 0x06, 0x4e, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x1d, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e, 0x47, 0x43, 0x54, 0x6f, 0x4c, 0x53, 0x2e, 0x43,
	0x6d, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72,
	0x65, 0x63, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x72, 0x65, 0x63, 0x76, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x6c, 0x61, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x69, 0x64, 0x22, 0x62, 0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x73, 0x67,
	0x12, 0x1d, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0b, 0x2e,
	0x47, 0x43, 0x54, 0x6f, 0x4c, 0x53, 0x2e, 0x43, 0x6d, 0x64, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x65, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x65, 0x6e, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x65, 0x63, 0x76, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x72, 0x65, 0x63, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x69, 0x64, 0x2a, 0x93, 0x01, 0x0a, 0x03, 0x43, 0x6d,
	0x64, 0x12, 0x19, 0x0a, 0x15, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x4c, 0x53, 0x46, 0x72, 0x6f,
	0x6d, 0x47, 0x43, 0x5f, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17,
	0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x4c, 0x53, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x43, 0x5f, 0x41,
	0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x16, 0x65, 0x4d, 0x73,
	0x67, 0x54, 0x6f, 0x4c, 0x53, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x43, 0x5f, 0x53, 0x65, 0x6e, 0x64,
	0x4d, 0x73, 0x67, 0x10, 0xa0, 0x01, 0x12, 0x1b, 0x0a, 0x16, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f,
	0x4c, 0x53, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x43, 0x5f, 0x52, 0x65, 0x61, 0x64, 0x4d, 0x73, 0x67,
	0x10, 0xa1, 0x01, 0x12, 0x1a, 0x0a, 0x15, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x47, 0x43, 0x46,
	0x72, 0x6f, 0x6d, 0x4c, 0x53, 0x5f, 0x4e, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x10, 0xd0, 0x01, 0x42,
	0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x47, 0x43, 0x54, 0x6f, 0x4c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_msg_proto_rawDescOnce sync.Once
	file_msg_proto_rawDescData = file_msg_proto_rawDesc
)

func file_msg_proto_rawDescGZIP() []byte {
	file_msg_proto_rawDescOnce.Do(func() {
		file_msg_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_proto_rawDescData)
	})
	return file_msg_proto_rawDescData
}

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_msg_proto_goTypes = []interface{}{
	(Cmd)(0),         // 0: GCToLS.Cmd
	(*AskLogin)(nil), // 1: GCToLS.AskLogin
	(*SendMsg)(nil),  // 2: GCToLS.SendMsg
	(*NetMsg)(nil),   // 3: GCToLS.NetMsg
	(*ReadMsg)(nil),  // 4: GCToLS.ReadMsg
}
var file_msg_proto_depIdxs = []int32{
	0, // 0: GCToLS.AskLogin.cmd:type_name -> GCToLS.Cmd
	0, // 1: GCToLS.SendMsg.cmd:type_name -> GCToLS.Cmd
	0, // 2: GCToLS.NetMsg.cmd:type_name -> GCToLS.Cmd
	0, // 3: GCToLS.ReadMsg.cmd:type_name -> GCToLS.Cmd
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AskLogin); i {
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
		file_msg_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendMsg); i {
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
		file_msg_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetMsg); i {
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
		file_msg_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadMsg); i {
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
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_msg_proto_goTypes,
		DependencyIndexes: file_msg_proto_depIdxs,
		EnumInfos:         file_msg_proto_enumTypes,
		MessageInfos:      file_msg_proto_msgTypes,
	}.Build()
	File_msg_proto = out.File
	file_msg_proto_rawDesc = nil
	file_msg_proto_goTypes = nil
	file_msg_proto_depIdxs = nil
}
