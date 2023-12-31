// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: idl/note.proto

package notedemo

import (
	context "context"
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

// BaseResp 定义
type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int64  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`
	StatusMessage string `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"`
	ServiceTime   int64  `protobuf:"varint,3,opt,name=service_time,json=serviceTime,proto3" json:"service_time,omitempty"`
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetStatusCode() int64 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResp) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *BaseResp) GetServiceTime() int64 {
	if x != nil {
		return x.ServiceTime
	}
	return 0
}

// Note 定义
type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId     int64  `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	UserId     int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UserName   string `protobuf:"bytes,3,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	UserAvatar string `protobuf:"bytes,4,opt,name=user_avatar,json=userAvatar,proto3" json:"user_avatar,omitempty"`
	Title      string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Content    string `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	CreateTime int64  `protobuf:"varint,7,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Note.ProtoReflect.Descriptor instead.
func (*Note) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{1}
}

func (x *Note) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

func (x *Note) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Note) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *Note) GetUserAvatar() string {
	if x != nil {
		return x.UserAvatar
	}
	return ""
}

func (x *Note) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Note) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Note) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

// CreateNoteRequest 定义
type CreateNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	UserId  int64  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *CreateNoteRequest) Reset() {
	*x = CreateNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteRequest) ProtoMessage() {}

func (x *CreateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteRequest.ProtoReflect.Descriptor instead.
func (*CreateNoteRequest) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{2}
}

func (x *CreateNoteRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateNoteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *CreateNoteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// CreateNoteResponse 定义
type CreateNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *CreateNoteResponse) Reset() {
	*x = CreateNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateNoteResponse) ProtoMessage() {}

func (x *CreateNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateNoteResponse.ProtoReflect.Descriptor instead.
func (*CreateNoteResponse) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{3}
}

func (x *CreateNoteResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

// DeleteNoteRequest 定义
type DeleteNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId int64 `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	UserId int64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *DeleteNoteRequest) Reset() {
	*x = DeleteNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNoteRequest) ProtoMessage() {}

func (x *DeleteNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNoteRequest.ProtoReflect.Descriptor instead.
func (*DeleteNoteRequest) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteNoteRequest) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

func (x *DeleteNoteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

// DeleteNoteResponse 定义
type DeleteNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *DeleteNoteResponse) Reset() {
	*x = DeleteNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteNoteResponse) ProtoMessage() {}

func (x *DeleteNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteNoteResponse.ProtoReflect.Descriptor instead.
func (*DeleteNoteResponse) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteNoteResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

// UpdateNoteRequest 定义
type UpdateNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteId  int64  `protobuf:"varint,1,opt,name=note_id,json=noteId,proto3" json:"note_id,omitempty"`
	UserId  int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title   string `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Content string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdateNoteRequest) Reset() {
	*x = UpdateNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteRequest) ProtoMessage() {}

func (x *UpdateNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteRequest.ProtoReflect.Descriptor instead.
func (*UpdateNoteRequest) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateNoteRequest) GetNoteId() int64 {
	if x != nil {
		return x.NoteId
	}
	return 0
}

func (x *UpdateNoteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UpdateNoteRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateNoteRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// UpdateNoteResponse 定义
type UpdateNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *UpdateNoteResponse) Reset() {
	*x = UpdateNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNoteResponse) ProtoMessage() {}

func (x *UpdateNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNoteResponse.ProtoReflect.Descriptor instead.
func (*UpdateNoteResponse) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateNoteResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

// MGetNoteRequest 定义
type MGetNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NoteIds []int64 `protobuf:"varint,1,rep,packed,name=note_ids,json=noteIds,proto3" json:"note_ids,omitempty"`
}

func (x *MGetNoteRequest) Reset() {
	*x = MGetNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MGetNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MGetNoteRequest) ProtoMessage() {}

func (x *MGetNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MGetNoteRequest.ProtoReflect.Descriptor instead.
func (*MGetNoteRequest) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{8}
}

func (x *MGetNoteRequest) GetNoteIds() []int64 {
	if x != nil {
		return x.NoteIds
	}
	return nil
}

// MGetNoteResponse 定义
type MGetNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notes    []*Note   `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	BaseResp *BaseResp `protobuf:"bytes,2,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *MGetNoteResponse) Reset() {
	*x = MGetNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MGetNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MGetNoteResponse) ProtoMessage() {}

func (x *MGetNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MGetNoteResponse.ProtoReflect.Descriptor instead.
func (*MGetNoteResponse) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{9}
}

func (x *MGetNoteResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *MGetNoteResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

// QueryNoteRequest 定义
type QueryNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SearchKey string `protobuf:"bytes,2,opt,name=search_key,json=searchKey,proto3" json:"search_key,omitempty"`
	Offset    int64  `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit     int64  `protobuf:"varint,4,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *QueryNoteRequest) Reset() {
	*x = QueryNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryNoteRequest) ProtoMessage() {}

func (x *QueryNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryNoteRequest.ProtoReflect.Descriptor instead.
func (*QueryNoteRequest) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{10}
}

func (x *QueryNoteRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *QueryNoteRequest) GetSearchKey() string {
	if x != nil {
		return x.SearchKey
	}
	return ""
}

func (x *QueryNoteRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *QueryNoteRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

// QueryNoteResponse 定义
type QueryNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Notes    []*Note   `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	Total    int64     `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	BaseResp *BaseResp `protobuf:"bytes,3,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
}

func (x *QueryNoteResponse) Reset() {
	*x = QueryNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_note_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryNoteResponse) ProtoMessage() {}

func (x *QueryNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_note_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryNoteResponse.ProtoReflect.Descriptor instead.
func (*QueryNoteResponse) Descriptor() ([]byte, []int) {
	return file_idl_note_proto_rawDescGZIP(), []int{11}
}

func (x *QueryNoteResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *QueryNoteResponse) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *QueryNoteResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

var File_idl_note_proto protoreflect.FileDescriptor

var file_idl_note_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x22, 0x75, 0x0a, 0x08, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xc7, 0x01,
	0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x75, 0x73, 0x65, 0x72,
	0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x5c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e,
	0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x62,
	0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x45, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22,
	0x41, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x75, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64,
	0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x41, 0x0a, 0x12, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x2c, 0x0a, 0x0f,
	0x4d, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x6e, 0x6f, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x03, 0x52, 0x07, 0x6e, 0x6f, 0x74, 0x65, 0x49, 0x64, 0x73, 0x22, 0x61, 0x0a, 0x10, 0x4d, 0x47,
	0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x12, 0x2b, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x78, 0x0a,
	0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x4b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x78, 0x0a, 0x11, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x05,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x42,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x32, 0xc9, 0x02, 0x0a, 0x0b, 0x4e, 0x6f, 0x74, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12,
	0x17, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x39, 0x0a, 0x08, 0x4d, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x15,
	0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x4d, 0x47, 0x65, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x4d, 0x47, 0x65,
	0x74, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c,
	0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x16, 0x2e, 0x6e, 0x6f,
	0x74, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x51, 0x75, 0x65, 0x72, 0x79,
	0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x6e, 0x6f, 0x74,
	0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x6e, 0x6f, 0x74, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a,
	0x1b, 0x63, 0x75, 0x72, 0x64, 0x64, 0x65, 0x6d, 0x6f, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78, 0x5f,
	0x67, 0x65, 0x6e, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x64, 0x65, 0x6d, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_note_proto_rawDescOnce sync.Once
	file_idl_note_proto_rawDescData = file_idl_note_proto_rawDesc
)

func file_idl_note_proto_rawDescGZIP() []byte {
	file_idl_note_proto_rawDescOnce.Do(func() {
		file_idl_note_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_note_proto_rawDescData)
	})
	return file_idl_note_proto_rawDescData
}

var file_idl_note_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_idl_note_proto_goTypes = []interface{}{
	(*BaseResp)(nil),           // 0: note.BaseResp
	(*Note)(nil),               // 1: note.Note
	(*CreateNoteRequest)(nil),  // 2: note.CreateNoteRequest
	(*CreateNoteResponse)(nil), // 3: note.CreateNoteResponse
	(*DeleteNoteRequest)(nil),  // 4: note.DeleteNoteRequest
	(*DeleteNoteResponse)(nil), // 5: note.DeleteNoteResponse
	(*UpdateNoteRequest)(nil),  // 6: note.UpdateNoteRequest
	(*UpdateNoteResponse)(nil), // 7: note.UpdateNoteResponse
	(*MGetNoteRequest)(nil),    // 8: note.MGetNoteRequest
	(*MGetNoteResponse)(nil),   // 9: note.MGetNoteResponse
	(*QueryNoteRequest)(nil),   // 10: note.QueryNoteRequest
	(*QueryNoteResponse)(nil),  // 11: note.QueryNoteResponse
}
var file_idl_note_proto_depIdxs = []int32{
	0,  // 0: note.CreateNoteResponse.base_resp:type_name -> note.BaseResp
	0,  // 1: note.DeleteNoteResponse.base_resp:type_name -> note.BaseResp
	0,  // 2: note.UpdateNoteResponse.base_resp:type_name -> note.BaseResp
	1,  // 3: note.MGetNoteResponse.notes:type_name -> note.Note
	0,  // 4: note.MGetNoteResponse.base_resp:type_name -> note.BaseResp
	1,  // 5: note.QueryNoteResponse.notes:type_name -> note.Note
	0,  // 6: note.QueryNoteResponse.base_resp:type_name -> note.BaseResp
	2,  // 7: note.NoteService.CreateNote:input_type -> note.CreateNoteRequest
	8,  // 8: note.NoteService.MGetNote:input_type -> note.MGetNoteRequest
	4,  // 9: note.NoteService.DeleteNote:input_type -> note.DeleteNoteRequest
	10, // 10: note.NoteService.QueryNote:input_type -> note.QueryNoteRequest
	6,  // 11: note.NoteService.UpdateNote:input_type -> note.UpdateNoteRequest
	3,  // 12: note.NoteService.CreateNote:output_type -> note.CreateNoteResponse
	9,  // 13: note.NoteService.MGetNote:output_type -> note.MGetNoteResponse
	5,  // 14: note.NoteService.DeleteNote:output_type -> note.DeleteNoteResponse
	11, // 15: note.NoteService.QueryNote:output_type -> note.QueryNoteResponse
	7,  // 16: note.NoteService.UpdateNote:output_type -> note.UpdateNoteResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_idl_note_proto_init() }
func file_idl_note_proto_init() {
	if File_idl_note_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_note_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
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
		file_idl_note_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Note); i {
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
		file_idl_note_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNoteRequest); i {
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
		file_idl_note_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateNoteResponse); i {
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
		file_idl_note_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNoteRequest); i {
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
		file_idl_note_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteNoteResponse); i {
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
		file_idl_note_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNoteRequest); i {
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
		file_idl_note_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNoteResponse); i {
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
		file_idl_note_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MGetNoteRequest); i {
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
		file_idl_note_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MGetNoteResponse); i {
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
		file_idl_note_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryNoteRequest); i {
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
		file_idl_note_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryNoteResponse); i {
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
			RawDescriptor: file_idl_note_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_note_proto_goTypes,
		DependencyIndexes: file_idl_note_proto_depIdxs,
		MessageInfos:      file_idl_note_proto_msgTypes,
	}.Build()
	File_idl_note_proto = out.File
	file_idl_note_proto_rawDesc = nil
	file_idl_note_proto_goTypes = nil
	file_idl_note_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.7.2. DO NOT EDIT.

type NoteService interface {
	CreateNote(ctx context.Context, req *CreateNoteRequest) (res *CreateNoteResponse, err error)
	MGetNote(ctx context.Context, req *MGetNoteRequest) (res *MGetNoteResponse, err error)
	DeleteNote(ctx context.Context, req *DeleteNoteRequest) (res *DeleteNoteResponse, err error)
	QueryNote(ctx context.Context, req *QueryNoteRequest) (res *QueryNoteResponse, err error)
	UpdateNote(ctx context.Context, req *UpdateNoteRequest) (res *UpdateNoteResponse, err error)
}
