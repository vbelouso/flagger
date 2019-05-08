// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/transformation.proto

package transformation

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type RouteTransformations struct {
	RequestTransformation  *Transformation `protobuf:"bytes,1,opt,name=request_transformation,json=requestTransformation,proto3" json:"request_transformation,omitempty"`
	ResponseTransformation *Transformation `protobuf:"bytes,2,opt,name=response_transformation,json=responseTransformation,proto3" json:"response_transformation,omitempty"`
	XXX_NoUnkeyedLiteral   struct{}        `json:"-"`
	XXX_unrecognized       []byte          `json:"-"`
	XXX_sizecache          int32           `json:"-"`
}

func (m *RouteTransformations) Reset()         { *m = RouteTransformations{} }
func (m *RouteTransformations) String() string { return proto.CompactTextString(m) }
func (*RouteTransformations) ProtoMessage()    {}
func (*RouteTransformations) Descriptor() ([]byte, []int) {
	return fileDescriptor_201f67ff59830de4, []int{0}
}
func (m *RouteTransformations) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RouteTransformations.Unmarshal(m, b)
}
func (m *RouteTransformations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RouteTransformations.Marshal(b, m, deterministic)
}
func (m *RouteTransformations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RouteTransformations.Merge(m, src)
}
func (m *RouteTransformations) XXX_Size() int {
	return xxx_messageInfo_RouteTransformations.Size(m)
}
func (m *RouteTransformations) XXX_DiscardUnknown() {
	xxx_messageInfo_RouteTransformations.DiscardUnknown(m)
}

var xxx_messageInfo_RouteTransformations proto.InternalMessageInfo

func (m *RouteTransformations) GetRequestTransformation() *Transformation {
	if m != nil {
		return m.RequestTransformation
	}
	return nil
}

func (m *RouteTransformations) GetResponseTransformation() *Transformation {
	if m != nil {
		return m.ResponseTransformation
	}
	return nil
}

// [#proto-status: experimental]
type Transformation struct {
	// Template is in the transformed request language domain
	// currently both are JSON
	//
	// Types that are valid to be assigned to TransformationType:
	//	*Transformation_TransformationTemplate
	//	*Transformation_HeaderBodyTransform
	TransformationType   isTransformation_TransformationType `protobuf_oneof:"transformation_type"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *Transformation) Reset()         { *m = Transformation{} }
func (m *Transformation) String() string { return proto.CompactTextString(m) }
func (*Transformation) ProtoMessage()    {}
func (*Transformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_201f67ff59830de4, []int{1}
}
func (m *Transformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Transformation.Unmarshal(m, b)
}
func (m *Transformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Transformation.Marshal(b, m, deterministic)
}
func (m *Transformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Transformation.Merge(m, src)
}
func (m *Transformation) XXX_Size() int {
	return xxx_messageInfo_Transformation.Size(m)
}
func (m *Transformation) XXX_DiscardUnknown() {
	xxx_messageInfo_Transformation.DiscardUnknown(m)
}

var xxx_messageInfo_Transformation proto.InternalMessageInfo

type isTransformation_TransformationType interface {
	isTransformation_TransformationType()
	Equal(interface{}) bool
}

type Transformation_TransformationTemplate struct {
	TransformationTemplate *TransformationTemplate `protobuf:"bytes,1,opt,name=transformation_template,json=transformationTemplate,proto3,oneof"`
}
type Transformation_HeaderBodyTransform struct {
	HeaderBodyTransform *HeaderBodyTransform `protobuf:"bytes,2,opt,name=header_body_transform,json=headerBodyTransform,proto3,oneof"`
}

func (*Transformation_TransformationTemplate) isTransformation_TransformationType() {}
func (*Transformation_HeaderBodyTransform) isTransformation_TransformationType()    {}

func (m *Transformation) GetTransformationType() isTransformation_TransformationType {
	if m != nil {
		return m.TransformationType
	}
	return nil
}

func (m *Transformation) GetTransformationTemplate() *TransformationTemplate {
	if x, ok := m.GetTransformationType().(*Transformation_TransformationTemplate); ok {
		return x.TransformationTemplate
	}
	return nil
}

func (m *Transformation) GetHeaderBodyTransform() *HeaderBodyTransform {
	if x, ok := m.GetTransformationType().(*Transformation_HeaderBodyTransform); ok {
		return x.HeaderBodyTransform
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Transformation) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Transformation_OneofMarshaler, _Transformation_OneofUnmarshaler, _Transformation_OneofSizer, []interface{}{
		(*Transformation_TransformationTemplate)(nil),
		(*Transformation_HeaderBodyTransform)(nil),
	}
}

func _Transformation_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Transformation)
	// transformation_type
	switch x := m.TransformationType.(type) {
	case *Transformation_TransformationTemplate:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.TransformationTemplate); err != nil {
			return err
		}
	case *Transformation_HeaderBodyTransform:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.HeaderBodyTransform); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Transformation.TransformationType has unexpected type %T", x)
	}
	return nil
}

func _Transformation_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Transformation)
	switch tag {
	case 1: // transformation_type.transformation_template
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TransformationTemplate)
		err := b.DecodeMessage(msg)
		m.TransformationType = &Transformation_TransformationTemplate{msg}
		return true, err
	case 2: // transformation_type.header_body_transform
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(HeaderBodyTransform)
		err := b.DecodeMessage(msg)
		m.TransformationType = &Transformation_HeaderBodyTransform{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Transformation_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Transformation)
	// transformation_type
	switch x := m.TransformationType.(type) {
	case *Transformation_TransformationTemplate:
		s := proto.Size(x.TransformationTemplate)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Transformation_HeaderBodyTransform:
		s := proto.Size(x.HeaderBodyTransform)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type Extraction struct {
	Header string `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// what information to extract. if extraction fails the result is
	// an empty value.
	Regex                string   `protobuf:"bytes,2,opt,name=regex,proto3" json:"regex,omitempty"`
	Subgroup             uint32   `protobuf:"varint,3,opt,name=subgroup,proto3" json:"subgroup,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Extraction) Reset()         { *m = Extraction{} }
func (m *Extraction) String() string { return proto.CompactTextString(m) }
func (*Extraction) ProtoMessage()    {}
func (*Extraction) Descriptor() ([]byte, []int) {
	return fileDescriptor_201f67ff59830de4, []int{2}
}
func (m *Extraction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Extraction.Unmarshal(m, b)
}
func (m *Extraction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Extraction.Marshal(b, m, deterministic)
}
func (m *Extraction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Extraction.Merge(m, src)
}
func (m *Extraction) XXX_Size() int {
	return xxx_messageInfo_Extraction.Size(m)
}
func (m *Extraction) XXX_DiscardUnknown() {
	xxx_messageInfo_Extraction.DiscardUnknown(m)
}

var xxx_messageInfo_Extraction proto.InternalMessageInfo

func (m *Extraction) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *Extraction) GetRegex() string {
	if m != nil {
		return m.Regex
	}
	return ""
}

func (m *Extraction) GetSubgroup() uint32 {
	if m != nil {
		return m.Subgroup
	}
	return 0
}

type TransformationTemplate struct {
	AdvancedTemplates bool `protobuf:"varint,1,opt,name=advanced_templates,json=advancedTemplates,proto3" json:"advanced_templates,omitempty"`
	// Extractors are in the origin request language domain
	Extractors map[string]*Extraction   `protobuf:"bytes,2,rep,name=extractors,proto3" json:"extractors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Headers    map[string]*InjaTemplate `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Types that are valid to be assigned to BodyTransformation:
	//	*TransformationTemplate_Body
	//	*TransformationTemplate_Passthrough
	//	*TransformationTemplate_MergeExtractorsToBody
	BodyTransformation   isTransformationTemplate_BodyTransformation `protobuf_oneof:"body_transformation"`
	XXX_NoUnkeyedLiteral struct{}                                    `json:"-"`
	XXX_unrecognized     []byte                                      `json:"-"`
	XXX_sizecache        int32                                       `json:"-"`
}

func (m *TransformationTemplate) Reset()         { *m = TransformationTemplate{} }
func (m *TransformationTemplate) String() string { return proto.CompactTextString(m) }
func (*TransformationTemplate) ProtoMessage()    {}
func (*TransformationTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_201f67ff59830de4, []int{3}
}
func (m *TransformationTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransformationTemplate.Unmarshal(m, b)
}
func (m *TransformationTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransformationTemplate.Marshal(b, m, deterministic)
}
func (m *TransformationTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransformationTemplate.Merge(m, src)
}
func (m *TransformationTemplate) XXX_Size() int {
	return xxx_messageInfo_TransformationTemplate.Size(m)
}
func (m *TransformationTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_TransformationTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_TransformationTemplate proto.InternalMessageInfo

type isTransformationTemplate_BodyTransformation interface {
	isTransformationTemplate_BodyTransformation()
	Equal(interface{}) bool
}

type TransformationTemplate_Body struct {
	Body *InjaTemplate `protobuf:"bytes,4,opt,name=body,proto3,oneof"`
}
type TransformationTemplate_Passthrough struct {
	Passthrough *Passthrough `protobuf:"bytes,5,opt,name=passthrough,proto3,oneof"`
}
type TransformationTemplate_MergeExtractorsToBody struct {
	MergeExtractorsToBody *MergeExtractorsToBody `protobuf:"bytes,6,opt,name=merge_extractors_to_body,json=mergeExtractorsToBody,proto3,oneof"`
}

func (*TransformationTemplate_Body) isTransformationTemplate_BodyTransformation()                  {}
func (*TransformationTemplate_Passthrough) isTransformationTemplate_BodyTransformation()           {}
func (*TransformationTemplate_MergeExtractorsToBody) isTransformationTemplate_BodyTransformation() {}

func (m *TransformationTemplate) GetBodyTransformation() isTransformationTemplate_BodyTransformation {
	if m != nil {
		return m.BodyTransformation
	}
	return nil
}

func (m *TransformationTemplate) GetAdvancedTemplates() bool {
	if m != nil {
		return m.AdvancedTemplates
	}
	return false
}

func (m *TransformationTemplate) GetExtractors() map[string]*Extraction {
	if m != nil {
		return m.Extractors
	}
	return nil
}

func (m *TransformationTemplate) GetHeaders() map[string]*InjaTemplate {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *TransformationTemplate) GetBody() *InjaTemplate {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_Body); ok {
		return x.Body
	}
	return nil
}

func (m *TransformationTemplate) GetPassthrough() *Passthrough {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_Passthrough); ok {
		return x.Passthrough
	}
	return nil
}

func (m *TransformationTemplate) GetMergeExtractorsToBody() *MergeExtractorsToBody {
	if x, ok := m.GetBodyTransformation().(*TransformationTemplate_MergeExtractorsToBody); ok {
		return x.MergeExtractorsToBody
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*TransformationTemplate) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _TransformationTemplate_OneofMarshaler, _TransformationTemplate_OneofUnmarshaler, _TransformationTemplate_OneofSizer, []interface{}{
		(*TransformationTemplate_Body)(nil),
		(*TransformationTemplate_Passthrough)(nil),
		(*TransformationTemplate_MergeExtractorsToBody)(nil),
	}
}

func _TransformationTemplate_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*TransformationTemplate)
	// body_transformation
	switch x := m.BodyTransformation.(type) {
	case *TransformationTemplate_Body:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Body); err != nil {
			return err
		}
	case *TransformationTemplate_Passthrough:
		_ = b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Passthrough); err != nil {
			return err
		}
	case *TransformationTemplate_MergeExtractorsToBody:
		_ = b.EncodeVarint(6<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.MergeExtractorsToBody); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("TransformationTemplate.BodyTransformation has unexpected type %T", x)
	}
	return nil
}

func _TransformationTemplate_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*TransformationTemplate)
	switch tag {
	case 4: // body_transformation.body
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(InjaTemplate)
		err := b.DecodeMessage(msg)
		m.BodyTransformation = &TransformationTemplate_Body{msg}
		return true, err
	case 5: // body_transformation.passthrough
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Passthrough)
		err := b.DecodeMessage(msg)
		m.BodyTransformation = &TransformationTemplate_Passthrough{msg}
		return true, err
	case 6: // body_transformation.merge_extractors_to_body
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(MergeExtractorsToBody)
		err := b.DecodeMessage(msg)
		m.BodyTransformation = &TransformationTemplate_MergeExtractorsToBody{msg}
		return true, err
	default:
		return false, nil
	}
}

func _TransformationTemplate_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*TransformationTemplate)
	// body_transformation
	switch x := m.BodyTransformation.(type) {
	case *TransformationTemplate_Body:
		s := proto.Size(x.Body)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransformationTemplate_Passthrough:
		s := proto.Size(x.Passthrough)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *TransformationTemplate_MergeExtractorsToBody:
		s := proto.Size(x.MergeExtractorsToBody)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

//
//custom functions:
//header_value(name) -> from the original headers
//extracted_value(name, index) -> from the extracted values
type InjaTemplate struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InjaTemplate) Reset()         { *m = InjaTemplate{} }
func (m *InjaTemplate) String() string { return proto.CompactTextString(m) }
func (*InjaTemplate) ProtoMessage()    {}
func (*InjaTemplate) Descriptor() ([]byte, []int) {
	return fileDescriptor_201f67ff59830de4, []int{4}
}
func (m *InjaTemplate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InjaTemplate.Unmarshal(m, b)
}
func (m *InjaTemplate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InjaTemplate.Marshal(b, m, deterministic)
}
func (m *InjaTemplate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InjaTemplate.Merge(m, src)
}
func (m *InjaTemplate) XXX_Size() int {
	return xxx_messageInfo_InjaTemplate.Size(m)
}
func (m *InjaTemplate) XXX_DiscardUnknown() {
	xxx_messageInfo_InjaTemplate.DiscardUnknown(m)
}

var xxx_messageInfo_InjaTemplate proto.InternalMessageInfo

func (m *InjaTemplate) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Passthrough struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Passthrough) Reset()         { *m = Passthrough{} }
func (m *Passthrough) String() string { return proto.CompactTextString(m) }
func (*Passthrough) ProtoMessage()    {}
func (*Passthrough) Descriptor() ([]byte, []int) {
	return fileDescriptor_201f67ff59830de4, []int{5}
}
func (m *Passthrough) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Passthrough.Unmarshal(m, b)
}
func (m *Passthrough) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Passthrough.Marshal(b, m, deterministic)
}
func (m *Passthrough) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Passthrough.Merge(m, src)
}
func (m *Passthrough) XXX_Size() int {
	return xxx_messageInfo_Passthrough.Size(m)
}
func (m *Passthrough) XXX_DiscardUnknown() {
	xxx_messageInfo_Passthrough.DiscardUnknown(m)
}

var xxx_messageInfo_Passthrough proto.InternalMessageInfo

type MergeExtractorsToBody struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MergeExtractorsToBody) Reset()         { *m = MergeExtractorsToBody{} }
func (m *MergeExtractorsToBody) String() string { return proto.CompactTextString(m) }
func (*MergeExtractorsToBody) ProtoMessage()    {}
func (*MergeExtractorsToBody) Descriptor() ([]byte, []int) {
	return fileDescriptor_201f67ff59830de4, []int{6}
}
func (m *MergeExtractorsToBody) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MergeExtractorsToBody.Unmarshal(m, b)
}
func (m *MergeExtractorsToBody) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MergeExtractorsToBody.Marshal(b, m, deterministic)
}
func (m *MergeExtractorsToBody) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MergeExtractorsToBody.Merge(m, src)
}
func (m *MergeExtractorsToBody) XXX_Size() int {
	return xxx_messageInfo_MergeExtractorsToBody.Size(m)
}
func (m *MergeExtractorsToBody) XXX_DiscardUnknown() {
	xxx_messageInfo_MergeExtractorsToBody.DiscardUnknown(m)
}

var xxx_messageInfo_MergeExtractorsToBody proto.InternalMessageInfo

type HeaderBodyTransform struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderBodyTransform) Reset()         { *m = HeaderBodyTransform{} }
func (m *HeaderBodyTransform) String() string { return proto.CompactTextString(m) }
func (*HeaderBodyTransform) ProtoMessage()    {}
func (*HeaderBodyTransform) Descriptor() ([]byte, []int) {
	return fileDescriptor_201f67ff59830de4, []int{7}
}
func (m *HeaderBodyTransform) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderBodyTransform.Unmarshal(m, b)
}
func (m *HeaderBodyTransform) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderBodyTransform.Marshal(b, m, deterministic)
}
func (m *HeaderBodyTransform) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderBodyTransform.Merge(m, src)
}
func (m *HeaderBodyTransform) XXX_Size() int {
	return xxx_messageInfo_HeaderBodyTransform.Size(m)
}
func (m *HeaderBodyTransform) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderBodyTransform.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderBodyTransform proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RouteTransformations)(nil), "transformation.plugins.gloo.solo.io.RouteTransformations")
	proto.RegisterType((*Transformation)(nil), "transformation.plugins.gloo.solo.io.Transformation")
	proto.RegisterType((*Extraction)(nil), "transformation.plugins.gloo.solo.io.Extraction")
	proto.RegisterType((*TransformationTemplate)(nil), "transformation.plugins.gloo.solo.io.TransformationTemplate")
	proto.RegisterMapType((map[string]*Extraction)(nil), "transformation.plugins.gloo.solo.io.TransformationTemplate.ExtractorsEntry")
	proto.RegisterMapType((map[string]*InjaTemplate)(nil), "transformation.plugins.gloo.solo.io.TransformationTemplate.HeadersEntry")
	proto.RegisterType((*InjaTemplate)(nil), "transformation.plugins.gloo.solo.io.InjaTemplate")
	proto.RegisterType((*Passthrough)(nil), "transformation.plugins.gloo.solo.io.Passthrough")
	proto.RegisterType((*MergeExtractorsToBody)(nil), "transformation.plugins.gloo.solo.io.MergeExtractorsToBody")
	proto.RegisterType((*HeaderBodyTransform)(nil), "transformation.plugins.gloo.solo.io.HeaderBodyTransform")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/plugins/transformation/transformation.proto", fileDescriptor_201f67ff59830de4)
}

var fileDescriptor_201f67ff59830de4 = []byte{
	// 605 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x95, 0xdd, 0x6e, 0xda, 0x30,
	0x14, 0xc7, 0x09, 0xb4, 0xac, 0x3d, 0xb4, 0xfb, 0x70, 0x1b, 0x88, 0xb8, 0x98, 0x50, 0x76, 0xc3,
	0x4d, 0x93, 0xb5, 0xbd, 0xa9, 0xba, 0x3b, 0x24, 0x44, 0xa6, 0xa9, 0xd2, 0x64, 0xa1, 0x69, 0xda,
	0x0d, 0x0a, 0xe0, 0x85, 0x40, 0x12, 0x67, 0xb6, 0x83, 0xe0, 0x29, 0xf6, 0x12, 0xbb, 0xd8, 0x73,
	0xed, 0x09, 0xf6, 0x08, 0x53, 0x9c, 0x00, 0x49, 0x16, 0x4d, 0x59, 0x7b, 0xe7, 0x63, 0xe7, 0xfc,
	0x7f, 0xe7, 0xcb, 0x31, 0x7c, 0x76, 0x5c, 0xb1, 0x88, 0xa6, 0xc6, 0x8c, 0xfa, 0x26, 0xa7, 0x1e,
	0xbd, 0x72, 0xa9, 0xe9, 0x78, 0x94, 0x9a, 0x21, 0xa3, 0x4b, 0x32, 0x13, 0x3c, 0xb1, 0xec, 0xd0,
	0x35, 0xd7, 0xd7, 0x66, 0xe8, 0x45, 0x8e, 0x1b, 0x70, 0x53, 0x30, 0x3b, 0xe0, 0x5f, 0x29, 0xf3,
	0x6d, 0xe1, 0xd2, 0xa0, 0x60, 0x1a, 0x21, 0xa3, 0x82, 0xa2, 0x37, 0xc5, 0xdd, 0xc4, 0xd7, 0x88,
	0xf5, 0x8c, 0x18, 0x65, 0xb8, 0xb4, 0x7b, 0xe9, 0x50, 0x87, 0xca, 0xef, 0xcd, 0x78, 0x95, 0xb8,
	0xea, 0xbf, 0x15, 0xb8, 0xc4, 0x34, 0x12, 0x64, 0x9c, 0x93, 0xe0, 0x68, 0x09, 0x6d, 0x46, 0xbe,
	0x45, 0x84, 0x8b, 0x49, 0x5e, 0x5d, 0x53, 0x7a, 0x4a, 0xbf, 0x75, 0x73, 0x6b, 0x54, 0x80, 0x1a,
	0x79, 0x55, 0xac, 0xa6, 0x92, 0xf9, 0x6d, 0xe4, 0x41, 0x87, 0x11, 0x1e, 0xd2, 0x80, 0x93, 0x22,
	0xac, 0xfe, 0x78, 0x58, 0x7b, 0xa7, 0x99, 0xdf, 0xd7, 0xbf, 0xd7, 0xe1, 0x79, 0x21, 0x80, 0x35,
	0x74, 0xf2, 0x80, 0x89, 0x20, 0x7e, 0xe8, 0xd9, 0x82, 0xa4, 0xd9, 0xbe, 0x7b, 0x44, 0x00, 0xe3,
	0x54, 0xc2, 0xaa, 0xe1, 0xb6, 0x28, 0x3d, 0x41, 0x01, 0xa8, 0x0b, 0x62, 0xcf, 0x09, 0x9b, 0x4c,
	0xe9, 0x7c, 0x7b, 0xc8, 0x3d, 0x4d, 0xfb, 0xae, 0x12, 0xd5, 0x92, 0x0a, 0x03, 0x3a, 0xdf, 0xee,
	0xf9, 0x56, 0x0d, 0x5f, 0x2c, 0xfe, 0xde, 0x1e, 0xa8, 0x70, 0x51, 0xcc, 0x73, 0x1b, 0x12, 0xfd,
	0x13, 0xc0, 0x70, 0x23, 0x98, 0x3d, 0x93, 0xc5, 0x68, 0x43, 0x33, 0xf1, 0x95, 0xb9, 0x9f, 0xe2,
	0xd4, 0x42, 0x97, 0x70, 0xcc, 0x88, 0x43, 0x36, 0x32, 0xb8, 0x53, 0x9c, 0x18, 0xa8, 0x0b, 0x27,
	0x3c, 0x9a, 0x3a, 0x8c, 0x46, 0xa1, 0xd6, 0xe8, 0x29, 0xfd, 0x73, 0xbc, 0xb7, 0xf5, 0x1f, 0x4d,
	0x68, 0x97, 0xd7, 0x04, 0x5d, 0x01, 0xb2, 0xe7, 0x6b, 0x3b, 0x98, 0x91, 0xf9, 0xbe, 0xd6, 0x5c,
	0x02, 0x4f, 0xf0, 0xab, 0xdd, 0xc9, 0xee, 0x6b, 0x8e, 0x56, 0x00, 0x24, 0x89, 0x90, 0x32, 0xae,
	0xd5, 0x7b, 0x8d, 0x7e, 0xeb, 0xe6, 0xc3, 0x13, 0x7a, 0x62, 0x0c, 0xf7, 0x6a, 0xc3, 0x40, 0xb0,
	0x2d, 0xce, 0xc8, 0xa3, 0x29, 0x3c, 0x4b, 0x52, 0xe6, 0x5a, 0x43, 0x92, 0xac, 0xa7, 0x90, 0x92,
	0xf6, 0xa4, 0x98, 0x9d, 0x30, 0x1a, 0xc1, 0x51, 0xdc, 0x72, 0xed, 0x48, 0x36, 0xfa, 0xba, 0x12,
	0xe0, 0x7d, 0xb0, 0xb4, 0x33, 0x43, 0x25, 0x05, 0xd0, 0x18, 0x5a, 0xa1, 0xcd, 0xb9, 0x58, 0x30,
	0x1a, 0x39, 0x0b, 0xed, 0x58, 0xea, 0xbd, 0xad, 0xa4, 0xf7, 0xf1, 0xe0, 0x67, 0xd5, 0x70, 0x56,
	0x06, 0x45, 0xa0, 0xf9, 0x84, 0x39, 0x64, 0x72, 0x28, 0xcb, 0x44, 0x50, 0x39, 0xa5, 0x5a, 0x53,
	0x22, 0xee, 0x2b, 0x21, 0x1e, 0x62, 0x91, 0x43, 0xad, 0xc7, 0x34, 0x9e, 0x47, 0xab, 0x86, 0x55,
	0xbf, 0xec, 0xa0, 0x1b, 0xc0, 0x8b, 0x42, 0x63, 0xd0, 0x4b, 0x68, 0xac, 0xc8, 0x36, 0x1d, 0xc5,
	0x78, 0x89, 0x86, 0x70, 0xbc, 0xb6, 0xbd, 0x88, 0xa4, 0x97, 0xc4, 0xac, 0x14, 0xc8, 0x61, 0xbe,
	0x71, 0xe2, 0x7d, 0x5f, 0xbf, 0x53, 0xba, 0x3e, 0x9c, 0x65, 0xdb, 0x53, 0x02, 0x1b, 0xe5, 0x61,
	0xff, 0xdf, 0xa8, 0x0c, 0x2e, 0xbe, 0x7e, 0xf9, 0x7b, 0x9e, 0xfc, 0x90, 0x74, 0x38, 0xcb, 0x7a,
	0x20, 0x04, 0x47, 0x82, 0x6c, 0x44, 0x1a, 0x86, 0x5c, 0xeb, 0xe7, 0xd0, 0xca, 0xb4, 0x4b, 0xef,
	0x80, 0x5a, 0x5a, 0x5a, 0x5d, 0x85, 0x8b, 0x92, 0xff, 0xc1, 0xe0, 0xe1, 0xe7, 0xaf, 0xd7, 0xca,
	0x97, 0x51, 0xb5, 0x17, 0x28, 0x5c, 0x39, 0xff, 0x7e, 0x85, 0xa6, 0x4d, 0xf9, 0x78, 0xdc, 0xfe,
	0x09, 0x00, 0x00, 0xff, 0xff, 0x4f, 0x65, 0x99, 0x3c, 0xd3, 0x06, 0x00, 0x00,
}

func (this *RouteTransformations) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RouteTransformations)
	if !ok {
		that2, ok := that.(RouteTransformations)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.RequestTransformation.Equal(that1.RequestTransformation) {
		return false
	}
	if !this.ResponseTransformation.Equal(that1.ResponseTransformation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Transformation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transformation)
	if !ok {
		that2, ok := that.(Transformation)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.TransformationType == nil {
		if this.TransformationType != nil {
			return false
		}
	} else if this.TransformationType == nil {
		return false
	} else if !this.TransformationType.Equal(that1.TransformationType) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Transformation_TransformationTemplate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transformation_TransformationTemplate)
	if !ok {
		that2, ok := that.(Transformation_TransformationTemplate)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.TransformationTemplate.Equal(that1.TransformationTemplate) {
		return false
	}
	return true
}
func (this *Transformation_HeaderBodyTransform) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Transformation_HeaderBodyTransform)
	if !ok {
		that2, ok := that.(Transformation_HeaderBodyTransform)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.HeaderBodyTransform.Equal(that1.HeaderBodyTransform) {
		return false
	}
	return true
}
func (this *Extraction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Extraction)
	if !ok {
		that2, ok := that.(Extraction)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Header != that1.Header {
		return false
	}
	if this.Regex != that1.Regex {
		return false
	}
	if this.Subgroup != that1.Subgroup {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TransformationTemplate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TransformationTemplate)
	if !ok {
		that2, ok := that.(TransformationTemplate)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.AdvancedTemplates != that1.AdvancedTemplates {
		return false
	}
	if len(this.Extractors) != len(that1.Extractors) {
		return false
	}
	for i := range this.Extractors {
		if !this.Extractors[i].Equal(that1.Extractors[i]) {
			return false
		}
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if !this.Headers[i].Equal(that1.Headers[i]) {
			return false
		}
	}
	if that1.BodyTransformation == nil {
		if this.BodyTransformation != nil {
			return false
		}
	} else if this.BodyTransformation == nil {
		return false
	} else if !this.BodyTransformation.Equal(that1.BodyTransformation) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TransformationTemplate_Body) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TransformationTemplate_Body)
	if !ok {
		that2, ok := that.(TransformationTemplate_Body)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Body.Equal(that1.Body) {
		return false
	}
	return true
}
func (this *TransformationTemplate_Passthrough) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TransformationTemplate_Passthrough)
	if !ok {
		that2, ok := that.(TransformationTemplate_Passthrough)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Passthrough.Equal(that1.Passthrough) {
		return false
	}
	return true
}
func (this *TransformationTemplate_MergeExtractorsToBody) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TransformationTemplate_MergeExtractorsToBody)
	if !ok {
		that2, ok := that.(TransformationTemplate_MergeExtractorsToBody)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.MergeExtractorsToBody.Equal(that1.MergeExtractorsToBody) {
		return false
	}
	return true
}
func (this *InjaTemplate) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*InjaTemplate)
	if !ok {
		that2, ok := that.(InjaTemplate)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Text != that1.Text {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Passthrough) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Passthrough)
	if !ok {
		that2, ok := that.(Passthrough)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MergeExtractorsToBody) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MergeExtractorsToBody)
	if !ok {
		that2, ok := that.(MergeExtractorsToBody)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *HeaderBodyTransform) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderBodyTransform)
	if !ok {
		that2, ok := that.(HeaderBodyTransform)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}