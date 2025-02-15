// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/project_domain_attributes.proto

package admin

import (
	fmt "fmt"
	math "math"

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

// Defines a set of custom matching attributes which defines resource defaults for a project and domain.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectDomainAttributes struct {
	// Unique project id for which this set of attributes will be applied.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id for which this set of attributes will be applied.
	Domain               string              `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	MatchingAttributes   *MatchingAttributes `protobuf:"bytes,3,opt,name=matching_attributes,json=matchingAttributes,proto3" json:"matching_attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ProjectDomainAttributes) Reset()         { *m = ProjectDomainAttributes{} }
func (m *ProjectDomainAttributes) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributes) ProtoMessage()    {}
func (*ProjectDomainAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{0}
}

func (m *ProjectDomainAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributes.Unmarshal(m, b)
}
func (m *ProjectDomainAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributes.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributes.Merge(m, src)
}
func (m *ProjectDomainAttributes) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributes.Size(m)
}
func (m *ProjectDomainAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributes proto.InternalMessageInfo

func (m *ProjectDomainAttributes) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectDomainAttributes) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ProjectDomainAttributes) GetMatchingAttributes() *MatchingAttributes {
	if m != nil {
		return m.MatchingAttributes
	}
	return nil
}

// Sets custom attributes for a project-domain combination.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectDomainAttributesUpdateRequest struct {
	// +required
	Attributes           *ProjectDomainAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ProjectDomainAttributesUpdateRequest) Reset()         { *m = ProjectDomainAttributesUpdateRequest{} }
func (m *ProjectDomainAttributesUpdateRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesUpdateRequest) ProtoMessage()    {}
func (*ProjectDomainAttributesUpdateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{1}
}

func (m *ProjectDomainAttributesUpdateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesUpdateRequest.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesUpdateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesUpdateRequest.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesUpdateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesUpdateRequest.Merge(m, src)
}
func (m *ProjectDomainAttributesUpdateRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesUpdateRequest.Size(m)
}
func (m *ProjectDomainAttributesUpdateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesUpdateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesUpdateRequest proto.InternalMessageInfo

func (m *ProjectDomainAttributesUpdateRequest) GetAttributes() *ProjectDomainAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Purposefully empty, may be populated in the future.
type ProjectDomainAttributesUpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectDomainAttributesUpdateResponse) Reset()         { *m = ProjectDomainAttributesUpdateResponse{} }
func (m *ProjectDomainAttributesUpdateResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesUpdateResponse) ProtoMessage()    {}
func (*ProjectDomainAttributesUpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{2}
}

func (m *ProjectDomainAttributesUpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesUpdateResponse.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesUpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesUpdateResponse.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesUpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesUpdateResponse.Merge(m, src)
}
func (m *ProjectDomainAttributesUpdateResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesUpdateResponse.Size(m)
}
func (m *ProjectDomainAttributesUpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesUpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesUpdateResponse proto.InternalMessageInfo

// Request to get an individual project domain attribute override.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectDomainAttributesGetRequest struct {
	// Unique project id which this set of attributes references.
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id which this set of attributes references.
	// +required
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Which type of matchable attributes to return.
	// +required
	ResourceType         MatchableResource `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.admin.MatchableResource" json:"resource_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProjectDomainAttributesGetRequest) Reset()         { *m = ProjectDomainAttributesGetRequest{} }
func (m *ProjectDomainAttributesGetRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesGetRequest) ProtoMessage()    {}
func (*ProjectDomainAttributesGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{3}
}

func (m *ProjectDomainAttributesGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesGetRequest.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesGetRequest.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesGetRequest.Merge(m, src)
}
func (m *ProjectDomainAttributesGetRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesGetRequest.Size(m)
}
func (m *ProjectDomainAttributesGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesGetRequest proto.InternalMessageInfo

func (m *ProjectDomainAttributesGetRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectDomainAttributesGetRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ProjectDomainAttributesGetRequest) GetResourceType() MatchableResource {
	if m != nil {
		return m.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

// Response to get an individual project domain attribute override.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectDomainAttributesGetResponse struct {
	Attributes           *ProjectDomainAttributes `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *ProjectDomainAttributesGetResponse) Reset()         { *m = ProjectDomainAttributesGetResponse{} }
func (m *ProjectDomainAttributesGetResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesGetResponse) ProtoMessage()    {}
func (*ProjectDomainAttributesGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{4}
}

func (m *ProjectDomainAttributesGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesGetResponse.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesGetResponse.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesGetResponse.Merge(m, src)
}
func (m *ProjectDomainAttributesGetResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesGetResponse.Size(m)
}
func (m *ProjectDomainAttributesGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesGetResponse proto.InternalMessageInfo

func (m *ProjectDomainAttributesGetResponse) GetAttributes() *ProjectDomainAttributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

// Request to delete a set matchable project domain attribute override.
// For more info on matchable attributes, see :ref:`ref_flyteidl.admin.MatchableAttributesConfiguration`
type ProjectDomainAttributesDeleteRequest struct {
	// Unique project id which this set of attributes references.
	// +required
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Unique domain id which this set of attributes references.
	// +required
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// Which type of matchable attributes to delete.
	// +required
	ResourceType         MatchableResource `protobuf:"varint,3,opt,name=resource_type,json=resourceType,proto3,enum=flyteidl.admin.MatchableResource" json:"resource_type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ProjectDomainAttributesDeleteRequest) Reset()         { *m = ProjectDomainAttributesDeleteRequest{} }
func (m *ProjectDomainAttributesDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesDeleteRequest) ProtoMessage()    {}
func (*ProjectDomainAttributesDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{5}
}

func (m *ProjectDomainAttributesDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesDeleteRequest.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesDeleteRequest.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesDeleteRequest.Merge(m, src)
}
func (m *ProjectDomainAttributesDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesDeleteRequest.Size(m)
}
func (m *ProjectDomainAttributesDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesDeleteRequest proto.InternalMessageInfo

func (m *ProjectDomainAttributesDeleteRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ProjectDomainAttributesDeleteRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ProjectDomainAttributesDeleteRequest) GetResourceType() MatchableResource {
	if m != nil {
		return m.ResourceType
	}
	return MatchableResource_TASK_RESOURCE
}

// Purposefully empty, may be populated in the future.
type ProjectDomainAttributesDeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProjectDomainAttributesDeleteResponse) Reset()         { *m = ProjectDomainAttributesDeleteResponse{} }
func (m *ProjectDomainAttributesDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*ProjectDomainAttributesDeleteResponse) ProtoMessage()    {}
func (*ProjectDomainAttributesDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ab0b551a649f05, []int{6}
}

func (m *ProjectDomainAttributesDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProjectDomainAttributesDeleteResponse.Unmarshal(m, b)
}
func (m *ProjectDomainAttributesDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProjectDomainAttributesDeleteResponse.Marshal(b, m, deterministic)
}
func (m *ProjectDomainAttributesDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectDomainAttributesDeleteResponse.Merge(m, src)
}
func (m *ProjectDomainAttributesDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_ProjectDomainAttributesDeleteResponse.Size(m)
}
func (m *ProjectDomainAttributesDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectDomainAttributesDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectDomainAttributesDeleteResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ProjectDomainAttributes)(nil), "flyteidl.admin.ProjectDomainAttributes")
	proto.RegisterType((*ProjectDomainAttributesUpdateRequest)(nil), "flyteidl.admin.ProjectDomainAttributesUpdateRequest")
	proto.RegisterType((*ProjectDomainAttributesUpdateResponse)(nil), "flyteidl.admin.ProjectDomainAttributesUpdateResponse")
	proto.RegisterType((*ProjectDomainAttributesGetRequest)(nil), "flyteidl.admin.ProjectDomainAttributesGetRequest")
	proto.RegisterType((*ProjectDomainAttributesGetResponse)(nil), "flyteidl.admin.ProjectDomainAttributesGetResponse")
	proto.RegisterType((*ProjectDomainAttributesDeleteRequest)(nil), "flyteidl.admin.ProjectDomainAttributesDeleteRequest")
	proto.RegisterType((*ProjectDomainAttributesDeleteResponse)(nil), "flyteidl.admin.ProjectDomainAttributesDeleteResponse")
}

func init() {
	proto.RegisterFile("flyteidl/admin/project_domain_attributes.proto", fileDescriptor_e8ab0b551a649f05)
}

var fileDescriptor_e8ab0b551a649f05 = []byte{
	// 342 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x53, 0x4f, 0x4b, 0xfb, 0x40,
	0x10, 0x65, 0x7f, 0x3f, 0xa8, 0x38, 0x6a, 0x0f, 0x11, 0x34, 0x78, 0x6a, 0x17, 0xa5, 0xbd, 0x98,
	0x40, 0x3d, 0x8a, 0x07, 0xa5, 0xd8, 0x93, 0x20, 0x51, 0x2f, 0x5e, 0x42, 0xfe, 0x8c, 0x69, 0x24,
	0xc9, 0xae, 0xc9, 0xe4, 0xd0, 0x0f, 0x23, 0xf4, 0xa3, 0x8a, 0xbb, 0xd9, 0xfe, 0xd3, 0x28, 0x82,
	0xe0, 0x2d, 0x93, 0xbc, 0x79, 0xef, 0xe5, 0x3d, 0x06, 0x9c, 0xa7, 0x6c, 0x46, 0x98, 0xc6, 0x99,
	0x1b, 0xc4, 0x79, 0x5a, 0xb8, 0xb2, 0x14, 0xcf, 0x18, 0x91, 0x1f, 0x8b, 0x3c, 0x48, 0x0b, 0x3f,
	0x20, 0x2a, 0xd3, 0xb0, 0x26, 0xac, 0x1c, 0x59, 0x0a, 0x12, 0x56, 0xd7, 0xe0, 0x1d, 0x85, 0x3f,
	0x1a, 0x6c, 0xec, 0xe7, 0x01, 0x45, 0xd3, 0x20, 0xcc, 0xd0, 0x2f, 0xb1, 0x12, 0x75, 0x19, 0xa1,
	0x5e, 0xe4, 0x73, 0x06, 0x87, 0xb7, 0x9a, 0x7c, 0xac, 0xb8, 0x2f, 0x17, 0xd4, 0x96, 0x0d, 0x5b,
	0x8d, 0xae, 0xcd, 0x7a, 0x6c, 0xb8, 0xed, 0x99, 0xd1, 0x3a, 0x80, 0x8e, 0x76, 0x62, 0xff, 0x53,
	0x1f, 0x9a, 0xc9, 0xba, 0x83, 0x7d, 0xa5, 0x94, 0x16, 0xc9, 0x8a, 0x47, 0xfb, 0x7f, 0x8f, 0x0d,
	0x77, 0x46, 0xdc, 0x59, 0x37, 0xe9, 0xdc, 0x34, 0xd0, 0xa5, 0xa4, 0x67, 0xe5, 0x1f, 0xde, 0x71,
	0x01, 0xc7, 0x2d, 0x0e, 0x1f, 0x64, 0x1c, 0x10, 0x7a, 0xf8, 0x52, 0x63, 0x45, 0xd6, 0x04, 0x60,
	0x45, 0x93, 0x29, 0xcd, 0xc1, 0xa6, 0x66, 0x0b, 0x93, 0xb7, 0xb2, 0xca, 0x07, 0x70, 0xf2, 0x8d,
	0x60, 0x25, 0x45, 0x51, 0x21, 0x7f, 0x65, 0xd0, 0x6f, 0x41, 0x4e, 0x90, 0x8c, 0xaf, 0x9f, 0xc7,
	0x78, 0x0d, 0x7b, 0xa6, 0x26, 0x9f, 0x66, 0x12, 0x55, 0x80, 0xdd, 0x51, 0xff, 0xd3, 0x00, 0xdf,
	0x5b, 0xf5, 0x1a, 0xb4, 0xb7, 0x6b, 0xf6, 0xee, 0x67, 0x12, 0x79, 0x0e, 0xfc, 0x2b, 0x7b, 0xfa,
	0x2f, 0x7e, 0x2f, 0xb7, 0x39, 0x6b, 0x6d, 0x6a, 0x8c, 0x19, 0x2e, 0x9b, 0xfa, 0xbb, 0x44, 0xda,
	0xab, 0x35, 0x0e, 0x75, 0x28, 0x57, 0x17, 0x8f, 0xe7, 0x49, 0x4a, 0xd3, 0x3a, 0x74, 0x22, 0x91,
	0xbb, 0x4a, 0x45, 0x94, 0x89, 0x7e, 0x70, 0x17, 0xc7, 0x95, 0x60, 0xe1, 0xca, 0xf0, 0x34, 0x11,
	0xee, 0xfa, 0xbd, 0x85, 0x1d, 0x75, 0x5d, 0x67, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x3f,
	0x53, 0x68, 0xc8, 0x03, 0x00, 0x00,
}
