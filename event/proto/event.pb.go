// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: proto/event.proto

package event

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminId            uint64  `protobuf:"varint,1,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
	ApplyRules         string  `protobuf:"bytes,2,opt,name=apply_rules,json=applyRules,proto3" json:"apply_rules,omitempty"`
	ProductQuantity    uint32  `protobuf:"varint,3,opt,name=product_quantity,json=productQuantity,proto3" json:"product_quantity,omitempty"`
	ProductItemPrice   float64 `protobuf:"fixed64,4,opt,name=product_item_price,json=productItemPrice,proto3" json:"product_item_price,omitempty"`
	Info               string  `protobuf:"bytes,5,opt,name=info,proto3" json:"info,omitempty"`
	ReviewUpperLimitAt string  `protobuf:"bytes,6,opt,name=review_upper_limit_at,json=reviewUpperLimitAt,proto3" json:"review_upper_limit_at,omitempty"`
	StartingAt         string  `protobuf:"bytes,7,opt,name=starting_at,json=startingAt,proto3" json:"starting_at,omitempty"`
	OverDueMaxTimes    uint64  `protobuf:"varint,8,opt,name=OverDueMaxTimes,proto3" json:"OverDueMaxTimes,omitempty"`
	Id                 uint64  `protobuf:"varint,9,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateEventRequest) Reset() {
	*x = CreateEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventRequest) ProtoMessage() {}

func (x *CreateEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventRequest.ProtoReflect.Descriptor instead.
func (*CreateEventRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{0}
}

func (x *CreateEventRequest) GetAdminId() uint64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

func (x *CreateEventRequest) GetApplyRules() string {
	if x != nil {
		return x.ApplyRules
	}
	return ""
}

func (x *CreateEventRequest) GetProductQuantity() uint32 {
	if x != nil {
		return x.ProductQuantity
	}
	return 0
}

func (x *CreateEventRequest) GetProductItemPrice() float64 {
	if x != nil {
		return x.ProductItemPrice
	}
	return 0
}

func (x *CreateEventRequest) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

func (x *CreateEventRequest) GetReviewUpperLimitAt() string {
	if x != nil {
		return x.ReviewUpperLimitAt
	}
	return ""
}

func (x *CreateEventRequest) GetStartingAt() string {
	if x != nil {
		return x.StartingAt
	}
	return ""
}

func (x *CreateEventRequest) GetOverDueMaxTimes() uint64 {
	if x != nil {
		return x.OverDueMaxTimes
	}
	return 0
}

func (x *CreateEventRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type CreateEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid      string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	IsSuccess bool   `protobuf:"varint,2,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
}

func (x *CreateEventResponse) Reset() {
	*x = CreateEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateEventResponse) ProtoMessage() {}

func (x *CreateEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateEventResponse.ProtoReflect.Descriptor instead.
func (*CreateEventResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{1}
}

func (x *CreateEventResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CreateEventResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

type EventItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UuidUrl string `protobuf:"bytes,1,opt,name=uuid_url,json=uuidUrl,proto3" json:"uuid_url,omitempty"`
	Info    string `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *EventItem) Reset() {
	*x = EventItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventItem) ProtoMessage() {}

func (x *EventItem) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventItem.ProtoReflect.Descriptor instead.
func (*EventItem) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{2}
}

func (x *EventItem) GetUuidUrl() string {
	if x != nil {
		return x.UuidUrl
	}
	return ""
}

func (x *EventItem) GetInfo() string {
	if x != nil {
		return x.Info
	}
	return ""
}

type FetchEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FetchEventsRequest) Reset() {
	*x = FetchEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchEventsRequest) ProtoMessage() {}

func (x *FetchEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchEventsRequest.ProtoReflect.Descriptor instead.
func (*FetchEventsRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{3}
}

type FetchEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*EventItem `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *FetchEventsResponse) Reset() {
	*x = FetchEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FetchEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FetchEventsResponse) ProtoMessage() {}

func (x *FetchEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FetchEventsResponse.ProtoReflect.Descriptor instead.
func (*FetchEventsResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{4}
}

func (x *FetchEventsResponse) GetEvents() []*EventItem {
	if x != nil {
		return x.Events
	}
	return nil
}

type CheckBlacklistRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Uuid   string `protobuf:"bytes,2,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *CheckBlacklistRequest) Reset() {
	*x = CheckBlacklistRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckBlacklistRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBlacklistRequest) ProtoMessage() {}

func (x *CheckBlacklistRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBlacklistRequest.ProtoReflect.Descriptor instead.
func (*CheckBlacklistRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{5}
}

func (x *CheckBlacklistRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CheckBlacklistRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CheckBlacklistResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link      string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	Situation int64  `protobuf:"varint,2,opt,name=situation,proto3" json:"situation,omitempty"`
}

func (x *CheckBlacklistResponse) Reset() {
	*x = CheckBlacklistResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckBlacklistResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckBlacklistResponse) ProtoMessage() {}

func (x *CheckBlacklistResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckBlacklistResponse.ProtoReflect.Descriptor instead.
func (*CheckBlacklistResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{6}
}

func (x *CheckBlacklistResponse) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *CheckBlacklistResponse) GetSituation() int64 {
	if x != nil {
		return x.Situation
	}
	return 0
}

type ApplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Link   string `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	UserId uint64 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Uuid   string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *ApplyRequest) Reset() {
	*x = ApplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyRequest) ProtoMessage() {}

func (x *ApplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyRequest.ProtoReflect.Descriptor instead.
func (*ApplyRequest) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{7}
}

func (x *ApplyRequest) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

func (x *ApplyRequest) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *ApplyRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type ApplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSuccess bool  `protobuf:"varint,1,opt,name=is_success,json=isSuccess,proto3" json:"is_success,omitempty"`
	Situation int64 `protobuf:"varint,2,opt,name=situation,proto3" json:"situation,omitempty"`
}

func (x *ApplyResponse) Reset() {
	*x = ApplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_event_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApplyResponse) ProtoMessage() {}

func (x *ApplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_event_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApplyResponse.ProtoReflect.Descriptor instead.
func (*ApplyResponse) Descriptor() ([]byte, []int) {
	return file_proto_event_proto_rawDescGZIP(), []int{8}
}

func (x *ApplyResponse) GetIsSuccess() bool {
	if x != nil {
		return x.IsSuccess
	}
	return false
}

func (x *ApplyResponse) GetSituation() int64 {
	if x != nil {
		return x.Situation
	}
	return 0
}

var File_proto_event_proto protoreflect.FileDescriptor

var file_proto_event_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x22, 0xcb, 0x02, 0x0a, 0x12, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x29, 0x0a,
	0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x5f, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x10, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x74, 0x65,
	0x6d, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x31, 0x0a, 0x15, 0x72, 0x65,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x75, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x72, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x55, 0x70, 0x70, 0x65, 0x72, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x41, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x41, 0x74, 0x12, 0x28,
	0x0a, 0x0f, 0x4f, 0x76, 0x65, 0x72, 0x44, 0x75, 0x65, 0x4d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0f, 0x4f, 0x76, 0x65, 0x72, 0x44, 0x75, 0x65,
	0x4d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x48, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0x3a, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x19, 0x0a, 0x08, 0x75, 0x75, 0x69, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x75, 0x75, 0x69, 0x64, 0x55, 0x72, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x14,
	0x0a, 0x12, 0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x3f, 0x0a, 0x13, 0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x06, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x44, 0x0a, 0x15, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x4a, 0x0a, 0x16, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x74,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69,
	0x74, 0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x4f, 0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x4c, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x74, 0x75,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x69, 0x74,
	0x75, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x86, 0x03, 0x0a, 0x0b, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x45, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x46,
	0x0a, 0x0b, 0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x19, 0x2e,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x46, 0x65, 0x74, 0x63, 0x68, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42,
	0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x17, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x44, 0x6f, 0x77, 0x6e, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x1c, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x42, 0x6c, 0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x6c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3c, 0x0a, 0x0d, 0x41, 0x70, 0x70, 0x6c, 0x79, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61,
	0x73, 0x65, 0x12, 0x13, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x41, 0x70, 0x70, 0x6c, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42,
	0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_event_proto_rawDescOnce sync.Once
	file_proto_event_proto_rawDescData = file_proto_event_proto_rawDesc
)

func file_proto_event_proto_rawDescGZIP() []byte {
	file_proto_event_proto_rawDescOnce.Do(func() {
		file_proto_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_event_proto_rawDescData)
	})
	return file_proto_event_proto_rawDescData
}

var file_proto_event_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_event_proto_goTypes = []interface{}{
	(*CreateEventRequest)(nil),     // 0: event.CreateEventRequest
	(*CreateEventResponse)(nil),    // 1: event.CreateEventResponse
	(*EventItem)(nil),              // 2: event.EventItem
	(*FetchEventsRequest)(nil),     // 3: event.FetchEventsRequest
	(*FetchEventsResponse)(nil),    // 4: event.FetchEventsResponse
	(*CheckBlacklistRequest)(nil),  // 5: event.CheckBlacklistRequest
	(*CheckBlacklistResponse)(nil), // 6: event.CheckBlacklistResponse
	(*ApplyRequest)(nil),           // 7: event.ApplyRequest
	(*ApplyResponse)(nil),          // 8: event.ApplyResponse
}
var file_proto_event_proto_depIdxs = []int32{
	2, // 0: event.FetchEventsResponse.events:type_name -> event.EventItem
	0, // 1: event.EventEngine.CreateEvent:input_type -> event.CreateEventRequest
	3, // 2: event.EventEngine.FetchEvents:input_type -> event.FetchEventsRequest
	5, // 3: event.EventEngine.CheckBlacklist:input_type -> event.CheckBlacklistRequest
	5, // 4: event.EventEngine.CheckBlacklistDowngrade:input_type -> event.CheckBlacklistRequest
	7, // 5: event.EventEngine.ApplyPurchase:input_type -> event.ApplyRequest
	1, // 6: event.EventEngine.CreateEvent:output_type -> event.CreateEventResponse
	4, // 7: event.EventEngine.FetchEvents:output_type -> event.FetchEventsResponse
	6, // 8: event.EventEngine.CheckBlacklist:output_type -> event.CheckBlacklistResponse
	6, // 9: event.EventEngine.CheckBlacklistDowngrade:output_type -> event.CheckBlacklistResponse
	8, // 10: event.EventEngine.ApplyPurchase:output_type -> event.ApplyResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_event_proto_init() }
func file_proto_event_proto_init() {
	if File_proto_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventRequest); i {
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
		file_proto_event_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateEventResponse); i {
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
		file_proto_event_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventItem); i {
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
		file_proto_event_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchEventsRequest); i {
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
		file_proto_event_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FetchEventsResponse); i {
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
		file_proto_event_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckBlacklistRequest); i {
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
		file_proto_event_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckBlacklistResponse); i {
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
		file_proto_event_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyRequest); i {
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
		file_proto_event_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApplyResponse); i {
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
			RawDescriptor: file_proto_event_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_event_proto_goTypes,
		DependencyIndexes: file_proto_event_proto_depIdxs,
		MessageInfos:      file_proto_event_proto_msgTypes,
	}.Build()
	File_proto_event_proto = out.File
	file_proto_event_proto_rawDesc = nil
	file_proto_event_proto_goTypes = nil
	file_proto_event_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// EventEngineClient is the client API for EventEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventEngineClient interface {
	CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error)
	FetchEvents(ctx context.Context, in *FetchEventsRequest, opts ...grpc.CallOption) (*FetchEventsResponse, error)
	CheckBlacklist(ctx context.Context, in *CheckBlacklistRequest, opts ...grpc.CallOption) (*CheckBlacklistResponse, error)
	CheckBlacklistDowngrade(ctx context.Context, in *CheckBlacklistRequest, opts ...grpc.CallOption) (*CheckBlacklistResponse, error)
	ApplyPurchase(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error)
}

type eventEngineClient struct {
	cc grpc.ClientConnInterface
}

func NewEventEngineClient(cc grpc.ClientConnInterface) EventEngineClient {
	return &eventEngineClient{cc}
}

func (c *eventEngineClient) CreateEvent(ctx context.Context, in *CreateEventRequest, opts ...grpc.CallOption) (*CreateEventResponse, error) {
	out := new(CreateEventResponse)
	err := c.cc.Invoke(ctx, "/event.EventEngine/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventEngineClient) FetchEvents(ctx context.Context, in *FetchEventsRequest, opts ...grpc.CallOption) (*FetchEventsResponse, error) {
	out := new(FetchEventsResponse)
	err := c.cc.Invoke(ctx, "/event.EventEngine/FetchEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventEngineClient) CheckBlacklist(ctx context.Context, in *CheckBlacklistRequest, opts ...grpc.CallOption) (*CheckBlacklistResponse, error) {
	out := new(CheckBlacklistResponse)
	err := c.cc.Invoke(ctx, "/event.EventEngine/CheckBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventEngineClient) CheckBlacklistDowngrade(ctx context.Context, in *CheckBlacklistRequest, opts ...grpc.CallOption) (*CheckBlacklistResponse, error) {
	out := new(CheckBlacklistResponse)
	err := c.cc.Invoke(ctx, "/event.EventEngine/CheckBlacklistDowngrade", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *eventEngineClient) ApplyPurchase(ctx context.Context, in *ApplyRequest, opts ...grpc.CallOption) (*ApplyResponse, error) {
	out := new(ApplyResponse)
	err := c.cc.Invoke(ctx, "/event.EventEngine/ApplyPurchase", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventEngineServer is the server API for EventEngine service.
type EventEngineServer interface {
	CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error)
	FetchEvents(context.Context, *FetchEventsRequest) (*FetchEventsResponse, error)
	CheckBlacklist(context.Context, *CheckBlacklistRequest) (*CheckBlacklistResponse, error)
	CheckBlacklistDowngrade(context.Context, *CheckBlacklistRequest) (*CheckBlacklistResponse, error)
	ApplyPurchase(context.Context, *ApplyRequest) (*ApplyResponse, error)
}

// UnimplementedEventEngineServer can be embedded to have forward compatible implementations.
type UnimplementedEventEngineServer struct {
}

func (*UnimplementedEventEngineServer) CreateEvent(context.Context, *CreateEventRequest) (*CreateEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}
func (*UnimplementedEventEngineServer) FetchEvents(context.Context, *FetchEventsRequest) (*FetchEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchEvents not implemented")
}
func (*UnimplementedEventEngineServer) CheckBlacklist(context.Context, *CheckBlacklistRequest) (*CheckBlacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBlacklist not implemented")
}
func (*UnimplementedEventEngineServer) CheckBlacklistDowngrade(context.Context, *CheckBlacklistRequest) (*CheckBlacklistResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckBlacklistDowngrade not implemented")
}
func (*UnimplementedEventEngineServer) ApplyPurchase(context.Context, *ApplyRequest) (*ApplyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyPurchase not implemented")
}

func RegisterEventEngineServer(s *grpc.Server, srv EventEngineServer) {
	s.RegisterService(&_EventEngine_serviceDesc, srv)
}

func _EventEngine_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventEngineServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventEngine/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventEngineServer).CreateEvent(ctx, req.(*CreateEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventEngine_FetchEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventEngineServer).FetchEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventEngine/FetchEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventEngineServer).FetchEvents(ctx, req.(*FetchEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventEngine_CheckBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckBlacklistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventEngineServer).CheckBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventEngine/CheckBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventEngineServer).CheckBlacklist(ctx, req.(*CheckBlacklistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventEngine_CheckBlacklistDowngrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckBlacklistRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventEngineServer).CheckBlacklistDowngrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventEngine/CheckBlacklistDowngrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventEngineServer).CheckBlacklistDowngrade(ctx, req.(*CheckBlacklistRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EventEngine_ApplyPurchase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ApplyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventEngineServer).ApplyPurchase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/event.EventEngine/ApplyPurchase",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventEngineServer).ApplyPurchase(ctx, req.(*ApplyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventEngine_serviceDesc = grpc.ServiceDesc{
	ServiceName: "event.EventEngine",
	HandlerType: (*EventEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _EventEngine_CreateEvent_Handler,
		},
		{
			MethodName: "FetchEvents",
			Handler:    _EventEngine_FetchEvents_Handler,
		},
		{
			MethodName: "CheckBlacklist",
			Handler:    _EventEngine_CheckBlacklist_Handler,
		},
		{
			MethodName: "CheckBlacklistDowngrade",
			Handler:    _EventEngine_CheckBlacklistDowngrade_Handler,
		},
		{
			MethodName: "ApplyPurchase",
			Handler:    _EventEngine_ApplyPurchase_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/event.proto",
}
