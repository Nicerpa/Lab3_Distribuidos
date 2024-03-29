// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: FulcrumProto/FulcrumProto.proto

package Lab3

import (
	empty "github.com/golang/protobuf/ptypes/empty"
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

type NewData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File   []string `protobuf:"bytes,1,rep,name=file,proto3" json:"file,omitempty"`
	Clock  []int32  `protobuf:"varint,2,rep,packed,name=clock,proto3" json:"clock,omitempty"`
	Planet string   `protobuf:"bytes,3,opt,name=planet,proto3" json:"planet,omitempty"`
}

func (x *NewData) Reset() {
	*x = NewData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewData) ProtoMessage() {}

func (x *NewData) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewData.ProtoReflect.Descriptor instead.
func (*NewData) Descriptor() ([]byte, []int) {
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{0}
}

func (x *NewData) GetFile() []string {
	if x != nil {
		return x.File
	}
	return nil
}

func (x *NewData) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

func (x *NewData) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusFlag int32 `protobuf:"varint,1,opt,name=statusFlag,proto3" json:"statusFlag,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{1}
}

func (x *Status) GetStatusFlag() int32 {
	if x != nil {
		return x.StatusFlag
	}
	return 0
}

type LogReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *LogReq) Reset() {
	*x = LogReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogReq) ProtoMessage() {}

func (x *LogReq) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogReq.ProtoReflect.Descriptor instead.
func (*LogReq) Descriptor() ([]byte, []int) {
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{2}
}

func (x *LogReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LogFile []string `protobuf:"bytes,1,rep,name=logFile,proto3" json:"logFile,omitempty"`
	Clock   []int32  `protobuf:"varint,2,rep,packed,name=clock,proto3" json:"clock,omitempty"`
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{3}
}

func (x *Log) GetLogFile() []string {
	if x != nil {
		return x.LogFile
	}
	return nil
}

func (x *Log) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

type PlanetList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []string `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *PlanetList) Reset() {
	*x = PlanetList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanetList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanetList) ProtoMessage() {}

func (x *PlanetList) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanetList.ProtoReflect.Descriptor instead.
func (*PlanetList) Descriptor() ([]byte, []int) {
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{4}
}

func (x *PlanetList) GetList() []string {
	if x != nil {
		return x.List
	}
	return nil
}

type PlanetData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet string `protobuf:"bytes,1,opt,name=planet,proto3" json:"planet,omitempty"`
}

func (x *PlanetData) Reset() {
	*x = PlanetData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanetData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanetData) ProtoMessage() {}

func (x *PlanetData) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanetData.ProtoReflect.Descriptor instead.
func (*PlanetData) Descriptor() ([]byte, []int) {
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{5}
}

func (x *PlanetData) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

type PlanetClock struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock []int32 `protobuf:"varint,1,rep,packed,name=clock,proto3" json:"clock,omitempty"`
}

func (x *PlanetClock) Reset() {
	*x = PlanetClock{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlanetClock) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlanetClock) ProtoMessage() {}

func (x *PlanetClock) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlanetClock.ProtoReflect.Descriptor instead.
func (*PlanetClock) Descriptor() ([]byte, []int) {
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{6}
}

func (x *PlanetClock) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

type CityData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet string `protobuf:"bytes,1,opt,name=planet,proto3" json:"planet,omitempty"`
	City   string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *CityData) Reset() {
	*x = CityData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityData) ProtoMessage() {}

func (x *CityData) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[7]
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
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{7}
}

func (x *CityData) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

func (x *CityData) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type CityRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock   []int32 `protobuf:"varint,1,rep,packed,name=clock,proto3" json:"clock,omitempty"`
	Rebelds int32   `protobuf:"varint,2,opt,name=rebelds,proto3" json:"rebelds,omitempty"`
}

func (x *CityRes) Reset() {
	*x = CityRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityRes) ProtoMessage() {}

func (x *CityRes) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[8]
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
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{8}
}

func (x *CityRes) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

func (x *CityRes) GetRebelds() int32 {
	if x != nil {
		return x.Rebelds
	}
	return 0
}

type NewCity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet  string `protobuf:"bytes,1,opt,name=planet,proto3" json:"planet,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Rebelds int32  `protobuf:"varint,3,opt,name=rebelds,proto3" json:"rebelds,omitempty"` // 0 si no se especifica
}

func (x *NewCity) Reset() {
	*x = NewCity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewCity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewCity) ProtoMessage() {}

func (x *NewCity) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewCity.ProtoReflect.Descriptor instead.
func (*NewCity) Descriptor() ([]byte, []int) {
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{9}
}

func (x *NewCity) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

func (x *NewCity) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *NewCity) GetRebelds() int32 {
	if x != nil {
		return x.Rebelds
	}
	return 0
}

type DelCity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet string `protobuf:"bytes,1,opt,name=planet,proto3" json:"planet,omitempty"`
	City   string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
}

func (x *DelCity) Reset() {
	*x = DelCity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelCity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelCity) ProtoMessage() {}

func (x *DelCity) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelCity.ProtoReflect.Descriptor instead.
func (*DelCity) Descriptor() ([]byte, []int) {
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{10}
}

func (x *DelCity) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

func (x *DelCity) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type UpCity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planet  string `protobuf:"bytes,1,opt,name=planet,proto3" json:"planet,omitempty"`
	City    string `protobuf:"bytes,2,opt,name=city,proto3" json:"city,omitempty"`
	Flag    int32  `protobuf:"varint,3,opt,name=flag,proto3" json:"flag,omitempty"`      // 1 -> modificar nombre; 0 -> modificar cantidad
	Newname string `protobuf:"bytes,4,opt,name=newname,proto3" json:"newname,omitempty"` // flag = 0 -> default value ""
	Num     int32  `protobuf:"varint,5,opt,name=num,proto3" json:"num,omitempty"`        // flag = 1 -> default value -1
}

func (x *UpCity) Reset() {
	*x = UpCity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpCity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpCity) ProtoMessage() {}

func (x *UpCity) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpCity.ProtoReflect.Descriptor instead.
func (*UpCity) Descriptor() ([]byte, []int) {
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{11}
}

func (x *UpCity) GetPlanet() string {
	if x != nil {
		return x.Planet
	}
	return ""
}

func (x *UpCity) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *UpCity) GetFlag() int32 {
	if x != nil {
		return x.Flag
	}
	return 0
}

func (x *UpCity) GetNewname() string {
	if x != nil {
		return x.Newname
	}
	return ""
}

func (x *UpCity) GetNum() int32 {
	if x != nil {
		return x.Num
	}
	return 0
}

type CityStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Clock []int32 `protobuf:"varint,1,rep,packed,name=clock,proto3" json:"clock,omitempty"`
}

func (x *CityStatus) Reset() {
	*x = CityStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CityStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CityStatus) ProtoMessage() {}

func (x *CityStatus) ProtoReflect() protoreflect.Message {
	mi := &file_FulcrumProto_FulcrumProto_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CityStatus.ProtoReflect.Descriptor instead.
func (*CityStatus) Descriptor() ([]byte, []int) {
	return file_FulcrumProto_FulcrumProto_proto_rawDescGZIP(), []int{12}
}

func (x *CityStatus) GetClock() []int32 {
	if x != nil {
		return x.Clock
	}
	return nil
}

var File_FulcrumProto_FulcrumProto_proto protoreflect.FileDescriptor

var file_FulcrumProto_FulcrumProto_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x46,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4b, 0x0a, 0x07,
	0x4e, 0x65, 0x77, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x22, 0x28, 0x0a, 0x06, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x6c, 0x61,
	0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46,
	0x6c, 0x61, 0x67, 0x22, 0x1c, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x35, 0x0a, 0x03, 0x4c, 0x6f, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x6f, 0x67, 0x46,
	0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x20, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x24, 0x0a, 0x0a, 0x50, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74,
	0x22, 0x23, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05,
	0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x36, 0x0a, 0x08, 0x43, 0x69, 0x74, 0x79, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x22, 0x39, 0x0a,
	0x07, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63,
	0x6b, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x4f, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x43,
	0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12,
	0x18, 0x0a, 0x07, 0x72, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x72, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x22, 0x35, 0x0a, 0x07, 0x44, 0x65, 0x6c,
	0x43, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x69, 0x74, 0x79,
	0x22, 0x74, 0x0a, 0x06, 0x55, 0x70, 0x43, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6c, 0x61, 0x6e,
	0x65, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x69, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x66, 0x6c, 0x61, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65,
	0x77, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6e, 0x75, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x03, 0x6e, 0x75, 0x6d, 0x22, 0x22, 0x0a, 0x0a, 0x43, 0x69, 0x74, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x63, 0x6b, 0x32, 0x8c, 0x04, 0x0a, 0x0d, 0x46,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x08,
	0x47, 0x65, 0x74, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x18, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72,
	0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x44, 0x61,
	0x74, 0x61, 0x1a, 0x19, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x43, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x00, 0x12,
	0x3c, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x46,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x69, 0x74, 0x79,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x15, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3c, 0x0a,
	0x07, 0x41, 0x64, 0x64, 0x43, 0x69, 0x74, 0x79, 0x12, 0x15, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72,
	0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x43, 0x69, 0x74, 0x79, 0x1a,
	0x18, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x12, 0x15, 0x2e, 0x46, 0x75, 0x6c, 0x63,
	0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c, 0x43, 0x69, 0x74, 0x79,
	0x1a, 0x18, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x0a,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x12, 0x14, 0x2e, 0x46, 0x75, 0x6c,
	0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x43, 0x69, 0x74, 0x79,
	0x1a, 0x18, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x43, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x11,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x18, 0x2e, 0x46, 0x75, 0x6c, 0x63,
	0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x4c,
	0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0a, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x4c, 0x6f, 0x67, 0x12, 0x14, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x46, 0x75, 0x6c, 0x63,
	0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4c, 0x6f, 0x67, 0x22, 0x00, 0x12, 0x39,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x46,
	0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x44,
	0x61, 0x74, 0x61, 0x1a, 0x14, 0x2e, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x11, 0x5a, 0x0f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4c, 0x61, 0x62, 0x33, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FulcrumProto_FulcrumProto_proto_rawDescOnce sync.Once
	file_FulcrumProto_FulcrumProto_proto_rawDescData = file_FulcrumProto_FulcrumProto_proto_rawDesc
)

func file_FulcrumProto_FulcrumProto_proto_rawDescGZIP() []byte {
	file_FulcrumProto_FulcrumProto_proto_rawDescOnce.Do(func() {
		file_FulcrumProto_FulcrumProto_proto_rawDescData = protoimpl.X.CompressGZIP(file_FulcrumProto_FulcrumProto_proto_rawDescData)
	})
	return file_FulcrumProto_FulcrumProto_proto_rawDescData
}

var file_FulcrumProto_FulcrumProto_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_FulcrumProto_FulcrumProto_proto_goTypes = []interface{}{
	(*NewData)(nil),     // 0: FulcrumProto.NewData
	(*Status)(nil),      // 1: FulcrumProto.Status
	(*LogReq)(nil),      // 2: FulcrumProto.LogReq
	(*Log)(nil),         // 3: FulcrumProto.Log
	(*PlanetList)(nil),  // 4: FulcrumProto.PlanetList
	(*PlanetData)(nil),  // 5: FulcrumProto.PlanetData
	(*PlanetClock)(nil), // 6: FulcrumProto.PlanetClock
	(*CityData)(nil),    // 7: FulcrumProto.CityData
	(*CityRes)(nil),     // 8: FulcrumProto.CityRes
	(*NewCity)(nil),     // 9: FulcrumProto.NewCity
	(*DelCity)(nil),     // 10: FulcrumProto.DelCity
	(*UpCity)(nil),      // 11: FulcrumProto.UpCity
	(*CityStatus)(nil),  // 12: FulcrumProto.CityStatus
	(*empty.Empty)(nil), // 13: google.protobuf.Empty
}
var file_FulcrumProto_FulcrumProto_proto_depIdxs = []int32{
	5,  // 0: FulcrumProto.FulcrumServer.GetClock:input_type -> FulcrumProto.PlanetData
	7,  // 1: FulcrumProto.FulcrumServer.GetNumber:input_type -> FulcrumProto.CityData
	9,  // 2: FulcrumProto.FulcrumServer.AddCity:input_type -> FulcrumProto.NewCity
	10, // 3: FulcrumProto.FulcrumServer.DeleteCity:input_type -> FulcrumProto.DelCity
	11, // 4: FulcrumProto.FulcrumServer.UpdateCity:input_type -> FulcrumProto.UpCity
	13, // 5: FulcrumProto.FulcrumServer.RequestPlanetList:input_type -> google.protobuf.Empty
	2,  // 6: FulcrumProto.FulcrumServer.RequestLog:input_type -> FulcrumProto.LogReq
	0,  // 7: FulcrumProto.FulcrumServer.UpdateFile:input_type -> FulcrumProto.NewData
	6,  // 8: FulcrumProto.FulcrumServer.GetClock:output_type -> FulcrumProto.PlanetClock
	8,  // 9: FulcrumProto.FulcrumServer.GetNumber:output_type -> FulcrumProto.CityRes
	12, // 10: FulcrumProto.FulcrumServer.AddCity:output_type -> FulcrumProto.CityStatus
	12, // 11: FulcrumProto.FulcrumServer.DeleteCity:output_type -> FulcrumProto.CityStatus
	12, // 12: FulcrumProto.FulcrumServer.UpdateCity:output_type -> FulcrumProto.CityStatus
	4,  // 13: FulcrumProto.FulcrumServer.RequestPlanetList:output_type -> FulcrumProto.PlanetList
	3,  // 14: FulcrumProto.FulcrumServer.RequestLog:output_type -> FulcrumProto.Log
	1,  // 15: FulcrumProto.FulcrumServer.UpdateFile:output_type -> FulcrumProto.Status
	8,  // [8:16] is the sub-list for method output_type
	0,  // [0:8] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_FulcrumProto_FulcrumProto_proto_init() }
func file_FulcrumProto_FulcrumProto_proto_init() {
	if File_FulcrumProto_FulcrumProto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FulcrumProto_FulcrumProto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewData); i {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogReq); i {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanetList); i {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanetData); i {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlanetClock); i {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewCity); i {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelCity); i {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpCity); i {
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
		file_FulcrumProto_FulcrumProto_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CityStatus); i {
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
			RawDescriptor: file_FulcrumProto_FulcrumProto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_FulcrumProto_FulcrumProto_proto_goTypes,
		DependencyIndexes: file_FulcrumProto_FulcrumProto_proto_depIdxs,
		MessageInfos:      file_FulcrumProto_FulcrumProto_proto_msgTypes,
	}.Build()
	File_FulcrumProto_FulcrumProto_proto = out.File
	file_FulcrumProto_FulcrumProto_proto_rawDesc = nil
	file_FulcrumProto_FulcrumProto_proto_goTypes = nil
	file_FulcrumProto_FulcrumProto_proto_depIdxs = nil
}
