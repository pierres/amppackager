// Code generated by protoc-gen-go. DO NOT EDIT.
// source: transformer/request/request.proto

package request // import "github.com/ampproject/amppackager/transformer/request"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// This should be kept in sync with HtmlFormat.Code in
// github.com/ampproject/amphtml/validator/validator.proto.
type Request_HtmlFormat int32

const (
	Request_UNKNOWN_CODE Request_HtmlFormat = 0
	Request_AMP          Request_HtmlFormat = 1
	Request_AMP4ADS      Request_HtmlFormat = 2
	Request_AMP4EMAIL    Request_HtmlFormat = 3
	Request_EXPERIMENTAL Request_HtmlFormat = 4
)

var Request_HtmlFormat_name = map[int32]string{
	0: "UNKNOWN_CODE",
	1: "AMP",
	2: "AMP4ADS",
	3: "AMP4EMAIL",
	4: "EXPERIMENTAL",
}

var Request_HtmlFormat_value = map[string]int32{
	"UNKNOWN_CODE": 0,
	"AMP":          1,
	"AMP4ADS":      2,
	"AMP4EMAIL":    3,
	"EXPERIMENTAL": 4,
}

func (x Request_HtmlFormat) String() string {
	return proto.EnumName(Request_HtmlFormat_name, int32(x))
}

func (Request_HtmlFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_762cce2ac5f73405, []int{0, 0}
}

type Request_TransformersConfig int32

const (
	// Execute the default list of transformers. For packager production
	// environments, this should be the config used.
	Request_DEFAULT Request_TransformersConfig = 0
	// Execute none, and simply parse and re-emit. Some normalization will be
	// performed regardless, including, but not limited to:
	// - HTML normalization (e.g. closing all non-void tags).
	// - removal of all comments
	// - lowercase-ing of attribute keys
	// - lexical sort of attribute keys
	// - text is escaped
	//
	// WARNING. THIS IS FOR TESTING PURPOSES ONLY.
	// Use of this setting in a packager production environment could produce
	// invalid transformed AMP when ingested by AMP caches.
	Request_NONE Request_TransformersConfig = 1
	// Execute the minimum needed for verification/validation.
	//
	// WARNING. FOR AMP CACHE USE ONLY.
	// Use of this setting in a packager production environment could produce
	// invalid transformed AMP when ingested by AMP caches.
	Request_VALIDATION Request_TransformersConfig = 2
	// Execute a custom set of transformers.
	//
	// WARNING. THIS IS FOR TESTING PURPOSES ONLY.
	// Use of this setting in a packager production environment could produce
	// invalid transformed AMP when ingested by AMP caches.
	Request_CUSTOM Request_TransformersConfig = 3
)

var Request_TransformersConfig_name = map[int32]string{
	0: "DEFAULT",
	1: "NONE",
	2: "VALIDATION",
	3: "CUSTOM",
}

var Request_TransformersConfig_value = map[string]int32{
	"DEFAULT":    0,
	"NONE":       1,
	"VALIDATION": 2,
	"CUSTOM":     3,
}

func (x Request_TransformersConfig) String() string {
	return proto.EnumName(Request_TransformersConfig_name, int32(x))
}

func (Request_TransformersConfig) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_762cce2ac5f73405, []int{0, 1}
}

// A Request encapsulates input and contextual parameters for use by the
// transformers.
type Request struct {
	// The AMP HTML document to transform.
	Html string `protobuf:"bytes,1,opt,name=html,proto3" json:"html,omitempty"`
	// The public URL of the document, i.e. the location that should appear in
	// the browser URL bar.
	DocumentUrl string `protobuf:"bytes,2,opt,name=document_url,json=documentUrl,proto3" json:"document_url,omitempty"`
	// The AMP runtime version.
	Rtv string `protobuf:"bytes,4,opt,name=rtv,proto3" json:"rtv,omitempty"`
	// The CSS contents to inline into the transformed HTML
	Css string `protobuf:"bytes,5,opt,name=css,proto3" json:"css,omitempty"`
	// Transformations are only run if the HTML tag contains the attribute
	// specifying one of the provided formats. If allowed_formats is empty, then
	// all non-experimental AMP formats are allowed.
	AllowedFormats []Request_HtmlFormat       `protobuf:"varint,7,rep,packed,name=allowed_formats,json=allowedFormats,proto3,enum=amp.transform.Request_HtmlFormat" json:"allowed_formats,omitempty"`
	Config         Request_TransformersConfig `protobuf:"varint,6,opt,name=config,proto3,enum=amp.transform.Request_TransformersConfig" json:"config,omitempty"`
	// If config == CUSTOM, this is the list of custom transformers to execute,
	// in the order provided here. Otherwise, this is ignored.
	Transformers []string `protobuf:"bytes,3,rep,name=transformers,proto3" json:"transformers,omitempty"`
	// The set of all transform version numbers accepted by the requestor. These
	// ranges should be non-overlapping and in descending order.
	//
	// The transfomer will select the highest one from this list that it supports.
	// If this set is empty, the transfomer will select the highest version it
	// supports. In both cases, it shouldn't include in-development versions.
	Versions             []*Request_VersionRange `protobuf:"bytes,8,rep,name=versions,proto3" json:"versions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_762cce2ac5f73405, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetHtml() string {
	if m != nil {
		return m.Html
	}
	return ""
}

func (m *Request) GetDocumentUrl() string {
	if m != nil {
		return m.DocumentUrl
	}
	return ""
}

func (m *Request) GetRtv() string {
	if m != nil {
		return m.Rtv
	}
	return ""
}

func (m *Request) GetCss() string {
	if m != nil {
		return m.Css
	}
	return ""
}

func (m *Request) GetAllowedFormats() []Request_HtmlFormat {
	if m != nil {
		return m.AllowedFormats
	}
	return nil
}

func (m *Request) GetConfig() Request_TransformersConfig {
	if m != nil {
		return m.Config
	}
	return Request_DEFAULT
}

func (m *Request) GetTransformers() []string {
	if m != nil {
		return m.Transformers
	}
	return nil
}

func (m *Request) GetVersions() []*Request_VersionRange {
	if m != nil {
		return m.Versions
	}
	return nil
}

// An inclusive range of version numbers.
type Request_VersionRange struct {
	Min                  int64    `protobuf:"varint,1,opt,name=min,proto3" json:"min,omitempty"`
	Max                  int64    `protobuf:"varint,2,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request_VersionRange) Reset()         { *m = Request_VersionRange{} }
func (m *Request_VersionRange) String() string { return proto.CompactTextString(m) }
func (*Request_VersionRange) ProtoMessage()    {}
func (*Request_VersionRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_762cce2ac5f73405, []int{0, 0}
}
func (m *Request_VersionRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request_VersionRange.Unmarshal(m, b)
}
func (m *Request_VersionRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request_VersionRange.Marshal(b, m, deterministic)
}
func (dst *Request_VersionRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request_VersionRange.Merge(dst, src)
}
func (m *Request_VersionRange) XXX_Size() int {
	return xxx_messageInfo_Request_VersionRange.Size(m)
}
func (m *Request_VersionRange) XXX_DiscardUnknown() {
	xxx_messageInfo_Request_VersionRange.DiscardUnknown(m)
}

var xxx_messageInfo_Request_VersionRange proto.InternalMessageInfo

func (m *Request_VersionRange) GetMin() int64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *Request_VersionRange) GetMax() int64 {
	if m != nil {
		return m.Max
	}
	return 0
}

// A Metadata is part of the transformers' response, and includes additional
// information either not present in or not easily accessible from the HTML. It
// should remain relatively small, as it undergoes a
// serialization/deserialization round-trip when the Go library is called from
// C.
type Metadata struct {
	// Absolute URLs of resources that should be preloaded when the AMP is
	// prefetched. In a signed exchange (SXG) context, these would be included as
	// `Link: rel=preload` headers, as these are used by the browser during SXG
	// prefetch:
	// https://github.com/WICG/webpackage/blob/master/explainer.md#prefetching-stops-here
	Preloads []*Metadata_Preload `protobuf:"bytes,1,rep,name=preloads,proto3" json:"preloads,omitempty"`
	// The version of the transforms that were performed on this output.
	Version              int64    `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata) Reset()         { *m = Metadata{} }
func (m *Metadata) String() string { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()    {}
func (*Metadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_762cce2ac5f73405, []int{1}
}
func (m *Metadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata.Unmarshal(m, b)
}
func (m *Metadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata.Marshal(b, m, deterministic)
}
func (dst *Metadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata.Merge(dst, src)
}
func (m *Metadata) XXX_Size() int {
	return xxx_messageInfo_Metadata.Size(m)
}
func (m *Metadata) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata proto.InternalMessageInfo

func (m *Metadata) GetPreloads() []*Metadata_Preload {
	if m != nil {
		return m.Preloads
	}
	return nil
}

func (m *Metadata) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type Metadata_Preload struct {
	// The URL of the resource to preload. Will be an absolute URL on the domain
	// of the target AMP cache.
	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	// The `as` attribute of the preload, as specified in
	// https://w3c.github.io/preload/#as-attribute and
	// https://html.spec.whatwg.org/multipage/semantics.html#attr-link-as. The
	// full list of potential values is specified in
	// https://fetch.spec.whatwg.org/#concept-request-destination, though for
	// the time being only "script" and "style" are allowed.
	As                   string   `protobuf:"bytes,2,opt,name=as,proto3" json:"as,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Metadata_Preload) Reset()         { *m = Metadata_Preload{} }
func (m *Metadata_Preload) String() string { return proto.CompactTextString(m) }
func (*Metadata_Preload) ProtoMessage()    {}
func (*Metadata_Preload) Descriptor() ([]byte, []int) {
	return fileDescriptor_762cce2ac5f73405, []int{1, 0}
}
func (m *Metadata_Preload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metadata_Preload.Unmarshal(m, b)
}
func (m *Metadata_Preload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metadata_Preload.Marshal(b, m, deterministic)
}
func (dst *Metadata_Preload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metadata_Preload.Merge(dst, src)
}
func (m *Metadata_Preload) XXX_Size() int {
	return xxx_messageInfo_Metadata_Preload.Size(m)
}
func (m *Metadata_Preload) XXX_DiscardUnknown() {
	xxx_messageInfo_Metadata_Preload.DiscardUnknown(m)
}

var xxx_messageInfo_Metadata_Preload proto.InternalMessageInfo

func (m *Metadata_Preload) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Metadata_Preload) GetAs() string {
	if m != nil {
		return m.As
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "amp.transform.Request")
	proto.RegisterType((*Request_VersionRange)(nil), "amp.transform.Request.VersionRange")
	proto.RegisterType((*Metadata)(nil), "amp.transform.Metadata")
	proto.RegisterType((*Metadata_Preload)(nil), "amp.transform.Metadata.Preload")
	proto.RegisterEnum("amp.transform.Request_HtmlFormat", Request_HtmlFormat_name, Request_HtmlFormat_value)
	proto.RegisterEnum("amp.transform.Request_TransformersConfig", Request_TransformersConfig_name, Request_TransformersConfig_value)
}

func init() { proto.RegisterFile("transformer/request/request.proto", fileDescriptor_762cce2ac5f73405) }

var fileDescriptor_762cce2ac5f73405 = []byte{
	// 500 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x5d, 0x73, 0xd2, 0x40,
	0x14, 0x6d, 0x08, 0x12, 0xb8, 0x50, 0xdc, 0xb9, 0x4f, 0x99, 0xbe, 0x08, 0xf8, 0x82, 0xe3, 0x4c,
	0x98, 0x41, 0x1d, 0x1f, 0x7c, 0x70, 0x22, 0xa4, 0x8a, 0x92, 0xc0, 0xa4, 0x50, 0x1d, 0x5f, 0x98,
	0x6d, 0xd8, 0x52, 0x34, 0x5f, 0xee, 0x2e, 0xb5, 0xbf, 0xc2, 0x5f, 0xe2, 0x8f, 0x74, 0x76, 0x13,
	0x68, 0xab, 0xf6, 0x69, 0xcf, 0x3d, 0xf7, 0x9c, 0xfd, 0x38, 0x73, 0x17, 0xba, 0x92, 0xd3, 0x54,
	0x5c, 0x66, 0x3c, 0x61, 0x7c, 0xc0, 0xd9, 0x8f, 0x1d, 0x13, 0x72, 0xbf, 0x3a, 0x39, 0xcf, 0x64,
	0x86, 0xc7, 0x34, 0xc9, 0x9d, 0x83, 0xac, 0xf7, 0xbb, 0x0a, 0x56, 0x58, 0x08, 0x10, 0xa1, 0x7a,
	0x25, 0x93, 0xd8, 0x36, 0x3a, 0x46, 0xbf, 0x11, 0x6a, 0x8c, 0x5d, 0x68, 0xad, 0xb3, 0x68, 0x97,
	0xb0, 0x54, 0xae, 0x76, 0x3c, 0xb6, 0x2b, 0xba, 0xd7, 0xdc, 0x73, 0x4b, 0x1e, 0x23, 0x01, 0x93,
	0xcb, 0x6b, 0xbb, 0xaa, 0x3b, 0x0a, 0x2a, 0x26, 0x12, 0xc2, 0x7e, 0x54, 0x30, 0x91, 0x10, 0xf8,
	0x11, 0x1e, 0xd3, 0x38, 0xce, 0x7e, 0xb2, 0xf5, 0x4a, 0x1d, 0x4b, 0xa5, 0xb0, 0xad, 0x8e, 0xd9,
	0x6f, 0x0f, 0xbb, 0xce, 0xbd, 0xfb, 0x38, 0xe5, 0x5d, 0x9c, 0x0f, 0x32, 0x89, 0x4f, 0xb5, 0x32,
	0x6c, 0x97, 0xce, 0xa2, 0x14, 0xe8, 0x42, 0x2d, 0xca, 0xd2, 0xcb, 0xed, 0xc6, 0xae, 0x75, 0x8c,
	0x7e, 0x7b, 0xf8, 0xec, 0x81, 0x2d, 0x16, 0xb7, 0x59, 0x88, 0x91, 0x36, 0x84, 0xa5, 0x11, 0x7b,
	0xd0, 0xba, 0x93, 0x94, 0xb0, 0xcd, 0x8e, 0xd9, 0x6f, 0x84, 0xf7, 0x38, 0x7c, 0x0b, 0xf5, 0x6b,
	0xc6, 0xc5, 0x36, 0x4b, 0x85, 0x5d, 0xef, 0x98, 0xfd, 0xe6, 0xf0, 0xe9, 0x03, 0x07, 0x9d, 0x17,
	0xb2, 0x90, 0xa6, 0x1b, 0x16, 0x1e, 0x4c, 0x27, 0x43, 0x68, 0xdd, 0xed, 0xa8, 0x54, 0x92, 0x6d,
	0xaa, 0xd3, 0x35, 0x43, 0x05, 0x35, 0x43, 0x6f, 0x74, 0xa6, 0x8a, 0xa1, 0x37, 0xbd, 0x25, 0xc0,
	0xed, 0xcb, 0x91, 0x40, 0x6b, 0x19, 0x7c, 0x0a, 0x66, 0x9f, 0x83, 0xd5, 0x68, 0x36, 0xf6, 0xc8,
	0x11, 0x5a, 0x60, 0xba, 0xfe, 0x9c, 0x18, 0xd8, 0x04, 0xcb, 0xf5, 0xe7, 0x2f, 0xdd, 0xf1, 0x19,
	0xa9, 0xe0, 0x31, 0x34, 0x54, 0xe1, 0xf9, 0xee, 0x64, 0x4a, 0x4c, 0x65, 0xf3, 0xbe, 0xcc, 0xbd,
	0x70, 0xe2, 0x7b, 0xc1, 0xc2, 0x9d, 0x92, 0x6a, 0xef, 0x3d, 0xe0, 0xbf, 0x69, 0xa8, 0x3d, 0xc6,
	0xde, 0xa9, 0xbb, 0x9c, 0x2e, 0xc8, 0x11, 0xd6, 0xa1, 0x1a, 0xcc, 0x02, 0x8f, 0x18, 0xd8, 0x06,
	0x38, 0x77, 0xa7, 0x93, 0xb1, 0xbb, 0x98, 0xcc, 0x02, 0x52, 0x41, 0x80, 0xda, 0x68, 0x79, 0xb6,
	0x98, 0xf9, 0xc4, 0xec, 0xfd, 0x32, 0xa0, 0xee, 0x33, 0x49, 0xd7, 0x54, 0x52, 0x7c, 0x03, 0xf5,
	0x9c, 0xb3, 0x38, 0xa3, 0x6b, 0x61, 0x1b, 0x3a, 0xa1, 0x27, 0x7f, 0x25, 0xb4, 0x97, 0x3a, 0xf3,
	0x42, 0x17, 0x1e, 0x0c, 0x68, 0x83, 0x55, 0x26, 0x55, 0xbe, 0x7f, 0x5f, 0x9e, 0x3c, 0x07, 0xab,
	0x94, 0xab, 0x80, 0xd4, 0xd0, 0x15, 0x03, 0xa9, 0x20, 0xb6, 0xa1, 0x42, 0x45, 0x39, 0x85, 0x15,
	0x2a, 0xde, 0xbd, 0xfe, 0xfa, 0x6a, 0xb3, 0x95, 0x57, 0xbb, 0x0b, 0x27, 0xca, 0x92, 0x01, 0x4d,
	0xf2, 0x9c, 0x67, 0xdf, 0x58, 0x24, 0x35, 0xa4, 0xd1, 0x77, 0xba, 0x61, 0x7c, 0xf0, 0x9f, 0x5f,
	0x71, 0x51, 0xd3, 0xdf, 0xe1, 0xc5, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x92, 0xb6, 0xa0, 0x1e,
	0x33, 0x03, 0x00, 0x00,
}
