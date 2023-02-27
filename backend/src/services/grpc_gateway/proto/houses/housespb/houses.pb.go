// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: houses/houses.proto

package housespb

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

type CreateHouseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=requestId,proto3" json:"requestId,omitempty"`
	House     *House `protobuf:"bytes,2,opt,name=house,proto3" json:"house,omitempty"`
}

func (x *CreateHouseRequest) Reset() {
	*x = CreateHouseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_houses_houses_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHouseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHouseRequest) ProtoMessage() {}

func (x *CreateHouseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_houses_houses_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHouseRequest.ProtoReflect.Descriptor instead.
func (*CreateHouseRequest) Descriptor() ([]byte, []int) {
	return file_houses_houses_proto_rawDescGZIP(), []int{0}
}

func (x *CreateHouseRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *CreateHouseRequest) GetHouse() *House {
	if x != nil {
		return x.House
	}
	return nil
}

type House struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *House) Reset() {
	*x = House{}
	if protoimpl.UnsafeEnabled {
		mi := &file_houses_houses_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *House) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*House) ProtoMessage() {}

func (x *House) ProtoReflect() protoreflect.Message {
	mi := &file_houses_houses_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use House.ProtoReflect.Descriptor instead.
func (*House) Descriptor() ([]byte, []int) {
	return file_houses_houses_proto_rawDescGZIP(), []int{1}
}

func (x *House) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_houses_houses_proto protoreflect.FileDescriptor

var file_houses_houses_proto_rawDesc = []byte{
	0x0a, 0x13, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x2f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x22, 0x57, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49,
	0x64, 0x12, 0x23, 0x0a, 0x05, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x2e, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52,
	0x05, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x22, 0x17, 0x0a, 0x05, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x32,
	0x44, 0x0a, 0x06, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0b, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x12, 0x1a, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x6f, 0x75, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x2e, 0x48, 0x6f,
	0x75, 0x73, 0x65, 0x22, 0x00, 0x42, 0x33, 0x5a, 0x31, 0x74, 0x75, 0x68, 0x6c, 0x61, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x68, 0x6f, 0x75, 0x73, 0x65,
	0x73, 0x2f, 0x68, 0x6f, 0x75, 0x73, 0x65, 0x73, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_houses_houses_proto_rawDescOnce sync.Once
	file_houses_houses_proto_rawDescData = file_houses_houses_proto_rawDesc
)

func file_houses_houses_proto_rawDescGZIP() []byte {
	file_houses_houses_proto_rawDescOnce.Do(func() {
		file_houses_houses_proto_rawDescData = protoimpl.X.CompressGZIP(file_houses_houses_proto_rawDescData)
	})
	return file_houses_houses_proto_rawDescData
}

var file_houses_houses_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_houses_houses_proto_goTypes = []interface{}{
	(*CreateHouseRequest)(nil), // 0: houses.CreateHouseRequest
	(*House)(nil),              // 1: houses.House
}
var file_houses_houses_proto_depIdxs = []int32{
	1, // 0: houses.CreateHouseRequest.house:type_name -> houses.House
	0, // 1: houses.Houses.CreateHouse:input_type -> houses.CreateHouseRequest
	1, // 2: houses.Houses.CreateHouse:output_type -> houses.House
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_houses_houses_proto_init() }
func file_houses_houses_proto_init() {
	if File_houses_houses_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_houses_houses_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHouseRequest); i {
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
		file_houses_houses_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*House); i {
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
			RawDescriptor: file_houses_houses_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_houses_houses_proto_goTypes,
		DependencyIndexes: file_houses_houses_proto_depIdxs,
		MessageInfos:      file_houses_houses_proto_msgTypes,
	}.Build()
	File_houses_houses_proto = out.File
	file_houses_houses_proto_rawDesc = nil
	file_houses_houses_proto_goTypes = nil
	file_houses_houses_proto_depIdxs = nil
}
