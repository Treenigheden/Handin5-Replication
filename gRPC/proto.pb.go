// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.4
// source: gRPC/proto.proto

package gRPC

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

type BidInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bid  int32 `protobuf:"varint,1,opt,name=bid,proto3" json:"bid,omitempty"`
	Port int32 `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *BidInput) Reset() {
	*x = BidInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidInput) ProtoMessage() {}

func (x *BidInput) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidInput.ProtoReflect.Descriptor instead.
func (*BidInput) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_proto_rawDescGZIP(), []int{0}
}

func (x *BidInput) GetBid() int32 {
	if x != nil {
		return x.Bid
	}
	return 0
}

func (x *BidInput) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount      int32 `protobuf:"varint,1,opt,name=amount,proto3" json:"amount,omitempty"`
	AuctionOver bool  `protobuf:"varint,2,opt,name=auctionOver,proto3" json:"auctionOver,omitempty"`
	Winner      int32 `protobuf:"varint,3,opt,name=winner,proto3" json:"winner,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_proto_rawDescGZIP(), []int{1}
}

func (x *Outcome) GetAmount() int32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Outcome) GetAuctionOver() bool {
	if x != nil {
		return x.AuctionOver
	}
	return false
}

func (x *Outcome) GetWinner() int32 {
	if x != nil {
		return x.Winner
	}
	return 0
}

type Confirmation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *Confirmation) Reset() {
	*x = Confirmation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Confirmation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Confirmation) ProtoMessage() {}

func (x *Confirmation) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Confirmation.ProtoReflect.Descriptor instead.
func (*Confirmation) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_proto_rawDescGZIP(), []int{2}
}

func (x *Confirmation) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type ConnectionAnnouncement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeID    int32 `protobuf:"varint,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	Timestamp int32 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *ConnectionAnnouncement) Reset() {
	*x = ConnectionAnnouncement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConnectionAnnouncement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConnectionAnnouncement) ProtoMessage() {}

func (x *ConnectionAnnouncement) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConnectionAnnouncement.ProtoReflect.Descriptor instead.
func (*ConnectionAnnouncement) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_proto_rawDescGZIP(), []int{3}
}

func (x *ConnectionAnnouncement) GetNodeID() int32 {
	if x != nil {
		return x.NodeID
	}
	return 0
}

func (x *ConnectionAnnouncement) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type UpdateAnnouncement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HighestBid       int32 `protobuf:"varint,1,opt,name=highestBid,proto3" json:"highestBid,omitempty"`
	HighestBidder    int32 `protobuf:"varint,2,opt,name=highestBidder,proto3" json:"highestBidder,omitempty"`
	AuctionIsOngoing bool  `protobuf:"varint,3,opt,name=AuctionIsOngoing,proto3" json:"AuctionIsOngoing,omitempty"`
	Timestamp        int32 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *UpdateAnnouncement) Reset() {
	*x = UpdateAnnouncement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAnnouncement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAnnouncement) ProtoMessage() {}

func (x *UpdateAnnouncement) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAnnouncement.ProtoReflect.Descriptor instead.
func (*UpdateAnnouncement) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateAnnouncement) GetHighestBid() int32 {
	if x != nil {
		return x.HighestBid
	}
	return 0
}

func (x *UpdateAnnouncement) GetHighestBidder() int32 {
	if x != nil {
		return x.HighestBidder
	}
	return 0
}

func (x *UpdateAnnouncement) GetAuctionIsOngoing() bool {
	if x != nil {
		return x.AuctionIsOngoing
	}
	return false
}

func (x *UpdateAnnouncement) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type AccessRequestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Granted   bool  `protobuf:"varint,1,opt,name=granted,proto3" json:"granted,omitempty"`
	Timestamp int32 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *AccessRequestResponse) Reset() {
	*x = AccessRequestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequestResponse) ProtoMessage() {}

func (x *AccessRequestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequestResponse.ProtoReflect.Descriptor instead.
func (*AccessRequestResponse) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_proto_rawDescGZIP(), []int{5}
}

func (x *AccessRequestResponse) GetGranted() bool {
	if x != nil {
		return x.Granted
	}
	return false
}

func (x *AccessRequestResponse) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type AccessRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeID    int32 `protobuf:"varint,1,opt,name=nodeID,proto3" json:"nodeID,omitempty"`
	Timestamp int32 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *AccessRequest) Reset() {
	*x = AccessRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessRequest) ProtoMessage() {}

func (x *AccessRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessRequest.ProtoReflect.Descriptor instead.
func (*AccessRequest) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_proto_rawDescGZIP(), []int{6}
}

func (x *AccessRequest) GetNodeID() int32 {
	if x != nil {
		return x.NodeID
	}
	return 0
}

func (x *AccessRequest) GetTimestamp() int32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gRPC_proto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_gRPC_proto_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_gRPC_proto_proto_rawDescGZIP(), []int{7}
}

var File_gRPC_proto_proto protoreflect.FileDescriptor

var file_gRPC_proto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x09, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x35, 0x22, 0x30, 0x0a,
	0x08, 0x42, 0x69, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x62, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x62, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22,
	0x5b, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x76, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x4f, 0x76, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x77, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x22, 0x28, 0x0a, 0x0c,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4e, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0xa4, 0x01, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0a, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64, 0x64, 0x65, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x68, 0x69, 0x67, 0x68, 0x65, 0x73, 0x74, 0x42, 0x69, 0x64,
	0x64, 0x65, 0x72, 0x12, 0x2a, 0x0a, 0x10, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x73,
	0x4f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x41,
	0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x73, 0x4f, 0x6e, 0x67, 0x6f, 0x69, 0x6e, 0x67, 0x12,
	0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x4f, 0x0a,
	0x15, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x45,
	0x0a, 0x0d, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32, 0xce,
	0x03, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a,
	0x03, 0x42, 0x69, 0x64, 0x12, 0x13, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x35,
	0x2e, 0x42, 0x69, 0x64, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x35, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x2e, 0x68,
	0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x35, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12,
	0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x35, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x12, 0x50, 0x0a, 0x12, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x21, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77,
	0x6f, 0x72, 0x6b, 0x35, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x35, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x48, 0x0a, 0x0e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72,
	0x6b, 0x35, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x17, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b,
	0x35, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4f,
	0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x12, 0x18, 0x2e, 0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x35, 0x2e,
	0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x35, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x06, 0x49, 0x45, 0x78, 0x69, 0x73, 0x74, 0x12, 0x10, 0x2e, 0x68, 0x6f, 0x6d, 0x65,
	0x77, 0x6f, 0x72, 0x6b, 0x35, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x68, 0x6f,
	0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x35, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x40, 0x0a,
	0x09, 0x49, 0x41, 0x6d, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x21, 0x2e, 0x68, 0x6f, 0x6d,
	0x65, 0x77, 0x6f, 0x72, 0x6b, 0x35, 0x2e, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x10, 0x2e,
	0x68, 0x6f, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x35, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x32,
	0x0c, 0x0a, 0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4e, 0x6f, 0x64, 0x65, 0x42, 0x08, 0x5a,
	0x06, 0x2e, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gRPC_proto_proto_rawDescOnce sync.Once
	file_gRPC_proto_proto_rawDescData = file_gRPC_proto_proto_rawDesc
)

func file_gRPC_proto_proto_rawDescGZIP() []byte {
	file_gRPC_proto_proto_rawDescOnce.Do(func() {
		file_gRPC_proto_proto_rawDescData = protoimpl.X.CompressGZIP(file_gRPC_proto_proto_rawDescData)
	})
	return file_gRPC_proto_proto_rawDescData
}

var file_gRPC_proto_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_gRPC_proto_proto_goTypes = []interface{}{
	(*BidInput)(nil),               // 0: homework5.BidInput
	(*Outcome)(nil),                // 1: homework5.Outcome
	(*Confirmation)(nil),           // 2: homework5.Confirmation
	(*ConnectionAnnouncement)(nil), // 3: homework5.ConnectionAnnouncement
	(*UpdateAnnouncement)(nil),     // 4: homework5.UpdateAnnouncement
	(*AccessRequestResponse)(nil),  // 5: homework5.AccessRequestResponse
	(*AccessRequest)(nil),          // 6: homework5.AccessRequest
	(*Empty)(nil),                  // 7: homework5.Empty
}
var file_gRPC_proto_proto_depIdxs = []int32{
	0, // 0: homework5.ServerNode.Bid:input_type -> homework5.BidInput
	7, // 1: homework5.ServerNode.Result:input_type -> homework5.Empty
	3, // 2: homework5.ServerNode.AnnounceConnection:input_type -> homework5.ConnectionAnnouncement
	4, // 3: homework5.ServerNode.AnnounceUpdate:input_type -> homework5.UpdateAnnouncement
	6, // 4: homework5.ServerNode.RequestLeadership:input_type -> homework5.AccessRequest
	7, // 5: homework5.ServerNode.IExist:input_type -> homework5.Empty
	3, // 6: homework5.ServerNode.IAmLeader:input_type -> homework5.ConnectionAnnouncement
	2, // 7: homework5.ServerNode.Bid:output_type -> homework5.Confirmation
	1, // 8: homework5.ServerNode.Result:output_type -> homework5.Outcome
	2, // 9: homework5.ServerNode.AnnounceConnection:output_type -> homework5.Confirmation
	2, // 10: homework5.ServerNode.AnnounceUpdate:output_type -> homework5.Confirmation
	5, // 11: homework5.ServerNode.RequestLeadership:output_type -> homework5.AccessRequestResponse
	7, // 12: homework5.ServerNode.IExist:output_type -> homework5.Empty
	7, // 13: homework5.ServerNode.IAmLeader:output_type -> homework5.Empty
	7, // [7:14] is the sub-list for method output_type
	0, // [0:7] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gRPC_proto_proto_init() }
func file_gRPC_proto_proto_init() {
	if File_gRPC_proto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gRPC_proto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidInput); i {
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
		file_gRPC_proto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
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
		file_gRPC_proto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Confirmation); i {
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
		file_gRPC_proto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConnectionAnnouncement); i {
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
		file_gRPC_proto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateAnnouncement); i {
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
		file_gRPC_proto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequestResponse); i {
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
		file_gRPC_proto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessRequest); i {
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
		file_gRPC_proto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
			RawDescriptor: file_gRPC_proto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_gRPC_proto_proto_goTypes,
		DependencyIndexes: file_gRPC_proto_proto_depIdxs,
		MessageInfos:      file_gRPC_proto_proto_msgTypes,
	}.Build()
	File_gRPC_proto_proto = out.File
	file_gRPC_proto_proto_rawDesc = nil
	file_gRPC_proto_proto_goTypes = nil
	file_gRPC_proto_proto_depIdxs = nil
}
