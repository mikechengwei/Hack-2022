// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: tools.proto

package msg

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

var File_tools_proto protoreflect.FileDescriptor

var file_tools_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x6f, 0x6f, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d,
	0x73, 0x67, 0x1a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x86, 0x01,
	0x0a, 0x0c, 0x54, 0x6f, 0x6f, 0x6c, 0x73, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x12, 0x32,
	0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x52, 0x65, 0x71, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x12, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72,
	0x22, 0x00, 0x12, 0x42, 0x0a, 0x16, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x6d,
	0x73, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x1a, 0x10, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x22, 0x00, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2e, 0x2f, 0x6d, 0x73, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_tools_proto_goTypes = []interface{}{
	(*ReqRegister)(nil),     // 0: msg.ReqRegister
	(*ReportTaskState)(nil), // 1: msg.ReportTaskState
	(*ReplyRegister)(nil),   // 2: msg.ReplyRegister
	(*ReplyReport)(nil),     // 3: msg.ReplyReport
}
var file_tools_proto_depIdxs = []int32{
	0, // 0: msg.ToolsManager.Register:input_type -> msg.ReqRegister
	1, // 1: msg.ToolsManager.ReportTaskCurrentState:input_type -> msg.ReportTaskState
	2, // 2: msg.ToolsManager.Register:output_type -> msg.ReplyRegister
	3, // 3: msg.ToolsManager.ReportTaskCurrentState:output_type -> msg.ReplyReport
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tools_proto_init() }
func file_tools_proto_init() {
	if File_tools_proto != nil {
		return
	}
	file_msg_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tools_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tools_proto_goTypes,
		DependencyIndexes: file_tools_proto_depIdxs,
	}.Build()
	File_tools_proto = out.File
	file_tools_proto_rawDesc = nil
	file_tools_proto_goTypes = nil
	file_tools_proto_depIdxs = nil
}
