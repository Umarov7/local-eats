// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: payment.proto

package payment

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

type NewPayment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId       string `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PaymentMethod string `protobuf:"bytes,2,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
	CardNumber    string `protobuf:"bytes,3,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	ExpiryDate    string `protobuf:"bytes,4,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	Cvv           string `protobuf:"bytes,5,opt,name=cvv,proto3" json:"cvv,omitempty"`
}

func (x *NewPayment) Reset() {
	*x = NewPayment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPayment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPayment) ProtoMessage() {}

func (x *NewPayment) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPayment.ProtoReflect.Descriptor instead.
func (*NewPayment) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{0}
}

func (x *NewPayment) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *NewPayment) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

func (x *NewPayment) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *NewPayment) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *NewPayment) GetCvv() string {
	if x != nil {
		return x.Cvv
	}
	return ""
}

type NewPaymentResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       string  `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Amount        float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Status        string  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	TransactionId string  `protobuf:"bytes,5,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	CreatedAt     string  `protobuf:"bytes,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *NewPaymentResp) Reset() {
	*x = NewPaymentResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPaymentResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPaymentResp) ProtoMessage() {}

func (x *NewPaymentResp) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPaymentResp.ProtoReflect.Descriptor instead.
func (*NewPaymentResp) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{1}
}

func (x *NewPaymentResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NewPaymentResp) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *NewPaymentResp) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *NewPaymentResp) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *NewPaymentResp) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *NewPaymentResp) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *ID) Reset() {
	*x = ID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ID) ProtoMessage() {}

func (x *ID) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ID.ProtoReflect.Descriptor instead.
func (*ID) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{2}
}

func (x *ID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type PaymentDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId       string  `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Amount        float32 `protobuf:"fixed32,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Status        string  `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	Method        string  `protobuf:"bytes,5,opt,name=method,proto3" json:"method,omitempty"`
	CardNumber    string  `protobuf:"bytes,6,opt,name=card_number,json=cardNumber,proto3" json:"card_number,omitempty"`
	ExpiryDate    string  `protobuf:"bytes,7,opt,name=expiry_date,json=expiryDate,proto3" json:"expiry_date,omitempty"`
	Cvv           string  `protobuf:"bytes,8,opt,name=cvv,proto3" json:"cvv,omitempty"`
	TransactionId string  `protobuf:"bytes,9,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id,omitempty"`
	CreatedAt     string  `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PaymentDetails) Reset() {
	*x = PaymentDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentDetails) ProtoMessage() {}

func (x *PaymentDetails) ProtoReflect() protoreflect.Message {
	mi := &file_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentDetails.ProtoReflect.Descriptor instead.
func (*PaymentDetails) Descriptor() ([]byte, []int) {
	return file_payment_proto_rawDescGZIP(), []int{3}
}

func (x *PaymentDetails) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PaymentDetails) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *PaymentDetails) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *PaymentDetails) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *PaymentDetails) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *PaymentDetails) GetCardNumber() string {
	if x != nil {
		return x.CardNumber
	}
	return ""
}

func (x *PaymentDetails) GetExpiryDate() string {
	if x != nil {
		return x.ExpiryDate
	}
	return ""
}

func (x *PaymentDetails) GetCvv() string {
	if x != nil {
		return x.Cvv
	}
	return ""
}

func (x *PaymentDetails) GetTransactionId() string {
	if x != nil {
		return x.TransactionId
	}
	return ""
}

func (x *PaymentDetails) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_payment_proto protoreflect.FileDescriptor

var file_payment_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0xa2, 0x01, 0x0a, 0x0a, 0x4e, 0x65, 0x77,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x61, 0x72,
	0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78,
	0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63,
	0x76, 0x76, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76, 0x76, 0x22, 0xb1, 0x01,
	0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x14, 0x0a, 0x02, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x9d, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72,
	0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1f, 0x0a,
	0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x61, 0x72, 0x64, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x1f,
	0x0a, 0x0b, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x63, 0x76, 0x76, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76,
	0x76, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x7a, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x3b, 0x0a, 0x0b, 0x4d, 0x61, 0x6b, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x13, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x4e, 0x65, 0x77, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x4e, 0x65, 0x77, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x32, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x0b, 0x2e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x49, 0x44, 0x1a, 0x17, 0x2e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x42, 0x12, 0x5a, 0x10, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_payment_proto_rawDescOnce sync.Once
	file_payment_proto_rawDescData = file_payment_proto_rawDesc
)

func file_payment_proto_rawDescGZIP() []byte {
	file_payment_proto_rawDescOnce.Do(func() {
		file_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_payment_proto_rawDescData)
	})
	return file_payment_proto_rawDescData
}

var file_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_payment_proto_goTypes = []interface{}{
	(*NewPayment)(nil),     // 0: payment.NewPayment
	(*NewPaymentResp)(nil), // 1: payment.NewPaymentResp
	(*ID)(nil),             // 2: payment.ID
	(*PaymentDetails)(nil), // 3: payment.PaymentDetails
}
var file_payment_proto_depIdxs = []int32{
	0, // 0: payment.Payment.MakePayment:input_type -> payment.NewPayment
	2, // 1: payment.Payment.GetPayment:input_type -> payment.ID
	1, // 2: payment.Payment.MakePayment:output_type -> payment.NewPaymentResp
	3, // 3: payment.Payment.GetPayment:output_type -> payment.PaymentDetails
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_payment_proto_init() }
func file_payment_proto_init() {
	if File_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_payment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPayment); i {
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
		file_payment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPaymentResp); i {
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
		file_payment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ID); i {
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
		file_payment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentDetails); i {
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
			RawDescriptor: file_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_payment_proto_goTypes,
		DependencyIndexes: file_payment_proto_depIdxs,
		MessageInfos:      file_payment_proto_msgTypes,
	}.Build()
	File_payment_proto = out.File
	file_payment_proto_rawDesc = nil
	file_payment_proto_goTypes = nil
	file_payment_proto_depIdxs = nil
}