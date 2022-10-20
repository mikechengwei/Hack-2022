// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.7
// source: msg.proto

package msg

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

type RespCode int32

const (
	RespCode_rc_OK RespCode = 0
)

// Enum value maps for RespCode.
var (
	RespCode_name = map[int32]string{
		0: "rc_OK",
	}
	RespCode_value = map[string]int32{
		"rc_OK": 0,
	}
)

func (x RespCode) Enum() *RespCode {
	p := new(RespCode)
	*p = x
	return p
}

func (x RespCode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RespCode) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[0].Descriptor()
}

func (RespCode) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[0]
}

func (x RespCode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RespCode.Descriptor instead.
func (RespCode) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

type TaskState int32

const (
	TaskState_ts_Init      TaskState = 0
	TaskState_ts_Create    TaskState = 1
	TaskState_ts_Dumpling  TaskState = 2
	TaskState_ts_Exception TaskState = 16
	TaskState_ts_Finish    TaskState = 32
)

// Enum value maps for TaskState.
var (
	TaskState_name = map[int32]string{
		0:  "ts_Init",
		1:  "ts_Create",
		2:  "ts_Dumpling",
		16: "ts_Exception",
		32: "ts_Finish",
	}
	TaskState_value = map[string]int32{
		"ts_Init":      0,
		"ts_Create":    1,
		"ts_Dumpling":  2,
		"ts_Exception": 16,
		"ts_Finish":    32,
	}
)

func (x TaskState) Enum() *TaskState {
	p := new(TaskState)
	*p = x
	return p
}

func (x TaskState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskState) Descriptor() protoreflect.EnumDescriptor {
	return file_msg_proto_enumTypes[1].Descriptor()
}

func (TaskState) Type() protoreflect.EnumType {
	return &file_msg_proto_enumTypes[1]
}

func (x TaskState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskState.Descriptor instead.
func (TaskState) EnumDescriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

type ClientInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key     string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Address string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
}

func (x *ClientInfo) Reset() {
	*x = ClientInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientInfo) ProtoMessage() {}

func (x *ClientInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ClientInfo.ProtoReflect.Descriptor instead.
func (*ClientInfo) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{0}
}

func (x *ClientInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *ClientInfo) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

type ServerInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskAddress string `protobuf:"bytes,1,opt,name=taskAddress,proto3" json:"taskAddress,omitempty"`
}

func (x *ServerInfo) Reset() {
	*x = ServerInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerInfo) ProtoMessage() {}

func (x *ServerInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ServerInfo.ProtoReflect.Descriptor instead.
func (*ServerInfo) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{1}
}

func (x *ServerInfo) GetTaskAddress() string {
	if x != nil {
		return x.TaskAddress
	}
	return ""
}

type TaskInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Key    string     `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Source *TableInfo `protobuf:"bytes,3,opt,name=source,proto3" json:"source,omitempty"`
	Target *TableInfo `protobuf:"bytes,4,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *TaskInfo) Reset() {
	*x = TaskInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskInfo) ProtoMessage() {}

func (x *TaskInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TaskInfo.ProtoReflect.Descriptor instead.
func (*TaskInfo) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{2}
}

func (x *TaskInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TaskInfo) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *TaskInfo) GetSource() *TableInfo {
	if x != nil {
		return x.Source
	}
	return nil
}

func (x *TaskInfo) GetTarget() *TableInfo {
	if x != nil {
		return x.Target
	}
	return nil
}

type TableInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Host     string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Port     int32  `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	Username string `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Db       string `protobuf:"bytes,5,opt,name=db,proto3" json:"db,omitempty"`
	Tbl      string `protobuf:"bytes,6,opt,name=tbl,proto3" json:"tbl,omitempty"`
}

func (x *TableInfo) Reset() {
	*x = TableInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableInfo) ProtoMessage() {}

func (x *TableInfo) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use TableInfo.ProtoReflect.Descriptor instead.
func (*TableInfo) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{3}
}

func (x *TableInfo) GetHost() string {
	if x != nil {
		return x.Host
	}
	return ""
}

func (x *TableInfo) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *TableInfo) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *TableInfo) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *TableInfo) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *TableInfo) GetTbl() string {
	if x != nil {
		return x.Tbl
	}
	return ""
}

type ReplyBase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rc      RespCode `protobuf:"varint,1,opt,name=rc,proto3,enum=msg.RespCode" json:"rc,omitempty"`
	RespMsg string   `protobuf:"bytes,2,opt,name=RespMsg,proto3" json:"RespMsg,omitempty"`
}

func (x *ReplyBase) Reset() {
	*x = ReplyBase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyBase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyBase) ProtoMessage() {}

func (x *ReplyBase) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyBase.ProtoReflect.Descriptor instead.
func (*ReplyBase) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{4}
}

func (x *ReplyBase) GetRc() RespCode {
	if x != nil {
		return x.Rc
	}
	return RespCode_rc_OK
}

func (x *ReplyBase) GetRespMsg() string {
	if x != nil {
		return x.RespMsg
	}
	return ""
}

type ReqRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cli *ClientInfo `protobuf:"bytes,1,opt,name=cli,proto3" json:"cli,omitempty"`
}

func (x *ReqRegister) Reset() {
	*x = ReqRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqRegister) ProtoMessage() {}

func (x *ReqRegister) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqRegister.ProtoReflect.Descriptor instead.
func (*ReqRegister) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{5}
}

func (x *ReqRegister) GetCli() *ClientInfo {
	if x != nil {
		return x.Cli
	}
	return nil
}

type ReplyRegister struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rc *ReplyBase `protobuf:"bytes,1,opt,name=rc,proto3" json:"rc,omitempty"`
}

func (x *ReplyRegister) Reset() {
	*x = ReplyRegister{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyRegister) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyRegister) ProtoMessage() {}

func (x *ReplyRegister) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyRegister.ProtoReflect.Descriptor instead.
func (*ReplyRegister) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{6}
}

func (x *ReplyRegister) GetRc() *ReplyBase {
	if x != nil {
		return x.Rc
	}
	return nil
}

type ReqNewTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cli    *ClientInfo `protobuf:"bytes,1,opt,name=cli,proto3" json:"cli,omitempty"`
	Task   *TaskInfo   `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
	Server *ServerInfo `protobuf:"bytes,3,opt,name=server,proto3" json:"server,omitempty"`
}

func (x *ReqNewTask) Reset() {
	*x = ReqNewTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqNewTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqNewTask) ProtoMessage() {}

func (x *ReqNewTask) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqNewTask.ProtoReflect.Descriptor instead.
func (*ReqNewTask) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{7}
}

func (x *ReqNewTask) GetCli() *ClientInfo {
	if x != nil {
		return x.Cli
	}
	return nil
}

func (x *ReqNewTask) GetTask() *TaskInfo {
	if x != nil {
		return x.Task
	}
	return nil
}

func (x *ReqNewTask) GetServer() *ServerInfo {
	if x != nil {
		return x.Server
	}
	return nil
}

type ReplyNewTask struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rc *ReplyBase `protobuf:"bytes,1,opt,name=rc,proto3" json:"rc,omitempty"`
}

func (x *ReplyNewTask) Reset() {
	*x = ReplyNewTask{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyNewTask) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyNewTask) ProtoMessage() {}

func (x *ReplyNewTask) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyNewTask.ProtoReflect.Descriptor instead.
func (*ReplyNewTask) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{8}
}

func (x *ReplyNewTask) GetRc() *ReplyBase {
	if x != nil {
		return x.Rc
	}
	return nil
}

type ReqReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cli  *ClientInfo `protobuf:"bytes,1,opt,name=cli,proto3" json:"cli,omitempty"`
	Task *TaskInfo   `protobuf:"bytes,2,opt,name=task,proto3" json:"task,omitempty"`
}

func (x *ReqReport) Reset() {
	*x = ReqReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqReport) ProtoMessage() {}

func (x *ReqReport) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqReport.ProtoReflect.Descriptor instead.
func (*ReqReport) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{9}
}

func (x *ReqReport) GetCli() *ClientInfo {
	if x != nil {
		return x.Cli
	}
	return nil
}

func (x *ReqReport) GetTask() *TaskInfo {
	if x != nil {
		return x.Task
	}
	return nil
}

type ReplyReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rc       *ReplyBase `protobuf:"bytes,1,opt,name=rc,proto3" json:"rc,omitempty"`
	State    TaskState  `protobuf:"varint,2,opt,name=state,proto3,enum=msg.TaskState" json:"state,omitempty"`
	Progress string     `protobuf:"bytes,3,opt,name=progress,proto3" json:"progress,omitempty"`
}

func (x *ReplyReport) Reset() {
	*x = ReplyReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyReport) ProtoMessage() {}

func (x *ReplyReport) ProtoReflect() protoreflect.Message {
	mi := &file_msg_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyReport.ProtoReflect.Descriptor instead.
func (*ReplyReport) Descriptor() ([]byte, []int) {
	return file_msg_proto_rawDescGZIP(), []int{10}
}

func (x *ReplyReport) GetRc() *ReplyBase {
	if x != nil {
		return x.Rc
	}
	return nil
}

func (x *ReplyReport) GetState() TaskState {
	if x != nil {
		return x.State
	}
	return TaskState_ts_Init
}

func (x *ReplyReport) GetProgress() string {
	if x != nil {
		return x.Progress
	}
	return ""
}

var File_msg_proto protoreflect.FileDescriptor

var file_msg_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x4c, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x2e,
	0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b,
	0x74, 0x61, 0x73, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x80,
	0x01, 0x0a, 0x08, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x26, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x74, 0x61, 0x72,
	0x67, 0x65, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x22, 0x8d, 0x01, 0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12,
	0x12, 0x0a, 0x04, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x68,
	0x6f, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12,
	0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x12,
	0x10, 0x0a, 0x03, 0x74, 0x62, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x74, 0x62,
	0x6c, 0x22, 0x44, 0x0a, 0x09, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x61, 0x73, 0x65, 0x12, 0x1d,
	0x0a, 0x02, 0x72, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x02, 0x72, 0x63, 0x12, 0x18, 0x0a,
	0x07, 0x52, 0x65, 0x73, 0x70, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x52, 0x65, 0x73, 0x70, 0x4d, 0x73, 0x67, 0x22, 0x30, 0x0a, 0x0b, 0x52, 0x65, 0x71, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x6c, 0x69, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x63, 0x6c, 0x69, 0x22, 0x2f, 0x0a, 0x0d, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x02, 0x72, 0x63,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x42, 0x61, 0x73, 0x65, 0x52, 0x02, 0x72, 0x63, 0x22, 0x7b, 0x0a, 0x0a, 0x52, 0x65,
	0x71, 0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x6c, 0x69, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x03, 0x63, 0x6c, 0x69, 0x12, 0x21, 0x0a, 0x04, 0x74,
	0x61, 0x73, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x12, 0x27,
	0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x22, 0x2e, 0x0a, 0x0c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x4e, 0x65, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x1e, 0x0a, 0x02, 0x72, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x02, 0x72, 0x63, 0x22, 0x51, 0x0a, 0x09, 0x52, 0x65, 0x71, 0x52, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x12, 0x21, 0x0a, 0x03, 0x63, 0x6c, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x03, 0x63, 0x6c, 0x69, 0x12, 0x21, 0x0a, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x6f, 0x0a, 0x0b, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x12, 0x1e, 0x0a, 0x02, 0x72, 0x63, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x72, 0x65, 0x70, 0x6c,
	0x79, 0x42, 0x61, 0x73, 0x65, 0x52, 0x02, 0x72, 0x63, 0x12, 0x24, 0x0a, 0x05, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0e, 0x2e, 0x6d, 0x73, 0x67, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x67, 0x72, 0x65, 0x73, 0x73, 0x2a, 0x15, 0x0a, 0x08, 0x52,
	0x65, 0x73, 0x70, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x72, 0x63, 0x5f, 0x4f, 0x4b,
	0x10, 0x00, 0x2a, 0x59, 0x0a, 0x09, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x0b, 0x0a, 0x07, 0x74, 0x73, 0x5f, 0x49, 0x6e, 0x69, 0x74, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x74, 0x73, 0x5f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x74,
	0x73, 0x5f, 0x44, 0x75, 0x6d, 0x70, 0x6c, 0x69, 0x6e, 0x67, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c,
	0x74, 0x73, 0x5f, 0x45, 0x78, 0x63, 0x65, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x10, 0x10, 0x12, 0x0d,
	0x0a, 0x09, 0x74, 0x73, 0x5f, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x10, 0x20, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2e, 0x2f, 0x6d, 0x73, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_msg_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_msg_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_msg_proto_goTypes = []interface{}{
	(RespCode)(0),         // 0: msg.RespCode
	(TaskState)(0),        // 1: msg.TaskState
	(*ClientInfo)(nil),    // 2: msg.ClientInfo
	(*ServerInfo)(nil),    // 3: msg.ServerInfo
	(*TaskInfo)(nil),      // 4: msg.TaskInfo
	(*TableInfo)(nil),     // 5: msg.TableInfo
	(*ReplyBase)(nil),     // 6: msg.replyBase
	(*ReqRegister)(nil),   // 7: msg.ReqRegister
	(*ReplyRegister)(nil), // 8: msg.ReplyRegister
	(*ReqNewTask)(nil),    // 9: msg.ReqNewTask
	(*ReplyNewTask)(nil),  // 10: msg.ReplyNewTask
	(*ReqReport)(nil),     // 11: msg.ReqReport
	(*ReplyReport)(nil),   // 12: msg.ReplyReport
}
var file_msg_proto_depIdxs = []int32{
	5,  // 0: msg.TaskInfo.source:type_name -> msg.TableInfo
	5,  // 1: msg.TaskInfo.target:type_name -> msg.TableInfo
	0,  // 2: msg.replyBase.rc:type_name -> msg.RespCode
	2,  // 3: msg.ReqRegister.cli:type_name -> msg.ClientInfo
	6,  // 4: msg.ReplyRegister.rc:type_name -> msg.replyBase
	2,  // 5: msg.ReqNewTask.cli:type_name -> msg.ClientInfo
	4,  // 6: msg.ReqNewTask.task:type_name -> msg.TaskInfo
	3,  // 7: msg.ReqNewTask.server:type_name -> msg.ServerInfo
	6,  // 8: msg.ReplyNewTask.rc:type_name -> msg.replyBase
	2,  // 9: msg.ReqReport.cli:type_name -> msg.ClientInfo
	4,  // 10: msg.ReqReport.task:type_name -> msg.TaskInfo
	6,  // 11: msg.ReplyReport.rc:type_name -> msg.replyBase
	1,  // 12: msg.ReplyReport.state:type_name -> msg.TaskState
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_msg_proto_init() }
func file_msg_proto_init() {
	if File_msg_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientInfo); i {
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
			switch v := v.(*ServerInfo); i {
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
			switch v := v.(*TaskInfo); i {
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
			switch v := v.(*TableInfo); i {
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
		file_msg_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyBase); i {
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
		file_msg_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqRegister); i {
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
		file_msg_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyRegister); i {
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
		file_msg_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqNewTask); i {
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
		file_msg_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyNewTask); i {
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
		file_msg_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqReport); i {
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
		file_msg_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyReport); i {
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
			NumEnums:      2,
			NumMessages:   11,
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
