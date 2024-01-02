// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.19.4
// source: template/v1/meeting.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreateMeetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Avatar       string `protobuf:"bytes,2,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Description  string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Example      string `protobuf:"bytes,4,opt,name=example,proto3" json:"example,omitempty"`
	TemplateFlow string `protobuf:"bytes,5,opt,name=templateFlow,proto3" json:"templateFlow,omitempty"`
	TemplateData string `protobuf:"bytes,6,opt,name=templateData,proto3" json:"templateData,omitempty"`
	CreatedBy    string `protobuf:"bytes,7,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
}

func (x *CreateMeetingRequest) Reset() {
	*x = CreateMeetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_v1_meeting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateMeetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateMeetingRequest) ProtoMessage() {}

func (x *CreateMeetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_v1_meeting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateMeetingRequest.ProtoReflect.Descriptor instead.
func (*CreateMeetingRequest) Descriptor() ([]byte, []int) {
	return file_template_v1_meeting_proto_rawDescGZIP(), []int{0}
}

func (x *CreateMeetingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateMeetingRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *CreateMeetingRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateMeetingRequest) GetExample() string {
	if x != nil {
		return x.Example
	}
	return ""
}

func (x *CreateMeetingRequest) GetTemplateFlow() string {
	if x != nil {
		return x.TemplateFlow
	}
	return ""
}

func (x *CreateMeetingRequest) GetTemplateData() string {
	if x != nil {
		return x.TemplateData
	}
	return ""
}

func (x *CreateMeetingRequest) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

type UpdateMeetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar       string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Example      string `protobuf:"bytes,5,opt,name=example,proto3" json:"example,omitempty"`
	TemplateFlow string `protobuf:"bytes,6,opt,name=templateFlow,proto3" json:"templateFlow,omitempty"`
	TemplateData string `protobuf:"bytes,7,opt,name=templateData,proto3" json:"templateData,omitempty"`
}

func (x *UpdateMeetingRequest) Reset() {
	*x = UpdateMeetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_v1_meeting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateMeetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateMeetingRequest) ProtoMessage() {}

func (x *UpdateMeetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_v1_meeting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateMeetingRequest.ProtoReflect.Descriptor instead.
func (*UpdateMeetingRequest) Descriptor() ([]byte, []int) {
	return file_template_v1_meeting_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateMeetingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateMeetingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateMeetingRequest) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *UpdateMeetingRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateMeetingRequest) GetExample() string {
	if x != nil {
		return x.Example
	}
	return ""
}

func (x *UpdateMeetingRequest) GetTemplateFlow() string {
	if x != nil {
		return x.TemplateFlow
	}
	return ""
}

func (x *UpdateMeetingRequest) GetTemplateData() string {
	if x != nil {
		return x.TemplateData
	}
	return ""
}

type DeleteMeetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteMeetingRequest) Reset() {
	*x = DeleteMeetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_v1_meeting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteMeetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteMeetingRequest) ProtoMessage() {}

func (x *DeleteMeetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_v1_meeting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteMeetingRequest.ProtoReflect.Descriptor instead.
func (*DeleteMeetingRequest) Descriptor() ([]byte, []int) {
	return file_template_v1_meeting_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteMeetingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MeetingInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Avatar       string `protobuf:"bytes,3,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Description  string `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Example      string `protobuf:"bytes,5,opt,name=example,proto3" json:"example,omitempty"`
	TemplateFlow string `protobuf:"bytes,6,opt,name=templateFlow,proto3" json:"templateFlow,omitempty"`
	TemplateData string `protobuf:"bytes,7,opt,name=templateData,proto3" json:"templateData,omitempty"`
	StarCount    int64  `protobuf:"varint,8,opt,name=starCount,proto3" json:"starCount,omitempty"`
	CreatedBy    string `protobuf:"bytes,9,opt,name=createdBy,proto3" json:"createdBy,omitempty"`
	CreatedTime  int64  `protobuf:"varint,10,opt,name=createdTime,proto3" json:"createdTime,omitempty"`
}

func (x *MeetingInfo) Reset() {
	*x = MeetingInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_v1_meeting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MeetingInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MeetingInfo) ProtoMessage() {}

func (x *MeetingInfo) ProtoReflect() protoreflect.Message {
	mi := &file_template_v1_meeting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MeetingInfo.ProtoReflect.Descriptor instead.
func (*MeetingInfo) Descriptor() ([]byte, []int) {
	return file_template_v1_meeting_proto_rawDescGZIP(), []int{3}
}

func (x *MeetingInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MeetingInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MeetingInfo) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *MeetingInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *MeetingInfo) GetExample() string {
	if x != nil {
		return x.Example
	}
	return ""
}

func (x *MeetingInfo) GetTemplateFlow() string {
	if x != nil {
		return x.TemplateFlow
	}
	return ""
}

func (x *MeetingInfo) GetTemplateData() string {
	if x != nil {
		return x.TemplateData
	}
	return ""
}

func (x *MeetingInfo) GetStarCount() int64 {
	if x != nil {
		return x.StarCount
	}
	return 0
}

func (x *MeetingInfo) GetCreatedBy() string {
	if x != nil {
		return x.CreatedBy
	}
	return ""
}

func (x *MeetingInfo) GetCreatedTime() int64 {
	if x != nil {
		return x.CreatedTime
	}
	return 0
}

type GetMeetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetMeetingRequest) Reset() {
	*x = GetMeetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_v1_meeting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeetingRequest) ProtoMessage() {}

func (x *GetMeetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_v1_meeting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeetingRequest.ProtoReflect.Descriptor instead.
func (*GetMeetingRequest) Descriptor() ([]byte, []int) {
	return file_template_v1_meeting_proto_rawDescGZIP(), []int{4}
}

func (x *GetMeetingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetMeetingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int64        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *MeetingInfo `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetMeetingReply) Reset() {
	*x = GetMeetingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_v1_meeting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMeetingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMeetingReply) ProtoMessage() {}

func (x *GetMeetingReply) ProtoReflect() protoreflect.Message {
	mi := &file_template_v1_meeting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMeetingReply.ProtoReflect.Descriptor instead.
func (*GetMeetingReply) Descriptor() ([]byte, []int) {
	return file_template_v1_meeting_proto_rawDescGZIP(), []int{5}
}

func (x *GetMeetingReply) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetMeetingReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *GetMeetingReply) GetData() *MeetingInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListMeetingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PageNum  int64  `protobuf:"varint,3,opt,name=pageNum,proto3" json:"pageNum,omitempty"`
	PageSize int64  `protobuf:"varint,4,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *ListMeetingRequest) Reset() {
	*x = ListMeetingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_v1_meeting_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMeetingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMeetingRequest) ProtoMessage() {}

func (x *ListMeetingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_template_v1_meeting_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMeetingRequest.ProtoReflect.Descriptor instead.
func (*ListMeetingRequest) Descriptor() ([]byte, []int) {
	return file_template_v1_meeting_proto_rawDescGZIP(), []int{6}
}

func (x *ListMeetingRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ListMeetingRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ListMeetingRequest) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListMeetingRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListMeetingReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32                  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string                 `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *ListMeetingReply_Data `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ListMeetingReply) Reset() {
	*x = ListMeetingReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_v1_meeting_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMeetingReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMeetingReply) ProtoMessage() {}

func (x *ListMeetingReply) ProtoReflect() protoreflect.Message {
	mi := &file_template_v1_meeting_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMeetingReply.ProtoReflect.Descriptor instead.
func (*ListMeetingReply) Descriptor() ([]byte, []int) {
	return file_template_v1_meeting_proto_rawDescGZIP(), []int{7}
}

func (x *ListMeetingReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListMeetingReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ListMeetingReply) GetData() *ListMeetingReply_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type ListMeetingReply_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64          `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*MeetingInfo `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *ListMeetingReply_Data) Reset() {
	*x = ListMeetingReply_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_v1_meeting_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMeetingReply_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMeetingReply_Data) ProtoMessage() {}

func (x *ListMeetingReply_Data) ProtoReflect() protoreflect.Message {
	mi := &file_template_v1_meeting_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMeetingReply_Data.ProtoReflect.Descriptor instead.
func (*ListMeetingReply_Data) Descriptor() ([]byte, []int) {
	return file_template_v1_meeting_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ListMeetingReply_Data) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListMeetingReply_Data) GetData() []*MeetingInfo {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_template_v1_meeting_proto protoreflect.FileDescriptor

var file_template_v1_meeting_proto_rawDesc = []byte{
	0x0a, 0x19, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe4, 0x01, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x46, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x22, 0xd6, 0x01, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74,
	0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x77,
	0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xab, 0x02, 0x0a,
	0x0b, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x46, 0x6c, 0x6f, 0x77, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x46, 0x6c, 0x6f, 0x77, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x42, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x69, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6e, 0x0a, 0x12, 0x4c, 0x69,
	0x73, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xc4, 0x01, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x3a, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x4e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12,
	0x30, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e,
	0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xd6, 0x04, 0x0a, 0x07, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x74, 0x0a,
	0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x25,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x6d, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x74, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x65,
	0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x20, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a,
	0x01, 0x2a, 0x1a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74,
	0x65, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x71, 0x0a, 0x0d, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x25, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x17, 0x2a, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x76, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x22, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x22, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1c, 0x12, 0x1a, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2f,
	0x7b, 0x69, 0x64, 0x7d, 0x12, 0x74, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x65, 0x74,
	0x69, 0x6e, 0x67, 0x12, 0x23, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x65, 0x74, 0x69, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d,
	0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1d, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x17, 0x12, 0x15, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61,
	0x74, 0x65, 0x2f, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x39, 0x0a, 0x0f, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x24, 0x61, 0x69, 0x2d, 0x6d, 0x65, 0x65, 0x74, 0x69, 0x6e, 0x67, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_template_v1_meeting_proto_rawDescOnce sync.Once
	file_template_v1_meeting_proto_rawDescData = file_template_v1_meeting_proto_rawDesc
)

func file_template_v1_meeting_proto_rawDescGZIP() []byte {
	file_template_v1_meeting_proto_rawDescOnce.Do(func() {
		file_template_v1_meeting_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_v1_meeting_proto_rawDescData)
	})
	return file_template_v1_meeting_proto_rawDescData
}

var file_template_v1_meeting_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_template_v1_meeting_proto_goTypes = []interface{}{
	(*CreateMeetingRequest)(nil),  // 0: api.template.v1.CreateMeetingRequest
	(*UpdateMeetingRequest)(nil),  // 1: api.template.v1.UpdateMeetingRequest
	(*DeleteMeetingRequest)(nil),  // 2: api.template.v1.DeleteMeetingRequest
	(*MeetingInfo)(nil),           // 3: api.template.v1.MeetingInfo
	(*GetMeetingRequest)(nil),     // 4: api.template.v1.GetMeetingRequest
	(*GetMeetingReply)(nil),       // 5: api.template.v1.GetMeetingReply
	(*ListMeetingRequest)(nil),    // 6: api.template.v1.ListMeetingRequest
	(*ListMeetingReply)(nil),      // 7: api.template.v1.ListMeetingReply
	(*ListMeetingReply_Data)(nil), // 8: api.template.v1.ListMeetingReply.Data
	(*BoolReply)(nil),             // 9: api.template.v1.BoolReply
}
var file_template_v1_meeting_proto_depIdxs = []int32{
	3, // 0: api.template.v1.GetMeetingReply.data:type_name -> api.template.v1.MeetingInfo
	8, // 1: api.template.v1.ListMeetingReply.data:type_name -> api.template.v1.ListMeetingReply.Data
	3, // 2: api.template.v1.ListMeetingReply.Data.data:type_name -> api.template.v1.MeetingInfo
	0, // 3: api.template.v1.Meeting.CreateMeeting:input_type -> api.template.v1.CreateMeetingRequest
	1, // 4: api.template.v1.Meeting.UpdateMeeting:input_type -> api.template.v1.UpdateMeetingRequest
	2, // 5: api.template.v1.Meeting.DeleteMeeting:input_type -> api.template.v1.DeleteMeetingRequest
	4, // 6: api.template.v1.Meeting.GetMeeting:input_type -> api.template.v1.GetMeetingRequest
	6, // 7: api.template.v1.Meeting.ListMeeting:input_type -> api.template.v1.ListMeetingRequest
	9, // 8: api.template.v1.Meeting.CreateMeeting:output_type -> api.template.v1.BoolReply
	9, // 9: api.template.v1.Meeting.UpdateMeeting:output_type -> api.template.v1.BoolReply
	9, // 10: api.template.v1.Meeting.DeleteMeeting:output_type -> api.template.v1.BoolReply
	5, // 11: api.template.v1.Meeting.GetMeeting:output_type -> api.template.v1.GetMeetingReply
	7, // 12: api.template.v1.Meeting.ListMeeting:output_type -> api.template.v1.ListMeetingReply
	8, // [8:13] is the sub-list for method output_type
	3, // [3:8] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_template_v1_meeting_proto_init() }
func file_template_v1_meeting_proto_init() {
	if File_template_v1_meeting_proto != nil {
		return
	}
	file_template_v1_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_template_v1_meeting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateMeetingRequest); i {
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
		file_template_v1_meeting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateMeetingRequest); i {
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
		file_template_v1_meeting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteMeetingRequest); i {
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
		file_template_v1_meeting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MeetingInfo); i {
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
		file_template_v1_meeting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeetingRequest); i {
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
		file_template_v1_meeting_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetMeetingReply); i {
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
		file_template_v1_meeting_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMeetingRequest); i {
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
		file_template_v1_meeting_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMeetingReply); i {
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
		file_template_v1_meeting_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMeetingReply_Data); i {
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
			RawDescriptor: file_template_v1_meeting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_template_v1_meeting_proto_goTypes,
		DependencyIndexes: file_template_v1_meeting_proto_depIdxs,
		MessageInfos:      file_template_v1_meeting_proto_msgTypes,
	}.Build()
	File_template_v1_meeting_proto = out.File
	file_template_v1_meeting_proto_rawDesc = nil
	file_template_v1_meeting_proto_goTypes = nil
	file_template_v1_meeting_proto_depIdxs = nil
}
