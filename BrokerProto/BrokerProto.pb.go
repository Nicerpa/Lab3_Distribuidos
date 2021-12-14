// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: BrokerProto/BrokerProto.proto

package Lab3

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

type LeiaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet string  `protobuf:"bytes,1,opt,name=planet,proto3" json:"planet,omitempty"`
	City   string  `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Ip     string  `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
	Clock  []int32 `protobuf:"varint,4,rep,packed,name=clock,proto3" json:"clock,omitempty"`
}

func (x *LeiaReq) Reset() {
	*x = LeiaReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BrokerProto_BrokerProto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LeiaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LeiaReq) ProtoMessage() {}

func (x *LeiaReq) ProtoReflect() protoreflect.Message {
	mi := &file_BrokerProto_BrokerProto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LeiaReq.ProtoReflect.Descriptor instead.
func (*LeiaReq) Descriptor() ([]byte, []int) {
	return file_BrokerProto_BrokerProto_proto_rawDescGZIP(), []int{0}
}

func (x *LeiaReq) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

func (x *LeiaReq) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *LeiaReq) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *LeiaReq) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

type Rebelds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rebelds int32   `protobuf:"varint,1,opt,name=rebelds,proto3" json:"rebelds,omitempty"`
	Ip      string  `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Clock   []int32 `protobuf:"varint,3,rep,packed,name=clock,proto3" json:"clock,omitempty"`
}

func (x *Rebelds) Reset() {
	*x = Rebelds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BrokerProto_BrokerProto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rebelds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rebelds) ProtoMessage() {}

func (x *Rebelds) ProtoReflect() protoreflect.Message {
	mi := &file_BrokerProto_BrokerProto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rebelds.ProtoReflect.Descriptor instead.
func (*Rebelds) Descriptor() ([]byte, []int) {
	return file_BrokerProto_BrokerProto_proto_rawDescGZIP(), []int{1}
}

func (x *Rebelds) GetRebelds() int32 {
	if x != nil {
		return x.Rebelds
	}
	return 0
}

func (x *Rebelds) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *Rebelds) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

type CityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet string  `protobuf:"bytes,1,opt,name=planet,proto3" json:"planet,omitempty"`
	Clock  []int32 `protobuf:"varint,2,rep,packed,name=clock,proto3" json:"clock,omitempty"`
	Ip     string  `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *CityData) Reset() {
	*x = CityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BrokerProto_BrokerProto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityData) ProtoMessage() {}

func (x *CityData) ProtoReflect() protoreflect.Message {
	mi := &file_BrokerProto_BrokerProto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityData.ProtoReflect.Descriptor instead.
func (*CityData) Descriptor() ([]byte, []int) {
	return file_BrokerProto_BrokerProto_proto_rawDescGZIP(), []int{2}
}

func (x *CityData) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

func (x *CityData) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

func (x *CityData) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type CityRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IpDes string `protobuf:"bytes,1,opt,name=ip_des,json=ipDes,proto3" json:"ip_des,omitempty"`
}

func (x *CityRes) Reset() {
	*x = CityRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_BrokerProto_BrokerProto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityRes) ProtoMessage() {}

func (x *CityRes) ProtoReflect() protoreflect.Message {
	mi := &file_BrokerProto_BrokerProto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityRes.ProtoReflect.Descriptor instead.
func (*CityRes) Descriptor() ([]byte, []int) {
	return file_BrokerProto_BrokerProto_proto_rawDescGZIP(), []int{3}
}

func (x *CityRes) GetIpDes() string {
	if x != nil {
		return x.IpDes
	}
	return ""
}

var File_BrokerProto_BrokerProto_proto protoreflect.FileDescriptor

var file_BrokerProto_BrokerProto_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5b, 0x0a, 0x07,
	0x4c, 0x65, 0x69, 0x61, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x49, 0x0a, 0x07, 0x52, 0x65, 0x62,
	0x65, 0x6c, 0x64, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x48, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x20,
	0x0a, 0x07, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x70, 0x5f,
	0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x70, 0x44, 0x65, 0x73,
	0x32, 0x8c, 0x01, 0x0a, 0x0c, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x40, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x62, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x14, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x65, 0x69, 0x61, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x42, 0x72,
	0x6f, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64,
	0x73, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x49, 0x50, 0x43, 0x69, 0x74, 0x79,
	0x12, 0x15, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x69, 0x74, 0x79, 0x44, 0x61, 0x74, 0x61, 0x1a, 0x14, 0x2e, 0x42, 0x72, 0x6f, 0x6b, 0x65, 0x72,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42,
	0x11, 0x5a, 0x0f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61,
	0x62, 0x33, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_BrokerProto_BrokerProto_proto_rawDescOnce sync.Once
	file_BrokerProto_BrokerProto_proto_rawDescData = file_BrokerProto_BrokerProto_proto_rawDesc
)

func file_BrokerProto_BrokerProto_proto_rawDescGZIP() []byte {
	file_BrokerProto_BrokerProto_proto_rawDescOnce.Do(func() {
		file_BrokerProto_BrokerProto_proto_rawDescData = protoimpl.X.CompressGZIP(file_BrokerProto_BrokerProto_proto_rawDescData)
	})
	return file_BrokerProto_BrokerProto_proto_rawDescData
}

var file_BrokerProto_BrokerProto_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_BrokerProto_BrokerProto_proto_goTypes = []interface{}{
	(*LeiaReq)(nil),  // 0: BrokerProto.LeiaReq
	(*Rebelds)(nil),  // 1: BrokerProto.Rebelds
	(*CityData)(nil), // 2: BrokerProto.CityData
	(*CityRes)(nil),  // 3: BrokerProto.CityRes
}
var file_BrokerProto_BrokerProto_proto_depIdxs = []int32{
	0, // 0: BrokerProto.BrokerServer.GetNumberRebelds:input_type -> BrokerProto.LeiaReq
	2, // 1: BrokerProto.BrokerServer.GetIPCity:input_type -> BrokerProto.CityData
	1, // 2: BrokerProto.BrokerServer.GetNumberRebelds:output_type -> BrokerProto.Rebelds
	3, // 3: BrokerProto.BrokerServer.GetIPCity:output_type -> BrokerProto.CityRes
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_BrokerProto_BrokerProto_proto_init() }
func file_BrokerProto_BrokerProto_proto_init() {
	if File_BrokerProto_BrokerProto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_BrokerProto_BrokerProto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LeiaReq); i {
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
		file_BrokerProto_BrokerProto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rebelds); i {
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
		file_BrokerProto_BrokerProto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityData); i {
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
		file_BrokerProto_BrokerProto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityRes); i {
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
			RawDescriptor: file_BrokerProto_BrokerProto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_BrokerProto_BrokerProto_proto_goTypes,
		DependencyIndexes: file_BrokerProto_BrokerProto_proto_depIdxs,
		MessageInfos:      file_BrokerProto_BrokerProto_proto_msgTypes,
	}.Build()
	File_BrokerProto_BrokerProto_proto = out.File
	file_BrokerProto_BrokerProto_proto_rawDesc = nil
	file_BrokerProto_BrokerProto_proto_goTypes = nil
	file_BrokerProto_BrokerProto_proto_depIdxs = nil
}
