// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/description_entity.proto

package admin

import (
	fmt "fmt"
	math "math"

	core "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
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

// The format of the long description
type DescriptionFormat int32

const (
	DescriptionFormat_DESCRIPTION_FORMAT_UNKNOWN  DescriptionFormat = 0
	DescriptionFormat_DESCRIPTION_FORMAT_MARKDOWN DescriptionFormat = 1
	DescriptionFormat_DESCRIPTION_FORMAT_HTML     DescriptionFormat = 2
	// python default documentation - comments is rst
	DescriptionFormat_DESCRIPTION_FORMAT_RST DescriptionFormat = 3
)

var DescriptionFormat_name = map[int32]string{
	0: "DESCRIPTION_FORMAT_UNKNOWN",
	1: "DESCRIPTION_FORMAT_MARKDOWN",
	2: "DESCRIPTION_FORMAT_HTML",
	3: "DESCRIPTION_FORMAT_RST",
}

var DescriptionFormat_value = map[string]int32{
	"DESCRIPTION_FORMAT_UNKNOWN":  0,
	"DESCRIPTION_FORMAT_MARKDOWN": 1,
	"DESCRIPTION_FORMAT_HTML":     2,
	"DESCRIPTION_FORMAT_RST":      3,
}

func (x DescriptionFormat) String() string {
	return proto.EnumName(DescriptionFormat_name, int32(x))
}

func (DescriptionFormat) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2715f55631fe48ea, []int{0}
}

// DescriptionEntity contains detailed description for the task/workflow.
// Documentation could provide insight into the algorithms, business use case, etc.
type DescriptionEntity struct {
	// id represents the unique identifier of the description entity.
	Id *core.Identifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// One-liner overview of the entity.
	ShortDescription string `protobuf:"bytes,2,opt,name=short_description,json=shortDescription,proto3" json:"short_description,omitempty"`
	// Full user description with formatting preserved.
	LongDescription *Description `protobuf:"bytes,3,opt,name=long_description,json=longDescription,proto3" json:"long_description,omitempty"`
	// Optional link to source code used to define this entity.
	SourceCode *SourceCode `protobuf:"bytes,4,opt,name=source_code,json=sourceCode,proto3" json:"source_code,omitempty"`
	// User-specified tags. These are arbitrary and can be used for searching
	// filtering and discovering tasks.
	Tags                 []string `protobuf:"bytes,5,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescriptionEntity) Reset()         { *m = DescriptionEntity{} }
func (m *DescriptionEntity) String() string { return proto.CompactTextString(m) }
func (*DescriptionEntity) ProtoMessage()    {}
func (*DescriptionEntity) Descriptor() ([]byte, []int) {
	return fileDescriptor_2715f55631fe48ea, []int{0}
}

func (m *DescriptionEntity) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescriptionEntity.Unmarshal(m, b)
}
func (m *DescriptionEntity) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescriptionEntity.Marshal(b, m, deterministic)
}
func (m *DescriptionEntity) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescriptionEntity.Merge(m, src)
}
func (m *DescriptionEntity) XXX_Size() int {
	return xxx_messageInfo_DescriptionEntity.Size(m)
}
func (m *DescriptionEntity) XXX_DiscardUnknown() {
	xxx_messageInfo_DescriptionEntity.DiscardUnknown(m)
}

var xxx_messageInfo_DescriptionEntity proto.InternalMessageInfo

func (m *DescriptionEntity) GetId() *core.Identifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DescriptionEntity) GetShortDescription() string {
	if m != nil {
		return m.ShortDescription
	}
	return ""
}

func (m *DescriptionEntity) GetLongDescription() *Description {
	if m != nil {
		return m.LongDescription
	}
	return nil
}

func (m *DescriptionEntity) GetSourceCode() *SourceCode {
	if m != nil {
		return m.SourceCode
	}
	return nil
}

func (m *DescriptionEntity) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

// Full user description with formatting preserved. This can be rendered
// by clients, such as the console or command line tools with in-tact
// formatting.
type Description struct {
	// Types that are valid to be assigned to Content:
	//	*Description_Value
	//	*Description_Uri
	Content isDescription_Content `protobuf_oneof:"content"`
	// Format of the long description
	Format DescriptionFormat `protobuf:"varint,3,opt,name=format,proto3,enum=flyteidl.admin.DescriptionFormat" json:"format,omitempty"`
	// Optional link to an icon for the entity
	IconLink             string   `protobuf:"bytes,4,opt,name=icon_link,json=iconLink,proto3" json:"icon_link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Description) Reset()         { *m = Description{} }
func (m *Description) String() string { return proto.CompactTextString(m) }
func (*Description) ProtoMessage()    {}
func (*Description) Descriptor() ([]byte, []int) {
	return fileDescriptor_2715f55631fe48ea, []int{1}
}

func (m *Description) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Description.Unmarshal(m, b)
}
func (m *Description) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Description.Marshal(b, m, deterministic)
}
func (m *Description) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Description.Merge(m, src)
}
func (m *Description) XXX_Size() int {
	return xxx_messageInfo_Description.Size(m)
}
func (m *Description) XXX_DiscardUnknown() {
	xxx_messageInfo_Description.DiscardUnknown(m)
}

var xxx_messageInfo_Description proto.InternalMessageInfo

type isDescription_Content interface {
	isDescription_Content()
}

type Description_Value struct {
	Value string `protobuf:"bytes,1,opt,name=value,proto3,oneof"`
}

type Description_Uri struct {
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof"`
}

func (*Description_Value) isDescription_Content() {}

func (*Description_Uri) isDescription_Content() {}

func (m *Description) GetContent() isDescription_Content {
	if m != nil {
		return m.Content
	}
	return nil
}

func (m *Description) GetValue() string {
	if x, ok := m.GetContent().(*Description_Value); ok {
		return x.Value
	}
	return ""
}

func (m *Description) GetUri() string {
	if x, ok := m.GetContent().(*Description_Uri); ok {
		return x.Uri
	}
	return ""
}

func (m *Description) GetFormat() DescriptionFormat {
	if m != nil {
		return m.Format
	}
	return DescriptionFormat_DESCRIPTION_FORMAT_UNKNOWN
}

func (m *Description) GetIconLink() string {
	if m != nil {
		return m.IconLink
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Description) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Description_Value)(nil),
		(*Description_Uri)(nil),
	}
}

// Link to source code used to define this entity
type SourceCode struct {
	Link                 string   `protobuf:"bytes,1,opt,name=link,proto3" json:"link,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceCode) Reset()         { *m = SourceCode{} }
func (m *SourceCode) String() string { return proto.CompactTextString(m) }
func (*SourceCode) ProtoMessage()    {}
func (*SourceCode) Descriptor() ([]byte, []int) {
	return fileDescriptor_2715f55631fe48ea, []int{2}
}

func (m *SourceCode) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceCode.Unmarshal(m, b)
}
func (m *SourceCode) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceCode.Marshal(b, m, deterministic)
}
func (m *SourceCode) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceCode.Merge(m, src)
}
func (m *SourceCode) XXX_Size() int {
	return xxx_messageInfo_SourceCode.Size(m)
}
func (m *SourceCode) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceCode.DiscardUnknown(m)
}

var xxx_messageInfo_SourceCode proto.InternalMessageInfo

func (m *SourceCode) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

// Represents a list of DescriptionEntities returned from the admin.
// See :ref:`ref_flyteidl.admin.DescriptionEntity` for more details
type DescriptionEntityList struct {
	// A list of DescriptionEntities returned based on the request.
	DescriptionEntities []*DescriptionEntity `protobuf:"bytes,1,rep,name=descriptionEntities,proto3" json:"descriptionEntities,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query. If there are no more results, this value will be empty.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescriptionEntityList) Reset()         { *m = DescriptionEntityList{} }
func (m *DescriptionEntityList) String() string { return proto.CompactTextString(m) }
func (*DescriptionEntityList) ProtoMessage()    {}
func (*DescriptionEntityList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2715f55631fe48ea, []int{3}
}

func (m *DescriptionEntityList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescriptionEntityList.Unmarshal(m, b)
}
func (m *DescriptionEntityList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescriptionEntityList.Marshal(b, m, deterministic)
}
func (m *DescriptionEntityList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescriptionEntityList.Merge(m, src)
}
func (m *DescriptionEntityList) XXX_Size() int {
	return xxx_messageInfo_DescriptionEntityList.Size(m)
}
func (m *DescriptionEntityList) XXX_DiscardUnknown() {
	xxx_messageInfo_DescriptionEntityList.DiscardUnknown(m)
}

var xxx_messageInfo_DescriptionEntityList proto.InternalMessageInfo

func (m *DescriptionEntityList) GetDescriptionEntities() []*DescriptionEntity {
	if m != nil {
		return m.DescriptionEntities
	}
	return nil
}

func (m *DescriptionEntityList) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Represents a request structure to retrieve a list of DescriptionEntities.
// See :ref:`ref_flyteidl.admin.DescriptionEntity` for more details
type DescriptionEntityListRequest struct {
	// Identifies the specific type of resource that this identifier corresponds to.
	ResourceType core.ResourceType `protobuf:"varint,1,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.core.ResourceType" json:"resource_type,omitempty"`
	// The identifier for the description entity.
	// +required
	Id *NamedEntityIdentifier `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Indicates the number of resources to be returned.
	// +required
	Limit uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query.
	// +optional
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	// Indicates a list of filters passed as string.
	// More info on constructing filters : <Link>
	// +optional
	Filters string `protobuf:"bytes,5,opt,name=filters,proto3" json:"filters,omitempty"`
	// Sort ordering for returned list.
	// +optional
	SortBy               *Sort    `protobuf:"bytes,6,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DescriptionEntityListRequest) Reset()         { *m = DescriptionEntityListRequest{} }
func (m *DescriptionEntityListRequest) String() string { return proto.CompactTextString(m) }
func (*DescriptionEntityListRequest) ProtoMessage()    {}
func (*DescriptionEntityListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2715f55631fe48ea, []int{4}
}

func (m *DescriptionEntityListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DescriptionEntityListRequest.Unmarshal(m, b)
}
func (m *DescriptionEntityListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DescriptionEntityListRequest.Marshal(b, m, deterministic)
}
func (m *DescriptionEntityListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DescriptionEntityListRequest.Merge(m, src)
}
func (m *DescriptionEntityListRequest) XXX_Size() int {
	return xxx_messageInfo_DescriptionEntityListRequest.Size(m)
}
func (m *DescriptionEntityListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DescriptionEntityListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DescriptionEntityListRequest proto.InternalMessageInfo

func (m *DescriptionEntityListRequest) GetResourceType() core.ResourceType {
	if m != nil {
		return m.ResourceType
	}
	return core.ResourceType_UNSPECIFIED
}

func (m *DescriptionEntityListRequest) GetId() *NamedEntityIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *DescriptionEntityListRequest) GetLimit() uint32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *DescriptionEntityListRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DescriptionEntityListRequest) GetFilters() string {
	if m != nil {
		return m.Filters
	}
	return ""
}

func (m *DescriptionEntityListRequest) GetSortBy() *Sort {
	if m != nil {
		return m.SortBy
	}
	return nil
}

func init() {
	proto.RegisterEnum("flyteidl.admin.DescriptionFormat", DescriptionFormat_name, DescriptionFormat_value)
	proto.RegisterType((*DescriptionEntity)(nil), "flyteidl.admin.DescriptionEntity")
	proto.RegisterType((*Description)(nil), "flyteidl.admin.Description")
	proto.RegisterType((*SourceCode)(nil), "flyteidl.admin.SourceCode")
	proto.RegisterType((*DescriptionEntityList)(nil), "flyteidl.admin.DescriptionEntityList")
	proto.RegisterType((*DescriptionEntityListRequest)(nil), "flyteidl.admin.DescriptionEntityListRequest")
}

func init() {
	proto.RegisterFile("flyteidl/admin/description_entity.proto", fileDescriptor_2715f55631fe48ea)
}

var fileDescriptor_2715f55631fe48ea = []byte{
	// 600 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x5d, 0xd2, 0xb5, 0x23, 0xb7, 0x6c, 0x74, 0x66, 0x8c, 0xd0, 0xa2, 0x51, 0x2a, 0x21, 0x0a,
	0x68, 0x89, 0x34, 0xc4, 0x03, 0x9a, 0x90, 0xd8, 0x57, 0xd5, 0x6a, 0xfd, 0x40, 0x6e, 0x11, 0x12,
	0x2f, 0x51, 0x9b, 0xb8, 0x99, 0xd5, 0x24, 0x2e, 0x8e, 0x8b, 0x94, 0x57, 0xc4, 0x23, 0xfc, 0x03,
	0x7e, 0x2c, 0x8a, 0xdd, 0x76, 0x69, 0x57, 0xed, 0xcd, 0x37, 0xf7, 0xdc, 0x73, 0xef, 0xf1, 0xb9,
	0x31, 0xbc, 0x1e, 0x07, 0x89, 0x20, 0xd4, 0x0b, 0xec, 0xa1, 0x17, 0xd2, 0xc8, 0xf6, 0x48, 0xec,
	0x72, 0x3a, 0x15, 0x94, 0x45, 0x0e, 0x89, 0x04, 0x15, 0x89, 0x35, 0xe5, 0x4c, 0x30, 0xb4, 0xb7,
	0x00, 0x5a, 0x12, 0x58, 0x3e, 0x5a, 0x16, 0xba, 0x8c, 0x13, 0x9b, 0x7a, 0x29, 0x7a, 0x4c, 0x09,
	0x57, 0xf8, 0x72, 0x65, 0x8d, 0xd8, 0x65, 0x61, 0xc8, 0x22, 0x95, 0xac, 0xfd, 0xd6, 0x61, 0xff,
	0xf2, 0xb6, 0xd3, 0x95, 0x6c, 0x84, 0xde, 0x80, 0x4e, 0x3d, 0x53, 0xab, 0x6a, 0xf5, 0xe2, 0xc9,
	0x33, 0x6b, 0xd9, 0x2f, 0xe5, 0xb7, 0x5a, 0x4b, 0x7e, 0xac, 0x53, 0x0f, 0xbd, 0x83, 0xfd, 0xf8,
	0x86, 0x71, 0xe1, 0x64, 0xe6, 0x35, 0xf5, 0xaa, 0x56, 0x37, 0x70, 0x49, 0x26, 0x32, 0xec, 0xa8,
	0x01, 0xa5, 0x80, 0x45, 0xfe, 0x0a, 0x36, 0x27, 0xbb, 0x54, 0xac, 0x55, 0x55, 0x56, 0xa6, 0x0c,
	0x3f, 0x4a, 0x8b, 0xb2, 0x3c, 0xa7, 0x50, 0x8c, 0xd9, 0x8c, 0xbb, 0xc4, 0x71, 0x99, 0x47, 0xcc,
	0x6d, 0x49, 0x51, 0x5e, 0xa7, 0xe8, 0x4b, 0xc8, 0x05, 0xf3, 0x08, 0x86, 0x78, 0x79, 0x46, 0x08,
	0xb6, 0xc5, 0xd0, 0x8f, 0xcd, 0x7c, 0x35, 0x57, 0x37, 0xb0, 0x3c, 0xd7, 0xfe, 0x69, 0x50, 0xcc,
	0x36, 0x38, 0x84, 0xfc, 0xcf, 0x61, 0x30, 0x23, 0xf2, 0x0e, 0x8c, 0xe6, 0x16, 0x56, 0x21, 0x42,
	0x90, 0x9b, 0x71, 0xaa, 0xf4, 0x35, 0xb7, 0x70, 0x1a, 0xa0, 0x8f, 0x50, 0x18, 0x33, 0x1e, 0x0e,
	0x85, 0x94, 0xb2, 0x77, 0xf2, 0xf2, 0x1e, 0x29, 0x0d, 0x09, 0xc4, 0xf3, 0x02, 0x54, 0x01, 0x83,
	0xba, 0x2c, 0x72, 0x02, 0x1a, 0x4d, 0xa4, 0x0a, 0x03, 0x3f, 0x48, 0x3f, 0xb4, 0x69, 0x34, 0x39,
	0x37, 0x60, 0xc7, 0x65, 0x91, 0x20, 0x91, 0xa8, 0x55, 0x01, 0xfa, 0x2b, 0x02, 0x64, 0x81, 0x9c,
	0x0d, 0xcb, 0x73, 0xed, 0x97, 0x06, 0x4f, 0xee, 0xf8, 0xd8, 0xa6, 0xb1, 0x40, 0x7d, 0x78, 0xec,
	0xad, 0x25, 0x28, 0x89, 0x4d, 0xad, 0x9a, 0xab, 0x17, 0xef, 0x9d, 0x55, 0x71, 0xe0, 0x4d, 0xd5,
	0xe8, 0x00, 0xf2, 0x82, 0x4d, 0xc8, 0xc2, 0x69, 0x15, 0xd4, 0xfe, 0xe8, 0xf0, 0x7c, 0xe3, 0x10,
	0x98, 0xfc, 0x98, 0x91, 0x58, 0xa0, 0xcf, 0xb0, 0xcb, 0xc9, 0xdc, 0x39, 0x91, 0x4c, 0xd5, 0xf5,
	0xee, 0x65, 0xcd, 0x97, 0x2b, 0x86, 0xe7, 0x98, 0x41, 0x32, 0x25, 0xf8, 0x21, 0xcf, 0x44, 0xe8,
	0x83, 0xdc, 0x4c, 0x5d, 0x1a, 0xfe, 0x6a, 0x7d, 0xf8, 0xee, 0x30, 0x24, 0x9e, 0xea, 0xba, 0xb6,
	0xa5, 0x07, 0x90, 0x0f, 0x68, 0x48, 0x95, 0x45, 0xbb, 0x58, 0x05, 0xb7, 0x2a, 0xb6, 0x33, 0x2a,
	0x90, 0x09, 0x3b, 0x63, 0x1a, 0x08, 0xc2, 0xd3, 0x15, 0x49, 0xbf, 0x2f, 0x42, 0x74, 0x0c, 0x3b,
	0x71, 0xba, 0xea, 0xa3, 0xc4, 0x2c, 0xc8, 0x09, 0x0e, 0xee, 0xae, 0x1c, 0x17, 0xb8, 0x90, 0x82,
	0xce, 0x93, 0xb7, 0x7f, 0xb5, 0x95, 0x7f, 0x4b, 0x79, 0x8f, 0x8e, 0xa0, 0x7c, 0x79, 0xd5, 0xbf,
	0xc0, 0xad, 0x2f, 0x83, 0x56, 0xaf, 0xeb, 0x34, 0x7a, 0xb8, 0x73, 0x36, 0x70, 0xbe, 0x76, 0xaf,
	0xbb, 0xbd, 0x6f, 0xdd, 0xd2, 0x16, 0x7a, 0x01, 0x95, 0x0d, 0xf9, 0xce, 0x19, 0xbe, 0xbe, 0x4c,
	0x01, 0x1a, 0xaa, 0xc0, 0xd3, 0x0d, 0x80, 0xe6, 0xa0, 0xd3, 0x2e, 0xe9, 0xa8, 0x0c, 0x87, 0x1b,
	0x92, 0xb8, 0x3f, 0x28, 0xe5, 0xce, 0x3f, 0x7d, 0x3f, 0xf5, 0xa9, 0xb8, 0x99, 0x8d, 0x2c, 0x97,
	0x85, 0xb6, 0x9c, 0x9c, 0x71, 0x5f, 0x1d, 0xec, 0xe5, 0x23, 0xe1, 0x93, 0xc8, 0x9e, 0x8e, 0x8e,
	0x7d, 0x66, 0xaf, 0xbe, 0x1b, 0xa3, 0x82, 0x7c, 0x31, 0xde, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0x01, 0x53, 0xea, 0xcd, 0xa9, 0x04, 0x00, 0x00,
}
