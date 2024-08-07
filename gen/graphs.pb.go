// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: graphs.proto

package grpc_todo

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type TimeRange int32

const (
	TimeRange_TIME_RANGE_SECOND       TimeRange = 0
	TimeRange_TIME_RANGE_MINUTE       TimeRange = 1
	TimeRange_TIME_RANGE_FIVE_MINUTES TimeRange = 2
	TimeRange_TIME_RANGE_DAY          TimeRange = 3
)

// Enum value maps for TimeRange.
var (
	TimeRange_name = map[int32]string{
		0: "TIME_RANGE_SECOND",
		1: "TIME_RANGE_MINUTE",
		2: "TIME_RANGE_FIVE_MINUTES",
		3: "TIME_RANGE_DAY",
	}
	TimeRange_value = map[string]int32{
		"TIME_RANGE_SECOND":       0,
		"TIME_RANGE_MINUTE":       1,
		"TIME_RANGE_FIVE_MINUTES": 2,
		"TIME_RANGE_DAY":          3,
	}
)

func (x TimeRange) Enum() *TimeRange {
	p := new(TimeRange)
	*p = x
	return p
}

func (x TimeRange) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeRange) Descriptor() protoreflect.EnumDescriptor {
	return file_graphs_proto_enumTypes[0].Descriptor()
}

func (TimeRange) Type() protoreflect.EnumType {
	return &file_graphs_proto_enumTypes[0]
}

func (x TimeRange) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeRange.Descriptor instead.
func (TimeRange) EnumDescriptor() ([]byte, []int) {
	return file_graphs_proto_rawDescGZIP(), []int{0}
}

type SymbolIdGraphs int32

const (
	SymbolIdGraphs_APPL SymbolIdGraphs = 0
	SymbolIdGraphs_NVDA SymbolIdGraphs = 1
	SymbolIdGraphs_TSLA SymbolIdGraphs = 2
)

// Enum value maps for SymbolIdGraphs.
var (
	SymbolIdGraphs_name = map[int32]string{
		0: "APPL",
		1: "NVDA",
		2: "TSLA",
	}
	SymbolIdGraphs_value = map[string]int32{
		"APPL": 0,
		"NVDA": 1,
		"TSLA": 2,
	}
)

func (x SymbolIdGraphs) Enum() *SymbolIdGraphs {
	p := new(SymbolIdGraphs)
	*p = x
	return p
}

func (x SymbolIdGraphs) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SymbolIdGraphs) Descriptor() protoreflect.EnumDescriptor {
	return file_graphs_proto_enumTypes[1].Descriptor()
}

func (SymbolIdGraphs) Type() protoreflect.EnumType {
	return &file_graphs_proto_enumTypes[1]
}

func (x SymbolIdGraphs) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SymbolIdGraphs.Descriptor instead.
func (SymbolIdGraphs) EnumDescriptor() ([]byte, []int) {
	return file_graphs_proto_rawDescGZIP(), []int{1}
}

type Candles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Candles []*Candle `protobuf:"bytes,1,rep,name=candles,proto3" json:"candles,omitempty"`
}

func (x *Candles) Reset() {
	*x = Candles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphs_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candles) ProtoMessage() {}

func (x *Candles) ProtoReflect() protoreflect.Message {
	mi := &file_graphs_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candles.ProtoReflect.Descriptor instead.
func (*Candles) Descriptor() ([]byte, []int) {
	return file_graphs_proto_rawDescGZIP(), []int{0}
}

func (x *Candles) GetCandles() []*Candle {
	if x != nil {
		return x.Candles
	}
	return nil
}

type Candle struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Symbol    SymbolIdGraphs         `protobuf:"varint,1,opt,name=symbol,proto3,enum=graphs.SymbolIdGraphs" json:"symbol,omitempty"`
	Open      uint32                 `protobuf:"varint,2,opt,name=open,proto3" json:"open,omitempty"`
	High      uint32                 `protobuf:"varint,3,opt,name=high,proto3" json:"high,omitempty"`
	Low       uint32                 `protobuf:"varint,4,opt,name=low,proto3" json:"low,omitempty"`
	Close     uint32                 `protobuf:"varint,5,opt,name=close,proto3" json:"close,omitempty"`
	TimeRange TimeRange              `protobuf:"varint,6,opt,name=time_range,json=timeRange,proto3,enum=graphs.TimeRange" json:"time_range,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Candle) Reset() {
	*x = Candle{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphs_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Candle) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Candle) ProtoMessage() {}

func (x *Candle) ProtoReflect() protoreflect.Message {
	mi := &file_graphs_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Candle.ProtoReflect.Descriptor instead.
func (*Candle) Descriptor() ([]byte, []int) {
	return file_graphs_proto_rawDescGZIP(), []int{1}
}

func (x *Candle) GetSymbol() SymbolIdGraphs {
	if x != nil {
		return x.Symbol
	}
	return SymbolIdGraphs_APPL
}

func (x *Candle) GetOpen() uint32 {
	if x != nil {
		return x.Open
	}
	return 0
}

func (x *Candle) GetHigh() uint32 {
	if x != nil {
		return x.High
	}
	return 0
}

func (x *Candle) GetLow() uint32 {
	if x != nil {
		return x.Low
	}
	return 0
}

func (x *Candle) GetClose() uint32 {
	if x != nil {
		return x.Close
	}
	return 0
}

func (x *Candle) GetTimeRange() TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return TimeRange_TIME_RANGE_SECOND
}

func (x *Candle) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type Symbol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        SymbolIdGraphs `protobuf:"varint,1,opt,name=id,proto3,enum=graphs.SymbolIdGraphs" json:"id,omitempty"`
	TimeRange TimeRange      `protobuf:"varint,2,opt,name=time_range,json=timeRange,proto3,enum=graphs.TimeRange" json:"time_range,omitempty"`
}

func (x *Symbol) Reset() {
	*x = Symbol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graphs_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Symbol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Symbol) ProtoMessage() {}

func (x *Symbol) ProtoReflect() protoreflect.Message {
	mi := &file_graphs_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Symbol.ProtoReflect.Descriptor instead.
func (*Symbol) Descriptor() ([]byte, []int) {
	return file_graphs_proto_rawDescGZIP(), []int{2}
}

func (x *Symbol) GetId() SymbolIdGraphs {
	if x != nil {
		return x.Id
	}
	return SymbolIdGraphs_APPL
}

func (x *Symbol) GetTimeRange() TimeRange {
	if x != nil {
		return x.TimeRange
	}
	return TimeRange_TIME_RANGE_SECOND
}

var File_graphs_proto protoreflect.FileDescriptor

var file_graphs_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x07, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x12,
	0x28, 0x0a, 0x07, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x52, 0x07, 0x63, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x73, 0x22, 0xff, 0x01, 0x0a, 0x06, 0x43, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x2e, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x49, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x73, 0x42, 0x08, 0xba, 0x48,
	0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x06, 0x73, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x12,
	0x0a, 0x04, 0x6f, 0x70, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x6f, 0x70,
	0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x69, 0x67, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x68, 0x69, 0x67, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x6f, 0x77, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x03, 0x6c, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x12, 0x30,
	0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x11, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65,
	0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x6c, 0x0a, 0x06, 0x53,
	0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x30, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x16, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f,
	0x6c, 0x49, 0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x73, 0x42, 0x08, 0xba, 0x48, 0x05, 0x82, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x12, 0x30, 0x0a, 0x0a, 0x74, 0x69, 0x6d, 0x65, 0x5f,
	0x72, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x67, 0x72,
	0x61, 0x70, 0x68, 0x73, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x2a, 0x6a, 0x0a, 0x09, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52,
	0x41, 0x4e, 0x47, 0x45, 0x5f, 0x53, 0x45, 0x43, 0x4f, 0x4e, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55,
	0x54, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a, 0x17, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e,
	0x47, 0x45, 0x5f, 0x46, 0x49, 0x56, 0x45, 0x5f, 0x4d, 0x49, 0x4e, 0x55, 0x54, 0x45, 0x53, 0x10,
	0x02, 0x12, 0x12, 0x0a, 0x0e, 0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f,
	0x44, 0x41, 0x59, 0x10, 0x03, 0x2a, 0x2e, 0x0a, 0x0e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x49,
	0x64, 0x47, 0x72, 0x61, 0x70, 0x68, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x41, 0x50, 0x50, 0x4c, 0x10,
	0x00, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x56, 0x44, 0x41, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x54,
	0x53, 0x4c, 0x41, 0x10, 0x02, 0x32, 0x79, 0x0a, 0x0c, 0x47, 0x72, 0x61, 0x70, 0x68, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x12, 0x0e, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x2e, 0x53, 0x79,
	0x6d, 0x62, 0x6f, 0x6c, 0x1a, 0x0f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x2e, 0x43, 0x61,
	0x6e, 0x64, 0x6c, 0x65, 0x73, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x79,
	0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x0e, 0x2e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x73, 0x2e, 0x53, 0x79, 0x6d, 0x62, 0x6f, 0x6c, 0x1a, 0x0e, 0x2e, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x73, 0x2e, 0x43, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x42, 0x68, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x73, 0x42, 0x0b,
	0x47, 0x72, 0x61, 0x70, 0x68, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x15, 0x65,
	0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d,
	0x74, 0x6f, 0x64, 0x6f, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x06, 0x47, 0x72, 0x61,
	0x70, 0x68, 0x73, 0xca, 0x02, 0x06, 0x47, 0x72, 0x61, 0x70, 0x68, 0x73, 0xe2, 0x02, 0x12, 0x47,
	0x72, 0x61, 0x70, 0x68, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x06, 0x47, 0x72, 0x61, 0x70, 0x68, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_graphs_proto_rawDescOnce sync.Once
	file_graphs_proto_rawDescData = file_graphs_proto_rawDesc
)

func file_graphs_proto_rawDescGZIP() []byte {
	file_graphs_proto_rawDescOnce.Do(func() {
		file_graphs_proto_rawDescData = protoimpl.X.CompressGZIP(file_graphs_proto_rawDescData)
	})
	return file_graphs_proto_rawDescData
}

var file_graphs_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_graphs_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_graphs_proto_goTypes = []any{
	(TimeRange)(0),                // 0: graphs.TimeRange
	(SymbolIdGraphs)(0),           // 1: graphs.SymbolIdGraphs
	(*Candles)(nil),               // 2: graphs.Candles
	(*Candle)(nil),                // 3: graphs.Candle
	(*Symbol)(nil),                // 4: graphs.Symbol
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_graphs_proto_depIdxs = []int32{
	3, // 0: graphs.Candles.candles:type_name -> graphs.Candle
	1, // 1: graphs.Candle.symbol:type_name -> graphs.SymbolIdGraphs
	0, // 2: graphs.Candle.time_range:type_name -> graphs.TimeRange
	5, // 3: graphs.Candle.created_at:type_name -> google.protobuf.Timestamp
	1, // 4: graphs.Symbol.id:type_name -> graphs.SymbolIdGraphs
	0, // 5: graphs.Symbol.time_range:type_name -> graphs.TimeRange
	4, // 6: graphs.GraphService.GetBySymbol:input_type -> graphs.Symbol
	4, // 7: graphs.GraphService.GetBySymbolStream:input_type -> graphs.Symbol
	2, // 8: graphs.GraphService.GetBySymbol:output_type -> graphs.Candles
	3, // 9: graphs.GraphService.GetBySymbolStream:output_type -> graphs.Candle
	8, // [8:10] is the sub-list for method output_type
	6, // [6:8] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_graphs_proto_init() }
func file_graphs_proto_init() {
	if File_graphs_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graphs_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Candles); i {
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
		file_graphs_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Candle); i {
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
		file_graphs_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Symbol); i {
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
			RawDescriptor: file_graphs_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_graphs_proto_goTypes,
		DependencyIndexes: file_graphs_proto_depIdxs,
		EnumInfos:         file_graphs_proto_enumTypes,
		MessageInfos:      file_graphs_proto_msgTypes,
	}.Build()
	File_graphs_proto = out.File
	file_graphs_proto_rawDesc = nil
	file_graphs_proto_goTypes = nil
	file_graphs_proto_depIdxs = nil
}
