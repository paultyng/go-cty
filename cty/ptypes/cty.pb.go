// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cty.proto

package ptypes

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type DynamicType int32

const (
	DynamicType_DYNAMIC_TYPE DynamicType = 0
)

var DynamicType_name = map[int32]string{
	0: "DYNAMIC_TYPE",
}

var DynamicType_value = map[string]int32{
	"DYNAMIC_TYPE": 0,
}

func (x DynamicType) String() string {
	return proto.EnumName(DynamicType_name, int32(x))
}

func (DynamicType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{0}
}

type PrimitiveType int32

const (
	PrimitiveType_STRING PrimitiveType = 0
	PrimitiveType_NUMBER PrimitiveType = 1
	PrimitiveType_BOOL   PrimitiveType = 2
)

var PrimitiveType_name = map[int32]string{
	0: "STRING",
	1: "NUMBER",
	2: "BOOL",
}

var PrimitiveType_value = map[string]int32{
	"STRING": 0,
	"NUMBER": 1,
	"BOOL":   2,
}

func (x PrimitiveType) String() string {
	return proto.EnumName(PrimitiveType_name, int32(x))
}

func (PrimitiveType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{1}
}

type CollectionKind int32

const (
	CollectionKind_LIST CollectionKind = 0
	CollectionKind_MAP  CollectionKind = 1
	CollectionKind_SET  CollectionKind = 2
)

var CollectionKind_name = map[int32]string{
	0: "LIST",
	1: "MAP",
	2: "SET",
}

var CollectionKind_value = map[string]int32{
	"LIST": 0,
	"MAP":  1,
	"SET":  2,
}

func (x CollectionKind) String() string {
	return proto.EnumName(CollectionKind_name, int32(x))
}

func (CollectionKind) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{2}
}

type UnknownValue int32

const (
	UnknownValue_UNKNOWN_VALUE UnknownValue = 0
)

var UnknownValue_name = map[int32]string{
	0: "UNKNOWN_VALUE",
}

var UnknownValue_value = map[string]int32{
	"UNKNOWN_VALUE": 0,
}

func (x UnknownValue) String() string {
	return proto.EnumName(UnknownValue_name, int32(x))
}

func (UnknownValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{3}
}

type NullValue int32

const (
	NullValue_NULL_VALUE NullValue = 0
)

var NullValue_name = map[int32]string{
	0: "NULL_VALUE",
}

var NullValue_value = map[string]int32{
	"NULL_VALUE": 0,
}

func (x NullValue) String() string {
	return proto.EnumName(NullValue_name, int32(x))
}

func (NullValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{4}
}

type InfinityValue int32

const (
	InfinityValue_POSITIVE_INFINITY InfinityValue = 0
	InfinityValue_NEGATIVE_INFINITY InfinityValue = 1
)

var InfinityValue_name = map[int32]string{
	0: "POSITIVE_INFINITY",
	1: "NEGATIVE_INFINITY",
}

var InfinityValue_value = map[string]int32{
	"POSITIVE_INFINITY": 0,
	"NEGATIVE_INFINITY": 1,
}

func (x InfinityValue) String() string {
	return proto.EnumName(InfinityValue_name, int32(x))
}

func (InfinityValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{5}
}

type Type struct {
	// Types that are valid to be assigned to Kind:
	//	*Type_DynamicType
	//	*Type_PrimitiveType
	//	*Type_CollectionType
	//	*Type_ObjectType
	//	*Type_TupleType
	Kind                 isType_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Type) Reset()         { *m = Type{} }
func (m *Type) String() string { return proto.CompactTextString(m) }
func (*Type) ProtoMessage()    {}
func (*Type) Descriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{0}
}

func (m *Type) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Type.Unmarshal(m, b)
}
func (m *Type) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Type.Marshal(b, m, deterministic)
}
func (m *Type) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Type.Merge(m, src)
}
func (m *Type) XXX_Size() int {
	return xxx_messageInfo_Type.Size(m)
}
func (m *Type) XXX_DiscardUnknown() {
	xxx_messageInfo_Type.DiscardUnknown(m)
}

var xxx_messageInfo_Type proto.InternalMessageInfo

type isType_Kind interface {
	isType_Kind()
}

type Type_DynamicType struct {
	DynamicType DynamicType `protobuf:"varint,1,opt,name=dynamic_type,json=dynamicType,proto3,enum=zclconf.cty.DynamicType,oneof"`
}

type Type_PrimitiveType struct {
	PrimitiveType PrimitiveType `protobuf:"varint,2,opt,name=primitive_type,json=primitiveType,proto3,enum=zclconf.cty.PrimitiveType,oneof"`
}

type Type_CollectionType struct {
	CollectionType *CollectionType `protobuf:"bytes,3,opt,name=collection_type,json=collectionType,proto3,oneof"`
}

type Type_ObjectType struct {
	ObjectType *ObjectType `protobuf:"bytes,4,opt,name=object_type,json=objectType,proto3,oneof"`
}

type Type_TupleType struct {
	TupleType *TupleType `protobuf:"bytes,5,opt,name=tuple_type,json=tupleType,proto3,oneof"`
}

func (*Type_DynamicType) isType_Kind() {}

func (*Type_PrimitiveType) isType_Kind() {}

func (*Type_CollectionType) isType_Kind() {}

func (*Type_ObjectType) isType_Kind() {}

func (*Type_TupleType) isType_Kind() {}

func (m *Type) GetKind() isType_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Type) GetDynamicType() DynamicType {
	if x, ok := m.GetKind().(*Type_DynamicType); ok {
		return x.DynamicType
	}
	return DynamicType_DYNAMIC_TYPE
}

func (m *Type) GetPrimitiveType() PrimitiveType {
	if x, ok := m.GetKind().(*Type_PrimitiveType); ok {
		return x.PrimitiveType
	}
	return PrimitiveType_STRING
}

func (m *Type) GetCollectionType() *CollectionType {
	if x, ok := m.GetKind().(*Type_CollectionType); ok {
		return x.CollectionType
	}
	return nil
}

func (m *Type) GetObjectType() *ObjectType {
	if x, ok := m.GetKind().(*Type_ObjectType); ok {
		return x.ObjectType
	}
	return nil
}

func (m *Type) GetTupleType() *TupleType {
	if x, ok := m.GetKind().(*Type_TupleType); ok {
		return x.TupleType
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Type) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Type_DynamicType)(nil),
		(*Type_PrimitiveType)(nil),
		(*Type_CollectionType)(nil),
		(*Type_ObjectType)(nil),
		(*Type_TupleType)(nil),
	}
}

type CollectionType struct {
	Kind                 CollectionKind `protobuf:"varint,1,opt,name=kind,proto3,enum=zclconf.cty.CollectionKind" json:"kind,omitempty"`
	Element              *Type          `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *CollectionType) Reset()         { *m = CollectionType{} }
func (m *CollectionType) String() string { return proto.CompactTextString(m) }
func (*CollectionType) ProtoMessage()    {}
func (*CollectionType) Descriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{1}
}

func (m *CollectionType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionType.Unmarshal(m, b)
}
func (m *CollectionType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionType.Marshal(b, m, deterministic)
}
func (m *CollectionType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionType.Merge(m, src)
}
func (m *CollectionType) XXX_Size() int {
	return xxx_messageInfo_CollectionType.Size(m)
}
func (m *CollectionType) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionType.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionType proto.InternalMessageInfo

func (m *CollectionType) GetKind() CollectionKind {
	if m != nil {
		return m.Kind
	}
	return CollectionKind_LIST
}

func (m *CollectionType) GetElement() *Type {
	if m != nil {
		return m.Element
	}
	return nil
}

type ObjectType struct {
	Attributes           map[string]*Type `protobuf:"bytes,1,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ObjectType) Reset()         { *m = ObjectType{} }
func (m *ObjectType) String() string { return proto.CompactTextString(m) }
func (*ObjectType) ProtoMessage()    {}
func (*ObjectType) Descriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{2}
}

func (m *ObjectType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectType.Unmarshal(m, b)
}
func (m *ObjectType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectType.Marshal(b, m, deterministic)
}
func (m *ObjectType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectType.Merge(m, src)
}
func (m *ObjectType) XXX_Size() int {
	return xxx_messageInfo_ObjectType.Size(m)
}
func (m *ObjectType) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectType.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectType proto.InternalMessageInfo

func (m *ObjectType) GetAttributes() map[string]*Type {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type TupleType struct {
	Types                []*Type  `protobuf:"bytes,1,rep,name=types,proto3" json:"types,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TupleType) Reset()         { *m = TupleType{} }
func (m *TupleType) String() string { return proto.CompactTextString(m) }
func (*TupleType) ProtoMessage()    {}
func (*TupleType) Descriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{3}
}

func (m *TupleType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TupleType.Unmarshal(m, b)
}
func (m *TupleType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TupleType.Marshal(b, m, deterministic)
}
func (m *TupleType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TupleType.Merge(m, src)
}
func (m *TupleType) XXX_Size() int {
	return xxx_messageInfo_TupleType.Size(m)
}
func (m *TupleType) XXX_DiscardUnknown() {
	xxx_messageInfo_TupleType.DiscardUnknown(m)
}

var xxx_messageInfo_TupleType proto.InternalMessageInfo

func (m *TupleType) GetTypes() []*Type {
	if m != nil {
		return m.Types
	}
	return nil
}

type Value struct {
	Type *Type `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	// Types that are valid to be assigned to Kind:
	//	*Value_UnknownValue
	//	*Value_NullValue
	//	*Value_NumberValue
	//	*Value_StringValue
	//	*Value_BoolValue
	//	*Value_ListValue
	//	*Value_MapValue
	//	*Value_SetValue
	//	*Value_ObjectValue
	//	*Value_TupleValue
	Kind                 isValue_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{4}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetType() *Type {
	if m != nil {
		return m.Type
	}
	return nil
}

type isValue_Kind interface {
	isValue_Kind()
}

type Value_UnknownValue struct {
	UnknownValue UnknownValue `protobuf:"varint,2,opt,name=unknown_value,json=unknownValue,proto3,enum=zclconf.cty.UnknownValue,oneof"`
}

type Value_NullValue struct {
	NullValue NullValue `protobuf:"varint,3,opt,name=null_value,json=nullValue,proto3,enum=zclconf.cty.NullValue,oneof"`
}

type Value_NumberValue struct {
	NumberValue *NumberValue `protobuf:"bytes,4,opt,name=number_value,json=numberValue,proto3,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,5,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,6,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,7,opt,name=list_value,json=listValue,proto3,oneof"`
}

type Value_MapValue struct {
	MapValue *MapValue `protobuf:"bytes,8,opt,name=map_value,json=mapValue,proto3,oneof"`
}

type Value_SetValue struct {
	SetValue *SetValue `protobuf:"bytes,9,opt,name=set_value,json=setValue,proto3,oneof"`
}

type Value_ObjectValue struct {
	ObjectValue *ObjectValue `protobuf:"bytes,10,opt,name=object_value,json=objectValue,proto3,oneof"`
}

type Value_TupleValue struct {
	TupleValue *TupleValue `protobuf:"bytes,11,opt,name=tuple_value,json=tupleValue,proto3,oneof"`
}

func (*Value_UnknownValue) isValue_Kind() {}

func (*Value_NullValue) isValue_Kind() {}

func (*Value_NumberValue) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_BoolValue) isValue_Kind() {}

func (*Value_ListValue) isValue_Kind() {}

func (*Value_MapValue) isValue_Kind() {}

func (*Value_SetValue) isValue_Kind() {}

func (*Value_ObjectValue) isValue_Kind() {}

func (*Value_TupleValue) isValue_Kind() {}

func (m *Value) GetKind() isValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Value) GetUnknownValue() UnknownValue {
	if x, ok := m.GetKind().(*Value_UnknownValue); ok {
		return x.UnknownValue
	}
	return UnknownValue_UNKNOWN_VALUE
}

func (m *Value) GetNullValue() NullValue {
	if x, ok := m.GetKind().(*Value_NullValue); ok {
		return x.NullValue
	}
	return NullValue_NULL_VALUE
}

func (m *Value) GetNumberValue() *NumberValue {
	if x, ok := m.GetKind().(*Value_NumberValue); ok {
		return x.NumberValue
	}
	return nil
}

func (m *Value) GetStringValue() string {
	if x, ok := m.GetKind().(*Value_StringValue); ok {
		return x.StringValue
	}
	return ""
}

func (m *Value) GetBoolValue() bool {
	if x, ok := m.GetKind().(*Value_BoolValue); ok {
		return x.BoolValue
	}
	return false
}

func (m *Value) GetListValue() *ListValue {
	if x, ok := m.GetKind().(*Value_ListValue); ok {
		return x.ListValue
	}
	return nil
}

func (m *Value) GetMapValue() *MapValue {
	if x, ok := m.GetKind().(*Value_MapValue); ok {
		return x.MapValue
	}
	return nil
}

func (m *Value) GetSetValue() *SetValue {
	if x, ok := m.GetKind().(*Value_SetValue); ok {
		return x.SetValue
	}
	return nil
}

func (m *Value) GetObjectValue() *ObjectValue {
	if x, ok := m.GetKind().(*Value_ObjectValue); ok {
		return x.ObjectValue
	}
	return nil
}

func (m *Value) GetTupleValue() *TupleValue {
	if x, ok := m.GetKind().(*Value_TupleValue); ok {
		return x.TupleValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_UnknownValue)(nil),
		(*Value_NullValue)(nil),
		(*Value_NumberValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_ListValue)(nil),
		(*Value_MapValue)(nil),
		(*Value_SetValue)(nil),
		(*Value_ObjectValue)(nil),
		(*Value_TupleValue)(nil),
	}
}

type NumberValue struct {
	// Types that are valid to be assigned to Kind:
	//	*NumberValue_InfinityValue
	//	*NumberValue_IntValue
	//	*NumberValue_DoubleValue
	//	*NumberValue_StringValue
	Kind                 isNumberValue_Kind `protobuf_oneof:"kind"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *NumberValue) Reset()         { *m = NumberValue{} }
func (m *NumberValue) String() string { return proto.CompactTextString(m) }
func (*NumberValue) ProtoMessage()    {}
func (*NumberValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{5}
}

func (m *NumberValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NumberValue.Unmarshal(m, b)
}
func (m *NumberValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NumberValue.Marshal(b, m, deterministic)
}
func (m *NumberValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NumberValue.Merge(m, src)
}
func (m *NumberValue) XXX_Size() int {
	return xxx_messageInfo_NumberValue.Size(m)
}
func (m *NumberValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NumberValue.DiscardUnknown(m)
}

var xxx_messageInfo_NumberValue proto.InternalMessageInfo

type isNumberValue_Kind interface {
	isNumberValue_Kind()
}

type NumberValue_InfinityValue struct {
	InfinityValue InfinityValue `protobuf:"varint,1,opt,name=infinity_value,json=infinityValue,proto3,enum=zclconf.cty.InfinityValue,oneof"`
}

type NumberValue_IntValue struct {
	IntValue int64 `protobuf:"zigzag64,2,opt,name=int_value,json=intValue,proto3,oneof"`
}

type NumberValue_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type NumberValue_StringValue struct {
	StringValue string `protobuf:"bytes,4,opt,name=string_value,json=stringValue,proto3,oneof"`
}

func (*NumberValue_InfinityValue) isNumberValue_Kind() {}

func (*NumberValue_IntValue) isNumberValue_Kind() {}

func (*NumberValue_DoubleValue) isNumberValue_Kind() {}

func (*NumberValue_StringValue) isNumberValue_Kind() {}

func (m *NumberValue) GetKind() isNumberValue_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *NumberValue) GetInfinityValue() InfinityValue {
	if x, ok := m.GetKind().(*NumberValue_InfinityValue); ok {
		return x.InfinityValue
	}
	return InfinityValue_POSITIVE_INFINITY
}

func (m *NumberValue) GetIntValue() int64 {
	if x, ok := m.GetKind().(*NumberValue_IntValue); ok {
		return x.IntValue
	}
	return 0
}

func (m *NumberValue) GetDoubleValue() float64 {
	if x, ok := m.GetKind().(*NumberValue_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *NumberValue) GetStringValue() string {
	if x, ok := m.GetKind().(*NumberValue_StringValue); ok {
		return x.StringValue
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NumberValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NumberValue_InfinityValue)(nil),
		(*NumberValue_IntValue)(nil),
		(*NumberValue_DoubleValue)(nil),
		(*NumberValue_StringValue)(nil),
	}
}

type ListValue struct {
	Values               []*Value `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListValue) Reset()         { *m = ListValue{} }
func (m *ListValue) String() string { return proto.CompactTextString(m) }
func (*ListValue) ProtoMessage()    {}
func (*ListValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{6}
}

func (m *ListValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListValue.Unmarshal(m, b)
}
func (m *ListValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListValue.Marshal(b, m, deterministic)
}
func (m *ListValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListValue.Merge(m, src)
}
func (m *ListValue) XXX_Size() int {
	return xxx_messageInfo_ListValue.Size(m)
}
func (m *ListValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ListValue.DiscardUnknown(m)
}

var xxx_messageInfo_ListValue proto.InternalMessageInfo

func (m *ListValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

type MapValue struct {
	Values               map[string]*Value `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MapValue) Reset()         { *m = MapValue{} }
func (m *MapValue) String() string { return proto.CompactTextString(m) }
func (*MapValue) ProtoMessage()    {}
func (*MapValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{7}
}

func (m *MapValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MapValue.Unmarshal(m, b)
}
func (m *MapValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MapValue.Marshal(b, m, deterministic)
}
func (m *MapValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MapValue.Merge(m, src)
}
func (m *MapValue) XXX_Size() int {
	return xxx_messageInfo_MapValue.Size(m)
}
func (m *MapValue) XXX_DiscardUnknown() {
	xxx_messageInfo_MapValue.DiscardUnknown(m)
}

var xxx_messageInfo_MapValue proto.InternalMessageInfo

func (m *MapValue) GetValues() map[string]*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

type SetValue struct {
	Values               []*Value `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetValue) Reset()         { *m = SetValue{} }
func (m *SetValue) String() string { return proto.CompactTextString(m) }
func (*SetValue) ProtoMessage()    {}
func (*SetValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{8}
}

func (m *SetValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetValue.Unmarshal(m, b)
}
func (m *SetValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetValue.Marshal(b, m, deterministic)
}
func (m *SetValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetValue.Merge(m, src)
}
func (m *SetValue) XXX_Size() int {
	return xxx_messageInfo_SetValue.Size(m)
}
func (m *SetValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SetValue.DiscardUnknown(m)
}

var xxx_messageInfo_SetValue proto.InternalMessageInfo

func (m *SetValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

type ObjectValue struct {
	Values               map[string]*Value `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ObjectValue) Reset()         { *m = ObjectValue{} }
func (m *ObjectValue) String() string { return proto.CompactTextString(m) }
func (*ObjectValue) ProtoMessage()    {}
func (*ObjectValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{9}
}

func (m *ObjectValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectValue.Unmarshal(m, b)
}
func (m *ObjectValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectValue.Marshal(b, m, deterministic)
}
func (m *ObjectValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectValue.Merge(m, src)
}
func (m *ObjectValue) XXX_Size() int {
	return xxx_messageInfo_ObjectValue.Size(m)
}
func (m *ObjectValue) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectValue.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectValue proto.InternalMessageInfo

func (m *ObjectValue) GetValues() map[string]*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

type TupleValue struct {
	Values               []*Value `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TupleValue) Reset()         { *m = TupleValue{} }
func (m *TupleValue) String() string { return proto.CompactTextString(m) }
func (*TupleValue) ProtoMessage()    {}
func (*TupleValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_187f7a7284e9b418, []int{10}
}

func (m *TupleValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TupleValue.Unmarshal(m, b)
}
func (m *TupleValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TupleValue.Marshal(b, m, deterministic)
}
func (m *TupleValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TupleValue.Merge(m, src)
}
func (m *TupleValue) XXX_Size() int {
	return xxx_messageInfo_TupleValue.Size(m)
}
func (m *TupleValue) XXX_DiscardUnknown() {
	xxx_messageInfo_TupleValue.DiscardUnknown(m)
}

var xxx_messageInfo_TupleValue proto.InternalMessageInfo

func (m *TupleValue) GetValues() []*Value {
	if m != nil {
		return m.Values
	}
	return nil
}

func init() {
	proto.RegisterEnum("zclconf.cty.DynamicType", DynamicType_name, DynamicType_value)
	proto.RegisterEnum("zclconf.cty.PrimitiveType", PrimitiveType_name, PrimitiveType_value)
	proto.RegisterEnum("zclconf.cty.CollectionKind", CollectionKind_name, CollectionKind_value)
	proto.RegisterEnum("zclconf.cty.UnknownValue", UnknownValue_name, UnknownValue_value)
	proto.RegisterEnum("zclconf.cty.NullValue", NullValue_name, NullValue_value)
	proto.RegisterEnum("zclconf.cty.InfinityValue", InfinityValue_name, InfinityValue_value)
	proto.RegisterType((*Type)(nil), "zclconf.cty.Type")
	proto.RegisterType((*CollectionType)(nil), "zclconf.cty.CollectionType")
	proto.RegisterType((*ObjectType)(nil), "zclconf.cty.ObjectType")
	proto.RegisterMapType((map[string]*Type)(nil), "zclconf.cty.ObjectType.AttributesEntry")
	proto.RegisterType((*TupleType)(nil), "zclconf.cty.TupleType")
	proto.RegisterType((*Value)(nil), "zclconf.cty.Value")
	proto.RegisterType((*NumberValue)(nil), "zclconf.cty.NumberValue")
	proto.RegisterType((*ListValue)(nil), "zclconf.cty.ListValue")
	proto.RegisterType((*MapValue)(nil), "zclconf.cty.MapValue")
	proto.RegisterMapType((map[string]*Value)(nil), "zclconf.cty.MapValue.ValuesEntry")
	proto.RegisterType((*SetValue)(nil), "zclconf.cty.SetValue")
	proto.RegisterType((*ObjectValue)(nil), "zclconf.cty.ObjectValue")
	proto.RegisterMapType((map[string]*Value)(nil), "zclconf.cty.ObjectValue.ValuesEntry")
	proto.RegisterType((*TupleValue)(nil), "zclconf.cty.TupleValue")
}

func init() { proto.RegisterFile("cty.proto", fileDescriptor_187f7a7284e9b418) }

var fileDescriptor_187f7a7284e9b418 = []byte{
	// 953 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x41, 0x8f, 0xda, 0x46,
	0x14, 0xc6, 0xc0, 0xee, 0xe2, 0x67, 0x20, 0xde, 0x51, 0xd2, 0xd2, 0x8d, 0xaa, 0x6c, 0x68, 0xaa,
	0x20, 0xda, 0xb2, 0x2a, 0x8d, 0x9a, 0x34, 0x6a, 0xa4, 0x02, 0x21, 0x1b, 0x14, 0x30, 0xc8, 0x98,
	0xad, 0xc8, 0x05, 0x81, 0xf1, 0x6e, 0xa6, 0x6b, 0xc6, 0x16, 0x8c, 0x53, 0xb9, 0xbf, 0xa3, 0xa7,
	0x1e, 0x7b, 0xa9, 0xd4, 0xbf, 0xd1, 0x1f, 0xd0, 0x5b, 0x7f, 0x4b, 0x8f, 0xd5, 0x8c, 0xc7, 0x83,
	0xbd, 0xf2, 0x1e, 0x72, 0xe9, 0x61, 0x25, 0xbf, 0xf7, 0xbe, 0xef, 0xe3, 0xbd, 0x99, 0xf7, 0xde,
	0x2c, 0xa8, 0x36, 0x0d, 0x5b, 0xfe, 0xd6, 0xa3, 0x1e, 0xd2, 0x7e, 0xb1, 0x5d, 0xdb, 0x23, 0x97,
	0x2d, 0x9b, 0x86, 0xf5, 0xbf, 0xf3, 0x50, 0xb4, 0x42, 0xdf, 0x41, 0x2f, 0xa0, 0xbc, 0x0e, 0xc9,
	0x72, 0x83, 0xed, 0x05, 0x0d, 0x7d, 0xa7, 0xa6, 0x9c, 0x2a, 0x8d, 0x6a, 0xbb, 0xd6, 0x4a, 0x80,
	0x5b, 0x2f, 0x23, 0x00, 0xc3, 0xbf, 0xce, 0x99, 0xda, 0x7a, 0x6f, 0xa2, 0x1e, 0x54, 0xfd, 0x2d,
	0xde, 0x60, 0x8a, 0xdf, 0x3b, 0x91, 0x40, 0x9e, 0x0b, 0x9c, 0xa4, 0x04, 0x26, 0x31, 0x44, 0x48,
	0x54, 0xfc, 0xa4, 0x03, 0xbd, 0x82, 0x3b, 0xb6, 0xe7, 0xba, 0x8e, 0x4d, 0xb1, 0x47, 0x22, 0x95,
	0xc2, 0xa9, 0xd2, 0xd0, 0xda, 0xf7, 0x53, 0x2a, 0x3d, 0x89, 0x11, 0x32, 0x55, 0x3b, 0xe5, 0x41,
	0xcf, 0x41, 0xf3, 0x56, 0x3f, 0x39, 0x36, 0x8d, 0x34, 0x8a, 0x5c, 0xe3, 0xe3, 0x94, 0xc6, 0x98,
	0xc7, 0x05, 0x1f, 0x3c, 0x69, 0xa1, 0xa7, 0x00, 0x34, 0xf0, 0x5d, 0x51, 0xc4, 0x01, 0xa7, 0x7e,
	0x94, 0xa2, 0x5a, 0x2c, 0x2c, 0x98, 0x2a, 0x8d, 0x8d, 0xee, 0x21, 0x14, 0xaf, 0x31, 0x59, 0xd7,
	0x09, 0x54, 0xd3, 0x09, 0xa2, 0xb3, 0x28, 0x22, 0x8e, 0xf4, 0xb6, 0x5a, 0xde, 0x60, 0xb2, 0x36,
	0x39, 0x10, 0x7d, 0x01, 0x47, 0x8e, 0xeb, 0x6c, 0x1c, 0x42, 0xf9, 0x29, 0x6a, 0xed, 0xe3, 0x74,
	0x02, 0xa1, 0xef, 0x98, 0x31, 0xa2, 0xfe, 0x87, 0x02, 0xb0, 0xaf, 0x06, 0x9d, 0x03, 0x2c, 0x29,
	0xdd, 0xe2, 0x55, 0x40, 0x9d, 0x5d, 0x4d, 0x39, 0x2d, 0x34, 0xb4, 0xf6, 0xe3, 0x5b, 0x4a, 0x6f,
	0x75, 0x24, 0xb2, 0x4f, 0xe8, 0x36, 0x34, 0x13, 0xd4, 0x93, 0x09, 0xdc, 0xb9, 0x11, 0x46, 0x3a,
	0x14, 0xae, 0x9d, 0x90, 0xd7, 0xa1, 0x9a, 0xec, 0x13, 0x3d, 0x86, 0x83, 0xf7, 0x4b, 0x37, 0x70,
	0x6e, 0xcf, 0x33, 0x8a, 0x3f, 0xcf, 0x3f, 0x53, 0xea, 0x4f, 0x40, 0x95, 0x67, 0xc7, 0x98, 0xec,
	0x84, 0xe3, 0x14, 0xb3, 0x98, 0x3c, 0x5e, 0xff, 0xa7, 0x08, 0x07, 0x17, 0x4c, 0x03, 0x7d, 0x0e,
	0x45, 0xd9, 0x9a, 0x99, 0x0c, 0x1e, 0x46, 0x3f, 0x40, 0x25, 0x20, 0xd7, 0xc4, 0xfb, 0x99, 0x2c,
	0xf6, 0xb9, 0x55, 0xdb, 0x9f, 0xa4, 0xf0, 0xb3, 0x08, 0xc1, 0x85, 0x5f, 0xe7, 0xcc, 0x72, 0x90,
	0xb0, 0x59, 0x0f, 0x90, 0xc0, 0x75, 0x05, 0xbd, 0xc0, 0xe9, 0xe9, 0x1e, 0x30, 0x02, 0xd7, 0x8d,
	0xb9, 0x2a, 0x89, 0x0d, 0x36, 0x44, 0x24, 0xd8, 0xac, 0x9c, 0xad, 0xa0, 0x46, 0x9d, 0x57, 0xbb,
	0x41, 0x65, 0x80, 0x98, 0xac, 0x91, 0xbd, 0x89, 0x3e, 0x83, 0xf2, 0x8e, 0x6e, 0x31, 0xb9, 0x12,
	0x74, 0xd6, 0x7d, 0x2a, 0x03, 0x45, 0xde, 0x08, 0xf4, 0x00, 0x60, 0xe5, 0x79, 0x71, 0x72, 0x87,
	0xa7, 0x4a, 0xa3, 0xc4, 0x92, 0x60, 0x3e, 0x99, 0xbd, 0x8b, 0x77, 0x54, 0x00, 0x8e, 0x32, 0x3a,
	0x78, 0x88, 0x77, 0x54, 0x66, 0xef, 0xc6, 0x06, 0x7a, 0x02, 0xea, 0x66, 0xe9, 0x0b, 0x5e, 0x89,
	0xf3, 0xee, 0xa5, 0x78, 0xa3, 0xa5, 0x1f, 0xd3, 0x4a, 0x1b, 0xf1, 0xcd, 0x58, 0x3b, 0x27, 0xfe,
	0x35, 0x35, 0x83, 0x35, 0x75, 0xe4, 0x8f, 0x95, 0x76, 0xe2, 0x9b, 0x9d, 0x94, 0x18, 0xd1, 0x88,
	0x08, 0x19, 0x27, 0x15, 0x35, 0xaa, 0x3c, 0x29, 0x6f, 0x6f, 0xb2, 0x09, 0x8f, 0xa6, 0x34, 0x62,
	0x6b, 0x19, 0x13, 0xce, 0x5b, 0x2d, 0x26, 0x47, 0x33, 0xcd, 0x2d, 0x39, 0xa8, 0x7f, 0x29, 0xa0,
	0x25, 0x2e, 0x83, 0xad, 0x30, 0x4c, 0x2e, 0x31, 0xc1, 0x34, 0x14, 0xb2, 0x4a, 0xc6, 0x0a, 0x1b,
	0x08, 0x48, 0xac, 0x5c, 0xc1, 0x49, 0x07, 0xfa, 0x14, 0x54, 0x4c, 0x68, 0xa2, 0xf1, 0x10, 0x2b,
	0x1b, 0x13, 0x2a, 0x6f, 0x78, 0xed, 0x05, 0x2b, 0x99, 0x38, 0xeb, 0x2d, 0x85, 0xef, 0x52, 0xee,
	0xcd, 0x6e, 0x83, 0x62, 0x46, 0x1b, 0xc8, 0x2a, 0x9e, 0x82, 0x2a, 0xaf, 0x13, 0x35, 0xe1, 0x90,
	0x53, 0x76, 0xb5, 0x3c, 0x9f, 0x2a, 0x94, 0x4a, 0x9d, 0x63, 0x4c, 0x81, 0xa8, 0xff, 0xaa, 0x40,
	0x29, 0xbe, 0x50, 0xf4, 0xdd, 0x0d, 0xe2, 0xc3, 0xcc, 0x7b, 0x8f, 0x14, 0xc4, 0xae, 0x10, 0x84,
	0x93, 0x11, 0x68, 0x09, 0x77, 0xc6, 0x8e, 0x68, 0xa4, 0x77, 0x44, 0x56, 0x4e, 0x89, 0x25, 0xf1,
	0x2d, 0x94, 0xe2, 0x86, 0xf9, 0xa0, 0x72, 0x7e, 0x53, 0x40, 0x4b, 0x34, 0x0c, 0xfa, 0x5e, 0x72,
	0xa3, 0x05, 0xf3, 0xe8, 0xb6, 0xd6, 0xfa, 0x3f, 0x8a, 0x7a, 0x06, 0xb0, 0x6f, 0xc7, 0x0f, 0x29,
	0xab, 0xf9, 0x00, 0xb4, 0xc4, 0xab, 0x8b, 0x74, 0x28, 0xbf, 0x9c, 0x1b, 0x9d, 0xd1, 0xa0, 0xb7,
	0xb0, 0xe6, 0x93, 0xbe, 0x9e, 0x6b, 0x7e, 0x0d, 0x95, 0xd4, 0xab, 0x8a, 0x00, 0x0e, 0xa7, 0x96,
	0x39, 0x30, 0xce, 0xf5, 0x1c, 0xfb, 0x36, 0x66, 0xa3, 0x6e, 0xdf, 0xd4, 0x15, 0x54, 0x82, 0x62,
	0x77, 0x3c, 0x1e, 0xea, 0xf9, 0xe6, 0x97, 0xc9, 0x17, 0x8a, 0x3d, 0x3b, 0x2c, 0x36, 0x1c, 0x4c,
	0x2d, 0x3d, 0x87, 0x8e, 0xa0, 0x30, 0xea, 0x4c, 0x74, 0x85, 0x7d, 0x4c, 0xfb, 0x96, 0x9e, 0x6f,
	0x3e, 0x84, 0x72, 0x72, 0x59, 0xa2, 0x63, 0xa8, 0xcc, 0x8c, 0x37, 0xc6, 0xf8, 0x47, 0x63, 0x71,
	0xd1, 0x19, 0xce, 0x58, 0x0e, 0xf7, 0x41, 0x95, 0x0b, 0x11, 0x55, 0x01, 0x8c, 0xd9, 0x70, 0x28,
	0x83, 0x2f, 0xa0, 0x92, 0x9a, 0x19, 0x74, 0x0f, 0x8e, 0x27, 0xe3, 0xe9, 0xc0, 0x1a, 0x5c, 0xf4,
	0x17, 0x03, 0xe3, 0xd5, 0xc0, 0x18, 0x58, 0x73, 0x3d, 0xc7, 0xdc, 0x46, 0xff, 0xbc, 0x93, 0x76,
	0x2b, 0xdd, 0x77, 0x70, 0xd7, 0xf6, 0x36, 0xf2, 0x84, 0xf8, 0xbf, 0x30, 0xab, 0xe0, 0xb2, 0x5b,
	0xea, 0x59, 0xf3, 0x09, 0x33, 0x26, 0xca, 0xdb, 0x47, 0x57, 0x98, 0xbe, 0x0b, 0x56, 0x2d, 0xdb,
	0xdb, 0x9c, 0x09, 0xe0, 0xd9, 0x95, 0xf7, 0x95, 0x4d, 0xc3, 0x33, 0xf6, 0xe7, 0xf3, 0x87, 0xe4,
	0x5f, 0x45, 0xf9, 0x3d, 0x5f, 0xe8, 0x59, 0xf3, 0x3f, 0xf3, 0x77, 0xdf, 0xf6, 0x86, 0x3d, 0xa6,
	0x36, 0x11, 0x6a, 0xad, 0x9e, 0x35, 0x5f, 0x1d, 0x72, 0xed, 0x6f, 0xfe, 0x0b, 0x00, 0x00, 0xff,
	0xff, 0xb0, 0xc1, 0x6d, 0x5e, 0x2b, 0x09, 0x00, 0x00,
}
