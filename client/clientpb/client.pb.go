// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: client/clientpb/client.proto

package clientpb

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

type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year  int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day   int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_clientpb_client_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_client_clientpb_client_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Date.ProtoReflect.Descriptor instead.
func (*Date) Descriptor() ([]byte, []int) {
	return file_client_clientpb_client_proto_rawDescGZIP(), []int{0}
}

func (x *Date) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Date) GetMonth() int32 {
	if x != nil {
		return x.Month
	}
	return 0
}

func (x *Date) GetDay() int32 {
	if x != nil {
		return x.Day
	}
	return 0
}

type ClientFirstCredRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Birthday *Date  `protobuf:"bytes,2,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Phone    string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *ClientFirstCredRequest) Reset() {
	*x = ClientFirstCredRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_clientpb_client_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientFirstCredRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientFirstCredRequest) ProtoMessage() {}

func (x *ClientFirstCredRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_clientpb_client_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientFirstCredRequest.ProtoReflect.Descriptor instead.
func (*ClientFirstCredRequest) Descriptor() ([]byte, []int) {
	return file_client_clientpb_client_proto_rawDescGZIP(), []int{1}
}

func (x *ClientFirstCredRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ClientFirstCredRequest) GetBirthday() *Date {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *ClientFirstCredRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type ClientFirstCredResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ClientFirstCredResponse) Reset() {
	*x = ClientFirstCredResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_clientpb_client_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientFirstCredResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientFirstCredResponse) ProtoMessage() {}

func (x *ClientFirstCredResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_clientpb_client_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientFirstCredResponse.ProtoReflect.Descriptor instead.
func (*ClientFirstCredResponse) Descriptor() ([]byte, []int) {
	return file_client_clientpb_client_proto_rawDescGZIP(), []int{2}
}

func (x *ClientFirstCredResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type ClientPhoneProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *ClientPhoneProofRequest) Reset() {
	*x = ClientPhoneProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_clientpb_client_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPhoneProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPhoneProofRequest) ProtoMessage() {}

func (x *ClientPhoneProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_clientpb_client_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPhoneProofRequest.ProtoReflect.Descriptor instead.
func (*ClientPhoneProofRequest) Descriptor() ([]byte, []int) {
	return file_client_clientpb_client_proto_rawDescGZIP(), []int{3}
}

func (x *ClientPhoneProofRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type ClientPhoneProofResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid string `protobuf:"bytes,1,opt,name=isValid,proto3" json:"isValid,omitempty"`
}

func (x *ClientPhoneProofResponse) Reset() {
	*x = ClientPhoneProofResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_clientpb_client_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPhoneProofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPhoneProofResponse) ProtoMessage() {}

func (x *ClientPhoneProofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_clientpb_client_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPhoneProofResponse.ProtoReflect.Descriptor instead.
func (*ClientPhoneProofResponse) Descriptor() ([]byte, []int) {
	return file_client_clientpb_client_proto_rawDescGZIP(), []int{4}
}

func (x *ClientPhoneProofResponse) GetIsValid() string {
	if x != nil {
		return x.IsValid
	}
	return ""
}

type ClientPhoneTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *ClientPhoneTokenRequest) Reset() {
	*x = ClientPhoneTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_clientpb_client_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPhoneTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPhoneTokenRequest) ProtoMessage() {}

func (x *ClientPhoneTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_client_clientpb_client_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPhoneTokenRequest.ProtoReflect.Descriptor instead.
func (*ClientPhoneTokenRequest) Descriptor() ([]byte, []int) {
	return file_client_clientpb_client_proto_rawDescGZIP(), []int{5}
}

func (x *ClientPhoneTokenRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *ClientPhoneTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ClientPhoneTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ClientPhoneTokenResponse) Reset() {
	*x = ClientPhoneTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_client_clientpb_client_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientPhoneTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientPhoneTokenResponse) ProtoMessage() {}

func (x *ClientPhoneTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_client_clientpb_client_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientPhoneTokenResponse.ProtoReflect.Descriptor instead.
func (*ClientPhoneTokenResponse) Descriptor() ([]byte, []int) {
	return file_client_clientpb_client_proto_rawDescGZIP(), []int{6}
}

func (x *ClientPhoneTokenResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

var File_client_clientpb_client_proto protoreflect.FileDescriptor

var file_client_clientpb_client_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x70,
	0x62, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x22, 0x42, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x22, 0x6c, 0x0a, 0x16, 0x43, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74, 0x43, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x08, 0x62, 0x69, 0x72, 0x74,
	0x68, 0x64, 0x61, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64,
	0x61, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x31, 0x0a, 0x17, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x46, 0x69, 0x72, 0x73, 0x74, 0x43, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x2f, 0x0a, 0x17, 0x43,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x34, 0x0a, 0x18,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x73, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x73, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x22, 0x45, 0x0a, 0x17, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x32, 0x0a, 0x18, 0x43, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0xff, 0x01,
	0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4f, 0x0a, 0x0a, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x65, 0x12, 0x1f, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x2e, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4c, 0x0a, 0x09, 0x46, 0x69, 0x72, 0x73, 0x74, 0x43, 0x72, 0x65, 0x64, 0x12, 0x1e, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x43, 0x72, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x46, 0x69, 0x72,
	0x73, 0x74, 0x43, 0x72, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4f,
	0x0a, 0x0a, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1f, 0x2e, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_client_clientpb_client_proto_rawDescOnce sync.Once
	file_client_clientpb_client_proto_rawDescData = file_client_clientpb_client_proto_rawDesc
)

func file_client_clientpb_client_proto_rawDescGZIP() []byte {
	file_client_clientpb_client_proto_rawDescOnce.Do(func() {
		file_client_clientpb_client_proto_rawDescData = protoimpl.X.CompressGZIP(file_client_clientpb_client_proto_rawDescData)
	})
	return file_client_clientpb_client_proto_rawDescData
}

var file_client_clientpb_client_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_client_clientpb_client_proto_goTypes = []interface{}{
	(*Date)(nil),                     // 0: client.Date
	(*ClientFirstCredRequest)(nil),   // 1: client.ClientFirstCredRequest
	(*ClientFirstCredResponse)(nil),  // 2: client.ClientFirstCredResponse
	(*ClientPhoneProofRequest)(nil),  // 3: client.ClientPhoneProofRequest
	(*ClientPhoneProofResponse)(nil), // 4: client.ClientPhoneProofResponse
	(*ClientPhoneTokenRequest)(nil),  // 5: client.ClientPhoneTokenRequest
	(*ClientPhoneTokenResponse)(nil), // 6: client.ClientPhoneTokenResponse
}
var file_client_clientpb_client_proto_depIdxs = []int32{
	0, // 0: client.ClientFirstCredRequest.birthday:type_name -> client.Date
	3, // 1: client.ClientService.PhoneProve:input_type -> client.ClientPhoneProofRequest
	1, // 2: client.ClientService.FirstCred:input_type -> client.ClientFirstCredRequest
	5, // 3: client.ClientService.PhoneToken:input_type -> client.ClientPhoneTokenRequest
	4, // 4: client.ClientService.PhoneProve:output_type -> client.ClientPhoneProofResponse
	2, // 5: client.ClientService.FirstCred:output_type -> client.ClientFirstCredResponse
	6, // 6: client.ClientService.PhoneToken:output_type -> client.ClientPhoneTokenResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_client_clientpb_client_proto_init() }
func file_client_clientpb_client_proto_init() {
	if File_client_clientpb_client_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_client_clientpb_client_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Date); i {
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
		file_client_clientpb_client_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientFirstCredRequest); i {
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
		file_client_clientpb_client_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientFirstCredResponse); i {
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
		file_client_clientpb_client_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPhoneProofRequest); i {
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
		file_client_clientpb_client_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPhoneProofResponse); i {
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
		file_client_clientpb_client_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPhoneTokenRequest); i {
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
		file_client_clientpb_client_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientPhoneTokenResponse); i {
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
			RawDescriptor: file_client_clientpb_client_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_client_clientpb_client_proto_goTypes,
		DependencyIndexes: file_client_clientpb_client_proto_depIdxs,
		MessageInfos:      file_client_clientpb_client_proto_msgTypes,
	}.Build()
	File_client_clientpb_client_proto = out.File
	file_client_clientpb_client_proto_rawDesc = nil
	file_client_clientpb_client_proto_goTypes = nil
	file_client_clientpb_client_proto_depIdxs = nil
}
