// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: api/helloworld/v1/classservice.proto

package v1

import (
	annotations "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of google/api/annotations.proto.

var E_Http = annotations.E_Http

var File_api_helloworld_v1_classservice_proto protoreflect.FileDescriptor

const file_api_helloworld_v1_classservice_proto_rawDesc = "" +
	"\n" +
	"$api/helloworld/v1/classservice.proto\x12\rhelloworld.v1\x1a\x1cgoogle/api/annotations.proto\x1a\x1dapi/helloworld/v1/class.proto2\xda\a\n" +
	"\fClassService\x12j\n" +
	"\vCreateClass\x12!.helloworld.v1.CreateClassRequest\x1a\x1f.helloworld.v1.CreateClassReply\"\x17\x82\xd3\xe4\x93\x02\x11:\x01*\"\fapi/v1/class\x12j\n" +
	"\vUpdateClass\x12!.helloworld.v1.UpdateClassRequest\x1a\x1f.helloworld.v1.UpdateClassReply\"\x17\x82\xd3\xe4\x93\x02\x11:\x01*\x1a\fapi/v1/class\x12q\n" +
	"\vDeleteClass\x12!.helloworld.v1.DeleteClassRequest\x1a\x1f.helloworld.v1.DeleteClassReply\"\x1e\x82\xd3\xe4\x93\x02\x18:\x01*\x1a\x13api/v1/class/delete\x12i\n" +
	"\tListClass\x12\x1f.helloworld.v1.ListClassRequest\x1a\x1d.helloworld.v1.ListClassReply\"\x1c\x82\xd3\xe4\x93\x02\x16:\x01*\"\x11api/v1/class/list\x12c\n" +
	"\bGetClass\x12\x1e.helloworld.v1.GetClassRequest\x1a\x1c.helloworld.v1.GetClassReply\"\x19\x82\xd3\xe4\x93\x02\x13\x12\x11api/v1/class/{id}\x12\x83\x01\n" +
	"\x10ExportClassExcel\x12&.helloworld.v1.ExportClassExcelRequest\x1a$.helloworld.v1.ExportClassExcelReply\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19/api/v1/class/export/{id}\x12\x97\x01\n" +
	"\x18ListExportClassExcelData\x12*.helloworld.v1.ListClassExcelReportRequest\x1a,.helloworld.v1.ListClassExcelReportDataReply\"!\x82\xd3\xe4\x93\x02\x1b\x12\x19api/v1/class/exports/list\x12\x8e\x01\n" +
	"\x14ExportListClassExcel\x12*.helloworld.v1.ExportListClassExcelRequest\x1a(.helloworld.v1.ExportListClassExcelReply\" \x82\xd3\xe4\x93\x02\x1a\x12\x18/api/v1/class/reportListB\x1eZ\x1cDemoApp/api/helloworld/v1;v1P\x00P\x01b\x06proto3"

var file_api_helloworld_v1_classservice_proto_goTypes = []any{
	(*CreateClassRequest)(nil),            // 0: helloworld.v1.CreateClassRequest
	(*UpdateClassRequest)(nil),            // 1: helloworld.v1.UpdateClassRequest
	(*DeleteClassRequest)(nil),            // 2: helloworld.v1.DeleteClassRequest
	(*ListClassRequest)(nil),              // 3: helloworld.v1.ListClassRequest
	(*GetClassRequest)(nil),               // 4: helloworld.v1.GetClassRequest
	(*ExportClassExcelRequest)(nil),       // 5: helloworld.v1.ExportClassExcelRequest
	(*ListClassExcelReportRequest)(nil),   // 6: helloworld.v1.ListClassExcelReportRequest
	(*ExportListClassExcelRequest)(nil),   // 7: helloworld.v1.ExportListClassExcelRequest
	(*CreateClassReply)(nil),              // 8: helloworld.v1.CreateClassReply
	(*UpdateClassReply)(nil),              // 9: helloworld.v1.UpdateClassReply
	(*DeleteClassReply)(nil),              // 10: helloworld.v1.DeleteClassReply
	(*ListClassReply)(nil),                // 11: helloworld.v1.ListClassReply
	(*GetClassReply)(nil),                 // 12: helloworld.v1.GetClassReply
	(*ExportClassExcelReply)(nil),         // 13: helloworld.v1.ExportClassExcelReply
	(*ListClassExcelReportDataReply)(nil), // 14: helloworld.v1.ListClassExcelReportDataReply
	(*ExportListClassExcelReply)(nil),     // 15: helloworld.v1.ExportListClassExcelReply
}
var file_api_helloworld_v1_classservice_proto_depIdxs = []int32{
	0,  // 0: helloworld.v1.ClassService.CreateClass:input_type -> helloworld.v1.CreateClassRequest
	1,  // 1: helloworld.v1.ClassService.UpdateClass:input_type -> helloworld.v1.UpdateClassRequest
	2,  // 2: helloworld.v1.ClassService.DeleteClass:input_type -> helloworld.v1.DeleteClassRequest
	3,  // 3: helloworld.v1.ClassService.ListClass:input_type -> helloworld.v1.ListClassRequest
	4,  // 4: helloworld.v1.ClassService.GetClass:input_type -> helloworld.v1.GetClassRequest
	5,  // 5: helloworld.v1.ClassService.ExportClassExcel:input_type -> helloworld.v1.ExportClassExcelRequest
	6,  // 6: helloworld.v1.ClassService.ListExportClassExcelData:input_type -> helloworld.v1.ListClassExcelReportRequest
	7,  // 7: helloworld.v1.ClassService.ExportListClassExcel:input_type -> helloworld.v1.ExportListClassExcelRequest
	8,  // 8: helloworld.v1.ClassService.CreateClass:output_type -> helloworld.v1.CreateClassReply
	9,  // 9: helloworld.v1.ClassService.UpdateClass:output_type -> helloworld.v1.UpdateClassReply
	10, // 10: helloworld.v1.ClassService.DeleteClass:output_type -> helloworld.v1.DeleteClassReply
	11, // 11: helloworld.v1.ClassService.ListClass:output_type -> helloworld.v1.ListClassReply
	12, // 12: helloworld.v1.ClassService.GetClass:output_type -> helloworld.v1.GetClassReply
	13, // 13: helloworld.v1.ClassService.ExportClassExcel:output_type -> helloworld.v1.ExportClassExcelReply
	14, // 14: helloworld.v1.ClassService.ListExportClassExcelData:output_type -> helloworld.v1.ListClassExcelReportDataReply
	15, // 15: helloworld.v1.ClassService.ExportListClassExcel:output_type -> helloworld.v1.ExportListClassExcelReply
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_api_helloworld_v1_classservice_proto_init() }
func file_api_helloworld_v1_classservice_proto_init() {
	if File_api_helloworld_v1_classservice_proto != nil {
		return
	}
	file_api_helloworld_v1_class_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_helloworld_v1_classservice_proto_rawDesc), len(file_api_helloworld_v1_classservice_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_helloworld_v1_classservice_proto_goTypes,
		DependencyIndexes: file_api_helloworld_v1_classservice_proto_depIdxs,
	}.Build()
	File_api_helloworld_v1_classservice_proto = out.File
	file_api_helloworld_v1_classservice_proto_goTypes = nil
	file_api_helloworld_v1_classservice_proto_depIdxs = nil
}
