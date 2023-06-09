// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.2
// source: paymentRequest.proto

package paymentRequest

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

type EachRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DealStock   int32 `protobuf:"varint,1,opt,name=deal_stock,json=dealStock,proto3" json:"deal_stock,omitempty"`
	TotalAmount int32 `protobuf:"varint,2,opt,name=total_amount,json=totalAmount,proto3" json:"total_amount,omitempty"`
}

func (x *EachRequest) Reset() {
	*x = EachRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentRequest_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EachRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EachRequest) ProtoMessage() {}

func (x *EachRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paymentRequest_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EachRequest.ProtoReflect.Descriptor instead.
func (*EachRequest) Descriptor() ([]byte, []int) {
	return file_paymentRequest_proto_rawDescGZIP(), []int{0}
}

func (x *EachRequest) GetDealStock() int32 {
	if x != nil {
		return x.DealStock
	}
	return 0
}

func (x *EachRequest) GetTotalAmount() int32 {
	if x != nil {
		return x.TotalAmount
	}
	return 0
}

type EachResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *EachResponse) Reset() {
	*x = EachResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentRequest_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EachResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EachResponse) ProtoMessage() {}

func (x *EachResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paymentRequest_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EachResponse.ProtoReflect.Descriptor instead.
func (*EachResponse) Descriptor() ([]byte, []int) {
	return file_paymentRequest_proto_rawDescGZIP(), []int{1}
}

func (x *EachResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *EachResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *EachResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegisterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Request      map[string]*EachRequest `protobuf:"bytes,1,rep,name=request,proto3" json:"request,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Address      string                  `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	CardNum      string                  `protobuf:"bytes,3,opt,name=card_num,json=cardNum,proto3" json:"card_num,omitempty"`
	ExpMonth     string                  `protobuf:"bytes,4,opt,name=exp_month,json=expMonth,proto3" json:"exp_month,omitempty"`
	ExpYear      string                  `protobuf:"bytes,5,opt,name=exp_year,json=expYear,proto3" json:"exp_year,omitempty"`
	Cvc          string                  `protobuf:"bytes,6,opt,name=cvc,proto3" json:"cvc,omitempty"`
	UserId       string                  `protobuf:"bytes,7,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Password     string                  `protobuf:"bytes,8,opt,name=password,proto3" json:"password,omitempty"`
	PaymentType  string                  `protobuf:"bytes,9,opt,name=payment_type,json=paymentType,proto3" json:"payment_type,omitempty"`
	Name         string                  `protobuf:"bytes,10,opt,name=name,proto3" json:"name,omitempty"`
	Itemname     string                  `protobuf:"bytes,11,opt,name=itemname,proto3" json:"itemname,omitempty"`
	Itemurl      string                  `protobuf:"bytes,12,opt,name=itemurl,proto3" json:"itemurl,omitempty"`
	Itemcategory int32                   `protobuf:"varint,13,opt,name=itemcategory,proto3" json:"itemcategory,omitempty"`
	Itemprice    int32                   `protobuf:"varint,14,opt,name=itemprice,proto3" json:"itemprice,omitempty"`
}

func (x *RegisterRequest) Reset() {
	*x = RegisterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentRequest_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterRequest) ProtoMessage() {}

func (x *RegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_paymentRequest_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterRequest.ProtoReflect.Descriptor instead.
func (*RegisterRequest) Descriptor() ([]byte, []int) {
	return file_paymentRequest_proto_rawDescGZIP(), []int{2}
}

func (x *RegisterRequest) GetRequest() map[string]*EachRequest {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *RegisterRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *RegisterRequest) GetCardNum() string {
	if x != nil {
		return x.CardNum
	}
	return ""
}

func (x *RegisterRequest) GetExpMonth() string {
	if x != nil {
		return x.ExpMonth
	}
	return ""
}

func (x *RegisterRequest) GetExpYear() string {
	if x != nil {
		return x.ExpYear
	}
	return ""
}

func (x *RegisterRequest) GetCvc() string {
	if x != nil {
		return x.Cvc
	}
	return ""
}

func (x *RegisterRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *RegisterRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *RegisterRequest) GetPaymentType() string {
	if x != nil {
		return x.PaymentType
	}
	return ""
}

func (x *RegisterRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RegisterRequest) GetItemname() string {
	if x != nil {
		return x.Itemname
	}
	return ""
}

func (x *RegisterRequest) GetItemurl() string {
	if x != nil {
		return x.Itemurl
	}
	return ""
}

func (x *RegisterRequest) GetItemcategory() int32 {
	if x != nil {
		return x.Itemcategory
	}
	return 0
}

func (x *RegisterRequest) GetItemprice() int32 {
	if x != nil {
		return x.Itemprice
	}
	return 0
}

type RegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   string                   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Response map[string]*EachResponse `protobuf:"bytes,2,rep,name=response,proto3" json:"response,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RegisterResponse) Reset() {
	*x = RegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_paymentRequest_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterResponse) ProtoMessage() {}

func (x *RegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_paymentRequest_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterResponse.ProtoReflect.Descriptor instead.
func (*RegisterResponse) Descriptor() ([]byte, []int) {
	return file_paymentRequest_proto_rawDescGZIP(), []int{3}
}

func (x *RegisterResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RegisterResponse) GetResponse() map[string]*EachResponse {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_paymentRequest_proto protoreflect.FileDescriptor

var file_paymentRequest_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0b, 0x45, 0x61, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x61, 0x6c, 0x5f, 0x73, 0x74,
	0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x64, 0x65, 0x61, 0x6c, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4c, 0x0a, 0x0c, 0x45, 0x61, 0x63, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0xf7, 0x03, 0x0a, 0x0f, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x07, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x70, 0x5f, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x78, 0x70, 0x4d,
	0x6f, 0x6e, 0x74, 0x68, 0x12, 0x19, 0x0a, 0x08, 0x65, 0x78, 0x70, 0x5f, 0x79, 0x65, 0x61, 0x72,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x78, 0x70, 0x59, 0x65, 0x61, 0x72, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x76, 0x63, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76,
	0x63, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x74, 0x65, 0x6d, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x74, 0x65,
	0x6d, 0x75, 0x72, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x69, 0x74, 0x65, 0x6d,
	0x75, 0x72, 0x6c, 0x12, 0x22, 0x0a, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x63, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0c, 0x69, 0x74, 0x65, 0x6d, 0x63,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x48, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x45, 0x61, 0x63, 0x68, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0xb3, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3b, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x1a, 0x4a, 0x0a, 0x0d, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x23, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x45, 0x61,
	0x63, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x49, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3a, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_paymentRequest_proto_rawDescOnce sync.Once
	file_paymentRequest_proto_rawDescData = file_paymentRequest_proto_rawDesc
)

func file_paymentRequest_proto_rawDescGZIP() []byte {
	file_paymentRequest_proto_rawDescOnce.Do(func() {
		file_paymentRequest_proto_rawDescData = protoimpl.X.CompressGZIP(file_paymentRequest_proto_rawDescData)
	})
	return file_paymentRequest_proto_rawDescData
}

var file_paymentRequest_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_paymentRequest_proto_goTypes = []interface{}{
	(*EachRequest)(nil),      // 0: EachRequest
	(*EachResponse)(nil),     // 1: EachResponse
	(*RegisterRequest)(nil),  // 2: RegisterRequest
	(*RegisterResponse)(nil), // 3: RegisterResponse
	nil,                      // 4: RegisterRequest.RequestEntry
	nil,                      // 5: RegisterResponse.ResponseEntry
}
var file_paymentRequest_proto_depIdxs = []int32{
	4, // 0: RegisterRequest.request:type_name -> RegisterRequest.RequestEntry
	5, // 1: RegisterResponse.response:type_name -> RegisterResponse.ResponseEntry
	0, // 2: RegisterRequest.RequestEntry.value:type_name -> EachRequest
	1, // 3: RegisterResponse.ResponseEntry.value:type_name -> EachResponse
	2, // 4: Transaction.CreateTransaction:input_type -> RegisterRequest
	3, // 5: Transaction.CreateTransaction:output_type -> RegisterResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_paymentRequest_proto_init() }
func file_paymentRequest_proto_init() {
	if File_paymentRequest_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_paymentRequest_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EachRequest); i {
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
		file_paymentRequest_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EachResponse); i {
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
		file_paymentRequest_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterRequest); i {
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
		file_paymentRequest_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterResponse); i {
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
			RawDescriptor: file_paymentRequest_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_paymentRequest_proto_goTypes,
		DependencyIndexes: file_paymentRequest_proto_depIdxs,
		MessageInfos:      file_paymentRequest_proto_msgTypes,
	}.Build()
	File_paymentRequest_proto = out.File
	file_paymentRequest_proto_rawDesc = nil
	file_paymentRequest_proto_goTypes = nil
	file_paymentRequest_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TransactionClient is the client API for Transaction service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TransactionClient interface {
	CreateTransaction(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
}

type transactionClient struct {
	cc grpc.ClientConnInterface
}

func NewTransactionClient(cc grpc.ClientConnInterface) TransactionClient {
	return &transactionClient{cc}
}

func (c *transactionClient) CreateTransaction(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/Transaction/CreateTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TransactionServer is the server API for Transaction service.
type TransactionServer interface {
	CreateTransaction(context.Context, *RegisterRequest) (*RegisterResponse, error)
}

// UnimplementedTransactionServer can be embedded to have forward compatible implementations.
type UnimplementedTransactionServer struct {
}

func (*UnimplementedTransactionServer) CreateTransaction(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTransaction not implemented")
}

func RegisterTransactionServer(s *grpc.Server, srv TransactionServer) {
	s.RegisterService(&_Transaction_serviceDesc, srv)
}

func _Transaction_CreateTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TransactionServer).CreateTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Transaction/CreateTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TransactionServer).CreateTransaction(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Transaction_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Transaction",
	HandlerType: (*TransactionServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTransaction",
			Handler:    _Transaction_CreateTransaction_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "paymentRequest.proto",
}
