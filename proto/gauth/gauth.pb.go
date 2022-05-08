// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: gauth.proto

package gauth

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

type Migration_Algorithm int32

const (
	Migration_ALGO_INVALID Migration_Algorithm = 0
	Migration_ALGO_SHA1    Migration_Algorithm = 1
)

// Enum value maps for Migration_Algorithm.
var (
	Migration_Algorithm_name = map[int32]string{
		0: "ALGO_INVALID",
		1: "ALGO_SHA1",
	}
	Migration_Algorithm_value = map[string]int32{
		"ALGO_INVALID": 0,
		"ALGO_SHA1":    1,
	}
)

func (x Migration_Algorithm) Enum() *Migration_Algorithm {
	p := new(Migration_Algorithm)
	*p = x
	return p
}

func (x Migration_Algorithm) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Migration_Algorithm) Descriptor() protoreflect.EnumDescriptor {
	return file_gauth_proto_enumTypes[0].Descriptor()
}

func (Migration_Algorithm) Type() protoreflect.EnumType {
	return &file_gauth_proto_enumTypes[0]
}

func (x Migration_Algorithm) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Migration_Algorithm.Descriptor instead.
func (Migration_Algorithm) EnumDescriptor() ([]byte, []int) {
	return file_gauth_proto_rawDescGZIP(), []int{0, 0}
}

type Migration_OtpType int32

const (
	Migration_OTP_INVALID Migration_OtpType = 0
	Migration_OTP_HOTP    Migration_OtpType = 1
	Migration_OTP_TOTP    Migration_OtpType = 2
)

// Enum value maps for Migration_OtpType.
var (
	Migration_OtpType_name = map[int32]string{
		0: "OTP_INVALID",
		1: "OTP_HOTP",
		2: "OTP_TOTP",
	}
	Migration_OtpType_value = map[string]int32{
		"OTP_INVALID": 0,
		"OTP_HOTP":    1,
		"OTP_TOTP":    2,
	}
)

func (x Migration_OtpType) Enum() *Migration_OtpType {
	p := new(Migration_OtpType)
	*p = x
	return p
}

func (x Migration_OtpType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Migration_OtpType) Descriptor() protoreflect.EnumDescriptor {
	return file_gauth_proto_enumTypes[1].Descriptor()
}

func (Migration_OtpType) Type() protoreflect.EnumType {
	return &file_gauth_proto_enumTypes[1]
}

func (x Migration_OtpType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Migration_OtpType.Descriptor instead.
func (Migration_OtpType) EnumDescriptor() ([]byte, []int) {
	return file_gauth_proto_rawDescGZIP(), []int{0, 1}
}

type Migration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OtpParameters []*Migration_OtpParameters `protobuf:"bytes,1,rep,name=otp_parameters,json=otpParameters,proto3" json:"otp_parameters,omitempty"`
	Version       int32                      `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	BatchSize     int32                      `protobuf:"varint,3,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	BatchIndex    int32                      `protobuf:"varint,4,opt,name=batch_index,json=batchIndex,proto3" json:"batch_index,omitempty"`
	BatchId       int32                      `protobuf:"varint,5,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
}

func (x *Migration) Reset() {
	*x = Migration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gauth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Migration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Migration) ProtoMessage() {}

func (x *Migration) ProtoReflect() protoreflect.Message {
	mi := &file_gauth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Migration.ProtoReflect.Descriptor instead.
func (*Migration) Descriptor() ([]byte, []int) {
	return file_gauth_proto_rawDescGZIP(), []int{0}
}

func (x *Migration) GetOtpParameters() []*Migration_OtpParameters {
	if x != nil {
		return x.OtpParameters
	}
	return nil
}

func (x *Migration) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *Migration) GetBatchSize() int32 {
	if x != nil {
		return x.BatchSize
	}
	return 0
}

func (x *Migration) GetBatchIndex() int32 {
	if x != nil {
		return x.BatchIndex
	}
	return 0
}

func (x *Migration) GetBatchId() int32 {
	if x != nil {
		return x.BatchId
	}
	return 0
}

type Migration_OtpParameters struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Secret    []byte              `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	Name      string              `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Issuer    string              `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Algorithm Migration_Algorithm `protobuf:"varint,4,opt,name=algorithm,proto3,enum=gauth.Migration_Algorithm" json:"algorithm,omitempty"`
	Digits    int32               `protobuf:"varint,5,opt,name=digits,proto3" json:"digits,omitempty"`
	Type      Migration_OtpType   `protobuf:"varint,6,opt,name=type,proto3,enum=gauth.Migration_OtpType" json:"type,omitempty"`
	Counter   int64               `protobuf:"varint,7,opt,name=counter,proto3" json:"counter,omitempty"`
}

func (x *Migration_OtpParameters) Reset() {
	*x = Migration_OtpParameters{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gauth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Migration_OtpParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Migration_OtpParameters) ProtoMessage() {}

func (x *Migration_OtpParameters) ProtoReflect() protoreflect.Message {
	mi := &file_gauth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Migration_OtpParameters.ProtoReflect.Descriptor instead.
func (*Migration_OtpParameters) Descriptor() ([]byte, []int) {
	return file_gauth_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Migration_OtpParameters) GetSecret() []byte {
	if x != nil {
		return x.Secret
	}
	return nil
}

func (x *Migration_OtpParameters) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Migration_OtpParameters) GetIssuer() string {
	if x != nil {
		return x.Issuer
	}
	return ""
}

func (x *Migration_OtpParameters) GetAlgorithm() Migration_Algorithm {
	if x != nil {
		return x.Algorithm
	}
	return Migration_ALGO_INVALID
}

func (x *Migration_OtpParameters) GetDigits() int32 {
	if x != nil {
		return x.Digits
	}
	return 0
}

func (x *Migration_OtpParameters) GetType() Migration_OtpType {
	if x != nil {
		return x.Type
	}
	return Migration_OTP_INVALID
}

func (x *Migration_OtpParameters) GetCounter() int64 {
	if x != nil {
		return x.Counter
	}
	return 0
}

var File_gauth_proto protoreflect.FileDescriptor

var file_gauth_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67,
	0x61, 0x75, 0x74, 0x68, 0x22, 0x9d, 0x04, 0x0a, 0x09, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x45, 0x0a, 0x0e, 0x6f, 0x74, 0x70, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65,
	0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x67, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x74, 0x70,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x0d, 0x6f, 0x74, 0x70, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x69,
	0x7a, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x6e, 0x64, 0x65,
	0x78, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x6e,
	0x64, 0x65, 0x78, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x1a, 0xed,
	0x01, 0x0a, 0x0d, 0x4f, 0x74, 0x70, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x06, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x69, 0x73, 0x73, 0x75, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68,
	0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x67, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x4d, 0x69, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69,
	0x74, 0x68, 0x6d, 0x52, 0x09, 0x61, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x64, 0x69, 0x67, 0x69, 0x74, 0x73, 0x12, 0x2c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x67, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4d, 0x69, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4f, 0x74, 0x70, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x22, 0x2c,
	0x0a, 0x09, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d, 0x12, 0x10, 0x0a, 0x0c, 0x41,
	0x4c, 0x47, 0x4f, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0d, 0x0a,
	0x09, 0x41, 0x4c, 0x47, 0x4f, 0x5f, 0x53, 0x48, 0x41, 0x31, 0x10, 0x01, 0x22, 0x36, 0x0a, 0x07,
	0x4f, 0x74, 0x70, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x4f, 0x54, 0x50, 0x5f, 0x49,
	0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x54, 0x50, 0x5f,
	0x48, 0x4f, 0x54, 0x50, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x54, 0x50, 0x5f, 0x54, 0x4f,
	0x54, 0x50, 0x10, 0x02, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61,
	0x75, 0x74, 0x68, 0x3b, 0x67, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_gauth_proto_rawDescOnce sync.Once
	file_gauth_proto_rawDescData = file_gauth_proto_rawDesc
)

func file_gauth_proto_rawDescGZIP() []byte {
	file_gauth_proto_rawDescOnce.Do(func() {
		file_gauth_proto_rawDescData = protoimpl.X.CompressGZIP(file_gauth_proto_rawDescData)
	})
	return file_gauth_proto_rawDescData
}

var file_gauth_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_gauth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_gauth_proto_goTypes = []interface{}{
	(Migration_Algorithm)(0),        // 0: gauth.Migration.Algorithm
	(Migration_OtpType)(0),          // 1: gauth.Migration.OtpType
	(*Migration)(nil),               // 2: gauth.Migration
	(*Migration_OtpParameters)(nil), // 3: gauth.Migration.OtpParameters
}
var file_gauth_proto_depIdxs = []int32{
	3, // 0: gauth.Migration.otp_parameters:type_name -> gauth.Migration.OtpParameters
	0, // 1: gauth.Migration.OtpParameters.algorithm:type_name -> gauth.Migration.Algorithm
	1, // 2: gauth.Migration.OtpParameters.type:type_name -> gauth.Migration.OtpType
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_gauth_proto_init() }
func file_gauth_proto_init() {
	if File_gauth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gauth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Migration); i {
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
		file_gauth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Migration_OtpParameters); i {
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
			RawDescriptor: file_gauth_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gauth_proto_goTypes,
		DependencyIndexes: file_gauth_proto_depIdxs,
		EnumInfos:         file_gauth_proto_enumTypes,
		MessageInfos:      file_gauth_proto_msgTypes,
	}.Build()
	File_gauth_proto = out.File
	file_gauth_proto_rawDesc = nil
	file_gauth_proto_goTypes = nil
	file_gauth_proto_depIdxs = nil
}
