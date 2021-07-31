// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.1
// source: proto/v1/fizzbuzz/fizzbuzz.proto

package fizzbuzz

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

type FizzBuzzServiceGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int1  uint64 `protobuf:"varint,1,opt,name=int1,proto3" json:"int1,omitempty"`
	Int2  uint64 `protobuf:"varint,2,opt,name=int2,proto3" json:"int2,omitempty"`
	Limit uint64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Str1  string `protobuf:"bytes,4,opt,name=str1,proto3" json:"str1,omitempty"`
	Str2  string `protobuf:"bytes,5,opt,name=str2,proto3" json:"str2,omitempty"`
}

func (x *FizzBuzzServiceGetRequest) Reset() {
	*x = FizzBuzzServiceGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FizzBuzzServiceGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FizzBuzzServiceGetRequest) ProtoMessage() {}

func (x *FizzBuzzServiceGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FizzBuzzServiceGetRequest.ProtoReflect.Descriptor instead.
func (*FizzBuzzServiceGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDescGZIP(), []int{0}
}

func (x *FizzBuzzServiceGetRequest) GetInt1() uint64 {
	if x != nil {
		return x.Int1
	}
	return 0
}

func (x *FizzBuzzServiceGetRequest) GetInt2() uint64 {
	if x != nil {
		return x.Int2
	}
	return 0
}

func (x *FizzBuzzServiceGetRequest) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FizzBuzzServiceGetRequest) GetStr1() string {
	if x != nil {
		return x.Str1
	}
	return ""
}

func (x *FizzBuzzServiceGetRequest) GetStr2() string {
	if x != nil {
		return x.Str2
	}
	return ""
}

type FizzBuzzServiceGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Words []string `protobuf:"bytes,6,rep,name=words,proto3" json:"words,omitempty"`
}

func (x *FizzBuzzServiceGetResponse) Reset() {
	*x = FizzBuzzServiceGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FizzBuzzServiceGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FizzBuzzServiceGetResponse) ProtoMessage() {}

func (x *FizzBuzzServiceGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FizzBuzzServiceGetResponse.ProtoReflect.Descriptor instead.
func (*FizzBuzzServiceGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDescGZIP(), []int{1}
}

func (x *FizzBuzzServiceGetResponse) GetWords() []string {
	if x != nil {
		return x.Words
	}
	return nil
}

type FizzBuzzServiceStatsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int1     uint64 `protobuf:"varint,1,opt,name=int1,proto3" json:"int1,omitempty"`
	Int2     uint64 `protobuf:"varint,2,opt,name=int2,proto3" json:"int2,omitempty"`
	Limit    uint64 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	Str1     string `protobuf:"bytes,4,opt,name=str1,proto3" json:"str1,omitempty"`
	Str2     string `protobuf:"bytes,5,opt,name=str2,proto3" json:"str2,omitempty"`
	Requests uint64 `protobuf:"varint,8,opt,name=requests,proto3" json:"requests,omitempty"`
}

func (x *FizzBuzzServiceStatsResponse) Reset() {
	*x = FizzBuzzServiceStatsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FizzBuzzServiceStatsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FizzBuzzServiceStatsResponse) ProtoMessage() {}

func (x *FizzBuzzServiceStatsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FizzBuzzServiceStatsResponse.ProtoReflect.Descriptor instead.
func (*FizzBuzzServiceStatsResponse) Descriptor() ([]byte, []int) {
	return file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDescGZIP(), []int{2}
}

func (x *FizzBuzzServiceStatsResponse) GetInt1() uint64 {
	if x != nil {
		return x.Int1
	}
	return 0
}

func (x *FizzBuzzServiceStatsResponse) GetInt2() uint64 {
	if x != nil {
		return x.Int2
	}
	return 0
}

func (x *FizzBuzzServiceStatsResponse) GetLimit() uint64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *FizzBuzzServiceStatsResponse) GetStr1() string {
	if x != nil {
		return x.Str1
	}
	return ""
}

func (x *FizzBuzzServiceStatsResponse) GetStr2() string {
	if x != nil {
		return x.Str2
	}
	return ""
}

func (x *FizzBuzzServiceStatsResponse) GetRequests() uint64 {
	if x != nil {
		return x.Requests
	}
	return 0
}

// Choosing between google.protobuf.Empty or an message without fields,
//The empty message has the advantage of being easier to extend without breakage in the future.
type FizzBuzzServiceStatsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FizzBuzzServiceStatsRequest) Reset() {
	*x = FizzBuzzServiceStatsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FizzBuzzServiceStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FizzBuzzServiceStatsRequest) ProtoMessage() {}

func (x *FizzBuzzServiceStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FizzBuzzServiceStatsRequest.ProtoReflect.Descriptor instead.
func (*FizzBuzzServiceStatsRequest) Descriptor() ([]byte, []int) {
	return file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDescGZIP(), []int{3}
}

var File_proto_v1_fizzbuzz_fizzbuzz_proto protoreflect.FileDescriptor

var file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x69, 0x7a, 0x7a, 0x62,
	0x75, 0x7a, 0x7a, 0x2f, 0x66, 0x69, 0x7a, 0x7a, 0x62, 0x75, 0x7a, 0x7a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x66, 0x69, 0x7a,
	0x7a, 0x62, 0x75, 0x7a, 0x7a, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x19, 0x46, 0x69, 0x7a, 0x7a, 0x42, 0x75, 0x7a, 0x7a,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x04, 0x69, 0x6e, 0x74, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x74, 0x32, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x04, 0x69, 0x6e, 0x74, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d,
	0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x74, 0x72, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73,
	0x74, 0x72, 0x31, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x72, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x73, 0x74, 0x72, 0x32, 0x22, 0x32, 0x0a, 0x1a, 0x46, 0x69, 0x7a, 0x7a, 0x42,
	0x75, 0x7a, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x18, 0x06,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x77, 0x6f, 0x72, 0x64, 0x73, 0x22, 0xa0, 0x01, 0x0a, 0x1c,
	0x46, 0x69, 0x7a, 0x7a, 0x42, 0x75, 0x7a, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x69, 0x6e, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04, 0x69, 0x6e, 0x74, 0x31,
	0x12, 0x12, 0x0a, 0x04, 0x69, 0x6e, 0x74, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x04,
	0x69, 0x6e, 0x74, 0x32, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74,
	0x72, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74, 0x72, 0x31, 0x12, 0x12,
	0x0a, 0x04, 0x73, 0x74, 0x72, 0x32, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x73, 0x74,
	0x72, 0x32, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x1d,
	0x0a, 0x1b, 0x46, 0x69, 0x7a, 0x7a, 0x42, 0x75, 0x7a, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x85, 0x02,
	0x0a, 0x0f, 0x46, 0x69, 0x7a, 0x7a, 0x42, 0x75, 0x7a, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x78, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x76, 0x31, 0x2e, 0x66, 0x69, 0x7a, 0x7a, 0x62, 0x75, 0x7a, 0x7a, 0x2e, 0x46, 0x69, 0x7a,
	0x7a, 0x42, 0x75, 0x7a, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x66, 0x69, 0x7a, 0x7a, 0x62, 0x75, 0x7a, 0x7a, 0x2e, 0x46, 0x69, 0x7a, 0x7a, 0x42,
	0x75, 0x7a, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22, 0x09, 0x2f,
	0x66, 0x69, 0x7a, 0x7a, 0x62, 0x75, 0x7a, 0x7a, 0x3a, 0x01, 0x2a, 0x12, 0x78, 0x0a, 0x05, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x12, 0x2e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x66, 0x69, 0x7a, 0x7a, 0x62, 0x75, 0x7a, 0x7a, 0x2e, 0x46, 0x69, 0x7a, 0x7a, 0x42, 0x75, 0x7a,
	0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x66, 0x69, 0x7a, 0x7a, 0x62, 0x75, 0x7a, 0x7a, 0x2e, 0x46, 0x69, 0x7a, 0x7a, 0x42, 0x75, 0x7a,
	0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x08, 0x12, 0x06, 0x2f,
	0x73, 0x74, 0x61, 0x74, 0x73, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x66, 0x69, 0x7a, 0x7a, 0x62, 0x75, 0x7a, 0x7a, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDescOnce sync.Once
	file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDescData = file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDesc
)

func file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDescGZIP() []byte {
	file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDescOnce.Do(func() {
		file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDescData)
	})
	return file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDescData
}

var file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proto_v1_fizzbuzz_fizzbuzz_proto_goTypes = []interface{}{
	(*FizzBuzzServiceGetRequest)(nil),    // 0: proto.v1.fizzbuzz.FizzBuzzServiceGetRequest
	(*FizzBuzzServiceGetResponse)(nil),   // 1: proto.v1.fizzbuzz.FizzBuzzServiceGetResponse
	(*FizzBuzzServiceStatsResponse)(nil), // 2: proto.v1.fizzbuzz.FizzBuzzServiceStatsResponse
	(*FizzBuzzServiceStatsRequest)(nil),  // 3: proto.v1.fizzbuzz.FizzBuzzServiceStatsRequest
}
var file_proto_v1_fizzbuzz_fizzbuzz_proto_depIdxs = []int32{
	0, // 0: proto.v1.fizzbuzz.FizzBuzzService.Get:input_type -> proto.v1.fizzbuzz.FizzBuzzServiceGetRequest
	3, // 1: proto.v1.fizzbuzz.FizzBuzzService.Stats:input_type -> proto.v1.fizzbuzz.FizzBuzzServiceStatsRequest
	1, // 2: proto.v1.fizzbuzz.FizzBuzzService.Get:output_type -> proto.v1.fizzbuzz.FizzBuzzServiceGetResponse
	2, // 3: proto.v1.fizzbuzz.FizzBuzzService.Stats:output_type -> proto.v1.fizzbuzz.FizzBuzzServiceStatsResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_v1_fizzbuzz_fizzbuzz_proto_init() }
func file_proto_v1_fizzbuzz_fizzbuzz_proto_init() {
	if File_proto_v1_fizzbuzz_fizzbuzz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FizzBuzzServiceGetRequest); i {
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
		file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FizzBuzzServiceGetResponse); i {
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
		file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FizzBuzzServiceStatsResponse); i {
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
		file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FizzBuzzServiceStatsRequest); i {
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
			RawDescriptor: file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_v1_fizzbuzz_fizzbuzz_proto_goTypes,
		DependencyIndexes: file_proto_v1_fizzbuzz_fizzbuzz_proto_depIdxs,
		MessageInfos:      file_proto_v1_fizzbuzz_fizzbuzz_proto_msgTypes,
	}.Build()
	File_proto_v1_fizzbuzz_fizzbuzz_proto = out.File
	file_proto_v1_fizzbuzz_fizzbuzz_proto_rawDesc = nil
	file_proto_v1_fizzbuzz_fizzbuzz_proto_goTypes = nil
	file_proto_v1_fizzbuzz_fizzbuzz_proto_depIdxs = nil
}
