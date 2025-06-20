// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: api/helloworld/v1/student.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ClassId       int64                  `protobuf:"varint,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateStudentRequest) Reset() {
	*x = CreateStudentRequest{}
	mi := &file_api_helloworld_v1_student_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentRequest) ProtoMessage() {}

func (x *CreateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_student_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentRequest.ProtoReflect.Descriptor instead.
func (*CreateStudentRequest) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_student_proto_rawDescGZIP(), []int{0}
}

func (x *CreateStudentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateStudentRequest) GetClassId() int64 {
	if x != nil {
		return x.ClassId
	}
	return 0
}

type UpdateStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                          // ID của học sinh cần cập nhật
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                       // Tên học sinh mới
	ClassId       int64                  `protobuf:"varint,3,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"` // ID lớp học mới
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStudentRequest) Reset() {
	*x = UpdateStudentRequest{}
	mi := &file_api_helloworld_v1_student_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentRequest) ProtoMessage() {}

func (x *UpdateStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_student_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentRequest.ProtoReflect.Descriptor instead.
func (*UpdateStudentRequest) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_student_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateStudentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateStudentRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateStudentRequest) GetClassId() int64 {
	if x != nil {
		return x.ClassId
	}
	return 0
}

type DeleteStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // ID của học sinh cần xóa
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteStudentRequest) Reset() {
	*x = DeleteStudentRequest{}
	mi := &file_api_helloworld_v1_student_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentRequest) ProtoMessage() {}

func (x *DeleteStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_student_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentRequest.ProtoReflect.Descriptor instead.
func (*DeleteStudentRequest) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_student_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteStudentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          uint32                 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`                                  // Số trang
	PageSize      uint32                 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`          // Số lượng bản ghi trên mỗi trang
	Name          *string                `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`                             // Tên học sinh để lọc
	ClassId       *int64                 `protobuf:"varint,4,opt,name=class_id,json=classId,proto3,oneof" json:"class_id,omitempty"`       // ID lớp học để lọc
	IsDeleted     *bool                  `protobuf:"varint,5,opt,name=is_deleted,json=isDeleted,proto3,oneof" json:"is_deleted,omitempty"` // Lọc theo trạng thái xóa
	Keyword       *string                `protobuf:"bytes,6,opt,name=keyword,proto3,oneof" json:"keyword,omitempty"`                       // Từ khóa tìm kiếm
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStudentRequest) Reset() {
	*x = ListStudentRequest{}
	mi := &file_api_helloworld_v1_student_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStudentRequest) ProtoMessage() {}

func (x *ListStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_student_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStudentRequest.ProtoReflect.Descriptor instead.
func (*ListStudentRequest) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_student_proto_rawDescGZIP(), []int{3}
}

func (x *ListStudentRequest) GetPage() uint32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListStudentRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListStudentRequest) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *ListStudentRequest) GetClassId() int64 {
	if x != nil && x.ClassId != nil {
		return *x.ClassId
	}
	return 0
}

func (x *ListStudentRequest) GetIsDeleted() bool {
	if x != nil && x.IsDeleted != nil {
		return *x.IsDeleted
	}
	return false
}

func (x *ListStudentRequest) GetKeyword() string {
	if x != nil && x.Keyword != nil {
		return *x.Keyword
	}
	return ""
}

type GetStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"` // ID của học sinh cần lấy thông tin
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStudentRequest) Reset() {
	*x = GetStudentRequest{}
	mi := &file_api_helloworld_v1_student_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentRequest) ProtoMessage() {}

func (x *GetStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_student_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentRequest.ProtoReflect.Descriptor instead.
func (*GetStudentRequest) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_student_proto_rawDescGZIP(), []int{4}
}

func (x *GetStudentRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ListStudentReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*StudentData         `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	Total         int64                  `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListStudentReply) Reset() {
	*x = ListStudentReply{}
	mi := &file_api_helloworld_v1_student_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListStudentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListStudentReply) ProtoMessage() {}

func (x *ListStudentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_student_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListStudentReply.ProtoReflect.Descriptor instead.
func (*ListStudentReply) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_student_proto_rawDescGZIP(), []int{5}
}

func (x *ListStudentReply) GetItems() []*StudentData {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *ListStudentReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type GetStudentReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Student       *StudentData           `protobuf:"bytes,1,opt,name=student,proto3" json:"student,omitempty"` // Thông tin học sinh
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetStudentReply) Reset() {
	*x = GetStudentReply{}
	mi := &file_api_helloworld_v1_student_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetStudentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStudentReply) ProtoMessage() {}

func (x *GetStudentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_student_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStudentReply.ProtoReflect.Descriptor instead.
func (*GetStudentReply) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_student_proto_rawDescGZIP(), []int{6}
}

func (x *GetStudentReply) GetStudent() *StudentData {
	if x != nil {
		return x.Student
	}
	return nil
}

type CreateStudentReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Thông báo kết quả tạo học sinh
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateStudentReply) Reset() {
	*x = CreateStudentReply{}
	mi := &file_api_helloworld_v1_student_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateStudentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStudentReply) ProtoMessage() {}

func (x *CreateStudentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_student_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStudentReply.ProtoReflect.Descriptor instead.
func (*CreateStudentReply) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_student_proto_rawDescGZIP(), []int{7}
}

func (x *CreateStudentReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type UpdateStudentReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Thông báo kết quả cập nhật học sinh
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateStudentReply) Reset() {
	*x = UpdateStudentReply{}
	mi := &file_api_helloworld_v1_student_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateStudentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateStudentReply) ProtoMessage() {}

func (x *UpdateStudentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_student_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateStudentReply.ProtoReflect.Descriptor instead.
func (*UpdateStudentReply) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_student_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateStudentReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteStudentReply struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"` // Thông báo kết quả xóa học sinh
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteStudentReply) Reset() {
	*x = DeleteStudentReply{}
	mi := &file_api_helloworld_v1_student_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteStudentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStudentReply) ProtoMessage() {}

func (x *DeleteStudentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_student_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStudentReply.ProtoReflect.Descriptor instead.
func (*DeleteStudentReply) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_student_proto_rawDescGZIP(), []int{9}
}

func (x *DeleteStudentReply) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type StudentData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ClassId       int64                  `protobuf:"varint,3,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`             // ID lớp học của học sinh
	ClassName     string                 `protobuf:"bytes,4,opt,name=class_name,json=className,proto3" json:"class_name,omitempty"`        // Tên lớp học của học sinh
	IsDeleted     *bool                  `protobuf:"varint,5,opt,name=is_deleted,json=isDeleted,proto3,oneof" json:"is_deleted,omitempty"` // Trạng thái xóa của học sinh
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StudentData) Reset() {
	*x = StudentData{}
	mi := &file_api_helloworld_v1_student_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StudentData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StudentData) ProtoMessage() {}

func (x *StudentData) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_student_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StudentData.ProtoReflect.Descriptor instead.
func (*StudentData) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_student_proto_rawDescGZIP(), []int{10}
}

func (x *StudentData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *StudentData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *StudentData) GetClassId() int64 {
	if x != nil {
		return x.ClassId
	}
	return 0
}

func (x *StudentData) GetClassName() string {
	if x != nil {
		return x.ClassName
	}
	return ""
}

func (x *StudentData) GetIsDeleted() bool {
	if x != nil && x.IsDeleted != nil {
		return *x.IsDeleted
	}
	return false
}

var File_api_helloworld_v1_student_proto protoreflect.FileDescriptor

const file_api_helloworld_v1_student_proto_rawDesc = "" +
	"\n" +
	"\x1fapi/helloworld/v1/student.proto\x12\rhelloworld.v1\x1a\x1dapi/helloworld/v1/class.proto\"E\n" +
	"\x14CreateStudentRequest\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x19\n" +
	"\bclass_id\x18\x02 \x01(\x03R\aclassId\"U\n" +
	"\x14UpdateStudentRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x19\n" +
	"\bclass_id\x18\x03 \x01(\x03R\aclassId\"&\n" +
	"\x14DeleteStudentRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"\xf2\x01\n" +
	"\x12ListStudentRequest\x12\x12\n" +
	"\x04page\x18\x01 \x01(\rR\x04page\x12\x1b\n" +
	"\tpage_size\x18\x02 \x01(\rR\bpageSize\x12\x17\n" +
	"\x04name\x18\x03 \x01(\tH\x00R\x04name\x88\x01\x01\x12\x1e\n" +
	"\bclass_id\x18\x04 \x01(\x03H\x01R\aclassId\x88\x01\x01\x12\"\n" +
	"\n" +
	"is_deleted\x18\x05 \x01(\bH\x02R\tisDeleted\x88\x01\x01\x12\x1d\n" +
	"\akeyword\x18\x06 \x01(\tH\x03R\akeyword\x88\x01\x01B\a\n" +
	"\x05_nameB\v\n" +
	"\t_class_idB\r\n" +
	"\v_is_deletedB\n" +
	"\n" +
	"\b_keyword\"#\n" +
	"\x11GetStudentRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"Z\n" +
	"\x10ListStudentReply\x120\n" +
	"\x05items\x18\x01 \x03(\v2\x1a.helloworld.v1.StudentDataR\x05items\x12\x14\n" +
	"\x05total\x18\x02 \x01(\x03R\x05total\"G\n" +
	"\x0fGetStudentReply\x124\n" +
	"\astudent\x18\x01 \x01(\v2\x1a.helloworld.v1.StudentDataR\astudent\".\n" +
	"\x12CreateStudentReply\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\".\n" +
	"\x12UpdateStudentReply\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\".\n" +
	"\x12DeleteStudentReply\x12\x18\n" +
	"\amessage\x18\x01 \x01(\tR\amessage\"\x9e\x01\n" +
	"\vStudentData\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x19\n" +
	"\bclass_id\x18\x03 \x01(\x03R\aclassId\x12\x1d\n" +
	"\n" +
	"class_name\x18\x04 \x01(\tR\tclassName\x12\"\n" +
	"\n" +
	"is_deleted\x18\x05 \x01(\bH\x00R\tisDeleted\x88\x01\x01B\r\n" +
	"\v_is_deletedB\x1eZ\x1cDemoApp/api/helloworld/v1;v1b\x06proto3"

var (
	file_api_helloworld_v1_student_proto_rawDescOnce sync.Once
	file_api_helloworld_v1_student_proto_rawDescData []byte
)

func file_api_helloworld_v1_student_proto_rawDescGZIP() []byte {
	file_api_helloworld_v1_student_proto_rawDescOnce.Do(func() {
		file_api_helloworld_v1_student_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_helloworld_v1_student_proto_rawDesc), len(file_api_helloworld_v1_student_proto_rawDesc)))
	})
	return file_api_helloworld_v1_student_proto_rawDescData
}

var file_api_helloworld_v1_student_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_helloworld_v1_student_proto_goTypes = []any{
	(*CreateStudentRequest)(nil), // 0: helloworld.v1.CreateStudentRequest
	(*UpdateStudentRequest)(nil), // 1: helloworld.v1.UpdateStudentRequest
	(*DeleteStudentRequest)(nil), // 2: helloworld.v1.DeleteStudentRequest
	(*ListStudentRequest)(nil),   // 3: helloworld.v1.ListStudentRequest
	(*GetStudentRequest)(nil),    // 4: helloworld.v1.GetStudentRequest
	(*ListStudentReply)(nil),     // 5: helloworld.v1.ListStudentReply
	(*GetStudentReply)(nil),      // 6: helloworld.v1.GetStudentReply
	(*CreateStudentReply)(nil),   // 7: helloworld.v1.CreateStudentReply
	(*UpdateStudentReply)(nil),   // 8: helloworld.v1.UpdateStudentReply
	(*DeleteStudentReply)(nil),   // 9: helloworld.v1.DeleteStudentReply
	(*StudentData)(nil),          // 10: helloworld.v1.StudentData
}
var file_api_helloworld_v1_student_proto_depIdxs = []int32{
	10, // 0: helloworld.v1.ListStudentReply.items:type_name -> helloworld.v1.StudentData
	10, // 1: helloworld.v1.GetStudentReply.student:type_name -> helloworld.v1.StudentData
	2,  // [2:2] is the sub-list for method output_type
	2,  // [2:2] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_helloworld_v1_student_proto_init() }
func file_api_helloworld_v1_student_proto_init() {
	if File_api_helloworld_v1_student_proto != nil {
		return
	}
	file_api_helloworld_v1_class_proto_init()
	file_api_helloworld_v1_student_proto_msgTypes[3].OneofWrappers = []any{}
	file_api_helloworld_v1_student_proto_msgTypes[10].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_helloworld_v1_student_proto_rawDesc), len(file_api_helloworld_v1_student_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_helloworld_v1_student_proto_goTypes,
		DependencyIndexes: file_api_helloworld_v1_student_proto_depIdxs,
		MessageInfos:      file_api_helloworld_v1_student_proto_msgTypes,
	}.Build()
	File_api_helloworld_v1_student_proto = out.File
	file_api_helloworld_v1_student_proto_goTypes = nil
	file_api_helloworld_v1_student_proto_depIdxs = nil
}
