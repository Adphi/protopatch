// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.3
// source: tests/message/rename.proto

// clang-format off

package message

import (
	_ "github.com/alta/protopatch/patch/go/message"
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

type NewName struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *NewName) Reset() {
	*x = NewName{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tests_message_rename_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewName) ProtoMessage() {}

func (x *NewName) ProtoReflect() protoreflect.Message {
	mi := &file_tests_message_rename_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OldName.ProtoReflect.Descriptor instead.
func (*NewName) Descriptor() ([]byte, []int) {
	return file_tests_message_rename_proto_rawDescGZIP(), []int{0}
}

func (x *NewName) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *NewName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_tests_message_rename_proto protoreflect.FileDescriptor

var file_tests_message_rename_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x74, 0x65, 0x73, 0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2f,
	0x72, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x74, 0x65,
	0x73, 0x74, 0x73, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x1a, 0x16, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f,
	0x67, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3c, 0x0a, 0x07, 0x4f, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a,
	0x0d, 0xca, 0xb5, 0x03, 0x09, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x2a,
	0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x6c, 0x74,
	0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x70, 0x61, 0x74, 0x63, 0x68, 0x2f, 0x74, 0x65, 0x73,
	0x74, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_tests_message_rename_proto_rawDescOnce sync.Once
	file_tests_message_rename_proto_rawDescData = file_tests_message_rename_proto_rawDesc
)

func file_tests_message_rename_proto_rawDescGZIP() []byte {
	file_tests_message_rename_proto_rawDescOnce.Do(func() {
		file_tests_message_rename_proto_rawDescData = protoimpl.X.CompressGZIP(file_tests_message_rename_proto_rawDescData)
	})
	return file_tests_message_rename_proto_rawDescData
}

var file_tests_message_rename_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tests_message_rename_proto_goTypes = []interface{}{
	(*NewName)(nil), // 0: tests.enum.OldName
}
var file_tests_message_rename_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tests_message_rename_proto_init() }
func file_tests_message_rename_proto_init() {
	if File_tests_message_rename_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tests_message_rename_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewName); i {
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
			RawDescriptor: file_tests_message_rename_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tests_message_rename_proto_goTypes,
		DependencyIndexes: file_tests_message_rename_proto_depIdxs,
		MessageInfos:      file_tests_message_rename_proto_msgTypes,
	}.Build()
	File_tests_message_rename_proto = out.File
	file_tests_message_rename_proto_rawDesc = nil
	file_tests_message_rename_proto_goTypes = nil
	file_tests_message_rename_proto_depIdxs = nil
}
