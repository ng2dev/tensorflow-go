// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/framework/attr_value.proto

package attr_value_go_proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	tensor_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_go_proto"
	tensor_shape_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/tensor_shape_go_proto"
	types_go_proto "github.com/tensorflow/tensorflow/tensorflow/go/core/framework/types_go_proto"
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

// Protocol buffer representing the value for an attr used to configure an Op.
// Comment indicates the corresponding attr type.  Only the field matching the
// attr type may be filled.
type AttrValue struct {
	// Types that are valid to be assigned to Value:
	//	*AttrValue_S
	//	*AttrValue_I
	//	*AttrValue_F
	//	*AttrValue_B
	//	*AttrValue_Type
	//	*AttrValue_Shape
	//	*AttrValue_Tensor
	//	*AttrValue_List
	//	*AttrValue_Func
	//	*AttrValue_Placeholder
	Value                isAttrValue_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AttrValue) Reset()         { *m = AttrValue{} }
func (m *AttrValue) String() string { return proto.CompactTextString(m) }
func (*AttrValue) ProtoMessage()    {}
func (*AttrValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_06e758bf81984406, []int{0}
}

func (m *AttrValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttrValue.Unmarshal(m, b)
}
func (m *AttrValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttrValue.Marshal(b, m, deterministic)
}
func (m *AttrValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttrValue.Merge(m, src)
}
func (m *AttrValue) XXX_Size() int {
	return xxx_messageInfo_AttrValue.Size(m)
}
func (m *AttrValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttrValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttrValue proto.InternalMessageInfo

type isAttrValue_Value interface {
	isAttrValue_Value()
}

type AttrValue_S struct {
	S []byte `protobuf:"bytes,2,opt,name=s,proto3,oneof"`
}

type AttrValue_I struct {
	I int64 `protobuf:"varint,3,opt,name=i,proto3,oneof"`
}

type AttrValue_F struct {
	F float32 `protobuf:"fixed32,4,opt,name=f,proto3,oneof"`
}

type AttrValue_B struct {
	B bool `protobuf:"varint,5,opt,name=b,proto3,oneof"`
}

type AttrValue_Type struct {
	Type types_go_proto.DataType `protobuf:"varint,6,opt,name=type,proto3,enum=tensorflow.DataType,oneof"`
}

type AttrValue_Shape struct {
	Shape *tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,7,opt,name=shape,proto3,oneof"`
}

type AttrValue_Tensor struct {
	Tensor *tensor_go_proto.TensorProto `protobuf:"bytes,8,opt,name=tensor,proto3,oneof"`
}

type AttrValue_List struct {
	List *AttrValue_ListValue `protobuf:"bytes,1,opt,name=list,proto3,oneof"`
}

type AttrValue_Func struct {
	Func *NameAttrList `protobuf:"bytes,10,opt,name=func,proto3,oneof"`
}

type AttrValue_Placeholder struct {
	Placeholder string `protobuf:"bytes,9,opt,name=placeholder,proto3,oneof"`
}

func (*AttrValue_S) isAttrValue_Value() {}

func (*AttrValue_I) isAttrValue_Value() {}

func (*AttrValue_F) isAttrValue_Value() {}

func (*AttrValue_B) isAttrValue_Value() {}

func (*AttrValue_Type) isAttrValue_Value() {}

func (*AttrValue_Shape) isAttrValue_Value() {}

func (*AttrValue_Tensor) isAttrValue_Value() {}

func (*AttrValue_List) isAttrValue_Value() {}

func (*AttrValue_Func) isAttrValue_Value() {}

func (*AttrValue_Placeholder) isAttrValue_Value() {}

func (m *AttrValue) GetValue() isAttrValue_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *AttrValue) GetS() []byte {
	if x, ok := m.GetValue().(*AttrValue_S); ok {
		return x.S
	}
	return nil
}

func (m *AttrValue) GetI() int64 {
	if x, ok := m.GetValue().(*AttrValue_I); ok {
		return x.I
	}
	return 0
}

func (m *AttrValue) GetF() float32 {
	if x, ok := m.GetValue().(*AttrValue_F); ok {
		return x.F
	}
	return 0
}

func (m *AttrValue) GetB() bool {
	if x, ok := m.GetValue().(*AttrValue_B); ok {
		return x.B
	}
	return false
}

func (m *AttrValue) GetType() types_go_proto.DataType {
	if x, ok := m.GetValue().(*AttrValue_Type); ok {
		return x.Type
	}
	return types_go_proto.DataType_DT_INVALID
}

func (m *AttrValue) GetShape() *tensor_shape_go_proto.TensorShapeProto {
	if x, ok := m.GetValue().(*AttrValue_Shape); ok {
		return x.Shape
	}
	return nil
}

func (m *AttrValue) GetTensor() *tensor_go_proto.TensorProto {
	if x, ok := m.GetValue().(*AttrValue_Tensor); ok {
		return x.Tensor
	}
	return nil
}

func (m *AttrValue) GetList() *AttrValue_ListValue {
	if x, ok := m.GetValue().(*AttrValue_List); ok {
		return x.List
	}
	return nil
}

func (m *AttrValue) GetFunc() *NameAttrList {
	if x, ok := m.GetValue().(*AttrValue_Func); ok {
		return x.Func
	}
	return nil
}

func (m *AttrValue) GetPlaceholder() string {
	if x, ok := m.GetValue().(*AttrValue_Placeholder); ok {
		return x.Placeholder
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AttrValue) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AttrValue_S)(nil),
		(*AttrValue_I)(nil),
		(*AttrValue_F)(nil),
		(*AttrValue_B)(nil),
		(*AttrValue_Type)(nil),
		(*AttrValue_Shape)(nil),
		(*AttrValue_Tensor)(nil),
		(*AttrValue_List)(nil),
		(*AttrValue_Func)(nil),
		(*AttrValue_Placeholder)(nil),
	}
}

// LINT.IfChange
type AttrValue_ListValue struct {
	S                    [][]byte                                  `protobuf:"bytes,2,rep,name=s,proto3" json:"s,omitempty"`
	I                    []int64                                   `protobuf:"varint,3,rep,packed,name=i,proto3" json:"i,omitempty"`
	F                    []float32                                 `protobuf:"fixed32,4,rep,packed,name=f,proto3" json:"f,omitempty"`
	B                    []bool                                    `protobuf:"varint,5,rep,packed,name=b,proto3" json:"b,omitempty"`
	Type                 []types_go_proto.DataType                 `protobuf:"varint,6,rep,packed,name=type,proto3,enum=tensorflow.DataType" json:"type,omitempty"`
	Shape                []*tensor_shape_go_proto.TensorShapeProto `protobuf:"bytes,7,rep,name=shape,proto3" json:"shape,omitempty"`
	Tensor               []*tensor_go_proto.TensorProto            `protobuf:"bytes,8,rep,name=tensor,proto3" json:"tensor,omitempty"`
	Func                 []*NameAttrList                           `protobuf:"bytes,9,rep,name=func,proto3" json:"func,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                  `json:"-"`
	XXX_unrecognized     []byte                                    `json:"-"`
	XXX_sizecache        int32                                     `json:"-"`
}

func (m *AttrValue_ListValue) Reset()         { *m = AttrValue_ListValue{} }
func (m *AttrValue_ListValue) String() string { return proto.CompactTextString(m) }
func (*AttrValue_ListValue) ProtoMessage()    {}
func (*AttrValue_ListValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_06e758bf81984406, []int{0, 0}
}

func (m *AttrValue_ListValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AttrValue_ListValue.Unmarshal(m, b)
}
func (m *AttrValue_ListValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AttrValue_ListValue.Marshal(b, m, deterministic)
}
func (m *AttrValue_ListValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AttrValue_ListValue.Merge(m, src)
}
func (m *AttrValue_ListValue) XXX_Size() int {
	return xxx_messageInfo_AttrValue_ListValue.Size(m)
}
func (m *AttrValue_ListValue) XXX_DiscardUnknown() {
	xxx_messageInfo_AttrValue_ListValue.DiscardUnknown(m)
}

var xxx_messageInfo_AttrValue_ListValue proto.InternalMessageInfo

func (m *AttrValue_ListValue) GetS() [][]byte {
	if m != nil {
		return m.S
	}
	return nil
}

func (m *AttrValue_ListValue) GetI() []int64 {
	if m != nil {
		return m.I
	}
	return nil
}

func (m *AttrValue_ListValue) GetF() []float32 {
	if m != nil {
		return m.F
	}
	return nil
}

func (m *AttrValue_ListValue) GetB() []bool {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *AttrValue_ListValue) GetType() []types_go_proto.DataType {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *AttrValue_ListValue) GetShape() []*tensor_shape_go_proto.TensorShapeProto {
	if m != nil {
		return m.Shape
	}
	return nil
}

func (m *AttrValue_ListValue) GetTensor() []*tensor_go_proto.TensorProto {
	if m != nil {
		return m.Tensor
	}
	return nil
}

func (m *AttrValue_ListValue) GetFunc() []*NameAttrList {
	if m != nil {
		return m.Func
	}
	return nil
}

// A list of attr names and their values. The whole list is attached
// with a string name.  E.g., MatMul[T=float].
type NameAttrList struct {
	Name                 string                `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Attr                 map[string]*AttrValue `protobuf:"bytes,2,rep,name=attr,proto3" json:"attr,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NameAttrList) Reset()         { *m = NameAttrList{} }
func (m *NameAttrList) String() string { return proto.CompactTextString(m) }
func (*NameAttrList) ProtoMessage()    {}
func (*NameAttrList) Descriptor() ([]byte, []int) {
	return fileDescriptor_06e758bf81984406, []int{1}
}

func (m *NameAttrList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NameAttrList.Unmarshal(m, b)
}
func (m *NameAttrList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NameAttrList.Marshal(b, m, deterministic)
}
func (m *NameAttrList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NameAttrList.Merge(m, src)
}
func (m *NameAttrList) XXX_Size() int {
	return xxx_messageInfo_NameAttrList.Size(m)
}
func (m *NameAttrList) XXX_DiscardUnknown() {
	xxx_messageInfo_NameAttrList.DiscardUnknown(m)
}

var xxx_messageInfo_NameAttrList proto.InternalMessageInfo

func (m *NameAttrList) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *NameAttrList) GetAttr() map[string]*AttrValue {
	if m != nil {
		return m.Attr
	}
	return nil
}

func init() {
	proto.RegisterType((*AttrValue)(nil), "tensorflow.AttrValue")
	proto.RegisterType((*AttrValue_ListValue)(nil), "tensorflow.AttrValue.ListValue")
	proto.RegisterType((*NameAttrList)(nil), "tensorflow.NameAttrList")
	proto.RegisterMapType((map[string]*AttrValue)(nil), "tensorflow.NameAttrList.AttrEntry")
}

func init() {
	proto.RegisterFile("tensorflow/core/framework/attr_value.proto", fileDescriptor_06e758bf81984406)
}

var fileDescriptor_06e758bf81984406 = []byte{
	// 526 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x94, 0x4f, 0x8b, 0xdb, 0x3c,
	0x10, 0xc6, 0x23, 0xcb, 0xc9, 0xc6, 0x4a, 0xd8, 0x37, 0x88, 0xb7, 0x54, 0x84, 0x42, 0x4d, 0xa0,
	0x45, 0x6c, 0x83, 0x43, 0xd3, 0x3f, 0x94, 0xde, 0x1a, 0x5a, 0xc8, 0xa1, 0x2c, 0x5b, 0x77, 0xe9,
	0xa1, 0x97, 0xe0, 0xa4, 0x72, 0x12, 0x36, 0x89, 0x8c, 0xac, 0x74, 0xc9, 0xb9, 0x87, 0x5e, 0xfb,
	0x39, 0xfa, 0x09, 0x7b, 0x2c, 0x33, 0xf2, 0x3a, 0x86, 0xae, 0x77, 0x6f, 0x33, 0xa3, 0xe7, 0x91,
	0xc6, 0x3f, 0x69, 0xcc, 0xce, 0xac, 0xda, 0xe5, 0xda, 0xa4, 0x1b, 0x7d, 0x3d, 0x5a, 0x68, 0xa3,
	0x46, 0xa9, 0x49, 0xb6, 0xea, 0x5a, 0x9b, 0xab, 0x51, 0x62, 0xad, 0x99, 0x7d, 0x4f, 0x36, 0x7b,
	0x15, 0x65, 0x46, 0x5b, 0xcd, 0xd9, 0x51, 0xdb, 0x7f, 0x5a, 0xef, 0x73, 0x2b, 0xce, 0xd3, 0x1f,
	0xde, 0xa7, 0x9b, 0xe5, 0xab, 0x24, 0x2b, 0x4e, 0xe8, 0x3f, 0xb9, 0x43, 0x7d, 0xc8, 0x54, 0xee,
	0x64, 0x83, 0x9f, 0x4d, 0x16, 0xbc, 0xb3, 0xd6, 0x7c, 0x81, 0xe6, 0xf8, 0x29, 0x23, 0xb9, 0xf0,
	0x42, 0x22, 0xbb, 0xd3, 0x46, 0x4c, 0x72, 0xc8, 0xd7, 0x82, 0x86, 0x44, 0x52, 0xc8, 0xd7, 0x90,
	0xa7, 0xc2, 0x0f, 0x89, 0xf4, 0x20, 0x4f, 0x21, 0x9f, 0x8b, 0x66, 0x48, 0x64, 0x1b, 0xf2, 0x39,
	0x3f, 0x63, 0x3e, 0x6c, 0x2e, 0x5a, 0x21, 0x91, 0xa7, 0xe3, 0xff, 0xa3, 0x63, 0x0f, 0xd1, 0xfb,
	0xc4, 0x26, 0x97, 0x87, 0x4c, 0x4d, 0x1b, 0x31, 0x6a, 0xf8, 0x4b, 0xd6, 0xc4, 0x7e, 0xc5, 0x49,
	0x48, 0x64, 0x67, 0xfc, 0xa8, 0x2a, 0xbe, 0xc4, 0xf0, 0x33, 0x2c, 0x5f, 0x40, 0x9b, 0xd3, 0x46,
	0xec, 0xc4, 0xfc, 0x39, 0x6b, 0x39, 0x9d, 0x68, 0xa3, 0xed, 0xe1, 0xbf, 0xb6, 0x1b, 0x47, 0x21,
	0xe4, 0xaf, 0x98, 0xbf, 0x59, 0xe7, 0x56, 0x10, 0x34, 0x3c, 0xae, 0x1a, 0xca, 0x2f, 0x8f, 0x3e,
	0xae, 0x73, 0x8b, 0x11, 0xf4, 0x07, 0x72, 0x1e, 0x31, 0x3f, 0xdd, 0xef, 0x16, 0x82, 0xa1, 0x4d,
	0x54, 0x6d, 0xe7, 0xc9, 0x56, 0x81, 0x15, 0x4c, 0xa0, 0x07, 0x1d, 0x1f, 0xb0, 0x4e, 0xb6, 0x49,
	0x16, 0x6a, 0xa5, 0x37, 0xdf, 0x94, 0x11, 0x41, 0x48, 0x64, 0x30, 0x6d, 0xc4, 0xd5, 0x62, 0xff,
	0x97, 0xc7, 0x82, 0xf2, 0x24, 0xde, 0x75, 0xb4, 0xa9, 0xec, 0x02, 0xeb, 0x9e, 0x63, 0x4d, 0x25,
	0x9d, 0x78, 0x3d, 0x02, 0xb4, 0x7b, 0x8e, 0x36, 0x95, 0x9e, 0xab, 0xa4, 0x50, 0x01, 0xde, 0x54,
	0xb6, 0x5d, 0x65, 0xce, 0x87, 0x25, 0x71, 0x5a, 0x47, 0x1c, 0xa5, 0x8e, 0xf9, 0xf8, 0xc8, 0x9c,
	0xde, 0xc7, 0xfc, 0x86, 0xf8, 0xa8, 0x42, 0x9c, 0xde, 0x41, 0xbc, 0xe4, 0x3d, 0x2c, 0xc0, 0x05,
	0x28, 0xaf, 0x05, 0xe7, 0xb0, 0x4d, 0x4e, 0x58, 0x13, 0x07, 0x63, 0xf0, 0x9b, 0xb0, 0x6e, 0x75,
	0x9d, 0x73, 0xe6, 0xef, 0x92, 0xad, 0xc2, 0x7b, 0x0b, 0x62, 0x8c, 0xf9, 0x6b, 0xe6, 0xc3, 0x2c,
	0x21, 0xb5, 0xce, 0x78, 0x50, 0xb7, 0x37, 0x5e, 0xec, 0x87, 0x9d, 0x35, 0x87, 0x18, 0xf5, 0xfd,
	0x73, 0xf7, 0xca, 0xb1, 0xc4, 0x7b, 0x8c, 0x5e, 0xa9, 0x43, 0xb1, 0x2f, 0x84, 0xfc, 0x59, 0xd1,
	0x04, 0xbe, 0xfd, 0xce, 0xf8, 0xc1, 0xad, 0x6f, 0x24, 0x76, 0x9a, 0xb7, 0xde, 0x1b, 0x32, 0xf9,
	0x41, 0x98, 0xd0, 0x66, 0x59, 0xd5, 0x95, 0xf3, 0x35, 0xf9, 0xaf, 0xb4, 0x20, 0x98, 0xfc, 0x82,
	0x7c, 0xfd, 0xb4, 0x5c, 0xdb, 0xd5, 0x7e, 0x1e, 0x2d, 0xf4, 0x76, 0x54, 0x19, 0xcc, 0xdb, 0xc3,
	0xa5, 0xae, 0xff, 0x7f, 0xcc, 0x96, 0x7a, 0x86, 0x93, 0xfb, 0x87, 0x90, 0x79, 0x0b, 0xa3, 0x17,
	0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x43, 0xf0, 0xb3, 0x7a, 0x04, 0x00, 0x00,
}
