// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: sqlc/v1/options.proto

package sqlcv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_sqlc_v1_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         80000,
		Name:          "sqlc.v1.sqlc_entity",
		Tag:           "bytes,80000,opt,name=sqlc_entity",
		Filename:      "sqlc/v1/options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.MessageOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         60001,
		Name:          "sqlc.v1.sqlc_request",
		Tag:           "varint,60001,opt,name=sqlc_request",
		Filename:      "sqlc/v1/options.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         80002,
		Name:          "sqlc.v1.sqlc_fk",
		Tag:           "bytes,80002,opt,name=sqlc_fk",
		Filename:      "sqlc/v1/options.proto",
	},
}

// Extension fields to descriptorpb.MessageOptions.
var (
	// optional string sqlc_entity = 80000;
	E_SqlcEntity = &file_sqlc_v1_options_proto_extTypes[0]
	// optional bool sqlc_request = 60001;
	E_SqlcRequest = &file_sqlc_v1_options_proto_extTypes[1]
)

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional string sqlc_fk = 80002;
	E_SqlcFk = &file_sqlc_v1_options_proto_extTypes[2]
)

var File_sqlc_v1_options_proto protoreflect.FileDescriptor

const file_sqlc_v1_options_proto_rawDesc = "" +
	"\n" +
	"\x15sqlc/v1/options.proto\x12\asqlc.v1\x1a google/protobuf/descriptor.proto:B\n" +
	"\vsqlc_entity\x12\x1f.google.protobuf.MessageOptions\x18\x80\xf1\x04 \x01(\tR\n" +
	"sqlcEntity:D\n" +
	"\fsqlc_request\x12\x1f.google.protobuf.MessageOptions\x18\xe1\xd4\x03 \x01(\bR\vsqlcRequest:8\n" +
	"\asqlc_fk\x12\x1d.google.protobuf.FieldOptions\x18\x82\xf1\x04 \x01(\tR\x06sqlcFkB\x8f\x01\n" +
	"\vcom.sqlc.v1B\fOptionsProtoP\x01Z5github.com/viqueen/protoc-gen-sqlc/api/sqlc/v1;sqlcv1\xa2\x02\x03SXX\xaa\x02\aSqlc.V1\xca\x02\aSqlc\\V1\xe2\x02\x13Sqlc\\V1\\GPBMetadata\xea\x02\bSqlc::V1b\x06proto3"

var file_sqlc_v1_options_proto_goTypes = []any{
	(*descriptorpb.MessageOptions)(nil), // 0: google.protobuf.MessageOptions
	(*descriptorpb.FieldOptions)(nil),   // 1: google.protobuf.FieldOptions
}
var file_sqlc_v1_options_proto_depIdxs = []int32{
	0, // 0: sqlc.v1.sqlc_entity:extendee -> google.protobuf.MessageOptions
	0, // 1: sqlc.v1.sqlc_request:extendee -> google.protobuf.MessageOptions
	1, // 2: sqlc.v1.sqlc_fk:extendee -> google.protobuf.FieldOptions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	0, // [0:3] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_sqlc_v1_options_proto_init() }
func file_sqlc_v1_options_proto_init() {
	if File_sqlc_v1_options_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_sqlc_v1_options_proto_rawDesc), len(file_sqlc_v1_options_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 3,
			NumServices:   0,
		},
		GoTypes:           file_sqlc_v1_options_proto_goTypes,
		DependencyIndexes: file_sqlc_v1_options_proto_depIdxs,
		ExtensionInfos:    file_sqlc_v1_options_proto_extTypes,
	}.Build()
	File_sqlc_v1_options_proto = out.File
	file_sqlc_v1_options_proto_goTypes = nil
	file_sqlc_v1_options_proto_depIdxs = nil
}
