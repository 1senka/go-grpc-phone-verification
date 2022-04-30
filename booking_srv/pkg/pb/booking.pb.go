// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.0
// source: pkg/pb/booking.proto

package pb

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

type BookingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	SessionType string `protobuf:"bytes,2,opt,name=sessionType,proto3" json:"sessionType,omitempty"`
	IsHotline   bool   `protobuf:"varint,3,opt,name=isHotline,proto3" json:"isHotline,omitempty"`
}

func (x *BookingRequest) Reset() {
	*x = BookingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_booking_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingRequest) ProtoMessage() {}

func (x *BookingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_booking_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingRequest.ProtoReflect.Descriptor instead.
func (*BookingRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_booking_proto_rawDescGZIP(), []int{0}
}

func (x *BookingRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BookingRequest) GetSessionType() string {
	if x != nil {
		return x.SessionType
	}
	return ""
}

func (x *BookingRequest) GetIsHotline() bool {
	if x != nil {
		return x.IsHotline
	}
	return false
}

type BookingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId        string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	SessionType   string `protobuf:"bytes,2,opt,name=sessionType,proto3" json:"sessionType,omitempty"`
	IsHotline     bool   `protobuf:"varint,3,opt,name=isHotline,proto3" json:"isHotline,omitempty"`
	BookingId     string `protobuf:"bytes,4,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
	BookingTime   string `protobuf:"bytes,5,opt,name=bookingTime,proto3" json:"bookingTime,omitempty"`
	BookingStatus string `protobuf:"bytes,6,opt,name=bookingStatus,proto3" json:"bookingStatus,omitempty"`
	Status        int32  `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	TherapistId   string `protobuf:"bytes,8,opt,name=therapistId,proto3" json:"therapistId,omitempty"`
	TherapistName string `protobuf:"bytes,9,opt,name=therapistName,proto3" json:"therapistName,omitempty"`
}

func (x *BookingResponse) Reset() {
	*x = BookingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_booking_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookingResponse) ProtoMessage() {}

func (x *BookingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_booking_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookingResponse.ProtoReflect.Descriptor instead.
func (*BookingResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_booking_proto_rawDescGZIP(), []int{1}
}

func (x *BookingResponse) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *BookingResponse) GetSessionType() string {
	if x != nil {
		return x.SessionType
	}
	return ""
}

func (x *BookingResponse) GetIsHotline() bool {
	if x != nil {
		return x.IsHotline
	}
	return false
}

func (x *BookingResponse) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

func (x *BookingResponse) GetBookingTime() string {
	if x != nil {
		return x.BookingTime
	}
	return ""
}

func (x *BookingResponse) GetBookingStatus() string {
	if x != nil {
		return x.BookingStatus
	}
	return ""
}

func (x *BookingResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *BookingResponse) GetTherapistId() string {
	if x != nil {
		return x.TherapistId
	}
	return ""
}

func (x *BookingResponse) GetTherapistName() string {
	if x != nil {
		return x.TherapistName
	}
	return ""
}

type Booking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BookingId     string `protobuf:"bytes,1,opt,name=bookingId,proto3" json:"bookingId,omitempty"`
	UserId        string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	SessionType   string `protobuf:"bytes,3,opt,name=sessionType,proto3" json:"sessionType,omitempty"`
	IsHotline     bool   `protobuf:"varint,4,opt,name=isHotline,proto3" json:"isHotline,omitempty"`
	BookingTime   string `protobuf:"bytes,5,opt,name=bookingTime,proto3" json:"bookingTime,omitempty"`
	BookingStatus string `protobuf:"bytes,6,opt,name=bookingStatus,proto3" json:"bookingStatus,omitempty"`
	Status        int32  `protobuf:"varint,7,opt,name=status,proto3" json:"status,omitempty"`
	TherapistId   string `protobuf:"bytes,8,opt,name=therapistId,proto3" json:"therapistId,omitempty"`
	TherapistName string `protobuf:"bytes,9,opt,name=therapistName,proto3" json:"therapistName,omitempty"`
	IsPaid        bool   `protobuf:"varint,10,opt,name=isPaid,proto3" json:"isPaid,omitempty"`
}

func (x *Booking) Reset() {
	*x = Booking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_booking_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Booking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Booking) ProtoMessage() {}

func (x *Booking) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_booking_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Booking.ProtoReflect.Descriptor instead.
func (*Booking) Descriptor() ([]byte, []int) {
	return file_pkg_pb_booking_proto_rawDescGZIP(), []int{2}
}

func (x *Booking) GetBookingId() string {
	if x != nil {
		return x.BookingId
	}
	return ""
}

func (x *Booking) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Booking) GetSessionType() string {
	if x != nil {
		return x.SessionType
	}
	return ""
}

func (x *Booking) GetIsHotline() bool {
	if x != nil {
		return x.IsHotline
	}
	return false
}

func (x *Booking) GetBookingTime() string {
	if x != nil {
		return x.BookingTime
	}
	return ""
}

func (x *Booking) GetBookingStatus() string {
	if x != nil {
		return x.BookingStatus
	}
	return ""
}

func (x *Booking) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Booking) GetTherapistId() string {
	if x != nil {
		return x.TherapistId
	}
	return ""
}

func (x *Booking) GetTherapistName() string {
	if x != nil {
		return x.TherapistName
	}
	return ""
}

func (x *Booking) GetIsPaid() bool {
	if x != nil {
		return x.IsPaid
	}
	return false
}

type GetBookingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *GetBookingListRequest) Reset() {
	*x = GetBookingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_booking_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingListRequest) ProtoMessage() {}

func (x *GetBookingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_booking_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingListRequest.ProtoReflect.Descriptor instead.
func (*GetBookingListRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_booking_proto_rawDescGZIP(), []int{3}
}

func (x *GetBookingListRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type TherapistGetFreeTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *TherapistGetFreeTimeRequest) Reset() {
	*x = TherapistGetFreeTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_booking_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TherapistGetFreeTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TherapistGetFreeTimeRequest) ProtoMessage() {}

func (x *TherapistGetFreeTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_booking_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TherapistGetFreeTimeRequest.ProtoReflect.Descriptor instead.
func (*TherapistGetFreeTimeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_booking_proto_rawDescGZIP(), []int{4}
}

func (x *TherapistGetFreeTimeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type TherapistGetFreeTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   int32       `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	FreeTime []*FreeTime `protobuf:"bytes,2,rep,name=free_time,json=freeTime,proto3" json:"free_time,omitempty"`
}

func (x *TherapistGetFreeTimeResponse) Reset() {
	*x = TherapistGetFreeTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_booking_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TherapistGetFreeTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TherapistGetFreeTimeResponse) ProtoMessage() {}

func (x *TherapistGetFreeTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_booking_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TherapistGetFreeTimeResponse.ProtoReflect.Descriptor instead.
func (*TherapistGetFreeTimeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_booking_proto_rawDescGZIP(), []int{5}
}

func (x *TherapistGetFreeTimeResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TherapistGetFreeTimeResponse) GetFreeTime() []*FreeTime {
	if x != nil {
		return x.FreeTime
	}
	return nil
}

type Date struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Year   int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	Month  int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	Day    int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	Hour   int32 `protobuf:"varint,4,opt,name=hour,proto3" json:"hour,omitempty"`
	Minute int32 `protobuf:"varint,5,opt,name=minute,proto3" json:"minute,omitempty"`
}

func (x *Date) Reset() {
	*x = Date{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_booking_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Date) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Date) ProtoMessage() {}

func (x *Date) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_booking_proto_msgTypes[6]
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
	return file_pkg_pb_booking_proto_rawDescGZIP(), []int{6}
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

func (x *Date) GetHour() int32 {
	if x != nil {
		return x.Hour
	}
	return 0
}

func (x *Date) GetMinute() int32 {
	if x != nil {
		return x.Minute
	}
	return 0
}

type TherapistSetFreeTimeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	FreeTime []*Date `protobuf:"bytes,2,rep,name=free_time,json=freeTime,proto3" json:"free_time,omitempty"`
}

func (x *TherapistSetFreeTimeRequest) Reset() {
	*x = TherapistSetFreeTimeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_booking_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TherapistSetFreeTimeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TherapistSetFreeTimeRequest) ProtoMessage() {}

func (x *TherapistSetFreeTimeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_booking_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TherapistSetFreeTimeRequest.ProtoReflect.Descriptor instead.
func (*TherapistSetFreeTimeRequest) Descriptor() ([]byte, []int) {
	return file_pkg_pb_booking_proto_rawDescGZIP(), []int{7}
}

func (x *TherapistSetFreeTimeRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *TherapistSetFreeTimeRequest) GetFreeTime() []*Date {
	if x != nil {
		return x.FreeTime
	}
	return nil
}

type TherapistSetFreeTimeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *TherapistSetFreeTimeResponse) Reset() {
	*x = TherapistSetFreeTimeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_booking_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TherapistSetFreeTimeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TherapistSetFreeTimeResponse) ProtoMessage() {}

func (x *TherapistSetFreeTimeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_booking_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TherapistSetFreeTimeResponse.ProtoReflect.Descriptor instead.
func (*TherapistSetFreeTimeResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_booking_proto_rawDescGZIP(), []int{8}
}

func (x *TherapistSetFreeTimeResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *TherapistSetFreeTimeResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

type FreeTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date        *Date  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	TherapistId string `protobuf:"bytes,2,opt,name=therapist_id,json=therapistId,proto3" json:"therapist_id,omitempty"`
}

func (x *FreeTime) Reset() {
	*x = FreeTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_booking_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FreeTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FreeTime) ProtoMessage() {}

func (x *FreeTime) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_booking_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FreeTime.ProtoReflect.Descriptor instead.
func (*FreeTime) Descriptor() ([]byte, []int) {
	return file_pkg_pb_booking_proto_rawDescGZIP(), []int{9}
}

func (x *FreeTime) GetDate() *Date {
	if x != nil {
		return x.Date
	}
	return nil
}

func (x *FreeTime) GetTherapistId() string {
	if x != nil {
		return x.TherapistId
	}
	return ""
}

type GetBookingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bookings []*Booking `protobuf:"bytes,1,rep,name=bookings,proto3" json:"bookings,omitempty"`
	Status   int32      `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetBookingListResponse) Reset() {
	*x = GetBookingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_pb_booking_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookingListResponse) ProtoMessage() {}

func (x *GetBookingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_pb_booking_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookingListResponse.ProtoReflect.Descriptor instead.
func (*GetBookingListResponse) Descriptor() ([]byte, []int) {
	return file_pkg_pb_booking_proto_rawDescGZIP(), []int{10}
}

func (x *GetBookingListResponse) GetBookings() []*Booking {
	if x != nil {
		return x.Bookings
	}
	return nil
}

func (x *GetBookingListResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_pkg_pb_booking_proto protoreflect.FileDescriptor

var file_pkg_pb_booking_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22,
	0x68, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69,
	0x73, 0x48, 0x6f, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x69, 0x73, 0x48, 0x6f, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0xaf, 0x02, 0x0a, 0x0f, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x73, 0x48, 0x6f, 0x74,
	0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x48, 0x6f,
	0x74, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69,
	0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x74, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74,
	0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x68, 0x65, 0x72, 0x61, 0x70,
	0x69, 0x73, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x74, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69,
	0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x68,
	0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0xbf, 0x02, 0x0a, 0x07,
	0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x48, 0x6f, 0x74, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x74, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x74, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74, 0x49, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x74, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x50, 0x61, 0x69, 0x64, 0x22, 0x2f, 0x0a,
	0x15, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2d,
	0x0a, 0x1b, 0x54, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74, 0x47, 0x65, 0x74, 0x46, 0x72,
	0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x66, 0x0a,
	0x1c, 0x54, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2e, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x46, 0x72, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x08, 0x66, 0x72, 0x65,
	0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x6e, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x6f, 0x75,
	0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x68, 0x6f, 0x75, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x6d, 0x69, 0x6e, 0x75, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6d,
	0x69, 0x6e, 0x75, 0x74, 0x65, 0x22, 0x59, 0x0a, 0x1b, 0x54, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69,
	0x73, 0x74, 0x53, 0x65, 0x74, 0x46, 0x72, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x09, 0x66, 0x72, 0x65, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x08, 0x66, 0x72, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x4e, 0x0a, 0x1c, 0x54, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74, 0x53, 0x65, 0x74,
	0x46, 0x72, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x50, 0x0a, 0x08, 0x46, 0x72, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x74, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x74, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74,
	0x49, 0x64, 0x22, 0x5e, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x08,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x52, 0x08, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x32, 0xe2, 0x02, 0x0a, 0x0e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c, 0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x12, 0x17, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x62, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x54, 0x68, 0x65,
	0x72, 0x61, 0x70, 0x69, 0x73, 0x74, 0x46, 0x72, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69,
	0x73, 0x74, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54,
	0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74, 0x47, 0x65, 0x74, 0x46, 0x72, 0x65, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x53,
	0x65, 0x74, 0x46, 0x72, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x24, 0x2e, 0x62, 0x6f, 0x6f,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x68, 0x65, 0x72, 0x61, 0x70, 0x69, 0x73, 0x74, 0x53, 0x65,
	0x74, 0x46, 0x72, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x25, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x54, 0x68, 0x65, 0x72, 0x61,
	0x70, 0x69, 0x73, 0x74, 0x53, 0x65, 0x74, 0x46, 0x72, 0x65, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_pb_booking_proto_rawDescOnce sync.Once
	file_pkg_pb_booking_proto_rawDescData = file_pkg_pb_booking_proto_rawDesc
)

func file_pkg_pb_booking_proto_rawDescGZIP() []byte {
	file_pkg_pb_booking_proto_rawDescOnce.Do(func() {
		file_pkg_pb_booking_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_pb_booking_proto_rawDescData)
	})
	return file_pkg_pb_booking_proto_rawDescData
}

var file_pkg_pb_booking_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_pkg_pb_booking_proto_goTypes = []interface{}{
	(*BookingRequest)(nil),               // 0: booking.BookingRequest
	(*BookingResponse)(nil),              // 1: booking.BookingResponse
	(*Booking)(nil),                      // 2: booking.Booking
	(*GetBookingListRequest)(nil),        // 3: booking.GetBookingListRequest
	(*TherapistGetFreeTimeRequest)(nil),  // 4: booking.TherapistGetFreeTimeRequest
	(*TherapistGetFreeTimeResponse)(nil), // 5: booking.TherapistGetFreeTimeResponse
	(*Date)(nil),                         // 6: booking.Date
	(*TherapistSetFreeTimeRequest)(nil),  // 7: booking.TherapistSetFreeTimeRequest
	(*TherapistSetFreeTimeResponse)(nil), // 8: booking.TherapistSetFreeTimeResponse
	(*FreeTime)(nil),                     // 9: booking.FreeTime
	(*GetBookingListResponse)(nil),       // 10: booking.GetBookingListResponse
}
var file_pkg_pb_booking_proto_depIdxs = []int32{
	9,  // 0: booking.TherapistGetFreeTimeResponse.free_time:type_name -> booking.FreeTime
	6,  // 1: booking.TherapistSetFreeTimeRequest.free_time:type_name -> booking.Date
	6,  // 2: booking.FreeTime.date:type_name -> booking.Date
	2,  // 3: booking.GetBookingListResponse.bookings:type_name -> booking.Booking
	0,  // 4: booking.BookingService.Booking:input_type -> booking.BookingRequest
	3,  // 5: booking.BookingService.GetBookingList:input_type -> booking.GetBookingListRequest
	4,  // 6: booking.BookingService.GetTherapistFreeTime:input_type -> booking.TherapistGetFreeTimeRequest
	7,  // 7: booking.BookingService.SetFreeTime:input_type -> booking.TherapistSetFreeTimeRequest
	1,  // 8: booking.BookingService.Booking:output_type -> booking.BookingResponse
	10, // 9: booking.BookingService.GetBookingList:output_type -> booking.GetBookingListResponse
	5,  // 10: booking.BookingService.GetTherapistFreeTime:output_type -> booking.TherapistGetFreeTimeResponse
	8,  // 11: booking.BookingService.SetFreeTime:output_type -> booking.TherapistSetFreeTimeResponse
	8,  // [8:12] is the sub-list for method output_type
	4,  // [4:8] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_pkg_pb_booking_proto_init() }
func file_pkg_pb_booking_proto_init() {
	if File_pkg_pb_booking_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_pb_booking_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingRequest); i {
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
		file_pkg_pb_booking_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookingResponse); i {
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
		file_pkg_pb_booking_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Booking); i {
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
		file_pkg_pb_booking_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookingListRequest); i {
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
		file_pkg_pb_booking_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TherapistGetFreeTimeRequest); i {
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
		file_pkg_pb_booking_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TherapistGetFreeTimeResponse); i {
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
		file_pkg_pb_booking_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_pkg_pb_booking_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TherapistSetFreeTimeRequest); i {
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
		file_pkg_pb_booking_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TherapistSetFreeTimeResponse); i {
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
		file_pkg_pb_booking_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FreeTime); i {
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
		file_pkg_pb_booking_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookingListResponse); i {
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
			RawDescriptor: file_pkg_pb_booking_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_pb_booking_proto_goTypes,
		DependencyIndexes: file_pkg_pb_booking_proto_depIdxs,
		MessageInfos:      file_pkg_pb_booking_proto_msgTypes,
	}.Build()
	File_pkg_pb_booking_proto = out.File
	file_pkg_pb_booking_proto_rawDesc = nil
	file_pkg_pb_booking_proto_goTypes = nil
	file_pkg_pb_booking_proto_depIdxs = nil
}
