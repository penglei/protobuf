// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto3_proto/proto3.proto

package proto3_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	test_proto "github.com/golang/protobuf/proto/test_proto"
	any "github.com/golang/protobuf/ptypes/any"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Message_Humour int32

const (
	Message_UNKNOWN     Message_Humour = 0
	Message_PUNS        Message_Humour = 1
	Message_SLAPSTICK   Message_Humour = 2
	Message_BILL_BAILEY Message_Humour = 3
)

var Message_Humour_name = map[int32]string{
	0: "UNKNOWN",
	1: "PUNS",
	2: "SLAPSTICK",
	3: "BILL_BAILEY",
}

var Message_Humour_value = map[string]int32{
	"UNKNOWN":     0,
	"PUNS":        1,
	"SLAPSTICK":   2,
	"BILL_BAILEY": 3,
}

func (x Message_Humour) String() string {
	return proto.EnumName(Message_Humour_name, int32(x))
}

func (Message_Humour) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{0, 0}
}

type Message struct {
	Name                 string                             `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hilarity             Message_Humour                     `protobuf:"varint,2,opt,name=hilarity,proto3,enum=proto3_proto.Message_Humour" json:"hilarity,omitempty"`
	HeightInCm           uint32                             `protobuf:"varint,3,opt,name=height_in_cm,json=heightInCm,proto3" json:"height_in_cm,omitempty"`
	Data                 []byte                             `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	ResultCount          int64                              `protobuf:"varint,7,opt,name=result_count,json=resultCount,proto3" json:"result_count,omitempty"`
	TrueScotsman         bool                               `protobuf:"varint,8,opt,name=true_scotsman,json=trueScotsman,proto3" json:"true_scotsman,omitempty"`
	Score                float32                            `protobuf:"fixed32,9,opt,name=score,proto3" json:"score,omitempty"`
	Key                  []uint64                           `protobuf:"varint,5,rep,packed,name=key,proto3" json:"key,omitempty"`
	ShortKey             []int32                            `protobuf:"varint,19,rep,packed,name=short_key,json=shortKey,proto3" json:"short_key,omitempty"`
	Nested               *Nested                            `protobuf:"bytes,6,opt,name=nested,proto3" json:"nested,omitempty"`
	RFunny               []Message_Humour                   `protobuf:"varint,16,rep,packed,name=r_funny,json=rFunny,proto3,enum=proto3_proto.Message_Humour" json:"r_funny,omitempty"`
	Terrain              map[string]*Nested                 `protobuf:"bytes,10,rep,name=terrain,proto3" json:"terrain,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Proto2Field          *test_proto.SubDefaults            `protobuf:"bytes,11,opt,name=proto2_field,json=proto2Field,proto3" json:"proto2_field,omitempty"`
	Proto2Value          map[string]*test_proto.SubDefaults `protobuf:"bytes,13,rep,name=proto2_value,json=proto2Value,proto3" json:"proto2_value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Anything             *any.Any                           `protobuf:"bytes,14,opt,name=anything,proto3" json:"anything,omitempty"`
	ManyThings           []*any.Any                         `protobuf:"bytes,15,rep,name=many_things,json=manyThings,proto3" json:"many_things,omitempty"`
	Submessage           *Message                           `protobuf:"bytes,17,opt,name=submessage,proto3" json:"submessage,omitempty"`
	Children             []*Message                         `protobuf:"bytes,18,rep,name=children,proto3" json:"children,omitempty"`
	StringMap            map[string]string                  `protobuf:"bytes,20,rep,name=string_map,json=stringMap,proto3" json:"string_map,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}
func (*Message) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{0}
}

func (m *Message) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Message.Unmarshal(m, b)
}
func (m *Message) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Message.Marshal(b, m, deterministic)
}
func (m *Message) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Message.Merge(m, src)
}
func (m *Message) XXX_Size() int {
	return xxx_messageInfo_Message.Size(m)
}
func (m *Message) XXX_DiscardUnknown() {
	xxx_messageInfo_Message.DiscardUnknown(m)
}

var xxx_messageInfo_Message proto.InternalMessageInfo

func (m *Message) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Message) GetHilarity() Message_Humour {
	if m != nil {
		return m.Hilarity
	}
	return Message_UNKNOWN
}

func (m *Message) GetHeightInCm() uint32 {
	if m != nil {
		return m.HeightInCm
	}
	return 0
}

func (m *Message) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Message) GetResultCount() int64 {
	if m != nil {
		return m.ResultCount
	}
	return 0
}

func (m *Message) GetTrueScotsman() bool {
	if m != nil {
		return m.TrueScotsman
	}
	return false
}

func (m *Message) GetScore() float32 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Message) GetKey() []uint64 {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Message) GetShortKey() []int32 {
	if m != nil {
		return m.ShortKey
	}
	return nil
}

func (m *Message) GetNested() *Nested {
	if m != nil {
		return m.Nested
	}
	return nil
}

func (m *Message) GetRFunny() []Message_Humour {
	if m != nil {
		return m.RFunny
	}
	return nil
}

func (m *Message) GetTerrain() map[string]*Nested {
	if m != nil {
		return m.Terrain
	}
	return nil
}

func (m *Message) GetProto2Field() *test_proto.SubDefaults {
	if m != nil {
		return m.Proto2Field
	}
	return nil
}

func (m *Message) GetProto2Value() map[string]*test_proto.SubDefaults {
	if m != nil {
		return m.Proto2Value
	}
	return nil
}

func (m *Message) GetAnything() *any.Any {
	if m != nil {
		return m.Anything
	}
	return nil
}

func (m *Message) GetManyThings() []*any.Any {
	if m != nil {
		return m.ManyThings
	}
	return nil
}

func (m *Message) GetSubmessage() *Message {
	if m != nil {
		return m.Submessage
	}
	return nil
}

func (m *Message) GetChildren() []*Message {
	if m != nil {
		return m.Children
	}
	return nil
}

func (m *Message) GetStringMap() map[string]string {
	if m != nil {
		return m.StringMap
	}
	return nil
}

type Nested struct {
	Bunny                string   `protobuf:"bytes,1,opt,name=bunny,proto3" json:"bunny,omitempty"`
	Cute                 bool     `protobuf:"varint,2,opt,name=cute,proto3" json:"cute,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nested) Reset()         { *m = Nested{} }
func (m *Nested) String() string { return proto.CompactTextString(m) }
func (*Nested) ProtoMessage()    {}
func (*Nested) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{1}
}

func (m *Nested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nested.Unmarshal(m, b)
}
func (m *Nested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nested.Marshal(b, m, deterministic)
}
func (m *Nested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nested.Merge(m, src)
}
func (m *Nested) XXX_Size() int {
	return xxx_messageInfo_Nested.Size(m)
}
func (m *Nested) XXX_DiscardUnknown() {
	xxx_messageInfo_Nested.DiscardUnknown(m)
}

var xxx_messageInfo_Nested proto.InternalMessageInfo

func (m *Nested) GetBunny() string {
	if m != nil {
		return m.Bunny
	}
	return ""
}

func (m *Nested) GetCute() bool {
	if m != nil {
		return m.Cute
	}
	return false
}

type MessageWithMap struct {
	ByteMapping          map[bool][]byte `protobuf:"bytes,1,rep,name=byte_mapping,json=byteMapping,proto3" json:"byte_mapping,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MessageWithMap) Reset()         { *m = MessageWithMap{} }
func (m *MessageWithMap) String() string { return proto.CompactTextString(m) }
func (*MessageWithMap) ProtoMessage()    {}
func (*MessageWithMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{2}
}

func (m *MessageWithMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageWithMap.Unmarshal(m, b)
}
func (m *MessageWithMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageWithMap.Marshal(b, m, deterministic)
}
func (m *MessageWithMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageWithMap.Merge(m, src)
}
func (m *MessageWithMap) XXX_Size() int {
	return xxx_messageInfo_MessageWithMap.Size(m)
}
func (m *MessageWithMap) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageWithMap.DiscardUnknown(m)
}

var xxx_messageInfo_MessageWithMap proto.InternalMessageInfo

func (m *MessageWithMap) GetByteMapping() map[bool][]byte {
	if m != nil {
		return m.ByteMapping
	}
	return nil
}

type IntMap struct {
	Rtt                  map[int32]int32 `protobuf:"bytes,1,rep,name=rtt,proto3" json:"rtt,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IntMap) Reset()         { *m = IntMap{} }
func (m *IntMap) String() string { return proto.CompactTextString(m) }
func (*IntMap) ProtoMessage()    {}
func (*IntMap) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{3}
}

func (m *IntMap) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMap.Unmarshal(m, b)
}
func (m *IntMap) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMap.Marshal(b, m, deterministic)
}
func (m *IntMap) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMap.Merge(m, src)
}
func (m *IntMap) XXX_Size() int {
	return xxx_messageInfo_IntMap.Size(m)
}
func (m *IntMap) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMap.DiscardUnknown(m)
}

var xxx_messageInfo_IntMap proto.InternalMessageInfo

func (m *IntMap) GetRtt() map[int32]int32 {
	if m != nil {
		return m.Rtt
	}
	return nil
}

type IntMaps struct {
	Maps                 []*IntMap `protobuf:"bytes,1,rep,name=maps,proto3" json:"maps,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *IntMaps) Reset()         { *m = IntMaps{} }
func (m *IntMaps) String() string { return proto.CompactTextString(m) }
func (*IntMaps) ProtoMessage()    {}
func (*IntMaps) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{4}
}

func (m *IntMaps) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IntMaps.Unmarshal(m, b)
}
func (m *IntMaps) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IntMaps.Marshal(b, m, deterministic)
}
func (m *IntMaps) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IntMaps.Merge(m, src)
}
func (m *IntMaps) XXX_Size() int {
	return xxx_messageInfo_IntMaps.Size(m)
}
func (m *IntMaps) XXX_DiscardUnknown() {
	xxx_messageInfo_IntMaps.DiscardUnknown(m)
}

var xxx_messageInfo_IntMaps proto.InternalMessageInfo

func (m *IntMaps) GetMaps() []*IntMap {
	if m != nil {
		return m.Maps
	}
	return nil
}

type TestUTF8 struct {
	Scalar string   `protobuf:"bytes,1,opt,name=scalar,proto3" json:"scalar,omitempty"`
	Vector []string `protobuf:"bytes,2,rep,name=vector,proto3" json:"vector,omitempty"`
	// Types that are valid to be assigned to Oneof:
	//	*TestUTF8_Field
	Oneof                isTestUTF8_Oneof `protobuf_oneof:"oneof"`
	MapKey               map[string]int64 `protobuf:"bytes,4,rep,name=map_key,json=mapKey,proto3" json:"map_key,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MapValue             map[int64]string `protobuf:"bytes,5,rep,name=map_value,json=mapValue,proto3" json:"map_value,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *TestUTF8) Reset()         { *m = TestUTF8{} }
func (m *TestUTF8) String() string { return proto.CompactTextString(m) }
func (*TestUTF8) ProtoMessage()    {}
func (*TestUTF8) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{5}
}

func (m *TestUTF8) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TestUTF8.Unmarshal(m, b)
}
func (m *TestUTF8) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TestUTF8.Marshal(b, m, deterministic)
}
func (m *TestUTF8) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TestUTF8.Merge(m, src)
}
func (m *TestUTF8) XXX_Size() int {
	return xxx_messageInfo_TestUTF8.Size(m)
}
func (m *TestUTF8) XXX_DiscardUnknown() {
	xxx_messageInfo_TestUTF8.DiscardUnknown(m)
}

var xxx_messageInfo_TestUTF8 proto.InternalMessageInfo

func (m *TestUTF8) GetScalar() string {
	if m != nil {
		return m.Scalar
	}
	return ""
}

func (m *TestUTF8) GetVector() []string {
	if m != nil {
		return m.Vector
	}
	return nil
}

type isTestUTF8_Oneof interface {
	isTestUTF8_Oneof()
}

type TestUTF8_Field struct {
	Field string `protobuf:"bytes,3,opt,name=field,proto3,oneof"`
}

func (*TestUTF8_Field) isTestUTF8_Oneof() {}

func (m *TestUTF8) GetOneof() isTestUTF8_Oneof {
	if m != nil {
		return m.Oneof
	}
	return nil
}

func (m *TestUTF8) GetField() string {
	if x, ok := m.GetOneof().(*TestUTF8_Field); ok {
		return x.Field
	}
	return ""
}

func (m *TestUTF8) GetMapKey() map[string]int64 {
	if m != nil {
		return m.MapKey
	}
	return nil
}

func (m *TestUTF8) GetMapValue() map[int64]string {
	if m != nil {
		return m.MapValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TestUTF8) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TestUTF8_Field)(nil),
	}
}

type Wrapped struct {
	Foo                  *wrappers.BoolValue `protobuf:"bytes,1,opt,name=foo,proto3" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *Wrapped) Reset()         { *m = Wrapped{} }
func (m *Wrapped) String() string { return proto.CompactTextString(m) }
func (*Wrapped) ProtoMessage()    {}
func (*Wrapped) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{6}
}

func (m *Wrapped) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Wrapped.Unmarshal(m, b)
}
func (m *Wrapped) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Wrapped.Marshal(b, m, deterministic)
}
func (m *Wrapped) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wrapped.Merge(m, src)
}
func (m *Wrapped) XXX_Size() int {
	return xxx_messageInfo_Wrapped.Size(m)
}
func (m *Wrapped) XXX_DiscardUnknown() {
	xxx_messageInfo_Wrapped.DiscardUnknown(m)
}

var xxx_messageInfo_Wrapped proto.InternalMessageInfo

func (m *Wrapped) GetFoo() *wrappers.BoolValue {
	if m != nil {
		return m.Foo
	}
	return nil
}

type Scalar struct {
	Foo                  bool     `protobuf:"varint,1,opt,name=foo,proto3" json:"foo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Scalar) Reset()         { *m = Scalar{} }
func (m *Scalar) String() string { return proto.CompactTextString(m) }
func (*Scalar) ProtoMessage()    {}
func (*Scalar) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c50d9b824d4ac38, []int{7}
}

func (m *Scalar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Scalar.Unmarshal(m, b)
}
func (m *Scalar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Scalar.Marshal(b, m, deterministic)
}
func (m *Scalar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Scalar.Merge(m, src)
}
func (m *Scalar) XXX_Size() int {
	return xxx_messageInfo_Scalar.Size(m)
}
func (m *Scalar) XXX_DiscardUnknown() {
	xxx_messageInfo_Scalar.DiscardUnknown(m)
}

var xxx_messageInfo_Scalar proto.InternalMessageInfo

func (m *Scalar) GetFoo() bool {
	if m != nil {
		return m.Foo
	}
	return false
}

func init() {
	proto.RegisterEnum("proto3_proto.Message_Humour", Message_Humour_name, Message_Humour_value)
	proto.RegisterType((*Message)(nil), "proto3_proto.Message")
	proto.RegisterMapType((map[string]*test_proto.SubDefaults)(nil), "proto3_proto.Message.Proto2ValueEntry")
	proto.RegisterMapType((map[string]string)(nil), "proto3_proto.Message.StringMapEntry")
	proto.RegisterMapType((map[string]*Nested)(nil), "proto3_proto.Message.TerrainEntry")
	proto.RegisterType((*Nested)(nil), "proto3_proto.Nested")
	proto.RegisterType((*MessageWithMap)(nil), "proto3_proto.MessageWithMap")
	proto.RegisterMapType((map[bool][]byte)(nil), "proto3_proto.MessageWithMap.ByteMappingEntry")
	proto.RegisterType((*IntMap)(nil), "proto3_proto.IntMap")
	proto.RegisterMapType((map[int32]int32)(nil), "proto3_proto.IntMap.RttEntry")
	proto.RegisterType((*IntMaps)(nil), "proto3_proto.IntMaps")
	proto.RegisterType((*TestUTF8)(nil), "proto3_proto.TestUTF8")
	proto.RegisterMapType((map[string]int64)(nil), "proto3_proto.TestUTF8.MapKeyEntry")
	proto.RegisterMapType((map[int64]string)(nil), "proto3_proto.TestUTF8.MapValueEntry")
	proto.RegisterType((*Wrapped)(nil), "proto3_proto.Wrapped")
	proto.RegisterType((*Scalar)(nil), "proto3_proto.Scalar")
}

func init() {
	proto.RegisterFile("proto3_proto/proto3.proto", fileDescriptor_1c50d9b824d4ac38)
}

var fileDescriptor_1c50d9b824d4ac38 = []byte{
	// 946 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x55, 0xff, 0x6e, 0xdb, 0x36,
	0x10, 0xae, 0x2c, 0x5b, 0x96, 0x4f, 0x4e, 0xea, 0x71, 0x69, 0xa7, 0x7a, 0x3f, 0xa0, 0x79, 0xc3,
	0x20, 0x0c, 0xad, 0xba, 0xb9, 0xc8, 0x96, 0xb5, 0xc5, 0x86, 0x38, 0x6b, 0x50, 0x23, 0xb1, 0x67,
	0xc8, 0xce, 0x82, 0xfd, 0x25, 0xd0, 0x36, 0x6d, 0x0b, 0xb3, 0x7e, 0x8c, 0xa4, 0x3a, 0xe8, 0x05,
	0xf6, 0x20, 0x7b, 0xa5, 0xbd, 0xd0, 0x40, 0x52, 0x4e, 0xe4, 0x4c, 0x5d, 0xfe, 0x12, 0xef, 0xf8,
	0x7d, 0xf7, 0x1d, 0xef, 0x8e, 0x14, 0x3c, 0x49, 0x69, 0xc2, 0x93, 0x17, 0x81, 0xfc, 0x3c, 0x57,
	0x86, 0x27, 0x3f, 0xa8, 0x5d, 0xde, 0xea, 0x3e, 0x59, 0x27, 0xc9, 0x7a, 0x4b, 0x14, 0x64, 0x9e,
	0xad, 0x9e, 0xe3, 0x38, 0x57, 0xc0, 0xee, 0x67, 0x77, 0xb7, 0xfe, 0xa4, 0x38, 0x4d, 0x09, 0x65,
	0xc5, 0xfe, 0x23, 0x4e, 0x18, 0x2f, 0x14, 0xc4, 0x52, 0xb9, 0x7b, 0x7f, 0xb5, 0xa0, 0x39, 0x22,
	0x8c, 0xe1, 0x35, 0x41, 0x08, 0xea, 0x31, 0x8e, 0x88, 0xad, 0x39, 0x9a, 0xdb, 0xf2, 0xe5, 0x1a,
	0x9d, 0x80, 0xb9, 0x09, 0xb7, 0x98, 0x86, 0x3c, 0xb7, 0x6b, 0x8e, 0xe6, 0x1e, 0xf6, 0x3f, 0xf1,
	0xca, 0x29, 0x79, 0x05, 0xd9, 0x7b, 0x9b, 0x45, 0x49, 0x46, 0xfd, 0x1b, 0x34, 0x72, 0xa0, 0xbd,
	0x21, 0xe1, 0x7a, 0xc3, 0x83, 0x30, 0x0e, 0x16, 0x91, 0xad, 0x3b, 0x9a, 0x7b, 0xe0, 0x83, 0xf2,
	0x0d, 0xe3, 0xb3, 0x48, 0xe8, 0x2d, 0x31, 0xc7, 0x76, 0xdd, 0xd1, 0xdc, 0xb6, 0x2f, 0xd7, 0xe8,
	0x73, 0x68, 0x53, 0xc2, 0xb2, 0x2d, 0x0f, 0x16, 0x49, 0x16, 0x73, 0xbb, 0xe9, 0x68, 0xae, 0xee,
	0x5b, 0xca, 0x77, 0x26, 0x5c, 0xe8, 0x0b, 0x38, 0xe0, 0x34, 0x23, 0x01, 0x5b, 0x24, 0x9c, 0x45,
	0x38, 0xb6, 0x4d, 0x47, 0x73, 0x4d, 0xbf, 0x2d, 0x9c, 0xd3, 0xc2, 0x87, 0x8e, 0xa0, 0xc1, 0x16,
	0x09, 0x25, 0x76, 0xcb, 0xd1, 0xdc, 0x9a, 0xaf, 0x0c, 0xd4, 0x01, 0xfd, 0x77, 0x92, 0xdb, 0x0d,
	0x47, 0x77, 0xeb, 0xbe, 0x58, 0xa2, 0x8f, 0xa1, 0xc5, 0x36, 0x09, 0xe5, 0x81, 0xf0, 0x7f, 0xe8,
	0xe8, 0x6e, 0xc3, 0x37, 0xa5, 0xe3, 0x82, 0xe4, 0xe8, 0x29, 0x18, 0x31, 0x61, 0x9c, 0x2c, 0x6d,
	0xc3, 0xd1, 0x5c, 0xab, 0x7f, 0xb4, 0x7f, 0xf4, 0xb1, 0xdc, 0xf3, 0x0b, 0x0c, 0x3a, 0x86, 0x26,
	0x0d, 0x56, 0x59, 0x1c, 0xe7, 0x76, 0xc7, 0xd1, 0xef, 0xad, 0x94, 0x41, 0xcf, 0x05, 0x16, 0xbd,
	0x86, 0x26, 0x27, 0x94, 0xe2, 0x30, 0xb6, 0xc1, 0xd1, 0x5d, 0xab, 0xdf, 0xab, 0xa6, 0xcd, 0x14,
	0xe8, 0x4d, 0xcc, 0x69, 0xee, 0xef, 0x28, 0xe8, 0x25, 0xa8, 0x09, 0xe9, 0x07, 0xab, 0x90, 0x6c,
	0x97, 0xb6, 0x25, 0x13, 0xfd, 0xc8, 0xbb, 0xed, 0xb6, 0x37, 0xcd, 0xe6, 0x3f, 0x93, 0x15, 0xce,
	0xb6, 0x9c, 0xf9, 0x96, 0x02, 0x9f, 0x0b, 0x2c, 0x1a, 0xde, 0x70, 0xdf, 0xe1, 0x6d, 0x46, 0xec,
	0x03, 0x29, 0xff, 0x55, 0xb5, 0xfc, 0x44, 0x22, 0x7f, 0x15, 0x40, 0x95, 0x42, 0x11, 0x4a, 0x7a,
	0xd0, 0x37, 0x60, 0xe2, 0x38, 0xe7, 0x9b, 0x30, 0x5e, 0xdb, 0x87, 0x45, 0xad, 0xd4, 0x40, 0x7a,
	0xbb, 0x81, 0xf4, 0x4e, 0xe3, 0xdc, 0xbf, 0x41, 0xa1, 0x63, 0xb0, 0x22, 0x1c, 0xe7, 0x81, 0xb4,
	0x98, 0xfd, 0x50, 0x6a, 0x57, 0x93, 0x40, 0x00, 0x67, 0x12, 0x87, 0x8e, 0x01, 0x58, 0x36, 0x8f,
	0x54, 0x52, 0xf6, 0x07, 0x52, 0xea, 0x51, 0x65, 0xc6, 0x7e, 0x09, 0x88, 0xbe, 0x05, 0x73, 0xb1,
	0x09, 0xb7, 0x4b, 0x4a, 0x62, 0x1b, 0x49, 0xa9, 0xf7, 0x90, 0x6e, 0x60, 0xe8, 0x0c, 0x80, 0x71,
	0x1a, 0xc6, 0xeb, 0x20, 0xc2, 0xa9, 0x7d, 0x24, 0x49, 0x5f, 0x56, 0xd7, 0x66, 0x2a, 0x71, 0x23,
	0x9c, 0xaa, 0xca, 0xb4, 0xd8, 0xce, 0xee, 0x4e, 0xa0, 0x5d, 0xee, 0xdb, 0x6e, 0x00, 0xd5, 0x0d,
	0x93, 0x03, 0xf8, 0x35, 0x34, 0x54, 0xf5, 0x6b, 0xff, 0x33, 0x62, 0x0a, 0xf2, 0xb2, 0x76, 0xa2,
	0x75, 0xaf, 0xa1, 0x73, 0xb7, 0x15, 0x15, 0x51, 0x9f, 0xed, 0x47, 0x7d, 0xef, 0x3c, 0x94, 0x02,
	0xbf, 0x86, 0xc3, 0xfd, 0x73, 0x54, 0x84, 0x3d, 0x2a, 0x87, 0x6d, 0x95, 0xd8, 0xbd, 0x9f, 0xc0,
	0x50, 0x73, 0x8d, 0x2c, 0x68, 0x5e, 0x8d, 0x2f, 0xc6, 0xbf, 0x5c, 0x8f, 0x3b, 0x0f, 0x90, 0x09,
	0xf5, 0xc9, 0xd5, 0x78, 0xda, 0xd1, 0xd0, 0x01, 0xb4, 0xa6, 0x97, 0xa7, 0x93, 0xe9, 0x6c, 0x78,
	0x76, 0xd1, 0xa9, 0xa1, 0x87, 0x60, 0x0d, 0x86, 0x97, 0x97, 0xc1, 0xe0, 0x74, 0x78, 0xf9, 0xe6,
	0xb7, 0x8e, 0xde, 0xeb, 0x83, 0xa1, 0x0e, 0x2b, 0x44, 0xe6, 0xf2, 0x16, 0x29, 0x61, 0x65, 0x88,
	0xc7, 0x62, 0x91, 0x71, 0xa5, 0x6c, 0xfa, 0x72, 0xdd, 0xfb, 0x5b, 0x83, 0xc3, 0xa2, 0x07, 0xd7,
	0x21, 0xdf, 0x8c, 0x70, 0x8a, 0x26, 0xd0, 0x9e, 0xe7, 0x9c, 0x88, 0x9e, 0xa5, 0x62, 0x18, 0x35,
	0xd9, 0xb7, 0x67, 0x95, 0x7d, 0x2b, 0x38, 0xde, 0x20, 0xe7, 0x64, 0xa4, 0xf0, 0xc5, 0x68, 0xcf,
	0x6f, 0x3d, 0xdd, 0x1f, 0xa1, 0x73, 0x17, 0x50, 0xae, 0x8c, 0x59, 0x51, 0x99, 0x76, 0xb9, 0x32,
	0x7f, 0x80, 0x31, 0x8c, 0xb9, 0xc8, 0xed, 0x39, 0xe8, 0x94, 0xf3, 0x22, 0xa5, 0x4f, 0xf7, 0x53,
	0x52, 0x10, 0xcf, 0xe7, 0x5c, 0xa5, 0x20, 0x90, 0xdd, 0xef, 0xc0, 0xdc, 0x39, 0xca, 0x92, 0x8d,
	0x0a, 0xc9, 0x46, 0x59, 0xf2, 0x05, 0x34, 0x55, 0x3c, 0x86, 0x5c, 0xa8, 0x47, 0x38, 0x65, 0x85,
	0xe8, 0x51, 0x95, 0xa8, 0x2f, 0x11, 0xbd, 0x7f, 0x6a, 0x60, 0xce, 0x08, 0xe3, 0x57, 0xb3, 0xf3,
	0x13, 0xf4, 0x18, 0x0c, 0xb6, 0xc0, 0x5b, 0x4c, 0x8b, 0x26, 0x14, 0x96, 0xf0, 0xbf, 0x23, 0x0b,
	0x9e, 0x50, 0xbb, 0xe6, 0xe8, 0xc2, 0xaf, 0x2c, 0xf4, 0x18, 0x1a, 0xea, 0xfd, 0x11, 0xaf, 0x7c,
	0xeb, 0xed, 0x03, 0x5f, 0x99, 0xe8, 0x15, 0x34, 0x23, 0x9c, 0xca, 0xc7, 0xb5, 0x5e, 0xf5, 0xb8,
	0xed, 0x04, 0xbd, 0x11, 0x4e, 0x2f, 0x48, 0xae, 0xce, 0x6e, 0x44, 0xd2, 0x40, 0xa7, 0xd0, 0x12,
	0x64, 0x75, 0xc8, 0x46, 0xd5, 0x05, 0x2c, 0xd3, 0x4b, 0x4f, 0x93, 0x19, 0x15, 0x66, 0xf7, 0x07,
	0xb0, 0x4a, 0x91, 0xef, 0x9b, 0x68, 0xbd, 0x7c, 0x1f, 0x5e, 0xc1, 0xc1, 0x5e, 0xd4, 0x32, 0x59,
	0xbf, 0xe7, 0x3a, 0x0c, 0x9a, 0xd0, 0x48, 0x62, 0x92, 0xac, 0x7a, 0xdf, 0x43, 0xf3, 0x5a, 0xfe,
	0x88, 0x97, 0xe8, 0x29, 0xe8, 0xab, 0x24, 0x91, 0x7c, 0xab, 0xdf, 0xfd, 0xcf, 0x4b, 0x37, 0x48,
	0x92, 0xad, 0x54, 0xf3, 0x05, 0xac, 0xd7, 0x05, 0x63, 0xaa, 0x6a, 0xde, 0xb9, 0xe5, 0x99, 0x72,
	0x6f, 0x6e, 0xa8, 0x22, 0xfc, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x22, 0xcb, 0x22, 0x23, 0x38, 0x08,
	0x00, 0x00,
}
