// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: _proto/offer.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File___proto_offer_proto protoreflect.FileDescriptor

var file___proto_offer_proto_rawDesc = []byte{
	0x0a, 0x12, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x32, 0x0e, 0x0a, 0x0c, 0x4f,
	0x66, 0x66, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x2e,
	0x2f, 0x6f, 0x66, 0x66, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var file___proto_offer_proto_goTypes = []interface{}{}
var file___proto_offer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file___proto_offer_proto_init() }
func file___proto_offer_proto_init() {
	if File___proto_offer_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file___proto_offer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file___proto_offer_proto_goTypes,
		DependencyIndexes: file___proto_offer_proto_depIdxs,
	}.Build()
	File___proto_offer_proto = out.File
	file___proto_offer_proto_rawDesc = nil
	file___proto_offer_proto_goTypes = nil
	file___proto_offer_proto_depIdxs = nil
}