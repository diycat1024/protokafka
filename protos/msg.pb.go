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

type MsgID int32

const (
	MsgID_eMsgToLSFromGC_Unknow   MsgID = 0
	MsgID_eMsgToLSFromGC_Begin    MsgID = 40960
	MsgID_eMsgToLSFromGC_AskLogin MsgID = 40961
	MsgID_eMsgToLSFromGC_End      MsgID = 40970
)

// Enum value maps for MsgID.
var (
	MsgID_name = map[int32]string{
		0:     "eMsgToLSFromGC_Unknow",
		40960: "eMsgToLSFromGC_Begin",
		40961: "eMsgToLSFromGC_AskLogin",
		40970: "eMsgToLSFromGC_End",
	}
	MsgID_value = map[string]int32{
		"eMsgToLSFromGC_Unknow":   0,
		"eMsgToLSFromGC_Begin":    40960,
		"eMsgToLSFromGC_AskLogin": 40961,
		"eMsgToLSFromGC_End":      40970,
	}
)

func (x MsgID) Enum() *MsgID {
	p := new(MsgID)
	*p = x
	return p
}

func (x MsgID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MsgID) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (MsgID) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x MsgID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MsgID.Descriptor instead.
func (MsgID) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

type AskLogin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msgid     MsgID  `protobuf:"varint,1,opt,name=msgid,proto3,enum=GCToLS.MsgID" json:"msgid,omitempty"`
	Platform  uint32 `protobuf:"varint,2,opt,name=platform,proto3" json:"platform,omitempty"`
	Uin       string `protobuf:"bytes,3,opt,name=uin,proto3" json:"uin,omitempty"`
	Sessionid string `protobuf:"bytes,4,opt,name=sessionid,proto3" json:"sessionid,omitempty"`
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

func (x *AskLogin) GetMsgid() MsgID {
	if x != nil {
		return x.Msgid
	}
	return MsgID_eMsgToLSFromGC_Unknow
}

func (x *AskLogin) GetPlatform() uint32 {
	if x != nil {
		return x.Platform
	}
	return 0
}

func (x *AskLogin) GetUin() string {
	if x != nil {
		return x.Uin
	}
	return ""
}

func (x *AskLogin) GetSessionid() string {
	if x != nil {
		return x.Sessionid
	}
	return ""
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x47, 0x43, 0x54,
	0x6f, 0x4c, 0x53, 0x22, 0x7b, 0x0a, 0x08, 0x41, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x23, 0x0a, 0x05, 0x6d, 0x73, 0x67, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d,
	0x2e, 0x47, 0x43, 0x54, 0x6f, 0x4c, 0x53, 0x2e, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x52, 0x05, 0x6d,
	0x73, 0x67, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x74, 0x66, 0x6f, 0x72, 0x6d,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75,
	0x69, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x64,
	0x2a, 0x77, 0x0a, 0x05, 0x4d, 0x73, 0x67, 0x49, 0x44, 0x12, 0x19, 0x0a, 0x15, 0x65, 0x4d, 0x73,
	0x67, 0x54, 0x6f, 0x4c, 0x53, 0x46, 0x72, 0x6f, 0x6d, 0x47, 0x43, 0x5f, 0x55, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x10, 0x00, 0x12, 0x1a, 0x0a, 0x14, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x4c, 0x53,
	0x46, 0x72, 0x6f, 0x6d, 0x47, 0x43, 0x5f, 0x42, 0x65, 0x67, 0x69, 0x6e, 0x10, 0x80, 0xc0, 0x02,
	0x12, 0x1d, 0x0a, 0x17, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x4c, 0x53, 0x46, 0x72, 0x6f, 0x6d,
	0x47, 0x43, 0x5f, 0x41, 0x73, 0x6b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x10, 0x81, 0xc0, 0x02, 0x12,
	0x18, 0x0a, 0x12, 0x65, 0x4d, 0x73, 0x67, 0x54, 0x6f, 0x4c, 0x53, 0x46, 0x72, 0x6f, 0x6d, 0x47,
	0x43, 0x5f, 0x45, 0x6e, 0x64, 0x10, 0x8a, 0xc0, 0x02, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x3b, 0x47,
	0x43, 0x54, 0x6f, 0x4c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_msg_proto_goTypes = []interface{}{
	(MsgID)(0),       // 0: GCToLS.MsgID
	(*AskLogin)(nil), // 1: GCToLS.AskLogin
}
var file_msg_proto_depIdxs = []int32{
	0, // 0: GCToLS.AskLogin.msgid:type_name -> GCToLS.MsgID
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
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
