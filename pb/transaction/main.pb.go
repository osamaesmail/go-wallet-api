// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: pb/transaction/main.proto

package transaction

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

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FromAccount string  `protobuf:"bytes,1,opt,name=FromAccount,proto3" json:"FromAccount,omitempty"`
	ToAccount   string  `protobuf:"bytes,2,opt,name=ToAccount,proto3" json:"ToAccount,omitempty"`
	Amount      float64 `protobuf:"fixed64,3,opt,name=Amount,proto3" json:"Amount,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_transaction_main_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transaction_main_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_pb_transaction_main_proto_rawDescGZIP(), []int{0}
}

func (x *CreateRequest) GetFromAccount() string {
	if x != nil {
		return x.FromAccount
	}
	return ""
}

func (x *CreateRequest) GetToAccount() string {
	if x != nil {
		return x.ToAccount
	}
	return ""
}

func (x *CreateRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

type ListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID  string `protobuf:"bytes,1,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Page    int64  `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	PerPage int64  `protobuf:"varint,3,opt,name=PerPage,proto3" json:"PerPage,omitempty"`
}

func (x *ListRequest) Reset() {
	*x = ListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_transaction_main_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequest) ProtoMessage() {}

func (x *ListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transaction_main_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequest.ProtoReflect.Descriptor instead.
func (*ListRequest) Descriptor() ([]byte, []int) {
	return file_pb_transaction_main_proto_rawDescGZIP(), []int{1}
}

func (x *ListRequest) GetUserID() string {
	if x != nil {
		return x.UserID
	}
	return ""
}

func (x *ListRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ListRequest) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID          string  `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	FromAccount string  `protobuf:"bytes,2,opt,name=FromAccount,proto3" json:"FromAccount,omitempty"`
	ToAccount   string  `protobuf:"bytes,3,opt,name=ToAccount,proto3" json:"ToAccount,omitempty"`
	Amount      float64 `protobuf:"fixed64,4,opt,name=Amount,proto3" json:"Amount,omitempty"`
	Currency    string  `protobuf:"bytes,5,opt,name=Currency,proto3" json:"Currency,omitempty"`
	CreatedAt   string  `protobuf:"bytes,6,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_transaction_main_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transaction_main_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_pb_transaction_main_proto_rawDescGZIP(), []int{2}
}

func (x *Response) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *Response) GetFromAccount() string {
	if x != nil {
		return x.FromAccount
	}
	return ""
}

func (x *Response) GetToAccount() string {
	if x != nil {
		return x.ToAccount
	}
	return ""
}

func (x *Response) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Response) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *Response) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type ResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data         []*Response `protobuf:"bytes,1,rep,name=Data,proto3" json:"Data,omitempty"`
	Page         int64       `protobuf:"varint,2,opt,name=Page,proto3" json:"Page,omitempty"`
	PerPage      int64       `protobuf:"varint,3,opt,name=PerPage,proto3" json:"PerPage,omitempty"`
	TotalPage    int64       `protobuf:"varint,4,opt,name=TotalPage,proto3" json:"TotalPage,omitempty"`
	TotalRecords int64       `protobuf:"varint,5,opt,name=TotalRecords,proto3" json:"TotalRecords,omitempty"`
}

func (x *ResponseList) Reset() {
	*x = ResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_transaction_main_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseList) ProtoMessage() {}

func (x *ResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_pb_transaction_main_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseList.ProtoReflect.Descriptor instead.
func (*ResponseList) Descriptor() ([]byte, []int) {
	return file_pb_transaction_main_proto_rawDescGZIP(), []int{3}
}

func (x *ResponseList) GetData() []*Response {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ResponseList) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *ResponseList) GetPerPage() int64 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *ResponseList) GetTotalPage() int64 {
	if x != nil {
		return x.TotalPage
	}
	return 0
}

func (x *ResponseList) GetTotalRecords() int64 {
	if x != nil {
		return x.TotalRecords
	}
	return 0
}

var File_pb_transaction_main_proto protoreflect.FileDescriptor

var file_pb_transaction_main_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61,
	0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x67, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x72, 0x6f,
	0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x46, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x54,
	0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x53, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x50,
	0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0xac, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x44, 0x12, 0x20, 0x0a, 0x0b, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x46, 0x72, 0x6f, 0x6d, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x6f, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa9, 0x01, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x04, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x50, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a,
	0x0c, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0c, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x73, 0x32, 0x8b, 0x01, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x3d, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x74, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3d, 0x0a, 0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x19, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x42,
	0x1c, 0x5a, 0x1a, 0x67, 0x6f, 0x2d, 0x61, 0x70, 0x69, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70,
	0x62, 0x2f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_transaction_main_proto_rawDescOnce sync.Once
	file_pb_transaction_main_proto_rawDescData = file_pb_transaction_main_proto_rawDesc
)

func file_pb_transaction_main_proto_rawDescGZIP() []byte {
	file_pb_transaction_main_proto_rawDescOnce.Do(func() {
		file_pb_transaction_main_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_transaction_main_proto_rawDescData)
	})
	return file_pb_transaction_main_proto_rawDescData
}

var file_pb_transaction_main_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_transaction_main_proto_goTypes = []interface{}{
	(*CreateRequest)(nil), // 0: transaction.CreateRequest
	(*ListRequest)(nil),   // 1: transaction.ListRequest
	(*Response)(nil),      // 2: transaction.Response
	(*ResponseList)(nil),  // 3: transaction.ResponseList
}
var file_pb_transaction_main_proto_depIdxs = []int32{
	2, // 0: transaction.ResponseList.Data:type_name -> transaction.Response
	0, // 1: transaction.Transaction.Create:input_type -> transaction.CreateRequest
	1, // 2: transaction.Transaction.List:input_type -> transaction.ListRequest
	2, // 3: transaction.Transaction.Create:output_type -> transaction.Response
	3, // 4: transaction.Transaction.List:output_type -> transaction.ResponseList
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_transaction_main_proto_init() }
func file_pb_transaction_main_proto_init() {
	if File_pb_transaction_main_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_transaction_main_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_pb_transaction_main_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListRequest); i {
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
		file_pb_transaction_main_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_pb_transaction_main_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResponseList); i {
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
			RawDescriptor: file_pb_transaction_main_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_transaction_main_proto_goTypes,
		DependencyIndexes: file_pb_transaction_main_proto_depIdxs,
		MessageInfos:      file_pb_transaction_main_proto_msgTypes,
	}.Build()
	File_pb_transaction_main_proto = out.File
	file_pb_transaction_main_proto_rawDesc = nil
	file_pb_transaction_main_proto_goTypes = nil
	file_pb_transaction_main_proto_depIdxs = nil
}