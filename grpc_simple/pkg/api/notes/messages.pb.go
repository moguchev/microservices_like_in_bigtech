// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.26.1
// source: api/notes/messages.proto

package notes

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

// NoteInfo - информация записи
type NoteInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// title - название записи
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	// content - содержимое записи
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *NoteInfo) Reset() {
	*x = NoteInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notes_messages_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoteInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoteInfo) ProtoMessage() {}

func (x *NoteInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoteInfo.ProtoReflect.Descriptor instead.
func (*NoteInfo) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{0}
}

func (x *NoteInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *NoteInfo) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

// Note - full note model
type Note struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id - уникальный идентификатор записи
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// info -
	Info *NoteInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *Note) Reset() {
	*x = Note{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notes_messages_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Note) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Note) ProtoMessage() {}

func (x *Note) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[1]
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
	return file_api_notes_messages_proto_rawDescGZIP(), []int{1}
}

func (x *Note) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Note) GetInfo() *NoteInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// SaveNoteRequest - запрос SaveNote
type SaveNoteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// info - информация записи (контент)
	Info *NoteInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *SaveNoteRequest) Reset() {
	*x = SaveNoteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notes_messages_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveNoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveNoteRequest) ProtoMessage() {}

func (x *SaveNoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveNoteRequest.ProtoReflect.Descriptor instead.
func (*SaveNoteRequest) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{2}
}

func (x *SaveNoteRequest) GetInfo() *NoteInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

// SaveNoteResponse - ответ SaveNote
type SaveNoteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// id - уникальный идентификатор записи
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *SaveNoteResponse) Reset() {
	*x = SaveNoteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notes_messages_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SaveNoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SaveNoteResponse) ProtoMessage() {}

func (x *SaveNoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SaveNoteResponse.ProtoReflect.Descriptor instead.
func (*SaveNoteResponse) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{3}
}

func (x *SaveNoteResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// ListNotesRequest - запрос ListNotes
type ListNotesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListNotesRequest) Reset() {
	*x = ListNotesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notes_messages_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesRequest) ProtoMessage() {}

func (x *ListNotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesRequest.ProtoReflect.Descriptor instead.
func (*ListNotesRequest) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{4}
}

// ListNotesResponse - ответ ListNotes
type ListNotesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// notes - все записи
	Notes []*Note `protobuf:"bytes,1,rep,name=notes,proto3" json:"notes,omitempty"`
	// Types that are assignable to Meta:
	//
	//	*ListNotesResponse_Foo_
	//	*ListNotesResponse_Bar_
	Meta isListNotesResponse_Meta `protobuf_oneof:"meta"`
}

func (x *ListNotesResponse) Reset() {
	*x = ListNotesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notes_messages_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesResponse) ProtoMessage() {}

func (x *ListNotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesResponse.ProtoReflect.Descriptor instead.
func (*ListNotesResponse) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{5}
}

func (x *ListNotesResponse) GetNotes() []*Note {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (m *ListNotesResponse) GetMeta() isListNotesResponse_Meta {
	if m != nil {
		return m.Meta
	}
	return nil
}

func (x *ListNotesResponse) GetFoo() *ListNotesResponse_Foo {
	if x, ok := x.GetMeta().(*ListNotesResponse_Foo_); ok {
		return x.Foo
	}
	return nil
}

func (x *ListNotesResponse) GetBar() *ListNotesResponse_Bar {
	if x, ok := x.GetMeta().(*ListNotesResponse_Bar_); ok {
		return x.Bar
	}
	return nil
}

type isListNotesResponse_Meta interface {
	isListNotesResponse_Meta()
}

type ListNotesResponse_Foo_ struct {
	Foo *ListNotesResponse_Foo `protobuf:"bytes,2,opt,name=foo,proto3,oneof"`
}

type ListNotesResponse_Bar_ struct {
	Bar *ListNotesResponse_Bar `protobuf:"bytes,3,opt,name=bar,proto3,oneof"`
}

func (*ListNotesResponse_Foo_) isListNotesResponse_Meta() {}

func (*ListNotesResponse_Bar_) isListNotesResponse_Meta() {}

type ListNotesResponse_Foo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int int64 `protobuf:"varint,1,opt,name=int,proto3" json:"int,omitempty"`
}

func (x *ListNotesResponse_Foo) Reset() {
	*x = ListNotesResponse_Foo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notes_messages_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotesResponse_Foo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesResponse_Foo) ProtoMessage() {}

func (x *ListNotesResponse_Foo) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesResponse_Foo.ProtoReflect.Descriptor instead.
func (*ListNotesResponse_Foo) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{5, 0}
}

func (x *ListNotesResponse_Foo) GetInt() int64 {
	if x != nil {
		return x.Int
	}
	return 0
}

type ListNotesResponse_Bar struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int int32 `protobuf:"varint,1,opt,name=int,proto3" json:"int,omitempty"`
}

func (x *ListNotesResponse_Bar) Reset() {
	*x = ListNotesResponse_Bar{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_notes_messages_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNotesResponse_Bar) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNotesResponse_Bar) ProtoMessage() {}

func (x *ListNotesResponse_Bar) ProtoReflect() protoreflect.Message {
	mi := &file_api_notes_messages_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNotesResponse_Bar.ProtoReflect.Descriptor instead.
func (*ListNotesResponse_Bar) Descriptor() ([]byte, []int) {
	return file_api_notes_messages_proto_rawDescGZIP(), []int{5, 1}
}

func (x *ListNotesResponse_Bar) GetInt() int32 {
	if x != nil {
		return x.Int
	}
	return 0
}

var File_api_notes_messages_proto protoreflect.FileDescriptor

var file_api_notes_messages_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x67, 0x75, 0x63, 0x68, 0x65, 0x76, 0x2e,
	0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x22, 0x3a, 0x0a, 0x08, 0x4e, 0x6f, 0x74,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x63, 0x0a, 0x04, 0x4e, 0x6f, 0x74, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x4b, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x67, 0x75, 0x63, 0x68, 0x65,
	0x76, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x5e, 0x0a, 0x0f, 0x53, 0x61,
	0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a,
	0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x67, 0x75, 0x63, 0x68, 0x65,
	0x76, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x22, 0x0a, 0x10, 0x53, 0x61,
	0x76, 0x65, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x12,
	0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0xcc, 0x02, 0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x49, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x67, 0x75, 0x63, 0x68, 0x65, 0x76, 0x2e, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4e, 0x6f, 0x74, 0x65, 0x52, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x12, 0x58, 0x0a, 0x03, 0x66, 0x6f, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x44, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f,
	0x67, 0x75, 0x63, 0x68, 0x65, 0x76, 0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x46, 0x6f, 0x6f, 0x48, 0x00, 0x52, 0x03, 0x66, 0x6f, 0x6f, 0x12, 0x58, 0x0a,
	0x03, 0x62, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x44, 0x2e, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x6f, 0x67, 0x75, 0x63, 0x68, 0x65, 0x76,
	0x2e, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e,
	0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x61, 0x72,
	0x48, 0x00, 0x52, 0x03, 0x62, 0x61, 0x72, 0x1a, 0x17, 0x0a, 0x03, 0x46, 0x6f, 0x6f, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x69, 0x6e, 0x74,
	0x1a, 0x17, 0x0a, 0x03, 0x42, 0x61, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6e, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x69, 0x6e, 0x74, 0x42, 0x06, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x42, 0x4b, 0x5a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6d, 0x6f, 0x67, 0x75, 0x63, 0x68, 0x65, 0x76, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x63, 0x73, 0x65, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x5f, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x3b, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_notes_messages_proto_rawDescOnce sync.Once
	file_api_notes_messages_proto_rawDescData = file_api_notes_messages_proto_rawDesc
)

func file_api_notes_messages_proto_rawDescGZIP() []byte {
	file_api_notes_messages_proto_rawDescOnce.Do(func() {
		file_api_notes_messages_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_notes_messages_proto_rawDescData)
	})
	return file_api_notes_messages_proto_rawDescData
}

var file_api_notes_messages_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_notes_messages_proto_goTypes = []any{
	(*NoteInfo)(nil),              // 0: github.com.moguchev.microservices.grpc_simple.NoteInfo
	(*Note)(nil),                  // 1: github.com.moguchev.microservices.grpc_simple.Note
	(*SaveNoteRequest)(nil),       // 2: github.com.moguchev.microservices.grpc_simple.SaveNoteRequest
	(*SaveNoteResponse)(nil),      // 3: github.com.moguchev.microservices.grpc_simple.SaveNoteResponse
	(*ListNotesRequest)(nil),      // 4: github.com.moguchev.microservices.grpc_simple.ListNotesRequest
	(*ListNotesResponse)(nil),     // 5: github.com.moguchev.microservices.grpc_simple.ListNotesResponse
	(*ListNotesResponse_Foo)(nil), // 6: github.com.moguchev.microservices.grpc_simple.ListNotesResponse.Foo
	(*ListNotesResponse_Bar)(nil), // 7: github.com.moguchev.microservices.grpc_simple.ListNotesResponse.Bar
}
var file_api_notes_messages_proto_depIdxs = []int32{
	0, // 0: github.com.moguchev.microservices.grpc_simple.Note.info:type_name -> github.com.moguchev.microservices.grpc_simple.NoteInfo
	0, // 1: github.com.moguchev.microservices.grpc_simple.SaveNoteRequest.info:type_name -> github.com.moguchev.microservices.grpc_simple.NoteInfo
	1, // 2: github.com.moguchev.microservices.grpc_simple.ListNotesResponse.notes:type_name -> github.com.moguchev.microservices.grpc_simple.Note
	6, // 3: github.com.moguchev.microservices.grpc_simple.ListNotesResponse.foo:type_name -> github.com.moguchev.microservices.grpc_simple.ListNotesResponse.Foo
	7, // 4: github.com.moguchev.microservices.grpc_simple.ListNotesResponse.bar:type_name -> github.com.moguchev.microservices.grpc_simple.ListNotesResponse.Bar
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_api_notes_messages_proto_init() }
func file_api_notes_messages_proto_init() {
	if File_api_notes_messages_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_notes_messages_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*NoteInfo); i {
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
		file_api_notes_messages_proto_msgTypes[1].Exporter = func(v any, i int) any {
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
		file_api_notes_messages_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SaveNoteRequest); i {
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
		file_api_notes_messages_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*SaveNoteResponse); i {
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
		file_api_notes_messages_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListNotesRequest); i {
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
		file_api_notes_messages_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ListNotesResponse); i {
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
		file_api_notes_messages_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*ListNotesResponse_Foo); i {
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
		file_api_notes_messages_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*ListNotesResponse_Bar); i {
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
	file_api_notes_messages_proto_msgTypes[5].OneofWrappers = []any{
		(*ListNotesResponse_Foo_)(nil),
		(*ListNotesResponse_Bar_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_notes_messages_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_api_notes_messages_proto_goTypes,
		DependencyIndexes: file_api_notes_messages_proto_depIdxs,
		MessageInfos:      file_api_notes_messages_proto_msgTypes,
	}.Build()
	File_api_notes_messages_proto = out.File
	file_api_notes_messages_proto_rawDesc = nil
	file_api_notes_messages_proto_goTypes = nil
	file_api_notes_messages_proto_depIdxs = nil
}
