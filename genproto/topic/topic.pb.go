// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.1
// source: topic.proto

package topic

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

type TopicNameReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *TopicNameReq) Reset() {
	*x = TopicNameReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicNameReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicNameReq) ProtoMessage() {}

func (x *TopicNameReq) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicNameReq.ProtoReflect.Descriptor instead.
func (*TopicNameReq) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{0}
}

func (x *TopicNameReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type TopicIdResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TopicIdResp) Reset() {
	*x = TopicIdResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TopicIdResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TopicIdResp) ProtoMessage() {}

func (x *TopicIdResp) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TopicIdResp.ProtoReflect.Descriptor instead.
func (*TopicIdResp) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{1}
}

func (x *TopicIdResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type CreateTopicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectId   string `protobuf:"bytes,1,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateTopicReq) Reset() {
	*x = CreateTopicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTopicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopicReq) ProtoMessage() {}

func (x *CreateTopicReq) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopicReq.ProtoReflect.Descriptor instead.
func (*CreateTopicReq) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{2}
}

func (x *CreateTopicReq) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *CreateTopicReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateTopicReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateTopicResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CreatedAt string `protobuf:"bytes,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *CreateTopicResp) Reset() {
	*x = CreateTopicResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTopicResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTopicResp) ProtoMessage() {}

func (x *CreateTopicResp) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTopicResp.ProtoReflect.Descriptor instead.
func (*CreateTopicResp) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{3}
}

func (x *CreateTopicResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CreateTopicResp) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type UpdateTopicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubjectId   string `protobuf:"bytes,2,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *UpdateTopicReq) Reset() {
	*x = UpdateTopicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTopicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTopicReq) ProtoMessage() {}

func (x *UpdateTopicReq) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTopicReq.ProtoReflect.Descriptor instead.
func (*UpdateTopicReq) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateTopicReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTopicReq) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *UpdateTopicReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateTopicReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UpdateTopicResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UpdatedAt string `protobuf:"bytes,2,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UpdateTopicResp) Reset() {
	*x = UpdateTopicResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTopicResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTopicResp) ProtoMessage() {}

func (x *UpdateTopicResp) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTopicResp.ProtoReflect.Descriptor instead.
func (*UpdateTopicResp) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTopicResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateTopicResp) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type DeleteTopicReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TopicId string `protobuf:"bytes,1,opt,name=topic_id,json=topicId,proto3" json:"topic_id,omitempty"`
}

func (x *DeleteTopicReq) Reset() {
	*x = DeleteTopicReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTopicReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTopicReq) ProtoMessage() {}

func (x *DeleteTopicReq) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTopicReq.ProtoReflect.Descriptor instead.
func (*DeleteTopicReq) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTopicReq) GetTopicId() string {
	if x != nil {
		return x.TopicId
	}
	return ""
}

type DeleteTopicResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteTopicResp) Reset() {
	*x = DeleteTopicResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTopicResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTopicResp) ProtoMessage() {}

func (x *DeleteTopicResp) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTopicResp.ProtoReflect.Descriptor instead.
func (*DeleteTopicResp) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteTopicResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type GetAllFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectId string `protobuf:"bytes,1,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
}

func (x *GetAllFilter) Reset() {
	*x = GetAllFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllFilter) ProtoMessage() {}

func (x *GetAllFilter) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllFilter.ProtoReflect.Descriptor instead.
func (*GetAllFilter) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllFilter) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

type GetAllTopicsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubjectId string `protobuf:"bytes,1,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	Limit     int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Page      int32  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
}

func (x *GetAllTopicsReq) Reset() {
	*x = GetAllTopicsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTopicsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTopicsReq) ProtoMessage() {}

func (x *GetAllTopicsReq) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTopicsReq.ProtoReflect.Descriptor instead.
func (*GetAllTopicsReq) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{9}
}

func (x *GetAllTopicsReq) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *GetAllTopicsReq) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllTopicsReq) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

type Topic struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	SubjectId   string `protobuf:"bytes,2,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	Name        string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedAt   string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Topic) Reset() {
	*x = Topic{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Topic) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Topic) ProtoMessage() {}

func (x *Topic) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Topic.ProtoReflect.Descriptor instead.
func (*Topic) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{10}
}

func (x *Topic) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Topic) GetSubjectId() string {
	if x != nil {
		return x.SubjectId
	}
	return ""
}

func (x *Topic) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Topic) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Topic) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type GetAllTopicsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Topics []*Topic `protobuf:"bytes,1,rep,name=topics,proto3" json:"topics,omitempty"`
	Limit  int32    `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Page   int32    `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Count  int32    `protobuf:"varint,4,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetAllTopicsResp) Reset() {
	*x = GetAllTopicsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_topic_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTopicsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTopicsResp) ProtoMessage() {}

func (x *GetAllTopicsResp) ProtoReflect() protoreflect.Message {
	mi := &file_topic_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTopicsResp.ProtoReflect.Descriptor instead.
func (*GetAllTopicsResp) Descriptor() ([]byte, []int) {
	return file_topic_proto_rawDescGZIP(), []int{11}
}

func (x *GetAllTopicsResp) GetTopics() []*Topic {
	if x != nil {
		return x.Topics
	}
	return nil
}

func (x *GetAllTopicsResp) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetAllTopicsResp) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllTopicsResp) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_topic_proto protoreflect.FileDescriptor

var file_topic_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x22, 0x22, 0x0a, 0x0c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x54, 0x6f, 0x70, 0x69,
	0x63, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40,
	0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x22, 0x75, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x2b, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x08, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74,
	0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x2d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x69, 0x6c, 0x74, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64,
	0x22, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x22, 0x8b, 0x01, 0x0a,
	0x05, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x78, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x24,
	0x0a, 0x06, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c,
	0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x06, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x32, 0xc6, 0x02, 0x0a, 0x0c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x12, 0x15, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x74, 0x6f,
	0x70, 0x69, 0x63, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x12, 0x3c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x12, 0x15, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x3c, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63,
	0x12, 0x15, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54,
	0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x3f, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x12,
	0x16, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f,
	0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x3b, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x42, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x13, 0x2e, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x2e, 0x54, 0x6f, 0x70,
	0x69, 0x63, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x2e, 0x54, 0x6f, 0x70, 0x69, 0x63, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70, 0x42, 0x10, 0x5a,
	0x0e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_topic_proto_rawDescOnce sync.Once
	file_topic_proto_rawDescData = file_topic_proto_rawDesc
)

func file_topic_proto_rawDescGZIP() []byte {
	file_topic_proto_rawDescOnce.Do(func() {
		file_topic_proto_rawDescData = protoimpl.X.CompressGZIP(file_topic_proto_rawDescData)
	})
	return file_topic_proto_rawDescData
}

var file_topic_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_topic_proto_goTypes = []any{
	(*TopicNameReq)(nil),     // 0: topic.TopicNameReq
	(*TopicIdResp)(nil),      // 1: topic.TopicIdResp
	(*CreateTopicReq)(nil),   // 2: topic.CreateTopicReq
	(*CreateTopicResp)(nil),  // 3: topic.CreateTopicResp
	(*UpdateTopicReq)(nil),   // 4: topic.UpdateTopicReq
	(*UpdateTopicResp)(nil),  // 5: topic.UpdateTopicResp
	(*DeleteTopicReq)(nil),   // 6: topic.DeleteTopicReq
	(*DeleteTopicResp)(nil),  // 7: topic.DeleteTopicResp
	(*GetAllFilter)(nil),     // 8: topic.GetAllFilter
	(*GetAllTopicsReq)(nil),  // 9: topic.GetAllTopicsReq
	(*Topic)(nil),            // 10: topic.Topic
	(*GetAllTopicsResp)(nil), // 11: topic.GetAllTopicsResp
}
var file_topic_proto_depIdxs = []int32{
	10, // 0: topic.GetAllTopicsResp.topics:type_name -> topic.Topic
	2,  // 1: topic.TopicService.CreateTopic:input_type -> topic.CreateTopicReq
	4,  // 2: topic.TopicService.UpdateTopic:input_type -> topic.UpdateTopicReq
	6,  // 3: topic.TopicService.DeleteTopic:input_type -> topic.DeleteTopicReq
	9,  // 4: topic.TopicService.GetAllTopics:input_type -> topic.GetAllTopicsReq
	0,  // 5: topic.TopicService.GetTopicIdByName:input_type -> topic.TopicNameReq
	3,  // 6: topic.TopicService.CreateTopic:output_type -> topic.CreateTopicResp
	5,  // 7: topic.TopicService.UpdateTopic:output_type -> topic.UpdateTopicResp
	7,  // 8: topic.TopicService.DeleteTopic:output_type -> topic.DeleteTopicResp
	11, // 9: topic.TopicService.GetAllTopics:output_type -> topic.GetAllTopicsResp
	1,  // 10: topic.TopicService.GetTopicIdByName:output_type -> topic.TopicIdResp
	6,  // [6:11] is the sub-list for method output_type
	1,  // [1:6] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_topic_proto_init() }
func file_topic_proto_init() {
	if File_topic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_topic_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TopicNameReq); i {
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
		file_topic_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TopicIdResp); i {
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
		file_topic_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTopicReq); i {
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
		file_topic_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTopicResp); i {
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
		file_topic_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTopicReq); i {
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
		file_topic_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTopicResp); i {
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
		file_topic_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTopicReq); i {
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
		file_topic_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTopicResp); i {
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
		file_topic_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllFilter); i {
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
		file_topic_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTopicsReq); i {
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
		file_topic_proto_msgTypes[10].Exporter = func(v any, i int) any {
			switch v := v.(*Topic); i {
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
		file_topic_proto_msgTypes[11].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTopicsResp); i {
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
			RawDescriptor: file_topic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_topic_proto_goTypes,
		DependencyIndexes: file_topic_proto_depIdxs,
		MessageInfos:      file_topic_proto_msgTypes,
	}.Build()
	File_topic_proto = out.File
	file_topic_proto_rawDesc = nil
	file_topic_proto_goTypes = nil
	file_topic_proto_depIdxs = nil
}
