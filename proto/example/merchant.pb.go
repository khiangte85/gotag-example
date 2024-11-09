// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: proto/example/merchant.proto

package modelpb

import (
	_ "github.com/srikrsna/protoc-gen-gotag/tagger"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MerchantStatus int32

const (
	MerchantStatus_ACTIVE    MerchantStatus = 0
	MerchantStatus_INACTIVE  MerchantStatus = 1
	MerchantStatus_SUSPENDED MerchantStatus = 2
	MerchantStatus_DELETED   MerchantStatus = 3
)

// Enum value maps for MerchantStatus.
var (
	MerchantStatus_name = map[int32]string{
		0: "ACTIVE",
		1: "INACTIVE",
		2: "SUSPENDED",
		3: "DELETED",
	}
	MerchantStatus_value = map[string]int32{
		"ACTIVE":    0,
		"INACTIVE":  1,
		"SUSPENDED": 2,
		"DELETED":   3,
	}
)

func (x MerchantStatus) Enum() *MerchantStatus {
	p := new(MerchantStatus)
	*p = x
	return p
}

func (x MerchantStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MerchantStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_example_merchant_proto_enumTypes[0].Descriptor()
}

func (MerchantStatus) Type() protoreflect.EnumType {
	return &file_proto_example_merchant_proto_enumTypes[0]
}

func (x MerchantStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MerchantStatus.Descriptor instead.
func (MerchantStatus) EnumDescriptor() ([]byte, []int) {
	return file_proto_example_merchant_proto_rawDescGZIP(), []int{0}
}

type Merchant struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// UUID of the merchant
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// UUID of the parent merchant
	ParentId string `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Email    string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	PhoneNo  string `protobuf:"bytes,6,opt,name=phone_no,json=phoneNo,proto3" json:"phone_no,omitempty"`
	// Is the merchant the owner of business, if true then parent_id is null
	// otherwise parent_id is the id of the owner merchant
	IsOwner   bool                   `protobuf:"varint,7,opt,name=is_owner,json=isOwner,proto3" json:"is_owner,omitempty"`
	Status    MerchantStatus         `protobuf:"varint,8,opt,name=status,proto3,enum=example.MerchantStatus" json:"status,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt *timestamppb.Timestamp `protobuf:"bytes,11,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Merchant) Reset() {
	*x = Merchant{}
	mi := &file_proto_example_merchant_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Merchant) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Merchant) ProtoMessage() {}

func (x *Merchant) ProtoReflect() protoreflect.Message {
	mi := &file_proto_example_merchant_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Merchant.ProtoReflect.Descriptor instead.
func (*Merchant) Descriptor() ([]byte, []int) {
	return file_proto_example_merchant_proto_rawDescGZIP(), []int{0}
}

func (x *Merchant) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Merchant) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

func (x *Merchant) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Merchant) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Merchant) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Merchant) GetPhoneNo() string {
	if x != nil {
		return x.PhoneNo
	}
	return ""
}

func (x *Merchant) GetIsOwner() bool {
	if x != nil {
		return x.IsOwner
	}
	return false
}

func (x *Merchant) GetStatus() MerchantStatus {
	if x != nil {
		return x.Status
	}
	return MerchantStatus_ACTIVE
}

func (x *Merchant) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Merchant) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

func (x *Merchant) GetDeletedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.DeletedAt
	}
	return nil
}

var File_proto_example_merchant_proto protoreflect.FileDescriptor

var file_proto_example_merchant_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f,
	0x6d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x2f, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbe, 0x03,
	0x0a, 0x08, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x09, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0x9a,
	0x84, 0x9e, 0x03, 0x0e, 0x64, 0x62, 0x3a, 0x22, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x22, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x6f, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0x9a, 0x84, 0x9e, 0x03, 0x0d, 0x64, 0x62, 0x3a, 0x22, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x6f, 0x22, 0x52, 0x07, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e,
	0x6f, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2f, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x39, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x12, 0x39, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x2a, 0x46,
	0x0a, 0x0e, 0x4d, 0x65, 0x72, 0x63, 0x68, 0x61, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x0a, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08,
	0x49, 0x4e, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x55,
	0x53, 0x50, 0x45, 0x4e, 0x44, 0x45, 0x44, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x44, 0x10, 0x03, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x69, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x74, 0x72, 0x65, 0x70,
	0x69, 0x73, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2f, 0x70, 0x62, 0x3b, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_example_merchant_proto_rawDescOnce sync.Once
	file_proto_example_merchant_proto_rawDescData = file_proto_example_merchant_proto_rawDesc
)

func file_proto_example_merchant_proto_rawDescGZIP() []byte {
	file_proto_example_merchant_proto_rawDescOnce.Do(func() {
		file_proto_example_merchant_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_example_merchant_proto_rawDescData)
	})
	return file_proto_example_merchant_proto_rawDescData
}

var file_proto_example_merchant_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_example_merchant_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_example_merchant_proto_goTypes = []any{
	(MerchantStatus)(0),           // 0: example.MerchantStatus
	(*Merchant)(nil),              // 1: example.Merchant
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_proto_example_merchant_proto_depIdxs = []int32{
	0, // 0: example.Merchant.status:type_name -> example.MerchantStatus
	2, // 1: example.Merchant.created_at:type_name -> google.protobuf.Timestamp
	2, // 2: example.Merchant.updated_at:type_name -> google.protobuf.Timestamp
	2, // 3: example.Merchant.deleted_at:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_proto_example_merchant_proto_init() }
func file_proto_example_merchant_proto_init() {
	if File_proto_example_merchant_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_example_merchant_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_example_merchant_proto_goTypes,
		DependencyIndexes: file_proto_example_merchant_proto_depIdxs,
		EnumInfos:         file_proto_example_merchant_proto_enumTypes,
		MessageInfos:      file_proto_example_merchant_proto_msgTypes,
	}.Build()
	File_proto_example_merchant_proto = out.File
	file_proto_example_merchant_proto_rawDesc = nil
	file_proto_example_merchant_proto_goTypes = nil
	file_proto_example_merchant_proto_depIdxs = nil
}